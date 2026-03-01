package tafexpr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComparison(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("1 && (0 || 2) && (2 >= 1) == 1 + 2"))
	assert.Equal(t, true, p.BoolValue)
}

func TestComparisonNeg(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("!(1 && (0 || 2) && (2 >= 1)) == !(1 + 2)"))
	assert.Equal(t, true, p.BoolValue)
}

func TestComparisonWithAdds(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("1+1<3"))
	assert.Equal(t, true, p.BoolValue)
}
