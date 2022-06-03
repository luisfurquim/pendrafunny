package main

import (
   "syscall/js"
   "github.com/luisfurquim/pendrafusion/jsnode"
)


func main() {
   var stop chan struct{}
   stop = make(chan struct{})
   js.Global().Set("pendrafusion", js.FuncOf(jsnode.ToPDF))
   <-stop
}

