package pendrafunny

import (
   "errors"
   "regexp"
   "golang.org/x/net/html"
   "golang.org/x/net/html/atom"
   "github.com/phpdave11/gofpdf"
   "github.com/luisfurquim/goose"
   "github.com/PuerkitoBio/goquery"
)

type Box struct {
   X, Y, W, H, LH float64
}

type Color struct {
   R, G, B int
}

type Format struct {
   Box Box
   FontFace string
   TextDecoration string
   FontSize float64
   FontColor Color
   TextAlign string
   LinkStr string
   LinkId int
}

type Converter struct {
   Doc *goquery.Document
   Pdf *gofpdf.Fpdf
   UseScript bool
   Format []Format
}

type Node interface {
}

type GooseG struct {
   Init goose.Alert
   Generate goose.Alert
}

var Goose GooseG

var Handler map[atom.Atom]func(*Converter, *html.Node)

var reSpc *regexp.Regexp = regexp.MustCompile(`\s+`)

var ErrParsing error = errors.New("Error parsing html data")
var ErrGenerating error = errors.New("Error generating pdf data")

