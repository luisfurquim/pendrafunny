package gqnode

import (
   "io"
   "github.com/PuerkitoBio/goquery"
)


func New(rd io.Reader) (*GQNode, error) {
   var n GQNode

   n, err = goquery.NewDocumentFromReader(rd)
   if err != nil {
      Goose.Init.Logf(1,"%s: %s", ErrParsing, err)
      return nil, err
   }


   return n, nil
}
