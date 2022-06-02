package pendrafusion

func (cnv *Converter) AddText(txt string) {
   var f *Format

   txt = reSpc.ReplaceAllString(txt, " ")
   Goose.Generate.Logf(5, "TEXT: %s", txt)
   f = &(cnv.Format[len(cnv.Format)-1])
   if f.LinkStr != "" {
      cnv.Pdf.WriteLinkString(f.Box.LH, txt, f.LinkStr)
   } else {
      cnv.Pdf.Write(f.Box.LH, txt)
   }
}
