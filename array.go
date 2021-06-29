package elti

type Array struct {
	vs []Value
}

func NewArray() *Array {
	return &Array{
		vs: make([]Value, 0),
	}
}

func (a *Array) PushBack(v Value) {
	a.vs = append(a.vs, v)
}

func (a *Array) Erase(index uint32) {
	a.vs = append(a.vs[:index], a.vs[index+1:]...)
}

func (a *Array) Size() uint32 {
	return uint32(len(a.vs))
}

func (a *Array) At(index uint32) Value {
	return a.vs[index]
}

// format : [ total_size | count | value-1 | ... | value-n ]
func (d *Array) seriValue(buf []byte) []byte {
	count := uint32(len(d.vs))
	var tmp_buf []byte = make([]byte, 0)
	tmp_buf = seriLength(count, tmp_buf)
	for _, each := range d.vs {
		vt := each.getValueType()
		tmp_buf = seriValueType(vt, tmp_buf)
		tmp_buf = each.seriValue(tmp_buf)
	}
	buf = seriLength(uint32(len(tmp_buf)), buf)
	buf = append(buf, tmp_buf...)
	return buf
}

func (d *Array) parseValue(buf []byte, begin uint32, pt ParseType) uint32 {
	d.vs = make([]Value, 0)
	_, new_begin := parseLength(buf, begin)
	count, new_begin := parseLength(buf, new_begin)
	var i uint32
	for i = 0; i < count; i++ {
		var vt ValueType
		vt, new_begin = parseValueType(buf, new_begin)
		v := valueFactory(vt, pt)
		new_begin = v.parseValue(buf, new_begin, pt)
		d.vs = append(d.vs, v)
	}
	return new_begin
}

func (d *Array) getValueType() ValueType {
	return ARRAY
}

func (d *Array) ToJson(bt BytesEncodeType) []byte {
	var result []byte
	result = append(result, '[')
	for _, each := range d.vs {
		result = append(result, each.ToJson(bt)...)
		result = append(result, ',')
	}
	if result[len(result)-1] == ',' {
		result = result[:len(result)-1]
	}
	result = append(result, ']')
	return result
}
