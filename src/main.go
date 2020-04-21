package main

import (
	"github.com/gopherjs/gopherjs/js"
)

var indexHTML = string(MustAsset("assets/index.html"))

func main() {
	js.Global.Get("document").Call("write", indexHTML)
	inputTextArea().Set("oninput", onInput)
}

func inputTextArea() *js.Object {
	return js.Global.Get("document").Call("querySelector", "#input")
}

func outputDiv() *js.Object {
	return js.Global.Get("document").Call("querySelector", "#output")
}

func onInput() {
	content := inputTextArea().Get("value").String()
	outputDiv().Set("innerText", content)
}
