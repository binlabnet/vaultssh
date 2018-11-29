package vs

import (
	"flag"
	"github.com/hashicorp/vault/api"
)

const (
	ADDKEY = "addkey"
	SSH    = "ssh"
)

type Internal struct {
	VaultClient *api.Client
	VaultToken  string
	PrivateKey  string
	PublicKey   string
}

type VSConfig struct {
	SigningRole    string
	Mode           string
	VaultAddress   string
	State          Internal
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
}

func Init() (vsConfig VSConfig, err error) {
	// flag.StringVar(&vsConfig.SigningRole, "signingRole", "ssh-client-signer/sign/regular-role", "ssh client signing role")
	// flag.StringVar(&vsConfig.SigningRole, "signingRole", "ssh-client-signer/roles/regular-role", "ssh client signing role")
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

	return vsConfig, err
}
