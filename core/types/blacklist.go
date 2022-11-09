package types

import "github.com/ethereum/go-ethereum/common"

// This is introduced because of the Tendermint IAVL Merkle Proof verification exploitation.
var NanoBlackList = []common.Address{
	common.HexToAddress("0x1337a83A3bf6964C249aE86b6E96fE429fA74C52"),

	common.HexToAddress("0x7dFd218943Da7E8Ef66339F5E0eBb7C5B7796B27"),
	common.HexToAddress("0x06244d0635c36057B65011C9C12019DB9d471F31"),
}
