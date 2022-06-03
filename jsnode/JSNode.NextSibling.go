package jsnode

import (
   "syscall/js"
   "github.com/luisfurquim/pendrafusion"
)

func (n JSNode) NextSibling() pendrafusion.Node {
   var child JSNode

   child.Node = n.Node.Get("nextSibling")
   child.CompSty = js.Global().Call("getComputedStyle", child.Node)
   return &child
}
