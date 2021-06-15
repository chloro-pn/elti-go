package elti

type DataRef struct {
	buf    []byte
	begin  uint32
	length uint32
	t      DataType
}

func NewDataRef(buf []byte, begin uint32, length uint32) *DataRef {
	return &DataRef{
		buf:    buf,
		begin:  begin,
		length: length,
	}
}

func (dr *DataRef) seriValue(buf []byte) []byte {
	buf = seriLength(uint32(len(dr.buf)+1), buf)
	buf = seriDataType(dr.t, buf)
	buf = append(buf, dr.buf[dr.begin:dr.begin+dr.length]...)
	return buf
}

func (dr *DataRef) parseValue(buf []byte, begin uint32, pt ParseType) uint32 {
	if pt != ParseRefOn {
		panic("DataRef.parseValue error. pt error.")
	}
	var new_begin uint32
	dr.buf = buf
	dr.length, new_begin = parseLength(buf, begin)
	dr.length -= 1
	dr.t, dr.begin = parseDataType(buf, new_begin)
	return dr.begin + dr.length
}

func (dr *DataRef) GetRef() []byte {
	return dr.buf[dr.begin : dr.begin+dr.length]
}

func (dr *DataRef) getValueType() ValueType {
	return DATAREF
}
