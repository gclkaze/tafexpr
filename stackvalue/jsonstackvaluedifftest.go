package stackvalue

import (
	"testing"

	"github.com/gclkaze/evalang-globals/globals"
	"github.com/gclkaze/evalang-globals/utils"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestJsonDifferenceIdentical(t *testing.T) {
	jsonString := utils.ReadFile("..\\jsonexprs\\simple.json")

	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	var payload globals.JSONStruct
	err := json.Unmarshal([]byte(*jsonString), &payload)
	assert.Equal(t, nil, err)

	var payload2 globals.JSONStruct
	err = json.Unmarshal([]byte(*jsonString), &payload2)
	assert.Equal(t, nil, err)

	st1 := NewJSONStackValue(payload)
	st2 := NewJSONStackValue(payload2)

	assert.Equal(t, true, st1.Equals(st2))
}

func TestJsonDifferenceFirstSuperset(t *testing.T) {
	jsonStringSuper := utils.ReadFile("..\\jsonexprs\\simplesuperset.json")
	jsonString := utils.ReadFile("..\\jsonexprs\\simple.json")

	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	var payload globals.JSONStruct
	err := json.Unmarshal([]byte(*jsonStringSuper), &payload)
	assert.Equal(t, nil, err)

	var payload2 globals.JSONStruct
	err = json.Unmarshal([]byte(*jsonString), &payload2)
	assert.Equal(t, nil, err)

	st1 := NewJSONStackValue(payload)
	st2 := NewJSONStackValue(payload2)

	assert.Equal(t, false, st1.Equals(st2))

}

func TestJsonDifferenceFirstEmpty(t *testing.T) {
	jsonStringEmpty := "{}"
	jsonString := utils.ReadFile("..\\jsonexprs\\simple.json")

	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	var payload globals.JSONStruct
	err := json.Unmarshal([]byte(jsonStringEmpty), &payload)
	assert.Equal(t, nil, err)

	var payload2 globals.JSONStruct
	err = json.Unmarshal([]byte(*jsonString), &payload2)
	assert.Equal(t, nil, err)

	st1 := NewJSONStackValue(payload)
	st2 := NewJSONStackValue(payload2)

	assert.Equal(t, false, st1.Equals(st2))
}

func TestJsonDifferenceIdenticalButKeysInDifferentOrder(t *testing.T) {
	jsonStringSuper := utils.ReadFile("..\\jsonexprs\\simplealtered.json")
	jsonString := utils.ReadFile("..\\jsonexprs\\simple.json")

	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	var payload globals.JSONStruct
	err := json.Unmarshal([]byte(*jsonStringSuper), &payload)
	assert.Equal(t, nil, err)

	var payload2 globals.JSONStruct
	err = json.Unmarshal([]byte(*jsonString), &payload2)
	assert.Equal(t, nil, err)

	st1 := NewJSONStackValue(payload)
	st2 := NewJSONStackValue(payload2)

	assert.Equal(t, true, st1.Equals(st2))
}

func TestJSONDifferenceCompletelyDifferent(t *testing.T) {
	jsonStringFirst := utils.ReadFile("..\\jsonexprs\\simplealtered.json")
	jsonStringSecond := "{\"a\":1,\"b\":[1,2,32,4]}"

	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	var payload globals.JSONStruct
	err := json.Unmarshal([]byte(*jsonStringFirst), &payload)
	assert.Equal(t, nil, err)

	var payload2 globals.JSONStruct
	err = json.Unmarshal([]byte(jsonStringSecond), &payload2)
	assert.Equal(t, nil, err)

	st1 := NewJSONStackValue(payload)
	st2 := NewJSONStackValue(payload2)

	assert.Equal(t, false, st1.Equals(st2))

}
