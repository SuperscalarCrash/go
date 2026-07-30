// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && loong32r

#include "textflag.h"

TEXT _rt0_loong32r_linux(SB),NOSPLIT|NOFRAME,$0
	MOVW	0(R3), R4
	ADD	$4, R3, R5
	JMP	runtime·rt0_go(SB)

// When building with -buildmode=c-archive or c-shared, the linker places
// this C ABI entry point in .init_array.
TEXT _rt0_loong32r_linux_lib(SB),NOSPLIT|NOFRAME,$0
	JMP	_rt0_loong32r_lib(SB)

// In external linking, the C runtime calls main with argc and argv in R4 and
// R5. This is an LA32 ILP32S entry point; it does not share an LA64 startup
// symbol or implementation.
TEXT main(SB),NOSPLIT|NOFRAME,$0
	JMP	runtime·rt0_go(SB)
