package elti

import (
	"reflect"
	"testing"
)

func TestElement(t *testing.T) {
	ele := NewElement("name", NewData([]byte{0x01, 0x04}))
	buf := make([]byte, 0)
	buf = ele.seri(buf)

	ele2 := &Element{
		key:   "",
		value: nil,
	}

	new_begin := ele2.parse(buf, 0)
	if new_begin != uint32(len(buf)) {
		t.Error("element parse error!")
	}
	if ele2.key != ele.key || reflect.DeepEqual(ele.value, ele2.value) == false {
		t.Error("element parse error!")
	}
}
