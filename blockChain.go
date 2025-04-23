package main

type Blockchain struct {
	blocks []*Block
}

// Retrieve Previous Block: The last block in the chain is retrieved.
// Create New Block: A new block is created using the provided data and the hash of the previous block.
// Append to Blockchain: The new block is appended to the blockchain.
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)

}
func NewGenesisBlock() *Block {
	return NewBlock("genesis block", []byte{})

}
