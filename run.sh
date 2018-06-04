#!/bin/bash
port=${1:?"requires an argument"}

rm *.dat
rm *.db

export NODE_ID=$port
echo $NODE_ID
ADDR=$(./Backbone CreateWallet | awk '{print $5}')

./Backbone CreateBlockchain -address $ADDR

ADDR2=$(./Backbone CreateWallet | awk '{print $5}')
echo "addr 1 : "  $ADDR
echo "addr 2 : "  $ADDR2

echo $ADDR > address.txt
echo $ADDR2 >> address.txt

cp blockchain_${port}.db blockchain_genesis.db
