package main

import (
	"bufio"
	"fmt"
	"github.com/richard-mauri/vaultssh/vs"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	vsapi := vs.Init()

	if vs.GetMode(vsapi) == vs.ADDKEY {
		if vs.GetPrivateKeyPath(vsapi) == "" {
			log.Printf("Please supply a value for the private key path\n")
			os.Exit(2)
		}
		privkey, err := ioutil.ReadFile(vs.GetPrivateKeyPath(vsapi))
		if err != nil {
			msg := fmt.Sprintf("Error reading from private key path; %v\n", err)
			log.Printf(msg)
			os.Exit(3)
		}
		vs.SetPrivateKey(vsapi, string(privkey))

		if vs.GetPublicKeyPath(vsapi) == "" {
			log.Printf("Please supply a value for the public key path\n")
			os.Exit(4)
		}
		pubkey, err := ioutil.ReadFile(vs.GetPublicKeyPath(vsapi))
		if err != nil {
			msg := fmt.Sprintf("Error reading from public key path; %v\n", err)
			log.Printf(msg)
			os.Exit(5)
		}
		vs.SetPublicKey(vsapi, string(pubkey))

		if vs.GetPasswd(vsapi) == "" {
			fmt.Printf("Enter vault userpass password: ")
			scanner := bufio.NewScanner(os.Stdin)
			err := scanner.Err()
			if err != nil {
				msg := fmt.Sprintf("Unable to read pasword; %v\n", err)
				log.Printf(msg)
				os.Exit(6)
			}
			if scanner.Scan() {
				vs.SetPasswd(vsapi, scanner.Text())
			}
		}

		err = vs.AddKeyPair(vsapi)
		if err != nil {
			msg := fmt.Sprintf("Unable to add key pair; %v\n", err)
			log.Printf(msg)
			os.Exit(7)
		}
	} else if vs.GetMode(vsapi) == vs.SSH {
		err := vs.StartSession(vsapi)
		if err != nil {
			msg := fmt.Sprintf("Unable to start session; %v\n", err)
			log.Printf(msg)
			os.Exit(8)
		}
	} else {
		msg := fmt.Sprintf("Illegal mode %s: must be either %s or %s\n", vs.GetMode(vsapi), vs.ADDKEY, vs.SSH)
		log.Printf(msg)
		os.Exit(9)
	}

	os.Exit(0)
}
