package blockchain

import (
	"blockchain/block"
	"github.com/boltdb/bolt"
	"errors"
)

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

const dbFile = "blotdb"
const blocksBucket = "blocksBucket"

func NewBlockchain() *Blockchain {

	var tip []byte
	db,err := bolt.Open(dbFile,0600,nil)

	if err != nil{

	}

	err = db.Update(func(tx *bolt.Tx) error {
		b :=  tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesis := NewGenesisBlock()
			b,err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				return errors.New("cannot create bucket")
			}

			err = b.Put(genesis.Hash,genesis.Serialize())

			err = b.Put([]byte("1"),genesis.Hash)
			tip = genesis.Hash
		}else {
			tip = b.Get([]byte("1"))
		}
		return nil
	})
	bc := Blockchain{tip,db}
	return &bc
	//return &Blockchain{
	//	[]*block.Block{NewGenesisBlock()},
	//}
}
