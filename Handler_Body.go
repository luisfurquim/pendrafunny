package pendrafusion

func Handler_Body(cnv *Converter, sel Node) {
   cnv.Pdf.AddPage()
   cnv.convert(sel)
}

