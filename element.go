package elti

type Element struct {
	key   string
	value Value
}

func NewElement(key string, v Value) *Element {
	return &Element{
		key:   key,
		value: v,
	}
}

// format : [ key_length | key | value_type | value ]
func (e *Element) seri(buf []byte) []byte {
	var length uint32 = uint32(len(e.key))
	buf = seriLength(length, buf)
	buf = append(buf, []byte(e.key)...)
	vt := e.value.getValueType()
	buf = seriValueType(vt, buf)
	buf = e.value.seriValue(buf)
	return buf
}

func (e *Element) parse(buf []byte, begin uint32) uint32 {
	var new_begin uint32
	e.key, new_begin = parseKey(buf, begin)
	vt, new_begin := parseValueType(buf, new_begin)
	e.value = valueFactory(vt)
	new_begin = e.value.parseValue(buf, new_begin)
	return new_begin
}
