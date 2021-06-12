package elti

type Value interface {
	seriValue(buf []byte) []byte
	parseValue(buf []byte, begin uint32) uint32
	getValueType() ValueType
}

func seriValueType(vt ValueType, buf []byte) []byte {
	v := byte(vt)
	buf = append(buf, v)
	return buf
}

func parseValueType(buf []byte, begin uint32) (ValueType, uint32) {
	vt := buf[begin]
	return ValueType(vt), begin + 1
}

func seriDataType(dt DataType, buf []byte) []byte {
	v := byte(dt)
	buf = append(buf, v)
	return buf
}

func parseDataType(buf []byte, begin uint32) (DataType, uint32) {
	dt := buf[begin]
	return DataType(dt), begin + 1
}

func parseKey(buf []byte, begin uint32) (string, uint32) {
	length, new_begin := parseLength(buf, begin)
	key := string(buf[new_begin : new_begin+length])
	new_begin += length
	return key, new_begin
}

func valueFactory(vt ValueType) Value {
	if vt == DATA {
		return NewData(nil)
	} else if vt == ARRAY {
		return NewArray()
	} else if vt == MAP {
		return NewMap()
	}
	return nil
}
