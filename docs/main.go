package main

import (
	"bufio"
	"strings"
	"syscall/js"

	"github.com/matsuyoshi30/godiff"
)

func showDiff(this js.Value, i []js.Value) interface{} {
	document := js.Global().Get("document")

	clearTable()

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

	if scanner1.Scan() {
		d, _ := godiff.NewDiff(scanner1.Text(), "")
		setDiff(d)
		for scanner1.Scan() {
			d, _ := godiff.NewDiff(scanner1.Text(), "")
			setDiff(d)
		}
	} else if scanner2.Scan() {
		d, _ := godiff.NewDiff("", scanner2.Text())
		setDiff(d)
		for scanner2.Scan() {
			d, _ := godiff.NewDiff("", scanner2.Text())
			setDiff(d)
		}
	}

	return nil
}

func clearTable() {
	document := js.Global().Get("document")

	th := document.Call("createElement", "th")
	tr := document.Call("createElement", "tr")
	tr.Call("appendChild", th)

	out1 := document.Call("getElementById", "output1")
	out1.Set("textContent", "")
	out1.Call("appendChild", tr)
	out2 := document.Call("getElementById", "output2")
	out2.Set("textContent", "")
	out2.Call("appendChild", tr)
}

func setDiff(d *godiff.Diff) {
	diffs := d.ShowDiff()

	lstr := emphasisStr(diffs[0], diffs[2])
	appendTableData("output1", lstr)
	rstr := emphasisStr(diffs[1], diffs[2])
	appendTableData("output2", rstr)
}

func emphasisStr(s, diff string) string {
	var retStr string
	for i := 0; i < len(s); i++ {
		if diff[i] != '-' {
			retStr += "<em>" + string(s[i]) + "</em>"
		} else {
			retStr += string(s[i])
		}
	}
	return retStr
}

func appendTableData(id, s string) {
	document := js.Global().Get("document")

	td := document.Call("createElement", "td")
	td.Set("innerHTML", s)
	tr := document.Call("createElement", "tr")
	tr.Set("className", "data"+id)
	tr.Call("appendChild", td)
	document.Call("getElementById", id).Call("appendChild", tr)
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
