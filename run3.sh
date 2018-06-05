export NODE_ID=3001

count=1
while read line; do
  export WALLET_$count=$line
  count=$(($count + 1))
done < address_3001.txt

export amount1=3
export amount2=1
export amount3=1
export amount4=1

echo "./Backbone Send -from ${WALLET_1} -to ${WALLET_3} -amount ${amount1}"
./Backbone Send -from $WALLET_1 -to $WALLET_2 -amount $amount1

echo "./Backbone Send -from ${WALLET_2} -to ${WALLET_4} -amount ${amount2}"
./Backbone Send -from $WALLET_3 -to $WALLET_4 -amount $amount2

#echo "./Backbone Send -from ${WALLET_2} -to ${WALLET_3} -amount ${amount3}"
#./Backbone Send -from $WALLET_2 -to $WALLET_3 -amount $amount3

#echo "./Backbone Send -from ${WALLET_1} -to ${WALLET_4} -amount ${amount4}"
#./Backbone Send -from $WALLET_1 -to $WALLET_4 -amount $amount4

# EXPECTED RESULT
# a1 = 3, a2 = 2, a3 = 2, a4 = 4
# W1:10 - 3 => W3: 0 + 3
# W2:10 - 2 => W4: 0 + 2
# W2:8 - 2 => W3: 3 + 2
# W1:7 - 4 => W4: 2 + 4
# ----------------------
# W1:3, W2:6, W3:5, W4:6