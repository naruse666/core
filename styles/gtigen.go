// Code generated by "goki generate"; DO NOT EDIT.

package styles

import (
	"goki.dev/gti"
	"goki.dev/ordmap"
)

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/girl/styles.Border",
	ShortName: "styles.Border",
	IDName:    "border",
	Doc:       "Border contains style parameters for borders",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Style", &gti.Field{Name: "Style", Type: "goki.dev/girl/styles.Sides[goki.dev/girl/styles.BorderStyles]", LocalType: "Sides[BorderStyles]", Doc: "prop: border-style = how to draw the border", Directives: gti.Directives{}, Tag: "xml:\"style\""}},
		{"Width", &gti.Field{Name: "Width", Type: "goki.dev/girl/styles.SideValues", LocalType: "SideValues", Doc: "prop: border-width = width of the border", Directives: gti.Directives{}, Tag: "xml:\"width\""}},
		{"Radius", &gti.Field{Name: "Radius", Type: "goki.dev/girl/styles.SideValues", LocalType: "SideValues", Doc: "prop: border-radius = rounding of the corners", Directives: gti.Directives{}, Tag: "xml:\"radius\""}},
		{"Color", &gti.Field{Name: "Color", Type: "goki.dev/girl/styles.SideColors", LocalType: "SideColors", Doc: "prop: border-color = color of the border", Directives: gti.Directives{}, Tag: "xml:\"color\""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/girl/styles.Shadow",
	ShortName: "styles.Shadow",
	IDName:    "shadow",
	Doc:       "style parameters for shadows",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"HOffset", &gti.Field{Name: "HOffset", Type: "goki.dev/girl/units.Value", LocalType: "units.Value", Doc: "prop: .h-offset = horizontal offset of shadow -- positive = right side, negative = left side", Directives: gti.Directives{}, Tag: "xml:\".h-offset\""}},
		{"VOffset", &gti.Field{Name: "VOffset", Type: "goki.dev/girl/units.Value", LocalType: "units.Value", Doc: "prop: .v-offset = vertical offset of shadow -- positive = below, negative = above", Directives: gti.Directives{}, Tag: "xml:\".v-offset\""}},
		{"Blur", &gti.Field{Name: "Blur", Type: "goki.dev/girl/units.Value", LocalType: "units.Value", Doc: "prop: .blur = blur radius -- higher numbers = more blurry", Directives: gti.Directives{}, Tag: "xml:\".blur\""}},
		{"Spread", &gti.Field{Name: "Spread", Type: "goki.dev/girl/units.Value", LocalType: "units.Value", Doc: "prop: .spread = spread radius -- positive number increases size of shadow, negative decreases size", Directives: gti.Directives{}, Tag: "xml:\".spread\""}},
		{"Color", &gti.Field{Name: "Color", Type: "image/color.RGBA", LocalType: "color.RGBA", Doc: "prop: .color = color of the shadow", Directives: gti.Directives{}, Tag: "xml:\".color\""}},
		{"Inset", &gti.Field{Name: "Inset", Type: "bool", LocalType: "bool", Doc: "prop: .inset = shadow is inset within box instead of outset outside of box", Directives: gti.Directives{}, Tag: "xml:\".inset\""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/girl/styles.Font",
	ShortName: "styles.Font",
	IDName:    "font",
	Doc:       "Font contains all font styling information.\nMost of font information is inherited.\nFont does not include all information needed\nfor rendering -- see [FontRender] for that.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Size", &gti.Field{Name: "Size", Type: "goki.dev/girl/units.Value", LocalType: "units.Value", Doc: "prop: font-size (inherited) = size of font to render -- convert to points when getting font to use", Directives: gti.Directives{}, Tag: "xml:\"font-size\" inherit:\"true\""}},
		{"Family", &gti.Field{Name: "Family", Type: "string", LocalType: "string", Doc: "prop: font-family = font family -- ordered list of comma-separated names from more general to more specific to use -- use split on , to parse", Directives: gti.Directives{}, Tag: "xml:\"font-family\" inherit:\"true\""}},
		{"Style", &gti.Field{Name: "Style", Type: "goki.dev/girl/styles.FontStyles", LocalType: "FontStyles", Doc: "prop: font-style (inherited) = style -- normal, italic, etc", Directives: gti.Directives{}, Tag: "xml:\"font-style\" inherit:\"true\""}},
		{"Weight", &gti.Field{Name: "Weight", Type: "goki.dev/girl/styles.FontWeights", LocalType: "FontWeights", Doc: "prop: font-weight (inherited) = weight: normal, bold, etc", Directives: gti.Directives{}, Tag: "xml:\"font-weight\" inherit:\"true\""}},
		{"Stretch", &gti.Field{Name: "Stretch", Type: "goki.dev/girl/styles.FontStretch", LocalType: "FontStretch", Doc: "prop: font-stretch = font stretch / condense options", Directives: gti.Directives{}, Tag: "xml:\"font-stretch\" inherit:\"true\""}},
		{"Variant", &gti.Field{Name: "Variant", Type: "goki.dev/girl/styles.FontVariants", LocalType: "FontVariants", Doc: "prop: font-variant = normal or small caps", Directives: gti.Directives{}, Tag: "xml:\"font-variant\" inherit:\"true\""}},
		{"Deco", &gti.Field{Name: "Deco", Type: "goki.dev/girl/styles.TextDecorations", LocalType: "TextDecorations", Doc: "prop: text-decoration = underline, line-through, etc -- not inherited", Directives: gti.Directives{}, Tag: "xml:\"text-decoration\""}},
		{"Shift", &gti.Field{Name: "Shift", Type: "goki.dev/girl/styles.BaselineShifts", LocalType: "BaselineShifts", Doc: "prop: baseline-shift = super / sub script -- not inherited", Directives: gti.Directives{}, Tag: "xml:\"baseline-shift\""}},
		{"Face", &gti.Field{Name: "Face", Type: "*goki.dev/girl/styles.FontFace", LocalType: "*FontFace", Doc: "full font information including enhanced metrics and actual font codes for drawing text -- this is a pointer into FontLibrary of loaded fonts", Directives: gti.Directives{}, Tag: "view:\"-\""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/girl/styles.FontRender",
	ShortName: "styles.FontRender",
	IDName:    "font-render",
	Doc:       "FontRender contains all font styling information\nthat is needed for SVG text rendering. It is passed to\nPaint and Style functions. It should typically not be\nused by end-user code -- see [Font] for that.\nIt stores all values as pointers so that they correspond\nto the values of the style object it was derived from.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Color", &gti.Field{Name: "Color", Type: "image/color.RGBA", LocalType: "color.RGBA", Doc: "prop: color (inherited) = text color -- also defines the currentColor variable value", Directives: gti.Directives{}, Tag: "xml:\"color\" inherit:\"true\""}},
		{"BackgroundColor", &gti.Field{Name: "BackgroundColor", Type: "goki.dev/colors.Full", LocalType: "colors.Full", Doc: "prop: background-color = background color -- not inherited, transparent by default", Directives: gti.Directives{}, Tag: "xml:\"background-color\""}},
		{"Opacity", &gti.Field{Name: "Opacity", Type: "float32", LocalType: "float32", Doc: "prop: opacity = alpha value to apply to the foreground and background of this element and all of its children", Directives: gti.Directives{}, Tag: "xml:\"opacity\""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Font", &gti.Field{Name: "Font", Type: "goki.dev/girl/styles.Font", LocalType: "Font", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/girl/styles.FontFace",
	ShortName: "styles.FontFace",
	IDName:    "font-face",
	Doc:       "FontFace is our enhanced Font Face structure which contains the enhanced computed\nmetrics in addition to the font.Face face",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Name", &gti.Field{Name: "Name", Type: "string", LocalType: "string", Doc: "The full FaceName that the font is accessed by", Directives: gti.Directives{}, Tag: ""}},
		{"Size", &gti.Field{Name: "Size", Type: "int", LocalType: "int", Doc: "The integer font size in raw dots", Directives: gti.Directives{}, Tag: ""}},
		{"Face", &gti.Field{Name: "Face", Type: "golang.org/x/image/font.Face", LocalType: "font.Face", Doc: "The system image.Font font rendering interface", Directives: gti.Directives{}, Tag: ""}},
		{"Metrics", &gti.Field{Name: "Metrics", Type: "goki.dev/girl/styles.FontMetrics", LocalType: "FontMetrics", Doc: "enhanced metric information for the font", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/girl/styles.FontMetrics",
	ShortName: "styles.FontMetrics",
	IDName:    "font-metrics",
	Doc:       "FontMetrics are our enhanced dot-scale font metrics compared to what is available in\nthe standard font.Metrics lib, including Ex and Ch being defined in terms of\nthe actual letter x and 0",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Height", &gti.Field{Name: "Height", Type: "float32", LocalType: "float32", Doc: "reference 1.0 spacing line height of font in dots -- computed from font as ascent + descent + lineGap, where lineGap is specified by the font as the recommended line spacing", Directives: gti.Directives{}, Tag: ""}},
		{"Em", &gti.Field{Name: "Em", Type: "float32", LocalType: "float32", Doc: "Em size of font -- this is NOT actually the width of the letter M, but rather the specified point size of the font (in actual display dots, not points) -- it does NOT include the descender and will not fit the entire height of the font", Directives: gti.Directives{}, Tag: ""}},
		{"Ex", &gti.Field{Name: "Ex", Type: "float32", LocalType: "float32", Doc: "Ex size of font -- this is the actual height of the letter x in the font", Directives: gti.Directives{}, Tag: ""}},
		{"Ch", &gti.Field{Name: "Ch", Type: "float32", LocalType: "float32", Doc: "Ch size of font -- this is the actual width of the 0 glyph in the font", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/girl/styles.Paint",
	ShortName: "styles.Paint",
	IDName:    "paint",
	Doc:       "Paint provides the styling parameters for SVG-style rendering",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Off", &gti.Field{Name: "Off", Type: "bool", LocalType: "bool", Doc: "prop: display:none -- node and everything below it are off, non-rendering", Directives: gti.Directives{}, Tag: ""}},
		{"Display", &gti.Field{Name: "Display", Type: "bool", LocalType: "bool", Doc: "todo big enum of how to display item -- controls layout etc", Directives: gti.Directives{}, Tag: "xml:\"display\""}},
		{"StrokeStyle", &gti.Field{Name: "StrokeStyle", Type: "goki.dev/girl/styles.Stroke", LocalType: "Stroke", Doc: "stroke (line drawing) parameters", Directives: gti.Directives{}, Tag: ""}},
		{"FillStyle", &gti.Field{Name: "FillStyle", Type: "goki.dev/girl/styles.Fill", LocalType: "Fill", Doc: "fill (region filling) parameters", Directives: gti.Directives{}, Tag: ""}},
		{"FontStyle", &gti.Field{Name: "FontStyle", Type: "goki.dev/girl/styles.FontRender", LocalType: "FontRender", Doc: "font also has global opacity setting, along with generic color, background-color settings, which can be copied into stroke / fill as needed", Directives: gti.Directives{}, Tag: ""}},
		{"TextStyle", &gti.Field{Name: "TextStyle", Type: "goki.dev/girl/styles.Text", LocalType: "Text", Doc: "font also has global opacity setting, along with generic color, background-color settings, which can be copied into stroke / fill as needed", Directives: gti.Directives{}, Tag: ""}},
		{"VecEff", &gti.Field{Name: "VecEff", Type: "goki.dev/girl/styles.VectorEffects", LocalType: "VectorEffects", Doc: "prop: vector-effect = various rendering special effects settings", Directives: gti.Directives{}, Tag: "xml:\"vector-effect\""}},
		{"XForm", &gti.Field{Name: "XForm", Type: "goki.dev/mat32/v2.Mat2", LocalType: "mat32.Mat2", Doc: "prop: transform = our additions to transform -- pushed to render state", Directives: gti.Directives{}, Tag: "xml:\"transform\""}},
		{"UnContext", &gti.Field{Name: "UnContext", Type: "goki.dev/girl/units.Context", LocalType: "units.Context", Doc: "units context -- parameters necessary for anchoring relative units", Directives: gti.Directives{}, Tag: "xml:\"-\""}},
		{"StyleSet", &gti.Field{Name: "StyleSet", Type: "bool", LocalType: "bool", Doc: "have the styles already been set?", Directives: gti.Directives{}, Tag: ""}},
		{"PropsNil", &gti.Field{Name: "PropsNil", Type: "bool", LocalType: "bool", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		{"dotsSet", &gti.Field{Name: "dotsSet", Type: "bool", LocalType: "bool", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		{"lastUnCtxt", &gti.Field{Name: "lastUnCtxt", Type: "goki.dev/girl/units.Context", LocalType: "units.Context", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/girl/styles.Sides",
	ShortName: "styles.Sides",
	IDName:    "sides",
	Doc:       "Sides contains values for each side or corner of a box.\nIf Sides contains sides, the struct field names correspond\ndirectly to the side values (ie: Top = top side value).\nIf Sides contains corners, the struct field names correspond\nto the corners as follows: Top = top left, Right = top right,\nBottom = bottom right, Left = bottom left.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Top", &gti.Field{Name: "Top", Type: "T", LocalType: "T", Doc: "top/top-left value", Directives: gti.Directives{}, Tag: "xml:\"top\""}},
		{"Right", &gti.Field{Name: "Right", Type: "T", LocalType: "T", Doc: "right/top-right value", Directives: gti.Directives{}, Tag: "xml:\"right\""}},
		{"Bottom", &gti.Field{Name: "Bottom", Type: "T", LocalType: "T", Doc: "bottom/bottom-right value", Directives: gti.Directives{}, Tag: "xml:\"bottom\""}},
		{"Left", &gti.Field{Name: "Left", Type: "T", LocalType: "T", Doc: "left/bottom-left value", Directives: gti.Directives{}, Tag: "xml:\"left\""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/girl/styles.SideValues",
	ShortName: "styles.SideValues",
	IDName:    "side-values",
	Doc:       "SideValues contains units.Value values for each side/corner of a box",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Value]", &gti.Field{Name: "Value]", Type: "goki.dev/girl/styles.Sides[goki.dev/girl/units.Value]", LocalType: "Sides[units.Value]", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/girl/styles.SideFloats",
	ShortName: "styles.SideFloats",
	IDName:    "side-floats",
	Doc:       "SideFloats contains float32 values for each side/corner of a box",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Sides[float32]", &gti.Field{Name: "Sides[float32]", Type: "goki.dev/girl/styles.Sides[float32]", LocalType: "Sides[float32]", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/girl/styles.SideColors",
	ShortName: "styles.SideColors",
	IDName:    "side-colors",
	Doc:       "SideColors contains color values for each side/corner of a box",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"RGBA]", &gti.Field{Name: "RGBA]", Type: "goki.dev/girl/styles.Sides[image/color.RGBA]", LocalType: "Sides[color.RGBA]", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/girl/styles.Style",
	ShortName: "styles.Style",
	IDName:    "style",
	Doc:       "Style has all the CSS-based style elements -- used for widget-type GUI objects.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"State", &gti.Field{Name: "State", Type: "goki.dev/girl/states.States", LocalType: "states.States", Doc: "State holds style-relevant state flags, for convenient styling access,\ngiven that styles typically depend on element states.", Directives: gti.Directives{}, Tag: ""}},
		{"Abilities", &gti.Field{Name: "Abilities", Type: "goki.dev/girl/abilities.Abilities", LocalType: "abilities.Abilities", Doc: "Abilities specifies the abilities of this element, which determine\nwhich kinds of states the element can express.\nThis is used by the goosi/events system.  Putting this info next\nto the State info makes it easy to configure and manage.", Directives: gti.Directives{}, Tag: ""}},
		{"Display", &gti.Field{Name: "Display", Type: "goki.dev/girl/styles.Display", LocalType: "Display", Doc: "Display controls how items are displayed, in terms of layout", Directives: gti.Directives{}, Tag: ""}},
		{"Cursor", &gti.Field{Name: "Cursor", Type: "goki.dev/cursors.Cursor", LocalType: "cursors.Cursor", Doc: "the cursor to switch to upon hovering over the element (inherited)", Directives: gti.Directives{}, Tag: ""}},
		{"ZIndex", &gti.Field{Name: "ZIndex", Type: "int", LocalType: "int", Doc: "ordering factor for rendering depth -- lower numbers rendered first.\nSort children according to this factor", Directives: gti.Directives{}, Tag: ""}},
		{"Align", &gti.Field{Name: "Align", Type: "goki.dev/girl/styles.XY[goki.dev/girl/styles.Align]", LocalType: "XY[Align]", Doc: "Align specifies the X, Y alignment of widget elements within a container", Directives: gti.Directives{}, Tag: "view:\"inline\""}},
		{"Pos", &gti.Field{Name: "Pos", Type: "goki.dev/girl/units.XY", LocalType: "units.XY", Doc: "position is only used for Layout = Nil cases", Directives: gti.Directives{}, Tag: "view:\"inline\""}},
		{"Min", &gti.Field{Name: "Min", Type: "goki.dev/girl/units.XY", LocalType: "units.XY", Doc: "Min is the minimum size of the actual content, exclusive of additional space\nfrom padding, border, margin; 0 = default is sum of Min for all content\n(which _includes_ space for all sub-elements).\nThis is equivalent to the Basis for the CSS flex styling model.", Directives: gti.Directives{}, Tag: "view:\"inline\""}},
		{"Max", &gti.Field{Name: "Max", Type: "goki.dev/girl/units.XY", LocalType: "units.XY", Doc: "Max is the maximum size of the actual content, exclusive of additional space\nfrom padding, border, margin; 0 = default provides no Max size constraint", Directives: gti.Directives{}, Tag: "view:\"inline\""}},
		{"Grow", &gti.Field{Name: "Grow", Type: "goki.dev/mat32/v2.Vec2", LocalType: "mat32.Vec2", Doc: "Grow is the proportional amount that the element can grow (stretch)\nif there is more space available.  0 = default = no growth.\nExtra available space is allocated as: Grow / sum (all Grow)", Directives: gti.Directives{}, Tag: ""}},
		{"Padding", &gti.Field{Name: "Padding", Type: "goki.dev/girl/styles.SideValues", LocalType: "SideValues", Doc: "Padding is the transparent space around central content of box,\nwhich is _included_ in the size of the standard box rendering.", Directives: gti.Directives{}, Tag: "view:\"inline\""}},
		{"Margin", &gti.Field{Name: "Margin", Type: "goki.dev/girl/styles.SideValues", LocalType: "SideValues", Doc: "Margin is the outer-most transparent space around box element,\nwhich is _excluded_ from standard box rendering.", Directives: gti.Directives{}, Tag: "view:\"inline\""}},
		{"FillMargin", &gti.Field{Name: "FillMargin", Type: "bool", LocalType: "bool", Doc: "FillMargin determines is whether to fill the margin with\nthe surrounding background color before rendering the element itself.\nThis is typically necessary to prevent text, border, and box shadow from rendering\nover themselves. It should be kept at its default value of true\nin most circumstances, but it can be set to false when the element\nis fully managed by something that is guaranteed to render the\nappropriate background color for the element.", Directives: gti.Directives{}, Tag: ""}},
		{"Border", &gti.Field{Name: "Border", Type: "goki.dev/girl/styles.Border", LocalType: "Border", Doc: "Border is a line border around the box element", Directives: gti.Directives{}, Tag: ""}},
		{"MaxBorder", &gti.Field{Name: "MaxBorder", Type: "goki.dev/girl/styles.Border", LocalType: "Border", Doc: "MaxBorder is the largest border that will ever be rendered\naround the element, the size of which is used for computing\nthe effective margin to allocate for the element.", Directives: gti.Directives{}, Tag: ""}},
		{"BoxShadow", &gti.Field{Name: "BoxShadow", Type: "[]goki.dev/girl/styles.Shadow", LocalType: "[]Shadow", Doc: "BoxShadow is the box shadows to render around box (can have multiple)", Directives: gti.Directives{}, Tag: ""}},
		{"MaxBoxShadow", &gti.Field{Name: "MaxBoxShadow", Type: "[]goki.dev/girl/styles.Shadow", LocalType: "[]Shadow", Doc: "MaxBoxShadow contains the largest shadows that will ever be rendered\naround the element, the size of which are used for computing the\neffective margin to allocate for the element.", Directives: gti.Directives{}, Tag: ""}},
		{"MainAxis", &gti.Field{Name: "MainAxis", Type: "goki.dev/mat32/v2.Dims", LocalType: "mat32.Dims", Doc: "MainAxis is the main axis along which elements are arranged by a layout.\nX = horizontal axis (default), Y = vertical axis.\nSee also [Wrap]", Directives: gti.Directives{}, Tag: ""}},
		{"Wrap", &gti.Field{Name: "Wrap", Type: "bool", LocalType: "bool", Doc: "Wrap causes elements to wrap around in the CrossAxis dimension\nto fit within sizing constraints (on by default).", Directives: gti.Directives{}, Tag: ""}},
		{"Overflow", &gti.Field{Name: "Overflow", Type: "goki.dev/girl/styles.XY[goki.dev/girl/styles.Overflow]", LocalType: "XY[Overflow]", Doc: "Overflow determines how to handle overflowing content in a layout.\nDefault is OverflowVisible.  Set to OverflowAuto to enable scrollbars.", Directives: gti.Directives{}, Tag: ""}},
		{"Gap", &gti.Field{Name: "Gap", Type: "goki.dev/girl/units.XY", LocalType: "units.XY", Doc: "For layout, extra space added between elements in the layout.", Directives: gti.Directives{}, Tag: ""}},
		{"Columns", &gti.Field{Name: "Columns", Type: "int", LocalType: "int", Doc: "For layout, number of columns to use in a grid layout.\nIf > 0, number of rows is computed as N elements / Columns.\nUsed as a constraint in layout if individual elements\ndo not specify their row, column positions", Directives: gti.Directives{}, Tag: ""}},
		{"ScrollBarWidth", &gti.Field{Name: "ScrollBarWidth", Type: "goki.dev/girl/units.Value", LocalType: "units.Value", Doc: "width of a layout scrollbar", Directives: gti.Directives{}, Tag: ""}},
		{"Row", &gti.Field{Name: "Row", Type: "int", LocalType: "int", Doc: "prop: row = specifies the row that this element should appear within a grid layout", Directives: gti.Directives{}, Tag: ""}},
		{"Col", &gti.Field{Name: "Col", Type: "int", LocalType: "int", Doc: "prop: col = specifies the column that this element should appear within a grid layout", Directives: gti.Directives{}, Tag: ""}},
		{"RowSpan", &gti.Field{Name: "RowSpan", Type: "int", LocalType: "int", Doc: "specifies the number of sequential rows that this element should occupy\nwithin a grid layout (todo: not currently supported)", Directives: gti.Directives{}, Tag: ""}},
		{"ColSpan", &gti.Field{Name: "ColSpan", Type: "int", LocalType: "int", Doc: "specifies the number of sequential columns that this element should occupy\nwithin a grid layout", Directives: gti.Directives{}, Tag: ""}},
		{"Color", &gti.Field{Name: "Color", Type: "image/color.RGBA", LocalType: "color.RGBA", Doc: "prop: color (inherited) = text color -- also defines the currentColor variable value", Directives: gti.Directives{}, Tag: "inherit:\"true\""}},
		{"BackgroundColor", &gti.Field{Name: "BackgroundColor", Type: "goki.dev/colors.Full", LocalType: "colors.Full", Doc: "prop: background-color = background color -- not inherited, transparent by default", Directives: gti.Directives{}, Tag: ""}},
		{"Opacity", &gti.Field{Name: "Opacity", Type: "float32", LocalType: "float32", Doc: "prop: opacity = alpha value to apply to the foreground and background of this element and all of its children", Directives: gti.Directives{}, Tag: ""}},
		{"StateLayer", &gti.Field{Name: "StateLayer", Type: "float32", LocalType: "float32", Doc: "StateLayer, if above zero, indicates to create a state layer over the element with this much opacity (on a scale of 0-1) and the\ncolor Color (or StateColor if it defined). It is automatically set based on State, but can be overridden in stylers.", Directives: gti.Directives{}, Tag: ""}},
		{"StateColor", &gti.Field{Name: "StateColor", Type: "image/color.RGBA", LocalType: "color.RGBA", Doc: "StateColor, if not the zero color, is the color to use for the StateLayer instead of Color. If you want to disable state layers\nfor an element, do not use this; instead, set StateLayer to 0.", Directives: gti.Directives{}, Tag: ""}},
		{"Font", &gti.Field{Name: "Font", Type: "goki.dev/girl/styles.Font", LocalType: "Font", Doc: "font parameters -- no xml prefix -- also has color, background-color", Directives: gti.Directives{}, Tag: ""}},
		{"Text", &gti.Field{Name: "Text", Type: "goki.dev/girl/styles.Text", LocalType: "Text", Doc: "text parameters -- no xml prefix", Directives: gti.Directives{}, Tag: ""}},
		{"UnContext", &gti.Field{Name: "UnContext", Type: "goki.dev/girl/units.Context", LocalType: "units.Context", Doc: "units context -- parameters necessary for anchoring relative units", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/girl/styles.Text",
	ShortName: "styles.Text",
	IDName:    "text",
	Doc:       "Text is used for layout-level (widget, html-style) text styling --\nFontStyle contains all the lower-level text rendering info used in SVG --\nmost of these are inherited",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Align", &gti.Field{Name: "Align", Type: "goki.dev/girl/styles.Align", LocalType: "Align", Doc: "prop: text-align (inherited) = how to align text, horizontally.\nThis *only* applies to the text within its containing element,\nand is typically relevant only for multi-line text:\nfor single-line text, if element does not have a specified size\nthat is different from the text size, then this has *no effect*.", Directives: gti.Directives{}, Tag: "xml:\"text-align\" inherit:\"true\""}},
		{"AlignV", &gti.Field{Name: "AlignV", Type: "goki.dev/girl/styles.Align", LocalType: "Align", Doc: "prop: text-vertical-align (inherited) = vertical alignment of text.\nThis is only applicable for SVG styling, not regular CSS / GoGi,\nwhich uses the global Align.Y.  It *only* applies to the text within\nits containing element: if that element does not have a specified size\nthat is different from the text size, then this has *no effect*.", Directives: gti.Directives{}, Tag: "xml:\"text-vertical-align\" inherit:\"true\""}},
		{"Anchor", &gti.Field{Name: "Anchor", Type: "goki.dev/girl/styles.TextAnchors", LocalType: "TextAnchors", Doc: "prop: text-anchor (inherited) = for svg rendering only:\ndetermines the alignment relative to text position coordinate.\nFor RTL start is right, not left, and start is top for TB", Directives: gti.Directives{}, Tag: "xml:\"text-anchor\" inherit:\"true\""}},
		{"LetterSpacing", &gti.Field{Name: "LetterSpacing", Type: "goki.dev/girl/units.Value", LocalType: "units.Value", Doc: "prop: letter-spacing = spacing between characters and lines", Directives: gti.Directives{}, Tag: "xml:\"letter-spacing\""}},
		{"WordSpacing", &gti.Field{Name: "WordSpacing", Type: "goki.dev/girl/units.Value", LocalType: "units.Value", Doc: "prop: word-spacing (inherited) = extra space to add between words", Directives: gti.Directives{}, Tag: "xml:\"word-spacing\" inherit:\"true\""}},
		{"LineHeight", &gti.Field{Name: "LineHeight", Type: "goki.dev/girl/units.Value", LocalType: "units.Value", Doc: "prop: line-height (inherited) = specified height of a line of text; text is centered within the overall lineheight; the standard way to specify line height is in terms of em", Directives: gti.Directives{}, Tag: "xml:\"line-height\" inherit:\"true\""}},
		{"WhiteSpace", &gti.Field{Name: "WhiteSpace", Type: "goki.dev/girl/styles.WhiteSpaces", LocalType: "WhiteSpaces", Doc: "prop: white-space (*not* inherited) specifies how white space is processed,\nand how lines are wrapped.  If set to WhiteSpaceNormal (default) lines are wrapped.\nSee info about interactions with Grow.X setting for this and the NoWrap case.", Directives: gti.Directives{}, Tag: "xml:\"white-space\""}},
		{"UnicodeBidi", &gti.Field{Name: "UnicodeBidi", Type: "goki.dev/girl/styles.UnicodeBidi", LocalType: "UnicodeBidi", Doc: "prop: unicode-bidi (inherited) = determines how to treat unicode bidirectional information", Directives: gti.Directives{}, Tag: "xml:\"unicode-bidi\" inherit:\"true\""}},
		{"Direction", &gti.Field{Name: "Direction", Type: "goki.dev/girl/styles.TextDirections", LocalType: "TextDirections", Doc: "prop: direction (inherited) = direction of text -- only applicable for unicode-bidi = bidi-override or embed -- applies to all text elements", Directives: gti.Directives{}, Tag: "xml:\"direction\" inherit:\"true\""}},
		{"WritingMode", &gti.Field{Name: "WritingMode", Type: "goki.dev/girl/styles.TextDirections", LocalType: "TextDirections", Doc: "prop: writing-mode (inherited) = overall writing mode -- only for text elements, not span", Directives: gti.Directives{}, Tag: "xml:\"writing-mode\" inherit:\"true\""}},
		{"OrientationVert", &gti.Field{Name: "OrientationVert", Type: "float32", LocalType: "float32", Doc: "prop: glyph-orientation-vertical (inherited) = for TBRL writing mode (only), determines orientation of alphabetic characters -- 90 is default (rotated) -- 0 means keep upright", Directives: gti.Directives{}, Tag: "xml:\"glyph-orientation-vertical\" inherit:\"true\""}},
		{"OrientationHoriz", &gti.Field{Name: "OrientationHoriz", Type: "float32", LocalType: "float32", Doc: "prop: glyph-orientation-horizontal (inherited) = for horizontal LR/RL writing mode (only), determines orientation of all characters -- 0 is default (upright)", Directives: gti.Directives{}, Tag: "xml:\"glyph-orientation-horizontal\" inherit:\"true\""}},
		{"Indent", &gti.Field{Name: "Indent", Type: "goki.dev/girl/units.Value", LocalType: "units.Value", Doc: "prop: text-indent (inherited) = how much to indent the first line in a paragraph", Directives: gti.Directives{}, Tag: "xml:\"text-indent\" inherit:\"true\""}},
		{"ParaSpacing", &gti.Field{Name: "ParaSpacing", Type: "goki.dev/girl/units.Value", LocalType: "units.Value", Doc: "prop: para-spacing (inherited) = extra spacing between paragraphs -- copied from Style.Margin per CSS spec if that is non-zero, else can be set directly with para-spacing", Directives: gti.Directives{}, Tag: "xml:\"para-spacing\" inherit:\"true\""}},
		{"TabSize", &gti.Field{Name: "TabSize", Type: "int", LocalType: "int", Doc: "prop: tab-size (inherited) = tab size, in number of characters", Directives: gti.Directives{}, Tag: "xml:\"tab-size\" inherit:\"true\""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/girl/styles.XY",
	ShortName: "styles.XY",
	IDName:    "xy",
	Doc:       "XY represents X,Y values",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"X", &gti.Field{Name: "X", Type: "T", LocalType: "T", Doc: "X is the horizontal axis value", Directives: gti.Directives{}, Tag: ""}},
		{"Y", &gti.Field{Name: "Y", Type: "T", LocalType: "T", Doc: "Y is the vertical axis value", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})
