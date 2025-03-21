# copystructure_perf
Performance comparison  of `copystructure.Copy(...)` between Wasm and native builds

## Basic usage

To build/execute both:
```
./run.sh
```

Runs both Wasm and native builds of a simple usage of `copystructure.Copy(...)` to show the performance differences between the two.

## Details

See the `wasm/` and `native/` directories for the respective code.

Both implement a simple usage of reading a "largish" dataset from `testdata.json`, then using `copystructure.Copy(...)` to duplicate this data.
The Wasm build tends to be significantly slower.


## Notes

Note: `github.com/mitchellh/copystructure` is deprecated.
The performance difference between Wasm and native code was noticed in existing code using `copystructure.Copy(...)`

This repo continues to use `copystructure.Copy(..)` to identify the fundamental difference why Wasm builds may be significantly slow than native.
