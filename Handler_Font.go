package pendrafusion

import (
   "fmt"
   "strings"
)

func Handler_Font(cnv *Converter, sel Node) {
   var val string
   var f *Format
   var parts []string
   var r, g, b int
   var newFont string
   var ok bool
   var sz float64

   f = &(cnv.Format[len(cnv.Format)-1])

   val, ok = sel.Attr("color")
   if ok {
      if parts = reRGB.FindStringSubmatch(val); len(parts)>0 {
         fmt.Sscanf(parts[1],"%d",&r)
         fmt.Sscanf(parts[2],"%d",&g)
         fmt.Sscanf(parts[3],"%d",&b)
      } else if parts = reHexColor.FindStringSubmatch(val); len(parts)>0 {
         fmt.Sscanf(parts[1],"%x",&r)
         fmt.Sscanf(parts[2],"%x",&g)
         fmt.Sscanf(parts[3],"%x",&b)
      } else {
      }
      f.FontColor = Color{R:r, G:g, B:b}
   }

   val, ok = sel.Attr("face")
   if ok {
      if newFont, ok = FontFace[strings.ToLower(val)]; ok {
         f.FontFace = newFont
      } else {
         f.FontFace = val
      }
   }

   val, ok = sel.Attr("size")
   if ok {
      sz = cnv.GetFontSize(val)
      if sz>0 {
         f.FontSize = sz
      }
   }

   cnv.Apply()
   cnv.convert(sel)
}

