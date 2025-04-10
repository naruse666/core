// Code generated by 'yaegi extract github.com/naruse666/core/base/fsx'. DO NOT EDIT.

package basesymbols

import (
	"github.com/naruse666/core/base/fsx"
	"reflect"
)

func init() {
	Symbols["github.com/naruse666/core/base/fsx/fsx"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"CopyFile":         reflect.ValueOf(fsx.CopyFile),
		"DirAndFile":       reflect.ValueOf(fsx.DirAndFile),
		"DirFS":            reflect.ValueOf(fsx.DirFS),
		"Dirs":             reflect.ValueOf(fsx.Dirs),
		"ExtSplit":         reflect.ValueOf(fsx.ExtSplit),
		"FileExists":       reflect.ValueOf(fsx.FileExists),
		"FileExistsFS":     reflect.ValueOf(fsx.FileExistsFS),
		"Filenames":        reflect.ValueOf(fsx.Filenames),
		"Files":            reflect.ValueOf(fsx.Files),
		"FindFilesOnPaths": reflect.ValueOf(fsx.FindFilesOnPaths),
		"GoSrcDir":         reflect.ValueOf(fsx.GoSrcDir),
		"HasFile":          reflect.ValueOf(fsx.HasFile),
		"LatestMod":        reflect.ValueOf(fsx.LatestMod),
		"RelativeFilePath": reflect.ValueOf(fsx.RelativeFilePath),
		"SplitRootPathFS":  reflect.ValueOf(fsx.SplitRootPathFS),
		"Sub":              reflect.ValueOf(fsx.Sub),

		// type definitions
		"Filename": reflect.ValueOf((*fsx.Filename)(nil)),
	}
}
