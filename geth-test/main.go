package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
)

func main() {
	fmt.Println("Hello")
	blockNonce := types.EncodeNonce(22)
	nonce := blockNonce.Uint64()
	marshalText, _ := blockNonce.MarshalText()

	fmt.Printf("blockNonce: %v\n", blockNonce)
	fmt.Printf("blockNonce hex: %#v\n", blockNonce)
	fmt.Printf("nonce: %v\n", nonce)
	fmt.Printf("marshalText: %v\n", marshalText)
	fmt.Printf("marshalText in string: %v\n", string(marshalText))

	fmt.Printf("-------------------------------\n")

	y := new(types.BlockNonce)
	y.UnmarshalText(marshalText)
	fmt.Printf("blockNonce: %v\n", y)
	fmt.Printf("blockNonce hex: %#v\n", y)
	fmt.Printf("nonce: %v\n", y.Uint64())
}
