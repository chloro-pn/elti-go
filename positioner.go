package elti

import "fmt"

type Positioner struct {
	buf         []byte
	begin_index uint32
	value_type  ValueType
	total_size  uint32
	find        bool
}

func NewPositioner(buf []byte, bi uint32, vt ValueType, find bool) *Positioner {
	p := &Positioner{
		buf:         buf,
		begin_index: bi,
		value_type:  vt,
		total_size:  0,
		find:        find,
	}
	p.total_size, p.begin_index = parseLength(p.buf, p.begin_index)
	return p
}

func (p *Positioner) Finded() bool {
	return p.find
}

func (p *Positioner) Attr(key string) *Positioner {
	if p.value_type != MAP {
		panic("Positioner.Attr error, type mismatch.")
	}
	current_index := p.begin_index
	count, current_index := parseLength(p.buf, current_index)
	var i uint32
	for i = 0; i < count; i++ {
		var key1 string
		key1, current_index = parseKey(p.buf, current_index)
		var vt ValueType
		vt, current_index = parseValueType(p.buf, current_index)
		if key1 == key {
			return NewPositioner(p.buf, current_index, vt, true)
		} else {
			var total_size uint32
			total_size, current_index = parseLength(p.buf, current_index)
			current_index += total_size
		}
	}
	return NewPositioner(nil, 0, INVALID, false)
}

func (p *Positioner) At(index uint32) *Positioner {
	if p.value_type != ARRAY {
		panic("Positioner.At error, type mismatch.")
	}
	current_index := p.begin_index
	count, current_index := parseLength(p.buf, current_index)
	if index >= count {
		return NewPositioner(nil, 0, INVALID, false)
	}
	var i uint32
	for i = 0; i < count; i++ {
		var vt ValueType
		vt, current_index = parseValueType(p.buf, current_index)
		if i == index {
			return NewPositioner(p.buf, current_index, vt, true)
		} else {
			var total_size uint32
			total_size, current_index = parseLength(p.buf, current_index)
			current_index += total_size
		}
	}
	return NewPositioner(nil, 0, INVALID, false)
}

func (p *Positioner) GetAsBytes() []byte {
	if p.value_type != DATA {
		panic("Positioiner.GetAsBytes error, type mismatch.")
	}
	dt := DataType(p.buf[p.begin_index])
	if dt != elti_bytes {
		panic(fmt.Sprintf("Positioner.GetAsBytes error, type_id = %d", dt))
	}
	return p.buf[p.begin_index+1 : p.begin_index+p.total_size]
}

func (p *Positioner) GetAsString() string {
	if p.value_type != DATA {
		panic("Positioiner.GetAsString error, type mismatch.")
	}
	dt := DataType(p.buf[p.begin_index])
	if dt != elti_string {
		panic(fmt.Sprintf("Positioner.GetAsBytes error, type_id = %d", dt))
	}
	return string(p.buf[p.begin_index+1 : p.begin_index+p.total_size])
}

func (p *Positioner) GetAsVarint() uint32 {
	if p.value_type != DATA {
		panic("Positioiner.GetAsVarint error, type mismatch.")
	}
	dt := DataType(p.buf[p.begin_index])
	if dt != elti_varint {
		panic(fmt.Sprintf("Positioner.GetAsVarint error, type_id = %d", dt))
	}
	result, new_begin := parseLength(p.buf, p.begin_index+1)
	if new_begin != (p.begin_index + p.total_size) {
		panic("Positioner.GetAsVarint error.")
	}
	return result
}

func (p *Positioner) GetAsBool() bool {
	if p.value_type != DATA {
		panic("Positioiner.GetAsBool error, type mismatch.")
	}
	dt := DataType(p.buf[p.begin_index])
	if dt != elti_bool {
		panic(fmt.Sprintf("Positioner.GetAsBytes error, type_id = %d", dt))
	}
	var result bool
	if p.buf[p.begin_index+1] == 0x00 {
		result = false
	} else if p.buf[p.begin_index+1] == 0x01 {
		result = true
	} else {
		panic("Positioner.GetAsBool error, invalid byte.")
	}
	return result
}

func (p *Positioner) GetAsUint8() uint8 {
	if p.value_type != DATA {
		panic("Positioiner.GetAsUint8 error, type mismatch.")
	}
	dt := DataType(p.buf[p.begin_index])
	if dt != elti_uint8 {
		panic(fmt.Sprintf("Positioner.GetAsUint8 error, type_id = %d", dt))
	}
	return bytes_to_uint8(p.buf[p.begin_index+1 : p.begin_index+p.total_size-1])
}

func (p *Positioner) GetAsInt8() int8 {
	if p.value_type != DATA {
		panic("Positioiner.GetAsInt8 error, type mismatch.")
	}
	dt := DataType(p.buf[p.begin_index])
	if dt != elti_int8 {
		panic(fmt.Sprintf("Positioner.GetAsInt8 error, type_id = %d", dt))
	}
	return bytes_to_int8(p.buf[p.begin_index+1 : p.begin_index+p.total_size-1])
}

func (p *Positioner) GetAsUint16() uint16 {
	if p.value_type != DATA {
		panic("Positioiner.GetAsUint16 error, type mismatch.")
	}
	dt := DataType(p.buf[p.begin_index])
	if dt != elti_uint16 {
		panic(fmt.Sprintf("Positioner.GetAsUint16 error, type_id = %d", dt))
	}
	return bytes_to_uint16(p.buf[p.begin_index+1 : p.begin_index+p.total_size-1])
}

func (p *Positioner) GetAsInt16() int16 {
	if p.value_type != DATA {
		panic("Positioiner.GetAsInt16 error, type mismatch.")
	}
	dt := DataType(p.buf[p.begin_index])
	if dt != elti_int16 {
		panic(fmt.Sprintf("Positioner.GetAsInt16 error, type_id = %d", dt))
	}
	return bytes_to_int16(p.buf[p.begin_index+1 : p.begin_index+p.total_size-1])
}

func (p *Positioner) GetAsUint32() uint32 {
	if p.value_type != DATA {
		panic("Positioiner.GetAsUint32 error, type mismatch.")
	}
	dt := DataType(p.buf[p.begin_index])
	if dt != elti_uint32 {
		panic(fmt.Sprintf("Positioner.GetAsUint32 error, type_id = %d", dt))
	}
	return bytes_to_uint32(p.buf[p.begin_index+1 : p.begin_index+p.total_size-1])
}

func (p *Positioner) GetAsInt32() int32 {
	if p.value_type != DATA {
		panic("Positioiner.GetAsInt32 error, type mismatch.")
	}
	dt := DataType(p.buf[p.begin_index])
	if dt != elti_int32 {
		panic(fmt.Sprintf("Positioner.GetAsInt32 error, type_id = %d", dt))
	}
	return bytes_to_int32(p.buf[p.begin_index+1 : p.begin_index+p.total_size-1])
}

func (p *Positioner) GetAsUint64() uint64 {
	if p.value_type != DATA {
		panic("Positioiner.GetAsUint64 error, type mismatch.")
	}
	dt := DataType(p.buf[p.begin_index])
	if dt != elti_bool {
		panic(fmt.Sprintf("Positioner.GetAsUint64 error, type_id = %d", dt))
	}
	return bytes_to_uint64(p.buf[p.begin_index+1 : p.begin_index+p.total_size-1])
}

func (p *Positioner) GetAsInt64() int64 {
	if p.value_type != DATA {
		panic("Positioiner.GetAsInt64 error, type mismatch.")
	}
	dt := DataType(p.buf[p.begin_index])
	if dt != elti_bool {
		panic(fmt.Sprintf("Positioner.GetAsInt64 error, type_id = %d", dt))
	}
	return bytes_to_int64(p.buf[p.begin_index+1 : p.begin_index+p.total_size-1])
}
