package merkletree

import (
	//"fmt"
	"hash/fnv"
	"strconv"
)

type MerkleTree struct {
	Root *node
}

type node struct {
	Hash  uint64
	Left  *node
	Right *node
}

func New(txns []string) *MerkleTree {
	return build(txns)
	//		Root: &node{},
}

func hashFnv1a(s ...string) uint64 {
	h := fnv.New64a()
	for i := range s {
		h.Write([]byte(s[i]))
	}
	return h.Sum64()
}

func build(txns []string) *MerkleTree {
	nparents := (len(txns) + 1) / 2

	parents := make([]*node, nparents)

	t := &MerkleTree{}

	// For each pair of transaction leaf nodes create a parent node.

	p := 0
	for i := 0; i < len(txns)-1; i += 2 {
		left := &node{
			hashFnv1a(txns[i]),
			nil,
			nil,
		}

		right := &node{
			hashFnv1a(txns[i+1]),
			nil,
			nil,
		}

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

	// If odd txns len now add the unpaired final txn.
	if len(txns)%2 != 0 {
		left := &node{
			hashFnv1a(txns[len(txns)-1]),
			nil,
			nil,
		}

		parent := &node{
			left.Hash,
			left,
			nil,
		}

		parents[len(parents)-1] = parent
	}

	return t
}
