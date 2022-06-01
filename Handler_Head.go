package html2pdf

import (
   "golang.org/x/net/html"
)

func Handler_Head(cnv *Converter, sel *html.Node) {
   cnv.convert(sel)
}

