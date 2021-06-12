package elti

type Data struct {
	v []byte
}

func NewData(v []byte) *Data {
	return &Data{
		v: v,
	}
}

func NewDataFromString(s string) *Data {
	return &Data{
		v: []byte(s),
	}
}

func NewDataFromVarint(n uint32) *Data {
	buf := make([]byte, 0)
	return &Data{
		v: seriLength(n, buf),
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
	}
}

func NewDataFromInt8(n int8) *Data {
	return &Data{
		v: int8_to_bytes(n),
	}
}

func NewDataFromUint8(n uint8) *Data {
	return &Data{
		v: uint8_to_bytes(n),
	}
}

func NewDataFromInt16(n int16) *Data {
	return &Data{
		v: int16_to_bytes(n),
	}
}

func NewDataFromUint16(n uint16) *Data {
	return &Data{
		v: uint16_to_bytes(n),
	}
}

func NewDataFromInt32(n int32) *Data {
	return &Data{
		v: int32_to_bytes(n),
	}
}

func NewDataFromUint32(n uint32) *Data {
	return &Data{
		v: uint32_to_bytes(n),
	}
}

func NewDataFromInt64(n int64) *Data {
	return &Data{
		v: int64_to_bytes(n),
	}
}

func NewDataFromUint64(n uint64) *Data {
	return &Data{
		v: uint64_to_bytes(n),
	}
}

func (d *Data) seriValue(buf []byte) []byte {
	total_length := uint32(len(d.v))
	buf = seriLength(total_length, buf)
	buf = append(buf, d.v...)
	return buf
}

func (d *Data) parseValue(buf []byte, begin uint32) uint32 {
	total_length, new_begin := parseLength(buf, begin)
	d.v = make([]byte, total_length)
	copy(d.v, buf[new_begin:new_begin+total_length])
	return new_begin + total_length
}

func (d *Data) getValueType() ValueType {
	return DATA
}

func (d *Data) BytesRef() []byte {
	return d.v
}

func (d *Data) GetAsBytes() []byte {
	buf := make([]byte, len(d.v))
	copy(buf, d.v)
	return buf
}

func (d *Data) GetAsString() string {
	return string(d.v)
}

func (d *Data) GetAsBool() bool {
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
	return bytes_to_int8(d.v)
}

func (d *Data) GetAsUint8() uint8 {
	return bytes_to_uint8(d.v)
}

func (d *Data) GetAsInt16() int16 {
	return bytes_to_int16(d.v)
}

func (d *Data) GetAsUint16() uint16 {
	return bytes_to_uint16(d.v)
}

func (d *Data) GetAsInt32() int32 {
	return bytes_to_int32(d.v)
}

func (d *Data) GetAsUint32() uint32 {
	return bytes_to_uint32(d.v)
}

func (d *Data) GetAsInt64() int64 {
	return bytes_to_int64(d.v)
}

func (d *Data) GetAsUint64() uint64 {
	return bytes_to_uint64(d.v)
}

func (d *Data) GetAsVarint() uint32 {
	result, new_begin := parseLength(d.v, 0)
	if new_begin != uint32(len(d.v)) {
		panic("Data.GetAsVarint error.")
	}
	return result
}
