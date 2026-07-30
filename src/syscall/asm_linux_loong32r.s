// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && loong32r

#include "textflag.h"

// rawVforkSyscall performs a syscall without entering the scheduler. Linux
// receives LA32R syscall arguments in R4-R9 and the syscall number in R11.
TEXT ·rawVforkSyscall(SB),NOSPLIT|NOFRAME,$0-24
	MOVW	a1+4(FP), R4
	MOVW	a2+8(FP), R5
	MOVW	a3+12(FP), R6
	MOVW	R0, R7
	MOVW	R0, R8
	MOVW	R0, R9
	MOVW	trap+0(FP), R11
	SYSCALL
	MOVW	$-4096, R12
	BGEU	R12, R4, rawVforkOK
	MOVW	$-1, R12
	MOVW	R12, r1+16(FP)
	SUB	R4, R0, R4
	MOVW	R4, err+20(FP)
	RET
rawVforkOK:
	MOVW	R4, r1+16(FP)
	MOVW	R0, err+20(FP)
	RET

TEXT ·rawSyscallNoError(SB),NOSPLIT|NOFRAME,$0-24
	MOVW	a1+4(FP), R4
	MOVW	a2+8(FP), R5
	MOVW	a3+12(FP), R6
	MOVW	R0, R7
	MOVW	R0, R8
	MOVW	R0, R9
	MOVW	trap+0(FP), R11
	SYSCALL
	MOVW	R4, r1+16(FP)
	MOVW	R0, r2+20(FP)
	RET
