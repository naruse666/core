// Code generated by "gosl"; DO NOT EDIT

package main

import (
	"embed"
	"unsafe"
	"cogentcore.org/core/gpu"
)

//go:embed shaders/*.wgsl
var shaders embed.FS

// ComputeGPU is the compute gpu device
var ComputeGPU *gpu.GPU

// UseGPU indicates whether to use GPU vs. CPU.
var UseGPU bool

// GPUSystem is a GPU compute System with kernels operating on the
// same set of data variables.
var GPUSystem *gpu.ComputeSystem

// GPUVars is an enum for GPU variables, for specifying what to sync.
type GPUVars int32 //enums:enum

const (
	ParamsVar GPUVars = 0
	DataVar GPUVars = 1
)

// GPUInit initializes the GPU compute system,
// configuring system(s), variables and kernels.
func GPUInit() {
	gp := gpu.NewComputeGPU()
	ComputeGPU = gp
	{
		sy := gpu.NewComputeSystem(gp, "Default")
		GPUSystem = sy
		gpu.NewComputePipelineShaderFS(shaders, "shaders/Compute.wgsl", sy)
		vars := sy.Vars()
		{
			sgp := vars.AddGroup(gpu.Storage)
			var vr *gpu.Var
			_ = vr
			vr = sgp.AddStruct("Params", int(unsafe.Sizeof(ParamStruct{})), 1, gpu.ComputeShader)
			vr = sgp.Add("Data", gpu.Float32, 1, gpu.ComputeShader)
			sgp.SetNValues(1)
		}
		sy.Config()
	}
}

// GPURelease releases the GPU compute system resources.
// Call this at program exit.
func GPURelease() {
	GPUSystem.Release()
	ComputeGPU.Release()
}

// RunCompute runs the Compute kernel with given number of elements,
// on either the CPU or GPU depending on the UseGPU variable.
// Can call multiple Run* kernels in a row, which are then all launched
// in the same command submission on the GPU, which is by far the most efficient.
// MUST call RunDone (with optional vars to sync) after all Run calls.
// Alternatively, a single-shot RunOneCompute call does Run and Done for a
// single run-and-sync case.
func RunCompute(n int) {
	if UseGPU {
		RunComputeGPU(n)
	} else {
		RunComputeCPU(n)
	}
}

// RunComputeGPU runs the Compute kernel on the GPU. See [RunCompute] for more info.
func RunComputeGPU(n int) {
	sy := GPUSystem
	pl := sy.ComputePipelines["Compute"]
	ce, _ := sy.BeginComputePass()
	pl.Dispatch1D(ce, n, 64)
}

// RunComputeCPU runs the Compute kernel on the CPU.
func RunComputeCPU(n int) {
	// todo: need threaded api -- not tensor
	for i := range n {
		Compute(uint32(i))
	}
}

// RunOneCompute runs the Compute kernel with given number of elements,
// on either the CPU or GPU depending on the UseGPU variable.
// This version then calls RunDone with the given variables to sync
// after the Run, for a single-shot Run-and-Done call. If multiple kernels
// can be run in sequence, it is much more efficient to do multiple Run*
// calls followed by a RunDone call.
func RunOneCompute(n int, syncVars ...GPUVars) {
	if UseGPU {
		RunComputeGPU(n)
		RunDone(syncVars...)
	} else {
		RunComputeCPU(n)
	}
}
// RunDone must be called after Run* calls to start compute kernels.
// This actually submits the kernel jobs to the GPU, and adds commands
// to synchronize the given variables back from the GPU to the CPU.
// After this function completes, the GPU results will be available in 
// the specified variables.
func RunDone(syncVars ...GPUVars) {
	if !UseGPU {
		return
	}
	sy := GPUSystem
	sy.ComputeEncoder.End()
	ReadFromGPU(syncVars...)
	sy.EndComputePass()
	SyncFromGPU(syncVars...)
}

// ToGPU copies given variables to the GPU for the system.
func ToGPU(vars ...GPUVars) {
	sy := GPUSystem
	syVars := sy.Vars()
	for _, vr := range vars {
		switch vr {
		case ParamsVar:
			v, _ := syVars.ValueByIndex(0, "Params", 0)
			gpu.SetValueFrom(v, Params)
		case DataVar:
			v, _ := syVars.ValueByIndex(0, "Data", 0)
			gpu.SetValueFrom(v, Data.Values)
		}
	}
}

// ReadFromGPU starts the process of copying vars to the GPU.
func ReadFromGPU(vars ...GPUVars) {
	sy := GPUSystem
	syVars := sy.Vars()
	for _, vr := range vars {
		switch vr {
		case ParamsVar:
			v, _ := syVars.ValueByIndex(0, "Params", 0)
			v.GPUToRead(sy.CommandEncoder)
		case DataVar:
			v, _ := syVars.ValueByIndex(0, "Data", 0)
			v.GPUToRead(sy.CommandEncoder)
		}
	}
}

// SyncFromGPU synchronizes vars from the GPU to the actual variable.
func SyncFromGPU(vars ...GPUVars) {
	sy := GPUSystem
	syVars := sy.Vars()
	for _, vr := range vars {
		switch vr {
		case ParamsVar:
			v, _ := syVars.ValueByIndex(0, "Params", 0)
			v.ReadSync()
			gpu.ReadToBytes(v, Params)
		case DataVar:
			v, _ := syVars.ValueByIndex(0, "Data", 0)
			v.ReadSync()
			gpu.ReadToBytes(v, Data.Values)
		}
	}
}

// GetParams returns a pointer to the given global variable: 
// [Params] []ParamStruct at given index.
// To ensure that values are updated on the GPU, you must call [SetParams].
// after all changes have been made.
func GetParams(idx uint32) *ParamStruct {
	return &Params[idx]
}
