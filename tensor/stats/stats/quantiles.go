// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stats

import (
	"math"

	"cogentcore.org/core/base/errors"
	"cogentcore.org/core/tensor"
)

// QuantilesFunc returns the given quantile(s) of non-NaN elements in given
// 1D tensor. Because sorting uses indexes, this only works for 1D case.
// If needed for a sub-space of values, that can be extracted through slicing
// and then used. Returns and logs an error if not 1D.
// qs are 0-1 values, 0 = min, 1 = max, .5 = median, etc.
// Uses linear interpolation.
// Because this requires a sort, it is more efficient to get as many quantiles
// as needed in one pass.
func QuantilesFunc(in, qs, out *tensor.Indexed) error {
	if in.Tensor.NumDims() != 1 {
		return errors.Log(errors.New("stats.QuantilesFunc: only 1D input tensors allowed"))
	}
	if qs.Tensor.NumDims() != 1 {
		return errors.Log(errors.New("stats.QuantilesFunc: only 1D quantile tensors allowed"))
	}
	sin := in.Clone()
	sin.ExcludeMissing1D()
	sin.Sort(tensor.Ascending)
	sz := len(sin.Indexes) - 1 // length of our own index list
	fsz := float64(sz)
	out.Tensor.SetShapeFrom(qs.Tensor)
	nq := qs.Tensor.Len()
	for i := range nq {
		q := qs.Tensor.Float1D(i)
		val := 0.0
		qi := q * fsz
		lwi := math.Floor(qi)
		lwii := int(lwi)
		if lwii >= sz {
			val = sin.FloatRowCell(sz, 0)
		} else if lwii < 0 {
			val = sin.FloatRowCell(0, 0)
		} else {
			phi := qi - lwi
			lwv := sin.FloatRowCell(lwii, 0)
			hiv := sin.FloatRowCell(lwii+1, 0)
			val = (1-phi)*lwv + phi*hiv
		}
		out.Tensor.SetFloat1D(val, i)
	}
	return nil
}

// MedianFunc computes the median (50% quantile) of tensor values.
// See [StatsFunc] for general information.
func MedianFunc(in, out *tensor.Indexed) {
	QuantilesFunc(in, tensor.NewIndexed(tensor.NewNumberFromSlice([]float64{.5})), out)
}

// Q1Func computes the first quantile (25%) of tensor values.
// See [StatsFunc] for general information.
func Q1Func(in, out *tensor.Indexed) {
	QuantilesFunc(in, tensor.NewIndexed(tensor.NewNumberFromSlice([]float64{.25})), out)
}

// Q3Func computes the third quantile (75%) of tensor values.
// See [StatsFunc] for general information.
func Q3Func(in, out *tensor.Indexed) {
	QuantilesFunc(in, tensor.NewIndexed(tensor.NewNumberFromSlice([]float64{.75})), out)
}
