package stackvalue

import "github.com/gclkaze/evalang-globals/globals"

type JSONBuilder struct {
	o globals.JSONStruct
	//parent interface{}
}

func NewJSONBuilder() *JSONBuilder {
	return &JSONBuilder{}
}

func (j *JSONBuilder) PutInteger(i int, key string) {
	j.o[key] = i
}

func (j *JSONBuilder) PutString(i string, key string) {
	j.o[key] = i
}

func (j *JSONBuilder) PutBool(i bool, key string) {
	j.o[key] = i
}

func (j *JSONBuilder) PutDouble(i float64, key string) {
	j.o[key] = i
}

func (j *JSONBuilder) PutJSON(i globals.JSONStruct, key string) {
	j.o[key] = globals.JSONStruct{}
}

func (j *JSONBuilder) GetJSON(key string) globals.JSONStruct {
	child, ok := j.o[key].(globals.JSONStruct)
	if !ok {
		return nil
	}
	return child
}
