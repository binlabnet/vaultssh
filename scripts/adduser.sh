#!/bin/bash

function usage {
    echo "Usage: adduser <username> <signingrole> [password]"
    echo
    echo "The VAULT_ADDR and VAULT_TOKEN env variables must be in scope"
    echo "The vault token must have root-like policies"
    exit 1
}

function main {

    if ([ "${VAULT_ADDR}" = "" ] || [ "${VAULT_TOKEN}" = "" ] || [ "${USERNAME}" = "" ] || [ "${ROLE}" = "" ])
    then
        usage
    fi

    if [ "${PASSWORD}" = "" ]; then
        echo -n "Password: "
        read -s PASSWORD
        echo
    fi

    vault policy write ssh-${USERNAME}-user - <<EOH1
    path "ssh/sign/${ROLE}" {
        capabilities = ["create","update"]
    }
    path "ssh/sign/${ROLE}" {
        capabilities = ["create","update"]
    }
    path "kv/users/${USERNAME}/*" {
        capabilities = ["create", "read", "update", "delete", "list"]
    }
    path "kv/data/users/${USERNAME}/*" {
        capabilities = ["create", "read", "update", "delete", "list"]
    }
EOH1

    vault write auth/userpass/users/${USERNAME} \
        password=${PASSWORD} \
        policies=ssh-${USERNAME}-user
}

USERNAME=$1
ROLE=$2
PASSWORD=$3

main
