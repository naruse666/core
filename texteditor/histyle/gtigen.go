// Code generated by "goki generate -add-types"; DO NOT EDIT.

package histyle

import (
	"cogentcore.org/core/gti"
)

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/texteditor/histyle.Trilean", IDName: "trilean", Doc: "Trilean value for StyleEntry value inheritance.", Directives: []gti.Directive{{Tool: "enums", Directive: "enum"}}})

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/texteditor/histyle.StyleEntry", IDName: "style-entry", Doc: "StyleEntry is one value in the map of highlight style values", Fields: []gti.Field{{Name: "Color", Doc: "text color"}, {Name: "Background", Doc: "background color"}, {Name: "Border", Doc: "border color? not sure what this is -- not really used"}, {Name: "Bold", Doc: "bold font"}, {Name: "Italic", Doc: "italic font"}, {Name: "Underline", Doc: "underline"}, {Name: "NoInherit", Doc: "don't inherit these settings from sub-category or category levels -- otherwise everything with a Pass is inherited"}}})

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/texteditor/histyle.Style", IDName: "style", Doc: "Style is a full style map of styles for different token.Tokens tag values"})

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/texteditor/histyle.Styles", IDName: "styles", Doc: "Styles is a collection of styles", Methods: []gti.Method{{Name: "OpenJSON", Doc: "Open hi styles from a JSON-formatted file. You can save and open\nstyles to / from files to share, experiment, transfer, etc.", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Args: []string{"filename"}, Returns: []string{"error"}}, {Name: "SaveJSON", Doc: "Save hi styles to a JSON-formatted file. You can save and open\nstyles to / from files to share, experiment, transfer, etc.", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Args: []string{"filename"}, Returns: []string{"error"}}}})

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/texteditor/histyle.Value", IDName: "value", Doc: "Value presents a button for selecting a highlight styling method", Embeds: []gti.Field{{Name: "ValueBase"}}})
