// Copyright 2023 Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package offscreen

import (
	"github.com/naruse666/core/system"
	"github.com/naruse666/core/system/driver/base"
)

// Window is the implementation of [system.Window] for the offscreen platform.
type Window struct {
	base.WindowMulti[*App, *Drawer]
}

func (w *Window) Screen() *system.Screen {
	return TheApp.Screen(0)
}
