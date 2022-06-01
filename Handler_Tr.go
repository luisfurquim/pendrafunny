package html2pdf

import (
   "golang.org/x/net/html"
)

func Handler_Tr(cnv *Converter, sel *html.Node) {
   Goose.Generate.Logf(5, "Tr: %s", sel.Data)
   cnv.convert(sel)
}

