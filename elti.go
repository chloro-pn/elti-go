package elti

import "fmt"

type Elti struct {
	root *Element
}

func NewElti(v Value) *Elti {
	return &Elti{
		root: NewElement("", v),
	}
}

func (e *Elti) SeriToBytes() []byte {
	buf := make([]byte, 0)
	buf = e.root.seri(buf)
	return buf
}

func (e *Elti) ToJson() []byte {
	return e.root.value.ToJson(Base64)
}

func (e *Elti) GetRoot() *ValueWrapper {
	return NewValueWrapper(e.root.value)
}

func ParseToElti(buf []byte, pt ParseType) *Elti {
	ele := NewElement("", nil)
	total_size := ele.parse(buf, 0, pt)
	if total_size != uint32(len(buf)) {
		panic(fmt.Sprintf("ParseToElti error, len(buf) == %d, total_size == %d.", len(buf), total_size))
	}
	return NewElti(ele.value)
}

func ParseToPositioner(buf []byte) *Positioner {
	length, new_begin := parseLength(buf, 0)
	new_begin += length
	vt, new_begin := parseValueType(buf, new_begin)
	return NewPositioner(buf, new_begin, vt, true)
}
