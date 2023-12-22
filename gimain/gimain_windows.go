// Copyright (c) 2018, The Goki Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package gimain

import (
	"goki.dev/gi/v2/keyfun"
)

func init() {
	keyfun.DefaultMap = keyfun.MapName("WindowsStd")
	keyfun.SetActiveMapName(keyfun.DefaultMap)
}
