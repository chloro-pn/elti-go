package elti

import "fmt"

type Map struct {
	kvs map[string]Value
}

func NewMap() *Map {
	return &Map{
		kvs: make(map[string]Value),
	}
}

func (m *Map) seriValue(buf []byte) []byte {
	tmp := make([]byte, 0)
	count := uint32(len(m.kvs))
	tmp = seriLength(count, tmp)
	for key, each := range m.kvs {
		ele := NewElement(key, each)
		tmp = ele.seri(tmp)
	}
	total_size := uint32(len(tmp))
	buf = seriLength(total_size, buf)
	buf = append(buf, tmp...)
	return buf
}

func (m *Map) parseValue(buf []byte, begin uint32, pt ParseType) uint32 {
	m.kvs = make(map[string]Value, 0)
	_, new_begin := parseLength(buf, begin)
	count, new_begin := parseLength(buf, new_begin)
	var i uint32
	for i = 0; i < count; i++ {
		ele := NewElement("", nil)
		new_begin = ele.parse(buf, new_begin, pt)
		_, ok := m.kvs[ele.key]
		if ok {
			panic(fmt.Sprintf("Map.parseValue error, duplicate key %s", ele.key))
		}
		m.kvs[ele.key] = ele.value
	}
	return new_begin
}

func (m *Map) getValueType() ValueType {
	return MAP
}

func (m *Map) Set(key string, v Value) {
	m.Erase(key)
	m.kvs[key] = v
}

func (m *Map) Erase(key string) {
	delete(m.kvs, key)
}

func (m *Map) Containes(key string) bool {
	_, ok := m.kvs[key]
	return ok
}

func (m *Map) Attr(key string) Value {
	v, ok := m.kvs[key]
	if !ok {
		panic(fmt.Sprintf("Map.Attr error, key %s does not exist.", key))
	}
	return v
}
