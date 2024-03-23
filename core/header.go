package core

import (
	"encoding/gob"
	"io"

	"github.com/Riyuze/riyu-net/types"
)

type Header struct {
	Version   uint
	PrevBlock types.Hash
	Timestamp uint64
	Height    uint32
	Nonce     uint64
}

func (h *Header) EncodeBinary(w io.Writer) error {
	enc := gob.NewEncoder(w)

	err := enc.Encode(h)
	if err != nil {
		return err
	}

	return nil
}

func (h *Header) DecodeBinary(r io.Reader) error {
	dec := gob.NewDecoder(r)

	err := dec.Decode(h)
	if err != nil {
		return err
	}

	return nil
}
