// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embed"

	"github.com/naruse666/core/content"
	"github.com/naruse666/core/core"
	"github.com/naruse666/core/htmlcore"
	_ "github.com/naruse666/core/yaegicore"
)

//go:embed content
var econtent embed.FS

func main() {
	b := core.NewBody("Cogent Content Example")
	ct := content.NewContent(b).SetContent(econtent)
	ct.Context.AddWikilinkHandler(htmlcore.GoDocWikilink("doc", "github.com/naruse666/core"))
	b.AddTopBar(func(bar *core.Frame) {
		core.NewToolbar(bar).Maker(ct.MakeToolbar)
	})
	b.RunMainWindow()
}
