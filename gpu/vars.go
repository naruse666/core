// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gpu

import (
	"fmt"
	"log"
	"strings"

	"cogentcore.org/core/base/indent"
	"cogentcore.org/core/vgpu/szalloc"

	vk "github.com/goki/WebGPU"
	"github.com/rajveermalviya/go-webgpu/wgpu"
)

// Vars are all the variables that are used by a pipeline,
// organized into Groups (optionally including the special VertexGroup
// or PushGroup).
// Vars are allocated to bindings sequentially in the order added.
type Vars struct {
	// map of Groups, by group number: VertexGroup is -2, PushGroup is -1,
	// rest are added incrementally.
	Groups map[int]*VarGroup

	// map of vars by different roles across all Groups, updated in Config(),
	// after all vars added.
	RoleMap map[VarRoles][]*Var

	// full set of BindGroupLayouts, one for each VarGroup >= 0
	Layouts []*wgpu.BindGroupLayout `display:"-"`

	// true if a VertexGroup has been added
	hasVertex bool `edit:"-"`

	// true if PushGroup has been added.  Note: not yet supported in WebGPU.
	hasPush bool `edit:"-"`

	// number of textures, at point of creating the DescLayout
	NTextures int
}

func (vs *Vars) Destroy(dev *Device) {
	for _, vg := range vs.Groups {
		vg.Destroy(dev)
	}
}

// AddVertexGroup adds a new Vertex Group.
// This is a special Group holding Vertex, Index vars
func (vs *Vars) AddVertexGroup() *VarGroup {
	if vs.Groups == nil {
		vs.Groups = make(map[int]*VarGroup)
	}
	vg := &VarGroup{Group: VertexGroup, ParentVars: vs}
	vs.Groups[VertexGroup] = vg
	vs.hasVertex = true
	return vg
}

// VertexGroup returns the Vertex Group -- a special Group holding Vertex, Index vars
func (vs *Vars) VertexGroup() *VarGroup {
	return vs.Groups[VertexGroup]
}

// AddPushGroup adds a new push constant Group -- this is a special Group holding
// values sent directly in the command buffer.
func (vs *Vars) AddPushGroup() *VarGroup {
	if vs.Groups == nil {
		vs.Groups = make(map[int]*VarGroup)
	}
	vg := &VarGroup{Group: PushGroup, ParentVars: vs}
	vs.Groups[PushGroup] = vg
	vs.hasPush = true
	return vg
}

// PushGroup returns the Push Group -- a special Group holding push constants
func (vs *Vars) PushGroup() *VarGroup {
	return vs.Groups[PushGroup]
}

// AddGroup adds a new non-Vertex Group for holding Uniforms, Storage, etc
// Groups are automatically numbered sequentially
func (vs *Vars) AddGroup() *VarGroup {
	if vs.Groups == nil {
		vs.Groups = make(map[int]*VarGroup)
	}
	idx := vs.NGroups()
	vg := &VarGroup{Group: idx, ParentVars: vs}
	vs.Groups[idx] = vg
	return vg
}

// VarByNameTry returns Var by name in given set number,
// returning error if not found
func (vs *Vars) VarByNameTry(set int, name string) (*Var, error) {
	vg, err := vs.GroupTry(set)
	if err != nil {
		return nil, err
	}
	return vg.VarByNameTry(name)
}

// ValueByNameTry returns value by first looking up variable name, then value name,
// within given set number, returning error if not found
func (vs *Vars) ValueByNameTry(set int, varName, valName string) (*Var, *Value, error) {
	vg, err := vs.GroupTry(set)
	if err != nil {
		return nil, nil, err
	}
	return vg.ValueByNameTry(varName, valName)
}

// ValueByIndexTry returns value by first looking up variable name, then value index,
// returning error if not found
func (vs *Vars) ValueByIndexTry(set int, varName string, valIndex int) (*Var, *Value, error) {
	vg, err := vs.GroupTry(set)
	if err != nil {
		return nil, nil, err
	}
	return vg.ValueByIndexTry(varName, valIndex)
}

// Config must be called after all variables have been added.
// Configures all Groups and also does validation, returning error
// does DescLayout too, so all ready for Pipeline config.
func (vs *Vars) Config() error {
	dev := vs.Mem.Device.Device
	ns := vs.NGroups()
	var cerr error
	vs.RoleMap = make(map[VarRoles][]*Var)
	for si := vs.StartGroup(); si < ns; si++ {
		vg := vs.Groups[si]
		if vg == nil {
			continue
		}
		err := vg.Config(dev)
		if err != nil {
			cerr = err
		}
		for ri, rl := range vg.RoleMap {
			vs.RoleMap[ri] = append(vs.RoleMap[ri], rl...)
		}
	}
	vs.BindLayout(dev)
	return cerr
}

// StringDoc returns info on variables
func (vs *Vars) StringDoc() string {
	ispc := 4
	var sb strings.Builder
	ns := vs.NGroups()
	for si := vs.StartGroup(); si < ns; si++ {
		vg := vs.Groups[si]
		if vg == nil {
			continue
		}
		sb.WriteString(fmt.Sprintf("Group: %d\n", vg.Group))

		for ri := Vertex; ri < VarRolesN; ri++ {
			rl, has := vg.RoleMap[ri]
			if !has || len(rl) == 0 {
				continue
			}
			sb.WriteString(fmt.Sprintf("%sRole: %s\n", indent.Spaces(1, ispc), ri.String()))
			for _, vr := range rl {
				sb.WriteString(fmt.Sprintf("%sVar: %s\n", indent.Spaces(2, ispc), vr.String()))
			}
		}
	}
	return sb.String()
}

// NGroups returns the number of regular non-VertexGroup sets
func (vs *Vars) NGroups() int {
	ex := 0
	if vs.hasVertex {
		ex++
	}
	if vs.hasPush {
		ex++
	}
	return len(vs.Groups) - ex
}

// StartGroup returns the starting set to use for iterating sets
func (vs *Vars) StartGroup() int {
	switch {
	case vs.hasVertex:
		return VertexGroup
	case vs.hasPush:
		return PushGroup
	default:
		return 0
	}
}

// GroupTry returns set by index, returning nil and error if not found
func (vs *Vars) GroupTry(set int) (*VarGroup, error) {
	vg, has := vs.Groups[set]
	if !has {
		err := fmt.Errorf("gpu.Vars:GroupTry set number %d not found", set)
		if Debug {
			log.Println(err)
		}
		return nil, err
	}
	return vg, nil
}

// VkVertexConfig returns WebGPU vertex config struct, for VertexGroup only!
// Note: there is no support for interleaved arrays so each binding and location
// is assigned the same sequential number, recorded in var Binding
func (vs *Vars) VkVertexConfig() *vk.PipelineVertexInputStateCreateInfo {
	if vs.hasVertex {
		return vs.Groups[VertexGroup].VkVertexConfig()
	}
	cfg := &vk.PipelineVertexInputStateCreateInfo{}
	cfg.SType = vk.StructureTypePipelineVertexInputStateCreateInfo
	return cfg
}

/*
// VkPushConfig returns WebGPU push constant ranges, only if PushGroup used.
func (vs *Vars) VkPushConfig() []vk.PushConstantRange {
	if vs.hasPush {
		return vs.Groups[PushGroup].VkPushConfig()
	}
	return nil
}
*/

///////////////////////////////////////////////////////////////////
// Binding, Layouts

// BindLayout configures the Layouts slice of BindGroupLayouts
// for all of the non-Vertex vars
func (vs *Vars) BindLayout(dev *Device) {
	vs.NTextures = 0
	if vs.NDescs < 1 {
		vs.NDescs = 1
	}
	nset := vs.NGroups()
	if nset == 0 {
		vs.Layouts = nil
		return
	}

	var lays []*wgpu.BindGroupLayout
	for si := 0; si < nset; si++ { // auto-skips vertex, push
		vg := vs.Groups[si]
		if vg == nil {
			continue
		}
		vg.BindLayout(dev, vs)
		vs.NTextures += vg.NTextures
		lays = append(lays, vg.Layout)
	}
}

/*
	bg, err = dev.Device.CreateBindGroup(&wgpu.BindGroupDescriptor{
		Label:  vl.Name,
		Layout: cameraBindGroupLayout,
		Entries: []wgpu.BindGroupEntry{{
			Binding: 0,
			Buffer:  vl.Bufer,
			Size:    wgpu.WholeSize,
		}},
	})
*/

// BindVertexValueName dynamically binds given VertexGroup value
// by name for given variable name.
// using given descIndex description set index (among the NDescs allocated).
//
// Value must have already been updated into device memory prior to this,
// ideally through a batch update prior to starting rendering, so that
// all the values are ready to be used during the render pass.
// This dynamically updates the offset to point to the specified val.
//
// Do NOT call BindValuesStart / End around this.
//
// returns error if not found.
func (vs *Vars) BindVertexValueName(varNm, valNm string) error {
	vg := vs.Groups[VertexGroup]
	vr, vl, err := vg.ValueByNameTry(varNm, valNm)
	if err != nil {
		return err
	}
	vr.BindValueIndex[vs.BindDescIndex] = vl.Index // this is then consumed by draw command
	return nil
}

// BindVertexValueIndex dynamically binds given VertexGroup value
// by index for given variable name.
// using given descIndex description set index (among the NDescs allocated).
//
// Value must have already been updated into device memory prior to this,
// ideally through a batch update prior to starting rendering, so that
// all the values are ready to be used during the render pass.
// This only dynamically updates the offset to point to the specified val.
//
// Do NOT call BindValuesStart / End around this.
//
// returns error if not found.
func (vs *Vars) BindVertexValueIndex(varNm string, valIndex int) error {
	vg := vs.Groups[VertexGroup]
	vr, vl, err := vg.ValueByIndexTry(varNm, valIndex)
	if err != nil {
		return err
	}
	vr.BindValueIndex[vs.BindDescIndex] = vl.Index // this is then consumed by draw command
	return nil
}

// TextureGroupSizeIndexes for texture at given index, allocated in groups by size
// using Values.AllocTexBySize, returns the indexes for the texture
// and layer to actually select the texture in the shader, and proportion
// of the Gp allocated texture size occupied by the texture.
func (vs *Vars) TextureGroupSizeIndexes(set int, varNm string, valIndex int) *szalloc.Indexes {
	vg, err := vs.GroupTry(set)
	if err != nil {
		return nil
	}
	return vg.TextureGroupSizeIndexes(vs, varNm, valIndex)
}

///////////////////////////////////////////////////////////
// Memory allocation

func (vs *Vars) MemSize(buff *Buffer) int {
	tsz := 0
	ns := vs.NGroups()
	for si := vs.StartGroup(); si < ns; si++ {
		vg := vs.Groups[si]
		if vg == nil {
			continue
		}
		for _, vr := range vg.Vars {
			if vr.Role.BuffType() != buff.Type {
				continue
			}
			tsz += vr.ValuesMemSize(buff.AlignBytes)
		}
	}
	return tsz
}

func (vs *Vars) MemSizeStorage(mm *Memory, alignBytes int) {
	ns := vs.NGroups()
	for si := vs.StartGroup(); si < ns; si++ {
		vg := vs.Groups[si]
		if vg == nil {
			continue
		}
		for _, vr := range vg.Vars {
			if vr.Role.BuffType() != StorageBuffer {
				continue
			}
			vr.MemSizeStorage(mm, alignBytes)
		}
	}
}

func (vs *Vars) AllocMem(buff *Buffer, offset int) int {
	ns := vs.NGroups()
	tsz := 0
	for si := vs.StartGroup(); si < ns; si++ {
		vg := vs.Groups[si]
		if vg == nil || vg.Group == PushGroup {
			continue
		}
		for _, vr := range vg.Vars {
			if vr.Role.BuffType() != buff.Type {
				continue
			}
			sz := vr.Values.AllocMem(vr, buff, offset)
			offset += sz
			tsz += sz
		}
	}
	return tsz
}

// Free resets the MemPtr for values, resets any self-owned resources (Textures)
func (vs *Vars) Free(buff *Buffer) {
	ns := vs.NGroups()
	for si := vs.StartGroup(); si < ns; si++ {
		vg := vs.Groups[si]
		if vg == nil || vg.Group == PushGroup {
			continue
		}
		for _, vr := range vg.Vars {
			if vr.Role.BuffType() != buff.Type {
				continue
			}
			vr.Values.Free()
		}
	}
}

// ModRegs returns the regions of Values that have been modified
func (vs *Vars) ModRegs(bt BufferTypes) []MemReg {
	ns := vs.NGroups()
	var mods []MemReg
	for si := vs.StartGroup(); si < ns; si++ {
		vg := vs.Groups[si]
		if vg == nil || vg.Group == PushSet {
			continue
		}
		for _, vr := range vg.Vars {
			if vr.Role.BuffType() != bt {
				continue
			}
			md := vr.Values.ModRegs(vr)
			mods = append(mods, md...)
		}
	}
	return mods
}

// ModRegStorage returns the regions of Storage Values that have been modified
func (vs *Vars) ModRegsStorage(bufIndex int, buff *Buffer) []MemReg {
	ns := vs.NSets()
	var mods []MemReg
	for si := vs.StartSet(); si < ns; si++ {
		vg := vs.SetMap[si]
		if vg == nil || vg.Set == PushSet {
			continue
		}
		for _, vr := range vg.Vars {
			if vr.Role.BuffType() != StorageBuffer {
				continue
			}
			if vr.StorageBuffer != bufIndex {
				continue
			}
			md := vr.Values.ModRegs(vr)
			mods = append(mods, md...)
		}
	}
	return mods
}

// AllocTextures allocates images on device memory
func (vs *Vars) AllocTextures(mm *Memory) {
	ns := vs.NSets()
	for si := vs.StartSet(); si < ns; si++ {
		vg := vs.SetMap[si]
		if vg == nil || vg.Set == PushSet {
			continue
		}
		for _, vr := range vg.Vars {
			if vr.Role != SampledTexture {
				continue
			}
			vr.Values.AllocTextures(mm)
		}
	}
}
