package pendrafusion

func Handler_SCRIPT(cnv *Converter, sel Node) {
   if cnv.UseScript {
      Goose.Generate.Logf(1, "SCRIPT: %s", sel.Data())
      cnv.convert(sel)
   }
}

