package main

import (
	"fmt"
	"syscall/js"

	"github.com/matsuyoshi30/godiff"
)

func showDiff(this js.Value, i []js.Value) interface{} {
	// get str
	str1 := getStr("input1")
	str2 := getStr("input2")

	// check diff
	d, err := godiff.NewDiff(str1, str2)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	diffs := d.ShowDiff()

	// set diff
	js.Global().Get("document").Call("getElementById", "output1").Set("textContent", diffs[0])
	js.Global().Get("document").Call("getElementById", "output2").Set("textContent", diffs[1])
	js.Global().Get("document").Call("getElementById", "output3").Set("textContent", diffs[2])

	return nil
}

func getStr(id string) string {
	return js.Global().Get("document").Call("getElementById", id).Get("value").String()
}

func registerCallbacks() {
	js.Global().Set("showDiff", js.FuncOf(showDiff))
}

var c chan struct{}

func init() {
	registerCallbacks()
	c = make(chan struct{}, 0)
}

func main() {
	<-c
}
