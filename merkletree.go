package merkletree

import (
	"fmt"
	"hash/fnv"
	"strconv"
)

type merkleTree struct {
	Root *node
}

type node struct {
	Hash  uint64
	Left  *node
	Right *node
}

// New takes a list of transactions and returns a merkle tree.
func New(txns []string) *merkleTree {
	nodes := txnsToLeafNodes(txns)

	for len(nodes) > 1 {
		nodes = makeParents(nodes)
	}

	t := &merkleTree{
		Root: nodes[0],
	}

	return t
}

// PrintTree prints the nodes of the merkle tree.
func (m *merkleTree) PrintTree() {
	n := m.Root
	level := []*node{n}
	printByLevel(level)
}

func printByLevel(level []*node) {
	fmt.Println("-----------------")
	if len(level) == 0 {
		return
	}

	buff := []*node{}

	for i, n := range level {
		if n != nil {
			fmt.Printf("> %p : %+v\n", level[i], level[i])
			buff = append(buff, n.Left)
			buff = append(buff, n.Right)
		}
	}

	if len(buff) == 0 {
		return
	}

	printByLevel(buff)
}

func hashFnv1a(s ...string) uint64 {
	h := fnv.New64a()
	for i := range s {
		h.Write([]byte(s[i]))
	}
	return h.Sum64()
}

func txnsToLeafNodes(txns []string) []*node {
	nodes := make([]*node, len(txns))

	for i, _ := range txns {
		n := &node{
			Hash:  hashFnv1a(txns[i]),
			Left:  nil,
			Right: nil,
		}

		nodes[i] = n
	}

	return nodes
}

func makeParents(nodes []*node) []*node {
	nodeCount := len(nodes)
	nParents := (nodeCount + 1) / 2
	parents := make([]*node, nParents)

	// For each pair of transaction leaf nodes create a parent node.
	pairedItems := nodeCount - (nodeCount % 2)

	p := 0
	for i := 0; i < pairedItems-1; i += 2 {
		left := nodes[i]
		right := nodes[i+1]

		parent := &node{
			hashFnv1a(
				strconv.FormatUint(left.Hash, 10),
				strconv.FormatUint(right.Hash, 10),
			),
			left,
			right,
		}

		parents[p] = parent
		p++
	}

	// Create parent for remaining unpaired node if an odd number of nodes.
	if nodeCount%2 != 0 {
		left := nodes[nodeCount-1]

		parent := &node{
			Hash:  left.Hash,
			Left:  left,
			Right: nil,
		}

		parents[len(parents)-1] = parent
	}

	return parents
}
