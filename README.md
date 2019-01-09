A Vault client supporting ssh remote commands, interactive sessions and scp file transfer using signed ssh certificates done all in-memory.

## vaultssh
 It is motivated by the need to be able to ssh from a shared bastion jump host and we :
 * Cannot store keys on disk (not secure)
 * Cannot use ssh agent forwarding (not secure)
 * Don't want to remember or manage a complicated passphrase
 
For this system to work, the ssh servers must be configured to trust the vault ca, which signs the users ssh cert.
 
The demo script captures all the steps so use that as reference.

## Operation
Here are the modes of operation:
1. addkey : user injects his ssh keypair into Vault *once* for subsequent ssh access
1. ssh: the user uses this mode to log into vault, sign his key and start an interactive ssh session
1. scpto: the user uses this mode to log into vault, sign his key and transfer files to a remote system
1. scpfrom: the user uses this mode to log into vault, sign his key and transfer files from a remote system

## Project Setup
* This uses "go mod" for build dependency management (go.11.4 is current and known to work ok)
* Makefile is used to drive the go build commands
* The GOPATH env variable must be defined so the installed vaultssh is installed under GOPATH/bin.
* The runtime demo requires dependencies including (at least these versions;others may work): Go (go1.11.4), Docker (18.09.0), Git (2.18.0), vault (>= 0.11.5)
* This project internally uses travis ci and goreleaser for CI builds and releases respectively.

## Demo
There is a demo.sh that invokes demo-build.sh and play.sh to build and use respectively.
The following video is an example of the play.sh session
[![asciicast](https://asciinema.org/a/217635.svg)](https://asciinema.org/a/217635)

## Usage
There are a couple bash scripts under scripts/
* configure.sh can be used to configure vault to enable userpass and signing. Run it after vault init.
* adduser.sh can be used to create a vault userpass account and configure policies to be able to sign and ssh.

## Example addkey usage (each user does this once; his vault password is prompted for)
* vaultssh -mode addkey -publicKeyPath ~/.ssh/id_rsa.pub -privateKeyPath ~/.ssh/id_rsa -username ubuntu

## Example ssh interactive usage:
* vaultssh -mode ssh -username ubuntu -sshServerHost infra1.foo.com

## Example scp to usage:
* vaultssh -mode scpto -username ubuntu -localPath /tmp/source.txt  -remotePath /home/ubuntu/source.txt -sshServerHost infra1.foo.com

## Example scp from usage:
* vaultssh -mode scpfrom -username ubuntu -localPath /tmp/source2.txt  -remotePath /home/ubuntu/source.txt -sshServerHost infra1.foo.com

```
Usage of vaultssh:
  -kvVersion int
    	vault kv verion (1 or 2) (default 1)
  -localPath string
    	fully qualified path to local file to scp to or from
  -mode string
    	one of: addkey | ssh | scpto | scpfrom (default "ssh")
  -passwd string
    	password for vault auth (will prompt if empty)
  -privateKeyPath string
    	fully qualified path to ssh private key file
  -publicKeyPath string
    	fully qualified path to ssh public key file
  -remoteCommand string
    	remote command to execute
  -remotePath string
    	fully qualified path to remote file to scp to or from
  -signingRole string
    	ssh client signing role (default "regular-role")
  -sshServerHost string
    	hostname to connect for ssh seesion (default "0.0.0.0")
  -sshServerPort int
    	port to connect for ssh session (default 22)
  -sshUsername string
    	username for ssh session (defaults to username value)
  -termType string
    	terminal type for session session (default "xterm-256color")
  -username string
    	username for vault auth (default "ubuntu")
  -v	print current version and exit
  -vaultAddress string
    	vault address (default "http://localhost:8200")
```

## License
Mozilla Public License, version 2.0

## TODO
* Test cases
* More configurability - ssh and scp options
* Support additional Vault user auth backends besides userpass
* cli more like ssh, scp

### GitHubPages: https://richard-mauri.github.io/vaultssh/

### CI Build: https://travis-ci.org/richard-mauri/vaultssh

### Releasing

### References links
* https://github.com/golang/go/wiki/Modules
* https://godoc.org/github.com/hashicorp/vault/api
* https://www.vaultproject.io/docs/secrets/ssh/signed-ssh-certificates.html
* https://github.com/hashicorp/vault/blob/master/command/ssh.go
* https://godoc.org/golang.org/x/crypto/ssh#example-PublicKeys
* https://dev.to/dmigwi/mocking-methods-in-go-5fg
* https://stackoverflow.com/questions/19167970/mock-functions-in-go
* https://caitiem.com/2016/08/18/a-quick-guide-to-testing-in-golang/
* https://github.com/golang/mock
* https://github.com/cweill/gotests
* https://goreleaser.com/
* https://asciinema.org/
* https://docstore.mik.ua/orelly/networking_2ndEd/ssh/index.htm
* https://docstore.mik.ua/orelly/networking_2ndEd/ssh/ch03_08.htm
