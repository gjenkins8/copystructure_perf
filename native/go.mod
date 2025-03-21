module github.com/gjenkins8/copystructure_wasm_repo/native

go 1.24.0

replace github.com/gjenkins8/copystructure_perf/copystructuretest => ../copystructuretest

require github.com/gjenkins8/copystructure_perf/copystructuretest v0.0.0

require (
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
)
