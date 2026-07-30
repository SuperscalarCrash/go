// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build loong32r

.file "gcc_loong32r.S"

/*
 * void crosscall1(void (*fn)(void), void (*setg_gcc)(void *g), void *g)
 *
 * Enter the Go toolchain from the LA32 ILP32S C ABI. Go code treats its
 * allocatable registers as caller-save, while the C ABI requires r22-r31 to
 * survive calls. The soft-float ABI has no callee-save floating registers.
 */
.globl crosscall1
.type crosscall1, @function
crosscall1:
	addi.w	$r3, $r3, -48
	st.w	$r1,  $r3, 0
	st.w	$r22, $r3, 4
	st.w	$r23, $r3, 8
	st.w	$r24, $r3, 12
	st.w	$r25, $r3, 16
	st.w	$r26, $r3, 20
	st.w	$r27, $r3, 24
	st.w	$r28, $r3, 28
	st.w	$r29, $r3, 32
	st.w	$r30, $r3, 36
	st.w	$r31, $r3, 40

	move	$r23, $r4
	move	$r4, $r6
	jirl	$r1, $r5, 0
	jirl	$r1, $r23, 0

	ld.w	$r22, $r3, 4
	ld.w	$r23, $r3, 8
	ld.w	$r24, $r3, 12
	ld.w	$r25, $r3, 16
	ld.w	$r26, $r3, 20
	ld.w	$r27, $r3, 24
	ld.w	$r28, $r3, 28
	ld.w	$r29, $r3, 32
	ld.w	$r30, $r3, 36
	ld.w	$r31, $r3, 40
	ld.w	$r1,  $r3, 0
	addi.w	$r3, $r3, 48
	jirl	$r0, $r1, 0
.size crosscall1, .-crosscall1

#ifdef __ELF__
.section .note.GNU-stack,"",%progbits
#endif
