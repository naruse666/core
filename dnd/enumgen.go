// Code generated by "enumgen"; DO NOT EDIT.

package dnd

import (
	"errors"
	"strconv"
	"strings"

	"goki.dev/enums"
)

var _ActionsValues = []Actions{0, 1, 2, 3, 4, 5, 6, 7, 8}

// ActionsN is the highest valid value
// for type Actions, plus one.
const ActionsN Actions = 9

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _ActionsNoOp() {
	var x [1]struct{}
	_ = x[NoAction-(0)]
	_ = x[Start-(1)]
	_ = x[DropOnTarget-(2)]
	_ = x[DropFmSource-(3)]
	_ = x[External-(4)]
	_ = x[Move-(5)]
	_ = x[Enter-(6)]
	_ = x[Exit-(7)]
	_ = x[Hover-(8)]
}

var _ActionsNameToValueMap = map[string]Actions{
	`NoAction`:     0,
	`noaction`:     0,
	`Start`:        1,
	`start`:        1,
	`DropOnTarget`: 2,
	`dropontarget`: 2,
	`DropFmSource`: 3,
	`dropfmsource`: 3,
	`External`:     4,
	`external`:     4,
	`Move`:         5,
	`move`:         5,
	`Enter`:        6,
	`enter`:        6,
	`Exit`:         7,
	`exit`:         7,
	`Hover`:        8,
	`hover`:        8,
}

var _ActionsDescMap = map[Actions]string{
	0: ``,
	1: `Start is triggered when criteria for DND starting have been met -- it is the chance for potential sources to start a DND event.`,
	2: `DropOnTarget is set when event is sent to the target where the item is dropped.`,
	3: `DropFmSource is set when event is sent back to the source after the target has been dropped on a valid target that did not ignore the event -- the source should check if Mod = DropMove, and typically delete itself in this case.`,
	4: `External is triggered from an external drop event`,
	5: `Move is sent whenever mouse is moving while dragging -- usually not needed.`,
	6: `Enter is sent when drag enters a given widget, in a FocusEvent.`,
	7: `Exit is sent when drag exits a given widget, in a FocusEvent. Exit from one widget always happens before entering another (so you can reset cursor to Not).`,
	8: `Hover is sent when drag is hovering over a widget without moving -- can use this for spring-loaded opening of items to drag into, for example.`,
}

var _ActionsMap = map[Actions]string{
	0: `NoAction`,
	1: `Start`,
	2: `DropOnTarget`,
	3: `DropFmSource`,
	4: `External`,
	5: `Move`,
	6: `Enter`,
	7: `Exit`,
	8: `Hover`,
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

var _DropModsValues = []DropMods{0, 1, 2, 3, 4}

// DropModsN is the highest valid value
// for type DropMods, plus one.
const DropModsN DropMods = 5

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _DropModsNoOp() {
	var x [1]struct{}
	_ = x[NoDropMod-(0)]
	_ = x[DropCopy-(1)]
	_ = x[DropMove-(2)]
	_ = x[DropLink-(3)]
	_ = x[DropIgnore-(4)]
}

var _DropModsNameToValueMap = map[string]DropMods{
	`NoDropMod`:  0,
	`nodropmod`:  0,
	`DropCopy`:   1,
	`dropcopy`:   1,
	`DropMove`:   2,
	`dropmove`:   2,
	`DropLink`:   3,
	`droplink`:   3,
	`DropIgnore`: 4,
	`dropignore`: 4,
}

var _DropModsDescMap = map[DropMods]string{
	0: ``,
	1: `Copy is the default and implies data is just copied -- receiver can do with it as they please and source does not need to take any further action`,
	2: `Move is signaled with a Shift or Meta key (by default) and implies that the source should delete itself when it receives the DropFmSource event action with this Mod value set -- receiver must update the Mod to reflect actual action taken, and be particularly careful with this one`,
	3: `Link can be any other kind of alternative action -- link is applicable to files (symbolic link)`,
	4: `Ignore means that the receiver chose to not process this drop`,
}

var _DropModsMap = map[DropMods]string{
	0: `NoDropMod`,
	1: `DropCopy`,
	2: `DropMove`,
	3: `DropLink`,
	4: `DropIgnore`,
}

// String returns the string representation
// of this DropMods value.
func (i DropMods) String() string {
	if str, ok := _DropModsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the DropMods value from its
// string representation, and returns an
// error if the string is invalid.
func (i *DropMods) SetString(s string) error {
	if val, ok := _DropModsNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _DropModsNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type DropMods")
}

// Int64 returns the DropMods value as an int64.
func (i DropMods) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the DropMods value from an int64.
func (i *DropMods) SetInt64(in int64) {
	*i = DropMods(in)
}

// Desc returns the description of the DropMods value.
func (i DropMods) Desc() string {
	if str, ok := _DropModsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// DropModsValues returns all possible values
// for the type DropMods.
func DropModsValues() []DropMods {
	return _DropModsValues
}

// Values returns all possible values
// for the type DropMods.
func (i DropMods) Values() []enums.Enum {
	res := make([]enums.Enum, len(_DropModsValues))
	for i, d := range _DropModsValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type DropMods.
func (i DropMods) IsValid() bool {
	_, ok := _DropModsMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i DropMods) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *DropMods) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}
