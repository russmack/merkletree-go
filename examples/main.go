package main

import (
	"fmt"

	"github.com/russmack/merkletree-go"
)

func main() {

	// Merkle tree is initialised with complete list of leaf node
	// transactions, ie there is no later adding of nodes.

	txns := []string{"aaa", "bbb", "ccc", "ddd", "eee", "fff", "ggg", "hhh", "iii"}

	m := merkletree.New(txns)

	fmt.Println("tree:", m)

	fmt.Println("tree:", m)
}
