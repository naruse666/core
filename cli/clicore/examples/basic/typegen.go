// Code generated by "core generate -add-types -add-methods"; DO NOT EDIT.

package main

import (
	"github.com/naruse666/core/types"
)

var _ = types.AddType(&types.Type{Name: "main.Config", IDName: "config", Directives: []types.Directive{{Tool: "go", Directive: "generate", Args: []string{"core", "generate", "-add-types", "-add-methods"}}}, Fields: []types.Field{{Name: "Name", Doc: "the name of the user"}, {Name: "Age", Doc: "the age of the user"}, {Name: "LikesGo", Doc: "whether the user likes Go"}, {Name: "BuildTarget", Doc: "the target platform to build for"}}})
