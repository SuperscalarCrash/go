// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package loong32r

import (
	"cmd/internal/obj"
	"cmd/internal/objabi"
	"fmt"
	"log"
	"slices"
)

// ctxt0 holds the state for one independently assembled LA32R function.
type ctxt0 struct {
	ctxt       *obj.Link
	newprog    obj.ProgAlloc
	cursym     *obj.LSym
	autosize   int32
	instoffset int64
	pc         int64
}

const (
	FuncAlign = 4
	loopAlign = 16
)

type Optab struct {
	as    obj.As
	from1 uint8
	reg   uint8
	from3 uint8
	to1   uint8
	to2   uint8
	type_ int8
	size  int8
	param int16
	flag  uint8
}

const (
	NOTUSETMP = 1 << iota
	branchLoopHead
)

// optab contains only LA32R encodings. Aliases with identical operand forms
// are installed by buildop.
var optab = []Optab{
	{obj.ATEXT, C_ADDR, C_NONE, C_NONE, C_TEXTSIZE, C_NONE, 0, 0, 0, 0},

	// General-register moves and integer ALU operations.
	{AMOVW, C_REG, C_NONE, C_NONE, C_REG, C_NONE, 1, 4, 0, 0},
	{AMOVB, C_REG, C_NONE, C_NONE, C_REG, C_NONE, 12, 8, 0, 0},
	{AMOVBU, C_REG, C_NONE, C_NONE, C_REG, C_NONE, 12, 4, 0, 0},
	{AMOVHU, C_REG, C_NONE, C_NONE, C_REG, C_NONE, 12, 8, 0, 0},

	{AADD, C_REG, C_REG, C_NONE, C_REG, C_NONE, 2, 4, 0, 0},
	{AADD, C_REG, C_NONE, C_NONE, C_REG, C_NONE, 2, 4, 0, 0},
	{AADD, C_S12CON, C_REG, C_NONE, C_REG, C_NONE, 4, 4, 0, 0},
	{AADD, C_S12CON, C_NONE, C_NONE, C_REG, C_NONE, 4, 4, 0, 0},
	{AADD, C_U12CON, C_REG, C_NONE, C_REG, C_NONE, 10, 8, 0, 0},
	{AADD, C_U12CON, C_NONE, C_NONE, C_REG, C_NONE, 10, 8, 0, 0},
	{AADD, C_32CON, C_REG, C_NONE, C_REG, C_NONE, 24, 12, 0, 0},
	{AADD, C_32CON, C_NONE, C_NONE, C_REG, C_NONE, 24, 12, 0, 0},

	{ASUB, C_REG, C_REG, C_NONE, C_REG, C_NONE, 2, 4, 0, 0},
	{ASUB, C_REG, C_NONE, C_NONE, C_REG, C_NONE, 2, 4, 0, 0},

	{ASGT, C_REG, C_REG, C_NONE, C_REG, C_NONE, 2, 4, 0, 0},
	{ASGT, C_REG, C_NONE, C_NONE, C_REG, C_NONE, 2, 4, 0, 0},
	{ASGT, C_S12CON, C_REG, C_NONE, C_REG, C_NONE, 4, 4, 0, 0},
	{ASGT, C_S12CON, C_NONE, C_NONE, C_REG, C_NONE, 4, 4, 0, 0},
	{ASGT, C_32CON, C_REG, C_NONE, C_REG, C_NONE, 24, 12, 0, 0},
	{ASGT, C_32CON, C_NONE, C_NONE, C_REG, C_NONE, 24, 12, 0, 0},

	{AAND, C_REG, C_REG, C_NONE, C_REG, C_NONE, 2, 4, 0, 0},
	{AAND, C_REG, C_NONE, C_NONE, C_REG, C_NONE, 2, 4, 0, 0},
	{AAND, C_U12CON, C_REG, C_NONE, C_REG, C_NONE, 4, 4, 0, 0},
	{AAND, C_U12CON, C_NONE, C_NONE, C_REG, C_NONE, 4, 4, 0, 0},
	{AAND, C_32CON, C_REG, C_NONE, C_REG, C_NONE, 24, 12, 0, 0},
	{AAND, C_32CON, C_NONE, C_NONE, C_REG, C_NONE, 24, 12, 0, 0},

	{ANOR, C_REG, C_REG, C_NONE, C_REG, C_NONE, 2, 4, 0, 0},
	{ANOR, C_REG, C_NONE, C_NONE, C_REG, C_NONE, 2, 4, 0, 0},
	{ANOR, C_32CON, C_REG, C_NONE, C_REG, C_NONE, 24, 12, 0, 0},
	{ANOR, C_32CON, C_NONE, C_NONE, C_REG, C_NONE, 24, 12, 0, 0},

	{ASLL, C_REG, C_REG, C_NONE, C_REG, C_NONE, 2, 4, 0, 0},
	{ASLL, C_REG, C_NONE, C_NONE, C_REG, C_NONE, 2, 4, 0, 0},
	{ASLL, C_U5CON, C_REG, C_NONE, C_REG, C_NONE, 16, 4, 0, 0},
	{ASLL, C_U5CON, C_NONE, C_NONE, C_REG, C_NONE, 16, 4, 0, 0},

	// Materialize constants and addresses.
	{ALU12IW, C_32CON, C_NONE, C_NONE, C_REG, C_NONE, 66, 4, 0, NOTUSETMP},
	{APCADDU12I, C_32CON, C_NONE, C_NONE, C_REG, C_NONE, 66, 4, 0, NOTUSETMP},
	{AMOVW, C_SACON, C_NONE, C_NONE, C_REG, C_NONE, 3, 4, REGSP, 0},
	{AMOVW, C_LACON, C_NONE, C_NONE, C_REG, C_NONE, 27, 12, REGSP, 0},
	{AMOVW, C_12CON, C_NONE, C_NONE, C_REG, C_NONE, 3, 4, REGZERO, 0},
	{AMOVW, C_32CON20_0, C_NONE, C_NONE, C_REG, C_NONE, 25, 4, 0, NOTUSETMP},
	{AMOVW, C_32CON, C_NONE, C_NONE, C_REG, C_NONE, 19, 8, 0, NOTUSETMP},
	{AMOVW, C_EXTADDR, C_NONE, C_NONE, C_REG, C_NONE, 52, 8, 0, NOTUSETMP},
	{AMOVW, C_GOTADDR, C_NONE, C_NONE, C_REG, C_NONE, 65, 8, 0, NOTUSETMP},

	// Integer loads and stores.
	{AMOVW, C_REG, C_NONE, C_NONE, C_SAUTO, C_NONE, 7, 4, REGSP, 0},
	{AMOVW, C_REG, C_NONE, C_NONE, C_SOREG_12, C_NONE, 7, 4, REGZERO, 0},
	{AMOVW, C_REG, C_NONE, C_NONE, C_LAUTO, C_NONE, 35, 12, REGSP, 0},
	{AMOVW, C_REG, C_NONE, C_NONE, C_LOREG_32, C_NONE, 35, 12, REGZERO, 0},
	{AMOVW, C_REG, C_NONE, C_NONE, C_ADDR, C_NONE, 50, 8, 0, 0},
	{AMOVW, C_REG, C_NONE, C_NONE, C_TLS_LE, C_NONE, 53, 16, 0, 0},
	{AMOVW, C_REG, C_NONE, C_NONE, C_TLS_IE, C_NONE, 56, 16, 0, 0},

	{AMOVW, C_SAUTO, C_NONE, C_NONE, C_REG, C_NONE, 8, 4, REGSP, 0},
	{AMOVW, C_SOREG_12, C_NONE, C_NONE, C_REG, C_NONE, 8, 4, REGZERO, 0},
	{AMOVW, C_LAUTO, C_NONE, C_NONE, C_REG, C_NONE, 36, 12, REGSP, 0},
	{AMOVW, C_LOREG_32, C_NONE, C_NONE, C_REG, C_NONE, 36, 12, REGZERO, 0},
	{AMOVW, C_ADDR, C_NONE, C_NONE, C_REG, C_NONE, 51, 8, 0, 0},
	{AMOVW, C_TLS_LE, C_NONE, C_NONE, C_REG, C_NONE, 54, 16, 0, 0},
	{AMOVW, C_TLS_IE, C_NONE, C_NONE, C_REG, C_NONE, 57, 16, 0, 0},

	{AMOVB, C_REG, C_NONE, C_NONE, C_SAUTO, C_NONE, 7, 4, REGSP, 0},
	{AMOVB, C_REG, C_NONE, C_NONE, C_SOREG_12, C_NONE, 7, 4, REGZERO, 0},
	{AMOVB, C_REG, C_NONE, C_NONE, C_LAUTO, C_NONE, 35, 12, REGSP, 0},
	{AMOVB, C_REG, C_NONE, C_NONE, C_LOREG_32, C_NONE, 35, 12, REGZERO, 0},
	{AMOVB, C_REG, C_NONE, C_NONE, C_ADDR, C_NONE, 50, 8, 0, 0},
	{AMOVB, C_REG, C_NONE, C_NONE, C_TLS_LE, C_NONE, 53, 16, 0, 0},
	{AMOVB, C_REG, C_NONE, C_NONE, C_TLS_IE, C_NONE, 56, 16, 0, 0},
	{AMOVB, C_SAUTO, C_NONE, C_NONE, C_REG, C_NONE, 8, 4, REGSP, 0},
	{AMOVB, C_SOREG_12, C_NONE, C_NONE, C_REG, C_NONE, 8, 4, REGZERO, 0},
	{AMOVB, C_LAUTO, C_NONE, C_NONE, C_REG, C_NONE, 36, 12, REGSP, 0},
	{AMOVB, C_LOREG_32, C_NONE, C_NONE, C_REG, C_NONE, 36, 12, REGZERO, 0},
	{AMOVB, C_ADDR, C_NONE, C_NONE, C_REG, C_NONE, 51, 8, 0, 0},
	{AMOVB, C_TLS_LE, C_NONE, C_NONE, C_REG, C_NONE, 54, 16, 0, 0},
	{AMOVB, C_TLS_IE, C_NONE, C_NONE, C_REG, C_NONE, 57, 16, 0, 0},

	{AMOVBU, C_SAUTO, C_NONE, C_NONE, C_REG, C_NONE, 8, 4, REGSP, 0},
	{AMOVBU, C_SOREG_12, C_NONE, C_NONE, C_REG, C_NONE, 8, 4, REGZERO, 0},
	{AMOVBU, C_LAUTO, C_NONE, C_NONE, C_REG, C_NONE, 36, 12, REGSP, 0},
	{AMOVBU, C_LOREG_32, C_NONE, C_NONE, C_REG, C_NONE, 36, 12, REGZERO, 0},
	{AMOVBU, C_ADDR, C_NONE, C_NONE, C_REG, C_NONE, 51, 8, 0, 0},
	{AMOVBU, C_TLS_LE, C_NONE, C_NONE, C_REG, C_NONE, 54, 16, 0, 0},
	{AMOVBU, C_TLS_IE, C_NONE, C_NONE, C_REG, C_NONE, 57, 16, 0, 0},

	{AMOVHU, C_SAUTO, C_NONE, C_NONE, C_REG, C_NONE, 8, 4, REGSP, 0},
	{AMOVHU, C_SOREG_12, C_NONE, C_NONE, C_REG, C_NONE, 8, 4, REGZERO, 0},
	{AMOVHU, C_LAUTO, C_NONE, C_NONE, C_REG, C_NONE, 36, 12, REGSP, 0},
	{AMOVHU, C_LOREG_32, C_NONE, C_NONE, C_REG, C_NONE, 36, 12, REGZERO, 0},
	{AMOVHU, C_ADDR, C_NONE, C_NONE, C_REG, C_NONE, 51, 8, 0, 0},
	{AMOVHU, C_TLS_LE, C_NONE, C_NONE, C_REG, C_NONE, 54, 16, 0, 0},
	{AMOVHU, C_TLS_IE, C_NONE, C_NONE, C_REG, C_NONE, 57, 16, 0, 0},

	// Control flow. Integer conditional branches always use the LA32R
	// two-register 16-bit form; LA64-only BEQZ/BNEZ forms are not used.
	{ABEQ, C_REG, C_REG, C_NONE, C_BRAN, C_NONE, 6, 4, 0, 0},
	{ABEQ, C_REG, C_NONE, C_NONE, C_BRAN, C_NONE, 6, 4, 0, 0},
	{AJMP, C_NONE, C_NONE, C_NONE, C_BRAN, C_NONE, 11, 4, 0, 0},
	{AJAL, C_NONE, C_NONE, C_NONE, C_BRAN, C_NONE, 11, 4, 0, 0},
	{AJMP, C_NONE, C_NONE, C_NONE, C_ZOREG, C_NONE, 18, 4, REGZERO, 0},
	{AJAL, C_NONE, C_NONE, C_NONE, C_ZOREG, C_NONE, 18, 4, REGLINK, 0},
	{AJIRL, C_32CON, C_REG, C_NONE, C_REG, C_NONE, 67, 4, 0, 0},

	// LL.W and SC.W are LA32R's complete base atomic instruction set.
	{ALL, C_SOREG_14, C_NONE, C_NONE, C_REG, C_NONE, 40, 4, 0, 0},
	{ALL, C_SAUTO_14, C_NONE, C_NONE, C_REG, C_NONE, 40, 4, REGSP, 0},
	{ASC, C_REG, C_NONE, C_NONE, C_SOREG_14, C_NONE, 40, 4, 0, 0},
	{ASC, C_REG, C_NONE, C_NONE, C_SAUTO_14, C_NONE, 40, 4, REGSP, 0},

	{ASYSCALL, C_NONE, C_NONE, C_NONE, C_NONE, C_NONE, 5, 4, 0, 0},
	{ASYSCALL, C_U15CON, C_NONE, C_NONE, C_NONE, C_NONE, 5, 4, 0, 0},
	{APRELD, C_SOREG_12, C_U5CON, C_NONE, C_NONE, C_NONE, 47, 4, 0, 0},
	{ARDTIMELW, C_NONE, C_NONE, C_NONE, C_REG, C_REG, 62, 4, 0, 0},
	{ARDTIMEID, C_REG, C_NONE, C_NONE, C_NONE, C_NONE, 63, 4, 0, 0},

	// Optional base floating-point extension.
	{AADDF, C_FREG, C_FREG, C_NONE, C_FREG, C_NONE, 2, 4, 0, 0},
	{AADDF, C_FREG, C_NONE, C_NONE, C_FREG, C_NONE, 2, 4, 0, 0},
	{ACMPEQF, C_FREG, C_FREG, C_NONE, C_FCCREG, C_NONE, 2, 4, 0, 0},
	{AABSF, C_FREG, C_NONE, C_NONE, C_FREG, C_NONE, 9, 4, 0, 0},
	{AFMADDF, C_FREG, C_FREG, C_FREG, C_FREG, C_NONE, 37, 4, 0, 0},
	{AFMADDF, C_FREG, C_FREG, C_NONE, C_FREG, C_NONE, 37, 4, 0, 0},
	{AFSEL, C_FCCREG, C_FREG, C_FREG, C_FREG, C_NONE, 33, 4, 0, 0},
	{AFSEL, C_FCCREG, C_FREG, C_NONE, C_FREG, C_NONE, 33, 4, 0, 0},
	{ABFPT, C_NONE, C_NONE, C_NONE, C_BRAN, C_NONE, 6, 4, 0, 0},
	{ABFPT, C_FCCREG, C_NONE, C_NONE, C_BRAN, C_NONE, 6, 4, 0, 0},

	{AMOVF, C_FREG, C_NONE, C_NONE, C_FREG, C_NONE, 9, 4, 0, 0},
	{AMOVF, C_SAUTO, C_NONE, C_NONE, C_FREG, C_NONE, 28, 4, REGSP, 0},
	{AMOVF, C_SOREG_12, C_NONE, C_NONE, C_FREG, C_NONE, 28, 4, REGZERO, 0},
	{AMOVF, C_LAUTO, C_NONE, C_NONE, C_FREG, C_NONE, 28, 12, REGSP, 0},
	{AMOVF, C_LOREG_32, C_NONE, C_NONE, C_FREG, C_NONE, 28, 12, REGZERO, 0},
	{AMOVF, C_ADDR, C_NONE, C_NONE, C_FREG, C_NONE, 51, 8, 0, 0},
	{AMOVF, C_FREG, C_NONE, C_NONE, C_SAUTO, C_NONE, 29, 4, REGSP, 0},
	{AMOVF, C_FREG, C_NONE, C_NONE, C_SOREG_12, C_NONE, 29, 4, REGZERO, 0},
	{AMOVF, C_FREG, C_NONE, C_NONE, C_LAUTO, C_NONE, 29, 12, REGSP, 0},
	{AMOVF, C_FREG, C_NONE, C_NONE, C_LOREG_32, C_NONE, 29, 12, REGZERO, 0},
	{AMOVF, C_FREG, C_NONE, C_NONE, C_ADDR, C_NONE, 50, 8, 0, 0},
	{AMOVW, C_REG, C_NONE, C_NONE, C_FREG, C_NONE, 30, 4, 0, 0},
	{AMOVW, C_FREG, C_NONE, C_NONE, C_REG, C_NONE, 30, 4, 0, 0},
	{AMOVW, C_12CON, C_NONE, C_NONE, C_FREG, C_NONE, 34, 8, 0, 0},

	{AWORD, C_32CON, C_NONE, C_NONE, C_NONE, C_NONE, 38, 4, 0, 0},
	{ANOOP, C_NONE, C_NONE, C_NONE, C_NONE, C_NONE, 49, 4, 0, 0},

	{obj.APCALIGN, C_U12CON, C_NONE, C_NONE, C_NONE, C_NONE, 0, 0, 0, 0},
	{obj.APCDATA, C_32CON, C_NONE, C_NONE, C_32CON, C_NONE, 0, 0, 0, 0},
	{obj.AFUNCDATA, C_U12CON, C_NONE, C_NONE, C_ADDR, C_NONE, 0, 0, 0, 0},
	{obj.ANOP, C_NONE, C_NONE, C_NONE, C_NONE, C_NONE, 0, 0, 0, 0},
	{obj.ANOP, C_32CON, C_NONE, C_NONE, C_NONE, C_NONE, 0, 0, 0, 0},
	{obj.ANOP, C_REG, C_NONE, C_NONE, C_NONE, C_NONE, 0, 0, 0, 0},
	{obj.ANOP, C_FREG, C_NONE, C_NONE, C_NONE, C_NONE, 0, 0, 0, 0},
}

func pcAlignPadLength(ctxt *obj.Link, pc, align int64) int {
	if align&(align-1) != 0 || align < 8 || align > 2048 {
		ctxt.Diag("instruction alignment must be a power of two in [8, 2048], got %d", align)
		return 0
	}
	return int(-pc & (align - 1))
}

var oprange [ALAST & obj.AMask][]Optab
var xcmp [C_NCLASS][C_NCLASS]bool

func span0(ctxt *obj.Link, cursym *obj.LSym, newprog obj.ProgAlloc) {
	if ctxt.Retpoline {
		ctxt.Diag("-spectre=ret not supported on loong32r")
		ctxt.Retpoline = false
	}
	p := cursym.Func().Text
	if p == nil || p.Link == nil {
		return
	}
	c := ctxt0{ctxt: ctxt, newprog: newprog, cursym: cursym, autosize: int32(p.To.Offset + ctxt.Arch.FixedFrameSize)}
	if oprange[AOR&obj.AMask] == nil {
		c.ctxt.Diag("loong32r ops not initialized; call buildop first")
	}

	pc := int64(0)
	p.Pc = 0
	for p = p.Link; p != nil; p = p.Link {
		p.Pc = pc
		o := c.oplook(p)
		m := int(o.size)
		if m == 0 {
			switch p.As {
			case obj.APCALIGN:
				m = pcAlignPadLength(ctxt, pc, p.From.Offset)
				if int32(p.From.Offset) > cursym.Func().Align {
					cursym.Func().Align = int32(p.From.Offset)
				}
			case obj.ANOP, obj.AFUNCDATA, obj.APCDATA:
				continue
			default:
				c.ctxt.Diag("zero-width instruction %v", p)
			}
		}
		pc += int64(m)
	}
	c.cursym.Size = pc

	for p = c.cursym.Func().Text.Link; p != nil; p = p.Link {
		if q := p.To.Target(); q != nil && q.Pc < p.Pc {
			q.Mark |= branchLoopHead
		}
	}

	for {
		rescan := false
		pc = 0
		prev := c.cursym.Func().Text
		for p = prev.Link; p != nil; prev, p = p, p.Link {
			p.Pc = pc
			o := c.oplook(p)
			if p.Mark&branchLoopHead != 0 && pc&(loopAlign-1) != 0 &&
				!(prev.As == obj.APCALIGN && prev.From.Offset >= loopAlign) {
				q := c.newprog()
				prev.Link, q.Link = q, p
				q.Pc = pc
				q.As = obj.APCALIGN
				q.From.SetConst(loopAlign)
				pc += int64(pcAlignPadLength(ctxt, pc, loopAlign))
				p.Pc = pc
				rescan = true
			}

			if o.type_ == 6 && p.To.Target() != nil {
				bound := int64(1 << 17) // signed offs16, scaled by four
				if p.As == ABFPT || p.As == ABFPF {
					bound = 1 << 22 // signed offs21, scaled by four
				}
				delta := p.To.Target().Pc - pc
				if delta < -bound || delta >= bound {
					// Invert the condition and branch around an unconditional B.
					target := p.To.Target()
					jump := c.newprog()
					next := c.newprog()
					jump.As = AJMP
					jump.Pos = p.Pos
					jump.To.Type = obj.TYPE_BRANCH
					jump.To.SetTarget(target)
					next.As = obj.ANOP
					next.Pos = p.Pos
					jump.Link, next.Link = next, p.Link
					p.Link = jump
					p.As = invertBranch(p.As)
					p.To.SetTarget(next)
					rescan = true
				}
			}

			m := int(o.size)
			if m == 0 {
				switch p.As {
				case obj.APCALIGN:
					m = pcAlignPadLength(ctxt, pc, p.From.Offset)
				case obj.ANOP, obj.AFUNCDATA, obj.APCDATA:
					continue
				default:
					c.ctxt.Diag("zero-width instruction %v", p)
				}
			}
			pc += int64(m)
		}
		c.cursym.Size = pc
		if !rescan {
			break
		}
	}

	pc += -pc & (FuncAlign - 1)
	c.cursym.Size = pc
	c.cursym.Grow(pc)
	bp := c.cursym.P
	var out [4]uint32
	for p = c.cursym.Func().Text.Link; p != nil; p = p.Link {
		c.pc = p.Pc
		o := c.oplook(p)
		if p.As == obj.APCALIGN {
			for n := pcAlignPadLength(ctxt, p.Pc, p.From.Offset); n > 0; n -= 4 {
				ctxt.Arch.ByteOrder.PutUint32(bp, OP_12IRR(c.opirr(AAND), 0, 0, 0))
				bp = bp[4:]
			}
			continue
		}
		if int(o.size) > 4*len(out) {
			log.Fatalf("loong32r output buffer too small for %v", p)
		}
		c.asmout(p, o, out[:])
		for i := 0; i < int(o.size)/4; i++ {
			ctxt.Arch.ByteOrder.PutUint32(bp, out[i])
			bp = bp[4:]
		}
	}

	obj.MarkUnsafePoints(ctxt, cursym.Func().Text, newprog, c.isUnsafePoint, c.isRestartable)
	for _, jt := range cursym.Func().JumpTables {
		for i, p := range jt.Targets {
			jt.Sym.WriteAddr(ctxt, int64(i)*4, 4, cursym, p.Pc)
		}
	}
}

func invertBranch(as obj.As) obj.As {
	switch as {
	case ABEQ:
		return ABNE
	case ABNE:
		return ABEQ
	case ABLT:
		return ABGE
	case ABGE:
		return ABLT
	case ABLTU:
		return ABGEU
	case ABGEU:
		return ABLTU
	case ABLTZ:
		return ABGEZ
	case ABGEZ:
		return ABLTZ
	case ABLEZ:
		return ABGTZ
	case ABGTZ:
		return ABLEZ
	case ABFPT:
		return ABFPF
	case ABFPF:
		return ABFPT
	default:
		return as
	}
}

func (c *ctxt0) isUnsafePoint(p *obj.Prog) bool {
	return p.From.Reg == REGTMP || p.To.Reg == REGTMP || p.Reg == REGTMP
}

func (c *ctxt0) isRestartable(p *obj.Prog) bool {
	if c.isUnsafePoint(p) {
		return false
	}
	o := c.oplook(p)
	return o.size > 4 && o.flag&NOTUSETMP == 0
}

func isint32(v int64) bool { return int64(int32(v)) == v }

func (c *ctxt0) aclass(a *obj.Addr) int {
	switch a.Type {
	case obj.TYPE_NONE:
		return C_NONE
	case obj.TYPE_REG:
		return c.rclass(a.Reg)
	case obj.TYPE_MEM:
		switch a.Name {
		case obj.NAME_EXTERN, obj.NAME_STATIC:
			if a.Sym == nil {
				return C_GOK
			}
			c.instoffset = a.Offset
			if a.Sym.Type == objabi.STLSBSS {
				if c.ctxt.Flag_shared {
					return C_TLS_IE
				}
				return C_TLS_LE
			}
			return C_ADDR
		case obj.NAME_AUTO:
			if a.Reg == REGSP {
				a.Reg = obj.REG_NONE
			}
			c.instoffset = int64(c.autosize) + a.Offset
			if -BIG_12 <= c.instoffset && c.instoffset < BIG_12 {
				return C_SAUTO
			}
			if -1<<15 <= c.instoffset && c.instoffset < 1<<15 {
				return C_SAUTO_14
			}
			return C_LAUTO
		case obj.NAME_PARAM:
			if a.Reg == REGSP {
				a.Reg = obj.REG_NONE
			}
			c.instoffset = int64(c.autosize) + a.Offset + c.ctxt.Arch.FixedFrameSize
			if -BIG_12 <= c.instoffset && c.instoffset < BIG_12 {
				return C_SAUTO
			}
			if -1<<15 <= c.instoffset && c.instoffset < 1<<15 {
				return C_SAUTO_14
			}
			return C_LAUTO
		case obj.NAME_NONE:
			if a.Index != 0 {
				return C_GOK // LA32R has no base indexed load/store instruction
			}
			c.instoffset = a.Offset
			if c.instoffset == 0 {
				return C_ZOREG
			}
			if -BIG_12 <= c.instoffset && c.instoffset < BIG_12 {
				return C_SOREG_12
			}
			if -1<<15 <= c.instoffset && c.instoffset < 1<<15 {
				return C_SOREG_14
			}
			if isint32(c.instoffset) {
				return C_LOREG_32
			}
			return C_GOK
		case obj.NAME_GOTREF:
			return C_GOTADDR
		}
	case obj.TYPE_TEXTSIZE:
		return C_TEXTSIZE
	case obj.TYPE_CONST, obj.TYPE_ADDR:
		switch a.Name {
		case obj.NAME_NONE:
			c.instoffset = a.Offset
			if a.Reg != 0 {
				if -BIG_12 <= c.instoffset && c.instoffset < BIG_12 {
					return C_SACON
				}
				if isint32(c.instoffset) {
					return C_LACON
				}
				return C_GOK
			}
		case obj.NAME_EXTERN, obj.NAME_STATIC:
			if a.Sym == nil {
				return C_GOK
			}
			c.instoffset = a.Offset
			if a.Sym.Type == objabi.STLSBSS {
				c.ctxt.Diag("taking the address of a TLS variable is not supported")
			}
			return C_EXTADDR
		case obj.NAME_AUTO:
			if a.Reg == REGSP {
				a.Reg = obj.REG_NONE
			}
			c.instoffset = int64(c.autosize) + a.Offset
			if -BIG_12 <= c.instoffset && c.instoffset < BIG_12 {
				return C_SACON
			}
			return C_LACON
		case obj.NAME_PARAM:
			if a.Reg == REGSP {
				a.Reg = obj.REG_NONE
			}
			c.instoffset = int64(c.autosize) + a.Offset + c.ctxt.Arch.FixedFrameSize
			if -BIG_12 <= c.instoffset && c.instoffset < BIG_12 {
				return C_SACON
			}
			return C_LACON
		default:
			return C_GOK
		}
		v := c.instoffset
		if v == 0 {
			return C_ZCON
		}
		if 0 <= v && v <= 31 {
			return C_U5CON
		}
		if -1<<11 <= v && v < 1<<11 {
			return C_S12CON
		}
		if 0 <= v && v < 1<<12 {
			return C_U12CON
		}
		if 0 <= v && v < 1<<15 {
			return C_U15CON
		}
		// A word constant may be written either as a signed int32 or as its
		// unsigned uint32 spelling. Both forms describe the same LA32R
		// register bit pattern.
		if v < -1<<31 || v >= 1<<32 {
			c.ctxt.Diag("constant %d does not fit in an LA32R word", v)
		}
		if uint32(v)&0xfff == 0 {
			return C_32CON20_0
		}
		return C_32CON
	case obj.TYPE_BRANCH:
		return C_BRAN
	}
	return C_GOK
}

func (c *ctxt0) rclass(r int16) int {
	switch {
	case REG_R0 <= r && r <= REG_R31:
		return C_REG
	case REG_F0 <= r && r <= REG_F31:
		return C_FREG
	case REG_FCSR0 <= r && r <= REG_FCSR3:
		return C_FCSRREG
	case REG_FCC0 <= r && r <= REG_FCC7:
		return C_FCCREG
	default:
		return C_GOK
	}
}

func oclass(a *obj.Addr) int { return int(a.Class) - 1 }

func (c *ctxt0) regoff(a *obj.Addr) int32 {
	c.instoffset = 0
	c.aclass(a)
	return int32(c.instoffset)
}

func prasm(p *obj.Prog) { fmt.Printf("%v\n", p) }

func (c *ctxt0) oplook(p *obj.Prog) *Optab {
	if oprange[AOR&obj.AMask] == nil {
		c.ctxt.Diag("loong32r ops not initialized; call buildop first")
	}
	if p.Optab != 0 {
		return &optab[p.Optab-1]
	}
	class := func(a *obj.Addr) int {
		v := int(a.Class)
		if v == 0 {
			v = c.aclass(a) + 1
			a.Class = int8(v)
		}
		return v - 1
	}
	a1, a4 := class(&p.From), class(&p.To)
	a2, a3, a5 := C_NONE, C_NONE, C_NONE
	if p.Reg != 0 {
		a2 = c.rclass(p.Reg)
	} else if len(p.RestArgs) > 0 {
		a2 = class(&p.RestArgs[0].Addr)
	}
	if len(p.RestArgs) > 1 {
		a3 = class(&p.RestArgs[1].Addr)
	} else if p.Reg != 0 && len(p.RestArgs) > 0 {
		a3 = class(&p.RestArgs[0].Addr)
	}
	if p.RegTo2 != 0 {
		a5 = C_REG
	}
	for i := range oprange[p.As&obj.AMask] {
		op := &oprange[p.As&obj.AMask][i]
		if xcmp[a1][op.from1] && xcmp[a2][op.reg] && xcmp[a3][op.from3] && xcmp[a4][op.to1] && xcmp[a5][op.to2] {
			p.Optab = uint16(cap(optab) - cap(oprange[p.As&obj.AMask]) + i + 1)
			return op
		}
	}
	c.ctxt.Diag("illegal combination %v %v %v %v %v %v: %v", p.As, DRconv(a1), DRconv(a2), DRconv(a3), DRconv(a4), DRconv(a5), p)
	prasm(p)
	return &Optab{as: obj.AUNDEF, type_: 49, size: 4}
}

func cmp(want, got int) bool {
	if want == got {
		return true
	}
	switch want {
	case C_32CON:
		return got == C_32CON20_0 || got == C_U15CON || got == C_12CON || got == C_S12CON || got == C_U12CON || got == C_U5CON || got == C_ZCON
	case C_32CON20_0:
		return got == C_ZCON
	case C_U15CON:
		return got == C_S12CON || got == C_U12CON || got == C_U5CON || got == C_ZCON
	case C_12CON:
		return got == C_S12CON || got == C_U12CON || got == C_U5CON || got == C_ZCON
	case C_S12CON:
		return got == C_U5CON || got == C_ZCON
	case C_U12CON:
		return got == C_U5CON || got == C_ZCON
	case C_U5CON:
		return got == C_ZCON
	case C_LACON:
		return got == C_SACON
	case C_LAUTO:
		return got == C_SAUTO_14 || got == C_SAUTO
	case C_SAUTO_14:
		return got == C_SAUTO
	case C_LOREG_32:
		return got == C_SOREG_14 || got == C_SOREG_12 || got == C_ZOREG
	case C_SOREG_14:
		return got == C_SOREG_12 || got == C_ZOREG
	case C_SOREG_12:
		return got == C_ZOREG
	case C_REG:
		return got == C_ZCON
	}
	return false
}

func ocmp(a, b Optab) int {
	if a.as != b.as {
		return int(a.as - b.as)
	}
	if a.from1 != b.from1 {
		return int(a.from1) - int(b.from1)
	}
	if a.reg != b.reg {
		return int(a.reg) - int(b.reg)
	}
	if a.from3 != b.from3 {
		return int(a.from3) - int(b.from3)
	}
	if a.to1 != b.to1 {
		return int(a.to1) - int(b.to1)
	}
	return int(a.to2) - int(b.to2)
}

func opset(a, base obj.As) { oprange[a&obj.AMask] = oprange[base&obj.AMask] }

func buildop(ctxt *obj.Link) {
	if ctxt.DiagFunc == nil {
		ctxt.DiagFunc = func(format string, args ...any) { log.Printf(format, args...) }
	}
	if oprange[AOR&obj.AMask] != nil {
		return
	}
	for i := range C_NCLASS {
		for j := range C_NCLASS {
			xcmp[i][j] = cmp(j, i)
		}
	}
	slices.SortFunc(optab, ocmp)
	for i := 0; i < len(optab); i++ {
		as, start := optab[i].as, i
		for i+1 < len(optab) && optab[i+1].as == as {
			i++
		}
		oprange[as&obj.AMask] = optab[start : i+1]
	}

	for _, as := range []obj.As{AMUL, AMULH, AMULHU, ADIV, ADIVU, AREM, AREMU} {
		opset(as, ASUB)
	}
	for _, as := range []obj.As{ASLT, ASLTU, ASGTU} {
		opset(as, ASGT)
	}
	for _, as := range []obj.As{AOR, AXOR} {
		opset(as, AAND)
	}
	for _, as := range []obj.As{ASRL, ASRA} {
		opset(as, ASLL)
	}
	for _, as := range []obj.As{AMOVH} {
		opset(as, AMOVB)
	}
	for _, as := range []obj.As{ABNE, ABLT, ABGE, ABLTU, ABGEU, ABLTZ, ABGEZ, ABLEZ, ABGTZ} {
		opset(as, ABEQ)
	}
	for _, as := range []obj.As{ADBAR, AIBAR, ABREAK} {
		opset(as, ASYSCALL)
	}
	opset(ARDTIMEHW, ARDTIMELW)

	for _, as := range []obj.As{AADDD, ASUBF, ASUBD, AMULF, AMULD, ADIVF, ADIVD, AFMAXF, AFMAXD, AFMINF, AFMIND, AFMAXAF, AFMAXAD, AFMINAF, AFMINAD, AFCOPYSGF, AFCOPYSGD} {
		opset(as, AADDF)
	}
	for _, as := range []obj.As{AABSD, ANEGF, ANEGD, ASQRTF, ASQRTD, AFRECIPF, AFRECIPD, AFRSQRTF, AFRSQRTD, AFCLASSF, AFCLASSD, AMOVWF, AMOVWD, ATRUNCFW, ATRUNCDW, AMOVFD, AMOVDF} {
		opset(as, AABSF)
	}
	for _, as := range []obj.As{ACMPEQD, ACMPGEF, ACMPGED, ACMPGTF, ACMPGTD} {
		opset(as, ACMPEQF)
	}
	for _, as := range []obj.As{AFMADDD, AFMSUBF, AFMSUBD, AFNMADDF, AFNMADDD, AFNMSUBF, AFNMSUBD} {
		opset(as, AFMADDF)
	}
	opset(AMOVD, AMOVF)
	opset(ABFPF, ABFPT)
	opset(obj.AUNDEF, ANOOP)
}

func OP_RRRR(op, rk, rj, ra, rd uint32) uint32 {
	return op | (ra&31)<<15 | (rk&31)<<10 | (rj&31)<<5 | rd&31
}
func OP_RRR(op, rk, rj, rd uint32) uint32 {
	return op | (rk&31)<<10 | (rj&31)<<5 | rd&31
}
func OP_RR(op, rj, rd uint32) uint32 { return op | (rj&31)<<5 | rd&31 }
func OP_16IRR(op, imm, rj, rd uint32) uint32 {
	return op | (imm&0xffff)<<10 | (rj&31)<<5 | rd&31
}
func OP_14IRR(op, imm, rj, rd uint32) uint32 {
	return op | (imm&0x3fff)<<10 | (rj&31)<<5 | rd&31
}
func OP_12IRR(op, imm, rj, rd uint32) uint32 {
	return op | (imm&0xfff)<<10 | (rj&31)<<5 | rd&31
}
func OP_IR(op, imm, rd uint32) uint32 { return op | (imm&0xfffff)<<5 | rd&31 }
func OP_15I(op, imm uint32) uint32    { return op | (imm&0x7fff)<<0 }
func OP_16IR_5I(op, imm, rj uint32) uint32 {
	return op | (imm&0xffff)<<10 | (rj&31)<<5 | (imm>>16)&31
}
func OP_12IR_5I(op, imm, rj, hint uint32) uint32 {
	return op | (imm&0xfff)<<10 | (rj&31)<<5 | hint&31
}
func OP_B_BL(op, imm uint32) uint32 {
	return op | (imm&0xffff)<<10 | (imm>>16)&0x3ff
}

func (c *ctxt0) asmout(p *obj.Prog, o *Optab, out []uint32) {
	for i := range out {
		out[i] = 0
	}
	reg := func(r, fallback int16) int16 {
		if r == 0 {
			return fallback
		}
		return r
	}
	switch o.type_ {
	case 0:
		return
	case 1:
		out[0] = OP_RRR(c.oprrr(AOR), REGZERO, uint32(p.From.Reg), uint32(p.To.Reg))
	case 2:
		out[0] = OP_RRR(c.oprrr(p.As), uint32(p.From.Reg), uint32(reg(p.Reg, p.To.Reg)), uint32(p.To.Reg))
	case 3:
		v := c.regoff(&p.From)
		base := reg(p.From.Reg, o.param)
		as := AADD
		if o.from1 == C_12CON && v >= 0 {
			as = AOR
		}
		out[0] = OP_12IRR(c.opirr(as), uint32(v), uint32(base), uint32(p.To.Reg))
	case 4:
		out[0] = OP_12IRR(c.opirr(p.As), uint32(c.regoff(&p.From)), uint32(reg(p.Reg, p.To.Reg)), uint32(p.To.Reg))
	case 5:
		out[0] = OP_15I(c.opi(p.As), uint32(c.regoff(&p.From)))
	case 6:
		v := int32(0)
		if p.To.Target() != nil {
			v = int32(p.To.Target().Pc-p.Pc) >> 2
		}
		if p.As == ABFPT || p.As == ABFPF {
			if (v<<11)>>11 != v {
				c.ctxt.Diag("LA32R floating branch too far: %v", p)
			}
			rj := p.From.Reg
			if rj == 0 {
				rj = REG_FCC0
			}
			out[0] = OP_16IR_5I(c.opirr(p.As), uint32(v), uint32(rj))
			break
		}
		if (v<<16)>>16 != v {
			c.ctxt.Diag("LA32R conditional branch too far: %v", p)
		}
		rj, rd := p.From.Reg, reg(p.Reg, REGZERO)
		if p.As == ABGTZ || p.As == ABLEZ {
			rj, rd = rd, rj
		}
		out[0] = OP_16IRR(c.opirr(p.As), uint32(v), uint32(rj), uint32(rd))
	case 7:
		out[0] = OP_12IRR(c.opirr(p.As), uint32(c.regoff(&p.To)), uint32(reg(p.To.Reg, o.param)), uint32(p.From.Reg))
	case 8:
		out[0] = OP_12IRR(c.opirr(-p.As), uint32(c.regoff(&p.From)), uint32(reg(p.From.Reg, o.param)), uint32(p.To.Reg))
	case 9:
		out[0] = OP_RR(c.oprr(p.As), uint32(p.From.Reg), uint32(p.To.Reg))
	case 10:
		v := c.regoff(&p.From)
		as := AOR
		if v < 0 {
			as = AADD
		}
		out[0] = OP_12IRR(c.opirr(as), uint32(v), REGZERO, REGTMP)
		out[1] = OP_RRR(c.oprrr(p.As), REGTMP, uint32(reg(p.Reg, p.To.Reg)), uint32(p.To.Reg))
	case 11:
		v := int32(0)
		if p.To.Target() != nil {
			v = int32(p.To.Target().Pc-p.Pc) >> 2
			if v < -1<<25 || v >= 1<<25 {
				c.ctxt.Diag("LA32R branch too far: %v", p)
			}
		}
		out[0] = OP_B_BL(c.opirr(p.As), uint32(v))
		if p.To.Sym != nil {
			c.cursym.AddRel(c.ctxt, obj.Reloc{Type: objabi.R_CALLLOONG32R, Off: int32(c.pc), Siz: 4, Sym: p.To.Sym, Add: p.To.Offset})
		}
	case 12:
		switch p.As {
		case AMOVB:
			out[0] = OP_16IRR(c.opirr(ASLL), 24, uint32(p.From.Reg), uint32(p.To.Reg))
			out[1] = OP_16IRR(c.opirr(ASRA), 24, uint32(p.To.Reg), uint32(p.To.Reg))
		case AMOVH:
			out[0] = OP_16IRR(c.opirr(ASLL), 16, uint32(p.From.Reg), uint32(p.To.Reg))
			out[1] = OP_16IRR(c.opirr(ASRA), 16, uint32(p.To.Reg), uint32(p.To.Reg))
		case AMOVBU:
			out[0] = OP_12IRR(c.opirr(AAND), 0xff, uint32(p.From.Reg), uint32(p.To.Reg))
		case AMOVHU:
			out[0] = OP_16IRR(c.opirr(ASLL), 16, uint32(p.From.Reg), uint32(p.To.Reg))
			out[1] = OP_16IRR(c.opirr(ASRL), 16, uint32(p.To.Reg), uint32(p.To.Reg))
		}
	case 16:
		out[0] = OP_16IRR(c.opirr(p.As), uint32(c.regoff(&p.From))&31, uint32(reg(p.Reg, p.To.Reg)), uint32(p.To.Reg))
	case 18:
		out[0] = OP_16IRR(c.opirr(AJIRL), 0, uint32(p.To.Reg), uint32(o.param))
		if p.As == obj.ACALL {
			c.cursym.AddRel(c.ctxt, obj.Reloc{Type: objabi.R_CALLIND, Off: int32(c.pc)})
		}
	case 19:
		v := c.regoff(&p.From)
		out[0] = OP_IR(c.opir(ALU12IW), uint32(v>>12), uint32(p.To.Reg))
		out[1] = OP_12IRR(c.opirr(AOR), uint32(v), uint32(p.To.Reg), uint32(p.To.Reg))
	case 24:
		v := c.regoff(&p.From)
		out[0] = OP_IR(c.opir(ALU12IW), uint32(v>>12), REGTMP)
		out[1] = OP_12IRR(c.opirr(AOR), uint32(v), REGTMP, REGTMP)
		out[2] = OP_RRR(c.oprrr(p.As), REGTMP, uint32(reg(p.Reg, p.To.Reg)), uint32(p.To.Reg))
	case 25:
		out[0] = OP_IR(c.opir(ALU12IW), uint32(c.regoff(&p.From)>>12), uint32(p.To.Reg))
	case 27:
		v := c.regoff(&p.From)
		out[0] = OP_IR(c.opir(ALU12IW), uint32(v>>12), REGTMP)
		out[1] = OP_12IRR(c.opirr(AOR), uint32(v), REGTMP, REGTMP)
		out[2] = OP_RRR(c.oprrr(AADD), REGTMP, uint32(reg(p.From.Reg, o.param)), uint32(p.To.Reg))
	case 28:
		v := c.regoff(&p.From)
		base := reg(p.From.Reg, o.param)
		if o.size == 4 {
			out[0] = OP_12IRR(c.opirr(-p.As), uint32(v), uint32(base), uint32(p.To.Reg))
		} else {
			out[0] = OP_IR(c.opir(ALU12IW), uint32((v+1<<11)>>12), REGTMP)
			out[1] = OP_RRR(c.oprrr(AADD), uint32(base), REGTMP, REGTMP)
			out[2] = OP_12IRR(c.opirr(-p.As), uint32(v), REGTMP, uint32(p.To.Reg))
		}
	case 29:
		v := c.regoff(&p.To)
		base := reg(p.To.Reg, o.param)
		if o.size == 4 {
			out[0] = OP_12IRR(c.opirr(p.As), uint32(v), uint32(base), uint32(p.From.Reg))
		} else {
			out[0] = OP_IR(c.opir(ALU12IW), uint32((v+1<<11)>>12), REGTMP)
			out[1] = OP_RRR(c.oprrr(AADD), uint32(base), REGTMP, REGTMP)
			out[2] = OP_12IRR(c.opirr(p.As), uint32(v), REGTMP, uint32(p.From.Reg))
		}
	case 30:
		out[0] = OP_RR(c.specialFpMovInst(oclass(&p.From), oclass(&p.To)), uint32(p.From.Reg), uint32(p.To.Reg))
	case 33:
		fk := p.To.Reg
		if len(p.RestArgs) > 0 {
			fk = p.RestArgs[0].Reg
		}
		out[0] = 0x340<<18 | (uint32(p.From.Reg)&7)<<15 | (uint32(fk)&31)<<10 | (uint32(p.Reg)&31)<<5 | uint32(p.To.Reg)&31
	case 34:
		v := c.regoff(&p.From)
		as := AADD
		if v >= 0 {
			as = AOR
		}
		out[0] = OP_12IRR(c.opirr(as), uint32(v), REGZERO, REGTMP)
		out[1] = OP_RR(c.specialFpMovInst(C_REG, C_FREG), REGTMP, uint32(p.To.Reg))
	case 35:
		v := c.regoff(&p.To)
		base := reg(p.To.Reg, o.param)
		out[0] = OP_IR(c.opir(ALU12IW), uint32((v+1<<11)>>12), REGTMP)
		out[1] = OP_RRR(c.oprrr(AADD), uint32(base), REGTMP, REGTMP)
		out[2] = OP_12IRR(c.opirr(p.As), uint32(v), REGTMP, uint32(p.From.Reg))
	case 36:
		v := c.regoff(&p.From)
		base := reg(p.From.Reg, o.param)
		out[0] = OP_IR(c.opir(ALU12IW), uint32((v+1<<11)>>12), REGTMP)
		out[1] = OP_RRR(c.oprrr(AADD), uint32(base), REGTMP, REGTMP)
		out[2] = OP_12IRR(c.opirr(-p.As), uint32(v), REGTMP, uint32(p.To.Reg))
	case 37:
		fa := p.To.Reg
		if len(p.RestArgs) > 0 {
			fa = p.RestArgs[0].Reg
		}
		out[0] = OP_RRRR(c.oprrrr(p.As), uint32(p.From.Reg), uint32(p.Reg), uint32(fa), uint32(p.To.Reg))
	case 38:
		out[0] = uint32(c.regoff(&p.From))
	case 40:
		if p.As == ALL {
			v := c.regoff(&p.From)
			if v&3 != 0 {
				c.ctxt.Diag("LL.W offset must be a multiple of four: %v", p)
			}
			out[0] = OP_14IRR(c.opirr(-ALL), uint32(v>>2), uint32(reg(p.From.Reg, o.param)), uint32(p.To.Reg))
		} else {
			v := c.regoff(&p.To)
			if v&3 != 0 {
				c.ctxt.Diag("SC.W offset must be a multiple of four: %v", p)
			}
			out[0] = OP_14IRR(c.opirr(ASC), uint32(v>>2), uint32(reg(p.To.Reg, o.param)), uint32(p.From.Reg))
		}
	case 47:
		hint := int64(p.Reg)
		if len(p.RestArgs) > 0 {
			hint = p.RestArgs[0].Offset
		}
		out[0] = OP_12IR_5I(c.opiir(APRELD), uint32(c.regoff(&p.From)), uint32(p.From.Reg), uint32(hint))
	case 49:
		if p.As == ANOOP {
			out[0] = OP_12IRR(c.opirr(AAND), 0, 0, 0)
		} else {
			out[0] = OP_15I(c.opi(ABREAK), 0)
		}
	case 50:
		out[0] = OP_IR(c.opir(APCADDU12I), 0, REGTMP)
		c.addRel(objabi.R_LOONG32R_ADDR_HI, c.pc, p.To.Sym, p.To.Offset)
		out[1] = OP_12IRR(c.opirr(p.As), 0, REGTMP, uint32(p.From.Reg))
		c.addRel(objabi.R_LOONG32R_ADDR_LO, c.pc+4, p.To.Sym, p.To.Offset)
	case 51:
		out[0] = OP_IR(c.opir(APCADDU12I), 0, REGTMP)
		c.addRel(objabi.R_LOONG32R_ADDR_HI, c.pc, p.From.Sym, p.From.Offset)
		out[1] = OP_12IRR(c.opirr(-p.As), 0, REGTMP, uint32(p.To.Reg))
		c.addRel(objabi.R_LOONG32R_ADDR_LO, c.pc+4, p.From.Sym, p.From.Offset)
	case 52:
		out[0] = OP_IR(c.opir(APCADDU12I), 0, uint32(p.To.Reg))
		c.addRel(objabi.R_LOONG32R_ADDR_HI, c.pc, p.From.Sym, p.From.Offset)
		out[1] = OP_12IRR(c.opirr(AADD), 0, uint32(p.To.Reg), uint32(p.To.Reg))
		c.addRel(objabi.R_LOONG32R_ADDR_LO, c.pc+4, p.From.Sym, p.From.Offset)
	case 53, 54:
		var sym *obj.LSym
		var add int64
		if o.type_ == 53 {
			sym, add = p.To.Sym, p.To.Offset
		} else {
			sym, add = p.From.Sym, p.From.Offset
		}
		out[0] = OP_IR(c.opir(ALU12IW), 0, REGTMP)
		c.addRel(objabi.R_LOONG32R_TLS_LE_HI, c.pc, sym, add)
		out[1] = OP_12IRR(c.opirr(AOR), 0, REGTMP, REGTMP)
		c.addRel(objabi.R_LOONG32R_TLS_LE_LO, c.pc+4, sym, add)
		out[2] = OP_RRR(c.oprrr(AADD), REG_R2, REGTMP, REGTMP)
		if o.type_ == 53 {
			out[3] = OP_12IRR(c.opirr(p.As), 0, REGTMP, uint32(p.From.Reg))
		} else {
			out[3] = OP_12IRR(c.opirr(-p.As), 0, REGTMP, uint32(p.To.Reg))
		}
	case 56, 57:
		var sym *obj.LSym
		if o.type_ == 56 {
			sym = p.To.Sym
		} else {
			sym = p.From.Sym
		}
		out[0] = OP_IR(c.opir(APCADDU12I), 0, REGTMP)
		c.addRel(objabi.R_LOONG32R_TLS_IE_HI, c.pc, sym, 0)
		out[1] = OP_12IRR(c.opirr(-AMOVW), 0, REGTMP, REGTMP)
		c.addRel(objabi.R_LOONG32R_TLS_IE_LO, c.pc+4, sym, 0)
		out[2] = OP_RRR(c.oprrr(AADD), REGTMP, REG_R2, REGTMP)
		if o.type_ == 56 {
			out[3] = OP_12IRR(c.opirr(p.As), 0, REGTMP, uint32(p.From.Reg))
		} else {
			out[3] = OP_12IRR(c.opirr(-p.As), 0, REGTMP, uint32(p.To.Reg))
		}
	case 62:
		// RDTIMEL.W/RDTIMEH.W name both architectural outputs in Plan 9
		// order: the time value first (rd), then the counter ID (rj).
		out[0] = OP_RR(c.oprr(p.As), uint32(p.RegTo2), uint32(p.To.Reg))
	case 63:
		out[0] = OP_RR(c.oprr(ARDTIMEID), uint32(p.From.Reg), REGZERO)
	case 65:
		out[0] = OP_IR(c.opir(APCADDU12I), 0, uint32(p.To.Reg))
		c.addRel(objabi.R_LOONG32R_GOT_HI, c.pc, p.From.Sym, 0)
		out[1] = OP_12IRR(c.opirr(-AMOVW), 0, uint32(p.To.Reg), uint32(p.To.Reg))
		c.addRel(objabi.R_LOONG32R_GOT_LO, c.pc+4, p.From.Sym, 0)
	case 66:
		v := int64(c.regoff(&p.From))
		if v < -1<<19 || v >= 1<<20 {
			c.ctxt.Diag("LA32R 20-bit immediate out of range: %v", p)
		}
		out[0] = OP_IR(c.opir(p.As), uint32(v), uint32(p.To.Reg))
	case 67:
		v := int64(c.regoff(&p.From))
		if v&3 != 0 || v < -1<<17 || v >= 1<<17 {
			c.ctxt.Diag("JIRL offset must be a signed, 4-byte-aligned 18-bit value: %v", p)
		}
		out[0] = OP_16IRR(c.opirr(AJIRL), uint32(v>>2), uint32(p.Reg), uint32(p.To.Reg))
	default:
		c.ctxt.Diag("unknown LA32R encoding type %d for %v", o.type_, p)
	}
}

func (c *ctxt0) addRel(typ objabi.RelocType, off int64, sym *obj.LSym, add int64) {
	c.cursym.AddRel(c.ctxt, obj.Reloc{Type: typ, Off: int32(off), Siz: 4, Sym: sym, Add: add})
}

func (c *ctxt0) oprrrr(as obj.As) uint32 {
	switch as {
	case AFMADDF:
		return 0x81 << 20
	case AFMADDD:
		return 0x82 << 20
	case AFMSUBF:
		return 0x85 << 20
	case AFMSUBD:
		return 0x86 << 20
	case AFNMADDF:
		return 0x89 << 20
	case AFNMADDD:
		return 0x8a << 20
	case AFNMSUBF:
		return 0x8d << 20
	case AFNMSUBD:
		return 0x8e << 20
	}
	c.ctxt.Diag("bad LA32R rrrr opcode %v", as)
	return 0
}

func (c *ctxt0) oprrr(as obj.As) uint32 {
	switch as {
	case AADD:
		return 0x20 << 15
	case ASUB:
		return 0x22 << 15
	case ASLT, ASGT:
		return 0x24 << 15
	case ASLTU, ASGTU:
		return 0x25 << 15
	case ANOR:
		return 0x28 << 15
	case AAND:
		return 0x29 << 15
	case AOR:
		return 0x2a << 15
	case AXOR:
		return 0x2b << 15
	case ASLL:
		return 0x2e << 15
	case ASRL:
		return 0x2f << 15
	case ASRA:
		return 0x30 << 15
	case AMUL:
		return 0x38 << 15
	case AMULH:
		return 0x39 << 15
	case AMULHU:
		return 0x3a << 15
	case ADIV:
		return 0x40 << 15
	case AREM:
		return 0x41 << 15
	case ADIVU:
		return 0x42 << 15
	case AREMU:
		return 0x43 << 15
	case AADDF:
		return 0x201 << 15
	case AADDD:
		return 0x202 << 15
	case ASUBF:
		return 0x205 << 15
	case ASUBD:
		return 0x206 << 15
	case AMULF:
		return 0x209 << 15
	case AMULD:
		return 0x20a << 15
	case ADIVF:
		return 0x20d << 15
	case ADIVD:
		return 0x20e << 15
	case AFMAXF:
		return 0x211 << 15
	case AFMAXD:
		return 0x212 << 15
	case AFMINF:
		return 0x215 << 15
	case AFMIND:
		return 0x216 << 15
	case AFMAXAF:
		return 0x219 << 15
	case AFMAXAD:
		return 0x21a << 15
	case AFMINAF:
		return 0x21d << 15
	case AFMINAD:
		return 0x21e << 15
	case AFCOPYSGF:
		return 0x225 << 15
	case AFCOPYSGD:
		return 0x226 << 15
	case ACMPEQF:
		return 0x0c1<<20 | 0x4<<15
	case ACMPEQD:
		return 0x0c2<<20 | 0x4<<15
	case ACMPGEF:
		return 0x0c1<<20 | 0x7<<15 // reversed operands: SLE
	case ACMPGED:
		return 0x0c2<<20 | 0x7<<15
	case ACMPGTF:
		return 0x0c1<<20 | 0x3<<15 // reversed operands: SLT
	case ACMPGTD:
		return 0x0c2<<20 | 0x3<<15
	}
	c.ctxt.Diag("bad LA32R rrr opcode %v", as)
	return 0
}

func (c *ctxt0) oprr(as obj.As) uint32 {
	switch as {
	case ARDTIMELW:
		return 0x18 << 10
	case ARDTIMEHW:
		return 0x19 << 10
	case ARDTIMEID:
		return 0x18 << 10
	case AABSF:
		return 0x4501 << 10
	case AABSD:
		return 0x4502 << 10
	case ANEGF:
		return 0x4505 << 10
	case ANEGD:
		return 0x4506 << 10
	case ASQRTF:
		return 0x4511 << 10
	case ASQRTD:
		return 0x4512 << 10
	case AFRECIPF:
		return 0x4515 << 10
	case AFRECIPD:
		return 0x4516 << 10
	case AFRSQRTF:
		return 0x4519 << 10
	case AFRSQRTD:
		return 0x451a << 10
	case AMOVF:
		return 0x4525 << 10
	case AMOVD:
		return 0x4526 << 10
	case AFCLASSF:
		return 0x450d << 10
	case AFCLASSD:
		return 0x450e << 10
	case AMOVWF:
		return 0x4744 << 10
	case AMOVWD:
		return 0x4748 << 10
	case ATRUNCFW:
		return 0x46a1 << 10
	case ATRUNCDW:
		return 0x46a2 << 10
	case AMOVFD:
		return 0x4649 << 10
	case AMOVDF:
		return 0x4646 << 10
	}
	c.ctxt.Diag("bad LA32R rr opcode %v", as)
	return 0
}

func (c *ctxt0) opi(as obj.As) uint32 {
	switch as {
	case ABREAK:
		return 0x54 << 15
	case ASYSCALL:
		return 0x56 << 15
	case ADBAR:
		return 0x70e4 << 15
	case AIBAR:
		return 0x70e5 << 15
	}
	c.ctxt.Diag("bad LA32R immediate opcode %v", as)
	return 0
}

func (c *ctxt0) opir(as obj.As) uint32 {
	switch as {
	case ALU12IW:
		return 0x0a << 25
	case APCADDU12I:
		return 0x0e << 25
	}
	c.ctxt.Diag("bad LA32R 20-bit opcode %v", as)
	return 0
}

func (c *ctxt0) opirr(as obj.As) uint32 {
	switch as {
	case AADD:
		return 0x00a << 22
	case ASLT, ASGT:
		return 0x008 << 22
	case ASLTU, ASGTU:
		return 0x009 << 22
	case AAND:
		return 0x00d << 22
	case AOR:
		return 0x00e << 22
	case AXOR:
		return 0x00f << 22
	case ASLL:
		return 0x00081 << 15
	case ASRL:
		return 0x00089 << 15
	case ASRA:
		return 0x00091 << 15
	case AJIRL:
		return 0x13 << 26
	case AJMP:
		return 0x14 << 26
	case AJAL:
		return 0x15 << 26
	case ABEQ:
		return 0x16 << 26
	case ABNE:
		return 0x17 << 26
	case ABLT, ABLTZ, ABGTZ:
		return 0x18 << 26
	case ABGE, ABGEZ, ABLEZ:
		return 0x19 << 26
	case ABLTU:
		return 0x1a << 26
	case ABGEU:
		return 0x1b << 26
	case ABFPF:
		return 0x12 << 26
	case ABFPT:
		return 0x12<<26 | 1<<8
	case AMOVB, AMOVBU:
		return 0x0a4 << 22
	case AMOVH, AMOVHU:
		return 0x0a5 << 22
	case AMOVW:
		return 0x0a6 << 22
	case -AMOVB:
		return 0x0a0 << 22
	case -AMOVH:
		return 0x0a1 << 22
	case -AMOVW:
		return 0x0a2 << 22
	case -AMOVBU:
		return 0x0a8 << 22
	case -AMOVHU:
		return 0x0a9 << 22
	case AMOVF:
		return 0x0ad << 22
	case AMOVD:
		return 0x0af << 22
	case -AMOVF:
		return 0x0ac << 22
	case -AMOVD:
		return 0x0ae << 22
	case -ALL:
		return 0x020 << 24
	case ASC:
		return 0x021 << 24
	}
	if as < 0 {
		c.ctxt.Diag("bad LA32R immediate/register opcode -%v", -as)
	} else {
		c.ctxt.Diag("bad LA32R immediate/register opcode %v", as)
	}
	return 0
}

func (c *ctxt0) opiir(as obj.As) uint32 {
	if as == APRELD {
		return 0x0ab << 22
	}
	c.ctxt.Diag("bad LA32R prefetch opcode %v", as)
	return 0
}

func (c *ctxt0) specialFpMovInst(from, to int) uint32 {
	switch {
	case from == C_REG && to == C_FREG:
		return 0x4529 << 10 // movgr2fr.w
	case from == C_FREG && to == C_REG:
		return 0x452d << 10 // movfr2gr.s
	}
	c.ctxt.Diag("bad LA32R floating move classes %d,%d", from, to)
	return 0
}
