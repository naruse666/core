// Code generated by 'yaegi extract cogentcore.org/core/tensor/vector'. DO NOT EDIT.

package symbols

import (
	"cogentcore.org/core/tensor/vector"
	"reflect"
)

func init() {
	Symbols["cogentcore.org/core/tensor/vector/vector"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Dot":    reflect.ValueOf(vector.Dot),
		"Mul":    reflect.ValueOf(vector.Mul),
		"MulOut": reflect.ValueOf(vector.MulOut),
		"NormL1": reflect.ValueOf(vector.NormL1),
		"NormL2": reflect.ValueOf(vector.NormL2),
		"Sum":    reflect.ValueOf(vector.Sum),
	}
}
