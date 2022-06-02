package pendrafusion

func Handler_Tr(cnv *Converter, sel Node) {
   Goose.Generate.Logf(5, "Tr: %s", sel.Data())
   cnv.convert(sel)
}

