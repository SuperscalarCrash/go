// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package loong32r implements the Go assembler backend for the independent
// LoongArch 32-bit Reduced (LA32R) instruction set.
//
// The integer register names are R0 through R31. R0 is the architectural zero
// register, R1 is the link register, R3 is SP, R22 holds g, R29 is the closure
// context register, and R30 is reserved as the assembler temporary.
//
// Integer mnemonics use Go's width-neutral names because every general-purpose
// register and integer operation is 32 bits. For example,
//
//	ADD R5, R4, R6       // add.w r6, r4, r5
//	MOVW 8(R4), R5      // ld.w r5, r4, 8
//	MOVW R5, 8(R4)      // st.w r5, r4, 8
//	SLL $3, R4, R5      // slli.w r5, r4, 3
//
// MOVB and MOVH loads are signed; MOVBU and MOVHU loads are unsigned. Stores
// select st.b, st.h, or st.w from the mnemonic. LL and SC expose the only
// atomic primitives in the LA32R base ISA.
//
// LA32R's optional chapter-3 floating-point instructions are accepted by the
// assembler, but linux/loong32r is compiled in soft-float mode and therefore
// does not require an FPU at run time.
//
// This package deliberately contains no LA64 integer, LSX, LASX, AMO, CRC, or
// other LoongArch extension instructions.
package loong32r
