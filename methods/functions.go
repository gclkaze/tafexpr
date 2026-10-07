package methods

import (
	"github.com/gclkaze/evalang-globals/globals/stackvalue"
	"github.com/gclkaze/evalang-globals/utils"
	mine "github.com/gclkaze/tafexpr/stackvalue"
)

var numeric = []stackvalue.StackValueType{stackvalue.INTEGER, stackvalue.DOUBLE}

// Free functions (called as name(args), with no receiver).
func init() {
	RegisterFunc("randomDoubleInRange", Method{Group: "random", Returns: "Double", Doc: "A random number between min and max.",
		Params: []Param{{Name: "min", Types: numeric}, {Name: "max", Types: numeric}},
		Fn: func(c *Call) (stackvalue.StackValue, error) {
			lo, err := c.Args[0].ToDouble()
			if err != nil {
				return nil, err
			}
			hi, err := c.Args[1].ToDouble()
			if err != nil {
				return nil, err
			}
			return mine.NewDoubleStackValue(utils.GetRandomDoubleInRange(lo, hi)), nil
		}})
}
