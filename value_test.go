package elti

import (
	"testing"
)

func TestParseValueType(t *testing.T) {
	var vt ValueType = MAP
	buf := []byte{byte(vt)}
	vt, new_begin := parseValueType(buf, uint32(0))
	if new_begin != 1 || vt != MAP {
		t.Error("parse value type error.")
	}
	buf = append(buf, byte(DATA))
	vt, new_begin = parseValueType(buf, new_begin)
	if new_begin != 2 || vt != DATA {
		t.Error("parse value type error.")
	}
}
