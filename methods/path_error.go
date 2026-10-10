package tafexpr

import (
	"strings"

	"github.com/gclkaze/evalang-globals/globals/stackvalue"
	"github.com/gclkaze/tafexpr/methods"
)

// pathError builds the message for a failed  $var.path  read. It names what the receiver holds
// and, when a method of that name exists for that type, suggests the call syntax
// ($list.length -> "Did you mean $list.length()?").
func (l *TAFArgumentListener) pathError(varName, path string) error {
	parentPath, field := "", path
	if i := strings.LastIndex(path, "."); i >= 0 {
		parentPath, field = path[:i], path[i+1:]
	}
	receiver := varName
	if parentPath != "" {
		receiver += "." + parentPath
	}
	return methods.FieldReadError(receiver, field, l.valueAt(varName, parentPath))
}

// valueAt returns the value of the receiver of the failed read (the variable itself, or the
// part of the path before the last segment), unwrapped so its real type shows; nil if unknown.
func (l *TAFArgumentListener) valueAt(varName, parentPath string) stackvalue.StackValue {
	var (
		v   stackvalue.StackValue
		err error
	)
	if parentPath == "" {
		v, err = l.VariableContext.GetVariableValue(varName, l.secretAware)
	} else {
		v, err = l.VariableContext.EvaluateJSONVariable(varName, parentPath, l.secretAware)
	}
	if err != nil || v == nil {
		return nil
	}
	return normalize(v)
}
