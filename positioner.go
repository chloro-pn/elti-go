package elti

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
	return p.buf[p.begin_index : p.begin_index+p.total_size]
}

func (p *Positioner) GetAsString() string {
	if p.value_type != DATA {
		panic("Positioiner.GetAsString error, type mismatch.")
	}
	return string(p.buf[p.begin_index : p.begin_index+p.total_size])
}

func (p *Positioner) GetAsVarint() uint32 {
	result, new_begin := parseLength(p.buf, 0)
	if new_begin != uint32(len(p.buf)) {
		panic("Positioner.GetAsVarint error.")
	}
	return result
}

func (p *Positioner) GetAsBool() bool {
	if p.value_type != DATA {
		panic("Positioiner.GetAsBool error, type mismatch.")
	}
	var result bool
	if p.buf[0] == 0x00 {
		result = false
	} else if p.buf[0] == 0x01 {
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
	return bytes_to_uint8(p.buf[p.begin_index : p.begin_index+p.total_size])
}

func (p *Positioner) GetAsInt8() int8 {
	if p.value_type != DATA {
		panic("Positioiner.GetAsInt8 error, type mismatch.")
	}
	return bytes_to_int8(p.buf[p.begin_index : p.begin_index+p.total_size])
}

func (p *Positioner) GetAsUint16() uint16 {
	if p.value_type != DATA {
		panic("Positioiner.GetAsUint16 error, type mismatch.")
	}
	return bytes_to_uint16(p.buf[p.begin_index : p.begin_index+p.total_size])
}

func (p *Positioner) GetAsInt16() int16 {
	if p.value_type != DATA {
		panic("Positioiner.GetAsInt16 error, type mismatch.")
	}
	return bytes_to_int16(p.buf[p.begin_index : p.begin_index+p.total_size])
}

func (p *Positioner) GetAsUint32() uint32 {
	if p.value_type != DATA {
		panic("Positioiner.GetAsUint32 error, type mismatch.")
	}
	return bytes_to_uint32(p.buf[p.begin_index : p.begin_index+p.total_size])
}

func (p *Positioner) GetAsInt32() int32 {
	if p.value_type != DATA {
		panic("Positioiner.GetAsInt32 error, type mismatch.")
	}
	return bytes_to_int32(p.buf[p.begin_index : p.begin_index+p.total_size])
}

func (p *Positioner) GetAsUint64() uint64 {
	if p.value_type != DATA {
		panic("Positioiner.GetAsUint64 error, type mismatch.")
	}
	return bytes_to_uint64(p.buf[p.begin_index : p.begin_index+p.total_size])
}

func (p *Positioner) GetAsInt64() int64 {
	if p.value_type != DATA {
		panic("Positioiner.GetAsInt64 error, type mismatch.")
	}
	return bytes_to_int64(p.buf[p.begin_index : p.begin_index+p.total_size])
}
