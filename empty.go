package mpt

import "encoding/hex"

const EthereumRootHash = "56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"

var (
	EmptyNodeRaw     = []byte{}
	EmptyNodeHash, _ = hex.DecodeString(EthereumRootHash)
)

func IsEmptyNode(node Node) bool {
	return node == nil
}
