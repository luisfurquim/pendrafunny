package html2pdf

import (
   "golang.org/x/net/html"
)

func GetAttr(key string, sel *html.Node) (string, bool) {
   var a html.Attribute

   for _, a = range sel.Attr {
      if a.Key == key {
         return a.Val, true
      }
   }

   return "", false
}

