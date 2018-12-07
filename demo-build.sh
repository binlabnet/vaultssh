#!/bin/bash

# Demonstrate secure interactive ssh session using vault signed certs where all operations are performed in memory.
# Vault, out of the box, shells out to ssh and ssh key must be present on disk, which is not good when you're
# forced to ssh to the destination host from a shared bastion jump host.
#
# Note, using an ssh passphrase is marginally better (passphrase can be guessed and forgotten)
#
# This self-contained demo script starts a vault server in dev mode and
# configures it as a ca for ssh signing and sets up a userpass account with a personal secret area to store the ssh keys.
# It then starts an ssh server which is configured to trust the vault ca.
# Next, it builds and launches the vaultssh client that starts an interactive cert based ssh session.

if [ "${GOPATH}" = "" ]; then
    echo "Please set the GOPATH environment variable; can't build vaultssh"
    exit 1
fi

export VAULT_ADDR=http://localhost:8200
export VAULT_TOKEN=roottoken

pkill vault >& /dev/null
docker kill sshtest >& /dev/null
docker rm sshtest >& /dev/null

vault server -dev -dev-root-token-id=roottoken > vaultinit.out &

sleep 1

scripts/configure.sh ubuntu

scripts/adduser.sh ubuntu regular-role newpasswd

docker build --tag sshtest .

docker run -p 6061:22 -d --name sshtest sshtest

BUILDTIME=$(date +%m-%d-%Y-%H:%M)

go install -ldflags "-X github.com/richard-mauri/vaultssh/vs.VersionString=v0.1.0-${BUILDTIME}"

go test github.com/richard-mauri/vaultssh/vs
