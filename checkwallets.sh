export NODE_ID=3001

read CENTRAL_NODE < address_3000.txt
count=1
while read line; do
  export WALLET_$count=$line
  count=$(($count + 1))
done < address_3001.txt
read MINER_NODE < address_3002.txt

./Backbone GetBalance -address $WALLET_1
./Backbone GetBalance -address $WALLET_2
./Backbone GetBalance -address $WALLET_3
./Backbone GetBalance -address $WALLET_4

./Backbone GetBalance -address $MINER_NODE
./Backbone GetBalance -address $CENTRAL_NODE