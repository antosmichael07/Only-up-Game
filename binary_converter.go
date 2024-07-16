package main

import (
	"encoding/binary"
	"math"
)

func float32_to_bytes(f float32) []byte {
	bytes := make([]byte, 4)
	bits := math.Float32bits(f)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}

func bytes_to_float32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}
