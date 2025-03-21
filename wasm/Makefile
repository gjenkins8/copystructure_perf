.DEFAULT: build
.PHONY: build vet run

PKG_SOURCE_FILES=$(shell go list -f '{{ $$dir := .Dir }}{{ range .GoFiles }}{{ printf "%s/%s " $$dir . }}{{ end }}' . )

copystructure_driver.wasm: $(PKG_SOURCE_FILES)
	GOOS=wasip1 GOARCH=wasm go build -buildmode=c-shared -o copystructure_driver.wasm .

build: copystructure_driver.wasm

vet:
	GOOS=wasip1 GOARCH=wasm go vet .

run:
	extism call copystructure_driver.wasm --wasi --memory-max 4096 copy_structure_driver