package pendrafusion

func Handler_Small(cnv *Converter, sel Node) {
   cnv.Format[len(cnv.Format)-1].FontSize = cnv.GetFontSize("2")
   cnv.Apply()
   cnv.convert(sel)
}

