#!/bin/bash

set -eu -o pipefail

printf "\n\n===== Running native build ====\n\n"
( GOOS= GOARCH= go test -benchmem -tags -shuffle=on -bench=. )

printf "\n\n===== Running Wasm build ======\n\n"
( GOOS=wasip1 GOARCH=wasm go test -benchmem -tags -shuffle=on -bench=. -exec="$(go env GOROOT)/lib/wasm/go_wasip1_wasm_exec" )
