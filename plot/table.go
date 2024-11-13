// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plot

import (
	"fmt"

	"cogentcore.org/core/base/errors"
	"cogentcore.org/core/tensor"
	"cogentcore.org/core/tensor/table"
)

// NewTablePlot returns a new Plot with all configuration based on given
// [table.Table] set of columns and associated metadata, which must have
// [Stylers] functions set (e.g., [SetStylersTo]) that at least set basic
// table parameters, including:
//   - On: Set the main (typically Role = Y) column On to include in plot.
//   - Role: Set the appropriate [Roles] role for this column (Y, X, etc).
//   - Group: Multiple columns used for a given Plotter type must be grouped
//     together with a common name (typically the name of the main Y axis),
//     e.g., for Low, High error bars, Size, Color, etc. If only one On column,
//     then Group can be empty and all other such columns will be grouped.
//   - Plotter: Determines the type of Plotter element to use, which in turn
//     determines the additional Roles that can be used within a Group.
func NewTablePlot(dt *table.Table) (*Plot, error) {
	nc := len(dt.Columns.Values)
	if nc == 0 {
		return nil, errors.New("plot.NewTablePlot: no columns in data table")
	}
	csty := make(map[tensor.Values]*Style, nc)
	gps := make(map[string][]tensor.Values, nc)
	var xt tensor.Values // get the _last_ role = X column -- most specific counter
	var errs []error
	var pstySt Style // overall PlotStyle accumulator
	pstySt.Defaults()
	for _, cl := range dt.Columns.Values {
		st := &Style{}
		st.Defaults()
		stl := GetStylersFrom(cl)
		if stl == nil {
			continue
		}
		csty[cl] = st
		stl.Run(st)
		stl.Run(&pstySt)
		gps[st.Group] = append(gps[st.Group], cl)
		if st.Role == X {
			xt = cl
		}
	}
	psty := pstySt.Plot
	globalX := false
	if psty.XAxis.Column != "" {
		xc := dt.Columns.At(psty.XAxis.Column)
		if xc != nil {
			xt = xc
			globalX = true
		} else {
			errs = append(errs, errors.New("XAxis.Column name not found: "+psty.XAxis.Column))
		}
	}
	doneGps := map[string]bool{}
	plt := New()
	var legends []Thumbnailer // candidates for legend adding -- only add if > 1
	var legLabels []string
	for ci, cl := range dt.Columns.Values {
		cnm := dt.Columns.Keys[ci]
		st := csty[cl]
		if st == nil || !st.On || st.Role == X {
			continue
		}
		lbl := cnm
		if st.Label != "" {
			lbl = st.Label
		}
		gp := st.Group
		if doneGps[gp] {
			continue
		}
		if gp != "" {
			doneGps[gp] = true
		}
		ptyp := "XY"
		if st.Plotter != "" {
			ptyp = string(st.Plotter)
		}
		pt, err := PlotterByType(ptyp)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		data := Data{st.Role: cl}
		gcols := gps[gp]
		gotReq := true
		if globalX {
			data[X] = xt
		}
		for _, rl := range pt.Required {
			if rl == st.Role || (rl == X && globalX) {
				continue
			}
			got := false
			for _, gc := range gcols {
				gst := csty[gc]
				if gst.Role == rl {
					if rl == Y {
						if !gst.On {
							continue
						}
					}
					data[rl] = gc
					got = true
					if rl != X { // get the last one for X
						break
					}
				}
			}
			if !got {
				if rl == X && xt != nil {
					data[rl] = xt
				} else {
					err = fmt.Errorf("plot.NewTablePlot: Required Role %q not found in Group %q, Plotter %q not added for Column: %q", rl.String(), gp, ptyp, cnm)
					errs = append(errs, err)
					gotReq = false
					fmt.Println(err)
				}
			}
		}
		if !gotReq {
			continue
		}
		for _, rl := range pt.Optional {
			if rl == st.Role { // should not happen
				continue
			}
			for _, gc := range gcols {
				gst := csty[gc]
				if gst.Role == rl {
					data[rl] = gc
					break
				}
			}
		}
		pl := pt.New(data)
		if pl != nil {
			plt.Add(pl)
			if !st.NoLegend {
				if tn, ok := pl.(Thumbnailer); ok {
					legends = append(legends, tn)
					legLabels = append(legLabels, lbl)
				}
			}
		} else {
			err = fmt.Errorf("plot.NewTablePlot: error in creating plotter type: %q", ptyp)
			errs = append(errs, err)
		}
	}
	if len(legends) > 1 {
		for i, l := range legends {
			plt.Legend.Add(legLabels[i], l)
		}
	}
	return plt, errors.Join(errs...)
}
