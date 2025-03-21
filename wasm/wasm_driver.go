package main

import (
	"fmt"
	"time"

	"github.com/gjenkins8/copystructure_perf/copystructuretest"
)

//go:wasmexport copy_structure_driver
func CopyStructureDriver() uint64 {

	testData := copystructuretest.LoadTestData()

	start := time.Now()
	result, err := copystructuretest.DoCopyStructure(testData)
	if err != nil {
		panic(err)
	}
	end := time.Now()
	fmt.Printf("DoCopyStructure len(r)=%d duration=%s\n", len(result), end.Sub(start))

	return 0
}

func main() {}
