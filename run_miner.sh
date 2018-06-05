export NODE_ID=3002

read MINER_NODE < address_3002.txt
./Backbone StartNode -miner ${MINER_NODE}