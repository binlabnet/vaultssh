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
# Next, it builds and launches various vaultssh client nstances to demonstrate the ssh and scp features.

. ./demo-build.sh

. ./play.sh
