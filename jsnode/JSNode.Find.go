package jsnode

import (
   "syscall/js"
   "github.com/luisfurquim/pendrafusion"
)

func (n JSNode) Find(string) []pendrafusion.Node {
   var a js.Value
   var l, i int
   var node JSNode
   var nodes []pendrafusion.Node

   a = n.Node.Get("querySelectorAll")
   l = a.Get("length").Int()
   if l == 0 {
      return []pendrafusion.Node{}
   }

   nodes = make([]pendrafusion.Node, l)

   for i=0; i<l; i++ {
      node.Node = a.Call("item", i)
      node.CompSty = js.Global().Call("getComputedStyle", node.Node)
      nodes = append(nodes, node)
   }

   return nodes

}
