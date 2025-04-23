package main

import (
	"bytes"
	"math/big"
)

const targetBits = 24

// block: The block we're trying to mine.
// target: The target value the hash must be less than.
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	//Lsh take a number in this case 1 and move to left making 232(256 - targetBits) zeros after 1 in bytes
	target.Lsh(target, uint(256-targetBits))
	// b is what we want to check and target the difficult
	pow := &ProofOfWork{b, target}
	return pow

}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.TimeStamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}
