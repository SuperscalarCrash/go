// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build loong32r

#include "textflag.h"

#define REGCTXT R29

// func memequal(a, b unsafe.Pointer, size uintptr) bool
TEXT runtime·memequal(SB),NOSPLIT|NOFRAME,$0-13
	MOVW	a+0(FP), R4
	MOVW	b+4(FP), R5
	BEQ	R4, R5, memequal_equal
	MOVW	size+8(FP), R6
	ADD	R6, R4, R7
memequal_loop:
	BEQ	R4, R7, memequal_equal
	MOVBU	0(R4), R8
	MOVBU	0(R5), R9
	BNE	R8, R9, memequal_notequal
	ADD	$1, R4
	ADD	$1, R5
	JMP	memequal_loop
memequal_notequal:
	MOVB	R0, ret+12(FP)
	RET
memequal_equal:
	MOVW	$1, R4
	MOVB	R4, ret+12(FP)
	RET

// func memequal_varlen(a, b unsafe.Pointer) bool
TEXT runtime·memequal_varlen(SB),NOSPLIT|NOFRAME,$0-9
	MOVW	a+0(FP), R4
	MOVW	b+4(FP), R5
	BEQ	R4, R5, memequal_var_equal
	MOVW	4(REGCTXT), R6
	ADD	R6, R4, R7
memequal_var_loop:
	BEQ	R4, R7, memequal_var_equal
	MOVBU	0(R4), R8
	MOVBU	0(R5), R9
	BNE	R8, R9, memequal_var_notequal
	ADD	$1, R4
	ADD	$1, R5
	JMP	memequal_var_loop
memequal_var_notequal:
	MOVB	R0, ret+8(FP)
	RET
memequal_var_equal:
	MOVW	$1, R4
	MOVB	R4, ret+8(FP)
	RET
