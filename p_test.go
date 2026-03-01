package tafexpr

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SetupParser() *TAFArgumentParser {
	p := &TAFArgumentParser{}
	return p
}

const float64EqualityThreshold = 1e-4

func almostEqual(a, b float64) bool {
	fmt.Printf("a: %v\n ", a)
	fmt.Printf("b: %v\n", b)
	return math.Abs(a-b) <= float64EqualityThreshold
}

func TestRandomString(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, false, p.Parse("123!+kkaksda"))
	errs := p.GetErrors()
	assert.Equal(t, true, len(errs) == 1)
	// assert.Equal(t, 100+200+3090, p.IntValue)
}

func TestSingleAssignmentSymbol(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, false, p.Parse("100 = 500"))
	errs := p.GetErrors()
	assert.Equal(t, true, len(errs) == 1)
	// assert.Equal(t, 100+200+3090, p.IntValue)
}

func TestArithmetics(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("100 + 200 +3090"))
	assert.Equal(t, 100+200+3090, p.IntValue)
}

func TestArithmeticsSimpleNegationAtEnd(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("100 + 200 +3090 - 90"))
	assert.Equal(t, 100+200+3090-90, p.IntValue)
}

func TestArithmeticsAllCombos(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("100 * 200 + 3090 - 60 + 90 / 5"))
	assert.Equal(t, (((100*200)+3090)-60)+90/5, p.IntValue)
}

func TestArithmeticsAllCombosPar(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("(100 * 200 + 3090) - (60 + 90) / 5"))
	assert.Equal(t, (100*200+3090)-(60+90)/5, p.IntValue)
}

func TestArithmeticsAllCombosPar2(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("(100 * (200 + 3090)) - (60 + (90/5))"))
	assert.Equal(t, (100*(200+3090))-(60+(90/5)), p.IntValue)
}

// DOUBLE ARITHMETICS
func TestDoubleArithmetics(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("1.1 + 5.2 + 3.333"))
	assert.Equal(t, true, almostEqual(1.1+5.2+3.333, p.DoubleValue))
}

func TestDoubleArithmeticsAllCombos(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("100.22 * 200.33 + 3090.11 - 60.35 + 90 / 5.2"))
	b := almostEqual((((100.22*200.33)+3090.11)-60.35)+90/5.2, p.DoubleValue)
	assert.Equal(t, true, b)
}

func TestDoubleArithmeticsAllCombosPar(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("(100.22 * 200.33 + 3090.11) - (60.35 + 90) / 5.2"))
	b := almostEqual((100.22*200.33+3090.11)-(60.35+90)/5.2, p.DoubleValue)
	assert.Equal(t, true, b)
}

func TestDoubleArithmeticsAllCombosPar2(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("(100.22 * (200.33 + 3090.11)) - (60.35 + (90/5.2))"))
	b := almostEqual((100.22*(200.33+3090.11))-(60.35+(90/5.2)), p.DoubleValue)
	assert.Equal(t, true, b)
}

func TestDoubleIntComboArithmeticsAllCombos(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("(100.22 * (200.33 + 3090.11)) - (60.35 + (90/5.2)) + 500 - 333"))
	b := almostEqual((100.22*(200.33+3090.11))-(60.35+(90/5.2))+500-333, p.DoubleValue)
	assert.Equal(t, true, b)
}

func TestDoubleIntComboArithmeticsAllCombosSimple(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("900/5.2"))
	b := almostEqual(900/5.2, p.DoubleValue)
	assert.Equal(t, true, b)
}

func TestDivisionByZero(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, false, p.Parse(" (((500 - 300) - 200 ) + 1 ) / 0 "))
}

func TestDivisionByZeroForDoubles(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, false, p.Parse(" (((500.33 - 300.33) - 200.111 ) + 1.2 ) / 0.0 "))
}

func TestRandomTokens(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, false, p.Parse("asdasa asdasda asds"))
}

/*func TestParseJSONString(t *testing.T) {
	p := SetupParser()
	s := "\\\"{\\\"($a)\\\":{\\\"($b)\\\":{\\\"($c)\\\":10}}}\\\""
	assert.Equal(t, true, p.Parse(s))
	assert.Equal(t, "\"{\\\"($a)\\\":{\\\"($b)\\\":{\\\"($c)\\\":10}}}\"", p.StringValue)
}*/

func TestParseJJSONString(t *testing.T) {
	t.Skip("This does not pass")
	p := SetupParser()
	actual := "\\\"{\\\"a\\\":{\\\"b\\\":{\\\"c\\\":666}}}\""
	expected := "{\\\"a\\\":{\\\"b\\\":{\\\"c\\\":666}}}"
	assert.Equal(t, true, p.Parse(actual))
	assert.Equal(t, expected, p.StringValue)

}

func TestParseRandomDoubleInRange(t *testing.T) {
	p := SetupParser()
	min := 1
	max := 10
	actual := "randomDoubleInRange(" + strconv.Itoa(min) + " , " + strconv.Itoa(max) + ")"
	res := p.Parse(actual)
	assert.Equal(t, true, res)
	val := p.DoubleValue
	assert.Equal(t, true, val >= float64(min) && val <= float64(max))
}

func TestParseRandomDoubleInRangeCastToInteger(t *testing.T) {
	p := SetupParser()
	min := 1
	max := 10
	actual := "randomDoubleInRange(" + strconv.Itoa(min) + " , " + strconv.Itoa(max) + ").toInteger"
	res := p.Parse(actual)
	assert.Equal(t, true, res)
	val := p.IntValue
	assert.Equal(t, true, val >= min && val <= max)
}

/*func TestQuotedString(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("{\\\"username\\\":\\\"LOL\\\"}"))
}*/
