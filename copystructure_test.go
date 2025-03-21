package copystructure_test

import (
	"testing"

	_ "embed"
	"encoding/json"

	"github.com/mitchellh/copystructure"
)

//go:embed testdata.json
var testDataJSON []byte

func loadTestData() map[string]any {
	var testData map[string]any
	if err := json.Unmarshal(testDataJSON, &testData); err != nil {
		panic(err)
	}

	return testData
}

func BenchmarkDoCopyStructure(b *testing.B) {
	testData := loadTestData()

	for b.Loop() {
		_, err := copystructure.Copy(testData)

		if err != nil {
			panic(err)
		}
	}
}
