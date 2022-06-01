package html2pdf

import (
   "golang.org/x/net/html"
)

func Handler_Br(cnv *Converter, sel *html.Node) {
   cnv.Pdf.Write(cnv.Format[len(cnv.Format)-1].Box.LH, "\n")
}

