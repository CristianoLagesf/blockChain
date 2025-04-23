package main

import (
	"bytes"
	"crypto/sha256"
	"time"

	"strconv"
)

// Timestamp: The time when the block was created.
// Data: The actual information or transactions the block holds.
// PrevBlockHash: The hash of the previous block, linking the blocks together.
// Hash: The unique identifier for the current block, derived from its contents.
type Block struct {
	TimeStamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) SetHask() {
	//convert to a byte slice
	timeStamp := []byte(strconv.FormatInt(b.TimeStamp, 10))
	// concatenated and converted to a byte
	header := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timeStamp}, []byte{})
	// created a hash using the var header
	hash := sha256.Sum256(header)

	b.Hash = hash[:]
}

// Initialize Block: A new Block is initialized with the current timestamp, provided data, and the hash of the previous block.
// Set Hash: The SetHash method computes and sets the block's hash.
// Return Block: The newly created block is returned.
func NewBlock(data string, prevBlockHask []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHask, []byte{}}
	block.SetHask()
	return block

}

func NewBlockChain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
