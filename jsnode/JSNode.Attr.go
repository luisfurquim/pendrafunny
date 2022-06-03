package jsnode

import (
   "syscall/js"
)

func (n JSNode) Attr(attr string) (string, bool) {
   var a js.Value

   a = n.Node.Call("getAttribute", attr)
   if a.IsNull() {
      return "", false
   }
   return  a.String(), true
}
// getComputedStyle