// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && loong32r && gc

#include "textflag.h"

// LA32R has an independent 32-bit Go ABI0 frame. Linux receives the syscall
// number in R11, arguments in R4-R9, and returns its result in R4.

TEXT ·Syscall(SB),NOSPLIT,$0-28
	JMP	syscall·Syscall(SB)

TEXT ·Syscall6(SB),NOSPLIT,$0-40
	JMP	syscall·Syscall6(SB)

TEXT ·SyscallNoError(SB),NOSPLIT,$0-24
	CALL	runtime·entersyscall(SB)
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
	CALL	runtime·exitsyscall(SB)
	RET

TEXT ·RawSyscall(SB),NOSPLIT,$0-28
	JMP	syscall·RawSyscall(SB)

TEXT ·RawSyscall6(SB),NOSPLIT,$0-40
	JMP	syscall·RawSyscall6(SB)

TEXT ·RawSyscallNoError(SB),NOSPLIT,$0-24
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
