package main

import (
	"encoding/binary"
	"testing"
)

func BenchmarkBinary(b *testing.B) {
	data := make([]byte, 4)
	for i := 0; i < b.N; i++ {
		binary.BigEndian.Uint16(data)
		binary.BigEndian.Uint16(data[2:])
	}
}
