package main

import (
	"fmt"

	"github.com/lkabell/go-blockchain/blockchain"
)

func main(){
	chain := blockchain.InitBlockChain()
	chain.AddBlock("This is a block")
	chain.AddBlock("This is the next block")

	for _, block := range chain.Blocks {
		fmt.Printf("block.PrevHash: %x\n", block.PrevHash)
		fmt.Printf("block.Data: %s\n", block.Data)
		fmt.Printf("block.Hash: %x\n\n", block.Hash)
	}
}