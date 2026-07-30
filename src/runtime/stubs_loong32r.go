// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build loong32r

package runtime

func load_g()
func save_g()
func spillArgs()
func unspillArgs()

//go:nosplit
func getfp() uintptr { return 0 }
