package pendrafusion

import (
   "errors"
   "regexp"
   "golang.org/x/net/html"
   "golang.org/x/net/html/atom"
   "github.com/phpdave11/gofpdf"
   "github.com/luisfurquim/goose"
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
   Doc Node
   Pdf *gofpdf.Fpdf
   UseScript bool
   Format []Format
}

type Node interface {
   Attr(string) (string, bool)
//   Attrs() []Attribute
   FirstChild() Node
   NextSibling() Node
   Type() html.NodeType
   Data() string
   DataAtom() atom.Atom
   Find(string) []Node
}

type Attribute interface {
   Key() string
   Val() string
}

type GooseG struct {
   Init goose.Alert
   Generate goose.Alert
}

var Goose GooseG

var reRGB *regexp.Regexp = regexp.MustCompile(`rgb\(([0-9]+),([0-9]+),([0-9]+)\)`)
var reHexColor *regexp.Regexp = regexp.MustCompile(`\#([0-9a-fA-F]{2})([0-9a-fA-F]{2})([0-9a-fA-F]{2})`)

var Handler map[atom.Atom]func(*Converter, Node)

var reSpc *regexp.Regexp = regexp.MustCompile(`\s+`)

var ErrGenerating error = errors.New("Error generating pdf data")

