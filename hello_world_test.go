package elti

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	str := elti_hello_world()
	if str != "hello world, elti!" {
		t.Error("elti_hello_world error!")
	}
}
