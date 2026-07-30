// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build loong32r

#include "textflag.h"

// memmove uses word transfers whenever all operands are word aligned. This is
// required by the garbage collector: an aligned pointer-sized move must never
// expose a partially copied pointer. LA32R does not permit unaligned word
// accesses, so all other cases use byte transfers.
//
// func memmove(to, from unsafe.Pointer, n uintptr)
TEXT runtime·memmove(SB),NOSPLIT|NOFRAME,$0-12
	MOVW	to+0(FP), R4
	MOVW	from+4(FP), R5
	MOVW	n+8(FP), R6
	BEQ	R4, R5, memmove_done
	BEQ	R6, memmove_done

	ADD	R6, R4, R7
	ADD	R6, R5, R8
	// Copy backward when dst is above src. Copying in either direction is
	// safe for disjoint ranges, so an explicit overlap test is unnecessary.
	SGTU	R4, R5, R9
	BNE	R9, memmove_backward

	OR	R4, R5, R9
	OR	R6, R9, R9
	AND	$3, R9, R9
	BNE	R9, memmove_forward_bytes

memmove_forward_words:
	BEQ	R6, memmove_done
	MOVW	0(R5), R10
	MOVW	R10, 0(R4)
	ADD	$4, R5
	ADD	$4, R4
	ADD	$-4, R6
	JMP	memmove_forward_words

memmove_forward_bytes:
	BEQ	R4, R7, memmove_done
	MOVBU	0(R5), R10
	MOVB	R10, 0(R4)
	ADD	$1, R5
	ADD	$1, R4
	JMP	memmove_forward_bytes

memmove_backward:
	OR	R4, R5, R9
	OR	R6, R9, R9
	AND	$3, R9, R9
	BNE	R9, memmove_backward_bytes

memmove_backward_words:
	BEQ	R6, memmove_done
	ADD	$-4, R8
	ADD	$-4, R7
	MOVW	0(R8), R10
	MOVW	R10, 0(R7)
	ADD	$-4, R6
	JMP	memmove_backward_words

memmove_backward_bytes:
	BEQ	R6, memmove_done
	ADD	$-1, R8
	ADD	$-1, R7
	MOVBU	0(R8), R10
	MOVB	R10, 0(R7)
	ADD	$-1, R6
	JMP	memmove_backward_bytes

memmove_done:
	RET
