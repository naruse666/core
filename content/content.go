// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package content provides a system for making content-focused
// apps and websites consisting of Markdown, HTML, and Cogent Core.
package content

//go:generate core generate

import (
	"io/fs"
	"strings"

	"cogentcore.org/core/base/errors"
	"cogentcore.org/core/base/fsx"
	"cogentcore.org/core/core"
	"cogentcore.org/core/htmlcore"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/tree"
)

// Content manages and displays the content of a set of pages.
type Content struct {
	core.Frame

	// Source is the source filesystem for the content.
	// It should be set using [Content.SetSource] or [Content.SetContent].
	Source fs.FS `set:"-"`

	// Context is the [htmlcore.Context] used to render the content,
	// which can be modified for things such as adding wikilink handlers.
	Context *htmlcore.Context `set:"-"`

	// pages are the pages that constitute the content.
	pages []*Page

	// pagesByName has the [Page] for each name transformed into lowercase.
	pagesByName map[string]*Page

	// pagesByURL has the [Page] for each URL.
	pagesByURL map[string]*Page

	// currentPage is the currently open page.
	currentPage *Page

	// renderedPage is the most recently rendered page.
	renderedPage *Page

	// headings are all of the heading elements on the currently rendered page.
	headings []*core.Text
}

func (ct *Content) Init() {
	ct.Frame.Init()
	ct.Context = htmlcore.NewContext()
	ct.Context.OpenURL = func(url string) {
		ct.Open(url)
	}
	ct.Context.AddWikilinkHandler(func(text string) (url string, label string) {
		name, label, has := strings.Cut(text, "|")
		if !has {
			label = text
		}
		if pg, ok := ct.pagesByName[strings.ToLower(name)]; ok {
			return pg.URL, label
		}
		return "", ""
	})

	ct.Styler(func(s *styles.Style) {
		s.Direction = styles.Column
		s.Grow.Set(1, 1)
	})

	ct.Maker(func(p *tree.Plan) {
		if ct.currentPage == nil {
			return
		}
		if ct.currentPage.Name != "" {
			tree.Add(p, func(w *core.Text) {
				w.SetType(core.TextDisplaySmall)
				w.Updater(func() {
					w.SetText(ct.currentPage.Name)
				})
			})
		}
		tree.Add(p, func(w *core.Frame) {
			w.Styler(func(s *styles.Style) {
				s.Direction = styles.Column
				s.Grow.Set(1, 1)
			})
			w.Updater(func() {
				errors.Log(ct.loadPage(w))
			})
		})
	})
}

// SetSource sets the source filesystem for the content.
func (ct *Content) SetSource(source fs.FS) *Content {
	ct.Source = source
	ct.pages = []*Page{}
	ct.pagesByName = map[string]*Page{}
	ct.pagesByURL = map[string]*Page{}
	errors.Log(fs.WalkDir(ct.Source, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path == "" || path == "." {
			return nil
		}
		pg, err := NewPage(ct.Source, path)
		if err != nil {
			return err
		}
		ct.pages = append(ct.pages, pg)
		ct.pagesByName[strings.ToLower(pg.Name)] = pg
		ct.pagesByURL[pg.URL] = pg
		return nil
	}))
	if root, ok := ct.pagesByURL[""]; ok {
		ct.currentPage = root
	} else {
		ct.currentPage = ct.pages[0]
	}
	return ct
}

// SetContent is a helper function that calls [Content.SetSource]
// with the "content" subdirectory of the given filesystem.
func (ct *Content) SetContent(content fs.FS) *Content {
	return ct.SetSource(fsx.Sub(content, "content"))
}

// Open opens the page with the given URL and updates the display.
// If no pages correspond to the URL, it is opened in the default browser.
func (ct *Content) Open(url string) *Content {
	pg, ok := ct.pagesByURL[url]
	if !ok {
		core.TheApp.OpenURL(url)
		return ct
	}
	ct.currentPage = pg
	ct.Update()
	return ct
}

// loadPage loads the current page content into the given frame if it is not already loaded.
func (ct *Content) loadPage(w *core.Frame) error {
	if ct.renderedPage == ct.currentPage {
		return nil
	}
	w.DeleteChildren()
	b, err := ct.currentPage.ReadContent()
	if err != nil {
		return err
	}
	err = htmlcore.ReadMD(ct.Context, w, b)
	if err != nil {
		return err
	}
	ct.getHeadings(w)
	ct.renderedPage = ct.currentPage
	return nil
}

// getHeadings sets [Content.headings] to be all of the headings in the given frame.
func (ct *Content) getHeadings(w *core.Frame) {
	ct.headings = []*core.Text{}
	w.WidgetWalkDown(func(cw core.Widget, cwb *core.WidgetBase) bool {
		tx, ok := cw.(*core.Text)
		if !ok {
			return tree.Continue
		}
		switch tx.Property("tag") {
		case "h1", "h2", "h3", "h4", "h5", "h6":
			ct.headings = append(ct.headings, tx)
		}
		return tree.Continue
	})
}
