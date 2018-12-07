package vs

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"log"
)

func (vsConfig *VSConfig) SignPubKeyAux(pubkey string) (signedCrt string, err error) {

	ssh := vsConfig.GetVaultClient().SSH()

	data := make(map[string]interface{})
	data["public_key"] = pubkey

	data["valid_principals"] = vsConfig.GetSshUsername() // comma-separated list of values
	data["cert_type"] = "user"
	data["extensions"] = map[string]string{
		"permit-X11-forwarding":   "",
		"permit-agent-forwarding": "",
		"permit-port-forwarding":  "",
		"permit-pty":              "",
		"permit-user-rc":          "",
	}

	log.Printf("Calling SignKey with role %s\n", vsConfig.GetSigningRole())
	secret, err := ssh.SignKey(vsConfig.GetSigningRole(), data)
	if err != nil {
		log.Printf("Error signing ssh key; %v\n", err)
		return signedCrt, err
	}

	signedCrt = secret.Data["signed_key"].(string)
	return signedCrt, err
}

func (vsConfig *VSConfig) getUserBase() string {
	if vsConfig.GetKvVersion() == 1 {
		return "secret"
	}
	return "kv"
}

func (vsConfig *VSConfig) VaultReadSSHKey() (pubkey, privkey string, err error) {
	base := vsConfig.getUserBase()
	path := fmt.Sprintf("%s/users/%s/keys/ssh", base, vsConfig.GetUsername())

	s, err := vsConfig.GetVaultClient().Logical().Read(path)
	if err != nil {
		log.Printf("Error reading ssh key pair to path: %s\n", path)
		return pubkey, privkey, err
	}

	pubkey = s.Data["crt"].(string)
	privkey = s.Data["key"].(string)
	return pubkey, privkey, err
}

func (vsConfig *VSConfig) VaultWriteSSHKey() (err error) {
	base := vsConfig.getUserBase()
	path := fmt.Sprintf("%s/users/%s/keys/ssh", base, vsConfig.GetUsername())

	secret := make(map[string]interface{})

	secret["key"] = vsConfig.GetPrivateKey()
	secret["crt"] = vsConfig.GetPublicKey()
	_, err = vsConfig.GetVaultClient().Logical().Write(path, secret)
	if err != nil {
		log.Printf("Error writing key pair to path: %s\n", path)
		return err
	}

	// confirmation
	s, err := vsConfig.GetVaultClient().Logical().Read(path)
	if err != nil {
		log.Printf("Error reading ssh key pair to path: %s\n", path)
		return err
	}
	if s == nil {
		log.Fatal("Vault internal error?; read in ssh key was nil")
	}

	return err
}

func (vsConfig *VSConfig) VaultLogin() (err error) {
	client, err := api.NewClient(&api.Config{
		Address: vsConfig.GetVaultAddress(),
	})
	if err != nil {
		return err
	}

	vsConfig.SetVaultClient(client)

	path := fmt.Sprintf("auth/userpass/login/%s", vsConfig.GetUsername())

	auth, err := vsConfig.GetVaultClient().Logical().Write(path, map[string]interface{}{
		"password": vsConfig.GetPasswd(),
	})
	if err != nil {
		return err
	}

	vt := auth.Auth.ClientToken
	vsConfig.SetVaultToken(vt)
	vsConfig.GetVaultClient().SetToken(vt)

	return err
}
