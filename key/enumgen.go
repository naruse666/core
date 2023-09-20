// Code generated by "enumgen"; DO NOT EDIT.

package key

import (
	"errors"
	"strconv"
	"strings"

	"goki.dev/enums"
)

var _ActionsValues = []Actions{0, 1, 2}

// ActionsN is the highest valid value
// for type Actions, plus one.
const ActionsN Actions = 3

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _ActionsNoOp() {
	var x [1]struct{}
	_ = x[NoAction-(0)]
	_ = x[Press-(1)]
	_ = x[Release-(2)]
}

var _ActionsNameToValueMap = map[string]Actions{
	`NoAction`: 0,
	`noaction`: 0,
	`Press`:    1,
	`press`:    1,
	`Release`:  2,
	`release`:  2,
}

var _ActionsDescMap = map[Actions]string{
	0: ``,
	1: ``,
	2: ``,
}

var _ActionsMap = map[Actions]string{
	0: `NoAction`,
	1: `Press`,
	2: `Release`,
}

// String returns the string representation
// of this Actions value.
func (i Actions) String() string {
	if str, ok := _ActionsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Actions value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Actions) SetString(s string) error {
	if val, ok := _ActionsNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _ActionsNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type Actions")
}

// Int64 returns the Actions value as an int64.
func (i Actions) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Actions value from an int64.
func (i *Actions) SetInt64(in int64) {
	*i = Actions(in)
}

// Desc returns the description of the Actions value.
func (i Actions) Desc() string {
	if str, ok := _ActionsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// ActionsValues returns all possible values
// for the type Actions.
func ActionsValues() []Actions {
	return _ActionsValues
}

// Values returns all possible values
// for the type Actions.
func (i Actions) Values() []enums.Enum {
	res := make([]enums.Enum, len(_ActionsValues))
	for i, d := range _ActionsValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Actions.
func (i Actions) IsValid() bool {
	_, ok := _ActionsMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Actions) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Actions) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}

var _ModifiersValues = []Modifiers{0, 1, 2, 3}

// ModifiersN is the highest valid value
// for type Modifiers, plus one.
const ModifiersN Modifiers = 4

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _ModifiersNoOp() {
	var x [1]struct{}
	_ = x[Shift-(0)]
	_ = x[Control-(1)]
	_ = x[Alt-(2)]
	_ = x[Meta-(3)]
}

var _ModifiersNameToValueMap = map[string]Modifiers{
	`Shift`:   0,
	`shift`:   0,
	`Control`: 1,
	`control`: 1,
	`Alt`:     2,
	`alt`:     2,
	`Meta`:    3,
	`meta`:    3,
}

var _ModifiersDescMap = map[Modifiers]string{
	0: ``,
	1: ``,
	2: ``,
	3: ``,
}

var _ModifiersMap = map[Modifiers]string{
	0: `Shift`,
	1: `Control`,
	2: `Alt`,
	3: `Meta`,
}

// String returns the string representation
// of this Modifiers value.
func (i Modifiers) String() string {
	if str, ok := _ModifiersMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Modifiers value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Modifiers) SetString(s string) error {
	if val, ok := _ModifiersNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _ModifiersNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type Modifiers")
}

// Int64 returns the Modifiers value as an int64.
func (i Modifiers) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Modifiers value from an int64.
func (i *Modifiers) SetInt64(in int64) {
	*i = Modifiers(in)
}

// Desc returns the description of the Modifiers value.
func (i Modifiers) Desc() string {
	if str, ok := _ModifiersDescMap[i]; ok {
		return str
	}
	return i.String()
}

// ModifiersValues returns all possible values
// for the type Modifiers.
func ModifiersValues() []Modifiers {
	return _ModifiersValues
}

// Values returns all possible values
// for the type Modifiers.
func (i Modifiers) Values() []enums.Enum {
	res := make([]enums.Enum, len(_ModifiersValues))
	for i, d := range _ModifiersValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Modifiers.
func (i Modifiers) IsValid() bool {
	_, ok := _ModifiersMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Modifiers) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Modifiers) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}
