// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && loong32r

package unix

import "syscall"

// LA32R's asm-generic syscall table exposes renameat2 but not renameat.
// at.go supplies a zero flags argument, preserving renameat semantics.
const renameatTrap uintptr = syscall.SYS_RENAMEAT2
