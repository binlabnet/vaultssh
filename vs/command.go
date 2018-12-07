package vs

import (
	"fmt"
	"io/ioutil"
	"log"
	"syscall"
	"golang.org/x/crypto/ssh/terminal"
)

func Initialize() (vsapi VsApi) {
	return NewVSConfig()
}

func Addkey(vsapi VsApi) (exitcode int) {
	exitcode = 0
	if GetPrivateKeyPath(vsapi) == "" {
		log.Printf("Please supply a value for the private key path\n")
		exitcode = 2
		return exitcode
	}
	privkey, err := ioutil.ReadFile(GetPrivateKeyPath(vsapi))
	if err != nil {
		msg := fmt.Sprintf("Error reading from private key path; %v\n", err)
		log.Printf(msg)
		exitcode = 3
		return exitcode
	}
	SetPrivateKey(vsapi, string(privkey))

	if GetPublicKeyPath(vsapi) == "" {
		log.Printf("Please supply a value for the public key path\n")
		exitcode = 4
		return exitcode
	}
	pubkey, err := ioutil.ReadFile(GetPublicKeyPath(vsapi))
	if err != nil {
		msg := fmt.Sprintf("Error reading from public key path; %v\n", err)
		log.Printf(msg)
		exitcode = 5
		return exitcode
	}
	SetPublicKey(vsapi, string(pubkey))

	if GetPasswd(vsapi) == "" {
		fmt.Printf("Enter vault userpass password: ")
		bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			msg := fmt.Sprintf("Unable to read pasword; %v\n", err)
			log.Printf(msg)
			exitcode = 6
			return exitcode
		}
		SetPasswd(vsapi, string(bytePassword))
	}

	err = AddKeyPair(vsapi)
	if err != nil {
		msg := fmt.Sprintf("Unable to add key pair; %v\n", err)
		log.Printf(msg)
		exitcode = 7
		return exitcode
	}
	return exitcode
}

func Ssh(vsapi VsApi) (exitcode int) {
	if GetPasswd(vsapi) == "" {
		fmt.Printf("Enter vault userpass password: ")
		bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			msg := fmt.Sprintf("Unable to read pasword; %v\n", err)
			log.Printf(msg)
			exitcode = 6
			return exitcode
		}
		SetPasswd(vsapi, string(bytePassword))
	}

	exitcode = 0
	err := StartSession(vsapi)
	if err != nil {
		msg := fmt.Sprintf("Unable to start session; %v\n", err)
		log.Printf(msg)
		exitcode = 8
	}
	return exitcode
}
