package main

import (
	"blockchain/blockchain"
	"fmt"
)

func main() {
	bc := blockchain.NewBlockchain()
	bc.AddBlock("Send 1BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _,block := range bc.GetBlocks() {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
