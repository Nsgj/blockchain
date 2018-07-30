package block

import (
	"strconv"
	"bytes"
	"time"
	"crypto/sha256"
	"blockchain/proofofwork"
)

type Block struct {
	Timestamp int64
	Data []byte
	PrevBlockHash []byte
	Hash []byte
	Nonce int
}
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func (b *Block)GetData() []byte {
	return b.Data
}
func (b *Block)GetNonce() int {
	return b.Nonce
}

func (b *Block)GetPrevBlockHash() []byte {
	return b.PrevBlockHash
}


func NewBlock(data string,prevBlockHash []byte)  *Block{

	block := &Block{time.Now().Unix(), []byte(data),
	prevBlockHash, []byte{},0}

	pow := proofofwork.NewProofOfWork(block)
	nonce,hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	//block.SetHash()
	return block

}
