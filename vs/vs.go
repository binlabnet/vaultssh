package vs

import (
	"flag"
	"github.com/hashicorp/vault/api"
)

type (
        VSConfig struct {
                SigningRole    string
                Mode           string
                VaultAddress   string
                PublicKeyPath  string
                PrivateKeyPath string
                SshServerHost  string
                SshServerPort  int
                TermType       string
                TermRows       int
                TermCols       int
                Username       string
                SshUsername    string
                Passwd         string
                VaultClient    *api.Client
                VaultToken     string
                PrivateKey     string
                PublicKey      string
        }
)

func (vsConfig *VSConfig) GetSigningRole() (string) {
	return vsConfig.SigningRole
}

func (vsConfig *VSConfig) SetSigningRole(role string) {
	vsConfig.SigningRole = role
}

func (vsConfig *VSConfig) GetMode() (string) {
	return vsConfig.Mode
}

func (vsConfig *VSConfig) SetMode(mode string) {
	vsConfig.Mode = mode
}

func (vsConfig *VSConfig) GetVaultAddress() (string) {
	return vsConfig.VaultAddress
}

func (vsConfig *VSConfig) SetVaultAddress(addr string) {
	vsConfig.VaultAddress = addr
}

func (vsConfig *VSConfig) GetPublicKeyPath() (string) {
	return vsConfig.PublicKeyPath
}

func (vsConfig *VSConfig) SetPublicKeyPath(keypath string) {
	vsConfig.PublicKeyPath = keypath
}

func (vsConfig *VSConfig) GetPrivateKeyPath() (string) {
	return vsConfig.PrivateKeyPath
}

func (vsConfig *VSConfig) SetPrivateKeyPath(keypath string) {
	vsConfig.PrivateKeyPath = keypath
}

func (vsConfig *VSConfig) GetSshServerHost() (string) {
	return vsConfig.SshServerHost
}

func (vsConfig *VSConfig) SetSshServerHost(host string) {
	vsConfig.SshServerHost = host
}

func (vsConfig *VSConfig) GetSshServerPort() (int) {
	return vsConfig.SshServerPort
}

func (vsConfig *VSConfig) SetSshServerPort(port int) {
	vsConfig.SshServerPort = port
}

func (vsConfig *VSConfig) GetTermType() (string) {
	return vsConfig.TermType
}

func (vsConfig *VSConfig) SetTermType(termtype string) {
	vsConfig.TermType = termtype
}

func (vsConfig *VSConfig) GetTermRows() (int) {
	return vsConfig.TermRows
}

func (vsConfig *VSConfig) SetTermRows(rows int) {
	vsConfig.TermRows = rows
}

func (vsConfig *VSConfig) GetTermCols() (int) {
	return vsConfig.TermCols
}

func (vsConfig *VSConfig) SetTermCols(cols int) {
	vsConfig.TermCols = cols
}

func (vsConfig *VSConfig) GetSshUsername() (string) {
	return vsConfig.SshUsername
}

func (vsConfig *VSConfig) SetSshUsername(username string) {
	vsConfig.SshUsername = username
}

func (vsConfig *VSConfig) GetUsername() (string) {
	return vsConfig.Username
}

func (vsConfig *VSConfig) SetUsername(username string) {
	vsConfig.Username = username
}

func (vsConfig *VSConfig) GetPasswd() (string) {
	return vsConfig.Passwd
}

func (vsConfig *VSConfig) SetPasswd(pw string) {
	vsConfig.Passwd = pw
}

func (vsConfig *VSConfig) GetVaultClient() (*api.Client) {
	return vsConfig.VaultClient
}

func (vsConfig *VSConfig) SetVaultClient(client *api.Client) {
	vsConfig.VaultClient = client
}

func (vsConfig *VSConfig) GetVaultToken() (string) {
	return vsConfig.VaultToken
}

func (vsConfig *VSConfig) SetVaultToken( token string) {
	vsConfig.VaultToken = token
}

func (vsConfig *VSConfig) GetPrivateKey() (string) {
	return vsConfig.PrivateKey
}

func (vsConfig *VSConfig) SetPrivateKey(privKey string) {
	vsConfig.PrivateKey = privKey
}

func (vsConfig *VSConfig) GetPublicKey() (string) {
	return vsConfig.PublicKey
}

func (vsConfig *VSConfig) SetPublicKey( pubKey string) {
	vsConfig.PublicKey = pubKey
}

func (vsConfig *VSConfig) AddKeyPair() (err error) {
	return vsConfig.AddKeyPairAux()
}

func (vsConfig *VSConfig) SignPubKey(pubKey string) (signedCrt string, err error) {
	return vsConfig.SignPubKeyAux(pubKey)
}

func (vsConfig *VSConfig) StartSession() (err error) {
	return vsConfig.StartSessionAux()
}

func Init() (*VSConfig) {
	var vsConfig VSConfig

	flag.StringVar(&vsConfig.SigningRole, "signingRole", "regular-role", "ssh client signing role")
	flag.StringVar(&vsConfig.Mode, "mode", ADDKEY, "one of: addkey | ssh")
	flag.StringVar(&vsConfig.VaultAddress, "vaultAddress", "http://localhost:8200", "vault address")
	flag.StringVar(&vsConfig.PublicKeyPath, "publicKeyPath", "", "path to ssh public key file")
	flag.StringVar(&vsConfig.PrivateKeyPath, "privateKeyPath", "", "path to ssh private key file")
	flag.StringVar(&vsConfig.SshServerHost, "sshServerHost", "0.0.0.0", "hostname to connect for ssh seesion")
	flag.StringVar(&vsConfig.Username, "username", "ubuntu", "username for vault auth")
	flag.StringVar(&vsConfig.Passwd, "passwd", "", "password for vault auth (will prompt if empty)")
	flag.StringVar(&vsConfig.SshUsername, "sshUsername", "ubuntu", "username for ssh session")
	flag.StringVar(&vsConfig.TermType, "termType", "xterm", "terminal type for session session")
	flag.IntVar(&vsConfig.SshServerPort, "sshServerPort", 22, "port to connect for ssh session")
	flag.IntVar(&vsConfig.TermRows, "termRows", 40, "numbr of rows in terminal")
	flag.IntVar(&vsConfig.TermCols, "termCols", 80, "numbr of columns in terminal")

	flag.Parse()

	return &vsConfig
}
