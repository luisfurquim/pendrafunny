package html2pdf

import (
   "golang.org/x/net/html"
)

func Handler_Body(cnv *Converter, sel *html.Node) {
   cnv.Pdf.AddPage()
   cnv.convert(sel)
}

