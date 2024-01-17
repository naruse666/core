// Code generated by "goki generate"; DO NOT EDIT.

package offscreen

import (
	"cogentcore.org/core/gti"
)

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/goosi/driver/offscreen.App", IDName: "app", Doc: "App is the [goosi.App] implementation for the offscreen platform", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Embeds: []gti.Field{{Name: "AppSingle"}}, Fields: []gti.Field{{Name: "TempDataDir", Doc: "TempDataDir is the path of the app data directory, used as the\nreturn value of [App.DataDir]. It is set to a temporary directory,\nas offscreen tests should not be dependent on user preferences and\nother data."}}})

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/goosi/driver/offscreen.Window", IDName: "window", Doc: "Window is the implementation of [goosi.Window] for the offscreen platform.", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Embeds: []gti.Field{{Name: "WindowSingle"}}})
