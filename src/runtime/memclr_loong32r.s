// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build loong32r

#include "textflag.h"

// func memclrNoHeapPointers(ptr unsafe.Pointer, n uintptr)
TEXT runtime·memclrNoHeapPointers(SB),NOSPLIT|NOFRAME,$0-8
	MOVW	ptr+0(FP), R4
	MOVW	n+4(FP), R5
	BEQ	R5, memclr_done
	OR	R4, R5, R6
	AND	$3, R6, R6
	BNE	R6, memclr_bytes

memclr_words:
	BEQ	R5, memclr_done
	MOVW	R0, 0(R4)
	ADD	$4, R4
	ADD	$-4, R5
	JMP	memclr_words

memclr_bytes:
	BEQ	R5, memclr_done
	MOVB	R0, 0(R4)
	ADD	$1, R4
	ADD	$-1, R5
	JMP	memclr_bytes

memclr_done:
	RET
