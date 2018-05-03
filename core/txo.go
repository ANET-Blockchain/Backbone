package core

import (
	"bytes"
	"encoding/gob"
	"log"
)

type TXOutput struct {
	Value      int
	PubKeyHash []byte
}

type TXOutputs struct {
	Outputs []TXOutput
}

func (out *TXOutput) Lock(address []byte) {
	pubKeyhash := Base58Decode(address)
	pubKeyhash = pubKeyhash[1 : len(pubKeyhash)-4]
	out.PubKeyHash = pubKeyhash
}

func (out *TXOutput) IsLockedWithKey(pubKeyhash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyhash) == 0
}

func NewTXOutput(value int, address string) *TXOutput {
	txo := &TXOutput{value, nil}
	txo.Lock([]byte(address))

	return txo
}

func (outs TXOutputs) Serialize() []byte {
	var buff bytes.Buffer

	encoder := gob.NewEncoder(&buff)
	err := encoder.Encode(outs)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func DeserializeOutputs(data []byte) TXOutputs {
	var outputs TXOutputs

	dec := gob.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(&outputs)
	if err != nil {
		log.Panic(err)
	}

	return outputs
}
