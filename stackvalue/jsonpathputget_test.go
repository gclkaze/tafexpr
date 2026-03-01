package stackvalue

import (
	"encoding/json"
	"testing"

	"github.com/Jeffail/gabs"
	"github.com/gclkaze/evalang/evalangparser/globals"
	"github.com/gclkaze/evalang/evalangparser/utils"
	"github.com/mdaverde/jsonpath"
	"github.com/stretchr/testify/assert"
)

func getJSON(t *testing.T) globals.JSONStruct {
	sample := utils.ReadFile("samplejsons\\sample.json")
	var payload globals.JSONStruct
	err := json.Unmarshal([]byte(*sample), &payload)
	assert.Nil(t, err)
	return payload
}

func getPlainJSON(t *testing.T) string {
	sample := utils.ReadFile("samplejsons\\sample.json")
	return *sample
}

func TestJSONPathGetSimple(t *testing.T) {

	payload := getJSON(t)
	value, err := jsonpath.Get(payload, "glossary.GlossDiv.GlossList.GlossEntry.GlossTerm")
	assert.Nil(t, err)
	assert.Equal(t, "Standard Generalized Markup Language", value)

	value, err = jsonpath.Get(payload, "glossary.GlossDiv.GlossList.GlossEntry.result")
	assert.Nil(t, err)
	assert.Equal(t, true, value)

	value, err = jsonpath.Get(payload, "glossary.GlossDiv.GlossList.GlossEntry.myId")
	assert.Nil(t, err)
	assert.Equal(t, float64(666), value)

	value, err = jsonpath.Get(payload, "glossary.GlossDiv.GlossList.GlossEntry.myFloatId")
	assert.Nil(t, err)
	assert.Equal(t, 666.666, value)

	value, err = jsonpath.Get(payload, "glossary.GlossDiv.GlossList.GlossEntry.GlossDef.GlossSeeAlso")
	assert.Nil(t, err)
	assert.Equal(t, []interface{}{"GML", "XML"}, value)

	value, err = jsonpath.Get(payload, "glossary.GlossDiv.GlossList.GlossEntry.GlossDef.GlossSeeAlso[1]")
	assert.Nil(t, err)
	assert.Equal(t, "XML", value)

	value, err = jsonpath.Get(payload, "glossary.GlossDiv.GlossList.GlossEntry.GlossTerm.DoesNotExist")
	assert.NotNil(t, err)
	assert.Nil(t, value)

	value, err = jsonpath.Get(payload, "glossary.GlossDiv.GlossList.GlossEntry.GlossDef.GlossSeeAlso[500]")
	assert.NotNil(t, err)
	assert.Nil(t, value)

	value, err = jsonpath.Get(payload, "glossary.GlossDiv[1].GlossList.GlossEntry.GlossDef.GlossSeeAlso[500]")
	assert.NotNil(t, err)
	assert.Nil(t, value)

	value, err = jsonpath.Get(payload, "glosdoesnotsary.GlossDiv[1].GlossList.GlossEntry.GlossDef.GlossSeeAlso[500]")
	assert.NotNil(t, err)
	assert.Nil(t, value)

}

func TestJSONPathGetSimpleGabs(t *testing.T) {

	payload := getPlainJSON(t)
	jsonParsed, err := gabs.ParseJSON([]byte(payload))
	assert.Nil(t, err)

	value := jsonParsed.Path("glossary.GlossDiv.GlossList.GlossEntry.GlossTerm")
	assert.Equal(t, "Standard Generalized Markup Language", value.Data())

	value = jsonParsed.Path("glossary.GlossDiv.GlossList.GlossEntry.result")
	assert.Equal(t, true, value.Data())

	value = jsonParsed.Path("glossary.GlossDiv.GlossList.GlossEntry.myId")
	assert.Equal(t, float64(666), value.Data())

	value = jsonParsed.Path("glossary.GlossDiv.GlossList.GlossEntry.myFloatId")
	assert.Equal(t, 666.666, value.Data())
	/*
		value = jsonParsed.Path("glossary.GlossDiv.GlossList.GlossEntry.GlossDef.GlossSeeAlso")
		assert.Equal(t, []interface{}{"GML", "XML"}, value.Data())

			value = jsonParsed.Search("glossary", "GlossDiv", "GlossList", "GlossEntry", "GlossDef", "GlossSeeAlso", "0")
			assert.Equal(t, "XML", value.Data())

			vv := jsonParsed.Path("glossary.GlossDiv.GlossList.GlossEntry.GlossDef.GlossSeeAlso.a")
			assert.Equal(t, "XML", vv.Data())

			value = jsonParsed.Path("glossary.GlossDiv.GlossList.GlossEntry.GlossTerm.DoesNotExist")
			assert.Nil(t, value.Data())

			value = jsonParsed.Path("glossary.GlossDiv.GlossList.GlossEntry.GlossDef.GlossSeeAlso[500]")
			assert.Nil(t, value.Data())

			value = jsonParsed.Path("glossary.GlossDiv[1].GlossList.GlossEntry.GlossDef.GlossSeeAlso[500]")
			assert.Nil(t, value.Data())

			value = jsonParsed.Path("glosdoesnotsary.GlossDiv[1].GlossList.GlossEntry.GlossDef.GlossSeeAlso[500]")
			assert.Nil(t, value.Data())*/

}

func TestJSONPathGetArrayOfArrays(t *testing.T) {
	payload := getJSON(t)

	value, err := jsonpath.Get(payload, "glossary.GlossDiv.GlossList.GlossEntry.GlossDef.arrayOfArrays.0.0.1")
	assert.Nil(t, err)
	assert.Equal(t, float64(2), value)

	value, err = jsonpath.Get(payload, "glossary.GlossDiv.GlossList.GlossEntry.GlossDef.arr[3].id")
	assert.Nil(t, err)
	assert.Equal(t, float64(4), value)

	value, err = jsonpath.Get(payload, "glossary.GlossDiv.GlossList.GlossEntry.GlossDef.arr[3].data.id[0]")
	assert.Nil(t, err)
	assert.Equal(t, float64(4), value)

	value, err = jsonpath.Get(payload, "glossary.GlossDiv.GlossList.GlossEntry.GlossDef.arr[3].data.id[4].another")
	assert.Nil(t, err)
	assert.Equal(t, "lol", value)

	value, err = jsonpath.Get(payload, "glossary.GlossDiv.GlossList.GlossEntry.GlossDef.arr.3.data.id.4.another")
	assert.Nil(t, err)
	assert.Equal(t, "lol", value)

	value, err = jsonpath.Get(payload, "glossary              .GlossDiv           .GlossList.GlossEntry            .GlossDef             .arr       .      3             .           data      .id   . 4 .another  ")
	assert.NotNil(t, err)
	assert.Equal(t, nil, value)

}

func TestJSONPathUpdateSimple(t *testing.T) {
	payload := getJSON(t)
	path := "glossary.GlossDiv.GlossList.GlossEntry.GlossTerm"
	pvalue := 123

	value, err := jsonpath.Get(payload, path)
	assert.Nil(t, err)
	assert.Equal(t, "Standard Generalized Markup Language", value)

	err = jsonpath.Set(&payload, path, pvalue)
	assert.Nil(t, err)

	value, err = jsonpath.Get(payload, path)
	assert.Nil(t, err)
	assert.Equal(t, pvalue, value)
}

func TestJSONPathWriteInexistingPath(t *testing.T) {
	payload := getJSON(t)
	path := "glossary.myUnknownPath.anotherUnknownPath"
	pvalue := 123

	_, err := jsonpath.Get(payload, path)
	assert.NotNil(t, err)

	err = jsonpath.Set(&payload, path, pvalue)
	assert.Nil(t, err)

	value, err := jsonpath.Get(payload, path)
	assert.Nil(t, err)
	assert.Equal(t, pvalue, value)
}

/*
//superfail
func TestJSONPathWriteInexistingPathInArray(t *testing.T) {
	payload := getJSON(t)
	path := "glossary.GlossDiv.GlossList.GlossEntry.GlossDef.GlossSeeAlso[3]"
	pvalue := 123

	_, err := jsonpath.Get(payload, path)
	assert.NotNil(t, err)

	err = jsonpath.Set(&payload, path, pvalue)
	assert.Nil(t, err)

	value, err := jsonpath.Get(payload, path)
	assert.Nil(t, err)
	assert.Equal(t, pvalue, value)
}*/

func TestJSONPathDelete(t *testing.T) {
	payload := getPlainJSON(t)
	path := "glossary.GlossDiv.GlossList"

	jsonParsed, err := gabs.ParseJSON([]byte(payload))
	assert.Nil(t, err)

	ok := jsonParsed.ExistsP(path)
	assert.True(t, ok)

	err = jsonParsed.DeleteP(path)
	assert.Nil(t, err)

	ok = jsonParsed.ExistsP(path)
	assert.False(t, ok)
}

/*func TestJSONPathDeleteObject(t *testing.T) {
	payload := getPlainJSON(t)
	path := "/glossary/GlossDiv/GlossList/GlossEntry/GlossDef/GlossSeeAlso/0"

	jsonParsed, err := gabs.ParseJSON([]byte(payload))
	assert.Nil(t, err)

	err = jsonParsed.Delete(path)
	assert.Nil(t, err)

	ok := jsonParsed.ExistsP(path)
	assert.False(t, ok)
}
*/
/*func TestPathElems(t *testing.T) {
	u := utils.NewJSONUtils()
	path := "$a.b.c.d[1][2][3].a.b[4]"
	res, paths := u.GetJSONVariableExpressionTokens(&path)
	assert.True(t, res)
	assert.NotNil(t, paths)

	assert.Equal(t, 0, (*paths)[0].Index)
	assert.Equal(t, globals.VARIABLE_DECL, (*paths)[0].Type)
	assert.Equal(t, "$a", (*paths)[0].Value)

	assert.Equal(t, 0, (*paths)[1].Index)
	assert.Equal(t, globals.PROPERTY, (*paths)[1].Type)
	assert.Equal(t, "b", (*paths)[1].Value)

	assert.Equal(t, 0, (*paths)[2].Index)
	assert.Equal(t, globals.PROPERTY, (*paths)[2].Type)
	assert.Equal(t, "c", (*paths)[2].Value)

	assert.Equal(t, 1, (*paths)[3].Index)
	assert.Equal(t, globals.INDEX, (*paths)[3].Type)
	assert.Equal(t, "d", (*paths)[3].Value)

	assert.Equal(t, 2, (*paths)[4].Index)
	assert.Equal(t, globals.INDEX, (*paths)[4].Type)
	assert.Equal(t, "", (*paths)[4].Value)

	assert.Equal(t, 3, (*paths)[5].Index)
	assert.Equal(t, globals.INDEX, (*paths)[5].Type)
	assert.Equal(t, "", (*paths)[5].Value)

	assert.Equal(t, 0, (*paths)[6].Index)
	assert.Equal(t, globals.PROPERTY, (*paths)[6].Type)
	assert.Equal(t, "a", (*paths)[6].Value)

	assert.Equal(t, 1, (*paths)[7].Index)
	assert.Equal(t, globals.INDEX, (*paths)[7].Type)
	assert.Equal(t, "d", (*paths)[7].Value)
}
*/
