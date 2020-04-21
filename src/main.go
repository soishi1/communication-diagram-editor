package main

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	"github.com/soishi1/communication-diagram-editor/src/source"
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
	parsed, err := source.Parse(content)
	if err != nil {
		outputDiv().Set("innerText", err.Error())
		return
	}
	outputDiv().Set("innerText", fmt.Sprintf("%+v\n", parsed))
}
