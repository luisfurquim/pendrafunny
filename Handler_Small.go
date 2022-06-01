package html2pdf

import (
   "golang.org/x/net/html"
)

func Handler_Small(cnv *Converter, sel *html.Node) {
   cnv.Format[len(cnv.Format)-1].FontSize = cnv.GetFontSize("2")
   cnv.Apply()
   cnv.convert(sel)
}

