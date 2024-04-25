// Copyright 2024 Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command docs provides documentation of Cogent Core,
// hosted at https://cogentcore.org/core.
package main

//go:generate core generate -pages content

import (
	"embed"
	"io/fs"

	"cogentcore.org/core/colors"
	"cogentcore.org/core/core"
	"cogentcore.org/core/errors"
	"cogentcore.org/core/events"
	"cogentcore.org/core/htmlview"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/pages"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/texteditor"
	"cogentcore.org/core/units"
	"cogentcore.org/core/views"
)

//go:embed content
var content embed.FS

//go:embed icon.svg
var icon []byte

//go:embed name.png
var name embed.FS

//go:embed image.png
var myImage embed.FS

//go:embed icon.svg
var mySVG embed.FS

//go:embed file.go
var myFile embed.FS

func main() {
	b := core.NewBody("Cogent Core Docs")
	pg := pages.NewPage(b).SetSource(errors.Log1(fs.Sub(content, "content")))
	pg.Context.WikilinkResolver = htmlview.PkgGoDevWikilink("cogentcore.org/core")
	b.AddAppBar(pg.AppBar)

	htmlview.ElementHandlers["home-header"] = homeHeader
	htmlview.ElementHandlers["get-started"] = func(ctx *htmlview.Context) bool {
		core.NewButton(ctx.BlockParent).SetText("Get Started").OnClick(func(e events.Event) {
			pg.OpenURL("/getting-started", true)
		}).Style(func(s *styles.Style) {
			s.Align.Self = styles.Center
		})
		return true
	}

	b.RunMainWindow()
}

func homeHeader(ctx *htmlview.Context) bool {
	ly := core.NewLayout(ctx.BlockParent).Style(func(s *styles.Style) {
		s.Direction = styles.Column
		s.CenterAll()
	})
	errors.Log(core.NewSVG(ly).ReadBytes(icon))
	img := core.NewImage(ly)
	errors.Log(img.OpenFS(name, "name.png"))
	img.Style(func(s *styles.Style) {
		x := func(uc *units.Context) float32 {
			return min(uc.Dp(612), uc.Vw(90))
		}
		s.Min.Set(units.Custom(x), units.Custom(func(uc *units.Context) float32 {
			return x(uc) * (128.0 / 612.0)
		}))
	})
	core.NewText(ly).SetType(core.TextHeadlineMedium).SetText("A cross-platform framework for building powerful, fast, and cogent 2D and 3D apps")

	blocks := core.NewLayout(ly).Style(func(s *styles.Style) {
		s.Display = styles.Grid
		s.Columns = 2
	})

	homeTextBlock(blocks, "CODE ONCE, RUN EVERYWHERE", "With Cogent Core, you can write your app once and it will instantly run on macOS, Windows, Linux, iOS, Android, and the Web, automatically scaling to any screen. Instead of struggling with platform-specific code in a multitude of languages, you can easily write and maintain a single pure Go codebase.")
	core.NewIcon(blocks).SetIcon(icons.Devices).Style(func(s *styles.Style) {
		s.Min.Set(units.Pw(40))
	})

	// we get the code example contained within the md file
	ctx.Node = ctx.Node.FirstChild.NextSibling
	ctx.BlockParent = core.NewLayout(blocks).Style(func(s *styles.Style) {
		s.Direction = styles.Column
	})
	htmlview.HandleElement(ctx)

	homeTextBlock(blocks, "EFFORTLESS ELEGANCE", "Cogent Core is built on Go, a high-level language designed for building elegant, readable, and scalable code with full type safety and a robust design that never gets in your way. Cogent Core makes it easy to get started with cross-platform app development in just two commands and seven lines of simple code.")

	homeTextBlock(blocks, "COMPLETELY CUSTOMIZABLE", "Cogent Core allows developers and users to fully customize apps to fit their unique needs and preferences through a robust styling system and a powerful color system that allow developers and users to instantly customize every aspect of the appearance and behavior of an app.")
	views.NewStructView(blocks).SetStruct(core.AppearanceSettings).OnChange(func(e events.Event) {
		core.UpdateSettings(blocks, core.AppearanceSettings)
	})

	texteditor.NewSoloEditor(blocks).Buffer.SetLang("go").SetTextString(`package main

func main() {
    fmt.Println("Hello, world!")
}
`)

	homeTextBlock(blocks, "ENDLESS FEATURES", "Cogent Core comes with a powerful set of advanced features that allow you to make almost anything, including fully featured text editors, video and audio players, interactive 3D graphics, customizable data plots, Markdown and HTML rendering, SVG and canvas vector graphics, and automatic views of any Go data structure for instant data binding and advanced app inspection.")

	homeTextBlock(blocks, "OPTIMIZED EXPERIENCE", "Every part of your development experience is guided by a comprehensive set of interactive example-based documentation, in-depth video tutorials, easy-to-use command line tools specialized for Cogent Core, and active support and development from the Cogent Core developers.")
	core.NewIcon(blocks).SetIcon(icons.PlayCircle).Style(func(s *styles.Style) {
		s.Min.Set(units.Pw(40))
	})

	core.NewIcon(blocks).SetIcon(icons.Bolt).Style(func(s *styles.Style) {
		s.Min.Set(units.Pw(40))
	})
	homeTextBlock(blocks, "EXTREMELY FAST", "Cogent Core is powered by Vulkan, a modern, cross-platform, high-performance graphics framework that allows apps to run on all platforms at extremely fast speeds. All Cogent Core apps compile to machine code, allowing them to run without any overhead.")

	homeTextBlock(blocks, "FREE AND OPEN SOURCE", "Cogent Core is completely free and open source under the permissive BSD-3 License, allowing you to use Cogent Core for any purpose, commercially or personally. We believe that software works best when everyone can use it.")
	core.NewIcon(blocks).SetIcon(icons.Code).Style(func(s *styles.Style) {
		s.Min.Set(units.Pw(40))
	})

	core.NewIcon(blocks).SetIcon(icons.GlobeAsia).Style(func(s *styles.Style) {
		s.Min.Set(units.Pw(40))
	})
	homeTextBlock(blocks, "USED AROUND THE WORLD", "Over six years of development, Cogent Core has been used and thoroughly tested by developers and scientists around the world for a wide variety of use cases. Cogent Core is a production-ready framework actively used to power everything from end-user apps to scientific research.")

	core.NewText(ly).SetType(core.TextDisplaySmall).SetText("<b>What can you make with Cogent Core?</b>")

	appBlocks := core.NewLayout(ly).Style(func(s *styles.Style) {
		s.Display = styles.Grid
		s.Columns = 2
	})

	homeTextBlock(appBlocks, "COGENT CODE", "Cogent Code is a fully featured Go IDE with support for syntax highlighting, code completion, symbol lookup, building and debugging, version control, keyboard shortcuts, and many other features.")
	core.NewIcon(appBlocks).SetIcon(icons.Code).Style(func(s *styles.Style) {
		s.Min.Set(units.Pw(40))
	})

	core.NewIcon(appBlocks).SetIcon(icons.Polyline).Style(func(s *styles.Style) {
		s.Min.Set(units.Pw(40))
	})
	homeTextBlock(appBlocks, "COGENT VECTOR", "Cogent Vector is a powerful vector graphics editor with complete support for shapes, paths, curves, text, images, gradients, groups, alignment, styling, importing, exporting, undo, redo, and various other features.")

	homeTextBlock(appBlocks, "COGENT MAIL", "Cogent Mail is a customizable email client with built-in Markdown support and an extensive set of keyboard shortcuts for advanced mail filing.")
	core.NewIcon(appBlocks).SetIcon(icons.Mail).Style(func(s *styles.Style) {
		s.Min.Set(units.Pw(40))
	})

	core.NewIcon(appBlocks).SetIcon(icons.Cognition).Style(func(s *styles.Style) {
		s.Min.Set(units.Pw(40))
	})
	homeTextBlock(appBlocks, "EMERGENT", "Emergent is a collection of biologically based 3D neural network models of the brain that power ongoing research in computational cognitive neuroscience.")

	return true
}

func homeTextBlock(parent core.Widget, title, text string) {
	block := core.NewLayout(parent).Style(func(s *styles.Style) {
		s.Direction = styles.Column
		s.Text.Align = styles.Start
		s.Min.X.Pw(50)
		s.Grow.Set(0, 0)
	})
	core.NewText(block).SetType(core.TextHeadlineLarge).SetText(title).Style(func(s *styles.Style) {
		s.Font.Weight = styles.WeightBold
		s.Color = colors.C(colors.Scheme.Primary.Base)
	})
	core.NewText(block).SetType(core.TextTitleLarge).SetText(text)
}
