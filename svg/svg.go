// Copyright (c) 2018, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package svg

//go:generate core generate

import (
	"image"
	"image/color"
	"strings"
	"sync"

	"github.com/naruse666/core/base/iox/imagex"
	"github.com/naruse666/core/colors"
	"github.com/naruse666/core/math32"
	"github.com/naruse666/core/paint"
	"github.com/naruse666/core/styles"
	"github.com/naruse666/core/styles/units"
	"github.com/naruse666/core/tree"
)

// SVG is an SVG object.
type SVG struct {
	// Name is the name of the SVG -- e.g., the filename if loaded
	Name string

	// the title of the svg
	Title string `xml:"title"`

	// the description of the svg
	Desc string `xml:"desc"`

	// Background is the image/color to fill the background with,
	// if any.
	Background image.Image

	// Color can be set to provide a default Fill and Stroke Color value
	Color image.Image

	// Size is size of image, Pos is offset within any parent viewport.
	// Node bounding boxes are based on 0 Pos offset within Pixels image
	Geom math32.Geom2DInt

	// physical width of the drawing, e.g., when printed.
	// Does not affect rendering, metadata.
	PhysicalWidth units.Value

	// physical height of the drawing, e.g., when printed.
	// Does not affect rendering, metadata.
	PhysicalHeight units.Value

	// InvertY, when applying the ViewBox transform, also flip the Y axis so that
	// the smallest Y value is at the bottom of the SVG box,
	// instead of being at the top as it is by default.
	InvertY bool

	// Translate specifies a translation to apply beyond what is specified in the SVG,
	// and its ViewBox transform.
	Translate math32.Vector2

	// Scale specifies a zoom scale factor to apply beyond what is specified in the SVG,
	// and its ViewBox transform.
	Scale float32

	// render state for rendering
	RenderState paint.State `copier:"-" json:"-" xml:"-" edit:"-"`

	// live pixels that we render into
	Pixels *image.RGBA `copier:"-" json:"-" xml:"-" edit:"-"`

	// all defs defined elements go here (gradients, symbols, etc)
	Defs *Group

	// Root is the root of the svg tree, which has the top-level viewbox and styles.
	Root *Root

	// map of def names to index. uses starting index to find element.
	// always updated after each search.
	DefIndexes map[string]int `display:"-" json:"-" xml:"-"`

	// map of unique numeric ids for all elements.
	// Used for allocating new unique id numbers, appended to end of elements.
	// See NewUniqueID, GatherIDs
	UniqueIDs map[int]struct{} `display:"-" json:"-" xml:"-"`

	// flag is set when the SVG is rendering
	IsRendering bool

	// mutex for protecting rendering
	RenderMu sync.Mutex `display:"-" json:"-" xml:"-"`
}

// NewSVG creates a SVG with Pixels Image of the specified width and height
func NewSVG(width, height int) *SVG {
	sv := &SVG{}
	sv.Config(width, height)
	return sv
}

// Config configures the SVG, setting image to given size
// and initializing all relevant fields.
func (sv *SVG) Config(width, height int) {
	sz := image.Point{width, height}
	sv.Geom.Size = sz
	sv.Scale = 1
	sv.Pixels = image.NewRGBA(image.Rectangle{Max: sz})
	sv.RenderState.Init(width, height, sv.Pixels)
	sv.Root = NewRoot()
	sv.Root.SetName("svg")
	sv.Defs = NewGroup()
	sv.Defs.SetName("defs")
}

// Resize resizes the viewport, creating a new image -- updates Geom Size
func (sv *SVG) Resize(nwsz image.Point) {
	if nwsz.X == 0 || nwsz.Y == 0 {
		return
	}
	if sv.Root == nil || sv.Root.This == nil {
		sv.Config(nwsz.X, nwsz.Y)
		return
	}
	if sv.Pixels != nil {
		ib := sv.Pixels.Bounds().Size()
		if ib == nwsz {
			sv.Geom.Size = nwsz // make sure
			return              // already good
		}
	}
	sv.Pixels = image.NewRGBA(image.Rectangle{Max: nwsz})
	sv.RenderState.Init(nwsz.X, nwsz.Y, sv.Pixels)
	sv.Geom.Size = nwsz // make sure
}

// DeleteAll deletes any existing elements in this svg
func (sv *SVG) DeleteAll() {
	if sv.Root == nil || sv.Root.This == nil {
		return
	}
	sv.Root.Paint.Defaults()
	sv.Root.DeleteChildren()
	sv.Defs.DeleteChildren()
}

// Base returns the current Color activated in the context.
// Color has support for special color names that are relative to
// this current color.
func (sv *SVG) Base() color.RGBA {
	return colors.AsRGBA(colors.ToUniform(sv.Background))
}

// ImageByURL finds a Node by an element name (URL-like path), and
// attempts to convert it to an [image.Image].
// Used for color styling based on url() value.
func (sv *SVG) ImageByURL(url string) image.Image {
	// TODO(kai): support taking snapshot of element as image in SVG.ImageByURL
	if sv == nil {
		return nil
	}
	val := url[4:]
	val = strings.TrimPrefix(strings.TrimSuffix(val, ")"), "#")
	def := sv.FindDefByName(val)
	if def != nil {
		if grad, ok := def.(*Gradient); ok {
			return grad.Grad
		}
	}
	ne := sv.FindNamedElement(val)
	if grad, ok := ne.(*Gradient); ok {
		return grad.Grad
	}
	return nil
}

func (sv *SVG) Style() {
	// set isDef
	sv.Defs.WalkDown(func(n tree.Node) bool {
		sn := n.(Node)
		sn.AsNodeBase().isDef = true
		sn.AsNodeBase().Style(sv)
		return tree.Continue
	})

	sv.Root.Paint.Defaults()
	if sv.Color != nil {
		// TODO(kai): consider handling non-uniform colors here
		c := colors.ToUniform(sv.Color)
		sv.Root.SetColorProperties("stroke", colors.AsHex(c))
		sv.Root.SetColorProperties("fill", colors.AsHex(c))
	}
	sv.SetUnitContext(&sv.Root.Paint, math32.Vector2{}, math32.Vector2{})

	sv.Root.WalkDown(func(k tree.Node) bool {
		sn := k.(Node)
		sn.AsNodeBase().Style(sv)
		return tree.Continue
	})
}

func (sv *SVG) Render() {
	sv.RenderMu.Lock()
	sv.IsRendering = true

	sv.Style()
	sv.SetRootTransform()

	rs := &sv.RenderState
	rs.PushBounds(sv.Pixels.Bounds())
	if sv.Background != nil {
		sv.FillViewport()
	}
	sv.Root.Render(sv)
	rs.PopBounds()
	sv.RenderMu.Unlock()
	sv.IsRendering = false
}

func (sv *SVG) FillViewport() {
	pc := &paint.Context{&sv.RenderState, &sv.Root.Paint}
	pc.FillBox(math32.Vector2{}, math32.FromPoint(sv.Geom.Size), sv.Background)
}

// SetRootTransform sets the Root node transform based on ViewBox, Translate, Scale
// parameters set on the SVG object.
func (sv *SVG) SetRootTransform() {
	vb := &sv.Root.ViewBox
	box := math32.FromPoint(sv.Geom.Size)
	if vb.Size.X == 0 {
		vb.Size.X = sv.PhysicalWidth.Dots
	}
	if vb.Size.Y == 0 {
		vb.Size.Y = sv.PhysicalHeight.Dots
	}
	_, trans, scale := vb.Transform(box)
	if sv.InvertY {
		scale.Y *= -1
	}
	trans.SetSub(vb.Min)
	trans.SetAdd(sv.Translate)
	scale.SetMulScalar(sv.Scale)
	pc := &sv.Root.Paint
	pc.Transform = pc.Transform.Scale(scale.X, scale.Y).Translate(trans.X, trans.Y)
	if sv.InvertY {
		pc.Transform.Y0 = -pc.Transform.Y0
	}
}

// SetDPITransform sets a scaling transform to compensate for
// a given LogicalDPI factor.
// svg rendering is done within a 96 DPI context.
func (sv *SVG) SetDPITransform(logicalDPI float32) {
	pc := &sv.Root.Paint
	dpisc := logicalDPI / 96.0
	pc.Transform = math32.Scale2D(dpisc, dpisc)
}

// SavePNG saves the Pixels to a PNG file
func (sv *SVG) SavePNG(fname string) error {
	return imagex.Save(sv.Pixels, fname)
}

// Root represents the root of an SVG tree.
type Root struct {
	Group

	// ViewBox defines the coordinate system for the drawing.
	// These units are mapped into the screen space allocated
	// for the SVG during rendering.
	ViewBox ViewBox
}

func (g *Root) SVGName() string { return "svg" }

func (g *Root) EnforceSVGName() bool { return false }

func (g *Root) NodeBBox(sv *SVG) image.Rectangle {
	// todo: return viewbox
	return sv.Geom.SizeRect()
}

// SetUnitContext sets the unit context based on size of viewport, element,
// and parent element (from bbox) and then caches everything out in terms of raw pixel
// dots for rendering -- call at start of render
func (sv *SVG) SetUnitContext(pc *styles.Paint, el, parent math32.Vector2) {
	pc.UnitContext.Defaults()
	pc.UnitContext.DPI = 96 // paint (SVG) context is always 96 = 1to1
	if sv.RenderState.Image != nil {
		sz := sv.RenderState.Image.Bounds().Size()
		pc.UnitContext.SetSizes(float32(sz.X), float32(sz.Y), el.X, el.Y, parent.X, parent.Y)
	} else {
		pc.UnitContext.SetSizes(0, 0, el.X, el.Y, parent.X, parent.Y)
	}
	pc.FontStyle.SetUnitContext(&pc.UnitContext)
	pc.ToDots()
}
