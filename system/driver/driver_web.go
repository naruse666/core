// Copyright 2023 Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js && !offscreen

package driver

import (
	"github.com/naruse666/core/system/driver/web"
)

func init() {
	web.Init()
}
