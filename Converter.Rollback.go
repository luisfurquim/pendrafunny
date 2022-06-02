package pendrafusion

func (cnv *Converter) Rollback() {
   var oldFmt, newFmt *Format
   var i int

   i = len(cnv.Format)
   newFmt = &cnv.Format[i-1]
   oldFmt = &cnv.Format[i-2]


   if newFmt.FontFace!=oldFmt.FontFace ||
      newFmt.TextDecoration!=oldFmt.TextDecoration ||
      newFmt.FontSize!=oldFmt.FontSize {
      cnv.Pdf.SetFont(oldFmt.FontFace, oldFmt.TextDecoration, oldFmt.FontSize)
   }

   if newFmt.FontColor != oldFmt.FontColor {
      cnv.Pdf.SetTextColor(oldFmt.FontColor.R,oldFmt.FontColor.G,oldFmt.FontColor.B)
   }
}

