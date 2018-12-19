package vs

import (
	"errors"
	"flag"
	"fmt"
	"github.com/hashicorp/vault/api"
	"log"
	"net/url"
	"os"
)

var (
	vsConfig      *VSConfig // Singleton: NewVSConfig "constructor" uses flags which can't be run > 1
	VersionString string
)

type VSConfig struct {
	remoteCommand  string
	localPath      string
	remotePath     string
	signingRole    string
	mode           string
	vaultAddress   string
	publicKeyPath  string
	privateKeyPath string
	sshServerHost  string
	sshServerPort  int
	termType       string
	username       string
	sshUsername    string
	passwd         string
	vaultClient    *api.Client
	vaultToken     string
	privateKey     string
	publicKey      string
	kvVersion      int
}

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
		msg := fmt.Sprintf("Invalid vaultAddress hostname")
		err = errors.New(msg)
		return err
	}
	if u.Port() == "" {
		msg := fmt.Sprintf("Invalid vaultAddress port")
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

func (vsConfig *VSConfig) SetLocalPath(localPath string) {
	vsConfig.localPath = localPath
}

func (vsConfig *VSConfig) GetLocalPath() string {
	return vsConfig.localPath
}

func (vsConfig *VSConfig) GetRemoteCommand() string {
	return vsConfig.remoteCommand
}

func (vsConfig *VSConfig) SetRemoteCommand(remoteCommand string) {
	vsConfig.remoteCommand = remoteCommand
}

func (vsConfig *VSConfig) GetRemotePath() string {
	return vsConfig.remotePath
}

func (vsConfig *VSConfig) SetRemotePath(remotepath string) {
	vsConfig.remotePath = remotepath
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

func (vsConfig *VSConfig) ScpSession() (err error) {
	return vsConfig.ScpSessionAux()
}

func (vsConfig *VSConfig) StartSession() (err error) {
	return vsConfig.StartSessionAux()
}

func (vsConfig *VSConfig) GetKvVersion() int {
	return vsConfig.kvVersion
}

func (vsConfig *VSConfig) SetKvVersion(kvVersion int) {
	vsConfig.kvVersion = kvVersion
}

// Really a singleton constructor as flags can't be parsed more than once
func NewVSConfig() *VSConfig {
	if vsConfig != nil {
		return vsConfig
	}

	vsConfig = new(VSConfig)

	version := flag.Bool("v", false, "print current version and exit")
	signingRole := flag.String("signingRole", "regular-role", "ssh client signing role")
	remoteCommand := flag.String("remoteCommand", "", "remote command to execute")
	mode := flag.String("mode", SSH, "one of: addkey | ssh | scpto | scpfrom")
	localPath := flag.String("localPath", "", "fully qualified path to local file to scp to or from")
	remotePath := flag.String("remotePath", "", "fully qualified path to remote file to scp to or from")
	vaultAddress := flag.String("vaultAddress", "http://localhost:8200", "vault address")
	publicKeyPath := flag.String("publicKeyPath", "", "fully qualified path to ssh public key file")
	privateKeyPath := flag.String("privateKeyPath", "", "fully qualified path to ssh private key file")
	sshServerHost := flag.String("sshServerHost", "0.0.0.0", "hostname to connect for ssh seesion")
	username := flag.String("username", "ubuntu", "username for vault auth")
	passwd := flag.String("passwd", "", "password for vault auth (will prompt if empty)")
	sshUsername := flag.String("sshUsername", "", "username for ssh session (defaults to username value)")
	termType := flag.String("termType", "xterm-256color", "terminal type for session session")
	sshServerPort := flag.Int("sshServerPort", 22, "port to connect for ssh session")
	kvVersion := flag.Int("kvVersion", 1, "vault kv verion (1 or 2)")

	flag.Parse()

	if *version {
		fmt.Println(VersionString)
		os.Exit(0)
	}

	// The reason we assign the fields using the Set API is to perform validations there
	vsConfig.SetSigningRole(*signingRole)
	vsConfig.SetMode(*mode)
	err := vsConfig.SetVaultAddress(*vaultAddress)
	if err != nil {
		log.Printf("Bad value for vaultAddress \"%s\"; %v\n", *vaultAddress, err)
		os.Exit(1)
	}
	vsConfig.SetRemoteCommand(*remoteCommand)
	vsConfig.SetLocalPath(*localPath)
	vsConfig.SetRemotePath(*remotePath)
	vsConfig.SetPublicKeyPath(*publicKeyPath)
	vsConfig.SetPrivateKeyPath(*privateKeyPath)
	vsConfig.SetSshServerHost(*sshServerHost)
	vsConfig.SetUsername(*username)
	vsConfig.SetPasswd(*passwd)
	if *sshUsername == "" {
		vsConfig.SetSshUsername(*username)
	} else {
		vsConfig.SetSshUsername(*sshUsername)
	}
	vsConfig.SetTermType(*termType)
	vsConfig.SetSshServerPort(*sshServerPort)
	vsConfig.SetKvVersion(*kvVersion)

	return vsConfig
}
