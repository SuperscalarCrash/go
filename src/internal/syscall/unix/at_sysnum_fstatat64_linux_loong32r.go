// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && loong32r

package unix

import "syscall"

const fstatatTrap uintptr = syscall.SYS_FSTATAT64
