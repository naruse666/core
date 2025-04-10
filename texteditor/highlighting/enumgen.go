// Code generated by "core generate -add-types"; DO NOT EDIT.

package highlighting

import (
	"github.com/naruse666/core/enums"
)

var _TrileanValues = []Trilean{0, 1, 2}

// TrileanN is the highest valid value for type Trilean, plus one.
const TrileanN Trilean = 3

var _TrileanValueMap = map[string]Trilean{`Pass`: 0, `Yes`: 1, `No`: 2}

var _TrileanDescMap = map[Trilean]string{0: ``, 1: ``, 2: ``}

var _TrileanMap = map[Trilean]string{0: `Pass`, 1: `Yes`, 2: `No`}

// String returns the string representation of this Trilean value.
func (i Trilean) String() string { return enums.String(i, _TrileanMap) }

// SetString sets the Trilean value from its string representation,
// and returns an error if the string is invalid.
func (i *Trilean) SetString(s string) error {
	return enums.SetString(i, s, _TrileanValueMap, "Trilean")
}

// Int64 returns the Trilean value as an int64.
func (i Trilean) Int64() int64 { return int64(i) }

// SetInt64 sets the Trilean value from an int64.
func (i *Trilean) SetInt64(in int64) { *i = Trilean(in) }

// Desc returns the description of the Trilean value.
func (i Trilean) Desc() string { return enums.Desc(i, _TrileanDescMap) }

// TrileanValues returns all possible values for the type Trilean.
func TrileanValues() []Trilean { return _TrileanValues }

// Values returns all possible values for the type Trilean.
func (i Trilean) Values() []enums.Enum { return enums.Values(_TrileanValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Trilean) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Trilean) UnmarshalText(text []byte) error { return enums.UnmarshalText(i, text, "Trilean") }
