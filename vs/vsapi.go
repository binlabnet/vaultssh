package vs

//go:generate mockgen -source=vsapi.go -package=github.com/richrd-mauri/vaultssh/vs -destination=vsapi_mock.go

import (
	"github.com/hashicorp/vault/api"
)

const (
	ADDKEY  = "addkey"
	SSH     = "ssh"
	SCPTO   = "scpto"
	SCPFROM = "scpfrom"
)

type (
	VsApi interface {
		GetSigningRole() string
		SetSigningRole(string)

		GetMode() string
		SetMode(string)

		GetVaultAddress() string
		SetVaultAddress(string) error

		GetPublicKeyPath() string
		SetPublicKeyPath(string)

		GetPrivateKeyPath() string
		SetPrivateKeyPath(string)

		GetSshServerHost() string
		SetSshServerHost(string)

		GetSshServerPort() int
		SetSshServerPort(int)

		GetTermType() string
		SetTermType(string)

		GetUsername() string
		SetUsername(string)

		GetSshUsername() string
		SetSshUsername(string)

		GetVaultClient() *api.Client
		SetVaultClient(*api.Client)

		GetVaultToken() string
		SetVaultToken(string)

		GetPasswd() string
		SetPasswd(string)

		GetPrivateKey() string
		SetPrivateKey(string)

		GetPublicKey() string
		SetPublicKey(string)

		GetKvVersion() int
		SetKvVersion(int)

		AddKeyPair() (err error)
		StartSession() (err error)
		ScpSession() (err error)
		SignPubKey(pubKey string) (signedCrt string, err error)
		VaultReadSSHKey() (pubKey, privKey string, err error)
		VaultWriteSSHKey() (err error)
		VaultLogin() (err error)
	}
)

func AddKeyPair(vsapi VsApi) (err error) {
	return vsapi.AddKeyPair()
}

func ScpSession(vsapi VsApi) (err error) {
	return vsapi.ScpSession()
}

func StartSession(vsapi VsApi) (err error) {
	return vsapi.StartSession()
}

func SignPubKey(vsapi VsApi, pubKey string) (signedCrt string, err error) {
	return vsapi.SignPubKey(pubKey)
}

func VaultReadSSHKey(vsapi VsApi) (pubKey, privKey string, err error) {
	return vsapi.VaultReadSSHKey()
}

func VaultWriteSSHKey(vsapi VsApi) (err error) {
	return vsapi.VaultWriteSSHKey()
}

func VaultLogin(vsapi VsApi) (err error) {
	return vsapi.VaultLogin()
}

func GetSigningRole(vsapi VsApi) string {
	return vsapi.GetSigningRole()
}

func SetSigningRole(vsapi VsApi, role string) {
	vsapi.SetSigningRole(role)
}

func GetMode(vsapi VsApi) string {
	return vsapi.GetMode()
}

func SetMode(vsapi VsApi, mode string) {
	vsapi.SetMode(mode)
}

func GetVaultAddress(vsapi VsApi) string {
	return vsapi.GetVaultAddress()
}

func SetVaultAddress(vsapi VsApi, addr string) error {
	return vsapi.SetVaultAddress(addr)
}

func GetPublicKeyPath(vsapi VsApi) string {
	return vsapi.GetPublicKeyPath()
}

func SetPublicKeyPath(vsapi VsApi, keypath string) {
	vsapi.SetPublicKeyPath(keypath)
}

func GetPrivateKeyPath(vsapi VsApi) string {
	return vsapi.GetPrivateKeyPath()
}

func SetPrivateKeyPath(vsapi VsApi, keypath string) {
	vsapi.SetPrivateKeyPath(keypath)
}

func GetSshServerHost(vsapi VsApi) string {
	return vsapi.GetSshServerHost()
}

func SetSshServerHost(vsapi VsApi, host string) {
	vsapi.SetSshServerHost(host)
}

func GetSshServerPort(vsapi VsApi) int {
	return vsapi.GetSshServerPort()
}

func SetSshServerPort(vsapi VsApi, port int) {
	vsapi.SetSshServerPort(port)
}

func GetTermType(vsapi VsApi) string {
	return vsapi.GetTermType()
}

func SetTermType(vsapi VsApi, termtype string) {
	vsapi.SetTermType(termtype)
}

func GetSshUsername(vsapi VsApi) string {
	return vsapi.GetSshUsername()
}

func SetSshUsername(vsapi VsApi, username string) {
	vsapi.SetSshUsername(username)
}

func GetUsername(vsapi VsApi) string {
	return vsapi.GetUsername()
}

func SetUsername(vsapi VsApi, username string) {
	vsapi.SetUsername(username)
}

func GetPasswd(vsapi VsApi) string {
	return vsapi.GetPasswd()
}

func SetPasswd(vsapi VsApi, pw string) {
	vsapi.SetPasswd(pw)
}

func GetVaultClient(vsapi VsApi) *api.Client {
	return vsapi.GetVaultClient()
}

func SetVaultClient(vsapi VsApi, client *api.Client) {
	vsapi.SetVaultClient(client)
}

func GetVaultToken(vsapi VsApi) string {
	return vsapi.GetVaultToken()
}

func SetVaultToken(vsapi VsApi, token string) {
	vsapi.SetVaultToken(token)
}

func GetPrivateKey(vsapi VsApi) string {
	return vsapi.GetPrivateKey()
}

func SetPrivateKey(vsapi VsApi, privKey string) {
	vsapi.SetPrivateKey(privKey)
}

func GetPublicKey(vsapi VsApi) string {
	return vsapi.GetPublicKey()
}

func SetPublicKey(vsapi VsApi, pubKey string) {
	vsapi.SetPublicKey(pubKey)
}

func GetKvVersion(vsapi VsApi) int {
	return vsapi.GetKvVersion()
}

func SetKvVersion(vsapi VsApi, kvVersion int) {
	vsapi.SetKvVersion(kvVersion)
}
