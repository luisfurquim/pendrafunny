package pendrafusion

import (
   "syscall/js"
)

func Convert(stop chan struct{}) (func(this js.Value, p []js.Value) interface{}) {
   return func(this js.Value, p []js.Value) interface{} {

     handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
         var dom, resolve, reject js.Value

         resolve = args[0]
         reject = args[1]

         if len(p) == 0 {
            dom = js.Global().Get("document")
         } else if len(p) == 1 {
            dom = p[0]
         } else {
            reject.Invoke(js.Global().Get("Error").New("Wrong parameter count! Needs just a dom object (none means document)"))
            return nil
         }

_ = dom
_ = resolve

         return nil
      })

      return js.Global().Get("Promise").New(handler)
   }
}

