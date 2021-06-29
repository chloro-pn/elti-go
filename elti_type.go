package elti

type DataType uint8

const (
	elti_bytes  DataType = 0x00
	elti_string DataType = 0x01
	elti_int8   DataType = 0x02
	elti_uint8  DataType = 0x03
	elti_int16  DataType = 0x04
	elti_uint16 DataType = 0x05
	elti_int32  DataType = 0x06
	elti_uint32 DataType = 0x07
	elti_int64  DataType = 0x08
	elti_uint64 DataType = 0x09
	elti_varint DataType = 0x0A
	elti_bool   DataType = 0x0B
	elti_custom DataType = 0x0C
)

type ValueType uint8

const (
	MAP ValueType = iota
	ARRAY
	DATA
	DATAREF
	INVALID
)

func parseValueTypeCheck(vt ValueType) bool {
	if vt == MAP || vt == ARRAY || vt == DATA {
		return true
	}
	return false
}

type ParseType uint8

const (
	ParseRefOn ParseType = iota
	ParseRefOff
)

type BytesEncodeType uint8

const (
	Base64 BytesEncodeType = iota
	HexStyle
)
