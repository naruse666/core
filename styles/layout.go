// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package styles

import (
	"goki.dev/girl/units"
)

// todo: for style
// Resize: user-resizability
// z-index

// CSS vs. Layout alignment
//
// CSS has align-self, align-items (for a container, provides a default for
// items) and align-content which only applies to lines in a flex layout (akin
// to a flow layout) -- there is a presumed horizontal aspect to these, except
// align-content, so they are subsumed in the AlignH parameter in this style.
// Vertical-align works as expected, and Text.Align uses left/center/right

// IMPORTANT: any changes here must be updated in style_props.go StyleLayoutFuncs

// ScrollBarWidthDefault is the default width of a scrollbar in pixels
var ScrollBarWidthDefault = float32(10)

func (s *Style) LayoutDefaults() {
	s.Gap.Set(units.Em(0.5))
	s.ScrollBarWidth.Dp(ScrollBarWidthDefault)
}

// ToDots runs ToDots on unit values, to compile down to raw pixels
func (s *Style) LayoutToDots(uc *units.Context) {
	s.Pos.ToDots(uc)
	s.Min.ToDots(uc)
	s.Max.ToDots(uc)
	s.Padding.ToDots(uc)
	s.Margin.ToDots(uc)
	s.Gap.ToDots(uc)
	s.ScrollBarWidth.ToDots(uc)
}

// Display determines how items are displayed
type Display int32 //enums:enum -trim-prefix Display

const (
	// Flex is the default layout model, based on a simplified version of the
	// CSS flex layout: uses MainAxis to specify the direction, Wrap for
	// wrapping of elements, and Min, Max, and Grow values on elements to
	// determine sizing.
	DisplayFlex Display = iota

	// Stacked is a stack of elements, with one on top that is visible
	DisplayStacked

	// Grid is the X, Y grid layout, with Columns specifying the number
	// of elements in the X axis.
	DisplayGrid

	// None means the item is not displayed: sets the Invisible state
	DisplayNone
)

// Align has all different types of alignment -- only some are applicable to
// different contexts, but there is also so much overlap that it makes sense
// to have them all in one list -- some are not standard CSS and used by
// layout
type Align int32 //enums:enum -trim-prefix Align

const (
	// Align items to the start (top, left) of layout
	AlignStart Align = iota

	// Align items to the end (bottom, right) of layout
	AlignEnd

	// Align all items centered around the center of layout space
	AlignCenter

	// Align to text baselines
	AlignBaseline

	// First and last are flush, equal space between remaining items
	AlignSpaceBetween

	// First and last have 1/2 space at edges, full space between remaining items
	AlignSpaceAround

	// Equal space at start, end, and between all items
	AlignSpaceEvenly
)

// overflow type -- determines what happens when there is too much stuff in a layout
type Overflow int32 //enums:enum -trim-prefix Overflow

const (
	// OverflowVisible makes the overflow visible, meaning that the size
	// of the container is always at least the Min size of its contents.
	// No scrollbars are shown.
	OverflowVisible Overflow = iota

	// OverflowHidden hides the overflow and doesn't present scrollbars.
	OverflowHidden

	// OverflowAuto automatically determines if scrollbars should be added to show
	// the overflow.  Scrollbars are added only if the actual content size is greater
	// than the currently available size.
	OverflowAuto

	// OverflowScroll means that scrollbars are always visible,
	// and is otherwise identical to Auto.  However, only during Viewport PrefSize call,
	// the actual content size is used -- otherwise it behaves just like Auto.
	OverflowScroll
)
