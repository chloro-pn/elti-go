package elti

import (
	"bytes"
	"encoding/binary"
)

func int8_to_bytes(n int8) []byte {
	return []byte{byte(n)}
}

func bytes_to_int8(b []byte) int8 {
	return int8(b[0])
}

func uint8_to_bytes(n uint8) []byte {
	return []byte{byte(n)}
}

func bytes_to_uint8(b []byte) uint8 {
	return uint8(b[0])
}

func int16_to_bytes(n int16) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, n)
	return buf.Bytes()
}

func bytes_to_int16(b []byte) int16 {
	buf := bytes.NewBuffer(b)
	var n int16
	binary.Read(buf, binary.BigEndian, &n)
	return n
}

func uint16_to_bytes(n uint16) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, n)
	return buf.Bytes()
}

func bytes_to_uint16(b []byte) uint16 {
	buf := bytes.NewBuffer(b)
	var n uint16
	binary.Read(buf, binary.BigEndian, &n)
	return n
}

func int32_to_bytes(n int32) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, n)
	return buf.Bytes()
}

func bytes_to_int32(b []byte) int32 {
	buf := bytes.NewBuffer(b)
	var n int32
	binary.Read(buf, binary.BigEndian, &n)
	return n
}

func uint32_to_bytes(n uint32) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, n)
	return buf.Bytes()
}

func bytes_to_uint32(b []byte) uint32 {
	buf := bytes.NewBuffer(b)
	var n uint32
	binary.Read(buf, binary.BigEndian, &n)
	return n
}

func int64_to_bytes(n int64) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, n)
	return buf.Bytes()
}

func bytes_to_int64(b []byte) int64 {
	buf := bytes.NewBuffer(b)
	var n int64
	binary.Read(buf, binary.BigEndian, &n)
	return n
}

func uint64_to_bytes(n uint64) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, n)
	return buf.Bytes()
}

func bytes_to_uint64(b []byte) uint64 {
	buf := bytes.NewBuffer(b)
	var n uint64
	binary.Read(buf, binary.BigEndian, &n)
	return n
}
