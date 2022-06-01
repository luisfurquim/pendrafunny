#!/bin/bash

GOOS=js GOARCH=wasm go build -o  build/pendrafunny.wasm build/pendrafunny.go
