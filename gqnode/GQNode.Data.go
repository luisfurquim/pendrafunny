package gqnode

import (
   "github.com/PuerkitoBio/goquery"
)


func (n GQNode) Data() string {
   if len(n.Nodes) > 0 {
      return n.Nodes[0].Data
   }

   return ""
}
