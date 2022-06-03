#!/bin/bash

GOOS=js GOARCH=wasm go build -o  build/pendrafusionjs.wasm build/pendrafusionjs.go
