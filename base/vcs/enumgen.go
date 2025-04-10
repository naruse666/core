// Code generated by "core generate"; DO NOT EDIT.

package vcs

import (
	"github.com/naruse666/core/enums"
)

var _FileStatusValues = []FileStatus{0, 1, 2, 3, 4, 5, 6}

// FileStatusN is the highest valid value for type FileStatus, plus one.
const FileStatusN FileStatus = 7

var _FileStatusValueMap = map[string]FileStatus{`Untracked`: 0, `Stored`: 1, `Modified`: 2, `Added`: 3, `Deleted`: 4, `Conflicted`: 5, `Updated`: 6}

var _FileStatusDescMap = map[FileStatus]string{0: `Untracked means file is not under VCS control`, 1: `Stored means file is stored under VCS control, and has not been modified in working copy`, 2: `Modified means file is under VCS control, and has been modified in working copy`, 3: `Added means file has just been added to VCS but is not yet committed`, 4: `Deleted means file has been deleted from VCS`, 5: `Conflicted means file is in conflict -- has not been merged`, 6: `Updated means file has been updated in the remote but not locally`}

var _FileStatusMap = map[FileStatus]string{0: `Untracked`, 1: `Stored`, 2: `Modified`, 3: `Added`, 4: `Deleted`, 5: `Conflicted`, 6: `Updated`}

// String returns the string representation of this FileStatus value.
func (i FileStatus) String() string { return enums.String(i, _FileStatusMap) }

// SetString sets the FileStatus value from its string representation,
// and returns an error if the string is invalid.
func (i *FileStatus) SetString(s string) error {
	return enums.SetString(i, s, _FileStatusValueMap, "FileStatus")
}

// Int64 returns the FileStatus value as an int64.
func (i FileStatus) Int64() int64 { return int64(i) }

// SetInt64 sets the FileStatus value from an int64.
func (i *FileStatus) SetInt64(in int64) { *i = FileStatus(in) }

// Desc returns the description of the FileStatus value.
func (i FileStatus) Desc() string { return enums.Desc(i, _FileStatusDescMap) }

// FileStatusValues returns all possible values for the type FileStatus.
func FileStatusValues() []FileStatus { return _FileStatusValues }

// Values returns all possible values for the type FileStatus.
func (i FileStatus) Values() []enums.Enum { return enums.Values(_FileStatusValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i FileStatus) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *FileStatus) UnmarshalText(text []byte) error {
	return enums.UnmarshalText(i, text, "FileStatus")
}

var _TypesValues = []Types{0, 1, 2, 3, 4}

// TypesN is the highest valid value for type Types, plus one.
const TypesN Types = 5

var _TypesValueMap = map[string]Types{`NoVCS`: 0, `novcs`: 0, `Git`: 1, `git`: 1, `Svn`: 2, `svn`: 2, `Bzr`: 3, `bzr`: 3, `Hg`: 4, `hg`: 4}

var _TypesDescMap = map[Types]string{0: ``, 1: ``, 2: ``, 3: ``, 4: ``}

var _TypesMap = map[Types]string{0: `NoVCS`, 1: `Git`, 2: `Svn`, 3: `Bzr`, 4: `Hg`}

// String returns the string representation of this Types value.
func (i Types) String() string { return enums.String(i, _TypesMap) }

// SetString sets the Types value from its string representation,
// and returns an error if the string is invalid.
func (i *Types) SetString(s string) error { return enums.SetStringLower(i, s, _TypesValueMap, "Types") }

// Int64 returns the Types value as an int64.
func (i Types) Int64() int64 { return int64(i) }

// SetInt64 sets the Types value from an int64.
func (i *Types) SetInt64(in int64) { *i = Types(in) }

// Desc returns the description of the Types value.
func (i Types) Desc() string { return enums.Desc(i, _TypesDescMap) }

// TypesValues returns all possible values for the type Types.
func TypesValues() []Types { return _TypesValues }

// Values returns all possible values for the type Types.
func (i Types) Values() []enums.Enum { return enums.Values(_TypesValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Types) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Types) UnmarshalText(text []byte) error { return enums.UnmarshalText(i, text, "Types") }
