# Dependencies
### boltdb
<code>go get github.com/boltdb/bolt/...</code>
### golang/crypto
<code>go get golang.com/x/crypto/... </code>

### How to test
NODE 3000
<code>sh run1.sh</code>

NODE 3001
./Backbone StartNode
downloading transactions
ctrl+c

<code>sh run2.sh</code>

NODE 3002
<code>sh run_miner.sh</code>

NODE 3001
<code>sh run3.sh</code>

View mining a new block at NODE 3002
Vheck the output at NODE 3000

NODE 3001
./Backbone StartNode
downloading the newly mined block
ctrl+c
<code>sh checkwallets.sh</code>
