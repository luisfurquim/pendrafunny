package jsnode

import (
   "fmt"
   "syscall/js"
   "github.com/luisfurquim/pendrafusion"
)


func ToPDF(this js.Value, p []js.Value) interface{} {
   var n JSNode

   if len(p) == 0 {
      n.Node = js.Global().Get("document").Call("querySelector", "HTML")
   } else {
      if p[0].InstanceOf(js.ValueOf("string")) {
         n.Node = js.Global().Get("document").Call("querySelector", p[0])
      } else {
         // Assumes that if it has this property then it should be a DOM Node
         // obviously it not a PROOF that it really is, but it should be enough
         // and is the caller responsability to call it with proper values...
         // so, call it with junk and junk is what you'll get
         if p[0].Get("firstChild").IsUndefined() {
            return nil
         }
         n.Node = p[0]
      }
   }

   n.CompSty = js.Global().Call("getComputedStyle", n.Node)

   return js.Global().Get("Promise").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
      var resolve, reject js.Value
      var buf pendrafusion.BufferReader
      var err error

      if len(args) != 2 {
         return nil
      }

      resolve = args[0]
      reject = args[1]

      go func() {
         buf, err = pendrafusion.NewConverter(n).Convert()
         if err != nil {
            reject.Invoke(js.Global().Get("Error").New(fmt.Sprintf("Error converting HTML node: %s", err)))
         } else {
            resolve.Invoke(string(buf.Bytes()))
         }
      }()

      return nil
   }))
}




/*
type Node interface {
//   Attrs() []Attribute
   Find(string) []Node
}

*/