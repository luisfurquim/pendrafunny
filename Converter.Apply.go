package pendrafusion

func (cnv *Converter) Apply() {
   var oldFmt, newFmt *Format
   var i int

   i = len(cnv.Format)
   newFmt = &cnv.Format[i-1]

   if i>1 {
      oldFmt = &cnv.Format[i-2]

      if newFmt.FontFace!=oldFmt.FontFace ||
         newFmt.TextDecoration!=oldFmt.TextDecoration ||
         newFmt.FontSize!=oldFmt.FontSize {
         cnv.Pdf.SetFont(newFmt.FontFace, newFmt.TextDecoration, newFmt.FontSize)
      }

      if newFmt.FontColor != oldFmt.FontColor {
         cnv.Pdf.SetTextColor(newFmt.FontColor.R,newFmt.FontColor.G,newFmt.FontColor.B)
      }
   } else {
      cnv.Pdf.SetFont(newFmt.FontFace, newFmt.TextDecoration, newFmt.FontSize)
      cnv.Pdf.SetTextColor(newFmt.FontColor.R,newFmt.FontColor.G,newFmt.FontColor.B)
   }
}

