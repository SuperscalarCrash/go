// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// func Syscall6(num, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2, errno uintptr)
//
// LA32R's Go ABI has no register arguments. Load its independent 32-bit ABI0
// frame, then convert to the Linux syscall ABI: number in R11, arguments in
// R4-R9, result in R4, and negative errno values in the range -1..-4095.
TEXT ·Syscall6(SB),NOSPLIT|NOFRAME,$0-40
	MOVW	num+0(FP), R11
	MOVW	a1+4(FP), R4
	MOVW	a2+8(FP), R5
	MOVW	a3+12(FP), R6
	MOVW	a4+16(FP), R7
	MOVW	a5+20(FP), R8
	MOVW	a6+24(FP), R9
	SYSCALL
	MOVW	$-4096, R12
	BGEU	R12, R4, syscall6_ok
	MOVW	$-1, R12
	MOVW	R12, r1+28(FP)
	MOVW	R0, r2+32(FP)
	SUB	R4, R0, R4
	MOVW	R4, errno+36(FP)
	RET
syscall6_ok:
	MOVW	R4, r1+28(FP)
	MOVW	R0, r2+32(FP)
	MOVW	R0, errno+36(FP)
	RET
