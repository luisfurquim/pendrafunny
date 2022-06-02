#!/bin/bash

GOOS=js GOARCH=wasm go build -o  build/pendrafusion.wasm build/pendrafusion.go
