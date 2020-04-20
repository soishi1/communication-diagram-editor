package main

import (
	"github.com/gopherjs/gopherjs/js"
)

var indexHTML = string(MustAsset("assets/index.html"))

func main() {
	js.Global.Get("document").Call("write", indexHTML)
}
