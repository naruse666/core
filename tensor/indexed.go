// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tensor

import (
	"cmp"
	"errors"
	"math"
	"math/rand"
	"slices"
	"sort"
	"strings"
)

// Indexed is an indexed wrapper around a tensor.Tensor that provides a
// specific view onto the Tensor defined by the set of indexes, which
// apply to the outer-most row dimension (with default row-major indexing).
// This is the universal representation of a homogenous data type in the
// [tensor] package framework, from scalar to vector, matrix, and beyond,
// because it can efficiently represent any kind of element with sufficient
// flexibility to enable a huge range of computations to be elegantly expressed.
// For example, sorting and filtering a tensor only requires
// updating the indexes while doing nothing to the Tensor itself.
// To produce a new [Tensor] that has its raw data actually organized according
// to the indexed order, call the [NewTensor] method.
// Use the [Set]FloatRowCell methods wherever possible, for the most efficient
// and natural indirection through the indexes.  The 1D methods on underlying
// tensor data do not indirect through the indexes and must be called directly
// on the [Tensor].
type Indexed struct { //types:add

	// Tensor that we are an indexed view onto.
	Tensor Tensor

	// Indexes are the indexes into Tensor rows.
	Indexes []int
}

// NewIndexed returns a new Indexed based on given tensor.
// If a list of indexes is passed, then our indexes are initialized
// as a copy of those.  This is used e.g. from a Indext Table column.
// Otherwise it is initialized with sequential indexes.
func NewIndexed(tsr Tensor, idxs ...[]int) *Indexed {
	ix := &Indexed{}
	if len(idxs) == 1 {
		ix.Tensor = tsr
		ix.Indexes = slices.Clone(idxs[0])
	} else {
		ix.SetTensor(tsr)
	}
	return ix
}

// SetTensor sets as indexes into given tensor with sequential initial indexes
func (ix *Indexed) SetTensor(tsr Tensor) {
	ix.Tensor = tsr
	ix.Sequential()
}

// Len returns the length of the index list
func (ix *Indexed) Len() int {
	return len(ix.Indexes)
}

// DeleteInvalid deletes all invalid indexes from the list.
// Call this if rows (could) have been deleted from tensor.
func (ix *Indexed) DeleteInvalid() {
	if ix.Tensor == nil || ix.Tensor.DimSize(0) <= 0 {
		ix.Indexes = nil
		return
	}
	ni := ix.Len()
	for i := ni - 1; i >= 0; i-- {
		if ix.Indexes[i] >= ix.Tensor.DimSize(0) {
			ix.Indexes = append(ix.Indexes[:i], ix.Indexes[i+1:]...)
		}
	}
}

// Sequential sets indexes to sequential row-wise indexes into tensor.
func (ix *Indexed) Sequential() { //types:add
	if ix.Tensor == nil || ix.Tensor.DimSize(0) <= 0 {
		ix.Indexes = nil
		return
	}
	ix.Indexes = make([]int, ix.Tensor.DimSize(0))
	for i := range ix.Indexes {
		ix.Indexes[i] = i
	}
}

// ExcludeMissing1D deletes indexes for a 1D tensor (only) where
// the values are missing, as indicated by NaN.
func (ix *Indexed) ExcludeMissing1D() { //types:add
	if ix.Tensor == nil || ix.Tensor.DimSize(0) <= 0 {
		ix.Indexes = nil
		return
	}
	if ix.Tensor.NumDims() > 1 {
		return
	}
	ni := ix.Len()
	for i := ni - 1; i >= 0; i-- {
		if math.IsNaN(ix.Tensor.Float1D(ix.Indexes[i])) {
			ix.Indexes = append(ix.Indexes[:i], ix.Indexes[i+1:]...)
		}
	}
}

// Permuted sets indexes to a permuted order -- if indexes already exist
// then existing list of indexes is permuted, otherwise a new set of
// permuted indexes are generated
func (ix *Indexed) Permuted() {
	if ix.Tensor == nil || ix.Tensor.DimSize(0) <= 0 {
		ix.Indexes = nil
		return
	}
	if len(ix.Indexes) == 0 {
		ix.Indexes = rand.Perm(ix.Tensor.DimSize(0))
	} else {
		rand.Shuffle(len(ix.Indexes), func(i, j int) {
			ix.Indexes[i], ix.Indexes[j] = ix.Indexes[j], ix.Indexes[i]
		})
	}
}

// AddIndex adds a new index to the list
func (ix *Indexed) AddIndex(idx int) {
	ix.Indexes = append(ix.Indexes, idx)
}

const (
	// Ascending specifies an ascending sort direction for tensor Sort routines
	Ascending = true

	// Descending specifies a descending sort direction for tensor Sort routines
	Descending = false
)

// SortFunc sorts the indexes into 1D Tensor using given compare function.
// Returns an error if called on a higher-dimensional tensor.
// The compare function operates directly on row numbers into the Tensor
// as these row numbers have already been projected through the indexes.
// cmp(a, b) should return a negative number when a < b, a positive
// number when a > b and zero when a == b.
func (ix *Indexed) SortFunc(cmp func(tsr Tensor, i, j int) int) error {
	if ix.Tensor.NumDims() > 1 {
		return errors.New("tensor Sorting is only for 1D tensors")
	}
	slices.SortFunc(ix.Indexes, func(a, b int) int {
		return cmp(ix.Tensor, ix.Indexes[a], ix.Indexes[b])
	})
	return nil
}

// SortIndexes sorts the indexes into our Tensor directly in
// numerical order, producing the native ordering, while preserving
// any filtering that might have occurred.
func (ix *Indexed) SortIndexes() {
	sort.Ints(ix.Indexes)
}

// Sort compare function for string values.
func CompareStrings(a, b string, ascending bool) int {
	cmp := strings.Compare(a, b)
	if !ascending {
		cmp = -cmp
	}
	return cmp
}

func CompareNumbers(a, b float64, ascending bool) int {
	cmp := cmp.Compare(a, b)
	if !ascending {
		cmp = -cmp
	}
	return cmp
}

// Sort does default alpha or numeric sort of 1D tensor based on data type.
// Returns an error if called on a higher-dimensional tensor.
func (ix *Indexed) Sort(ascending bool) error {
	if ix.Tensor.NumDims() > 1 {
		return errors.New("tensor Sorting is only for 1D tensors")
	}
	if ix.Tensor.IsString() {
		ix.SortFunc(func(tsr Tensor, i, j int) int {
			return CompareStrings(tsr.String1D(i), tsr.String1D(j), ascending)
		})
	} else {
		ix.SortFunc(func(tsr Tensor, i, j int) int {
			return CompareNumbers(tsr.Float1D(i), tsr.Float1D(j), ascending)
		})
	}
	return nil
}

// SortStableFunc stably sorts the indexes of 1D Tensor using given compare function.
// The compare function operates directly on row numbers into the Tensor
// as these row numbers have already been projected through the indexes.
// cmp(a, b) should return a negative number when a < b, a positive
// number when a > b and zero when a == b.
// It is *essential* that it always returns 0 when the two are equal
// for the stable function to actually work.
func (ix *Indexed) SortStableFunc(cmp func(tsr Tensor, i, j int) int) {
	slices.SortStableFunc(ix.Indexes, func(a, b int) int {
		return cmp(ix.Tensor, ix.Indexes[a], ix.Indexes[b])
	})
}

// SortStable does default alpha or numeric stable sort
// of 1D tensor based on data type.
// Returns an error if called on a higher-dimensional tensor.
func (ix *Indexed) SortStable(ascending bool) error {
	if ix.Tensor.NumDims() > 1 {
		return errors.New("tensor Sorting is only for 1D tensors")
	}
	if ix.Tensor.IsString() {
		ix.SortStableFunc(func(tsr Tensor, i, j int) int {
			return CompareStrings(tsr.String1D(i), tsr.String1D(j), ascending)
		})
	} else {
		ix.SortStableFunc(func(tsr Tensor, i, j int) int {
			return CompareNumbers(tsr.Float1D(i), tsr.Float1D(j), ascending)
		})
	}
	return nil
}

// FilterFunc is a function used for filtering that returns
// true if Tensor row should be included in the current filtered
// view of the tensor, and false if it should be removed.
type FilterFunc func(tsr Tensor, row int) bool

// Filter filters the indexes into our Tensor using given Filter function.
// The Filter function operates directly on row numbers into the Tensor
// as these row numbers have already been projected through the indexes.
func (ix *Indexed) Filter(filterer func(tsr Tensor, row int) bool) {
	sz := len(ix.Indexes)
	for i := sz - 1; i >= 0; i-- { // always go in reverse for filtering
		if !filterer(ix.Tensor, ix.Indexes[i]) { // delete
			ix.Indexes = append(ix.Indexes[:i], ix.Indexes[i+1:]...)
		}
	}
}

// NewTensor returns a new tensor with column data organized according to
// the indexes.
func (ix *Indexed) NewTensor() Tensor {
	rows := len(ix.Indexes)
	nt := ix.Tensor.Clone()
	nt.SetNumRows(rows)
	return nt
}

// Clone returns a copy of the current Indexed view into the same
// underlying Tensor as the source, with its own index memory.
func (ix *Indexed) Clone() *Indexed {
	nix := &Indexed{}
	nix.CopyFrom(ix)
	return nix
}

// CopyFrom copies from given other Indexed (we have our own unique copy of indexes).
func (ix *Indexed) CopyFrom(oix *Indexed) {
	ix.Tensor = oix.Tensor
	ix.Indexes = slices.Clone(oix.Indexes)
}

// AddRows adds n rows to end of underlying Tensor, and to the indexes in this view
func (ix *Indexed) AddRows(n int) { //types:add
	stidx := ix.Tensor.DimSize(0)
	ix.Tensor.SetNumRows(stidx + n)
	for i := stidx; i < stidx+n; i++ {
		ix.Indexes = append(ix.Indexes, i)
	}
}

// InsertRows adds n rows to end of underlying Tensor, and to the indexes starting at
// given index in this view
func (ix *Indexed) InsertRows(at, n int) {
	stidx := ix.Tensor.DimSize(0)
	ix.Tensor.SetNumRows(stidx + n)
	nw := make([]int, n, n+len(ix.Indexes)-at)
	for i := 0; i < n; i++ {
		nw[i] = stidx + i
	}
	ix.Indexes = append(ix.Indexes[:at], append(nw, ix.Indexes[at:]...)...)
}

// DeleteRows deletes n rows of indexes starting at given index in the list of indexes
func (ix *Indexed) DeleteRows(at, n int) {
	ix.Indexes = append(ix.Indexes[:at], ix.Indexes[at+n:]...)
}

// Swap switches the indexes for i and j
func (ix *Indexed) Swap(i, j int) {
	ix.Indexes[i], ix.Indexes[j] = ix.Indexes[j], ix.Indexes[i]
}

// Float returns the value of given index as a float64.
// The first index value is indirected through the indexes.
func (ix *Indexed) Float(i []int) float64 {
	ic := slices.Clone(i)
	ic[0] = ix.Indexes[ic[0]]
	return ix.Tensor.Float(ic)
}

// SetFloat sets the value of given index as a float64
// The first index value is indirected through the [Indexes].
func (ix *Indexed) SetFloat(i []int, val float64) {
	ic := slices.Clone(i)
	ic[0] = ix.Indexes[ic[0]]
	ix.Tensor.SetFloat(ic, val)
}

// FloatRowCell returns the value at given row and cell,
// where row is outer-most dim, and cell is 1D index into remaining inner dims.
// Row is indirected through the [Indexes].
// This is the preferred interface for all Indexed operations.
func (ix *Indexed) FloatRowCell(row, cell int) float64 {
	return ix.Tensor.FloatRowCell(ix.Indexes[row], cell)
}

// SetFloatRowCell sets the value at given row and cell,
// where row is outer-most dim, and cell is 1D index into remaining inner dims.
// Row is indirected through the [Indexes].
// This is the preferred interface for all Indexed operations.
func (ix *Indexed) SetFloatRowCell(row, cell int, val float64) {
	ix.Tensor.SetFloatRowCell(ix.Indexes[row], cell, val)
}
