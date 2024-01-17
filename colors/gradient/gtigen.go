// Code generated by "goki generate"; DO NOT EDIT.

package gradient

import (
	"cogentcore.org/core/colors"
	"cogentcore.org/core/gti"
	"cogentcore.org/core/mat32"
)

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/colors/gradient.Base", IDName: "base", Doc: "Base contains the data and logic common to all gradient types.", Directives: []gti.Directive{{Tool: "gti", Directive: "add", Args: []string{"-setters"}}}, Fields: []gti.Field{{Name: "Stops", Doc: "the stops for the gradient; use AddStop to add stops"}, {Name: "Spread", Doc: "the spread method used for the gradient if it stops before the end"}, {Name: "Blend", Doc: "the colorspace algorithm to use for blending colors"}, {Name: "Units", Doc: "the units to use for the gradient"}, {Name: "Box", Doc: "the bounding box of the object with the gradient; this is used when rendering\ngradients with [Units] of [ObjectBoundingBox]."}, {Name: "Transform", Doc: "Transform is the transformation matrix applied to the gradient's points."}, {Name: "ObjectMatrix", Doc: "ObjectMatrix is the computed effective object transformation matrix for a gradient\nwith [Units] of [ObjectBoundingBox]. It should not be set by end users."}}})

// SetSpread sets the [Base.Spread]:
// the spread method used for the gradient if it stops before the end
func (t *Base) SetSpread(v Spreads) *Base { t.Spread = v; return t }

// SetBlend sets the [Base.Blend]:
// the colorspace algorithm to use for blending colors
func (t *Base) SetBlend(v colors.BlendTypes) *Base { t.Blend = v; return t }

// SetUnits sets the [Base.Units]:
// the units to use for the gradient
func (t *Base) SetUnits(v Units) *Base { t.Units = v; return t }

// SetBox sets the [Base.Box]:
// the bounding box of the object with the gradient; this is used when rendering
// gradients with [Units] of [ObjectBoundingBox].
func (t *Base) SetBox(v mat32.Box2) *Base { t.Box = v; return t }

// SetTransform sets the [Base.Transform]:
// Transform is the transformation matrix applied to the gradient's points.
func (t *Base) SetTransform(v mat32.Mat2) *Base { t.Transform = v; return t }

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/colors/gradient.Linear", IDName: "linear", Doc: "Linear represents a linear gradient. It implements the [image.Image] interface.", Directives: []gti.Directive{{Tool: "gti", Directive: "add", Args: []string{"-setters"}}}, Embeds: []gti.Field{{Name: "Base"}}, Fields: []gti.Field{{Name: "Start", Doc: "the starting point of the gradient (x1 and y1 in SVG)"}, {Name: "End", Doc: "the ending point of the gradient (x2 and y2 in SVG)"}, {Name: "EffStart", Doc: "EffStart is the computed effective transformed starting point of the gradient.\nIt should not be set by end users."}, {Name: "EffEnd", Doc: "EffEnd is the computed effective transformed ending point of the gradient.\nIt should not be set by end users."}}})

// SetStart sets the [Linear.Start]:
// the starting point of the gradient (x1 and y1 in SVG)
func (t *Linear) SetStart(v mat32.Vec2) *Linear { t.Start = v; return t }

// SetEnd sets the [Linear.End]:
// the ending point of the gradient (x2 and y2 in SVG)
func (t *Linear) SetEnd(v mat32.Vec2) *Linear { t.End = v; return t }

// SetSpread sets the [Linear.Spread]
func (t *Linear) SetSpread(v Spreads) *Linear { t.Spread = v; return t }

// SetBlend sets the [Linear.Blend]
func (t *Linear) SetBlend(v colors.BlendTypes) *Linear { t.Blend = v; return t }

// SetUnits sets the [Linear.Units]
func (t *Linear) SetUnits(v Units) *Linear { t.Units = v; return t }

// SetBox sets the [Linear.Box]
func (t *Linear) SetBox(v mat32.Box2) *Linear { t.Box = v; return t }

// SetTransform sets the [Linear.Transform]
func (t *Linear) SetTransform(v mat32.Mat2) *Linear { t.Transform = v; return t }

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/colors/gradient.Radial", IDName: "radial", Doc: "Radial represents a radial gradient. It implements the [image.Image] interface.", Directives: []gti.Directive{{Tool: "gti", Directive: "add", Args: []string{"-setters"}}}, Embeds: []gti.Field{{Name: "Base"}}, Fields: []gti.Field{{Name: "Center", Doc: "the center point of the gradient (cx and cy in SVG)"}, {Name: "Focal", Doc: "the focal point of the gradient (fx and fy in SVG)"}, {Name: "Radius", Doc: "the radius of the gradient (rx and ry in SVG)"}}})

// SetCenter sets the [Radial.Center]:
// the center point of the gradient (cx and cy in SVG)
func (t *Radial) SetCenter(v mat32.Vec2) *Radial { t.Center = v; return t }

// SetFocal sets the [Radial.Focal]:
// the focal point of the gradient (fx and fy in SVG)
func (t *Radial) SetFocal(v mat32.Vec2) *Radial { t.Focal = v; return t }

// SetRadius sets the [Radial.Radius]:
// the radius of the gradient (rx and ry in SVG)
func (t *Radial) SetRadius(v mat32.Vec2) *Radial { t.Radius = v; return t }

// SetSpread sets the [Radial.Spread]
func (t *Radial) SetSpread(v Spreads) *Radial { t.Spread = v; return t }

// SetBlend sets the [Radial.Blend]
func (t *Radial) SetBlend(v colors.BlendTypes) *Radial { t.Blend = v; return t }

// SetUnits sets the [Radial.Units]
func (t *Radial) SetUnits(v Units) *Radial { t.Units = v; return t }

// SetBox sets the [Radial.Box]
func (t *Radial) SetBox(v mat32.Box2) *Radial { t.Box = v; return t }

// SetTransform sets the [Radial.Transform]
func (t *Radial) SetTransform(v mat32.Mat2) *Radial { t.Transform = v; return t }
