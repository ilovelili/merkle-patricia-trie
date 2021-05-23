package mpt

import "github.com/ethereum/go-ethereum/rlp"

// Key-values of the Ethereum state are used as paths on the MPT.
// Nibble is the unit used to distinguish key values in the MPT,
// so each node can have up to 16 branches.
// Additionally, since a node has its own value, a branch node is an array of 17 items composed of 1 node value and 16 branches.
type Node interface {
	Hash() []byte
	Raw() []interface{}
}

func Hash(node Node) []byte {
	if IsEmptyNode(node) {
		return EmptyNodeHash
	}
	return node.Hash()
}

func Serialize(node Node) []byte {
	var raw interface{}

	if IsEmptyNode(node) {
		raw = EmptyNodeRaw
	} else {
		raw = node.Raw()
	}

	rlp, err := rlp.EncodeToBytes(raw)
	if err != nil {
		panic(err)
	}

	return rlp
}
