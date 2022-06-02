package pendrafusion

import (
   "math"
)

// Arbitrary values, feel free to fix it to more realist values
var FontSizes []float64 = []float64{
   0,   // invalid value
   6,
   9,
   12,
   16,
   20,
   26,
   32,
}

func (cnv *Converter) GetFontSize(txt string) float64 {
   var i byte
   var ndx int
   var curr float64

   if txt == "" {
      return 0
   }

   i = txt[len(txt)-1]
   if i<49 || i>57 {
      return 0
   }

   i -= 48

   curr = cnv.Format[len(cnv.Format)-1].FontSize
   for ; ndx<len(FontSizes); ndx++ {
      if FontSizes[ndx]>=curr {
         break
      }
   }

   if ndx == len(FontSizes) {
      ndx--
      Goose.Generate.Fatalf(0,"curr: %f, new: %s", curr, txt)
   }

   if math.Abs(FontSizes[ndx]-curr) > math.Abs(FontSizes[ndx-1]-curr) {
      ndx --
   }


   if txt[0] == '+' {
      if (ndx+int(i)) > 7 {
         return FontSizes[7]
      }
      return FontSizes[ndx+int(i)]
   } else if txt[0] == '-' {
      if (ndx-int(i)) < 1 {
         return FontSizes[1]
      }
      return FontSizes[ndx-int(i)]
   }

   return FontSizes[int(i)]
}

