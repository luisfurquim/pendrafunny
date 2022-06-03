package jsnode

import (
   "syscall/js"
   "github.com/luisfurquim/pendrafusion"
)

func (n JSNode) FirstChild() pendrafusion.Node {
   var child JSNode

   child.Node = n.Node.Get("firstChild")
   child.CompSty = js.Global().Call("getComputedStyle", child.Node)
   return &child
}
