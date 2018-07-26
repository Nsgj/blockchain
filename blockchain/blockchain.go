package blockchain

import "myblockchain/block"

type Blockchain struct {
	blocks []*block.Block
}

func (bc *Blockchain) AddBlock(data string)  {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := block.NewBlock(data,prevBlock.Hash)
	bc.blocks = append(bc.blocks,newBlock)
}

func (bc *Blockchain) GetBlocks() []*block.Block {
	return bc.blocks
}

func NewGenesisBlock() *block.Block {
	return block.NewBlock("Genesis Block",[]byte{})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{
		[]*block.Block{NewGenesisBlock()},
	}
}
