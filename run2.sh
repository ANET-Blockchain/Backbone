#!/bin/bash
#port=${1:?"requires an argument"}
#cp blockchain_${port}.db blockchain_genesis.db

export NODE_ID=3001

read CENTRAL_NODE < address_3000.txt
count=1
while read line; do
  export WALLET_$count=$line
  count=$(($count + 1))
done < address_3001.txt

#./Backbone StartNode

./Backbone GetBalance -address $WALLET_1
./Backbone GetBalance -address $WALLET_2
echo "CENTRAL_NODE"
./Backbone GetBalance -address $CENTRAL_NODE

