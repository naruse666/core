// Code generated by "core generate"; DO NOT EDIT.

package cmd

import (
	"cogentcore.org/core/gti"
)

var _ = gti.AddFunc(&gti.Func{Name: "cogentcore.org/core/core/cmd.Build", Doc: "Build builds an executable for the package\nat the config path for the config platforms.", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Args: []string{"c"}, Returns: []string{"error"}})

var _ = gti.AddFunc(&gti.Func{Name: "cogentcore.org/core/core/cmd.Changed", Doc: "Changed concurrently prints all of the repositories within this directory\nthat have been changed and need to be updated in Git.", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Args: []string{"c"}, Returns: []string{"error"}})

var _ = gti.AddFunc(&gti.Func{Name: "cogentcore.org/core/core/cmd.Install", Doc: "Install installs the package on the local system.\nIt uses the same config info as build.", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Args: []string{"c"}, Returns: []string{"error"}})

var _ = gti.AddFunc(&gti.Func{Name: "cogentcore.org/core/core/cmd.Log", Doc: "Log prints the logs from your app running on Android to the terminal.\nAndroid is the only supported platform for log; use the -debug flag on\nrun for other platforms.", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Args: []string{"c"}, Returns: []string{"error"}})

var _ = gti.AddFunc(&gti.Func{Name: "cogentcore.org/core/core/cmd.Pack", Doc: "Pack builds and packages the app for the target platform.\nFor android, ios, and web, it is equivalent to build.", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Args: []string{"c"}, Returns: []string{"error"}})

var _ = gti.AddFunc(&gti.Func{Name: "cogentcore.org/core/core/cmd.Pull", Doc: "Pull concurrently pulls all of the Git repositories within the current directory.", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Args: []string{"c"}, Returns: []string{"error"}})

var _ = gti.AddFunc(&gti.Func{Name: "cogentcore.org/core/core/cmd.Release", Doc: "Release releases the project with the specified git version tag.", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Args: []string{"c"}, Returns: []string{"error"}})

var _ = gti.AddFunc(&gti.Func{Name: "cogentcore.org/core/core/cmd.NextRelease", Doc: "NextRelease releases the project with the current git version\ntag incremented by one patch version.", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Args: []string{"c"}, Returns: []string{"error"}})

var _ = gti.AddFunc(&gti.Func{Name: "cogentcore.org/core/core/cmd.Run", Doc: "Run builds and runs the config package. It also displays the logs generated\nby the app. It uses the same config info as build.", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Args: []string{"c"}, Returns: []string{"error"}})

var _ = gti.AddFunc(&gti.Func{Name: "cogentcore.org/core/core/cmd.Setup", Doc: "Setup installs platform-specific dependencies for the current platform.\nIt only needs to be called once per system.", Directives: []gti.Directive{{Tool: "gti", Directive: "add"}}, Args: []string{"c"}, Returns: []string{"error"}})
