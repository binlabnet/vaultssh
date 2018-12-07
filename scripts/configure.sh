#!/bin/bash

function usage {
    echo "Usage: configure <default user>"
    echo
    echo "The VAULT_ADDR and VAULT_TOKEN env variables must be in scope"
    echo "The vault token must have root-like policies"
    exit 1
}

function main {

    if ([ "${VAULT_ADDR}" = "" ] || [ "${VAULT_TOKEN}" = "" ] || [ "${DEFAULT_USER}" = "" ] )
    then
        usage
    fi

    vault secrets enable -version=1 kv
    vault secrets enable -path=ssh ssh
    vault auth enable userpass

    vault write -field=public_key ssh/config/ca generate_signing_key=true

    vault read -field=public_key ssh/config/ca > ./trusted-user-ca-keys.pem

    vault write ssh/roles/regular-role -<<EOH
    {
      "allow_user_certificates": true,
      "allowed_users": "*",
      "default_extensions": [
        {
          "permit-pty": ""
        }
      ],
      "key_type": "ca",
      "default_user": "${DEFAULT_USER}",
      "ttl": "30m0s"
    }
EOH
}

DEFAULT_USER=$1
main
