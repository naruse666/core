// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package table

import (
	"fmt"
	"slices"
	"sort"
	"strings"

	"cogentcore.org/core/base/errors"
)

// SplitAgg contains aggregation results for splits
type SplitAgg struct {

	// the name of the aggregation operation performed, e.g., Sum, Mean, etc
	Name string

	// column index on which the aggregation was performed -- results will have same shape as cells in this column
	ColumnIndex int

	// aggregation results -- outer index is length of splits, inner is the length of the cell shape for the column
	Aggs [][]float64
}

// Splits is a list of indexed views into a given Table, that represent a particular
// way of splitting up the data, e.g., whenever a given column value changes.
//
// It is functionally equivalent to the MultiIndex in python's pandas: it has multiple
// levels of indexes as listed in the Levels field, which then have corresponding
// Values for each split.  These index levels can be re-ordered, and new Splits or
// Indexeds's can be created from subsets of the existing levels.  The Values are
// stored simply as string values, as this is the most general type and often
// index values are labels etc.
//
// For Splits created by the splits.GroupBy function for example, each index Level is
// the column name that the data was grouped by, and the Values for each split are then
// the values of those columns.  However, any arbitrary set of levels and values can
// be used, e.g., as in the splits.GroupByFunc function.
//
// Conceptually, a given Split always contains the full "outer product" of all the
// index levels -- there is one split for each unique combination of values along each
// index level.  Thus, removing one level collapses across those values and moves the
// corresponding indexes into the remaining split indexes.
//
// You can Sort and Filter based on the index values directly, to reorganize the splits
// and drop particular index values, etc.
//
// Splits also maintains Aggs aggregate values for each split, which can be computed using
// standard aggregation methods over data columns, using the split.Agg* functions.
//
// The table code contains the structural methods for managing the Splits data.
// See split package for end-user methods to generate different kinds of splits,
// and perform aggregations, etc.
type Splits struct {

	// the list of index views for each split
	Splits []*Indexed

	// levels of indexes used to organize the splits -- each split contains the full outer product across these index levels.  for example, if the split was generated by grouping over column values, then these are the column names in order of grouping.  the splits are not automatically sorted hierarchically by these levels but e.g., the GroupBy method produces that result -- use the Sort methods to explicitly sort.
	Levels []string

	// the values of the index levels associated with each split.  The outer dimension is the same length as Splits, and the inner dimension is the levels.
	Values [][]string

	// aggregate results, one for each aggregation operation performed -- split-level data is contained within each SplitAgg struct -- deleting a split removes these aggs but adding new splits just invalidates all existing aggs (they are automatically deleted).
	Aggs []*SplitAgg

	// current Less function used in sorting
	lessFunc SplitsLessFunc `copier:"-" display:"-" xml:"-" json:"-"`
}

// SplitsLessFunc is a function used for sort comparisons that returns
// true if split i is less than split j
type SplitsLessFunc func(spl *Splits, i, j int) bool

// Len returns number of splits
func (spl *Splits) Len() int {
	return len(spl.Splits)
}

// Table returns the table from the first split (should be same for all)
// returns nil if no splits yet
func (spl *Splits) Table() *Table {
	if len(spl.Splits) == 0 {
		return nil
	}
	return spl.Splits[0].Table
}

// New adds a new split to the list for given table, and with associated
// values, which are copied before saving into Values list, and any number of rows
// from the table associated with this split (also copied).
// Any existing Aggs are deleted by this.
func (spl *Splits) New(dt *Table, values []string, rows ...int) *Indexed {
	spl.Aggs = nil
	ix := &Indexed{Table: dt}
	spl.Splits = append(spl.Splits, ix)
	if len(rows) > 0 {
		ix.Indexes = append(ix.Indexes, slices.Clone(rows)...)
	}
	if len(values) > 0 {
		spl.Values = append(spl.Values, slices.Clone(values))
	} else {
		spl.Values = append(spl.Values, nil)
	}
	return ix
}

// ByValue finds split indexes by matching to split values, returns nil if not found.
// values are used in order as far as they go and any remaining values are assumed
// to match, and any empty values will match anything.  Can use this to access different
// subgroups within overall set of splits.
func (spl *Splits) ByValue(values []string) []int {
	var matches []int
	for si, sn := range spl.Values {
		sz := min(len(sn), len(values))
		match := true
		for j := 0; j < sz; j++ {
			if values[j] == "" {
				continue
			}
			if values[j] != sn[j] {
				match = false
				break
			}
		}
		if match {
			matches = append(matches, si)
		}
	}
	return matches
}

// Delete deletes split at given index -- use this to coordinate deletion
// of Splits, Values, and Aggs values for given split
func (spl *Splits) Delete(idx int) {
	spl.Splits = append(spl.Splits[:idx], spl.Splits[idx+1:]...)
	spl.Values = append(spl.Values[:idx], spl.Values[idx+1:]...)
	for _, ag := range spl.Aggs {
		ag.Aggs = append(ag.Aggs[:idx], ag.Aggs[idx+1:]...)
	}
}

// Filter removes any split for which given function returns false
func (spl *Splits) Filter(fun func(idx int) bool) {
	sz := len(spl.Splits)
	for si := sz - 1; si >= 0; si-- {
		if !fun(si) {
			spl.Delete(si)
		}
	}
}

// Sort sorts the splits according to the given Less function.
func (spl *Splits) Sort(lessFunc func(spl *Splits, i, j int) bool) {
	spl.lessFunc = lessFunc
	sort.Sort(spl)
}

// SortLevels sorts the splits according to the current index level ordering of values
// i.e., first index level is outer sort dimension, then within that is the next, etc
func (spl *Splits) SortLevels() {
	spl.Sort(func(sl *Splits, i, j int) bool {
		vli := sl.Values[i]
		vlj := sl.Values[j]
		for k := range vli {
			if vli[k] < vlj[k] {
				return true
			} else if vli[k] > vlj[k] {
				return false
			} // fallthrough
		}
		return false
	})
}

// SortOrder sorts the splits according to the given ordering of index levels
// which can be a subset as well
func (spl *Splits) SortOrder(order []int) error {
	if len(order) == 0 || len(order) > len(spl.Levels) {
		return fmt.Errorf("table.Splits SortOrder: order length == 0 or > Levels")
	}
	spl.Sort(func(sl *Splits, i, j int) bool {
		vli := sl.Values[i]
		vlj := sl.Values[j]
		for k := range order {
			if vli[order[k]] < vlj[order[k]] {
				return true
			} else if vli[order[k]] > vlj[order[k]] {
				return false
			} // fallthrough
		}
		return false
	})
	return nil
}

// ReorderLevels re-orders the index levels according to the given new ordering indexes
// e.g., []int{1,0} will move the current level 0 to level 1, and 1 to level 0
// no checking is done to ensure these are sensible beyond basic length test --
// behavior undefined if so.  Typically you want to call SortLevels after this.
func (spl *Splits) ReorderLevels(order []int) error {
	nlev := len(spl.Levels)
	if len(order) != nlev {
		return fmt.Errorf("table.Splits ReorderLevels: order length != Levels")
	}
	old := make([]string, nlev)
	copy(old, spl.Levels)
	for i := range order {
		spl.Levels[order[i]] = old[i]
	}
	for si := range spl.Values {
		copy(old, spl.Values[si])
		for i := range order {
			spl.Values[si][order[i]] = old[i]
		}
	}
	return nil
}

// ExtractLevels returns a new Splits that only has the given levels of indexes,
// in their given order, with the other levels removed and their corresponding indexes
// merged into the appropriate remaining levels.
// Any existing aggregation data is not retained in the new splits.
func (spl *Splits) ExtractLevels(levels []int) (*Splits, error) {
	nlv := len(levels)
	if nlv == 0 || nlv >= len(spl.Levels) {
		return nil, fmt.Errorf("table.Splits ExtractLevels: levels length == 0 or >= Levels")
	}
	aggs := spl.Aggs
	spl.Aggs = nil
	ss := spl.Clone()
	spl.Aggs = aggs
	ss.SortOrder(levels)
	// now just do the grouping by levels values
	lstValues := make([]string, nlv)
	curValues := make([]string, nlv)
	var curIx *Indexed
	nsp := len(ss.Splits)
	for si := nsp - 1; si >= 0; si-- {
		diff := false
		for li := range levels {
			vl := ss.Values[si][levels[li]]
			curValues[li] = vl
			if vl != lstValues[li] {
				diff = true
			}
		}
		if diff || curIx == nil {
			curIx = ss.Splits[si]
			copy(lstValues, curValues)
			ss.Values[si] = slices.Clone(curValues)
		} else {
			curIx.Indexes = append(curIx.Indexes, ss.Splits[si].Indexes...) // absorb
			ss.Delete(si)
		}
	}
	ss.Levels = make([]string, nlv)
	for li := range levels {
		ss.Levels[li] = spl.Levels[levels[li]]
	}
	return ss, nil
}

// Clone returns a cloned copy of our SplitAgg
func (sa *SplitAgg) Clone() *SplitAgg {
	nsa := &SplitAgg{}
	nsa.CopyFrom(sa)
	return nsa
}

// CopyFrom copies from other SplitAgg -- we get our own unique copy of everything
func (sa *SplitAgg) CopyFrom(osa *SplitAgg) {
	sa.Name = osa.Name
	sa.ColumnIndex = osa.ColumnIndex
	nags := len(osa.Aggs)
	if nags > 0 {
		sa.Aggs = make([][]float64, nags)
		for si := range osa.Aggs {
			sa.Aggs[si] = slices.Clone(osa.Aggs[si])
		}
	}
}

// Clone returns a cloned copy of our splits
func (spl *Splits) Clone() *Splits {
	nsp := &Splits{}
	nsp.CopyFrom(spl)
	return nsp
}

// CopyFrom copies from other Splits -- we get our own unique copy of everything
func (spl *Splits) CopyFrom(osp *Splits) {
	spl.Splits = make([]*Indexed, len(osp.Splits))
	spl.Values = make([][]string, len(osp.Values))
	for si := range osp.Splits {
		spl.Splits[si] = osp.Splits[si].Clone()
		spl.Values[si] = slices.Clone(osp.Values[si])
	}
	spl.Levels = slices.Clone(osp.Levels)

	nag := len(osp.Aggs)
	if nag > 0 {
		spl.Aggs = make([]*SplitAgg, nag)
		for ai := range osp.Aggs {
			spl.Aggs[ai] = osp.Aggs[ai].Clone()
		}
	}
}

// AddAgg adds a new set of aggregation results for the Splits
func (spl *Splits) AddAgg(name string, colIndex int) *SplitAgg {
	ag := &SplitAgg{Name: name, ColumnIndex: colIndex}
	spl.Aggs = append(spl.Aggs, ag)
	return ag
}

// DeleteAggs deletes all existing aggregation data
func (spl *Splits) DeleteAggs() {
	spl.Aggs = nil
}

// AggByName returns Agg results for given name, which does NOT include the
// column name, just the name given to the Agg result
// (e.g., Mean for a standard Mean agg).
// Returns error message if not found.
func (spl *Splits) AggByName(name string) (*SplitAgg, error) {
	for _, ag := range spl.Aggs {
		if ag.Name == name {
			return ag, nil
		}
	}
	return nil, fmt.Errorf("table.Splits AggByName: agg results named: %v not found", name)
}

// AggByColumnName returns Agg results for given column name,
// optionally including :Name agg name appended, where Name
// is the name given to the Agg result (e.g., Mean for a standard Mean agg).
// Returns error message if not found.
func (spl *Splits) AggByColumnName(name string) (*SplitAgg, error) {
	dt := spl.Table()
	if dt == nil {
		return nil, fmt.Errorf("table.Splits AggByColumnName: table nil")
	}
	nmsp := strings.Split(name, ":")
	colIndex, err := dt.ColumnIndex(nmsp[0])
	if err != nil {
		return nil, err
	}
	for _, ag := range spl.Aggs {
		if ag.ColumnIndex != colIndex {
			continue
		}
		if len(nmsp) == 2 && nmsp[1] != ag.Name {
			continue
		}
		return ag, nil
	}
	return nil, fmt.Errorf("table.Splits AggByColumnName: agg results named: %v not found", name)
}

// SetLevels sets the Levels index names -- must match actual index dimensionality
// of the Values.  This is automatically done by e.g., GroupBy, but must be done
// manually if creating custom indexes.
func (spl *Splits) SetLevels(levels ...string) {
	spl.Levels = levels
}

// use these for arg to ArgsToTable*
const (
	// ColumnNameOnly means resulting agg table just has the original column name, no aggregation name
	ColumnNameOnly bool = true
	// AddAggName means resulting agg table columns have aggregation name appended
	AddAggName = false
)

// AggsToTable returns a Table containing this Splits' aggregate data.
// Must have Levels and Aggs all created as in the split.Agg* methods.
// if colName == ColumnNameOnly, then the name of the columns for the Table
// is just the corresponding agg column name -- otherwise it also includes
// the name of the aggregation function with a : divider (e.g., Name:Mean)
func (spl *Splits) AggsToTable(colName bool) *Table {
	nsp := len(spl.Splits)
	if nsp == 0 {
		return nil
	}
	dt := spl.Splits[0].Table
	st := NewTable().SetNumRows(nsp)
	for _, cn := range spl.Levels {
		oc, _ := dt.ColumnByName(cn)
		if oc != nil {
			st.AddColumnOfType(oc.DataType(), cn)
		} else {
			st.AddStringColumn(cn)
		}
	}
	for _, ag := range spl.Aggs {
		col := dt.Columns[ag.ColumnIndex]
		an := dt.ColumnNames[ag.ColumnIndex]
		if colName == AddAggName {
			an += ":" + ag.Name
		}
		st.AddFloat64TensorColumn(an, col.Shape().Sizes[1:]...)
	}
	for si := range spl.Splits {
		cidx := 0
		for ci := range spl.Levels {
			col := st.Columns[cidx]
			col.SetString1D(spl.Values[si][ci], si)
			cidx++
		}
		for _, ag := range spl.Aggs {
			col := st.Columns[cidx]
			_, csz := col.RowCellSize()
			sti := si * csz
			av := ag.Aggs[si]
			for j, a := range av {
				col.SetFloat1D(a, sti+j)
			}
			cidx++
		}
	}
	return st
}

// AggsToTableCopy returns a Table containing this Splits' aggregate data
// and a copy of the first row of data for each split for all non-agg cols,
// which is useful for recording other data that goes along with aggregated values.
// Must have Levels and Aggs all created as in the split.Agg* methods.
// if colName == ColumnNameOnly, then the name of the columns for the Table
// is just the corresponding agg column name -- otherwise it also includes
// the name of the aggregation function with a : divider (e.g., Name:Mean)
func (spl *Splits) AggsToTableCopy(colName bool) *Table {
	nsp := len(spl.Splits)
	if nsp == 0 {
		return nil
	}
	dt := spl.Splits[0].Table
	st := NewTable().SetNumRows(nsp)
	exmap := make(map[string]struct{})
	for _, cn := range spl.Levels {
		st.AddStringColumn(cn)
		exmap[cn] = struct{}{}
	}
	for _, ag := range spl.Aggs {
		col := dt.Columns[ag.ColumnIndex]
		an := dt.ColumnNames[ag.ColumnIndex]
		exmap[an] = struct{}{}
		if colName == AddAggName {
			an += ":" + ag.Name
		}
		st.AddFloat64TensorColumn(an, col.Shape().Sizes[1:]...)
	}
	var cpcol []string
	for _, cn := range dt.ColumnNames {
		if _, ok := exmap[cn]; !ok {
			cpcol = append(cpcol, cn)
			col := errors.Log1(dt.ColumnByName(cn))
			st.AddColumn(col.Clone(), cn)
		}
	}
	for si, sidx := range spl.Splits {
		cidx := 0
		for ci := range spl.Levels {
			col := st.Columns[cidx]
			col.SetString1D(spl.Values[si][ci], si)
			cidx++
		}
		for _, ag := range spl.Aggs {
			col := st.Columns[cidx]
			_, csz := col.RowCellSize()
			sti := si * csz
			av := ag.Aggs[si]
			for j, a := range av {
				col.SetFloat1D(a, sti+j)
			}
			cidx++
		}
		if len(sidx.Indexes) > 0 {
			stidx := sidx.Indexes[0]
			for _, cn := range cpcol {
				st.CopyCell(cn, si, dt, cn, stidx)
			}
		}
	}
	return st
}

// Less calls the LessFunc for sorting
func (spl *Splits) Less(i, j int) bool {
	return spl.lessFunc(spl, i, j)
}

// Swap switches the indexes for i and j
func (spl *Splits) Swap(i, j int) {
	spl.Splits[i], spl.Splits[j] = spl.Splits[j], spl.Splits[i]
	spl.Values[i], spl.Values[j] = spl.Values[j], spl.Values[i]
	for _, ag := range spl.Aggs {
		ag.Aggs[i], ag.Aggs[j] = ag.Aggs[j], ag.Aggs[i]
	}
}
