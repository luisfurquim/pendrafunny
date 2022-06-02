package pendrafusion

import (
   "fmt"
)

func Handler_Table(cnv *Converter, sel Node) {
   var f *Format
   var w float64
   var val string
   var ok bool

   f = &(cnv.Format[len(cnv.Format)-1])

   if val, ok = sel.Attr("width"); ok && len(val)>0 {
      if val[len(val)-1] == '%' {
         fmt.Sscanf(val[:len(val)-1], "%f", &w)
         f.Box.W = f.Box.W * w / 100
      } else if len(val)>=3 {
         if val[len(val)-2:] == "pt" {
            fmt.Sscanf(val[:len(val)-2], "%f", &w)
            f.Box.W = w/2.83464
         } else {
            fmt.Sscanf(val[:len(val)-2], "%f", &f.Box.W)
         }
      }

      if (f.Box.W + cnv.Pdf.GetX()) > cnv.Format[0].Box.W {
         cnv.Pdf.Write(f.Box.LH, "\n")
      }
   } else {
      cnv.Pdf.Write(f.Box.LH, "\n")
      //f.Box.W = f.Box.W - cnv.Pdf.GetX()
   }

   cnv.Apply()

   Goose.Generate.Logf(5, "Table: %s", sel.Data())
   cnv.convert(sel)
   cnv.Pdf.Write(f.Box.LH, "\n")
}

