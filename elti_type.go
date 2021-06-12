package elti

const (
	elti_bytes  uint8 = 0x00
	elti_string uint8 = 0x01
	elti_int8   uint8 = 0x02
	elti_uint8  uint8 = 0x03
	elti_int16  uint8 = 0x04
	elti_uint16 uint8 = 0x05
	elti_int32  uint8 = 0x06
	elti_uint32 uint8 = 0x07
	elti_int64  uint8 = 0x08
	elti_uint64 uint8 = 0x09
	elti_varint uint8 = 0x0A
	elti_bool   uint8 = 0x0B
	elti_custom uint8 = 0x0C
)

type ValueType uint8

const (
	MAP ValueType = iota
	ARRAY
	DATA
	INVALID
)
