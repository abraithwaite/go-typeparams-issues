package main

import (
	"encoding"
)

type Seralizable interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

type SerDeString string

func (s *SerDeString) UnmarshalBinary(in []byte) error {
	*s = SerDeString(in)
	return nil
}

func (s SerDeString) MarshalBinary() ([]byte, error) {
	return []byte(s), nil
}


type GenericSerializable[T Seralizable] struct {
	Key string
	Value T
}

func (g GenericSerializable[T]) Send() {
	_, _ = g.Value.MarshalBinary()
}

func main() {
	val := SerDeString("asdf")
	x := GenericSerializable[*SerDeString]{
		Value: &val,
	}
	x.Send()
}
