package elti

type ValueWrapper struct {
	v Value
}

func NewValueWrapper(v Value) *ValueWrapper {
	return &ValueWrapper{v: v}
}

func (vw *ValueWrapper) Attr(key string) *ValueWrapper {
	vt, ok := vw.v.(*Map)
	if !ok {
		panic("ValueWrapper.Attr error, type mismatch.")
	}
	return NewValueWrapper(vt.Attr(key))
}

func (vw *ValueWrapper) At(index uint32) *ValueWrapper {
	vt, ok := vw.v.(*Array)
	if !ok {
		panic("ValueWrapper.At error, type mismatch.")
	}
	return NewValueWrapper(vt.At(index))
}

func (vw *ValueWrapper) Size() uint32 {
	vt, ok := vw.v.(*Array)
	if !ok {
		panic("ValueWrapper.Size error, type mismatch.")
	}
	return vt.Size()
}

func (vw *ValueWrapper) GetValueType() ValueType {
	return vw.v.getValueType()
}

func (vw *ValueWrapper) GetAsBytes() []byte {
	vt, ok := vw.v.(*Data)
	if !ok {
		panic("ValueWrapper.GetAsBytes error, type mismatch")
	}
	return vt.GetAsBytes()
}

func (vw *ValueWrapper) GetAsString() string {
	vt, ok := vw.v.(*Data)
	if !ok {
		panic("ValueWrapper.GetAsString error, type mismatch")
	}
	return vt.GetAsString()
}

func (vw *ValueWrapper) GetAsVarint() uint32 {
	vt, ok := vw.v.(*Data)
	if !ok {
		panic("ValueWrapper.GetAsVarint error, type mismatch")
	}
	return vt.GetAsVarint()
}
