package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/require"
	"testing"
)

//https://github.com/zhangchiqing/merkle-patricia-trie

func byteArrToHexArr(bytes []byte) []string {
	hexArr := make([]string, len(bytes))

	for i, b := range bytes {
		hexArr[i] = fmt.Sprintf("%02x", b)
	}

	return hexArr
}

func TestRLP(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		var arr = []interface{}{"cat", "dog"}

		rl1, err := rlp.EncodeToBytes(arr[0])
		require.NoError(t, err)
		rl2, err := rlp.EncodeToBytes(arr[1])
		require.NoError(t, err)
		rl3, err := rlp.EncodeToBytes(arr)
		require.NoError(t, err)

		l1 := len(rl1)
		l2 := len(rl2)
		l3 := len(rl3)
		t.Logf("%v - %v", byteArrToHexArr(rl1), l1)
		t.Logf("%v - %v", byteArrToHexArr(rl2), l2)
		t.Logf("%v - %v", byteArrToHexArr(rl3), l3)

		var dec1 string
		var dec2 string
		rlp.DecodeBytes(rl1, &dec1)
		rlp.DecodeBytes(rl2, &dec2)
		require.Equal(t, dec1, arr[0])
		require.Equal(t, dec2, arr[1])

		dec3 := []string{}
		rlp.DecodeBytes(rl3, &dec3)

		t.Logf("%v", dec1)
		t.Logf("%v", dec2)
		t.Logf("%v", dec3)
	})

	//t.Run("test2", func(t *testing.T) {
	//	input := "dog"
	//	rl, err := rlp.EncodeToBytes(input)
	//	require.NoError(t, err)
	//
	//	l := len(rl)
	//	t.Logf("%v - %v", byteArrToHexArr(rl), l)
	//})
}
