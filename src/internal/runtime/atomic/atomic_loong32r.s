// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build loong32r

#include "textflag.h"

// All operations use only the LA32R LL.W/SC.W pair and base barriers.

TEXT ·Cas(SB),NOSPLIT,$0-13
	MOVW	ptr+0(FP), R4
	MOVW	old+4(FP), R5
	MOVW	new+8(FP), R6
	DBAR	$0
cas_loop:
	LL	(R4), R8
	BNE	R5, R8, cas_fail
	MOVW	R6, R7
	SC	R7, (R4)
	BEQ	R7, cas_loop
	DBAR	$0
	MOVB	R7, ret+12(FP)
	RET
cas_fail:
	DBAR	$0
	MOVB	R0, ret+12(FP)
	RET

TEXT ·Xadd(SB),NOSPLIT,$0-12
	MOVW	ptr+0(FP), R4
	MOVW	delta+4(FP), R5
	DBAR	$0
xadd_loop:
	LL	(R4), R6
	ADD	R5, R6, R7
	SC	R7, (R4)
	BEQ	R7, xadd_loop
	DBAR	$0
	ADD	R5, R6, R4
	MOVW	R4, ret+8(FP)
	RET

TEXT ·Xchg(SB),NOSPLIT,$0-12
	MOVW	ptr+0(FP), R4
	MOVW	value+4(FP), R5
	DBAR	$0
xchg_loop:
	LL	(R4), R6
	MOVW	R5, R7
	SC	R7, (R4)
	BEQ	R7, xchg_loop
	DBAR	$0
	MOVW	R6, ret+8(FP)
	RET

TEXT ·Load(SB),NOSPLIT,$0-8
	MOVW	ptr+0(FP), R4
	DBAR	$0
	MOVW	(R4), R4
	DBAR	$0
	MOVW	R4, ret+4(FP)
	RET

TEXT ·Load8(SB),NOSPLIT,$0-5
	MOVW	ptr+0(FP), R4
	DBAR	$0
	MOVBU	(R4), R4
	DBAR	$0
	MOVB	R4, ret+4(FP)
	RET

TEXT ·Store(SB),NOSPLIT,$0-8
	MOVW	ptr+0(FP), R4
	MOVW	value+4(FP), R5
	DBAR	$0
	MOVW	R5, (R4)
	DBAR	$0
	RET

TEXT ·Store8(SB),NOSPLIT,$0-5
	MOVW	ptr+0(FP), R4
	MOVBU	value+4(FP), R5
	DBAR	$0
	MOVB	R5, (R4)
	DBAR	$0
	RET

TEXT ·Or(SB),NOSPLIT,$0-8
	MOVW	ptr+0(FP), R4
	MOVW	value+4(FP), R5
	DBAR	$0
or_loop:
	LL	(R4), R6
	OR	R5, R6, R7
	SC	R7, (R4)
	BEQ	R7, or_loop
	DBAR	$0
	RET

TEXT ·And(SB),NOSPLIT,$0-8
	MOVW	ptr+0(FP), R4
	MOVW	value+4(FP), R5
	DBAR	$0
and_loop:
	LL	(R4), R6
	AND	R5, R6, R7
	SC	R7, (R4)
	BEQ	R7, and_loop
	DBAR	$0
	RET

TEXT ·Or32(SB),NOSPLIT,$0-12
	MOVW	ptr+0(FP), R4
	MOVW	value+4(FP), R5
	DBAR	$0
or32_loop:
	LL	(R4), R6
	OR	R5, R6, R7
	SC	R7, (R4)
	BEQ	R7, or32_loop
	DBAR	$0
	MOVW	R6, ret+8(FP)
	RET

TEXT ·And32(SB),NOSPLIT,$0-12
	MOVW	ptr+0(FP), R4
	MOVW	value+4(FP), R5
	DBAR	$0
and32_loop:
	LL	(R4), R6
	AND	R5, R6, R7
	SC	R7, (R4)
	BEQ	R7, and32_loop
	DBAR	$0
	MOVW	R6, ret+8(FP)
	RET

TEXT ·spinLock(SB),NOSPLIT,$0-4
	MOVW	state+0(FP), R4
	DBAR	$0
lock_loop:
	LL	(R4), R5
	BNE	R5, lock_loop
	MOVW	$1, R6
	SC	R6, (R4)
	BEQ	R6, lock_loop
	DBAR	$0
	RET

TEXT ·spinUnlock(SB),NOSPLIT,$0-4
	MOVW	state+0(FP), R4
	DBAR	$0
	MOVW	R0, (R4)
	DBAR	$0
	RET

TEXT ·Casint32(SB),NOSPLIT|NOFRAME,$0-13
	JMP	·Cas(SB)
TEXT ·Casint64(SB),NOSPLIT|NOFRAME,$0-21
	JMP	·Cas64(SB)
TEXT ·Casuintptr(SB),NOSPLIT|NOFRAME,$0-13
	JMP	·Cas(SB)
TEXT ·Casp1(SB),NOSPLIT|NOFRAME,$0-13
	JMP	·Cas(SB)
TEXT ·CasRel(SB),NOSPLIT|NOFRAME,$0-13
	JMP	·Cas(SB)
TEXT ·Loadint32(SB),NOSPLIT|NOFRAME,$0-8
	JMP	·Load(SB)
TEXT ·Loadint64(SB),NOSPLIT|NOFRAME,$0-12
	JMP	·Load64(SB)
TEXT ·Loaduintptr(SB),NOSPLIT|NOFRAME,$0-8
	JMP	·Load(SB)
TEXT ·Loaduint(SB),NOSPLIT|NOFRAME,$0-8
	JMP	·Load(SB)
TEXT ·Loadp(SB),NOSPLIT|NOFRAME,$0-8
	JMP	·Load(SB)
TEXT ·LoadAcq(SB),NOSPLIT|NOFRAME,$0-8
	JMP	·Load(SB)
TEXT ·LoadAcquintptr(SB),NOSPLIT|NOFRAME,$0-8
	JMP	·Load(SB)
TEXT ·Storeint32(SB),NOSPLIT|NOFRAME,$0-8
	JMP	·Store(SB)
TEXT ·Storeint64(SB),NOSPLIT|NOFRAME,$0-12
	JMP	·Store64(SB)
TEXT ·Storeuintptr(SB),NOSPLIT|NOFRAME,$0-8
	JMP	·Store(SB)
TEXT ·StorepNoWB(SB),NOSPLIT|NOFRAME,$0-8
	JMP	·Store(SB)
TEXT ·StoreRel(SB),NOSPLIT|NOFRAME,$0-8
	JMP	·Store(SB)
TEXT ·StoreReluintptr(SB),NOSPLIT|NOFRAME,$0-8
	JMP	·Store(SB)
TEXT ·Xaddint32(SB),NOSPLIT|NOFRAME,$0-12
	JMP	·Xadd(SB)
TEXT ·Xaddint64(SB),NOSPLIT|NOFRAME,$0-20
	JMP	·Xadd64(SB)
TEXT ·Xadduintptr(SB),NOSPLIT|NOFRAME,$0-12
	JMP	·Xadd(SB)
TEXT ·Xchgint32(SB),NOSPLIT|NOFRAME,$0-12
	JMP	·Xchg(SB)
TEXT ·Xchgint64(SB),NOSPLIT|NOFRAME,$0-20
	JMP	·Xchg64(SB)
TEXT ·Xchguintptr(SB),NOSPLIT|NOFRAME,$0-12
	JMP	·Xchg(SB)
TEXT ·Anduintptr(SB),NOSPLIT|NOFRAME,$0-12
	JMP	·And32(SB)
TEXT ·Oruintptr(SB),NOSPLIT|NOFRAME,$0-12
	JMP	·Or32(SB)
