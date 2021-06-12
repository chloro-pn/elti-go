package elti

import (
	"fmt"
	"testing"
)

func TestSeriLength(t *testing.T) {
	var buf []byte
	var length uint32 = 2
	buf = seriLength(length, buf)
	if len(buf) != 1 || buf[0] != 2 {
		t.Error("seriLength error!")
	}
	buf = make([]byte, 0)
	length = 128
	buf = seriLength(length, buf)
	fmt.Println(len(buf))
	if len(buf) != 2 || buf[1] != 0x01 {
		t.Error("seriLength error!")
	}
}

func bufCheck(buf []byte) bool {
	//buf : [0111 1111 0000 0000 0000 0001]
	if len(buf) != 3 {
		return false
	}
	if buf[0] != 0x7F || buf[1] != 0x80 || buf[2] != 0x01 {
		return false
	}
	return true
}
func TestParseLength(t *testing.T) {
	buf := make([]byte, 0)
	buf = seriLength(uint32(127), buf)
	buf = seriLength(uint32(128), buf)
	// now buf : [0111 1111 0000 0000 0000 0001]
	if bufCheck(buf) == false {
		t.Error("parseLength error!")
	}
	length, new_begin := parseLength(buf, 0)
	if new_begin != 1 || length != 127 {
		t.Error("parseLength error!")
	}
	length, new_begin = parseLength(buf, new_begin)
	if new_begin != 3 || length != 128 {
		t.Error("parseLength error!")
	}
}
