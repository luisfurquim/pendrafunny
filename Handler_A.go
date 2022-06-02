package pendrafusion

import (
   "golang.org/x/net/html"
)

func Handler_A(cnv *Converter, sel Node) {
   var f *Format
   var val string
   var ok bool

   if sel.FirstChild() == nil {
      return
   }

   if sel.FirstChild().Type() == html.TextNode {
      f = &(cnv.Format[len(cnv.Format)-1])

      if val, ok = sel.Attr("href"); ok {
         f.LinkStr = val
      }

      if f.LinkStr != "" {
         f.TextDecoration = SetTextDecoration(f.TextDecoration, 'U')
         f.FontColor = Color{R:0,G:0,B:255}
         cnv.Apply()
      } else {
         cnv.Pdf.Write(f.Box.LH, sel.FirstChild().Data())
      }

      Goose.Generate.Logf(5, "A: %s", sel.Data())
      cnv.convert(sel)
   }
}

