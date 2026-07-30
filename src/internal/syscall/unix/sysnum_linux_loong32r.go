// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && loong32r

package unix

// LA32R uses Linux's asm-generic syscall number space.
const (
	getrandomTrap       uintptr = 278
	copyFileRangeTrap   uintptr = 285
	pidfdSendSignalTrap uintptr = 424
	pidfdOpenTrap       uintptr = 434
	openat2Trap         uintptr = 437
)
