#!/bin/bash

# Demonstrate secure interactive ssh session, ssh remote commands, scp to and scp from using vault ca signed ssh keys where all operations are performed in memory.

PATH=$GOPATH/bin:$PATH
SLEEP=1

function echomsg {
	echo -e ${MSG}
}

function echocmdstart {
	echo -e "\033[38;5;50m"
}

function echoend {
	echo -e "\033[0m"
}

function playstep {
	echomsg
	sleep ${SLEEP}
	echocmdstart
	(set -x; ${CMD})
	echoend
	sleep ${SLEEP}
}

MSG="1. addkey: demonstrate how a user injects his ssh key pair into Vault.

Ideally, the vaultssh -addkey command would be invoked from a PC that has network connectivity and firewall acl to Vault.
Notice the Vault password supplied in clear text. This is for the demo;vaultssh will prompt for the password when it is not supplied."

CMD="vaultssh -mode addkey -publicKeyPath ${HOME}/.ssh/id_rsa.pub -privateKeyPath ${HOME}/.ssh/id_rsa -username ubuntu -passwd newpasswd"
playstep

# Punt on getting this sequence working with the demo function (escaping hell)
MSG="2. ssh remote command demo"
echomsg
sleep ${SLEEP}
echocmdstart
(set -x; vaultssh -mode ssh -sshServerPort 6061 -username ubuntu -passwd newpasswd -remoteCommand 'ls -l')
echoend
sleep ${SLEEP}

# Punt on getting this sequence working with the demo function (escaping hell)
MSG="Ok, now let's create a tar file in preparation for the scpto demo that will follow"
echomsg
sleep ${SLEEP}
(set -x; rm ./vendor*.tar.gz >& /dev/null)
(set -x; tar czf ./vendor.tar.gz ./vendor)
(set -x; ls -l ./vendor*.tar.gz)
(set -x; sum ./vendor*.tar.gz)
echoend
sleep ${SLEEP}

MSG="3. scpto: demonstrate how to copy a file to the remote host"
CMD="vaultssh -mode scpto -sshServerPort 6061 -username ubuntu -passwd newpasswd -localPath ./vendor.tar.gz -remotePath /home/ubuntu/"
playstep

MSG="4. scpfrom: demonstrate how to copy a file from the remote host"
CMD="vaultssh -mode scpfrom -sshServerPort 6061 -username ubuntu -passwd newpasswd -localPath ./vendor2.tar.gz -remotePath /home/ubuntu/vendor.tar.gz"
playstep

MSG="Confirm that the original tar file is identical to the copy"
echomsg
sleep ${SLEEP}
(set -x; ls -l ./vendor*.tar.gz; sum ./vendor*.tar.gz)
echoend
sleep ${SLEEP}

MSG="5. ssh interactive session demo"
CMD="vaultssh -mode ssh -sshServerPort 6061 -username ubuntu -passwd newpasswd"
playstep

echo "Done"
