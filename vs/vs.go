package vs

import (
	"flag"
	"github.com/hashicorp/vault/api"
)

type (
	VSConfig struct {
		signingRole    string
		mode           string
		vaultAddress   string
		publicKeyPath  string
		privateKeyPath string
		sshServerHost  string
		sshServerPort  int
		termType       string
		termRows       int
		termCols       int
		username       string
		sshUsername    string
		passwd         string
		vaultClient    *api.Client
		vaultToken     string
		privateKey     string
		publicKey      string
	}
)

func (vsConfig *VSConfig) GetSigningRole() string {
	return vsConfig.signingRole
}

func (vsConfig *VSConfig) SetSigningRole(role string) {
	vsConfig.signingRole = role
}

func (vsConfig *VSConfig) GetMode() string {
	return vsConfig.mode
}

func (vsConfig *VSConfig) SetMode(mode string) {
	vsConfig.mode = mode
}

func (vsConfig *VSConfig) GetVaultAddress() string {
	return vsConfig.vaultAddress
}

func (vsConfig *VSConfig) SetVaultAddress(addr string) {
	vsConfig.vaultAddress = addr
}

func (vsConfig *VSConfig) GetPublicKeyPath() string {
	return vsConfig.publicKeyPath
}

func (vsConfig *VSConfig) SetPublicKeyPath(keypath string) {
	vsConfig.publicKeyPath = keypath
}

func (vsConfig *VSConfig) GetPrivateKeyPath() string {
	return vsConfig.privateKeyPath
}

func (vsConfig *VSConfig) SetPrivateKeyPath(keypath string) {
	vsConfig.privateKeyPath = keypath
}

func (vsConfig *VSConfig) GetSshServerHost() string {
	return vsConfig.sshServerHost
}

func (vsConfig *VSConfig) SetSshServerHost(host string) {
	vsConfig.sshServerHost = host
}

func (vsConfig *VSConfig) GetSshServerPort() int {
	return vsConfig.sshServerPort
}

func (vsConfig *VSConfig) SetSshServerPort(port int) {
	vsConfig.sshServerPort = port
}

func (vsConfig *VSConfig) GetTermType() string {
	return vsConfig.termType
}

func (vsConfig *VSConfig) SetTermType(termtype string) {
	vsConfig.termType = termtype
}

func (vsConfig *VSConfig) GetTermRows() int {
	return vsConfig.termRows
}

func (vsConfig *VSConfig) SetTermRows(rows int) {
	vsConfig.termRows = rows
}

func (vsConfig *VSConfig) GetTermCols() int {
	return vsConfig.termCols
}

func (vsConfig *VSConfig) SetTermCols(cols int) {
	vsConfig.termCols = cols
}

func (vsConfig *VSConfig) GetSshUsername() string {
	return vsConfig.sshUsername
}

func (vsConfig *VSConfig) SetSshUsername(username string) {
	vsConfig.sshUsername = username
}

func (vsConfig *VSConfig) GetUsername() string {
	return vsConfig.username
}

func (vsConfig *VSConfig) SetUsername(username string) {
	vsConfig.username = username
}

func (vsConfig *VSConfig) GetPasswd() string {
	return vsConfig.passwd
}

func (vsConfig *VSConfig) SetPasswd(pw string) {
	vsConfig.passwd = pw
}

func (vsConfig *VSConfig) GetVaultClient() *api.Client {
	return vsConfig.vaultClient
}

func (vsConfig *VSConfig) SetVaultClient(client *api.Client) {
	vsConfig.vaultClient = client
}

func (vsConfig *VSConfig) GetVaultToken() string {
	return vsConfig.vaultToken
}

func (vsConfig *VSConfig) SetVaultToken(token string) {
	vsConfig.vaultToken = token
}

func (vsConfig *VSConfig) GetPrivateKey() string {
	return vsConfig.privateKey
}

func (vsConfig *VSConfig) SetPrivateKey(privKey string) {
	vsConfig.privateKey = privKey
}

func (vsConfig *VSConfig) GetPublicKey() string {
	return vsConfig.publicKey
}

func (vsConfig *VSConfig) SetPublicKey(pubKey string) {
	vsConfig.publicKey = pubKey
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

func Init() *VSConfig {
	var vsConfig VSConfig

	flag.StringVar(&vsConfig.signingRole, "signingRole", "regular-role", "ssh client signing role")
	flag.StringVar(&vsConfig.mode, "mode", ADDKEY, "one of: addkey | ssh")
	flag.StringVar(&vsConfig.vaultAddress, "vaultAddress", "http://localhost:8200", "vault address")
	flag.StringVar(&vsConfig.publicKeyPath, "publicKeyPath", "", "path to ssh public key file")
	flag.StringVar(&vsConfig.privateKeyPath, "privateKeyPath", "", "path to ssh private key file")
	flag.StringVar(&vsConfig.sshServerHost, "sshServerHost", "0.0.0.0", "hostname to connect for ssh seesion")
	flag.StringVar(&vsConfig.username, "username", "ubuntu", "username for vault auth")
	flag.StringVar(&vsConfig.passwd, "passwd", "", "password for vault auth (will prompt if empty)")
	flag.StringVar(&vsConfig.sshUsername, "sshUsername", "ubuntu", "username for ssh session")
	flag.StringVar(&vsConfig.termType, "termType", "xterm", "terminal type for session session")
	flag.IntVar(&vsConfig.sshServerPort, "sshServerPort", 22, "port to connect for ssh session")
	flag.IntVar(&vsConfig.termRows, "termRows", 40, "numbr of rows in terminal")
	flag.IntVar(&vsConfig.termCols, "termCols", 80, "numbr of columns in terminal")

	flag.Parse()

	return &vsConfig
}
