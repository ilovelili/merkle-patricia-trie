package mpt

import (
	"github.com/ethereum/go-ethereum/crypto"
)

// In the MPT, there is one more type of nodes apart from the branch nodes and the leaf nodes.
// They are extension nodes. An extension node is an optimized node of the branch node.
// In the Ethereum state, quite frequently, there are branch nodes that have only one child node.
// This is the reason why the MPT compresses branch nodes that contain only one child into extension nodes that have a path and the hash of the child.
type ExtensionNode struct {
	Path []Nibble
	Next Node
}

func NewExtensionNode(nibbles []Nibble, next Node) *ExtensionNode {
	return &ExtensionNode{
		Path: nibbles,
		Next: next,
	}
}

func (e ExtensionNode) Hash() []byte {
	return crypto.Keccak256(e.Serialize())
}

func (e ExtensionNode) Raw() []interface{} {
	hashes := make([]interface{}, 2)
	hashes[0] = ToBytes(ToPrefixed(e.Path, false))
	if len(Serialize(e.Next)) >= 32 {
		hashes[1] = e.Next.Hash()
	} else {
		hashes[1] = e.Next.Raw()
	}
	return hashes
}

func (e ExtensionNode) Serialize() []byte {
	return Serialize(e)
}
