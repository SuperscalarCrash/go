// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build loong32r

package cpu

// CacheLinePadSize is a conservative cache-line size for LA32R systems.
const CacheLinePadSize = 64

// LA32R has no optional instruction-set feature used by the Go baseline.
func doinit() {}
