package jsnode

import (
   "golang.org/x/net/html/atom"
)

func (n JSNode) DataAtom() atom.Atom {
   return atom.Lookup([]byte(n.Node.Get("tagName").String()))
}
