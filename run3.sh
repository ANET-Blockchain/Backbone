export NODE_ID=3001

count=1
while read line; do
  export WALLET_$count=$line
  count=$(($count + 1))
done < address_3001.txt

echo "./Backbone Send -from ${WALLET_1} -to ${WALLET_3} -amount 2"
./Backbone Send -from $WALLET_1 -to $WALLET_3 -amount 2
echo "./Backbone Send -from ${WALLET_1} -to ${WALLET_4} -amount 3"
./Backbone Send -from $WALLET_1 -to $WALLET_4 -amount 3