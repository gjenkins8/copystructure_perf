package copystructuretest

import (
	_ "embed"
	"encoding/json"

	"github.com/mitchellh/copystructure"
)

//go:embed testdata.json
var testDataJSON []byte

func LoadTestData() map[string]any {
	var testData map[string]any
	if err := json.Unmarshal(testDataJSON, &testData); err != nil {
		panic(err)
	}

	return testData
}

// DoCopyStructure is a wrapper around copystructure.Copy
func DoCopyStructure(src any) (map[string]any, error) {
	result, err := copystructure.Copy(src)
	if err != nil {
		return nil, err
	}

	return result.(map[string]any), nil
}
