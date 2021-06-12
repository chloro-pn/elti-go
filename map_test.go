package elti

import (
	"testing"
)

func TestMap(t *testing.T) {
	m := NewMap()
	m.Set("name", NewDataFromString("bob"))
	m.Set("age", NewDataFromUint8(25))
	buf := make([]byte, 0)
	buf = m.seriValue(buf)

	m2 := NewMap()
	new_begin := m2.parseValue(buf, 0)
	if new_begin != uint32(len(buf)) {
		t.Error("test map error.")
	}
	if m2.Attr("name").getValueType() != DATA {
		t.Error("test map error.")
	}
	if m2.Attr("name").(*Data).GetAsString() != "bob" {
		t.Error("test map error.")
	}
	if m2.Attr("age").(*Data).GetAsUint8() != 25 {
		t.Error("test map error.")
	}
}
