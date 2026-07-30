// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The encodings in this file were produced independently with the GNU LA32R
// assembler and checked against the LA32R reference manual. They intentionally
// exercise LA32R instructions and operand forms rather than LA64 test cases.

#include "../../../../../runtime/textflag.h"

TEXT asmtest(SB),DUPOK|NOSPLIT,$0
	// Three-register integer instructions.
	ADD	R4, R5, R6		// a6101000
	SUB	R4, R5, R6		// a6101100
	SLT	R4, R5, R6		// a6101200
	SLTU	R4, R5, R6		// a6901200
	AND	R4, R5, R6		// a6901400
	OR	R4, R5, R6		// a6101500
	NOR	R4, R5, R6		// a6101400
	XOR	R4, R5, R6		// a6901500
	MUL	R4, R5, R6		// a6101c00
	MULH	R4, R5, R6		// a6901c00
	MULHU	R4, R5, R6		// a6101d00
	DIV	R4, R5, R6		// a6102000
	DIVU	R4, R5, R6		// a6102100
	REM	R4, R5, R6		// a6902000
	REMU	R4, R5, R6		// a6902100
	SLL	R4, R5, R6		// a6101700
	SRL	R4, R5, R6		// a6901700
	SRA	R4, R5, R6		// a6101800

	// Immediate integer and upper-immediate instructions.
	ADD	$-7, R5, R6		// a6e4bf02
	SGT	$-7, R5, R6		// a6e43f02
	SGTU	$-7, R5, R6		// a6e47f02
	AND	$2748, R5, R6		// a6f06a03
	OR	$2748, R5, R6		// a6f0aa03
	XOR	$2748, R5, R6		// a6f0ea03
	SLL	$7, R5, R6		// a69c4000
	SRL	$7, R5, R6		// a69c4400
	SRA	$7, R5, R6		// a69c4800
	LU12IW	$74565, R6		// a6682414
	PCADDU12I	$74565, R6	// a668241c
	MOVW	$2147483648, R6	// 06000015
	MOVW	$4294967295, R7	// e7ffff15e7fcbf03

	// Integer loads, stores, and the complete LA32R atomic subset.
	MOVB	-16(R5), R6		// a6c03f28
	MOVBU	-16(R5), R6		// a6c03f2a
	MOVH	-16(R5), R6		// a6c07f28
	MOVHU	-16(R5), R6		// a6c07f2a
	MOVW	-16(R5), R6		// a6c0bf28
	MOVB	R6, -16(R5)		// a6c03f29
	MOVH	R6, -16(R5)		// a6c07f29
	MOVW	R6, -16(R5)		// a6c0bf29
	LL	16(R5), R6		// a6100020
	SC	R6, 16(R5)		// a6100021
	LL	32764(R5), R6		// a6fc7f20
	LL	-32768(R5), R6		// a6008020
	SC	R6, 32764(R5)		// a6fc7f21
	SC	R6, -32768(R5)		// a6008021
	PRELD	-16(R5), $7		// a7c0ff2a

	// Barriers, traps, timer access, and indirect control transfer.
	DBAR	$18			// 12007238
	IBAR	$0			// 00807238
	SYSCALL	$291			// 23012b00
	BREAK	$291			// 23012a00
	RDTIMELW	R4, R5		// a4600000
	RDTIMEHW	R4, R5		// a4640000
	RDTIMEID	R4		// 80600000
	JIRL	$16, R4, R1		// 8110004c

	// LA32R uses two-register conditional branches.
	BEQ	R4, R5, 1(PC)		// 85040058
	BNE	R4, R5, 1(PC)		// 8504005c
	BLT	R4, R5, 1(PC)		// 85040060
	BGE	R4, R5, 1(PC)		// 85040064
	BLTU	R4, R5, 1(PC)		// 85040068
	BGEU	R4, R5, 1(PC)		// 8504006c
	WORD	$0			// 00000000
	JMP	1(PC)			// 00040050
	JAL	1(PC)			// CALL 1(PC) // 00040054

	// Optional base floating-point extension. The Go baseline remains soft-float,
	// but the assembler still accepts the architectural instructions.
	ADDF	F4, F5, F6		// a6900001
	ADDD	F4, F5, F6		// a6100101
	SUBF	F4, F5, F6		// a6900201
	MULF	F4, F5, F6		// a6900401
	DIVF	F4, F5, F6		// a6900601
	FMADDF	F4, F5, F7, F6		// a6901308
	ABSF	F5, F6			// a6041401
	NEGF	F5, F6			// a6141401
	SQRTF	F5, F6			// a6441401
	FRECIPF	F5, F6			// a6541401
	FRSQRTF	F5, F6			// a6641401
	FMAXF	F4, F5, F6		// a6900801
	FMINF	F4, F5, F6		// a6900a01
	CMPEQF	F4, F5, FCC3		// a310120c
	FSEL	FCC3, F5, F4, F6	// a690010d
	MOVF	-16(R5), F6		// a6c03f2b
	MOVF	F6, -16(R5)		// a6c07f2b
	MOVW	R5, F6			// a6a41401
	MOVW	F5, R6			// a6b41401
	TRUNCFW	F5, F6			// a6841a01
	MOVFD	F5, F6			// a6241901
	MOVDF	F5, F6			// a6181901
	BFPF	FCC3, 1(PC)		// 60040048
	BFPT	FCC3, 1(PC)		// 60050048

	WORD	$305419896		// 78563412
