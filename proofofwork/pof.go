package proofofwork

import (
	"myblockchain/block"
	"math/big"
)

const targetBits  = 24

type ProofOfWork struct {
	block *block.Block
	target *big.Int
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {

}

func NewProofOfWork(b *block.Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target,uint(256-targetBits))

	pow := &ProofOfWork{b,target}
	return pow
}
