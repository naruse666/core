// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package system

import (
	"image"
	"image/draw"

	"cogentcore.org/core/math32"
)

// Drawer is an interface for image/draw style image compositing
// functionality, which is implemented for the GPU in
// [*cogentcore.org/core/gpu/gpudraw.Drawer] and in offscreen drivers.
// This is used for compositing the stack of images that together comprise
// the content of a window.
type Drawer interface {
	// DestBounds returns the bounds of the render destination
	DestBounds() image.Rectangle

	// Start starts recording a sequence of draw / fill actions,
	// which will be performed on the GPU at End().
	// This must be called prior to any Drawer operations.
	Start()

	// End ends image drawing rendering process on render target.
	End()

	// Copy copies the given Go source image to the render target, with the
	// same semantics as golang.org/x/image/draw.Copy, with the destination
	// implicit in the Drawer target.
	//   - Must have called Start first!
	//   - dp is the destination point.
	//   - src is the source image. If an image.Uniform, fast Fill is done.
	//   - sr is the source region, if zero full src is used; must have for Uniform.
	//   - op is the drawing operation: Src = copy source directly (blit),
	//     Over = alpha blend with existing.
	Copy(dp image.Point, src image.Image, sr image.Rectangle, op draw.Op)

	// Scale copies the given Go source image to the render target,
	// scaling the region defined by src and sr to the destination
	// such that sr in src-space is mapped to dr in dst-space.
	// with the same semantics as golang.org/x/image/draw.Scale, with the
	// destination implicit in the Drawer target.
	// If src image is an
	//   - Must have called Start first!
	//   - dr is the destination rectangle; if zero uses full dest image.
	//   - src is the source image. Uniform does not work (or make sense) here.
	//   - sr is the source region, if zero full src is used; must have for Uniform.
	//   - op is the drawing operation: Src = copy source directly (blit),
	//     Over = alpha blend with existing.
	Scale(dr image.Rectangle, src image.Image, sr image.Rectangle, op draw.Op)

	// Transform copies the given Go source image to the render target,
	// with the same semantics as golang.org/x/image/draw.Transform, with the
	// destination implicit in the Drawer target.
	//   - xform is the transform mapping source to destination coordinates.
	//   - src is the source image. Uniform does not work (or make sense) here.
	//   - sr is the source region, if zero full src is used; must have for Uniform.
	//   - op is the drawing operation: Src = copy source directly (blit),
	//     Over = alpha blend with existing.
	Transform(xform math32.Matrix3, src image.Image, sr image.Rectangle, op draw.Op)

	// Surface is the gpu device being drawn to.
	// Could be nil on unsupported devices (web).
	Surface() any
}

// DrawerBase is a base implementation of [Drawer] with basic no-ops
// for most methods. Embedders need to implement DestBounds and End.
type DrawerBase struct {
	// Image is the target render image
	Image *image.RGBA
}

// Copy copies the given Go source image to the render target, with the
// same semantics as golang.org/x/image/draw.Copy, with the destination
// implicit in the Drawer target.
//   - Must have called Start first!
//   - dp is the destination point.
//   - src is the source image. If an image.Uniform, fast Fill is done.
//   - sr is the source region, if zero full src is used; must have for Uniform.
//   - op is the drawing operation: Src = copy source directly (blit),
//     Over = alpha blend with existing.
func (dw *DrawerBase) Copy(dp image.Point, src image.Image, sr image.Rectangle, op draw.Op) {
	draw.Draw(dw.Image, image.Rectangle{dp, dp.Add(src.Bounds().Size())}, src, sr.Min, op)
}

// Scale copies the given Go source image to the render target,
// scaling the region defined by src and sr to the destination
// such that sr in src-space is mapped to dr in dst-space.
// with the same semantics as golang.org/x/image/draw.Scale, with the
// destination implicit in the Drawer target.
// If src image is an
//   - Must have called Start first!
//   - dr is the destination rectangle; if zero uses full dest image.
//   - src is the source image. Uniform does not work (or make sense) here.
//   - sr is the source region, if zero full src is used; must have for Uniform.
//   - op is the drawing operation: Src = copy source directly (blit),
//     Over = alpha blend with existing.
func (dw *DrawerBase) Scale(dr image.Rectangle, src image.Image, sr image.Rectangle, op draw.Op) {
	// todo: use drawmatrix and x/image to implement scale.
	draw.Draw(dw.Image, dr, src, sr.Min, op)
}

// Transform copies the given Go source image to the render target,
// with the same semantics as golang.org/x/image/draw.Transform, with the
// destination implicit in the Drawer target.
//   - xform is the transform mapping source to destination coordinates.
//   - src is the source image. Uniform does not work (or make sense) here.
//   - sr is the source region, if zero full src is used; must have for Uniform.
//   - op is the drawing operation: Src = copy source directly (blit),
//     Over = alpha blend with existing.
func (dw *DrawerBase) Transform(xform math32.Matrix3, src image.Image, sr image.Rectangle, op draw.Op) {
	// todo: use drawmatrix and x/image to implement transform
	draw.Draw(dw.Image, sr, src, sr.Min, op)
}

// Start starts recording a sequence of draw / fill actions,
// which will be performed on the GPU at End().
// This must be called prior to any Drawer operations.
func (dw *DrawerBase) Start() {
	// no-op
}

func (dw *DrawerBase) Surface() any {
	// no-op
	return nil
}
