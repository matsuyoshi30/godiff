package main

import (
	"bufio"
	"strings"
	"syscall/js"

	"github.com/matsuyoshi30/godiff"
)

func showDiff(this js.Value, i []js.Value) interface{} {
	document := js.Global().Get("document")
	input1 := document.Call("getElementById", "input1")
	input2 := document.Call("getElementById", "input2")

	str1 := input1.Get("value").String()
	str2 := input2.Get("value").String()

	scanner1 := bufio.NewScanner(strings.NewReader(str1))
	scanner2 := bufio.NewScanner(strings.NewReader(str2))

	for scanner1.Scan() && scanner2.Scan() {
		d, _ := godiff.NewDiff(scanner1.Text(), scanner2.Text())
		setDiff(d)
	}

	// TODO: implement show diff when len(str1) is different from len(str2)
	// currently only use for same length str1 and str2

	return nil
}

func setDiff(d *godiff.Diff) {
	document := js.Global().Get("document")

	diffs := d.ShowDiff()
	td1 := document.Call("createElement", "td")
	td1.Set("innerHTML", diffs[0])
	tr1 := document.Call("createElement", "tr")
	tr1.Call("appendChild", td1)
	document.Call("getElementById", "output1").Call("appendChild", tr1)
	td2 := document.Call("createElement", "td")
	td2.Set("innerHTML", diffs[1])
	tr2 := document.Call("createElement", "tr")
	tr2.Call("appendChild", td2)
	document.Call("getElementById", "output2").Call("appendChild", tr2)

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
