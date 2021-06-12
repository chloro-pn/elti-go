package elti

import "testing"

func TestEltiType(t *testing.T) {
	str := elti_string
	if str != 0x01 {
		t.Error("elti_type error!")
	}

	str = elti_bytes
	if str != 0x00 {
		t.Error("elti_type error!")
	}
}

func TestValueType(t *testing.T) {
	ty := MAP
	if ty != 0 {
		t.Error("value_type error!")
	}
	ty = ARRAY
	if ty != 1 {
		t.Error("value_type error!")
	}
	ty = DATA
	if ty != 2 {
		t.Error("value_type error!")
	}
}
