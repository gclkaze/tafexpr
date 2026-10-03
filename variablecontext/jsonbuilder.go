package variablecontext

import "github.com/gclkaze/evalang-globals/globals"

// JSONObjectBuilder accumulates key/value pairs and produces a
// globals.JSONObjectGen (a map[string]any under the hood, since
// JSONObjectGen = interface{}) once you're done adding to it.
type JSONObjectBuilder struct {
	values map[string]any
}

func NewJSONObjectBuilder() *JSONObjectBuilder {
	return &JSONObjectBuilder{values: make(map[string]any)}
}

// Set adds or overwrites a single key/value pair, returning the builder
// itself so calls can be chained: builder.Set("a", 1).Set("b", "x").
func (b *JSONObjectBuilder) Set(key string, value any) *JSONObjectBuilder {
	b.values[key] = value
	return b
}

// SetAll merges every entry from the given map into the builder,
// overwriting any existing keys with the same name.
func (b *JSONObjectBuilder) SetAll(entries map[string]any) *JSONObjectBuilder {
	for k, v := range entries {
		b.values[k] = v
	}
	return b
}

// Build returns the accumulated key/value pairs as a globals.JSONObjectGen.
func (b *JSONObjectBuilder) Build() globals.JSONObjectGen {
	return b.values
}
