package gqnode

import (
   "errors"
   "github.com/luisfurquim/goose"
   "github.com/PuerkitoBio/goquery"
)

type GQNode struct {
   goquery.Selection
}

type GooseG struct {
   Init goose.Alert
   Generate goose.Alert
}

var Goose GooseG

var ErrParsing error = errors.New("Error parsing html data")

