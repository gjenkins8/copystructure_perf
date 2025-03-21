# copystructure_perf
Performance comparison  of `copystructure.Copy(...)` between Wasm and native builds

## Basic usage

To build/execute both Wasm and native benchmarks:
```
./run.sh
```

## Details

Test loads a "largish" example data set from `testdata.json`, then uses `go test` benchmarks to measure the performance of `copystructure.Copy(...)`.
For both `GOOS=wasip1 GOARCH=wasm`, and native (`GOOS` and `GOARCH` unset) builds.

The Wasm build tends to be significantly slower.


## Notes

Note: `github.com/mitchellh/copystructure` is deprecated.
The performance difference between Wasm and native code was noticed in existing code using `copystructure.Copy(...)`

This repo continues to use `copystructure.Copy(..)` to identify the fundamental difference why Wasm builds may be significantly slow than native.
