package main

import "github.com/chloro-pn/elti-go"
import "fmt"

func main() {
	v := elti.NewDataFromString("hello")
	v2 := elti.NewDataFromString("world")
	arr := elti.NewArray()
	arr.PushBack(v)
	arr.PushBack(v2)
	el := elti.NewElti(arr)
	buf := el.SeriToBytes()

	el2 := elti.ParseToElti(buf)
	root := el2.GetRoot()
	fmt.Printf("%s, %s!", root.At(0).GetAsString(), root.At(1).GetAsString())
}
