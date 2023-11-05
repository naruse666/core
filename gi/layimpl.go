// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gi

import (
	"fmt"
	"image"

	"goki.dev/girl/styles"
	"goki.dev/ki/v2"
	"goki.dev/mat32/v2"
)

// Layout uses 2 Size passes, 2 Position passes:
//
// 1. SizeUp (bottom-up): gathers sizes from our Children & Parts,
//    based only on Min style sizes and actual content sizing.
//    Flexible elements (e.g., Text, Flex Wrap, TopAppBar) allocate
//    optimistically along their main axis, up to any optional Max size
//
// 2. SizeDown (top-down, multiple iterations possible): assigns sizes based
//    on allocated parent avail size, giving extra space based on Grow factors,
//    and flexible elements wrap / config to fit top-down constraint along main
//    axis, producing a (new) top-down size expanding in cross axis as needed
//    (or removing items that don't fit, etc).  Wrap & Grid layouts assign
//    X,Y index coordinates to items during this pass.
//
//    This iterates until the resulting sizes values are stable.
//
// 3. Position: uses the final sizes to position everything within the layout
//    according to alignment settings.
//
// 4. ScenePos: scene-based position and final BBox is computed based on scrollbar
//    position.  This step can be performed when scrolling.

// LaySize has sizing information for content and total size, and grow factors
type LaySize struct {
	// Content is size of the contents (children, parts) of the widget,
	// excluding all additional spacing (padding etc)
	Content mat32.Vec2 `view:"inline"`

	// Total is total size of the widget,
	// including all additional spacing (padding, margin, scrollbars)a
	// but excluding gap, spacing managed by the layout.
	Total mat32.Vec2 `view:"inline"`
}

func (ls *LaySize) String() string {
	return fmt.Sprintf("Content: %v, \tTotal: %v", ls.Content, ls.Total)
}

// SetContentMax sets the Content size subject to given Max values
// which only apply if they are > 0.  Values are automatically
// rounded up to Ceil ints
func (ls *LaySize) SetContentMax(sz mat32.Vec2, mx mat32.Vec2) {
	ls.Content = sz
	ls.Content.SetMinPos(mx)
	ls.Content.SetToCeil()
}

// SetContentToFit sets the Content size subject to given Max values
// which only apply if they are > 0.  Values are automatically
// rounded up to Ceil ints
func (ls *LaySize) SetContentToFit(sz mat32.Vec2, mx mat32.Vec2) {
	ls.Content.SetMaxPos(sz)
	ls.Content.SetMinPos(mx)
	ls.Content.SetToCeil()
}

// SetTotalFromContent sets the Total size as Content plus given extra space
func (ls *LaySize) SetTotalFromContent(spc mat32.Vec2) {
	ls.Total = ls.Content.Add(spc)
}

// SetContentFromTotal sets the Content from Total size subtracting extra
func (ls *LaySize) SetContentFromTotal(spc mat32.Vec2) {
	ls.Content = ls.Total.Sub(spc)
}

// LayoutState contains the the layout state for each widget,
// Set by the parent Layout during the Layout process.
type LayoutState struct {
	// Size is the size of the element (updated during different passes,
	// and holding the final computed size).
	Size LaySize `view:"inline"`

	// Cell is the logical X, Y index coordinates (row, col) of element
	// within its parent layout
	Cell image.Point

	// Pos is position, relative to parent Content size space
	Pos mat32.Vec2

	// Scroll is additional scrolling offset within our parent layout
	Scroll mat32.Vec2

	// ScPos is position, relative to overall Scene that we render into,
	// including effects of scroll offset
	ScPos mat32.Vec2

	// 2D bounding box for Total size occupied within parent Scene object that we render onto.
	// These are the pixels we draw into, filtered through parent bounding boxes
	// (empty for invisible). Used for render Bounds clipping.
	BBox image.Rectangle `edit:"-" copy:"-" json:"-" xml:"-" set:"-"`

	// 2D bounding box for our content, which excludes our padding, margin, etc.
	ContentBBox image.Rectangle `edit:"-" copy:"-" json:"-" xml:"-" set:"-"`
}

// Reset is called at start of layout process -- resets all values back to 0
func (ls *LayoutState) Reset() {
	// ls.Actual.Reset()
	// ls.Final.Reset()
	// ls.BU.Reset()
	// ls.TD.Reset()
}

func (ls LayoutState) String() string {
	// return "Actual: " + ls.Actual.String() + "\nFinal:  " + ls.Final.String() +
	// 	+"\nBU:    " + ls.BU.String() + "\nTD:    " + ls.TD.String() + "\n"
	return ""
}

// LayImplCell holds the layout implementation data for one cell
type LayImplCell struct {
	// Total size of cell
	Size mat32.Vec2

	// GrowSum is the sum of grow factors along each dimension
	GrowSum mat32.Vec2
}

func (ls *LayImplCell) String() string {
	return fmt.Sprintf("Sizet: %v, \tGrowSum: %g", ls.Size, ls.GrowSum)
}

func (ls *LayImplCell) Reset() {
	ls.Size.SetZero()
	ls.GrowSum.SetZero()
}

// LayImplState has internal state for implementing layout
type LayImplState struct {
	// Cells is number of cells along each dimension,
	// computed for each layout type.
	Cells image.Point `edit:"-" copy:"-" json:"-" xml:"-" set:"-"`

	// line breaks for wrap layout
	WrapBreaks []int `edit:"-" copy:"-" json:"-" xml:"-"`

	// sizes has the data for the columns in [0] and rows in [1]:
	// col Size.X = max(X over rows) (cross axis), .Y = sum(Y over rows) (main axis for col)
	// row Size.X = sum(X over cols) (main axis for row), .Y = max(Y over cols) (cross axis)
	Sizes [2][]LayImplCell `edit:"-" copy:"-" json:"-" xml:"-" set:"-"`

	// ScrollSize has the scrollbar sizes (widths) for each dim, which adds extra space.
	// If there is a vertical scrollbar, X has width; if horizontal, Y has "width" = height
	ScrollSize mat32.Vec2

	// GapSize has the extra gap sizing between elements for each dimension
	GapSize mat32.Vec2
}

// InitSizes initializes the Sizes based on Cells geom
func (ls *LayImplState) InitSizes() {
	for d := mat32.X; d < mat32.DimsN; d++ {
		n := mat32.PointDim(ls.Cells, d)
		if len(ls.Sizes[d]) != n {
			ls.Sizes[d] = make([]LayImplCell, n)
		}
		for i := 0; i < n; i++ {
			ls.Sizes[i].Reset()
		}
	}
	ls.WrapBreaks = nil
	ls.ScrollSize.SetZero()
	ls.GapSize.SetZero()
}

// StackTopWidget returns the StackTop element as a widget
func (ly *Layout) StackTopWidget() (Widget, *WidgetBase) {
	sn, err := ly.ChildTry(ly.StackTop)
	if err != nil {
		return nil, nil
	}
	return AsWidget(sn)
}

//////////////////////////////////////////////////////////////////////
//		SizeUp

func (wb *WidgetBase) SizeUp(sc *Scene) {
	wb.SizeUpWidget(sc)
}

// SizeUpWidget is the standard Widget SizeUp pass
func (wb *WidgetBase) SizeUpWidget(sc *Scene) {
	wb.SizeFromStyle()
	wb.SizeUpParts(sc)
	wb.Alloc.Size.SetTotalFromContent(wb.Styles.BoxSpace().Size())
}

// SizeFromStyle sets the initial Content size from Min style size,
// subject to Max constraints.
func (wb *WidgetBase) SizeFromStyle() {
	wb.Alloc.Size.SetContentMax(wb.Styles.Min.Dots(), wb.Styles.Max.Dots())
}

// SizeUpParts adjusts the Content size to hold the Parts layout if present
func (wb *WidgetBase) SizeUpParts(sc *Scene) {
	if wb.Parts == nil {
		return
	}
	wb.Parts.SizeUp(sc)
	wb.Alloc.Size.SetContentToFit(wb.Parts.Alloc.Size.Total, wb.Styles.Max.Dots())
}

func (ly *Layout) SizeUp(sc *Scene) {
	ly.SizeUpLay(sc)
}

// SizeUpLay is the Layout standard SizeUp pass
func (ly *Layout) SizeUpLay(sc *Scene) {
	ly.SizeFromStyle()
	ly.SizeUpKids(sc)
	ly.SetInitCells()
	ly.SizeUpCells(sc)
	ly.Alloc.Size.SetTotalFromContent(ly.Styles.BoxSpace().Size())
}

// SizeUpKids calls SizeUp on all the children of this node
func (ly *Layout) SizeUpKids(sc *Scene) {
	ly.WidgetKidsIter(func(i int, kwi Widget, kwb *WidgetBase) bool {
		kwi.SizeUp(sc)
		return ki.Continue
	})
	return
}

// SetInitCells sets the initial default assignment of cell indexes
// to each widget, based on layout type.
func (ly *Layout) SetInitCells() {
	switch {
	case ly.Styles.Display == styles.DisplayFlex:
		ly.SetInitCellsFlex()
	case ly.Styles.Display == styles.DisplayStacked:
		ly.SetInitCellsStacked()
	case ly.Styles.Display == styles.DisplayGrid:
		ly.SetInitCellsGrid()
	default:
		ly.SetInitCellsStacked() // whatever
	}
}

func (ly *Layout) SetInitCellsFlex() {
	ma := ly.Styles.MainAxis
	ca := ma.OrthoDim()
	idx := 0
	ly.WidgetKidsIter(func(i int, kwi Widget, kwb *WidgetBase) bool {
		mat32.SetPointDim(&kwb.Alloc.Cell, ma, idx)
		idx++
		return ki.Continue
	})
	mat32.SetPointDim(&ly.LayImpl.Cells, ma, idx)
	mat32.SetPointDim(&ly.LayImpl.Cells, ca, 1)
}

func (ly *Layout) SetInitCellsStacked() {
	ly.WidgetKidsIter(func(i int, kwi Widget, kwb *WidgetBase) bool {
		kwb.Alloc.Cell.Set(0, 0)
		return ki.Continue
	})
	ly.LayImpl.Cells.Set(1, 1)
}

func (ly *Layout) SetInitCellsGrid() {
	n := len(*ly.Children())
	cols := ly.Styles.Columns
	if cols == 0 {
		cols = int(mat32.Sqrt(float32(n)))
	}
	rows = n / cols
	for rows*cols < n {
		rows++
	}
	ly.LayImpl.Cells.Set(cols, rows)
	ci := 0
	ri := 0
	ly.WidgetKidsIter(func(i int, kwi Widget, kwb *WidgetBase) bool {
		kwb.Alloc.Cell.Set(ci, ri)
		ci++
		cs := kwb.Styles.ColSpan
		if cs > 1 {
			ci += cs - 1
		}
		if ci >= cols {
			ci = 0
			ri++
		}
		return ki.Continue
	})
}

// todo: wrap requires a different non-grid logic -- no constraint of same sizes across rows

// SizeUpCells gathers sizing data for cells
func (ly *Layout) SizeUpCells(sc *Scene) {
	if ly.Styles.Display == styles.DisplayStacked {
		ly.SizeUpCellsStacked(sc)
		return
	}
	// r   0   1   col X = max(X over rows), Y = sum(Y over rows)
	//   +--+--+
	// 0 |  |  |   row X = sum(X over cols), Y = max(Y over cols)
	//   +--+--+
	// 1 |  |  |
	//   +--+--+

	ly.LayImpl.InitSizes()
	ly.WidgetKidsIter(func(i int, kwi Widget, kwb *WidgetBase) bool {
		ci := kwb.Alloc.Cell
		sz := kwb.Alloc.Size.Total
		grw := kwb.Styles.Grow
		for ma := mat32.X; ma < mat32.DimsN; ma++ { // main axis = X then Y
			ca := ma.OrthoDim()          // cross axis = Y then X
			mi := mat32.PointDim(ci, ma) // X, Y
			ci := mat32.PointDim(ci, ca) // Y, X
			md := &ly.LayImpl.Sizes[mi]  // X, Y
			cd := &ly.LayImpl.Sizes[ci]  // Y, X
			msz := sz.Dim(mi)            // main axis size dim: X, Y
			mx := md.Size.Dim(mi)
			mx = max(mx, msz) // Col, max widths of all elements; Row, max heights of all elements
			md.Size.SetDim(mi, mx)
			sm := cd.Size.Dim(ci)
			sm += msz
			cd.Size.SetDim(ci, sm)            // Row, sum widths of all elements; Col, sum heights of all elements
			cd.GrowSum.Dim(mi) += grw.Dim(mi) // row.X = sum of col grows; col.Y = sum of row grows
			md.GrowSum.Dim(mi) += grw.Dim(mi) // col.X = sum of row grows; row.Y = sum of col grows
		}
		return ki.Continue
	})
	if ly.Styles.Display == styles.DisplayFlex {
		r0 := ly.LayImpl.Sizes[1]
		fmt.Println("row 0:", r0)
	}

	var csz mat32.Vec2
	for ma := mat32.X; ma < mat32.DimsN; ma++ { // main axis = X then Y
		ca := ma.OrthoDim()       // cross axis = Y then X
		n := ly.LayImpl.Cells[ma] // cols, rows
		sum := float32(0)
		for i := 0; i < n; i++ {
			md := &ly.LayImpl.Sizes[mi] // X, Y
			mx := md.Dim(mi)
			sum += mx // sum of maxes
		}
		csz.SetDim(ma, sum)
	}
	if LayoutTrace {
		fmt.Println(ly, " Content Size:", csz)
	}
	ly.Alloc.Size.SetContentToFit(csz, ly.Styles.Max.Dots())
}

// SizeUpCellsStacked
func (ly *Layout) SizeUpCellsStacked(sc *Scene) {
	_, kwb := StackTopWidget()
	if kwb == nil {
		return
	}
	ly.LayImpl.Sizes[0].Size = kwb.Alloc.Size.Total
	ly.LayImpl.Sizes[1].Size = kwb.Alloc.Size.Total
	ly.LayImpl.Sizes[0].GrowSum = kwb.Styles.Grow
	ly.LayImpl.Sizes[0].GrowSum = kwb.Styles.Grow
	ly.Alloc.Size.SetContentToFit(ly.LayImpl.Sizes[0].Size, ly.Styles.Max.Dots())
}

//////////////////////////////////////////////////////////////////////
//		SizeDown

func (wb *WidgetBase) SizeDown(sc *Scene, iter int, allocTotal mat32.Vec2) bool {
	return wb.SizeDownWidget(sc, iter, allocTotal)
}

// SizeDownWidget is the standard widget implementation of SizeDown,
// computing updated Content size from given total allocation
// and giving that content to its parts if they exist.
func (wb *WidgetBase) SizeDownWidget(sc *Scene, iter int, allocTotal mat32.Vec2) bool {
	wb.Alloc.Total = allocTotal
	wb.Alloc.SetContentFromTotal(wb.Styles.BoxSpace())
	redo := wb.SizeDownParts(sc, iter, wb.Alloc.Content) // give our content to parts
	return redo
}

func (wb *WidgetBase) SizeDownParts(sc *Scene, iter int, allocTotal mat32.Vec2) bool {
	if wb.Parts == nil {
		return false
	}
	return wb.Parts.SizeDown(sc, iter, allocTotal)
}

func (ly *Layout) SizeDown(sc *Scene, iter int, allocTotal mat32.Vec2) bool {
	return ly.SizeDownLay(sc, iter, allocTotal)
}

// SizeDownLay is the Layout standard SizeDown pass, returning true if another
// iteration is required.  It allocates sizes to fit given parent-allocated
// total size.
func (ly *Layout) SizeDownLay(sc *Scene, iter int, allocTotal mat32.Vec2) bool {
	// todo: need to add gap sizes!
	ourTotal := ly.Alloc.Size.Total
	var totSz mat32.Vec2
	switch ly.Styles.Overflow {
	case styles.OverflowVisible:
		totSz = ourTotal.Max(allocTotal) // get everything we need or are given
		// no scrollbars
	case styles.OverflowHidden:
		totSz = allocTotal // get what we are given
		// no scrollbars
	case styles.OverflowAuto, styles.OverflowScroll:
		totSz = ourTotal.Max(allocTotal) // get everything we need or are given
		if ourTotal.X > allocTotal.X {
			ly.LayImpl.ScrollSize.Y = ly.Styles.ScrollBarWidth.Dots()
		}
		if ourTotal.Y > allocTotal.Y {
			ly.LayImpl.ScrollSize.X = ly.Styles.ScrollBarWidth.Dots()
		}
	}
	prevContent := ly.Alloc.Size.Content
	ly.Alloc.Size.Total = totSz
	extot := ly.LayImpl.ScrollSize.Add(ly.LayImpl.GapSize).Add(ly.Styles.BoxSpace())
	ly.Alloc.SetContentFromTotal(extot)

	conDiff := ly.Alloc.size.Content.Sub(prevContent)
	redo := conDiff.X != 0 || condDiff.Y != 0 // size changed
	if conDiff.X > 0 || conDiff.Y > 0 {
		ly.SizeDownGrow(conDiff)
		re := ly.SizeDownAlloc(sc, iter)
		if re {
			redo = true
		}
	}
	return redo
}

// SizeDownGrow grows the element sizes based on total extra
func (ly *Layout) SizeDownGrow(diff mat32.Vec2) {
	if ly.Styles.Display == styles.DisplayFlex && ly.Styles.Wrap {
		ly.SizeDownWrap() // first recompute wrap
		// todo: use special version of grow
	}
	for ma := mat32.X; ma < mat32.DimsN; ma++ { // main axis = X then Y
		ca := ma.OrthoDim()   // cross axis = Y then X
		extra := diff.Dim(ca) // col.Y = extra height for rows in this col; row.X = extra width for cols
		if extra <= 0 {
			continue
		}
		md := &ly.LayImpl.Sizes[mi] // col data; row data
		gsum := md.GrowSum.Dim(ci)  // col.Y = sum of row grows for col; row.X = sum of col grows for row
		if gsum <= 0 {
			continue
		}
		nca := mat32.Pointdim(ly.LayImpl.Cells, ca) // n rows; n cols
		for ci := 0; ci < nca; ci++ {
			cd := &ly.LayImpl.Sizes[ca][ci] // row; col
			gr := cd.GrowSum.Dim(ci)        // row.Y = sum of row grows for row
			ex := extra * (gr / gsum)
			fmt.Println(ly, "ma dim:", ma, "ca:", gr, "prop:", gr/gsum)
			cur := cd.Size.Dim(ci)
			cd.Size.SetDim(cur + ex)
		}
	}
}

// SizeDownAlloc calls size down on kids with newly allocated sizes.
// returns true if any report needing a redo.
func (ly *Layout) SizeDownAlloc(sc *Scene, iter int) bool {
	if ly.Styles.Display == styles.DisplayStacked {
		ly.SizeDownAllocStacked(sc)
		return
	}
	// todo: wrap needs special case
	return ly.SizeDownAllocCells(sc, iter)
}

// SizeDownAllocCells calls size down on kids with newly allocated sizes.
// returns true if any report needing a redo.
func (ly *Layout) SizeDownAllocCells(sc *Scene, iter int) bool {
	redo := false
	ly.WidgetKidsIter(func(i int, kwi Widget, kwb *WidgetBase) bool {
		ci := kwb.Alloc.Cell
		sz := kwb.Alloc.Size.Total
		xa := ly.LayImpl.Sizes[0][ci.X].X // column width
		ya := ly.LayImpl.Sizes[1][ci.Y].Y // row height
		asz := mat32.NewVec2(xa, ya)
		if asz != sz {
			fmt.Println(kwi, "td resize, was:", sz, "alloc:", asz)
			redo = true
		}
		re := kwi.SizeDown(sc, iter, asz)
		redo = redo || re
		return ki.Continue
	})
	return redo
}

//////////////////////////////////////////////////////////////////////
//		Position

// Position: uses the final sizes to position everything within layouts
// according to alignment settings.
func (wb *WidgetBase) Position(sc *Scene) {
	wb.PositionWidget(sc)
}

func (wb *WidgetBase) PositionWidget(sc *Scene) {
	wb.PositionParts(sc)
}

func (wb *WidgetBase) PositionParts(sc *Scene) {
	if wb.Parts == nil {
		return
	}
	wb.Parts.Position(sc)
}

// Position: uses the final sizes to position everything within layouts
// according to alignment settings.
func (ly *Layout) Position(sc *Scene) {
	ly.PositionLay(sc)
}

func (ly *Layout) PositionLay(sc *Scene) {
	// todo
}

//////////////////////////////////////////////////////////////////////
//		ScenePos

// todo: scene overwrites this to start fresh, rest get info from parent

// ScenePos: scene-based position and final BBox is computed based on
// parents accumulated position and scrollbar position.
// This step can be performed when scrolling after updating Scroll.
func (wb *WidgetBase) ScenePos(sc *Scene) {
	wb.ScenePosWidget(sc)
}

func (wb *WidgetBase) ScenePosWidget(sc *Scene) {
	// todo: do our bbox, pass content bbox which we compute
	wb.ScenePosParts(sc)
}

func (wb *WidgetBase) ScenePosParts(sc *Scene) {
	if wb.Parts == nil {
		return
	}
	wb.Parts.ScenePos(sc)
}

// ScenePos: scene-based position and final BBox is computed based on
// parents accumulated position and scrollbar position.
// This step can be performed when scrolling after updating Scroll.
func (ly *Layout) ScenePos(sc *Scene) {
	ly.ScenePosLay(sc)
}

func (ly *Layout) ScenePosLay(sc *Scene) {
	// todo
}

/*
// LayoutSharedDim implements calculations to layout for the shared dimension
// (i.e., Vertical for Horizontal layout). Returns pos and size.
func LayoutSharedDimImpl(ly *Layout, avail, need, pref, max float32, spc styles.SideFloats, al styles.Align) (pos, size float32) {
	usePref := true
	targ := pref
	extra := avail - targ
	if extra < -0.1 { // not fitting in pref, go with min
		usePref = false
		targ = need
		extra = avail - targ
	}
	extra = mat32.Max(extra, 0.0) // no negatives

	stretchNeed := false // stretch relative to need
	stretchMax := false  // only stretch Max = neg

	if usePref && extra > 0.0 { // have some stretch extra
		if max < 0.0 {
			stretchMax = true // only stretch those marked as infinitely stretchy
		}
	} else if extra > 0.0 { // extra relative to Need
		stretchNeed = true // stretch relative to need
	}

	if ly.Lay == LayoutHoriz || ly.Lay == LayoutHorizFlow {
		pos = spc.Pos().Dim(mat32.Y)
	} else {
		pos = spc.Pos().Dim(mat32.X)
	}

	size = need
	if usePref {
		size = pref
	}
	if stretchMax || stretchNeed {
		size += extra
	} else {
		if styles.IsAlignMiddle(al) {
			pos += 0.5 * extra
		} else if styles.IsAlignEnd(al) {
			pos += extra
		} else if al == styles.AlignJustify { // treat justify as stretch
			size += extra
		}
	}

	// if LayoutTrace {
	// 	fmt.Printf("ly %v avail: %v targ: %v, extra %v, strMax: %v, strNeed: %v, pos: %v size: %v spc: %v\n", ly.Nm, avail, targ, extra, stretchMax, stretchNeed, pos, size, spc)
	// }

	return
}

// LayoutSharedDim lays out items along a shared dimension, where all elements
// share the same space, e.g., Horiz for a Vert layout, and vice-versa.
func LayoutSharedDim(ly *Layout, dim mat32.Dims) {
	spc := ly.BoxSpace()
	avail := ly.LayState.Alloc.Size.Dim(dim) - spc.Size().Dim(dim)
	for i, c := range ly.Kids {
		if c == nil {
			continue
		}
		ni := c.(Widget).AsWidget()
		if ni == nil {
			continue
		}
		if ly.Lay == LayoutStacked && ly.Is(LayoutStackTopOnly) && i != ly.StackTop {
			continue
		}
		ni.StyMu.RLock()
		al := ni.Styles.AlignDim(dim)
		ni.StyMu.RUnlock()
		pref := ni.LayState.Size.Pref.Dim(dim)
		need := ni.LayState.Size.Need.Dim(dim)
		max := ni.LayState.Size.Max.Dim(dim)
		pos, size := LayoutSharedDimImpl(ly, avail, need, pref, max, spc, al)
		ni.LayState.Alloc.Size.SetDim(dim, size)
		ni.LayState.Alloc.PosRel.SetDim(dim, pos)
	}
}

// LayoutAlongDim lays out all children along given dim -- only affects that dim --
// e.g., use LayoutSharedDim for other dim.
func LayoutAlongDim(ly *Layout, dim mat32.Dims) {
	sz := len(ly.Kids)
	if sz == 0 {
		return
	}

	elspc := float32(sz-1) * ly.Styles.Gap.Dots
	al := ly.Styles.AlignDim(dim)
	spc := ly.BoxSpace()
	exspc := spc.Size().Dim(dim) + elspc
	avail := ly.LayState.Alloc.Size.Dim(dim) - exspc
	pref := ly.LayState.Size.Pref.Dim(dim) - exspc
	need := ly.LayState.Size.Need.Dim(dim) - exspc

	targ := pref
	usePref := true
	extra := avail - targ
	if extra < -0.1 { // not fitting in pref, go with need
		usePref = false
		targ = need
		extra = avail - targ
	}
	extra = mat32.Max(extra, 0.0) // no negatives

	nstretch := 0
	stretchTot := float32(0.0)
	stretchNeed := false        // stretch relative to need
	stretchMax := false         // only stretch Max = neg
	addSpace := false           // apply extra toward spacing -- for justify
	if usePref && extra > 0.0 { // have some stretch extra
		for _, c := range ly.Kids {
			if c == nil {
				continue
			}
			ni := c.(Widget).AsWidget()
			if ni == nil {
				continue
			}
			if ni.LayState.Size.HasMaxStretch(dim) { // negative = stretch
				nstretch++
				stretchTot += ni.LayState.Size.Pref.Dim(dim)
			}
		}
		if nstretch > 0 {
			stretchMax = true // only stretch those marked as infinitely stretchy
		}
	} else if extra > 0.0 { // extra relative to Need
		for _, c := range ly.Kids {
			if c == nil {
				continue
			}
			ni := c.(Widget).AsWidget()
			if ni == nil {
				continue
			}
			if ni.LayState.Size.HasMaxStretch(dim) || ni.LayState.Size.CanStretchNeed(dim) {
				nstretch++
				stretchTot += ni.LayState.Size.Pref.Dim(dim)
			}
		}
		if nstretch > 0 {
			stretchNeed = true // stretch relative to need
		}
	}

	extraSpace := float32(0.0)
	if sz > 1 && extra > 0.0 && al == styles.AlignJustify && !stretchNeed && !stretchMax {
		addSpace = true
		// if neither, then just distribute as spacing for justify
		extraSpace = extra / float32(sz-1)
	}

	// now arrange everyone
	pos := spc.Pos().Dim(dim)

	// todo: need a direction setting too
	if styles.IsAlignEnd(al) && !stretchNeed && !stretchMax {
		pos += extra
	}

	if LayoutTrace {
		fmt.Printf("Layout: %v Along dim %v, avail: %v elspc: %v need: %v pref: %v targ: %v, extra %v, strMax: %v, strNeed: %v, nstr %v, strTot %v\n", ly.Path(), dim, avail, elspc, need, pref, targ, extra, stretchMax, stretchNeed, nstretch, stretchTot)
	}

	for i, c := range ly.Kids {
		if c == nil {
			continue
		}
		ni := c.(Widget).AsWidget()
		if ni == nil {
			continue
		}
		size := ni.LayState.Size.Need.Dim(dim)
		if usePref {
			size = ni.LayState.Size.Pref.Dim(dim)
		}
		if stretchMax { // negative = stretch
			if ni.LayState.Size.HasMaxStretch(dim) { // in proportion to pref
				size += extra * (ni.LayState.Size.Pref.Dim(dim) / stretchTot)
			}
		} else if stretchNeed {
			if ni.LayState.Size.HasMaxStretch(dim) || ni.LayState.Size.CanStretchNeed(dim) {
				size += extra * (ni.LayState.Size.Pref.Dim(dim) / stretchTot)
			}
		} else if addSpace { // implies align justify
			if i > 0 {
				pos += extraSpace
			}
		}

		ni.LayState.Alloc.Size.SetDim(dim, size)
		ni.LayState.Alloc.PosRel.SetDim(dim, pos)
		if LayoutTrace {
			fmt.Printf("Layout: %v Child: %v, pos: %v, size: %v, need: %v, pref: %v\n", ly.Path(), ni.Nm, pos, size, ni.LayState.Size.Need.Dim(dim), ni.LayState.Size.Pref.Dim(dim))
		}
		pos += size + ly.Styles.Gap.Dots
	}
}

//////////////////////////////////////////////////////
//  Wrap

// GatherSizesFlow is size first pass: gather the size information from the children
func GatherSizesFlow(ly *Layout, iter int) {
	sz := len(ly.Kids)
	if sz == 0 {
		return
	}

	if iter > 0 {
		ly.ChildrenUpdateSizes() // essential to call this
		prv := ly.ChildSize
		ly.LayState.Size.Need = prv
		ly.LayState.Size.Pref = prv
		ly.LayState.Alloc.Size = prv
		ly.LayState.UpdateSizes() // enforce max and normal ordering, etc
		if LayoutTrace {
			fmt.Printf("Size:   %v iter 1 fix size: %v\n", ly.Path(), prv)
		}
		return
	}

	sumPref, sumNeed, maxPref, maxNeed := GatherSizesSumMax(ly)

	// for flow, the need is always *maxNeed* (i.e., a single item)
	// and the pref is based on styled pref estimate

	sdim := LaySummedDim(ly.Lay)
	odim := mat32.OtherDim(sdim)
	pref := ly.LayState.Size.Pref.Dim(sdim)
	// opref := ly.LayState.Size.Pref.Dim(odim) // not using

	if pref == 0 {
		pref = 6 * maxNeed.Dim(sdim) // 6 items preferred backup default
	}
	if pref == 0 {
		pref = 200 // final backstop
	}

	if LayoutTrace {
		fmt.Printf("Size:   %v flow pref start: %v\n", ly.Path(), pref)
	}

	sNeed := sumNeed.Dim(sdim)
	nNeed := float32(1)
	tNeed := sNeed
	oNeed := maxNeed.Dim(odim)
	if sNeed > pref {
		tNeed = pref
		nNeed = mat32.Ceil(sNeed / tNeed)
		oNeed *= nNeed
	}

	sPref := sumPref.Dim(sdim)
	nPref := float32(1)
	tPref := sPref
	oPref := maxPref.Dim(odim)
	if sPref > pref {
		if sNeed > pref {
			tPref = pref
			nPref = mat32.Ceil(sPref / tPref)
			oPref *= nPref
		} else {
			tPref = sNeed // go with need
		}
	}

	ly.LayState.Size.Need.SetMaxDim(sdim, maxNeed.Dim(sdim)) // Need min = max of single item
	ly.LayState.Size.Pref.SetMaxDim(sdim, tPref)
	ly.LayState.Size.Need.SetMaxDim(odim, oNeed)
	ly.LayState.Size.Pref.SetMaxDim(odim, oPref)

	spc := ly.BoxSpace()
	ly.LayState.Size.Need.SetAdd(spc.Size())
	ly.LayState.Size.Pref.SetAdd(spc.Size())

	elspc := float32(0.0)
	if sz >= 2 {
		elspc = float32(sz-1) * ly.Styles.Gap.Dots
	}
	if LaySumDim(ly.Lay, mat32.X) {
		ly.LayState.Size.Need.X += elspc
		ly.LayState.Size.Pref.X += elspc
	}
	if LaySumDim(ly.Lay, mat32.Y) {
		ly.LayState.Size.Need.Y += elspc
		ly.LayState.Size.Pref.Y += elspc
	}

	ly.LayState.UpdateSizes() // enforce max and normal ordering, etc
	if LayoutTrace {
		fmt.Printf("Size:   %v gather sizes need: %v, pref: %v, elspc: %v\n", ly.Path(), ly.LayState.Size.Need, ly.LayState.Size.Pref, elspc)
	}
}

// todo: grid does not process spans at all yet -- assumes = 1

// GatherSizesGrid is size first pass: gather the size information from the
// children, grid version
func GatherSizesGrid(ly *Layout) {
	if len(ly.Kids) == 0 {
		return
	}

	cols := ly.Styles.Columns
	rows := 0

	sz := len(ly.Kids)
	// collect overall size
	for _, c := range ly.Kids {
		if c == nil {
			continue
		}
		ni := c.(Widget).AsWidget()
		if ni == nil {
			continue
		}
		ni.StyMu.RLock()
		st := &ni.Styles
		ni.StyMu.RUnlock()
		if st.Col > 0 {
			cols = max(cols, st.Col+st.ColSpan)
		}
		if st.Row > 0 {
			rows = max(rows, st.Row+st.RowSpan)
		}
	}

	if cols == 0 {
		cols = int(mat32.Sqrt(float32(sz))) // whatever -- not well defined
	}
	if rows == 0 {
		rows = sz / cols
	}
	for rows*cols < sz { // not defined to have multiple items per cell -- make room for everyone
		rows++
	}

	ly.GridSize.X = cols
	ly.GridSize.Y = rows

	if len(ly.GridData[Row]) != rows {
		ly.GridData[Row] = make([]GridData, rows)
	}
	if len(ly.GridData[Col]) != cols {
		ly.GridData[Col] = make([]GridData, cols)
	}

	for i := range ly.GridData[Row] {
		gd := &ly.GridData[Row][i]
		gd.SizeNeed = 0
		gd.SizePref = 0
	}
	for i := range ly.GridData[Col] {
		gd := &ly.GridData[Col][i]
		gd.SizeNeed = 0
		gd.SizePref = 0
	}

	col := 0
	row := 0
	for _, c := range ly.Kids {
		if c == nil {
			continue
		}
		ni := c.(Widget).AsWidget()
		if ni == nil {
			continue
		}
		ni.LayState.UpdateSizes()
		ni.StyMu.RLock()
		st := &ni.Styles
		ni.StyMu.RUnlock()
		if st.Col > 0 {
			col = st.Col
		}
		if st.Row > 0 {
			row = st.Row
		}

		rgd := &(ly.GridData[Row][row])
		cgd := &(ly.GridData[Col][col])

		// todo: need to deal with span in sums..
		mat32.SetMax(&(rgd.SizeNeed), ni.LayState.Size.Need.Y)
		mat32.SetMax(&(rgd.SizePref), ni.LayState.Size.Pref.Y)
		mat32.SetMax(&(cgd.SizeNeed), ni.LayState.Size.Need.X)
		mat32.SetMax(&(cgd.SizePref), ni.LayState.Size.Pref.X)

		// for max: any -1 stretch dominates, else accumulate any max
		if rgd.SizeMax >= 0 {
			if ni.LayState.Size.Max.Y < 0 { // stretch
				rgd.SizeMax = -1
			} else {
				mat32.SetMax(&(rgd.SizeMax), ni.LayState.Size.Max.Y)
			}
		}
		if cgd.SizeMax >= 0 {
			if ni.LayState.Size.Max.Y < 0 { // stretch
				cgd.SizeMax = -1
			} else {
				mat32.SetMax(&(cgd.SizeMax), ni.LayState.Size.Max.X)
			}
		}

		col++
		if col >= cols { // todo: really only works if NO items specify row,col or ALL do..
			col = 0
			row++
			if row >= rows { // wrap-around.. no other good option
				row = 0
			}
		}
	}

	prefSizing := false
	sc := ly.Sc
	if sc != nil && sc.Is(ScPrefSizing) {
		prefSizing = ly.Styles.Overflow == styles.OverflowScroll // special case
	}

	// if there aren't existing prefs, we need to compute size
	if prefSizing || ly.LayState.Size.Pref.X == 0 || ly.LayState.Size.Pref.Y == 0 {
		sbw := ly.Styles.ScrollBarWidth.Dots
		maxRow := len(ly.GridData[Row])
		maxCol := len(ly.GridData[Col])
		if prefSizing {
			maxRow = min(LayoutPrefMaxRows, maxRow)
			maxCol = min(LayoutPrefMaxCols, maxCol)
		}

		// Y = sum across rows which have max's
		var sumPref, sumNeed mat32.Vec2
		for i := 0; i < maxRow; i++ {
			gd := ly.GridData[Row][i]
			sumNeed.SetAddDim(mat32.Y, gd.SizeNeed)
			sumPref.SetAddDim(mat32.Y, gd.SizePref)
		}
		// X = sum across cols which have max's
		for i := 0; i < maxCol; i++ {
			gd := ly.GridData[Col][i]
			sumNeed.SetAddDim(mat32.X, gd.SizeNeed)
			sumPref.SetAddDim(mat32.X, gd.SizePref)
		}

		if prefSizing {
			sumNeed.X += sbw
			sumNeed.Y += sbw
			sumPref.X += sbw
			sumPref.Y += sbw
		}

		if ly.LayState.Size.Pref.X == 0 || prefSizing {
			ly.LayState.Size.Need.X = mat32.Max(ly.LayState.Size.Need.X, sumNeed.X)
			ly.LayState.Size.Pref.X = mat32.Max(ly.LayState.Size.Pref.X, sumPref.X)
		} else { // use target size from style otherwise
			ly.LayState.Size.Need.X = ly.LayState.Size.Pref.X
		}
		if ly.LayState.Size.Pref.Y == 0 || prefSizing {
			ly.LayState.Size.Need.Y = mat32.Max(ly.LayState.Size.Need.Y, sumNeed.Y)
			ly.LayState.Size.Pref.Y = mat32.Max(ly.LayState.Size.Pref.Y, sumPref.Y)
		} else { // use target size from style otherwise
			ly.LayState.Size.Need.Y = ly.LayState.Size.Pref.Y
		}
	} else { // neither are zero so we use those directly
		ly.LayState.Size.Need = ly.LayState.Size.Pref
	}

	spc := ly.BoxSpace()
	ly.LayState.Size.Need.SetAdd(spc.Size())
	ly.LayState.Size.Pref.SetAdd(spc.Size())

	ly.LayState.Size.Need.X += float32(cols-1) * ly.Styles.Gap.Dots
	ly.LayState.Size.Pref.X += float32(cols-1) * ly.Styles.Gap.Dots
	ly.LayState.Size.Need.Y += float32(rows-1) * ly.Styles.Gap.Dots
	ly.LayState.Size.Pref.Y += float32(rows-1) * ly.Styles.Gap.Dots

	ly.LayState.UpdateSizes() // enforce max and normal ordering, etc
	if LayoutTrace {
		fmt.Printf("Size:   %v gather sizes grid need: %v, pref: %v\n", ly.Path(), ly.LayState.Size.Need, ly.LayState.Size.Pref)
	}
}

////////////////////////////////////////////////////////////////////////////////////////
//     Layout children

// LayoutFlow manages the flow layout along given dimension
// returns true if needs another iteration (only if iter == 0)
func LayoutFlow(ly *Layout, dim mat32.Dims, iter int) bool {
	ly.FlowBreaks = nil
	sz := len(ly.Kids)
	if sz == 0 {
		return false
	}

	elspc := float32(sz-1) * ly.Styles.Gap.Dots
	spc := ly.BoxSpace()
	exspc := spc.Size().Dim(dim) + elspc

	avail := ly.LayState.Alloc.Size.Dim(dim) - exspc
	odim := mat32.OtherDim(dim)

	// SidesTODO: might be odim
	pos := spc.Pos().Dim(dim)
	for i, c := range ly.Kids {
		if c == nil {
			continue
		}
		ni := c.(Widget).AsWidget()
		if ni == nil {
			continue
		}
		size := ni.LayState.Size.Need.Dim(dim)
		if pos+size > avail {
			ly.FlowBreaks = append(ly.FlowBreaks, i)
			pos = spc.Pos().Dim(dim)
		}
		ni.LayState.Alloc.Size.SetDim(dim, size)
		ni.LayState.Alloc.PosRel.SetDim(dim, pos)
		if LayoutTrace {
			fmt.Printf("Layout: %v Child: %v, pos: %v, size: %v, need: %v, pref: %v\n", ly.Path(), ni.Nm, pos, size, ni.LayState.Size.Need.Dim(dim), ni.LayState.Size.Pref.Dim(dim))
		}
		pos += size + ly.Styles.Gap.Dots
	}
	ly.FlowBreaks = append(ly.FlowBreaks, len(ly.Kids))

	nrows := len(ly.FlowBreaks)
	oavail := ly.LayState.Alloc.Size.Dim(odim) - exspc
	oavPerRow := oavail / float32(nrows)
	ci := 0
	rpos := float32(0)
	var nsz mat32.Vec2
	for _, bi := range ly.FlowBreaks {
		rmax := float32(0)
		for i := ci; i < bi; i++ {
			c := ly.Kids[i]
			if c == nil {
				continue
			}
			ni := c.(Widget).AsWidget()
			if ni == nil {
				continue
			}
			ni.StyMu.RLock()
			al := ni.Styles.AlignDim(odim)
			ni.StyMu.RUnlock()
			pref := ni.LayState.Size.Pref.Dim(odim)
			need := ni.LayState.Size.Need.Dim(odim)
			max := ni.LayState.Size.Max.Dim(odim)
			pos, size := LayoutSharedDimImpl(ly, oavPerRow, need, pref, max, spc, al)
			ni.LayState.Alloc.Size.SetDim(odim, size)
			ni.LayState.Alloc.PosRel.SetDim(odim, rpos+pos)
			rmax = mat32.Max(rmax, size)
			nsz.X = mat32.Max(nsz.X, ni.LayState.Alloc.PosRel.X+ni.LayState.Alloc.Size.X)
			nsz.Y = mat32.Max(nsz.Y, ni.LayState.Alloc.PosRel.Y+ni.LayState.Alloc.Size.Y)
		}
		rpos += rmax + ly.Styles.Gap.Dots
		ci = bi
	}
	ly.LayState.Size.Need = nsz
	ly.LayState.Size.Pref = nsz
	if LayoutTrace {
		fmt.Printf("Layout: %v Flow final size: %v\n", ly.Path(), nsz)
	}
	// if nrows == 1 {
	// 	return false
	// }
	return true
}

// LayoutGridDim lays out grid data along each dimension (row, Y; col, X),
// same as LayoutAlongDim.  For cols, X has width prefs of each -- turn that
// into an actual allocated width for each column, and likewise for rows.
func LayoutGridDim(ly *Layout, rowcol RowCol, dim mat32.Dims) {
	gds := ly.GridData[rowcol]
	sz := len(gds)
	if sz == 0 {
		return
	}
	elspc := float32(sz-1) * ly.Styles.Gap.Dots
	al := ly.Styles.AlignDim(dim)
	spc := ly.BoxSpace()
	exspc := spc.Size().Dim(dim) + elspc
	avail := ly.LayState.Alloc.Size.Dim(dim) - exspc
	pref := ly.LayState.Size.Pref.Dim(dim) - exspc
	need := ly.LayState.Size.Need.Dim(dim) - exspc

	targ := pref
	usePref := true
	extra := avail - targ
	if extra < -0.1 { // not fitting in pref, go with need
		usePref = false
		targ = need
		extra = avail - targ
	}
	extra = mat32.Max(extra, 0.0) // no negatives

	nstretch := 0
	stretchTot := float32(0.0)
	stretchNeed := false        // stretch relative to need
	stretchMax := false         // only stretch Max = neg
	addSpace := false           // apply extra toward spacing -- for justify
	if usePref && extra > 0.0 { // have some stretch extra
		for _, gd := range gds {
			if gd.SizeMax < 0 { // stretch
				nstretch++
				stretchTot += gd.SizePref
			}
		}
		if nstretch > 0 {
			stretchMax = true // only stretch those marked as infinitely stretchy
		}
	} else if extra > 0.0 { // extra relative to Need
		for _, gd := range gds {
			if gd.SizeMax < 0 || gd.SizePref > gd.SizeNeed {
				nstretch++
				stretchTot += gd.SizePref
			}
		}
		if nstretch > 0 {
			stretchNeed = true // stretch relative to need
		}
	}

	extraSpace := float32(0.0)
	if sz > 1 && extra > 0.0 && al == styles.AlignJustify && !stretchNeed && !stretchMax {
		addSpace = true
		// if neither, then just distribute as spacing for justify
		extraSpace = extra / float32(sz-1)
	}

	// now arrange everyone
	pos := spc.Pos().Dim(dim)

	// todo: need a direction setting too
	if styles.IsAlignEnd(al) && !stretchNeed && !stretchMax {
		pos += extra
	}

	if LayoutTrace {
		fmt.Printf("Layout Grid Dim: %v All on dim %v, avail: %v need: %v pref: %v targ: %v, extra %v, strMax: %v, strNeed: %v, nstr %v, strTot %v\n", ly.Path(), dim, avail, need, pref, targ, extra, stretchMax, stretchNeed, nstretch, stretchTot)
	}

	for i := range gds {
		gd := &gds[i]
		size := gd.SizeNeed
		if usePref {
			size = gd.SizePref
		}
		if stretchMax { // negative = stretch
			if gd.SizeMax < 0 { // in proportion to pref
				size += extra * (gd.SizePref / stretchTot)
			}
		} else if stretchNeed {
			if gd.SizeMax < 0 || gd.SizePref > gd.SizeNeed {
				size += extra * (gd.SizePref / stretchTot)
			}
		} else if addSpace { // implies align justify
			if i > 0 {
				pos += extraSpace
			}
		}

		gd.AllocSize = size
		gd.AllocPosRel = pos
		if LayoutTrace {
			fmt.Printf("Grid %v pos: %v, size: %v\n", rowcol, pos, size)
		}
		pos += size + ly.Styles.Gap.Dots
	}
}

// LayoutGridLay manages overall grid layout of children
func LayoutGridLay(ly *Layout) {
	sz := len(ly.Kids)
	if sz == 0 {
		return
	}

	LayoutGridDim(ly, Row, mat32.Y)
	LayoutGridDim(ly, Col, mat32.X)

	col := 0
	row := 0
	cols := ly.GridSize.X
	rows := ly.GridSize.Y

	if cols*rows != ly.NumChildren() {
		GatherSizesGrid(ly)
	}

	for _, c := range ly.Kids {
		if c == nil {
			continue
		}
		ni := c.(Widget).AsWidget()
		if ni == nil {
			continue
		}

		ni.StyMu.RLock()
		st := &ni.Styles
		ni.StyMu.RUnlock()
		if st.Col > 0 {
			col = st.Col
		}
		if st.Row > 0 {
			row = st.Row
		}

		{ // col, X dim
			dim := mat32.X
			gd := ly.GridData[Col][col]
			avail := gd.AllocSize
			al := st.AlignDim(dim)
			pref := ni.LayState.Size.Pref.Dim(dim)
			need := ni.LayState.Size.Need.Dim(dim)
			max := ni.LayState.Size.Max.Dim(dim)
			pos, size := LayoutSharedDimImpl(ly, avail, need, pref, max, styles.SideFloats{}, al)
			ni.LayState.Alloc.Size.SetDim(dim, size)
			ni.LayState.Alloc.PosRel.SetDim(dim, pos+gd.AllocPosRel)

		}
		{ // row, Y dim
			dim := mat32.Y
			gd := ly.GridData[Row][row]
			avail := gd.AllocSize
			al := st.AlignDim(dim)
			pref := ni.LayState.Size.Pref.Dim(dim)
			need := ni.LayState.Size.Need.Dim(dim)
			max := ni.LayState.Size.Max.Dim(dim)
			pos, size := LayoutSharedDimImpl(ly, avail, need, pref, max, styles.SideFloats{}, al)
			ni.LayState.Alloc.Size.SetDim(dim, size)
			ni.LayState.Alloc.PosRel.SetDim(dim, pos+gd.AllocPosRel)
		}

		if LayoutTrace {
			fmt.Printf("Layout: %v grid col: %v row: %v pos: %v size: %v\n", ly.Path(), col, row, ni.LayState.Alloc.PosRel, ni.LayState.Alloc.Size)
		}

		col++
		if col >= cols { // todo: really only works if NO items specify row,col or ALL do..
			col = 0
			row++
			if row >= rows { // wrap-around.. no other good option
				row = 0
			}
		}
	}
}

// FinalizeLayout is final pass through children to finalize the layout,
// computing summary size stats
func (ly *Layout) FinalizeLayout() {
	ly.ChildSize = mat32.Vec2Zero
	if ly.Lay == LayoutStacked && ly.Is(LayoutStackTopOnly) {
		sn, err := ly.ChildTry(ly.StackTop)
		if err != nil {
			return
		}
		ni := sn.(Widget).AsWidget()
		if ni == nil {
			return
		}
		ly.ChildSize.SetMax(ni.LayState.Alloc.PosRel.Add(ni.LayState.Alloc.Size))
		ni.LayState.Alloc.SizeOrig = ni.LayState.Alloc.Size
		return
	}
	for _, c := range ly.Kids {
		if c == nil {
			continue
		}
		ni := c.(Widget).AsWidget()
		if ni == nil {
			continue
		}
		ly.ChildSize.SetMax(ni.LayState.Alloc.PosRel.Add(ni.LayState.Alloc.Size))
		ni.LayState.Alloc.SizeOrig = ni.LayState.Alloc.Size
	}
}

*/
