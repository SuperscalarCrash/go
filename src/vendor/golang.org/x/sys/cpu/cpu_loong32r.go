// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build loong32r

package cpu

// LA32R exposes no optional ISA extensions used by this package. A 64-byte
// pad is conservative for the implementations targeted by this port.
const cacheLineSize = 64

func initOptions() {}
