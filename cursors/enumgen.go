// Code generated by "core generate"; DO NOT EDIT.

package cursors

import (
	"github.com/naruse666/core/enums"
)

var _CursorValues = []Cursor{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39}

// CursorN is the highest valid value for type Cursor, plus one.
const CursorN Cursor = 40

var _CursorValueMap = map[string]Cursor{`none`: 0, `arrow`: 1, `context-menu`: 2, `help`: 3, `pointer`: 4, `progress`: 5, `wait`: 6, `cell`: 7, `crosshair`: 8, `text`: 9, `vertical-text`: 10, `alias`: 11, `copy`: 12, `move`: 13, `not-allowed`: 14, `grab`: 15, `grabbing`: 16, `resize-col`: 17, `resize-row`: 18, `resize-up`: 19, `resize-right`: 20, `resize-down`: 21, `resize-left`: 22, `resize-n`: 23, `resize-e`: 24, `resize-s`: 25, `resize-w`: 26, `resize-ne`: 27, `resize-nw`: 28, `resize-se`: 29, `resize-sw`: 30, `resize-ew`: 31, `resize-ns`: 32, `resize-nesw`: 33, `resize-nwse`: 34, `zoom-in`: 35, `zoom-out`: 36, `screenshot-selection`: 37, `screenshot-window`: 38, `poof`: 39}

var _CursorDescMap = map[Cursor]string{0: `None indicates no preference for a cursor; will typically be inherited`, 1: `Arrow is a standard arrow cursor, which is the default window cursor`, 2: `ContextMenu indicates that a context menu is available`, 3: `Help indicates that help information is available`, 4: `Pointer is a pointing hand that indicates a link or an interactive element`, 5: `Progress indicates that the app is busy in the background, but can still be interacted with (use [Wait] to indicate that it can&#39;t be interacted with)`, 6: `Wait indicates that the app is busy and can not be interacted with (use [Progress] to indicate that it can be interacted with)`, 7: `Cell indicates a table cell, especially one that can be selected`, 8: `Crosshair is a cross cursor that typically indicates precision selection, such as in an image`, 9: `Text is an I-Beam that indicates text that can be selected`, 10: `VerticalText is a sideways I-Beam that indicates vertical text that can be selected`, 11: `Alias indicates that a shortcut or alias will be created`, 12: `Copy indicates that a copy of something will be created`, 13: `Move indicates that something is being moved`, 14: `NotAllowed indicates that something can not be done`, 15: `Grab indicates that something can be grabbed`, 16: `Grabbing indicates that something is actively being grabbed`, 17: `ResizeCol indicates that something can be resized in the horizontal direction`, 18: `ResizeRow indicates that something can be resized in the vertical direction`, 19: `ResizeUp indicates that something can be resized in the upper direction`, 20: `ResizeRight indicates that something can be resized in the right direction`, 21: `ResizeDown indicates that something can be resized in the downward direction`, 22: `ResizeLeft indicates that something can be resized in the left direction`, 23: `ResizeN indicates that something can be resized in the upper direction`, 24: `ResizeE indicates that something can be resized in the right direction`, 25: `ResizeS indicates that something can be resized in the downward direction`, 26: `ResizeW indicates that something can be resized in the left direction`, 27: `ResizeNE indicates that something can be resized in the upper-right direction`, 28: `ResizeNW indicates that something can be resized in the upper-left direction`, 29: `ResizeSE indicates that something can be resized in the lower-right direction`, 30: `ResizeSW indicates that something can be resized in the lower-left direction`, 31: `ResizeEW indicates that something can be resized bidirectionally in the right-left direction`, 32: `ResizeNS indicates that something can be resized bidirectionally in the top-bottom direction`, 33: `ResizeNESW indicates that something can be resized bidirectionally in the top-right to bottom-left direction`, 34: `ResizeNWSE indicates that something can be resized bidirectionally in the top-left to bottom-right direction`, 35: `ZoomIn indicates that something can be zoomed in`, 36: `ZoomOut indicates that something can be zoomed out`, 37: `ScreenshotSelection indicates that a screenshot selection box is being selected`, 38: `ScreenshotWindow indicates that a screenshot is being taken of an entire window`, 39: `Poof indicates that an item will dissapear when it is released`}

var _CursorMap = map[Cursor]string{0: `none`, 1: `arrow`, 2: `context-menu`, 3: `help`, 4: `pointer`, 5: `progress`, 6: `wait`, 7: `cell`, 8: `crosshair`, 9: `text`, 10: `vertical-text`, 11: `alias`, 12: `copy`, 13: `move`, 14: `not-allowed`, 15: `grab`, 16: `grabbing`, 17: `resize-col`, 18: `resize-row`, 19: `resize-up`, 20: `resize-right`, 21: `resize-down`, 22: `resize-left`, 23: `resize-n`, 24: `resize-e`, 25: `resize-s`, 26: `resize-w`, 27: `resize-ne`, 28: `resize-nw`, 29: `resize-se`, 30: `resize-sw`, 31: `resize-ew`, 32: `resize-ns`, 33: `resize-nesw`, 34: `resize-nwse`, 35: `zoom-in`, 36: `zoom-out`, 37: `screenshot-selection`, 38: `screenshot-window`, 39: `poof`}

// String returns the string representation of this Cursor value.
func (i Cursor) String() string { return enums.String(i, _CursorMap) }

// SetString sets the Cursor value from its string representation,
// and returns an error if the string is invalid.
func (i *Cursor) SetString(s string) error { return enums.SetString(i, s, _CursorValueMap, "Cursor") }

// Int64 returns the Cursor value as an int64.
func (i Cursor) Int64() int64 { return int64(i) }

// SetInt64 sets the Cursor value from an int64.
func (i *Cursor) SetInt64(in int64) { *i = Cursor(in) }

// Desc returns the description of the Cursor value.
func (i Cursor) Desc() string { return enums.Desc(i, _CursorDescMap) }

// CursorValues returns all possible values for the type Cursor.
func CursorValues() []Cursor { return _CursorValues }

// Values returns all possible values for the type Cursor.
func (i Cursor) Values() []enums.Enum { return enums.Values(_CursorValues) }

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Cursor) MarshalText() ([]byte, error) { return []byte(i.String()), nil }

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Cursor) UnmarshalText(text []byte) error { return enums.UnmarshalText(i, text, "Cursor") }
