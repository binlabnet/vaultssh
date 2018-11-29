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

	vsConfig, err := vs.Init()
	if err != nil {
		msg := fmt.Sprintf("Unable to initialize; %v\n", err)
		log.Printf(msg)
		os.Exit(1)
	}

	if vsConfig.Mode == vs.ADDKEY {
		if vsConfig.PrivateKeyPath == "" {
			log.Printf("Please supply a value for the private key path\n")
			os.Exit(2)
		}
		privkey, err := ioutil.ReadFile(vsConfig.PrivateKeyPath)
		if err != nil {
			msg := fmt.Sprintf("Error reading from private key path; %v\n", err)
			log.Printf(msg)
			os.Exit(3)
		}
		vsConfig.State.PrivateKey = string(privkey)

		if vsConfig.PublicKeyPath == "" {
			log.Printf("Please supply a value for the public key path\n")
			os.Exit(4)
		}
		pubkey, err := ioutil.ReadFile(vsConfig.PublicKeyPath)
		if err != nil {
			msg := fmt.Sprintf("Error reading from public key path; %v\n", err)
			log.Printf(msg)
			os.Exit(5)
		}
		vsConfig.State.PublicKey = string(pubkey)

		if vsConfig.Passwd == "" {
			fmt.Printf("Enter vault userpass password: ")
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Err() != nil {
				msg := fmt.Sprintf("Unable to read pasword; %v\n", err)
				log.Printf(msg)
				os.Exit(6)
			}
			if scanner.Scan() {
				vsConfig.Passwd = scanner.Text()
			}
		}

		err = vsConfig.AddKeyPair()
		if err != nil {
			msg := fmt.Sprintf("Unable to add key pair; %v\n", err)
			log.Printf(msg)
			os.Exit(7)
		}
	} else if vsConfig.Mode == vs.SSH {
		err = vsConfig.StartSession()
		if err != nil {
			msg := fmt.Sprintf("Unable to start session; %v\n", err)
			log.Printf(msg)
			os.Exit(8)
		}
	} else {
		msg := fmt.Sprintf("Illegal mode %s: must be either %s or %s\n", vsConfig.Mode, vs.ADDKEY, vs.SSH)
		log.Printf(msg)
		os.Exit(9)
	}

	os.Exit(0)
}
