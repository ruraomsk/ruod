#!/bin/bash
echo 'Compiling for teleofis'
GOOS=linux GOARCH=arm  go build
if [ $? -ne 0 ]; then
	echo 'An error has occurred! Aborting the script execution...'
	exit 1
fi
echo 'Copy ruod to device'
tar -czvf ruod.tar.gz potop
scp ruod.tar.gz root@185.27.195.194:/root 
scp  ruod.tar.gz root@192.168.88.1:/root 
scp ruod.sh root@192.168.88.1:/root 
