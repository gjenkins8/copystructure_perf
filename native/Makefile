.DEFAULT: build
.PHONY: build vet run

PKG_SOURCE_FILES=$(shell go list -f '{{ $$dir := .Dir }}{{ range .GoFiles }}{{ printf "%s/%s " $$dir . }}{{ end }}' . )

copystructure_driver: $(PKG_SOURCE_FILES)
	go build -o copystructure_driver .

build: copystructure_driver

vet:
	go vet .

run:
	./copystructure_driver