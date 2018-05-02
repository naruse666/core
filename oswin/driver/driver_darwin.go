// Copyright 2018 The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// based on golang.org/x/exp/shiny:
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin

package driver

import (
	"github.com/goki/goki/gi/oswin"
	"github.com/goki/goki/gi/oswin/driver/gldriver"
)

func main(f func(oswin.App)) {
	gldriver.Main(f)
}
