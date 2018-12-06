A Vault client supporting interactive ssh sessions using signed ssh certificates done all in-memory.

## vaultssh
 It is motivated by the need to be able to ssh from a shared bastion jump host and we :
 * Cannot store keys on disk
 * Don't want to remember or manage a complicated passphrase
 
 For this system to work, the ssh servers must be configured to trust the vault ca, which signs the users ssh cert.
 
 The demo script captures all the steps so use that as reference.

## Operation
There are two modes of operation:
1. addkey : user injects his ssh keypair into Vault *once* for subsequent ssh access
1. ssh: the user uses this mode to log into vault, sign his key and start an interactive ssh session

## Project Setup
* The build and demo dependencies include: Go (go1.11.2), Docker (18.09.0), Git (2.18.0), dep (v0.5.0), vault (0.11.5)
* Run "dep ensure" to populate vendor dependencies
* Of course, make sure you have GOPATH defined and use "go install" to build.

## Demo
The demo.sh starts a vault dev server, an ssh server, configures both, stores ssh keys and uses them to start a session.
Enter "exit" to abort the ssh terminal session.
Or, watch this:
[![asciicast](https://asciinema.org/a/hGrgVLfcCWYAOo4J0G92QvcWz.svg)](https://asciinema.org/a/hGrgVLfcCWYAOo4J0G92QvcWz)

## Usage
```
Usage of vaultssh:
  -mode string
    	one of: addkey | ssh (default "addkey")
  -passwd string
    	password for vault auth (will prompt if empty)
  -privateKeyPath string
    	path to ssh private key file
  -publicKeyPath string
    	path to ssh public key file
  -signingRole string
    	ssh client signing role (default "regular-role")
  -sshServerHost string
    	hostname to connect for ssh seesion (default "0.0.0.0")
  -sshServerPort int
    	port to connect for ssh session (default 22)
  -sshUsername string
    	username for ssh session (default "ubuntu")
  -termCols int
    	numbr of columns in terminal (default 80)
  -termRows int
    	numbr of rows in terminal (default 40)
  -termType string
    	terminal type for session session (default "xterm")
  -username string
    	username for vault auth (default "ubuntu")
  -vaultAddress string
    	vault address (default "http://localhost:8200")
```

## License
Mozilla Public License, version 2.0

## TODO
* Test cases
* More configurability
* Support additional Vault user auth backends besides userpass

### GitHubPages: https://richard-mauri.github.io/vaultssh/

### References links
* https://golang.github.io/dep/
* https://godoc.org/github.com/hashicorp/vault/api
* https://www.vaultproject.io/docs/secrets/ssh/signed-ssh-certificates.html
* https://github.com/hashicorp/vault/blob/master/command/ssh.go
* https://godoc.org/golang.org/x/crypto/ssh#example-PublicKeys
* https://dev.to/dmigwi/mocking-methods-in-go-5fg
* https://stackoverflow.com/questions/19167970/mock-functions-in-go
* https://caitiem.com/2016/08/18/a-quick-guide-to-testing-in-golang/
* https://github.com/golang/mock
* https://github.com/cweill/gotests
