// Copyright (c) 2018, The Goki Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gi

import (
	"image"
	"log/slog"

	"cogentcore.org/core/colors"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/svg"
	"golang.org/x/image/draw"
)

// Icon contains a svg.SVG element.
// The rendered version is cached for a given size.
// Icons do not render a background or border independent of their SVG object.
type Icon struct {
	WidgetBase

	// icon name that has been set.
	Icon icons.Icon `set:"-"`

	// file name for the loaded icon, if loaded
	Filename string `set:"-"`

	// SVG drawing of the icon
	SVG svg.SVG `set:"-"`
}

func (ic *Icon) CopyFieldsFrom(frm any) {
	fr := frm.(*Icon)
	ic.WidgetBase.CopyFieldsFrom(&fr.WidgetBase)
	ic.Icon = fr.Icon
	ic.Filename = fr.Filename
}

func (ic *Icon) OnInit() {
	ic.WidgetBase.OnInit()
	ic.SetStyles()
}

func (ic *Icon) SetStyles() {
	ic.SVG.Norm = true
	ic.SVG.Scale = 1
	ic.Style(func(s *styles.Style) {
		s.Min.X.Dp(16)
		s.Min.Y.Dp(16)
	})
}

// SetIcon sets the icon, logging error if not found.
// Does nothing if IconName is already == icon name.
func (ic *Icon) SetIcon(icon icons.Icon) *Icon {
	_, err := ic.SetIconTry(icon)
	if err != nil {
		slog.Error("error opening icon named", "name", icon, "err", err)
	}
	return ic
}

// SetIconUpdate sets the icon and sets flag to render.
// Does nothing if IconName is already == icon name.
func (ic *Icon) SetIconUpdate(icon icons.Icon) *Icon {
	ic.SetIcon(icon)
	ic.SetNeedsRender(true)
	return ic
}

// SetIconTry sets the icon, returning error
// message if not found etc, and returning true if a new icon was actually set.
// Does nothing and returns false if IconName is already == icon name.
func (ic *Icon) SetIconTry(icon icons.Icon) (bool, error) {
	if !icon.IsValid() {
		ic.SVG.DeleteAll()
		ic.Config()
		return false, nil
	}
	if ic.SVG.Root.HasChildren() && ic.Icon == icon {
		// fmt.Println("icon already set:", icon)
		return false, nil
	}
	fnm := icon.Filename()
	ic.SVG.Config(2, 2)
	err := ic.SVG.OpenFS(icons.Icons, fnm)
	if err != nil {
		ic.Config()
		return false, err
	}
	ic.Icon = icon
	// fmt.Println("icon set:", icon)
	return true, nil

}

func (ic *Icon) DrawIntoScene() {
	if ic.SVG.Pixels == nil {
		return
	}
	r := ic.Geom.ContentBBox
	sp := ic.Geom.ScrollOffset()
	draw.Draw(ic.Sc.Pixels, r, ic.SVG.Pixels, sp, draw.Over)
}

// RenderSVG renders the SVG to Pixels if needs update
func (ic *Icon) RenderSVG() {
	rc := ic.Sc.RenderCtx()
	sv := &ic.SVG
	sz := ic.Geom.Size.Actual.Content.ToPoint()
	clr := colors.ApplyOpacity(ic.Styles.Color, ic.Styles.Opacity)
	if !rc.HasFlag(RenderRebuild) && sv.Pixels != nil { // if rebuilding rebuild..
		isz := sv.Pixels.Bounds().Size()
		// if nothing has changed, we don't need to re-render
		if isz == sz && sv.Name == string(ic.Icon) && colors.ToUniform(sv.Color) == clr {
			return
		}
	}
	// todo: units context from us to SVG??

	if sz == (image.Point{}) {
		return
	}
	// ensure that we have new pixels to render to in order to prevent
	// us from rendering over ourself
	sv.Pixels = image.NewRGBA(image.Rectangle{Max: sz})
	sv.RenderState.Init(sz.X, sz.Y, sv.Pixels)
	sv.Geom.Size = sz // make sure

	sv.Resize(sz) // does Config if needed

	// TODO(kai): what about gradient icons?
	sv.Color = colors.C(clr)

	sv.Scale = 1
	sv.Render()
	sv.Name = string(ic.Icon)
	// fmt.Println("re-rendered icon:", sv.Name, "size:", sz)
}

func (ic *Icon) Render() {
	ic.RenderSVG()
	if ic.PushBounds() {
		ic.RenderChildren()
		ic.DrawIntoScene()
		ic.PopBounds()
	}
}
