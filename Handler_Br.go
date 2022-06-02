package pendrafusion

func Handler_Br(cnv *Converter, sel Node) {
   cnv.Pdf.Write(cnv.Format[len(cnv.Format)-1].Box.LH, "\n")
}

