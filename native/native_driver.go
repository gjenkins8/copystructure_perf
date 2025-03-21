package main

import (
	"fmt"
	"time"

	"github.com/gjenkins8/copystructure_perf/copystructuretest"
)

func main() {

	testData := copystructuretest.LoadTestData()

	start := time.Now()
	result, err := copystructuretest.DoCopyStructure(testData)
	if err != nil {
		panic(err)
	}
	end := time.Now()
	fmt.Printf("DoCopyStructure len(r)=%d duration=%s\n", len(result), end.Sub(start))
}
