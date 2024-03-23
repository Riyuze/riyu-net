package core

import (
	"bytes"
	"crypto/sha256"
	"io"

	"github.com/Riyuze/riyu-net/types"
)

type Block struct {
	header       Header
	transactions []Transaction

	hash types.Hash
}

func (b *Block) Hash() types.Hash {
	buf := bytes.Buffer{}
	b.header.EncodeBinary(&buf)

	if b.hash.IsZero() {
		b.hash = types.Hash(sha256.Sum256(buf.Bytes()))
	}

	return b.hash
}

func (b *Block) EncodeBinary(w io.Writer) error {
	err := b.header.EncodeBinary(w)
	if err != nil {
		return err
	}

	for _, tx := range b.transactions {
		err := tx.EncodeBinary(w)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Block) DecodeBinary(r io.Reader) error {
	err := b.header.DecodeBinary(r)
	if err != nil {
		return err
	}

	for _, tx := range b.transactions {
		err := tx.DecodeBinary(r)
		if err != nil {
			return err
		}
	}

	return nil
}
