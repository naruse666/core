// Copyright 2024 Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gpudraw

import (
	"embed"
	"image/draw"
	"unsafe"

	"cogentcore.org/core/base/errors"
	"cogentcore.org/core/base/slicesx"
	"cogentcore.org/core/gpu"
	"github.com/rajveermalviya/go-webgpu/wgpu"
)

//go:embed shaders/*.wgsl
var shaders embed.FS

// ConfigPipeline configures graphics settings on the pipeline
func (dw *Drawer) ConfigPipeline(pl *gpu.GraphicsPipeline) {
	pl.SetGraphicsDefaults()
	pl.SetCullMode(wgpu.CullModeNone)
	// if dw.YIsDown {
	// 	pl.SetFrontFace(wgpu.FrontFaceCCW)
	// } else {
	pl.SetFrontFace(wgpu.FrontFaceCW)
	// }
}

// configSystem configures GPUDraw sytem
func (dw *Drawer) configSystem(gp *gpu.GPU, dev *gpu.Device, renderFormat *gpu.TextureFormat) {
	dw.YIsDown = false
	dw.opList = slicesx.SetLength(dw.opList, AllocChunk) // allocate
	dw.opList = dw.opList[:0]

	dw.Sys = gpu.NewGraphicsSystem(gp, "gpudraw", dev)
	dw.Sys.ConfigRender(renderFormat, gpu.UndefType)
	sy := dw.Sys
	// sy.SetClearColor(color.RGBA{50, 50, 50, 255})

	// note: requires different pipelines for src vs. over draw op modes
	dspl := sy.AddGraphicsPipeline("drawsrc")
	dw.ConfigPipeline(dspl)
	dspl.SetColorBlend(false)

	dopl := sy.AddGraphicsPipeline("drawover")
	dw.ConfigPipeline(dopl)
	dopl.SetColorBlend(true) // default

	fpl := sy.AddGraphicsPipeline("fill")
	dw.ConfigPipeline(dopl)

	sh := dspl.AddShader("draw")
	sh.OpenFileFS(shaders, "shaders/draw.wgsl")
	dspl.AddEntry(sh, gpu.VertexShader, "vs_main")
	dspl.AddEntry(sh, gpu.FragmentShader, "fs_main")

	sh = dopl.AddShader("draw")
	sh.OpenFileFS(shaders, "shaders/draw.wgsl")
	dopl.AddEntry(sh, gpu.VertexShader, "vs_main")
	dopl.AddEntry(sh, gpu.FragmentShader, "fs_main")

	sh = fpl.AddShader("fill")
	sh.OpenFileFS(shaders, "shaders/fill.wgsl")
	fpl.AddEntry(sh, gpu.VertexShader, "vs_main")
	fpl.AddEntry(sh, gpu.FragmentShader, "fs_main")

	vgp := sy.Vars.AddVertexGroup()
	mgp := sy.Vars.AddGroup(gpu.Uniform, "Matrix")         // 0
	tgp := sy.Vars.AddGroup(gpu.SampledTexture, "Texture") // 1

	posv := vgp.Add("Pos", gpu.Float32Vector2, 0, gpu.VertexShader)
	idxv := vgp.Add("Index", gpu.Uint16, 0, gpu.VertexShader)
	idxv.Role = gpu.Index

	mv := mgp.AddStruct("Matrix", int(unsafe.Sizeof(Matrix{})), 1, gpu.VertexShader, gpu.FragmentShader)
	mv.DynamicOffset = true

	tgp.Add("TexSampler", gpu.TextureRGBA32, 1, gpu.FragmentShader)

	vgp.SetNValues(1)
	mgp.SetNValues(1)
	tgp.SetNValues(AllocChunk)

	sy.Config()

	rectPos := posv.Values.Values[0]
	gpu.SetValueFrom(rectPos, []float32{
		0.0, 0.0,
		0.0, 1.0,
		1.0, 0.0,
		1.0, 1.0})

	rectIndex := idxv.Values.Values[0]
	gpu.SetValueFrom(rectIndex, []uint16{0, 1, 2, 2, 1, 3})

	vl := sy.Vars.ValueByIndex(0, "Matrix", 0)
	vl.DynamicN = AllocChunk
}

func (dw *Drawer) drawAll() error {
	sy := dw.Sys

	vl := sy.Vars.ValueByIndex(0, "Matrix", 0)
	vl.WriteDynamicBuffer()

	var view *wgpu.TextureView
	var err error
	if dw.surface != nil {
		view, err = dw.surface.AcquireNextTexture()
		if errors.Log(err) != nil {
			return err
		}
	} else {
		// todo!
		// sy.ResetBeginRenderPassNoClear(cmd, dw.Frame.Frames[0], descIndex)
	}

	mvr := sy.Vars.VarByName(0, "Matrix")
	mvl := mvr.Values.Values[0]
	tvr := sy.Vars.VarByName(1, "TexSampler")
	tvr.Values.Current = 0

	cmd := sy.NewCommandEncoder()
	rp := sy.BeginRenderPassNoClear(cmd, view) // NoClear

	imgIdx := 0
	lastOp := draw.Op(-1)
	_ = lastOp
	for i, op := range dw.opList {
		var pl *gpu.GraphicsPipeline
		switch op {
		case draw.Src:
			pl = sy.GraphicsPipelines["drawsrc"]
		case draw.Over:
			pl = sy.GraphicsPipelines["drawover"]
		default:
			pl = sy.GraphicsPipelines["fill"]
		}
		mvl.DynamicIndex = i
		if op != Fill {
			tvr.Values.Current = imgIdx
			imgIdx++
		}
		if op != lastOp {
			pl.BindPipeline(rp)
			lastOp = op
		} else {
			pl.BindAllGroups(rp)
		}
		pl.BindDrawIndexed(rp)
	}
	rp.End()
	if dw.surface != nil {
		dw.surface.SubmitRender(cmd)
		dw.surface.Present()
	} else {
		// dw.Frame.SubmitRender(cmd)
		// dw.Frame.WaitForRender()
	}
	return nil
}
