// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build loong32r && gc

#include "textflag.h"

TEXT ·RewindAndSetgid(SB),NOSPLIT|NOFRAME,$0-0
	// Rewind stack pointer so anything that happens on the stack
	// will clobber the test pattern created by the caller.
	ADD	$(1024*8), R3

	// Ask signaller to setgid. DBAR 0 is a full memory barrier on LA32R.
	MOVW	$1, R12
	DBAR	$0
	MOVW	R12, ·Baton(SB)
	DBAR	$0

	// Wait for setgid completion.
loop:
	DBAR	$0
	MOVW	·Baton(SB), R12
	OR	R13, R13, R13 // hint that we're in a spin loop
	BNE	R12, loop
	DBAR	$0

	// Restore stack.
	ADD	$(-1024*8), R3
	RET
