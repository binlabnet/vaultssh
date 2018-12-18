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

. ./demo-build.sh

# Typically, vaultssh -mode addkey is run on users pc and there is network access to the vault server

$GOPATH/bin/vaultssh -mode addkey -publicKeyPath ~/.ssh/id_rsa.pub -privateKeyPath ~/.ssh/id_rsa -username ubuntu -passwd newpasswd

# Typically, vaultssh -mode ssh is run on bastion host which also has network access to the vault server
# (Note: the system will prompt for the vault userpass credential if not provided)

rm ./vendor*.tar.gz >& /dev/null
tar czf ./vendor.tar.gz ./vendor
ls -l ./vendor*.tar.gz

echo "Copying vendor tar file to remote"
$GOPATH/bin/vaultssh -mode scpto -sshServerPort 6061 -username ubuntu -passwd newpasswd -localPath ./vendor.tar.gz -remotePath /home/ubuntu/

echo "Copying vendor tar file from remote"
$GOPATH/bin/vaultssh -mode scpfrom -sshServerPort 6061 -username ubuntu -passwd newpasswd -localPath ./vendor2.tar.gz -remotePath /home/ubuntu/vendor.tar.gz

ls -l ./vendor*.tar.gz
sum ./vendor*.tar.gz

$GOPATH/bin/vaultssh -mode ssh -sshServerPort 6061 -username ubuntu -passwd newpasswd
