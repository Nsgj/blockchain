package proofofwork

import (
	"math/big"
	"bytes"
	"strconv"
	"fmt"
	"math"
	"crypto/sha256"
)

type Blockin interface {
	GetData() []byte
	GetPrevBlockHash() []byte
	GetNonce() int
}

const targetBits  = 24

type ProofOfWork struct {
	block Blockin
	target *big.Int
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join([][]byte{
		pow.block.GetPrevBlockHash(),
		pow.block.GetData(),
		IntToHex(int64(targetBits)),
		IntToHex(int64(nonce)),
	},
	[]byte{})
	return data
}

func (pow *ProofOfWork)Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.GetNonce())
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}

func (pow *ProofOfWork)Run() (int,[]byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	maxNonce := math.MaxInt64

	fmt.Printf("Mining the block containing \"%s\"\n",pow.block.GetData())

	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x",hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1{
			break
		}else {
			nonce++
		}
		}
		fmt.Print("\n\n")
	return nonce,hash[:]
}

func IntToHex(n int64) []byte {
	return []byte(strconv.FormatInt(n, 16))
}

func NewProofOfWork(b Blockin) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target,uint(256-targetBits))

	pow := &ProofOfWork{b,target}
	return pow
}

