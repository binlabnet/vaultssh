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
# Next, it builds snd launches the vaultssh client that starts an interactive cert based ssh session.

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

# The following steps should be done by an admin user 
vault secrets enable -version=1 kv
#vault secrets enable -path=ssh-client-signer ssh
vault secrets enable -path=ssh ssh
vault auth enable userpass

vault policy write ssh-ubuntu-user - <<"EOH1"
path "ssh/sign/regular-role" {
    capabilities = ["create","update"]
}
path "ssh/sign/regular-role" {
    capabilities = ["create","update"]
}
path "kv/users/ubuntu/*" {
    capabilities = ["create", "read", "update", "delete", "list"]
}
path "kv/data/users/ubuntu/*" {
    capabilities = ["create", "read", "update", "delete", "list"]
}
EOH1

vault write auth/userpass/users/ubuntu \
    password=newpasswd \
    policies=ssh-ubuntu-user

#vault write -field=public_key ssh-client-signer/config/ca generate_signing_key=true
vault write -field=public_key ssh/config/ca generate_signing_key=true

#vault read -field=public_key ssh-client-signer/config/ca > ./trusted-user-ca-keys.pem
vault read -field=public_key ssh/config/ca > ./trusted-user-ca-keys.pem

docker build --tag sshtest .

docker run -p 6061:22 -d --name sshtest sshtest

#docker exec -it sshtest /bin/bash

#vault write ssh-client-signer/roles/regular-role -<<"EOH"
vault write ssh/roles/regular-role -<<"EOH"
{
  "allow_user_certificates": true,
  "allowed_users": "*",
  "default_extensions": [
    {
      "permit-pty": ""
    }
  ],
  "key_type": "ca",
  "default_user": "ubuntu",
  "ttl": "30m0s"
}
EOH

# The following needs to be replaced by new two-phase vaultssh machanism that will:
# 1 addkey
#   a) PREREQUISITE: vault login the user (userpass will be supported first)
#   b) write the user's ssh key-pair to kv/users/<username>/keys
# 2 ssh
#   a) PREREQUISITE: vault login the user (userpass will be supported first)
#   b) vault read the keys, sign the crt it and keep the result only in memory
#   c) start the ssh interactive session using signed crt (the ssh server must have sshd preconfigured  to trust the ca crt)
# 

go install

# In a real scenario the user wishing to ssh would perform addkey once and the ssh often (not back to back like here)
# Typically, vaultssh -mode addkey is run on users pc and there is network access to the vault server
$ vaultssh -mode ssh would be run on the bastion host, which also has access to the vault server
$GOPATH/bin/vaultssh -mode addkey -publicKeyPath ~/.ssh/id_rsa.pub -privateKeyPath ~/.ssh/id_rsa -username ubuntu -passwd newpasswd
$GOPATH/bin/vaultssh -mode ssh -sshServerPort 6061 -username ubuntu -passwd newpasswd
