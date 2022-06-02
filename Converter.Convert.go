package pendrafusion

import (
   "golang.org/x/net/html"
)

func (cnv *Converter) convert(sel Node) {
   var h func(*Converter, Node)
   var ok bool
   var child Node
   var stack int

   for child=sel.FirstChild(); child!=nil; child=child.NextSibling() {
      if child.Type() == html.TextNode {
         Goose.Generate.Logf(5, "Text: %s", child.Data())
         cnv.AddText(child.Data())
      } else if child.Type() == html.ElementNode {
         if h, ok = Handler[child.DataAtom()]; ok {
            Goose.Generate.Logf(5, "has handler: %s", child.Data())
            stack = len(cnv.Format)
            cnv.Format = append(cnv.Format, cnv.Format[stack-1])
            h(cnv, child)
            cnv.Rollback()
            cnv.Format = cnv.Format[:stack]
         } else {
            Goose.Generate.Logf(1, "no handler: %s", child.Data())
            cnv.convert(child)
         }
      }
   }
}


func (cnv *Converter) Convert() (BufferReader, error) {
   var buf *Buffer
   var err error

   cnv.Pdf.SetFont("notoSans", "", 12)
   _, cnv.Format[0].Box.LH = cnv.Pdf.GetFontSize()

   Goose.Generate.Logf(0,"lineHt: %f", cnv.Format[0].Box.LH)

   cnv.convert(cnv.Doc.Find("HTML")[0])

   buf = NewBuffer()
   err = cnv.Pdf.Output(buf)
   if err != nil {
      Goose.Generate.Logf(1,"%s: %s", ErrGenerating, err)
      return nil, err
   }

   return buf, nil
}
