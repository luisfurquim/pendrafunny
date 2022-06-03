package jsnode

import (
   "golang.org/x/net/html"
)

func (n JSNode) Type() html.NodeType {
   return html.NodeType(n.Node.Get("nodeType").Int())
}
