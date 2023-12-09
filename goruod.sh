#!/bin/ash
mkdir /tmp/rura
tar -xzvf ruod.tar.gz -C "/tmp/rura"
cp config.json /tmp/rura/
cd /tmp/rura
while true
do
    ./ruod > /dev/null 2>/dev/null
    cp config.json /root
done 
