package elti

import "testing"

func TestElti(t *testing.T) {
	m := NewMap()
	m.Set("name", NewDataFromString("bob"))
	m.Set("age", NewDataFromUint8(24))
	m.Set("varint-test", NewDataFromVarint(15))
	arr := NewArray()
	arr.PushBack(NewDataFromString("math"))
	arr.PushBack(NewDataFromString("english"))
	m.Set("source", arr)
	e := NewElti(m)
	buf := e.SeriToBytes()

	e2 := ParseToElti(buf)
	if e2.GetRoot().GetValueType() != MAP {
		t.Error("elti test error.")
	}
	if e2.GetRoot().Attr("source").At(0).GetAsString() != "math" {
		t.Error("elti test error.")
	}
	if e2.GetRoot().Attr("varint-test").GetAsVarint() != 15 {
		t.Error("elti test error.")
	}

	p := ParseToPositioner(buf)
	if p.Attr("name").GetAsString() != "bob" {
		t.Error("elti test error.")
	}
	if p.Attr("source").At(0).GetAsString() != "math" {
		t.Error("elti test error.")
	}
	if p.Attr("varint-test").GetAsVarint() != 15 {
		t.Error("elti test perror.")
	}
}
