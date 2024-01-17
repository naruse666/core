// Code generated by "goki generate"; DO NOT EDIT.

package base

import (
	"cogentcore.org/core/gti"
)

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/goosi/driver/base.App", IDName: "app", Doc: "App contains the data and logic common to all implementations of [goosi.App].", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Fields: []gti.Field{{Name: "This", Doc: "This is the App as a [goosi.App] interface, which preserves the actual identity\nof the app when calling interface methods in the base App."}, {Name: "Mu", Doc: "Mu is the main mutex protecting access to app operations, including [App.RunOnMain] functions."}, {Name: "MainQueue", Doc: "MainQueue is the queue of functions to call on the main loop.\nTo add to it, use [App.RunOnMain]."}, {Name: "MainDone", Doc: "MainDone is a channel on which is a signal is sent when the main\nloop of the app should be terminated."}, {Name: "Nm", Doc: "Nm is the name of the app."}, {Name: "Abt", Doc: "Abt is the about information for the app."}, {Name: "OpenFls", Doc: "OpenFls are files that have been set by the operating system to open at startup."}, {Name: "Quitting", Doc: "Quitting is whether the app is quitting and thus closing all of the windows"}, {Name: "QuitReqFunc", Doc: "QuitReqFunc is a function to call when a quit is requested"}, {Name: "QuitCleanFunc", Doc: "QuitCleanFunc is a function to call when the app is about to quit"}, {Name: "Dark", Doc: "Dark is whether the system color theme is dark (as opposed to light)"}}})

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/goosi/driver/base.AppMulti", IDName: "app-multi", Doc: "AppMulti contains the data and logic common to all implementations of [goosi.App]\non multi-window platforms (desktop), as opposed to single-window\nplatforms (mobile, web, and offscreen), for which you should use [AppSingle]. An AppMulti is associated\nwith a corresponding type of [goosi.Window]. The [goosi.Window]\ntype should embed [WindowMulti].", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Embeds: []gti.Field{{Name: "App"}}, Fields: []gti.Field{{Name: "Windows", Doc: "Windows are the windows associated with the app"}, {Name: "Screens", Doc: "Screens are the screens associated with the app"}, {Name: "AllScreens", Doc: "AllScreens is a unique list of all screens ever seen, from which\ninformation can be got if something is missing in [AppMulti.Screens]"}, {Name: "CtxWindow", Doc: "CtxWindow is a dynamically set context window used for some operations"}}})

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/goosi/driver/base.AppSingle", IDName: "app-single", Doc: "AppSingle contains the data and logic common to all implementations of [goosi.App]\non single-window platforms (mobile, web, and offscreen), as opposed to multi-window\nplatforms (desktop), for which you should use [AppMulti]. An AppSingle is associated\nwith a corresponding type of [goosi.Drawer] and [goosi.Window]. The [goosi.Window]\ntype should embed [WindowSingle].", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Embeds: []gti.Field{{Name: "App"}}, Fields: []gti.Field{{Name: "EvMgr", Doc: "EvMgr is the event manager for the app"}, {Name: "Draw", Doc: "Draw is the single [goosi.Drawer] used for the app."}, {Name: "Win", Doc: "Win is the single [goosi.Window] associated with the app."}, {Name: "Scrn", Doc: "Scrn is the single [goosi.Screen] associated with the app."}, {Name: "Insets", Doc: "Insets are the size of any insets on the sides of the screen."}}})

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/goosi/driver/base.Window", IDName: "window", Doc: "Window contains the data and logic common to all implementations of [goosi.Window].\nA Window is associated with a corresponding [goosi.App] type.", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Fields: []gti.Field{{Name: "This", Doc: "This is the Window as a [goosi.Window] interface, which preserves the actual identity\nof the window when calling interface methods in the base Window."}, {Name: "App", Doc: "App is the [goosi.App] associated with the window."}, {Name: "Mu", Doc: "Mu is the main mutex protecting access to window operations, including [Window.RunOnWin] functions."}, {Name: "WinClose", Doc: "WinClose is a channel on which a single is sent to indicate that the\nwindow should close."}, {Name: "CloseReqFunc", Doc: "CloseReqFunc is the function to call on a close request"}, {Name: "CloseCleanFunc", Doc: "CloseCleanFunc is the function to call to close the window"}, {Name: "Nm", Doc: "Nm is the name of the window"}, {Name: "Titl", Doc: "Titl is the title of the window"}, {Name: "Flgs", Doc: "Flgs contains the flags associated with the window"}, {Name: "FPS", Doc: "FPS is the FPS (frames per second) for rendering the window"}, {Name: "DestroyGPUFunc", Doc: "DestroyGPUFunc should be set to a function that will destroy GPU resources\nin the main thread prior to destroying the drawer\nand the surface; otherwise it is difficult to\nensure that the proper ordering of destruction applies."}, {Name: "CursorEnabled", Doc: "CursorEnabled is whether the cursor is currently enabled"}}})

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/goosi/driver/base.WindowMulti", IDName: "window-multi", Doc: "WindowMulti contains the data and logic common to all implementations of [goosi.Window]\non multi-window platforms (desktop), as opposed to single-window\nplatforms (mobile, web, and offscreen), for which you should use [WindowSingle].\nA WindowMulti is associated with a corresponding [goosi.App] type.\nThe [goosi.App] type should embed [AppMulti].", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Embeds: []gti.Field{{Name: "Window"}}, Fields: []gti.Field{{Name: "EvMgr", Doc: "EvMgr is the event manager for the window"}, {Name: "Draw", Doc: "Draw is the [goosi.Drawer] used for this window."}, {Name: "Pos", Doc: "Pos is the position of the window"}, {Name: "WnSize", Doc: "WnSize is the size of the window in window manager coordinates"}, {Name: "PixSize", Doc: "PixSize is the pixel size of the window in raw display dots"}, {Name: "DevicePixelRatio", Doc: "DevicePixelRatio is a factor that scales the screen's\n\"natural\" pixel coordinates into actual device pixels.\nOn OS-X, it is backingScaleFactor = 2.0 on \"retina\""}, {Name: "PhysDPI", Doc: "PhysicalDPI is the physical dots per inch of the screen,\nfor generating true-to-physical-size output.\nIt is computed as 25.4 * (PixSize.X / PhysicalSize.X)\nwhere 25.4 is the number of mm per inch."}, {Name: "LogDPI", Doc: "LogicalDPI is the logical dots per inch of the screen,\nwhich is used for all rendering.\nIt is: transient zoom factor * screen-specific multiplier * PhysicalDPI"}}})

var _ = gti.AddType(&gti.Type{Name: "cogentcore.org/core/goosi/driver/base.WindowSingle", IDName: "window-single", Doc: "WindowSingle contains the data and logic common to all implementations of [goosi.Window]\non single-window platforms (mobile, web, and offscreen), as opposed to multi-window\nplatforms (desktop), for which you should use [WindowSingle].\nA WindowSingle is associated with a corresponding [AppSingler] type.", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Embeds: []gti.Field{{Name: "Window"}}})
