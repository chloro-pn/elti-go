package elti

import (
	"bytes"
	"encoding/base64"
	"testing"
)

func TestBase64Decode(t *testing.T) {
	var buf []byte = []byte("hello")
	var buff bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &buff)
	encoder.Write(buf)
	encoder.Close()
	if string(buff.Bytes()) != "aGVsbG8=" {
		t.Error("test base64 decode error.")
	}
}
