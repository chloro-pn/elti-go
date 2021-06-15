package elti

import (
	"bytes"
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	arr := NewArray()
	arr.PushBack(NewDataFromInt8(int8(2)))
	arr.PushBack(NewData([]byte{0x00, 0x01, 0x02}))
	arr.PushBack(NewDataFromString("hello"))
	buf := []byte{}
	buf = arr.seriValue(buf)

	arr2 := NewArray()
	total_size := arr2.parseValue(buf, 0, ParseRefOff)
	fmt.Printf("%d %d", total_size, len(buf))
	if total_size != uint32(len(buf)) {
		t.Error("test array error.")
	}
	if arr2.Size() != arr.Size() {
		t.Error("test array error.")
	}
	if arr2.At(0).getValueType() != DATA {
		t.Error("test array error.")
	}
	v, ok := arr2.At(0).(*Data)
	if !ok {
		t.Error("test array error.")
	}
	n := bytes_to_int8(v.BytesRef())
	if n != 2 {
		t.Error("test array error.")
	}
	if !bytes.Equal([]byte{0x00, 0x01, 0x02}, arr2.At(1).(*Data).BytesRef()) {
		t.Error("test array error.")
	}
	if string(arr2.At(2).(*Data).BytesRef()) != "hello" {
		t.Error("test array error.")
	}
	arr2.Erase(0)
	if arr2.Size() != 2 {
		t.Error("test array error.")
	}
	if !bytes.Equal([]byte{0x00, 0x01, 0x02}, arr2.At(0).(*Data).BytesRef()) {
		t.Error("test array error.")
	}
}
