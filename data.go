package elti

import "fmt"

type Data struct {
	v []byte
	t DataType
}

func NewData(v []byte) *Data {
	return &Data{
		v: v,
		t: elti_bytes,
	}
}

func NewDataFromString(s string) *Data {
	return &Data{
		v: []byte(s),
		t: elti_string,
	}
}

func NewDataFromVarint(n uint32) *Data {
	buf := make([]byte, 0)
	return &Data{
		v: seriLength(n, buf),
		t: elti_varint,
	}
}

func NewDataFromBool(b bool) *Data {
	var tmp byte
	if b == true {
		tmp = 0x01
	} else {
		tmp = 0x00
	}
	return &Data{
		v: []byte{tmp},
		t: elti_bool,
	}
}

func NewDataFromInt8(n int8) *Data {
	return &Data{
		v: int8_to_bytes(n),
		t: elti_int8,
	}
}

func NewDataFromUint8(n uint8) *Data {
	return &Data{
		v: uint8_to_bytes(n),
		t: elti_uint8,
	}
}

func NewDataFromInt16(n int16) *Data {
	return &Data{
		v: int16_to_bytes(n),
		t: elti_int16,
	}
}

func NewDataFromUint16(n uint16) *Data {
	return &Data{
		v: uint16_to_bytes(n),
		t: elti_uint16,
	}
}

func NewDataFromInt32(n int32) *Data {
	return &Data{
		v: int32_to_bytes(n),
		t: elti_int32,
	}
}

func NewDataFromUint32(n uint32) *Data {
	return &Data{
		v: uint32_to_bytes(n),
		t: elti_uint32,
	}
}

func NewDataFromInt64(n int64) *Data {
	return &Data{
		v: int64_to_bytes(n),
		t: elti_int64,
	}
}

func NewDataFromUint64(n uint64) *Data {
	return &Data{
		v: uint64_to_bytes(n),
		t: elti_uint64,
	}
}

func (d *Data) seriValue(buf []byte) []byte {
	total_length := uint32(len(d.v))
	buf = seriLength(total_length+1, buf)
	buf = seriDataType(d.t, buf)
	buf = append(buf, d.v...)
	return buf
}

func (d *Data) parseValue(buf []byte, begin uint32) uint32 {
	total_length, new_begin := parseLength(buf, begin)
	if total_length == 0 {
		panic("Data.parseValue error, total_length == 0.")
	}
	d.t, new_begin = parseDataType(buf, new_begin)
	d.v = make([]byte, total_length-1)
	copy(d.v, buf[new_begin:new_begin+total_length-1])
	return new_begin + total_length - 1
}

func (d *Data) getValueType() ValueType {
	return DATA
}

func (d *Data) BytesRef() []byte {
	return d.v
}

func (d *Data) GetAsBytes() []byte {
	if d.t != elti_bytes {
		panic(fmt.Sprintf("Data.GetAsBytes error, real type_id = %d", d.t))
	}
	buf := make([]byte, len(d.v))
	copy(buf, d.v)
	return buf
}

func (d *Data) GetAsString() string {
	if d.t != elti_string {
		panic(fmt.Sprintf("Data.GetAsString error, real type_id = %d", d.t))
	}
	return string(d.v)
}

func (d *Data) GetAsBool() bool {
	if d.t != elti_bool {
		panic(fmt.Sprintf("Data.GetAsBool error, real type_id = %d", d.t))
	}
	var result bool
	if d.v[0] == 0x00 {
		result = false
	} else if d.v[0] == 0x01 {
		result = true
	} else {
		panic("Data.GetAsBool error.")
	}
	return result
}

func (d *Data) GetAsInt8() int8 {
	if d.t != elti_int8 {
		panic(fmt.Sprintf("Data.GetAsInt8 error, real type_id = %d", d.t))
	}
	return bytes_to_int8(d.v)
}

func (d *Data) GetAsUint8() uint8 {
	if d.t != elti_uint8 {
		panic(fmt.Sprintf("Data.GetAsUint8 error, real type_id = %d", d.t))
	}
	return bytes_to_uint8(d.v)
}

func (d *Data) GetAsInt16() int16 {
	if d.t != elti_int16 {
		panic(fmt.Sprintf("Data.GetAsInt16 error, real type_id = %d", d.t))
	}
	return bytes_to_int16(d.v)
}

func (d *Data) GetAsUint16() uint16 {
	if d.t != elti_uint16 {
		panic(fmt.Sprintf("Data.GetAsUint16 error, real type_id = %d", d.t))
	}
	return bytes_to_uint16(d.v)
}

func (d *Data) GetAsInt32() int32 {
	if d.t != elti_int32 {
		panic(fmt.Sprintf("Data.GetAsInt32 error, real type_id = %d", d.t))
	}
	return bytes_to_int32(d.v)
}

func (d *Data) GetAsUint32() uint32 {
	if d.t != elti_uint32 {
		panic(fmt.Sprintf("Data.GetAsUint32 error, real type_id = %d", d.t))
	}
	return bytes_to_uint32(d.v)
}

func (d *Data) GetAsInt64() int64 {
	if d.t != elti_int64 {
		panic(fmt.Sprintf("Data.GetAsInt64 error, real type_id = %d", d.t))
	}
	return bytes_to_int64(d.v)
}

func (d *Data) GetAsUint64() uint64 {
	if d.t != elti_uint64 {
		panic(fmt.Sprintf("Data.GetAsUint64 error, real type_id = %d", d.t))
	}
	return bytes_to_uint64(d.v)
}

func (d *Data) GetAsVarint() uint32 {
	if d.t != elti_varint {
		panic(fmt.Sprintf("Data.GetAsVarint error, real type_id = %d", d.t))
	}
	result, new_begin := parseLength(d.v, 0)
	if new_begin != uint32(len(d.v)) {
		panic("Data.GetAsVarint error.")
	}
	return result
}
