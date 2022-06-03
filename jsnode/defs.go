package jsnode

import (
   "errors"
   "syscall/js"
   "github.com/luisfurquim/goose"
)

type JSNode struct {
   Node js.Value
   CompSty js.Value
}

type GooseG struct {
   Init goose.Alert
   Generate goose.Alert
}

var Goose GooseG

var ErrParsing error = errors.New("Error parsing html data")

