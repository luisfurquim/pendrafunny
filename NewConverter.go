package pendrafunny

import (
   "io"
   "github.com/phpdave11/gofpdf"
   "github.com/PuerkitoBio/goquery"
)


func NewConverter(in io.Reader) (*Converter, error) {
   var cnv Converter
   var err error

   cnv.Doc, err = goquery.NewDocumentFromReader(in)
   if err != nil {
      Goose.Init.Logf(1,"%s: %s", ErrParsing, err)
      return nil, err
   }

   cnv.Pdf = gofpdf.New("P", "mm", "A4", "")
   cnv.Pdf.AddUTF8Font("notoSans", "", "./fonts/NotoSans-Regular.ttf")
   cnv.Pdf.AddUTF8Font("notoSans", "B", "./fonts/NotoSans-Bold.ttf")

/*
   cnv.Pdf.AddUTF8Font("notoSans", "I", "./fonts/NotoSans-Italics.ttf")
   Goose.Generate.Logf(1, "H: %f", cnv.CurrHeight)
   cnv.Pdf.AddUTF8Font("notoSans", "BI", "./fonts/NotoSans-BoldItalics.ttf")
   Goose.Generate.Logf(1, "H: %f", cnv.CurrHeight)
*/

   cnv.Format = []Format{Format{
      FontFace: "notoSans",
      FontSize: 12,
   }}

   cnv.Format[0].Box.W, cnv.Format[0].Box.H = cnv.Pdf.GetPageSize()
   _, cnv.Format[0].Box.LH = cnv.Pdf.GetFontSize()

   cnv.Apply()

   return &cnv, nil
}

