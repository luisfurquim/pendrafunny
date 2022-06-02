package pendrafusion

func Handler_Big(cnv *Converter, sel Node) {
   cnv.Format[len(cnv.Format)-1].FontSize = cnv.GetFontSize("5")
   cnv.Apply()
   cnv.convert(sel)
}

