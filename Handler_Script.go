package html2pdf

import (
   "golang.org/x/net/html"
)

func Handler_SCRIPT(cnv *Converter, sel *html.Node) {
   if cnv.UseScript {
      Goose.Generate.Logf(1, "SCRIPT: %s", sel.Data)
      cnv.convert(sel)
   }
}

