// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

TEXT errors(SB),$0
	JIRL	$2, R4, R1		// ERROR "JIRL offset must be a signed, 4-byte-aligned 18-bit value"
	JIRL	$131072, R4, R1		// ERROR "JIRL offset must be a signed, 4-byte-aligned 18-bit value"
	JIRL	$-131076, R4, R1	// ERROR "JIRL offset must be a signed, 4-byte-aligned 18-bit value"
	LU12IW	$1048576, R4		// ERROR "LA32R 20-bit immediate out of range"
	LU12IW	$-524289, R4		// ERROR "LA32R 20-bit immediate out of range"
	PCADDU12I	$1048576, R4	// ERROR "LA32R 20-bit immediate out of range"
	LL	1(R5), R4		// ERROR "LL.W offset must be a multiple of four"
	SC	R4, 1(R5)		// ERROR "SC.W offset must be a multiple of four"
	PRELD	0(R5), $32		// ERROR "illegal combination PRELD"
