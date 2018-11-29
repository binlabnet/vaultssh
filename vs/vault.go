package vs

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"log"
)

func (vsConfig *VSConfig) SignPubKey(pubkey string) (signedCrt string, err error) {

	ssh := vsConfig.State.VaultClient.SSH()

	data := make(map[string]interface{})
	data["public_key"] = pubkey

	data["valid_principals"] = vsConfig.SshUsername // comma-separated list of values
	data["cert_type"] = "user"
	data["extensions"] = map[string]string{
		"permit-X11-forwarding":   "",
		"permit-agent-forwarding": "",
		"permit-port-forwarding":  "",
		"permit-pty":              "",
		"permit-user-rc":          "",
	}

	log.Printf("Calling SignKey with role %s\n", vsConfig.SigningRole)
	secret, err := ssh.SignKey(vsConfig.SigningRole, data)
	if err != nil {
		log.Printf("Error signing ssh key; %v\n", err)
		return signedCrt, err
	}

	signedCrt = secret.Data["signed_key"].(string)
	return signedCrt, err
}

func (vsConfig *VSConfig) VaultReadSSHKey() (pubkey, privkey string, err error) {
	path := fmt.Sprintf("kv/users/%s/keys/ssh", vsConfig.Username)

	s, err := vsConfig.State.VaultClient.Logical().Read(path)
	if err != nil {
		log.Printf("Error reading ssh key pair to path: %s\n", path)
		return pubkey, privkey, err
	}

	pubkey = s.Data["crt"].(string)
	privkey = s.Data["key"].(string)
	return pubkey, privkey, err
}

func (vsConfig *VSConfig) VaultWriteSSHKey() (err error) {
	path := fmt.Sprintf("kv/users/%s/keys/ssh", vsConfig.Username)

	secret := make(map[string]interface{})

	secret["key"] = vsConfig.State.PrivateKey
	secret["crt"] = vsConfig.State.PublicKey
	_, err = vsConfig.State.VaultClient.Logical().Write(path, secret)
	if err != nil {
		log.Printf("Error writing key pair to path: %s\n", path)
		return err
	}

	// confirmation
	s, err := vsConfig.State.VaultClient.Logical().Read(path)
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
		Address: vsConfig.VaultAddress,
	})
	if err != nil {
		return err
	}

	vsConfig.State.VaultClient = client

	path := fmt.Sprintf("auth/userpass/login/%s", vsConfig.Username)

	auth, err := vsConfig.State.VaultClient.Logical().Write(path, map[string]interface{}{
		"password": vsConfig.Passwd,
	})
	if err != nil {
		return err
	}

	vsConfig.State.VaultToken = auth.Auth.ClientToken
	vsConfig.State.VaultClient.SetToken(vsConfig.State.VaultToken)

	return err
}
