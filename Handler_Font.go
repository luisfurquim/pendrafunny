package html2pdf

import (
   "fmt"
   "regexp"
   "strings"
   "golang.org/x/net/html"
)

var reRGB *regexp.Regexp = regexp.MustCompile(`rgb\(([0-9]+),([0-9]+),([0-9]+)\)`)
var reHexColor *regexp.Regexp = regexp.MustCompile(`\#([0-9a-fA-F]{2})([0-9a-fA-F]{2})([0-9a-fA-F]{2})`)

func Handler_Font(cnv *Converter, sel *html.Node) {
   var a html.Attribute
   var f *Format
   var parts []string
   var r, g, b int
   var newFont string
   var ok bool
   var sz float64

   f = &(cnv.Format[len(cnv.Format)-1])

   for _, a = range sel.Attr {
      switch a.Key {
      case "color":
         if parts = reRGB.FindStringSubmatch(a.Val); len(parts)>0 {
            fmt.Sscanf(parts[1],"%d",&r)
            fmt.Sscanf(parts[2],"%d",&g)
            fmt.Sscanf(parts[3],"%d",&b)
         } else if parts = reHexColor.FindStringSubmatch(a.Val); len(parts)>0 {
            fmt.Sscanf(parts[1],"%x",&r)
            fmt.Sscanf(parts[2],"%x",&g)
            fmt.Sscanf(parts[3],"%x",&b)
         } else {
         }
         f.FontColor = Color{R:r, G:g, B:b}
      case "face":
         if newFont, ok = FontFace[strings.ToLower(a.Val)]; ok {
            f.FontFace = newFont
         } else {
            f.FontFace = a.Val
         }
      case "size":
         sz = cnv.GetFontSize(a.Val)
         if sz>0 {
            f.FontSize = sz
         }
      }
   }

   cnv.Apply()

   cnv.convert(sel)
}

