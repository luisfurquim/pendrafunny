package main

import (
   "syscall/js"
   "github.com/luisfurquim/pendrafunny"
)


func main() {
   var stop chan struct{}
   stop = make(chan struct{})
   js.Global().Set("pendrafunny", js.FuncOf(pendrafunny.Convert(stop)))
   <-stop
}

