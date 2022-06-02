package pendrafusion

func Handler_P(cnv *Converter, sel Node) {
   var f *Format

   f = &(cnv.Format[len(cnv.Format)-1])
   f.Box.W = f.Box.W - cnv.Pdf.GetX()
   cnv.Apply()

   Goose.Generate.Logf(5, "P: %s", sel.Data())
   cnv.convert(sel)
   cnv.Pdf.Write(f.Box.LH, "\n")
}

