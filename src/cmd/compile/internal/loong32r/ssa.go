// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package loong32r

import (
	"math"

	"cmd/compile/internal/base"
	"cmd/compile/internal/ir"
	"cmd/compile/internal/logopt"
	"cmd/compile/internal/ssa"
	"cmd/compile/internal/ssagen"
	"cmd/compile/internal/types"
	"cmd/internal/obj"
	"cmd/internal/obj/loong32r"
	"internal/abi"
)

// isFPreg reports whether r is an FP register.
func isFPreg(r int16) bool {
	return loong32r.REG_F0 <= r && r <= loong32r.REG_F31
}

func dbar(s *ssagen.State) *obj.Prog {
	p := s.Prog(loong32r.ADBAR)
	p.From.SetConst(0)
	return p
}

// loadByType returns the load instruction of the given type.
func loadByType(t *types.Type, r int16) obj.As {
	if isFPreg(r) {
		if t.Size() == 4 { // float32 or int32
			return loong32r.AMOVF
		} else { // float64 or int64
			return loong32r.AMOVD
		}
	} else {
		switch t.Size() {
		case 1:
			if t.IsSigned() {
				return loong32r.AMOVB
			} else {
				return loong32r.AMOVBU
			}
		case 2:
			if t.IsSigned() {
				return loong32r.AMOVH
			} else {
				return loong32r.AMOVHU
			}
		case 4:
			return loong32r.AMOVW
		}
	}
	panic("bad load type")
}

// storeByType returns the store instruction of the given type.
func storeByType(t *types.Type, r int16) obj.As {
	if isFPreg(r) {
		if t.Size() == 4 { // float32 or int32
			return loong32r.AMOVF
		} else { // float64 or int64
			return loong32r.AMOVD
		}
	} else {
		switch t.Size() {
		case 1:
			return loong32r.AMOVB
		case 2:
			return loong32r.AMOVH
		case 4:
			return loong32r.AMOVW
		}
	}
	panic("bad store type")
}

func ssaGenValue(s *ssagen.State, v *ssa.Value) {
	switch v.Op {
	case ssa.OpCopy, ssa.OpLOONG32RMOVWreg:
		t := v.Type
		if t.IsMemory() {
			return
		}
		x := v.Args[0].Reg()
		y := v.Reg()
		if x == y {
			return
		}
		as := loong32r.AMOVW
		if isFPreg(x) && isFPreg(y) {
			as = loong32r.AMOVF
			if t.Size() == 8 {
				as = loong32r.AMOVD
			}
		}

		p := s.Prog(as)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = x
		p.To.Type = obj.TYPE_REG
		p.To.Reg = y
	case ssa.OpLOONG32RMOVWnop:
		// nothing to do
	case ssa.OpLoadReg:
		if v.Type.IsFlags() {
			v.Fatalf("load flags not implemented: %v", v.LongString())
			return
		}
		r := v.Reg()
		p := s.Prog(loadByType(v.Type, r))
		ssagen.AddrAuto(&p.From, v.Args[0])
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
	case ssa.OpStoreReg:
		if v.Type.IsFlags() {
			v.Fatalf("store flags not implemented: %v", v.LongString())
			return
		}
		r := v.Args[0].Reg()
		p := s.Prog(storeByType(v.Type, r))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = r
		ssagen.AddrAuto(&p.To, v)
	case ssa.OpLOONG32RADD,
		ssa.OpLOONG32RSUB,
		ssa.OpLOONG32RAND,
		ssa.OpLOONG32ROR,
		ssa.OpLOONG32RXOR,
		ssa.OpLOONG32RNOR,
		ssa.OpLOONG32RSLL,
		ssa.OpLOONG32RSRL,
		ssa.OpLOONG32RSRA,
		ssa.OpLOONG32RADDF,
		ssa.OpLOONG32RADDD,
		ssa.OpLOONG32RSUBF,
		ssa.OpLOONG32RSUBD,
		ssa.OpLOONG32RMULF,
		ssa.OpLOONG32RMULD,
		ssa.OpLOONG32RDIVF,
		ssa.OpLOONG32RDIVD,
		ssa.OpLOONG32RMUL:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg()
		p.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpLOONG32RSGT,
		ssa.OpLOONG32RSGTU:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg()
		p.Reg = v.Args[1].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpLOONG32RSGTzero,
		ssa.OpLOONG32RSGTUzero:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg()
		p.Reg = loong32r.REGZERO
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpLOONG32RADDconst,
		ssa.OpLOONG32RSUBconst,
		ssa.OpLOONG32RANDconst,
		ssa.OpLOONG32RORconst,
		ssa.OpLOONG32RXORconst,
		ssa.OpLOONG32RNORconst,
		ssa.OpLOONG32RSLLconst,
		ssa.OpLOONG32RSRLconst,
		ssa.OpLOONG32RSRAconst,
		ssa.OpLOONG32RSGTconst,
		ssa.OpLOONG32RSGTUconst:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpLOONG32RMULT,
		ssa.OpLOONG32RMULTU:
		high := loong32r.AMULH
		if v.Op == ssa.OpLOONG32RMULTU {
			high = loong32r.AMULHU
		}
		p := s.Prog(high)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg()
		p.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg0()
		p = s.Prog(loong32r.AMUL)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg()
		p.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg1()
	case ssa.OpLOONG32RDIV,
		ssa.OpLOONG32RDIVU:
		rem, div := loong32r.AREM, loong32r.ADIV
		if v.Op == ssa.OpLOONG32RDIVU {
			rem, div = loong32r.AREMU, loong32r.ADIVU
		}
		p := s.Prog(rem)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg()
		p.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg0()
		p = s.Prog(div)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg()
		p.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg1()
	case ssa.OpLOONG32RMOVWconst:
		r := v.Reg()
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
		if isFPreg(r) {
			// Constants cannot be moved directly into FP registers.
			p.To.Reg = loong32r.REGTMP
			p = s.Prog(loong32r.AMOVW)
			p.From.Type = obj.TYPE_REG
			p.From.Reg = loong32r.REGTMP
			p.To.Type = obj.TYPE_REG
			p.To.Reg = r
		}
	case ssa.OpLOONG32RMOVFconst,
		ssa.OpLOONG32RMOVDconst:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_FCONST
		p.From.Val = math.Float64frombits(uint64(v.AuxInt))
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpLOONG32RCMOVZ:
		branch := s.Prog(loong32r.ABNE)
		branch.From.Type = obj.TYPE_REG
		branch.From.Reg = v.Args[2].Reg()
		branch.To.Type = obj.TYPE_BRANCH
		p := s.Prog(loong32r.AMOVW)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
		done := s.Prog(obj.ANOP)
		branch.To.SetTarget(done)
	case ssa.OpLOONG32RCMOVZzero:
		branch := s.Prog(loong32r.ABNE)
		branch.From.Type = obj.TYPE_REG
		branch.From.Reg = v.Args[1].Reg()
		branch.To.Type = obj.TYPE_BRANCH
		p := s.Prog(loong32r.AMOVW)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = loong32r.REGZERO
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
		done := s.Prog(obj.ANOP)
		branch.To.SetTarget(done)
	case ssa.OpLOONG32RCMPEQF,
		ssa.OpLOONG32RCMPEQD,
		ssa.OpLOONG32RCMPGEF,
		ssa.OpLOONG32RCMPGED,
		ssa.OpLOONG32RCMPGTF,
		ssa.OpLOONG32RCMPGTD:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg()
		p.Reg = v.Args[1].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = loong32r.REG_FCC0
	case ssa.OpLOONG32RMOVWaddr:
		p := s.Prog(loong32r.AMOVW)
		p.From.Type = obj.TYPE_ADDR
		p.From.Reg = v.Args[0].Reg()
		var wantreg string
		// MOVW $sym+off(base), R
		// the assembler expands it as the following:
		// - base is SP (R3): add the constant offset to SP
		//               when the constant is large, REGTMP (R30) may be used
		// - base is SB: load external address with relocation
		switch v.Aux.(type) {
		default:
			v.Fatalf("aux is of unknown type %T", v.Aux)
		case *obj.LSym:
			wantreg = "SB"
			ssagen.AddAux(&p.From, v)
		case *ir.Name:
			wantreg = "SP"
			ssagen.AddAux(&p.From, v)
		case nil:
			// No sym, just MOVW $off(SP), R
			wantreg = "SP"
			p.From.Offset = v.AuxInt
		}
		if reg := v.Args[0].RegName(); reg != wantreg {
			v.Fatalf("bad reg %s for symbol type %T, want %s", reg, v.Aux, wantreg)
		}
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpLOONG32RMOVBload,
		ssa.OpLOONG32RMOVBUload,
		ssa.OpLOONG32RMOVHload,
		ssa.OpLOONG32RMOVHUload,
		ssa.OpLOONG32RMOVWload,
		ssa.OpLOONG32RMOVFload,
		ssa.OpLOONG32RMOVDload:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg()
		ssagen.AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpLOONG32RMOVBstore,
		ssa.OpLOONG32RMOVHstore,
		ssa.OpLOONG32RMOVWstore,
		ssa.OpLOONG32RMOVFstore,
		ssa.OpLOONG32RMOVDstore:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg()
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg()
		ssagen.AddAux(&p.To, v)
	case ssa.OpLOONG32RMOVBstorezero,
		ssa.OpLOONG32RMOVHstorezero,
		ssa.OpLOONG32RMOVWstorezero:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = loong32r.REGZERO
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg()
		ssagen.AddAux(&p.To, v)
	case ssa.OpLOONG32RMOVBreg,
		ssa.OpLOONG32RMOVBUreg,
		ssa.OpLOONG32RMOVHreg,
		ssa.OpLOONG32RMOVHUreg:
		a := v.Args[0]
		for a.Op == ssa.OpCopy || a.Op == ssa.OpLOONG32RMOVWreg || a.Op == ssa.OpLOONG32RMOVWnop {
			a = a.Args[0]
		}
		if a.Op == ssa.OpLoadReg {
			t := a.Type
			switch {
			case v.Op == ssa.OpLOONG32RMOVBreg && t.Size() == 1 && t.IsSigned(),
				v.Op == ssa.OpLOONG32RMOVBUreg && t.Size() == 1 && !t.IsSigned(),
				v.Op == ssa.OpLOONG32RMOVHreg && t.Size() == 2 && t.IsSigned(),
				v.Op == ssa.OpLOONG32RMOVHUreg && t.Size() == 2 && !t.IsSigned():
				// arg is a proper-typed load, already zero/sign-extended, don't extend again
				if v.Reg() == v.Args[0].Reg() {
					return
				}
				p := s.Prog(loong32r.AMOVW)
				p.From.Type = obj.TYPE_REG
				p.From.Reg = v.Args[0].Reg()
				p.To.Type = obj.TYPE_REG
				p.To.Reg = v.Reg()
				return
			default:
			}
		}
		fallthrough
	case ssa.OpLOONG32RMOVWF,
		ssa.OpLOONG32RMOVWD,
		ssa.OpLOONG32RTRUNCFW,
		ssa.OpLOONG32RTRUNCDW,
		ssa.OpLOONG32RMOVFD,
		ssa.OpLOONG32RMOVDF,
		ssa.OpLOONG32RMOVWfpgp,
		ssa.OpLOONG32RMOVWgpfp,
		ssa.OpLOONG32RNEGF,
		ssa.OpLOONG32RNEGD,
		ssa.OpLOONG32RABSD,
		ssa.OpLOONG32RSQRTF,
		ssa.OpLOONG32RSQRTD:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpLOONG32RCTZ:
		// Isolate the lowest set bit, then count shifts until it vanishes.
		// The zero input is handled separately and returns 32.
		out, tmp := v.Reg(), v.RegTmp()
		p := s.Prog(loong32r.ASUB)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = out
		p.Reg = loong32r.REGZERO
		p.To.Type = obj.TYPE_REG
		p.To.Reg = tmp
		p = s.Prog(loong32r.AAND)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = tmp
		p.Reg = out
		p.To.Type = obj.TYPE_REG
		p.To.Reg = out
		zero := s.Prog(loong32r.ABEQ)
		zero.From.Type = obj.TYPE_REG
		zero.From.Reg = out
		zero.To.Type = obj.TYPE_BRANCH
		p = s.Prog(loong32r.AMOVW)
		p.From.SetConst(0)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = tmp
		loop := s.Prog(loong32r.ASRL)
		loop.From.SetConst(1)
		loop.Reg = out
		loop.To.Type = obj.TYPE_REG
		loop.To.Reg = out
		done := s.Prog(loong32r.ABEQ)
		done.From.Type = obj.TYPE_REG
		done.From.Reg = out
		done.To.Type = obj.TYPE_BRANCH
		p = s.Prog(loong32r.AADD)
		p.From.SetConst(1)
		p.Reg = tmp
		p.To.Type = obj.TYPE_REG
		p.To.Reg = tmp
		back := s.Prog(loong32r.AJMP)
		back.To.Type = obj.TYPE_BRANCH
		back.To.SetTarget(loop)
		zeroValue := s.Prog(loong32r.AMOVW)
		zeroValue.From.SetConst(32)
		zeroValue.To.Type = obj.TYPE_REG
		zeroValue.To.Reg = tmp
		finish := s.Prog(loong32r.AMOVW)
		finish.From.Type = obj.TYPE_REG
		finish.From.Reg = tmp
		finish.To.Type = obj.TYPE_REG
		finish.To.Reg = out
		zero.To.SetTarget(zeroValue)
		done.To.SetTarget(finish)
	case ssa.OpLOONG32RBITLEN:
		// Shift the input down while counting significant bits.
		out, tmp := v.Reg(), v.RegTmp()
		p := s.Prog(loong32r.AMOVW)
		p.From.SetConst(0)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = tmp
		loop := s.Prog(loong32r.ABEQ)
		loop.From.Type = obj.TYPE_REG
		loop.From.Reg = out
		loop.To.Type = obj.TYPE_BRANCH
		p = s.Prog(loong32r.ASRL)
		p.From.SetConst(1)
		p.Reg = out
		p.To.Type = obj.TYPE_REG
		p.To.Reg = out
		p = s.Prog(loong32r.AADD)
		p.From.SetConst(1)
		p.Reg = tmp
		p.To.Type = obj.TYPE_REG
		p.To.Reg = tmp
		back := s.Prog(loong32r.AJMP)
		back.To.Type = obj.TYPE_BRANCH
		back.To.SetTarget(loop)
		finish := s.Prog(loong32r.AMOVW)
		finish.From.Type = obj.TYPE_REG
		finish.From.Reg = tmp
		finish.To.Type = obj.TYPE_REG
		finish.To.Reg = out
		loop.To.SetTarget(finish)
	case ssa.OpLOONG32RNEG:
		// SUB from REGZERO
		p := s.Prog(loong32r.ASUB)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg()
		p.Reg = loong32r.REGZERO
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpLOONG32RLoweredZero:
		// SUBU	$4, R1
		// MOVW	R0, 4(R1)
		// ADDU	$4, R1
		// BNE	Rarg1, R1, -2(PC)
		// arg1 is the address of the last element to zero
		var sz int64
		var mov obj.As
		switch {
		case v.AuxInt%4 == 0:
			sz = 4
			mov = loong32r.AMOVW
		case v.AuxInt%2 == 0:
			sz = 2
			mov = loong32r.AMOVH
		default:
			sz = 1
			mov = loong32r.AMOVB
		}
		p := s.Prog(loong32r.ASUB)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = sz
		p.To.Type = obj.TYPE_REG
		p.To.Reg = loong32r.REG_R20
		p2 := s.Prog(mov)
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = loong32r.REGZERO
		p2.To.Type = obj.TYPE_MEM
		p2.To.Reg = loong32r.REG_R20
		p2.To.Offset = sz
		p3 := s.Prog(loong32r.AADD)
		p3.From.Type = obj.TYPE_CONST
		p3.From.Offset = sz
		p3.To.Type = obj.TYPE_REG
		p3.To.Reg = loong32r.REG_R20
		p4 := s.Prog(loong32r.ABNE)
		p4.From.Type = obj.TYPE_REG
		p4.From.Reg = v.Args[1].Reg()
		p4.Reg = loong32r.REG_R20
		p4.To.Type = obj.TYPE_BRANCH
		p4.To.SetTarget(p2)
	case ssa.OpLOONG32RLoweredMove:
		// SUBU	$4, R1
		// MOVW	4(R1), Rtmp
		// MOVW	Rtmp, (R2)
		// ADDU	$4, R1
		// ADDU	$4, R2
		// BNE	Rarg2, R1, -4(PC)
		// arg2 is the address of the last element of src
		var sz int64
		var mov obj.As
		switch {
		case v.AuxInt%4 == 0:
			sz = 4
			mov = loong32r.AMOVW
		case v.AuxInt%2 == 0:
			sz = 2
			mov = loong32r.AMOVH
		default:
			sz = 1
			mov = loong32r.AMOVB
		}
		p := s.Prog(loong32r.ASUB)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = sz
		p.To.Type = obj.TYPE_REG
		p.To.Reg = loong32r.REG_R21
		p2 := s.Prog(mov)
		p2.From.Type = obj.TYPE_MEM
		p2.From.Reg = loong32r.REG_R21
		p2.From.Offset = sz
		p2.To.Type = obj.TYPE_REG
		p2.To.Reg = loong32r.REGTMP
		p3 := s.Prog(mov)
		p3.From.Type = obj.TYPE_REG
		p3.From.Reg = loong32r.REGTMP
		p3.To.Type = obj.TYPE_MEM
		p3.To.Reg = loong32r.REG_R20
		p4 := s.Prog(loong32r.AADD)
		p4.From.Type = obj.TYPE_CONST
		p4.From.Offset = sz
		p4.To.Type = obj.TYPE_REG
		p4.To.Reg = loong32r.REG_R21
		p5 := s.Prog(loong32r.AADD)
		p5.From.Type = obj.TYPE_CONST
		p5.From.Offset = sz
		p5.To.Type = obj.TYPE_REG
		p5.To.Reg = loong32r.REG_R20
		p6 := s.Prog(loong32r.ABNE)
		p6.From.Type = obj.TYPE_REG
		p6.From.Reg = v.Args[2].Reg()
		p6.Reg = loong32r.REG_R21
		p6.To.Type = obj.TYPE_BRANCH
		p6.To.SetTarget(p2)
	case ssa.OpLOONG32RCALLstatic, ssa.OpLOONG32RCALLclosure, ssa.OpLOONG32RCALLinter:
		s.Call(v)
	case ssa.OpLOONG32RCALLtail:
		s.TailCall(v)
	case ssa.OpLOONG32RLoweredWB:
		p := s.Prog(obj.ACALL)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		// AuxInt encodes how many buffer entries we need.
		p.To.Sym = ir.Syms.GCWriteBarrier[v.AuxInt-1]

	case ssa.OpLOONG32RLoweredPanicBoundsRR, ssa.OpLOONG32RLoweredPanicBoundsRC, ssa.OpLOONG32RLoweredPanicBoundsCR, ssa.OpLOONG32RLoweredPanicBoundsCC,
		ssa.OpLOONG32RLoweredPanicExtendRR, ssa.OpLOONG32RLoweredPanicExtendRC:
		// Compute the constant we put in the PCData entry for this call.
		code, signed := ssa.BoundsKind(v.AuxInt).Code()
		xIsReg := false
		yIsReg := false
		xVal := 0
		yVal := 0
		extend := false
		switch v.Op {
		case ssa.OpLOONG32RLoweredPanicBoundsRR:
			xIsReg = true
			xVal = int(v.Args[0].Reg() - loong32r.REG_R4)
			yIsReg = true
			yVal = int(v.Args[1].Reg() - loong32r.REG_R4)
		case ssa.OpLOONG32RLoweredPanicExtendRR:
			extend = true
			xIsReg = true
			hi := int(v.Args[0].Reg() - loong32r.REG_R4)
			lo := int(v.Args[1].Reg() - loong32r.REG_R4)
			xVal = hi<<2 + lo // encode 2 register numbers
			yIsReg = true
			yVal = int(v.Args[2].Reg() - loong32r.REG_R4)
		case ssa.OpLOONG32RLoweredPanicBoundsRC:
			xIsReg = true
			xVal = int(v.Args[0].Reg() - loong32r.REG_R4)
			c := v.Aux.(ssa.PanicBoundsC).C
			if c >= 0 && c <= abi.BoundsMaxConst {
				yVal = int(c)
			} else {
				// Move constant to a register
				yIsReg = true
				if yVal == xVal {
					yVal = 1
				}
				p := s.Prog(loong32r.AMOVW)
				p.From.Type = obj.TYPE_CONST
				p.From.Offset = c
				p.To.Type = obj.TYPE_REG
				p.To.Reg = loong32r.REG_R4 + int16(yVal)
			}
		case ssa.OpLOONG32RLoweredPanicExtendRC:
			extend = true
			xIsReg = true
			hi := int(v.Args[0].Reg() - loong32r.REG_R4)
			lo := int(v.Args[1].Reg() - loong32r.REG_R4)
			xVal = hi<<2 + lo // encode 2 register numbers
			c := v.Aux.(ssa.PanicBoundsC).C
			if c >= 0 && c <= abi.BoundsMaxConst {
				yVal = int(c)
			} else {
				// Move constant to a register
				for yVal == hi || yVal == lo {
					yVal++
				}
				p := s.Prog(loong32r.AMOVW)
				p.From.Type = obj.TYPE_CONST
				p.From.Offset = c
				p.To.Type = obj.TYPE_REG
				p.To.Reg = loong32r.REG_R4 + int16(yVal)
			}
		case ssa.OpLOONG32RLoweredPanicBoundsCR:
			yIsReg = true
			yVal = int(v.Args[0].Reg() - loong32r.REG_R4)
			c := v.Aux.(ssa.PanicBoundsC).C
			if c >= 0 && c <= abi.BoundsMaxConst {
				xVal = int(c)
			} else if signed && int64(int32(c)) == c || !signed && int64(uint32(c)) == c {
				// Move constant to a register
				xIsReg = true
				if xVal == yVal {
					xVal = 1
				}
				p := s.Prog(loong32r.AMOVW)
				p.From.Type = obj.TYPE_CONST
				p.From.Offset = c
				p.To.Type = obj.TYPE_REG
				p.To.Reg = loong32r.REG_R4 + int16(xVal)
			} else {
				// Move constant to two registers
				extend = true
				xIsReg = true
				hi := 0
				lo := 1
				if hi == yVal {
					hi = 2
				}
				if lo == yVal {
					lo = 2
				}
				xVal = hi<<2 + lo
				p := s.Prog(loong32r.AMOVW)
				p.From.Type = obj.TYPE_CONST
				p.From.Offset = c >> 32
				p.To.Type = obj.TYPE_REG
				p.To.Reg = loong32r.REG_R4 + int16(hi)
				p = s.Prog(loong32r.AMOVW)
				p.From.Type = obj.TYPE_CONST
				p.From.Offset = int64(int32(c))
				p.To.Type = obj.TYPE_REG
				p.To.Reg = loong32r.REG_R4 + int16(lo)
			}
		case ssa.OpLOONG32RLoweredPanicBoundsCC:
			c := v.Aux.(ssa.PanicBoundsCC).Cx
			if c >= 0 && c <= abi.BoundsMaxConst {
				xVal = int(c)
			} else if signed && int64(int32(c)) == c || !signed && int64(uint32(c)) == c {
				// Move constant to a register
				xIsReg = true
				p := s.Prog(loong32r.AMOVW)
				p.From.Type = obj.TYPE_CONST
				p.From.Offset = c
				p.To.Type = obj.TYPE_REG
				p.To.Reg = loong32r.REG_R4 + int16(xVal)
			} else {
				// Move constant to two registers
				extend = true
				xIsReg = true
				hi := 0
				lo := 1
				xVal = hi<<2 + lo
				p := s.Prog(loong32r.AMOVW)
				p.From.Type = obj.TYPE_CONST
				p.From.Offset = c >> 32
				p.To.Type = obj.TYPE_REG
				p.To.Reg = loong32r.REG_R4 + int16(hi)
				p = s.Prog(loong32r.AMOVW)
				p.From.Type = obj.TYPE_CONST
				p.From.Offset = int64(int32(c))
				p.To.Type = obj.TYPE_REG
				p.To.Reg = loong32r.REG_R4 + int16(lo)
			}
			c = v.Aux.(ssa.PanicBoundsCC).Cy
			if c >= 0 && c <= abi.BoundsMaxConst {
				yVal = int(c)
			} else {
				// Move constant to a register
				yIsReg = true
				yVal = 2
				p := s.Prog(loong32r.AMOVW)
				p.From.Type = obj.TYPE_CONST
				p.From.Offset = c
				p.To.Type = obj.TYPE_REG
				p.To.Reg = loong32r.REG_R4 + int16(yVal)
			}
		}
		c := abi.BoundsEncode(code, signed, xIsReg, yIsReg, xVal, yVal)

		p := s.Prog(obj.APCDATA)
		p.From.SetConst(abi.PCDATA_PanicBounds)
		p.To.SetConst(int64(c))
		p = s.Prog(obj.ACALL)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		if extend {
			p.To.Sym = ir.Syms.PanicExtend
		} else {
			p.To.Sym = ir.Syms.PanicBounds
		}

	case ssa.OpLOONG32RLoweredAtomicLoad8,
		ssa.OpLOONG32RLoweredAtomicLoad32:
		dbar(s)

		var op obj.As
		switch v.Op {
		case ssa.OpLOONG32RLoweredAtomicLoad8:
			op = loong32r.AMOVB
		case ssa.OpLOONG32RLoweredAtomicLoad32:
			op = loong32r.AMOVW
		}
		p := s.Prog(op)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg0()

		dbar(s)
	case ssa.OpLOONG32RLoweredAtomicStore8,
		ssa.OpLOONG32RLoweredAtomicStore32:
		dbar(s)

		var op obj.As
		switch v.Op {
		case ssa.OpLOONG32RLoweredAtomicStore8:
			op = loong32r.AMOVB
		case ssa.OpLOONG32RLoweredAtomicStore32:
			op = loong32r.AMOVW
		}
		p := s.Prog(op)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg()
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg()

		dbar(s)
	case ssa.OpLOONG32RLoweredAtomicStorezero:
		dbar(s)

		p := s.Prog(loong32r.AMOVW)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = loong32r.REGZERO
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg()

		dbar(s)
	case ssa.OpLOONG32RLoweredAtomicExchange:
		// SYNC
		// MOVW Rarg1, Rtmp
		// LL	(Rarg0), Rout
		// SC	Rtmp, (Rarg0)
		// BEQ	Rtmp, -3(PC)
		// SYNC
		dbar(s)

		p := s.Prog(loong32r.AMOVW)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = loong32r.REGTMP

		p1 := s.Prog(loong32r.ALL)
		p1.From.Type = obj.TYPE_MEM
		p1.From.Reg = v.Args[0].Reg()
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = v.Reg0()

		p2 := s.Prog(loong32r.ASC)
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = loong32r.REGTMP
		p2.To.Type = obj.TYPE_MEM
		p2.To.Reg = v.Args[0].Reg()

		p3 := s.Prog(loong32r.ABEQ)
		p3.From.Type = obj.TYPE_REG
		p3.From.Reg = loong32r.REGTMP
		p3.To.Type = obj.TYPE_BRANCH
		p3.To.SetTarget(p)

		dbar(s)
	case ssa.OpLOONG32RLoweredAtomicAdd:
		// SYNC
		// LL	(Rarg0), Rout
		// ADDU Rarg1, Rout, Rtmp
		// SC	Rtmp, (Rarg0)
		// BEQ	Rtmp, -3(PC)
		// SYNC
		// ADDU Rarg1, Rout
		dbar(s)

		p := s.Prog(loong32r.ALL)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg0()

		p1 := s.Prog(loong32r.AADD)
		p1.From.Type = obj.TYPE_REG
		p1.From.Reg = v.Args[1].Reg()
		p1.Reg = v.Reg0()
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = loong32r.REGTMP

		p2 := s.Prog(loong32r.ASC)
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = loong32r.REGTMP
		p2.To.Type = obj.TYPE_MEM
		p2.To.Reg = v.Args[0].Reg()

		p3 := s.Prog(loong32r.ABEQ)
		p3.From.Type = obj.TYPE_REG
		p3.From.Reg = loong32r.REGTMP
		p3.To.Type = obj.TYPE_BRANCH
		p3.To.SetTarget(p)

		dbar(s)

		p4 := s.Prog(loong32r.AADD)
		p4.From.Type = obj.TYPE_REG
		p4.From.Reg = v.Args[1].Reg()
		p4.Reg = v.Reg0()
		p4.To.Type = obj.TYPE_REG
		p4.To.Reg = v.Reg0()

	case ssa.OpLOONG32RLoweredAtomicAddconst:
		// SYNC
		// LL	(Rarg0), Rout
		// ADDU $auxInt, Rout, Rtmp
		// SC	Rtmp, (Rarg0)
		// BEQ	Rtmp, -3(PC)
		// SYNC
		// ADDU $auxInt, Rout
		dbar(s)

		p := s.Prog(loong32r.ALL)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg0()

		p1 := s.Prog(loong32r.AADD)
		p1.From.Type = obj.TYPE_CONST
		p1.From.Offset = v.AuxInt
		p1.Reg = v.Reg0()
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = loong32r.REGTMP

		p2 := s.Prog(loong32r.ASC)
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = loong32r.REGTMP
		p2.To.Type = obj.TYPE_MEM
		p2.To.Reg = v.Args[0].Reg()

		p3 := s.Prog(loong32r.ABEQ)
		p3.From.Type = obj.TYPE_REG
		p3.From.Reg = loong32r.REGTMP
		p3.To.Type = obj.TYPE_BRANCH
		p3.To.SetTarget(p)

		dbar(s)

		p4 := s.Prog(loong32r.AADD)
		p4.From.Type = obj.TYPE_CONST
		p4.From.Offset = v.AuxInt
		p4.Reg = v.Reg0()
		p4.To.Type = obj.TYPE_REG
		p4.To.Reg = v.Reg0()

	case ssa.OpLOONG32RLoweredAtomicAnd,
		ssa.OpLOONG32RLoweredAtomicOr:
		// SYNC
		// LL	(Rarg0), Rtmp
		// AND/OR	Rarg1, Rtmp
		// SC	Rtmp, (Rarg0)
		// BEQ	Rtmp, -3(PC)
		// SYNC
		dbar(s)

		p := s.Prog(loong32r.ALL)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = loong32r.REGTMP

		p1 := s.Prog(v.Op.Asm())
		p1.From.Type = obj.TYPE_REG
		p1.From.Reg = v.Args[1].Reg()
		p1.Reg = loong32r.REGTMP
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = loong32r.REGTMP

		p2 := s.Prog(loong32r.ASC)
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = loong32r.REGTMP
		p2.To.Type = obj.TYPE_MEM
		p2.To.Reg = v.Args[0].Reg()

		p3 := s.Prog(loong32r.ABEQ)
		p3.From.Type = obj.TYPE_REG
		p3.From.Reg = loong32r.REGTMP
		p3.To.Type = obj.TYPE_BRANCH
		p3.To.SetTarget(p)

		dbar(s)

	case ssa.OpLOONG32RLoweredAtomicCas:
		// MOVW $0, Rout
		// SYNC
		// LL	(Rarg0), Rtmp
		// BNE	Rtmp, Rarg1, 4(PC)
		// MOVW Rarg2, Rout
		// SC	Rout, (Rarg0)
		// BEQ	Rout, -4(PC)
		// SYNC
		p := s.Prog(loong32r.AMOVW)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = loong32r.REGZERO
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg0()

		dbar(s)

		p1 := s.Prog(loong32r.ALL)
		p1.From.Type = obj.TYPE_MEM
		p1.From.Reg = v.Args[0].Reg()
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = loong32r.REGTMP

		p2 := s.Prog(loong32r.ABNE)
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = v.Args[1].Reg()
		p2.Reg = loong32r.REGTMP
		p2.To.Type = obj.TYPE_BRANCH

		p3 := s.Prog(loong32r.AMOVW)
		p3.From.Type = obj.TYPE_REG
		p3.From.Reg = v.Args[2].Reg()
		p3.To.Type = obj.TYPE_REG
		p3.To.Reg = v.Reg0()

		p4 := s.Prog(loong32r.ASC)
		p4.From.Type = obj.TYPE_REG
		p4.From.Reg = v.Reg0()
		p4.To.Type = obj.TYPE_MEM
		p4.To.Reg = v.Args[0].Reg()

		p5 := s.Prog(loong32r.ABEQ)
		p5.From.Type = obj.TYPE_REG
		p5.From.Reg = v.Reg0()
		p5.To.Type = obj.TYPE_BRANCH
		p5.To.SetTarget(p1)

		// Both the successful and compare-failed paths must execute the
		// trailing barrier before later memory operations can proceed.
		post := dbar(s)
		p2.To.SetTarget(post)

	case ssa.OpLOONG32RLoweredNilCheck:
		// Issue a load which will fault if arg is nil.
		p := s.Prog(loong32r.AMOVB)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg()
		ssagen.AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = loong32r.REGTMP
		if logopt.Enabled() {
			logopt.LogOpt(v.Pos, "nilcheck", "genssa", v.Block.Func.Name)
		}
		if base.Debug.Nil != 0 && v.Pos.Line() > 1 { // v.Pos.Line()==1 in generated wrappers
			base.WarnfAt(v.Pos, "generated nil check")
		}
	case ssa.OpLOONG32RFPFlagTrue,
		ssa.OpLOONG32RFPFlagFalse:
		// Materialize FCC0 using a short conditional branch.
		branch := loong32r.ABFPF
		if v.Op == ssa.OpLOONG32RFPFlagFalse {
			branch = loong32r.ABFPT
		}
		p := s.Prog(loong32r.AMOVW)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = loong32r.REGZERO
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
		p1 := s.Prog(branch)
		p1.To.Type = obj.TYPE_BRANCH
		p2 := s.Prog(loong32r.AMOVW)
		p2.From.SetConst(1)
		p2.To.Type = obj.TYPE_REG
		p2.To.Reg = v.Reg()
		done := s.Prog(obj.ANOP)
		p1.To.SetTarget(done)

	case ssa.OpLOONG32RLoweredGetClosurePtr:
		// Closure pointer is R29 (loong32r.REGCTXT).
		ssagen.CheckLoweredGetClosurePtr(v)
	case ssa.OpLOONG32RLoweredGetCallerSP:
		// caller's SP is FixedFrameSize below the address of the first arg
		p := s.Prog(loong32r.AMOVW)
		p.From.Type = obj.TYPE_ADDR
		p.From.Offset = -base.Ctxt.Arch.FixedFrameSize
		p.From.Name = obj.NAME_PARAM
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpLOONG32RLoweredGetCallerPC:
		p := s.Prog(obj.AGETCALLERPC)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()
	case ssa.OpLOONG32RLoweredPubBarrier:
		dbar(s)
	case ssa.OpClobber, ssa.OpClobberReg:
		// TODO: implement for clobberdead experiment. Nop is ok for now.
	default:
		v.Fatalf("genValue not implemented: %s", v.LongString())
	}
}

var blockJump = map[ssa.BlockKind]struct {
	asm, invasm obj.As
}{
	ssa.BlockLOONG32REQ:  {loong32r.ABEQ, loong32r.ABNE},
	ssa.BlockLOONG32RNE:  {loong32r.ABNE, loong32r.ABEQ},
	ssa.BlockLOONG32RLTZ: {loong32r.ABLTZ, loong32r.ABGEZ},
	ssa.BlockLOONG32RGEZ: {loong32r.ABGEZ, loong32r.ABLTZ},
	ssa.BlockLOONG32RLEZ: {loong32r.ABLEZ, loong32r.ABGTZ},
	ssa.BlockLOONG32RGTZ: {loong32r.ABGTZ, loong32r.ABLEZ},
	ssa.BlockLOONG32RFPT: {loong32r.ABFPT, loong32r.ABFPF},
	ssa.BlockLOONG32RFPF: {loong32r.ABFPF, loong32r.ABFPT},
}

func ssaGenBlock(s *ssagen.State, b, next *ssa.Block) {
	switch b.Kind {
	case ssa.BlockPlain, ssa.BlockDefer:
		if b.Succs[0].Block() != next {
			p := s.Prog(obj.AJMP)
			p.To.Type = obj.TYPE_BRANCH
			s.Branches = append(s.Branches, ssagen.Branch{P: p, B: b.Succs[0].Block()})
		}
	case ssa.BlockExit, ssa.BlockRetJmp:
	case ssa.BlockRet:
		s.Prog(obj.ARET)
	case ssa.BlockLOONG32REQ, ssa.BlockLOONG32RNE,
		ssa.BlockLOONG32RLTZ, ssa.BlockLOONG32RGEZ,
		ssa.BlockLOONG32RLEZ, ssa.BlockLOONG32RGTZ,
		ssa.BlockLOONG32RFPT, ssa.BlockLOONG32RFPF:
		jmp := blockJump[b.Kind]
		var p *obj.Prog
		switch next {
		case b.Succs[0].Block():
			p = s.Br(jmp.invasm, b.Succs[1].Block())
		case b.Succs[1].Block():
			p = s.Br(jmp.asm, b.Succs[0].Block())
		default:
			if b.Likely != ssa.BranchUnlikely {
				p = s.Br(jmp.asm, b.Succs[0].Block())
				s.Br(obj.AJMP, b.Succs[1].Block())
			} else {
				p = s.Br(jmp.invasm, b.Succs[1].Block())
				s.Br(obj.AJMP, b.Succs[0].Block())
			}
		}
		if !b.Controls[0].Type.IsFlags() {
			p.From.Type = obj.TYPE_REG
			p.From.Reg = b.Controls[0].Reg()
		}
	default:
		b.Fatalf("branch not implemented: %s", b.LongString())
	}
}
