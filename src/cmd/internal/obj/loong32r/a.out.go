// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file describes the LA32R register and instruction name spaces.
package loong32r

import "cmd/internal/obj"

//go:generate go run ../stringer.go -i $GOFILE -o anames.go -p loong32r

const (
	NSNAME = 8
	NSYM   = 50
	NREG   = 32
	NFREG  = 32
)

const (
	REG_R0 = obj.RBaseLOONG32R + iota // each encoded register bank starts at a multiple of 32
	REG_R1
	REG_R2
	REG_R3
	REG_R4
	REG_R5
	REG_R6
	REG_R7
	REG_R8
	REG_R9
	REG_R10
	REG_R11
	REG_R12
	REG_R13
	REG_R14
	REG_R15
	REG_R16
	REG_R17
	REG_R18
	REG_R19
	REG_R20
	REG_R21
	REG_R22
	REG_R23
	REG_R24
	REG_R25
	REG_R26
	REG_R27
	REG_R28
	REG_R29
	REG_R30
	REG_R31

	REG_F0
	REG_F1
	REG_F2
	REG_F3
	REG_F4
	REG_F5
	REG_F6
	REG_F7
	REG_F8
	REG_F9
	REG_F10
	REG_F11
	REG_F12
	REG_F13
	REG_F14
	REG_F15
	REG_F16
	REG_F17
	REG_F18
	REG_F19
	REG_F20
	REG_F21
	REG_F22
	REG_F23
	REG_F24
	REG_F25
	REG_F26
	REG_F27
	REG_F28
	REG_F29
	REG_F30
	REG_F31

	REG_FCSR0
	REG_FCSR1
	REG_FCSR2
	REG_FCSR3
	regFCSR4
	regFCSR5
	regFCSR6
	regFCSR7
	regFCSR8
	regFCSR9
	regFCSR10
	regFCSR11
	regFCSR12
	regFCSR13
	regFCSR14
	regFCSR15
	regFCSR16
	regFCSR17
	regFCSR18
	regFCSR19
	regFCSR20
	regFCSR21
	regFCSR22
	regFCSR23
	regFCSR24
	regFCSR25
	regFCSR26
	regFCSR27
	regFCSR28
	regFCSR29
	regFCSR30
	regFCSR31

	REG_FCC0
	REG_FCC1
	REG_FCC2
	REG_FCC3
	REG_FCC4
	REG_FCC5
	REG_FCC6
	REG_FCC7

	REG_LAST = REG_FCC7

	REG_SPECIAL = REG_FCSR0

	REGZERO = REG_R0
	REGLINK = REG_R1
	REGSP   = REG_R3
	REGCTXT = REG_R29
	REGG    = REG_R22
	REGTMP  = REG_R30
	FREGRET = REG_F0
)

var LOONG32RDWARFRegisters = map[int16]int16{}

func init() {
	for r := int16(REG_R0); r <= REG_R31; r++ {
		LOONG32RDWARFRegisters[r] = r - REG_R0
	}
	for r := int16(REG_F0); r <= REG_F31; r++ {
		LOONG32RDWARFRegisters[r] = r - REG_F0 + 32
	}
}

const (
	BIG_12 = 1 << 11
	BIG_32 = 1 << 31
)

const (
	LABEL  = 1 << 0
	LEAF   = 1 << 1
	SYNC   = 1 << 2
	BRANCH = 1 << 3
)

//go:generate go run ../mkcnames.go -i a.out.go -o cnames.go -p loong32r
const (
	C_NONE = iota
	C_REG
	C_FREG
	C_FCSRREG
	C_FCCREG

	C_ZCON
	C_U5CON
	C_U12CON
	C_S12CON
	C_12CON
	C_U15CON
	C_32CON20_0
	C_32CON

	C_SACON
	C_LACON
	C_EXTADDR
	C_BRAN
	C_SAUTO
	C_SAUTO_14
	C_LAUTO
	C_ZOREG
	C_SOREG_12
	C_SOREG_14
	C_LOREG_32
	C_ADDR
	C_TLS_LE
	C_TLS_IE
	C_GOTADDR
	C_TEXTSIZE

	C_GOK
	C_NCLASS
)

// The integer instructions below are exactly the LA32R base integer ISA from
// chapter 2 of the LA32R reference manual. A few Plan 9 pseudo-instructions
// (MOV*, NEG, SGT*, zero-relative branches, JMP/JAL/RET and WORD) lower only to
// those base instructions.
const (
	AADD = obj.ABaseLoong32r + obj.A_ARCHSPECIFIC + iota
	ASUB
	ASLT
	ASLTU
	ASGT  // pseudo: SLT with the two inputs exchanged
	ASGTU // pseudo: SLTU with the two inputs exchanged
	ALU12IW
	APCADDU12I
	AAND
	AOR
	ANOR
	AXOR
	AMUL
	AMULH
	AMULHU
	ADIV
	ADIVU
	AREM
	AREMU
	ASLL
	ASRL
	ASRA

	ABEQ
	ABNE
	ABLT
	ABGE
	ABLTU
	ABGEU
	ABLTZ // pseudo: BLT rj, R0
	ABGEZ // pseudo: BGE rj, R0
	ABLEZ // pseudo: BGE R0, rj
	ABGTZ // pseudo: BLT R0, rj
	AJIRL

	AMOVB
	AMOVBU
	AMOVH
	AMOVHU
	AMOVW
	APRELD
	ALL
	ASC
	ADBAR
	AIBAR
	ASYSCALL
	ABREAK
	ARDTIMELW
	ARDTIMEHW
	ARDTIMEID
	ANOOP
	AWORD

	// Optional base floating-point extension from chapter 3. The GOARCH
	// baseline remains soft-float, so generated Go code never requires it.
	AADDF
	AADDD
	ASUBF
	ASUBD
	AMULF
	AMULD
	ADIVF
	ADIVD
	AFMADDF
	AFMADDD
	AFMSUBF
	AFMSUBD
	AFNMADDF
	AFNMADDD
	AFNMSUBF
	AFNMSUBD
	AFMAXF
	AFMAXD
	AFMINF
	AFMIND
	AFMAXAF
	AFMAXAD
	AFMINAF
	AFMINAD
	AABSF
	AABSD
	ANEGF
	ANEGD
	ASQRTF
	ASQRTD
	AFRECIPF
	AFRECIPD
	AFRSQRTF
	AFRSQRTD
	AFCOPYSGF
	AFCOPYSGD
	AFCLASSF
	AFCLASSD
	ACMPEQF
	ACMPEQD
	ACMPGEF
	ACMPGED
	ACMPGTF
	ACMPGTD
	AMOVF
	AMOVD
	AMOVWF
	AMOVWD
	ATRUNCFW
	ATRUNCDW
	AMOVFD
	AMOVDF
	AFSEL
	ABFPF // BCEQZ
	ABFPT // BCNEZ

	ALAST
)

// Aliases used by compiler code and assembly listings.
const (
	AJMP  = obj.AJMP
	AJAL  = obj.ACALL
	ARET  = obj.ARET
	AJALR = AJIRL
)

func init() {
	obj.RegisterOpcode(obj.ABaseLoong32r, Anames)
}
