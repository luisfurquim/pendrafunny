package jsnode

func (n JSNode) Data() string {
   return n.Node.Get("nodeValue").String()
}
