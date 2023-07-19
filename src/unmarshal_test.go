package src

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

func TestUnmarshalJSONFromInterface(t *testing.T) {
	var jsonMap interface{}

	jsonString := `{
		"known":"I do",
		"unknown":"I don't"
	}`

	_ = json.Unmarshal([]byte(jsonString), &jsonMap)

	result := UnmarshalJSONFromInterface[MyStructure](jsonMap)

	assert.Equal(t, "I do", result.Known)
	assert.Equal(t, "I don't", result.Fields["unknown"])
}
func TestUnmarshalJSON(t *testing.T) {
	jsonString := `{
		"known":"I do",
		"unknown":"I don't"
	}`
	result := UnmarshalJSON[MyStructure]([]byte(jsonString))

	assert.Equal(t, "I do", result.Known)
	assert.Equal(t, "I don't", result.Fields["unknown"])
}

type MyStructure struct {
	Fields map[string]interface{} `json:"fields"`
	Known  string                 `json:"known"`
}
