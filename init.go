package pendrafusion

import (
   "golang.org/x/net/html/atom"
)

func init() {
   Handler = map[atom.Atom]func(*Converter, Node){
      atom.Header: Handler_Head,
      atom.Body: Handler_Body,
      atom.P: Handler_P,
      atom.A: Handler_A,
      atom.Small: Handler_Small,
      atom.Big: Handler_Big,
      atom.Br: Handler_Br,
      atom.Strike: Handler_Strike,
      atom.Table: Handler_Table,
      atom.Tr: Handler_Tr,
      atom.B: Handler_Strong,
      atom.Strong: Handler_Strong,
      atom.Font: Handler_Font,
      atom.Script: Handler_SCRIPT,
   }
}

