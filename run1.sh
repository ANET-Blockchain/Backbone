#!/bin/bash

rm *.dat
rm *.db
rm address*.*
rm ./Backbone
go build

export NODE_ID=3000
ADDR=$(./Backbone CreateWallet | awk '{print $5}')

./Backbone CreateBlockchain -address $ADDR

ADDR2=$(./Backbone CreateWallet | awk '{print $5}')
echo "------------------------------"
echo "NODE_ID = 3000"
echo "CENTRAL_NODE : "  $ADDR

echo $ADDR > address_3000.txt
#echo $ADDR2 >> address_3000.txt

cp blockchain_3000.db blockchain_genesis.db
cp blockchain_genesis.db blockchain_3001.db
cp blockchain_genesis.db blockchain_3002.db

export NODE_ID=3001
ADDR=$(./Backbone CreateWallet | awk '{print $5}')
ADDR2=$(./Backbone CreateWallet | awk '{print $5}')
ADDR3=$(./Backbone CreateWallet | awk '{print $5}')
ADDR4=$(./Backbone CreateWallet | awk '{print $5}')
echo "------------------------------"
echo "NODE_ID = 3001"
echo "WALLET_1 : "  $ADDR
echo "WALLET_2 : "  $ADDR2
echo "WALLET_3 : "  $ADDR3
echo "WALLET_4 : "  $ADDR4

echo $ADDR > address_3001.txt
echo $ADDR2 >> address_3001.txt
echo $ADDR3 >> address_3001.txt
echo $ADDR4 >> address_3001.txt

export NODE_ID=3002
ADDR=$(./Backbone CreateWallet | awk '{print $5}')
echo "------------------------------"
echo "NODE_ID = 3002"
echo "MINER_NODE : "  $ADDR
echo $ADDR > address_3002.txt
#echo $ADDR2 >> address_3002.txt

export NODE_ID=3000

read CENTRAL_NODE < address_3000.txt

count=1
while read line; do
  export WALLET_$count=$line
  count=$(($count + 1))
done < address_3001.txt
read MINER_NODE < address_3002.txt

export amount1=10
export amount2=10

echo "./Backbone Send -from ${CENTRAL_NODE} -to ${WALLET_1} -amount ${amount1} -mine"
./Backbone Send -from $CENTRAL_NODE -to $WALLET_1 -amount $amount1 -mine
echo "./Backbone Send -from ${CENTRAL_NODE} -to ${WALLET_2} -amount ${amount2} -mine"
./Backbone Send -from $CENTRAL_NODE -to $WALLET_2 -amount $amount2 -mine

./Backbone StartNode