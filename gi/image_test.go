// Copyright (c) 2023, The Goki Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gi

import (
	"image"
	"path/filepath"
	"testing"

	"cogentcore.org/core/grr"
	"cogentcore.org/core/mat32"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/units"
)

var testImagePath = Filename(filepath.Join("..", "logo", "goki_logo.png"))

func TestImageBasic(t *testing.T) {
	sc := NewScene()
	fr := NewFrame(sc)
	img := NewImage(fr)
	grr.Test(t, img.OpenImage(testImagePath))
	sc.AssertRender(t, filepath.Join("image", "basic"))
}

func TestImageCropped(t *testing.T) {
	sc := NewScene()
	sc.Style(func(s *styles.Style) {
		s.Max.Set(units.Dp(75))
	})
	fr := NewFrame(sc).Style(func(s *styles.Style) {
		s.Overflow.Set(styles.OverflowAuto)
	})
	img := NewImage(fr)
	grr.Test(t, img.OpenImage(testImagePath))
	sc.AssertRender(t, filepath.Join("image", "cropped"))
}

func TestImageScrolled(t *testing.T) {
	sc := NewScene()
	sc.Style(func(s *styles.Style) {
		s.Max.Set(units.Dp(75))
	})
	fr := NewFrame(sc).Style(func(s *styles.Style) {
		s.Overflow.Set(styles.OverflowAuto)
	})
	img := NewImage(fr)
	grr.Test(t, img.OpenImage(testImagePath))
	sc.AssertRender(t, filepath.Join("image", "scrolled"), func() {
		sc.GoosiEventMgr().Scroll(image.Pt(10, 10), mat32.V2(2, 3))
	})
}
