// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build loong32r

package goarch

const (
	_ArchFamily          = LOONG32R
	_DefaultPhysPageSize = 4096
	_PCQuantum           = 4
	_MinFrameSize        = 4
	_StackAlign          = 16
)
