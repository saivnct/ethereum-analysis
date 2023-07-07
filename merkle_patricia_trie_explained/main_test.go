package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

//https://github.com/zhangchiqing/merkle-patricia-trie

func TestMain(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		require.Equal(t, 1, 1)
	})
}
