// Code generated by "core generate"; DO NOT EDIT.

package indent

import (
	"github.com/naruse666/core/enums"
)

var _CharacterValues = []Character{0, 1}

// CharacterN is the highest valid value for type Character, plus one.
const CharacterN Character = 2

var _CharacterValueMap = map[string]Character{`Tab`: 0, `Space`: 1}

var _CharacterDescMap = map[Character]string{0: `Tab indicates to use tabs for indentation.`, 1: `Space indicates to use spaces for indentation.`}

var _CharacterMap = map[Character]string{0: `Tab`, 1: `Space`}

// String returns the string representation of this Character value.
func (i Character) String() string { return enums.String(i, _CharacterMap) }

// SetString sets the Character value from its string representation,
// and returns an error if the string is invalid.
func (i *Character) SetString(s string) error {
	return enums.SetString(i, s, _CharacterValueMap, "Character")
}

// Int64 returns the Character value as an int64.
func (i Character) Int64() int64 { return int64(i) }

// SetInt64 sets the Character value from an int64.
func (i *Character) SetInt64(in int64) { *i = Character(in) }

// Desc returns the description of the Character value.
func (i Character) Desc() string { return enums.Desc(i, _CharacterDescMap) }

// CharacterValues returns all possible values for the type Character.
func CharacterValues() []Character { return _CharacterValues }

// Values returns all possible values for the type Character.
func (i Character) Values() []enums.Enum { return enums.Values(_CharacterValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Character) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Character) UnmarshalText(text []byte) error {
	return enums.UnmarshalText(i, text, "Character")
}
