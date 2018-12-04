package vs

import (
	"flag"
	"fmt"
	"errors"
	"log"
	"os"
	"github.com/hashicorp/vault/api"
	"net/url"
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

func (vsConfig *VSConfig) SetVaultAddress(addr string) (err error) {
	u, err := url.ParseRequestURI(addr)
	if err != nil {
		return err
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		msg := fmt.Sprintf("Invalid vaultAddress scheme \"%s\"; it must be http or https", u.Scheme)
		err = errors.New(msg)
		return err
	}
	if u.Hostname() == "" {
		msg := fmt.Sprintf("Invalid vaultAddress hostname" )
		err = errors.New(msg)
		return err
	}
	if u.Port() == "" {
		msg := fmt.Sprintf("Invalid vaultAddress port" )
		err = errors.New(msg)
		return err
	}
	vsConfig.vaultAddress = addr
	return err
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

func InitFlags() *VSConfig {
	var vsConfig VSConfig

	signingRole := flag.String("signingRole", "regular-role", "ssh client signing role")
	mode := flag.String("mode", SSH, "one of: addkey | ssh")
	vaultAddress := flag.String("vaultAddress", "http://localhost:8200", "vault address")
	publicKeyPath := flag.String("publicKeyPath", "", "path to ssh public key file")
	privateKeyPath := flag.String("privateKeyPath", "", "path to ssh private key file")
	sshServerHost := flag.String("sshServerHost", "0.0.0.0", "hostname to connect for ssh seesion")
	username := flag.String("username", "ubuntu", "username for vault auth")
	passwd := flag.String("passwd", "", "password for vault auth (will prompt if empty)")
	sshUsername := flag.String("sshUsername", "ubuntu", "username for ssh session")
	termType := flag.String("termType", "xterm", "terminal type for session session")
	sshServerPort := flag.Int("sshServerPort", 22, "port to connect for ssh session")
	termRows := flag.Int("termRows", 40, "numbr of rows in terminal")
	termCols := flag.Int("termCols", 80, "numbr of columns in terminal")

	flag.Parse()

	// The reason we assign the fields using the Set API is to perform validations there
	vsConfig.SetSigningRole(*signingRole)
	vsConfig.SetMode(*mode)
	err := vsConfig.SetVaultAddress(*vaultAddress)
	if err != nil {
		log.Printf("Bad value for vaultAddress \"%s\"; %v\n", *vaultAddress, err)
		os.Exit(1)
	}
	vsConfig.SetPublicKeyPath(*publicKeyPath)
	vsConfig.SetPrivateKeyPath(*privateKeyPath)
	vsConfig.SetSshServerHost(*sshServerHost)
	vsConfig.SetUsername(*username)
	vsConfig.SetPasswd(*passwd)
	vsConfig.SetSshUsername(*sshUsername)
	vsConfig.SetTermType(*termType)
	vsConfig.SetSshServerPort(*sshServerPort)
	vsConfig.SetTermRows(*termRows)
	vsConfig.SetTermCols(*termCols)

	return &vsConfig
}
