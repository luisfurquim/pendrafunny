package main

import (
   "syscall/js"
   "github.com/luisfurquim/pendrafusion"
)


func main() {
   var stop chan struct{}
   stop = make(chan struct{})
   js.Global().Set("pendrafusion", js.FuncOf(pendrafusion.Convert(stop)))
   <-stop
}

