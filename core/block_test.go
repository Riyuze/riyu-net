package core

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/Riyuze/riyu-net/types"
	"github.com/stretchr/testify/assert"
)

func TestHeader_Encode_Decode(t *testing.T) {
	h := Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: uint64(time.Now().UnixNano()),
		Height:    10,
		Nonce:     123456,
	}

	buf := bytes.Buffer{}

	h.EncodeBinary(&buf)
	fmt.Printf("%+v\n", h)

	hDecode := Header{}

	hDecode.DecodeBinary(&buf)
	fmt.Printf("%+v\n", hDecode)

	assert.Equal(t, h, hDecode)
}

func TestBlock_Encode_Decode(t *testing.T) {
	b := Block{
		header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: uint64(time.Now().UnixNano()),
			Height:    10,
			Nonce:     123456,
		},
		transactions: nil,
	}

	buf := bytes.Buffer{}

	b.EncodeBinary(&buf)
	fmt.Printf("%+v\n", b)

	bDecode := Block{}

	bDecode.DecodeBinary(&buf)
	fmt.Printf("%+v\n", bDecode)

	assert.Equal(t, b, bDecode)
}

func TestBlockHash(t *testing.T) {
	b := Block{
		header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: uint64(time.Now().UnixNano()),
			Height:    10,
			Nonce:     123456,
		},
		transactions: nil,
	}

	h := b.Hash()
	fmt.Println(h)

	assert.False(t, h.IsZero())
}
