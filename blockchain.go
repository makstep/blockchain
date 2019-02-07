package main

// Blockchain struct
type Blockchain struct {
	blocks []*Block
}

// AddBlock creates a block with a string and adds into a blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)

	bc.blocks = append(bc.blocks, newBlock)
}

// NewBlockchain creates an genesis block and a blockchain
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
