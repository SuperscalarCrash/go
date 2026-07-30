// Code generated from _gen/LOONG32R.rules using 'go generate'; DO NOT EDIT.

package ssa

import "cmd/compile/internal/types"

func rewriteValueLOONG32R(v *Value) bool {
	switch v.Op {
	case OpAbs:
		v.Op = OpLOONG32RABSD
		return true
	case OpAdd16:
		v.Op = OpLOONG32RADD
		return true
	case OpAdd32:
		v.Op = OpLOONG32RADD
		return true
	case OpAdd32F:
		v.Op = OpLOONG32RADDF
		return true
	case OpAdd32withcarry:
		return rewriteValueLOONG32R_OpAdd32withcarry(v)
	case OpAdd64F:
		v.Op = OpLOONG32RADDD
		return true
	case OpAdd8:
		v.Op = OpLOONG32RADD
		return true
	case OpAddPtr:
		v.Op = OpLOONG32RADD
		return true
	case OpAddr:
		return rewriteValueLOONG32R_OpAddr(v)
	case OpAnd16:
		v.Op = OpLOONG32RAND
		return true
	case OpAnd32:
		v.Op = OpLOONG32RAND
		return true
	case OpAnd8:
		v.Op = OpLOONG32RAND
		return true
	case OpAndB:
		v.Op = OpLOONG32RAND
		return true
	case OpAtomicAdd32:
		v.Op = OpLOONG32RLoweredAtomicAdd
		return true
	case OpAtomicAnd32:
		v.Op = OpLOONG32RLoweredAtomicAnd
		return true
	case OpAtomicAnd8:
		return rewriteValueLOONG32R_OpAtomicAnd8(v)
	case OpAtomicCompareAndSwap32:
		v.Op = OpLOONG32RLoweredAtomicCas
		return true
	case OpAtomicExchange32:
		v.Op = OpLOONG32RLoweredAtomicExchange
		return true
	case OpAtomicLoad32:
		v.Op = OpLOONG32RLoweredAtomicLoad32
		return true
	case OpAtomicLoad8:
		v.Op = OpLOONG32RLoweredAtomicLoad8
		return true
	case OpAtomicLoadPtr:
		v.Op = OpLOONG32RLoweredAtomicLoad32
		return true
	case OpAtomicOr32:
		v.Op = OpLOONG32RLoweredAtomicOr
		return true
	case OpAtomicOr8:
		return rewriteValueLOONG32R_OpAtomicOr8(v)
	case OpAtomicStore32:
		v.Op = OpLOONG32RLoweredAtomicStore32
		return true
	case OpAtomicStore8:
		v.Op = OpLOONG32RLoweredAtomicStore8
		return true
	case OpAtomicStorePtrNoWB:
		v.Op = OpLOONG32RLoweredAtomicStore32
		return true
	case OpAvg32u:
		return rewriteValueLOONG32R_OpAvg32u(v)
	case OpBitLen16:
		return rewriteValueLOONG32R_OpBitLen16(v)
	case OpBitLen32:
		return rewriteValueLOONG32R_OpBitLen32(v)
	case OpBitLen8:
		return rewriteValueLOONG32R_OpBitLen8(v)
	case OpClosureCall:
		v.Op = OpLOONG32RCALLclosure
		return true
	case OpCom16:
		return rewriteValueLOONG32R_OpCom16(v)
	case OpCom32:
		return rewriteValueLOONG32R_OpCom32(v)
	case OpCom8:
		return rewriteValueLOONG32R_OpCom8(v)
	case OpConst16:
		return rewriteValueLOONG32R_OpConst16(v)
	case OpConst32:
		return rewriteValueLOONG32R_OpConst32(v)
	case OpConst32F:
		v.Op = OpLOONG32RMOVFconst
		return true
	case OpConst64F:
		v.Op = OpLOONG32RMOVDconst
		return true
	case OpConst8:
		return rewriteValueLOONG32R_OpConst8(v)
	case OpConstBool:
		return rewriteValueLOONG32R_OpConstBool(v)
	case OpConstNil:
		return rewriteValueLOONG32R_OpConstNil(v)
	case OpCtz16:
		return rewriteValueLOONG32R_OpCtz16(v)
	case OpCtz16NonZero:
		v.Op = OpCtz32
		return true
	case OpCtz32:
		return rewriteValueLOONG32R_OpCtz32(v)
	case OpCtz32NonZero:
		v.Op = OpCtz32
		return true
	case OpCtz8:
		return rewriteValueLOONG32R_OpCtz8(v)
	case OpCtz8NonZero:
		v.Op = OpCtz32
		return true
	case OpCvt32Fto32:
		v.Op = OpLOONG32RTRUNCFW
		return true
	case OpCvt32Fto64F:
		v.Op = OpLOONG32RMOVFD
		return true
	case OpCvt32to32F:
		v.Op = OpLOONG32RMOVWF
		return true
	case OpCvt32to64F:
		v.Op = OpLOONG32RMOVWD
		return true
	case OpCvt64Fto32:
		v.Op = OpLOONG32RTRUNCDW
		return true
	case OpCvt64Fto32F:
		v.Op = OpLOONG32RMOVDF
		return true
	case OpCvtBoolToUint8:
		v.Op = OpCopy
		return true
	case OpDiv16:
		return rewriteValueLOONG32R_OpDiv16(v)
	case OpDiv16u:
		return rewriteValueLOONG32R_OpDiv16u(v)
	case OpDiv32:
		return rewriteValueLOONG32R_OpDiv32(v)
	case OpDiv32F:
		v.Op = OpLOONG32RDIVF
		return true
	case OpDiv32u:
		return rewriteValueLOONG32R_OpDiv32u(v)
	case OpDiv64F:
		v.Op = OpLOONG32RDIVD
		return true
	case OpDiv8:
		return rewriteValueLOONG32R_OpDiv8(v)
	case OpDiv8u:
		return rewriteValueLOONG32R_OpDiv8u(v)
	case OpEq16:
		return rewriteValueLOONG32R_OpEq16(v)
	case OpEq32:
		return rewriteValueLOONG32R_OpEq32(v)
	case OpEq32F:
		return rewriteValueLOONG32R_OpEq32F(v)
	case OpEq64F:
		return rewriteValueLOONG32R_OpEq64F(v)
	case OpEq8:
		return rewriteValueLOONG32R_OpEq8(v)
	case OpEqB:
		return rewriteValueLOONG32R_OpEqB(v)
	case OpEqPtr:
		return rewriteValueLOONG32R_OpEqPtr(v)
	case OpGetCallerPC:
		v.Op = OpLOONG32RLoweredGetCallerPC
		return true
	case OpGetCallerSP:
		v.Op = OpLOONG32RLoweredGetCallerSP
		return true
	case OpGetClosurePtr:
		v.Op = OpLOONG32RLoweredGetClosurePtr
		return true
	case OpHmul32:
		return rewriteValueLOONG32R_OpHmul32(v)
	case OpHmul32u:
		return rewriteValueLOONG32R_OpHmul32u(v)
	case OpInterCall:
		v.Op = OpLOONG32RCALLinter
		return true
	case OpIsInBounds:
		return rewriteValueLOONG32R_OpIsInBounds(v)
	case OpIsNonNil:
		return rewriteValueLOONG32R_OpIsNonNil(v)
	case OpIsSliceInBounds:
		return rewriteValueLOONG32R_OpIsSliceInBounds(v)
	case OpLOONG32RADD:
		return rewriteValueLOONG32R_OpLOONG32RADD(v)
	case OpLOONG32RADDconst:
		return rewriteValueLOONG32R_OpLOONG32RADDconst(v)
	case OpLOONG32RAND:
		return rewriteValueLOONG32R_OpLOONG32RAND(v)
	case OpLOONG32RANDconst:
		return rewriteValueLOONG32R_OpLOONG32RANDconst(v)
	case OpLOONG32RCMOVZ:
		return rewriteValueLOONG32R_OpLOONG32RCMOVZ(v)
	case OpLOONG32RCMOVZzero:
		return rewriteValueLOONG32R_OpLOONG32RCMOVZzero(v)
	case OpLOONG32RLoweredAtomicAdd:
		return rewriteValueLOONG32R_OpLOONG32RLoweredAtomicAdd(v)
	case OpLOONG32RLoweredAtomicStore32:
		return rewriteValueLOONG32R_OpLOONG32RLoweredAtomicStore32(v)
	case OpLOONG32RLoweredPanicBoundsRC:
		return rewriteValueLOONG32R_OpLOONG32RLoweredPanicBoundsRC(v)
	case OpLOONG32RLoweredPanicBoundsRR:
		return rewriteValueLOONG32R_OpLOONG32RLoweredPanicBoundsRR(v)
	case OpLOONG32RLoweredPanicExtendRC:
		return rewriteValueLOONG32R_OpLOONG32RLoweredPanicExtendRC(v)
	case OpLOONG32RLoweredPanicExtendRR:
		return rewriteValueLOONG32R_OpLOONG32RLoweredPanicExtendRR(v)
	case OpLOONG32RMOVBUload:
		return rewriteValueLOONG32R_OpLOONG32RMOVBUload(v)
	case OpLOONG32RMOVBUreg:
		return rewriteValueLOONG32R_OpLOONG32RMOVBUreg(v)
	case OpLOONG32RMOVBload:
		return rewriteValueLOONG32R_OpLOONG32RMOVBload(v)
	case OpLOONG32RMOVBreg:
		return rewriteValueLOONG32R_OpLOONG32RMOVBreg(v)
	case OpLOONG32RMOVBstore:
		return rewriteValueLOONG32R_OpLOONG32RMOVBstore(v)
	case OpLOONG32RMOVBstorezero:
		return rewriteValueLOONG32R_OpLOONG32RMOVBstorezero(v)
	case OpLOONG32RMOVDload:
		return rewriteValueLOONG32R_OpLOONG32RMOVDload(v)
	case OpLOONG32RMOVDstore:
		return rewriteValueLOONG32R_OpLOONG32RMOVDstore(v)
	case OpLOONG32RMOVFload:
		return rewriteValueLOONG32R_OpLOONG32RMOVFload(v)
	case OpLOONG32RMOVFstore:
		return rewriteValueLOONG32R_OpLOONG32RMOVFstore(v)
	case OpLOONG32RMOVHUload:
		return rewriteValueLOONG32R_OpLOONG32RMOVHUload(v)
	case OpLOONG32RMOVHUreg:
		return rewriteValueLOONG32R_OpLOONG32RMOVHUreg(v)
	case OpLOONG32RMOVHload:
		return rewriteValueLOONG32R_OpLOONG32RMOVHload(v)
	case OpLOONG32RMOVHreg:
		return rewriteValueLOONG32R_OpLOONG32RMOVHreg(v)
	case OpLOONG32RMOVHstore:
		return rewriteValueLOONG32R_OpLOONG32RMOVHstore(v)
	case OpLOONG32RMOVHstorezero:
		return rewriteValueLOONG32R_OpLOONG32RMOVHstorezero(v)
	case OpLOONG32RMOVWload:
		return rewriteValueLOONG32R_OpLOONG32RMOVWload(v)
	case OpLOONG32RMOVWnop:
		return rewriteValueLOONG32R_OpLOONG32RMOVWnop(v)
	case OpLOONG32RMOVWreg:
		return rewriteValueLOONG32R_OpLOONG32RMOVWreg(v)
	case OpLOONG32RMOVWstore:
		return rewriteValueLOONG32R_OpLOONG32RMOVWstore(v)
	case OpLOONG32RMOVWstorezero:
		return rewriteValueLOONG32R_OpLOONG32RMOVWstorezero(v)
	case OpLOONG32RMUL:
		return rewriteValueLOONG32R_OpLOONG32RMUL(v)
	case OpLOONG32RNEG:
		return rewriteValueLOONG32R_OpLOONG32RNEG(v)
	case OpLOONG32RNOR:
		return rewriteValueLOONG32R_OpLOONG32RNOR(v)
	case OpLOONG32RNORconst:
		return rewriteValueLOONG32R_OpLOONG32RNORconst(v)
	case OpLOONG32ROR:
		return rewriteValueLOONG32R_OpLOONG32ROR(v)
	case OpLOONG32RORconst:
		return rewriteValueLOONG32R_OpLOONG32RORconst(v)
	case OpLOONG32RSGT:
		return rewriteValueLOONG32R_OpLOONG32RSGT(v)
	case OpLOONG32RSGTU:
		return rewriteValueLOONG32R_OpLOONG32RSGTU(v)
	case OpLOONG32RSGTUconst:
		return rewriteValueLOONG32R_OpLOONG32RSGTUconst(v)
	case OpLOONG32RSGTUzero:
		return rewriteValueLOONG32R_OpLOONG32RSGTUzero(v)
	case OpLOONG32RSGTconst:
		return rewriteValueLOONG32R_OpLOONG32RSGTconst(v)
	case OpLOONG32RSGTzero:
		return rewriteValueLOONG32R_OpLOONG32RSGTzero(v)
	case OpLOONG32RSLL:
		return rewriteValueLOONG32R_OpLOONG32RSLL(v)
	case OpLOONG32RSLLconst:
		return rewriteValueLOONG32R_OpLOONG32RSLLconst(v)
	case OpLOONG32RSRA:
		return rewriteValueLOONG32R_OpLOONG32RSRA(v)
	case OpLOONG32RSRAconst:
		return rewriteValueLOONG32R_OpLOONG32RSRAconst(v)
	case OpLOONG32RSRL:
		return rewriteValueLOONG32R_OpLOONG32RSRL(v)
	case OpLOONG32RSRLconst:
		return rewriteValueLOONG32R_OpLOONG32RSRLconst(v)
	case OpLOONG32RSUB:
		return rewriteValueLOONG32R_OpLOONG32RSUB(v)
	case OpLOONG32RSUBconst:
		return rewriteValueLOONG32R_OpLOONG32RSUBconst(v)
	case OpLOONG32RXOR:
		return rewriteValueLOONG32R_OpLOONG32RXOR(v)
	case OpLOONG32RXORconst:
		return rewriteValueLOONG32R_OpLOONG32RXORconst(v)
	case OpLeq16:
		return rewriteValueLOONG32R_OpLeq16(v)
	case OpLeq16U:
		return rewriteValueLOONG32R_OpLeq16U(v)
	case OpLeq32:
		return rewriteValueLOONG32R_OpLeq32(v)
	case OpLeq32F:
		return rewriteValueLOONG32R_OpLeq32F(v)
	case OpLeq32U:
		return rewriteValueLOONG32R_OpLeq32U(v)
	case OpLeq64F:
		return rewriteValueLOONG32R_OpLeq64F(v)
	case OpLeq8:
		return rewriteValueLOONG32R_OpLeq8(v)
	case OpLeq8U:
		return rewriteValueLOONG32R_OpLeq8U(v)
	case OpLess16:
		return rewriteValueLOONG32R_OpLess16(v)
	case OpLess16U:
		return rewriteValueLOONG32R_OpLess16U(v)
	case OpLess32:
		return rewriteValueLOONG32R_OpLess32(v)
	case OpLess32F:
		return rewriteValueLOONG32R_OpLess32F(v)
	case OpLess32U:
		return rewriteValueLOONG32R_OpLess32U(v)
	case OpLess64F:
		return rewriteValueLOONG32R_OpLess64F(v)
	case OpLess8:
		return rewriteValueLOONG32R_OpLess8(v)
	case OpLess8U:
		return rewriteValueLOONG32R_OpLess8U(v)
	case OpLoad:
		return rewriteValueLOONG32R_OpLoad(v)
	case OpLocalAddr:
		return rewriteValueLOONG32R_OpLocalAddr(v)
	case OpLsh16x16:
		return rewriteValueLOONG32R_OpLsh16x16(v)
	case OpLsh16x32:
		return rewriteValueLOONG32R_OpLsh16x32(v)
	case OpLsh16x64:
		return rewriteValueLOONG32R_OpLsh16x64(v)
	case OpLsh16x8:
		return rewriteValueLOONG32R_OpLsh16x8(v)
	case OpLsh32x16:
		return rewriteValueLOONG32R_OpLsh32x16(v)
	case OpLsh32x32:
		return rewriteValueLOONG32R_OpLsh32x32(v)
	case OpLsh32x64:
		return rewriteValueLOONG32R_OpLsh32x64(v)
	case OpLsh32x8:
		return rewriteValueLOONG32R_OpLsh32x8(v)
	case OpLsh8x16:
		return rewriteValueLOONG32R_OpLsh8x16(v)
	case OpLsh8x32:
		return rewriteValueLOONG32R_OpLsh8x32(v)
	case OpLsh8x64:
		return rewriteValueLOONG32R_OpLsh8x64(v)
	case OpLsh8x8:
		return rewriteValueLOONG32R_OpLsh8x8(v)
	case OpMod16:
		return rewriteValueLOONG32R_OpMod16(v)
	case OpMod16u:
		return rewriteValueLOONG32R_OpMod16u(v)
	case OpMod32:
		return rewriteValueLOONG32R_OpMod32(v)
	case OpMod32u:
		return rewriteValueLOONG32R_OpMod32u(v)
	case OpMod8:
		return rewriteValueLOONG32R_OpMod8(v)
	case OpMod8u:
		return rewriteValueLOONG32R_OpMod8u(v)
	case OpMove:
		return rewriteValueLOONG32R_OpMove(v)
	case OpMul16:
		v.Op = OpLOONG32RMUL
		return true
	case OpMul32:
		v.Op = OpLOONG32RMUL
		return true
	case OpMul32F:
		v.Op = OpLOONG32RMULF
		return true
	case OpMul32uhilo:
		v.Op = OpLOONG32RMULTU
		return true
	case OpMul64F:
		v.Op = OpLOONG32RMULD
		return true
	case OpMul8:
		v.Op = OpLOONG32RMUL
		return true
	case OpNeg16:
		v.Op = OpLOONG32RNEG
		return true
	case OpNeg32:
		v.Op = OpLOONG32RNEG
		return true
	case OpNeg32F:
		v.Op = OpLOONG32RNEGF
		return true
	case OpNeg64F:
		v.Op = OpLOONG32RNEGD
		return true
	case OpNeg8:
		v.Op = OpLOONG32RNEG
		return true
	case OpNeq16:
		return rewriteValueLOONG32R_OpNeq16(v)
	case OpNeq32:
		return rewriteValueLOONG32R_OpNeq32(v)
	case OpNeq32F:
		return rewriteValueLOONG32R_OpNeq32F(v)
	case OpNeq64F:
		return rewriteValueLOONG32R_OpNeq64F(v)
	case OpNeq8:
		return rewriteValueLOONG32R_OpNeq8(v)
	case OpNeqB:
		v.Op = OpLOONG32RXOR
		return true
	case OpNeqPtr:
		return rewriteValueLOONG32R_OpNeqPtr(v)
	case OpNilCheck:
		v.Op = OpLOONG32RLoweredNilCheck
		return true
	case OpNot:
		return rewriteValueLOONG32R_OpNot(v)
	case OpOffPtr:
		return rewriteValueLOONG32R_OpOffPtr(v)
	case OpOr16:
		v.Op = OpLOONG32ROR
		return true
	case OpOr32:
		v.Op = OpLOONG32ROR
		return true
	case OpOr8:
		v.Op = OpLOONG32ROR
		return true
	case OpOrB:
		v.Op = OpLOONG32ROR
		return true
	case OpPanicBounds:
		v.Op = OpLOONG32RLoweredPanicBoundsRR
		return true
	case OpPanicExtend:
		v.Op = OpLOONG32RLoweredPanicExtendRR
		return true
	case OpPubBarrier:
		v.Op = OpLOONG32RLoweredPubBarrier
		return true
	case OpRotateLeft16:
		return rewriteValueLOONG32R_OpRotateLeft16(v)
	case OpRotateLeft32:
		return rewriteValueLOONG32R_OpRotateLeft32(v)
	case OpRotateLeft64:
		return rewriteValueLOONG32R_OpRotateLeft64(v)
	case OpRotateLeft8:
		return rewriteValueLOONG32R_OpRotateLeft8(v)
	case OpRound32F:
		v.Op = OpCopy
		return true
	case OpRound64F:
		v.Op = OpCopy
		return true
	case OpRsh16Ux16:
		return rewriteValueLOONG32R_OpRsh16Ux16(v)
	case OpRsh16Ux32:
		return rewriteValueLOONG32R_OpRsh16Ux32(v)
	case OpRsh16Ux64:
		return rewriteValueLOONG32R_OpRsh16Ux64(v)
	case OpRsh16Ux8:
		return rewriteValueLOONG32R_OpRsh16Ux8(v)
	case OpRsh16x16:
		return rewriteValueLOONG32R_OpRsh16x16(v)
	case OpRsh16x32:
		return rewriteValueLOONG32R_OpRsh16x32(v)
	case OpRsh16x64:
		return rewriteValueLOONG32R_OpRsh16x64(v)
	case OpRsh16x8:
		return rewriteValueLOONG32R_OpRsh16x8(v)
	case OpRsh32Ux16:
		return rewriteValueLOONG32R_OpRsh32Ux16(v)
	case OpRsh32Ux32:
		return rewriteValueLOONG32R_OpRsh32Ux32(v)
	case OpRsh32Ux64:
		return rewriteValueLOONG32R_OpRsh32Ux64(v)
	case OpRsh32Ux8:
		return rewriteValueLOONG32R_OpRsh32Ux8(v)
	case OpRsh32x16:
		return rewriteValueLOONG32R_OpRsh32x16(v)
	case OpRsh32x32:
		return rewriteValueLOONG32R_OpRsh32x32(v)
	case OpRsh32x64:
		return rewriteValueLOONG32R_OpRsh32x64(v)
	case OpRsh32x8:
		return rewriteValueLOONG32R_OpRsh32x8(v)
	case OpRsh8Ux16:
		return rewriteValueLOONG32R_OpRsh8Ux16(v)
	case OpRsh8Ux32:
		return rewriteValueLOONG32R_OpRsh8Ux32(v)
	case OpRsh8Ux64:
		return rewriteValueLOONG32R_OpRsh8Ux64(v)
	case OpRsh8Ux8:
		return rewriteValueLOONG32R_OpRsh8Ux8(v)
	case OpRsh8x16:
		return rewriteValueLOONG32R_OpRsh8x16(v)
	case OpRsh8x32:
		return rewriteValueLOONG32R_OpRsh8x32(v)
	case OpRsh8x64:
		return rewriteValueLOONG32R_OpRsh8x64(v)
	case OpRsh8x8:
		return rewriteValueLOONG32R_OpRsh8x8(v)
	case OpSelect0:
		return rewriteValueLOONG32R_OpSelect0(v)
	case OpSelect1:
		return rewriteValueLOONG32R_OpSelect1(v)
	case OpSignExt16to32:
		v.Op = OpLOONG32RMOVHreg
		return true
	case OpSignExt8to16:
		v.Op = OpLOONG32RMOVBreg
		return true
	case OpSignExt8to32:
		v.Op = OpLOONG32RMOVBreg
		return true
	case OpSignmask:
		return rewriteValueLOONG32R_OpSignmask(v)
	case OpSlicemask:
		return rewriteValueLOONG32R_OpSlicemask(v)
	case OpSqrt:
		v.Op = OpLOONG32RSQRTD
		return true
	case OpSqrt32:
		v.Op = OpLOONG32RSQRTF
		return true
	case OpStaticCall:
		v.Op = OpLOONG32RCALLstatic
		return true
	case OpStore:
		return rewriteValueLOONG32R_OpStore(v)
	case OpSub16:
		v.Op = OpLOONG32RSUB
		return true
	case OpSub32:
		v.Op = OpLOONG32RSUB
		return true
	case OpSub32F:
		v.Op = OpLOONG32RSUBF
		return true
	case OpSub32withcarry:
		return rewriteValueLOONG32R_OpSub32withcarry(v)
	case OpSub64F:
		v.Op = OpLOONG32RSUBD
		return true
	case OpSub8:
		v.Op = OpLOONG32RSUB
		return true
	case OpSubPtr:
		v.Op = OpLOONG32RSUB
		return true
	case OpTailCall:
		v.Op = OpLOONG32RCALLtail
		return true
	case OpTrunc16to8:
		v.Op = OpCopy
		return true
	case OpTrunc32to16:
		v.Op = OpCopy
		return true
	case OpTrunc32to8:
		v.Op = OpCopy
		return true
	case OpWB:
		v.Op = OpLOONG32RLoweredWB
		return true
	case OpXor16:
		v.Op = OpLOONG32RXOR
		return true
	case OpXor32:
		v.Op = OpLOONG32RXOR
		return true
	case OpXor8:
		v.Op = OpLOONG32RXOR
		return true
	case OpZero:
		return rewriteValueLOONG32R_OpZero(v)
	case OpZeroExt16to32:
		v.Op = OpLOONG32RMOVHUreg
		return true
	case OpZeroExt8to16:
		v.Op = OpLOONG32RMOVBUreg
		return true
	case OpZeroExt8to32:
		v.Op = OpLOONG32RMOVBUreg
		return true
	case OpZeromask:
		return rewriteValueLOONG32R_OpZeromask(v)
	}
	return false
}
func rewriteValueLOONG32R_OpAdd32withcarry(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Add32withcarry <t> x y c)
	// result: (ADD c (ADD <t> x y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		c := v_2
		v.reset(OpLOONG32RADD)
		v0 := b.NewValue0(v.Pos, OpLOONG32RADD, t)
		v0.AddArg2(x, y)
		v.AddArg2(c, v0)
		return true
	}
}
func rewriteValueLOONG32R_OpAddr(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Addr {sym} base)
	// result: (MOVWaddr {sym} base)
	for {
		sym := auxToSym(v.Aux)
		base := v_0
		v.reset(OpLOONG32RMOVWaddr)
		v.Aux = symToAux(sym)
		v.AddArg(base)
		return true
	}
}
func rewriteValueLOONG32R_OpAtomicAnd8(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (AtomicAnd8 ptr val mem)
	// cond: !config.BigEndian
	// result: (LoweredAtomicAnd (AND <typ.UInt32Ptr> (MOVWconst [^3]) ptr) (OR <typ.UInt32> (SLL <typ.UInt32> (ZeroExt8to32 val) (SLLconst <typ.UInt32> [3] (ANDconst <typ.UInt32> [3] ptr))) (NORconst [0] <typ.UInt32> (SLL <typ.UInt32> (MOVWconst [0xff]) (SLLconst <typ.UInt32> [3] (ANDconst <typ.UInt32> [3] ptr))))) mem)
	for {
		ptr := v_0
		val := v_1
		mem := v_2
		if !(!config.BigEndian) {
			break
		}
		v.reset(OpLOONG32RLoweredAtomicAnd)
		v0 := b.NewValue0(v.Pos, OpLOONG32RAND, typ.UInt32Ptr)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v1.AuxInt = int32ToAuxInt(^3)
		v0.AddArg2(v1, ptr)
		v2 := b.NewValue0(v.Pos, OpLOONG32ROR, typ.UInt32)
		v3 := b.NewValue0(v.Pos, OpLOONG32RSLL, typ.UInt32)
		v4 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v4.AddArg(val)
		v5 := b.NewValue0(v.Pos, OpLOONG32RSLLconst, typ.UInt32)
		v5.AuxInt = int32ToAuxInt(3)
		v6 := b.NewValue0(v.Pos, OpLOONG32RANDconst, typ.UInt32)
		v6.AuxInt = int32ToAuxInt(3)
		v6.AddArg(ptr)
		v5.AddArg(v6)
		v3.AddArg2(v4, v5)
		v7 := b.NewValue0(v.Pos, OpLOONG32RNORconst, typ.UInt32)
		v7.AuxInt = int32ToAuxInt(0)
		v8 := b.NewValue0(v.Pos, OpLOONG32RSLL, typ.UInt32)
		v9 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v9.AuxInt = int32ToAuxInt(0xff)
		v8.AddArg2(v9, v5)
		v7.AddArg(v8)
		v2.AddArg2(v3, v7)
		v.AddArg3(v0, v2, mem)
		return true
	}
	// match: (AtomicAnd8 ptr val mem)
	// cond: config.BigEndian
	// result: (LoweredAtomicAnd (AND <typ.UInt32Ptr> (MOVWconst [^3]) ptr) (OR <typ.UInt32> (SLL <typ.UInt32> (ZeroExt8to32 val) (SLLconst <typ.UInt32> [3] (ANDconst <typ.UInt32> [3] (XORconst <typ.UInt32> [3] ptr)))) (NORconst [0] <typ.UInt32> (SLL <typ.UInt32> (MOVWconst [0xff]) (SLLconst <typ.UInt32> [3] (ANDconst <typ.UInt32> [3] (XORconst <typ.UInt32> [3] ptr)))))) mem)
	for {
		ptr := v_0
		val := v_1
		mem := v_2
		if !(config.BigEndian) {
			break
		}
		v.reset(OpLOONG32RLoweredAtomicAnd)
		v0 := b.NewValue0(v.Pos, OpLOONG32RAND, typ.UInt32Ptr)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v1.AuxInt = int32ToAuxInt(^3)
		v0.AddArg2(v1, ptr)
		v2 := b.NewValue0(v.Pos, OpLOONG32ROR, typ.UInt32)
		v3 := b.NewValue0(v.Pos, OpLOONG32RSLL, typ.UInt32)
		v4 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v4.AddArg(val)
		v5 := b.NewValue0(v.Pos, OpLOONG32RSLLconst, typ.UInt32)
		v5.AuxInt = int32ToAuxInt(3)
		v6 := b.NewValue0(v.Pos, OpLOONG32RANDconst, typ.UInt32)
		v6.AuxInt = int32ToAuxInt(3)
		v7 := b.NewValue0(v.Pos, OpLOONG32RXORconst, typ.UInt32)
		v7.AuxInt = int32ToAuxInt(3)
		v7.AddArg(ptr)
		v6.AddArg(v7)
		v5.AddArg(v6)
		v3.AddArg2(v4, v5)
		v8 := b.NewValue0(v.Pos, OpLOONG32RNORconst, typ.UInt32)
		v8.AuxInt = int32ToAuxInt(0)
		v9 := b.NewValue0(v.Pos, OpLOONG32RSLL, typ.UInt32)
		v10 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v10.AuxInt = int32ToAuxInt(0xff)
		v9.AddArg2(v10, v5)
		v8.AddArg(v9)
		v2.AddArg2(v3, v8)
		v.AddArg3(v0, v2, mem)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpAtomicOr8(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (AtomicOr8 ptr val mem)
	// cond: !config.BigEndian
	// result: (LoweredAtomicOr (AND <typ.UInt32Ptr> (MOVWconst [^3]) ptr) (SLL <typ.UInt32> (ZeroExt8to32 val) (SLLconst <typ.UInt32> [3] (ANDconst <typ.UInt32> [3] ptr))) mem)
	for {
		ptr := v_0
		val := v_1
		mem := v_2
		if !(!config.BigEndian) {
			break
		}
		v.reset(OpLOONG32RLoweredAtomicOr)
		v0 := b.NewValue0(v.Pos, OpLOONG32RAND, typ.UInt32Ptr)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v1.AuxInt = int32ToAuxInt(^3)
		v0.AddArg2(v1, ptr)
		v2 := b.NewValue0(v.Pos, OpLOONG32RSLL, typ.UInt32)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v3.AddArg(val)
		v4 := b.NewValue0(v.Pos, OpLOONG32RSLLconst, typ.UInt32)
		v4.AuxInt = int32ToAuxInt(3)
		v5 := b.NewValue0(v.Pos, OpLOONG32RANDconst, typ.UInt32)
		v5.AuxInt = int32ToAuxInt(3)
		v5.AddArg(ptr)
		v4.AddArg(v5)
		v2.AddArg2(v3, v4)
		v.AddArg3(v0, v2, mem)
		return true
	}
	// match: (AtomicOr8 ptr val mem)
	// cond: config.BigEndian
	// result: (LoweredAtomicOr (AND <typ.UInt32Ptr> (MOVWconst [^3]) ptr) (SLL <typ.UInt32> (ZeroExt8to32 val) (SLLconst <typ.UInt32> [3] (ANDconst <typ.UInt32> [3] (XORconst <typ.UInt32> [3] ptr)))) mem)
	for {
		ptr := v_0
		val := v_1
		mem := v_2
		if !(config.BigEndian) {
			break
		}
		v.reset(OpLOONG32RLoweredAtomicOr)
		v0 := b.NewValue0(v.Pos, OpLOONG32RAND, typ.UInt32Ptr)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v1.AuxInt = int32ToAuxInt(^3)
		v0.AddArg2(v1, ptr)
		v2 := b.NewValue0(v.Pos, OpLOONG32RSLL, typ.UInt32)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v3.AddArg(val)
		v4 := b.NewValue0(v.Pos, OpLOONG32RSLLconst, typ.UInt32)
		v4.AuxInt = int32ToAuxInt(3)
		v5 := b.NewValue0(v.Pos, OpLOONG32RANDconst, typ.UInt32)
		v5.AuxInt = int32ToAuxInt(3)
		v6 := b.NewValue0(v.Pos, OpLOONG32RXORconst, typ.UInt32)
		v6.AuxInt = int32ToAuxInt(3)
		v6.AddArg(ptr)
		v5.AddArg(v6)
		v4.AddArg(v5)
		v2.AddArg2(v3, v4)
		v.AddArg3(v0, v2, mem)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpAvg32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Avg32u <t> x y)
	// result: (ADD (SRLconst <t> (SUB <t> x y) [1]) y)
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpLOONG32RADD)
		v0 := b.NewValue0(v.Pos, OpLOONG32RSRLconst, t)
		v0.AuxInt = int32ToAuxInt(1)
		v1 := b.NewValue0(v.Pos, OpLOONG32RSUB, t)
		v1.AddArg2(x, y)
		v0.AddArg(v1)
		v.AddArg2(v0, y)
		return true
	}
}
func rewriteValueLOONG32R_OpBitLen16(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (BitLen16 x)
	// result: (BitLen32 (ZeroExt16to32 x))
	for {
		x := v_0
		v.reset(OpBitLen32)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpBitLen32(v *Value) bool {
	v_0 := v.Args[0]
	// match: (BitLen32 x)
	// result: (BITLEN x)
	for {
		x := v_0
		v.reset(OpLOONG32RBITLEN)
		v.AddArg(x)
		return true
	}
}
func rewriteValueLOONG32R_OpBitLen8(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (BitLen8 x)
	// result: (BitLen32 (ZeroExt8to32 x))
	for {
		x := v_0
		v.reset(OpBitLen32)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpCom16(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Com16 x)
	// result: (NORconst [0] x)
	for {
		x := v_0
		v.reset(OpLOONG32RNORconst)
		v.AuxInt = int32ToAuxInt(0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueLOONG32R_OpCom32(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Com32 x)
	// result: (NORconst [0] x)
	for {
		x := v_0
		v.reset(OpLOONG32RNORconst)
		v.AuxInt = int32ToAuxInt(0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueLOONG32R_OpCom8(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Com8 x)
	// result: (NORconst [0] x)
	for {
		x := v_0
		v.reset(OpLOONG32RNORconst)
		v.AuxInt = int32ToAuxInt(0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueLOONG32R_OpConst16(v *Value) bool {
	// match: (Const16 [val])
	// result: (MOVWconst [int32(val)])
	for {
		val := auxIntToInt16(v.AuxInt)
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(int32(val))
		return true
	}
}
func rewriteValueLOONG32R_OpConst32(v *Value) bool {
	// match: (Const32 [val])
	// result: (MOVWconst [int32(val)])
	for {
		val := auxIntToInt32(v.AuxInt)
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(int32(val))
		return true
	}
}
func rewriteValueLOONG32R_OpConst8(v *Value) bool {
	// match: (Const8 [val])
	// result: (MOVWconst [int32(val)])
	for {
		val := auxIntToInt8(v.AuxInt)
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(int32(val))
		return true
	}
}
func rewriteValueLOONG32R_OpConstBool(v *Value) bool {
	// match: (ConstBool [t])
	// result: (MOVWconst [b2i32(t)])
	for {
		t := auxIntToBool(v.AuxInt)
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(b2i32(t))
		return true
	}
}
func rewriteValueLOONG32R_OpConstNil(v *Value) bool {
	// match: (ConstNil)
	// result: (MOVWconst [0])
	for {
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
}
func rewriteValueLOONG32R_OpCtz16(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Ctz16 x)
	// result: (Ctz32 (Or32 <typ.UInt32> x (MOVWconst [1<<16])))
	for {
		x := v_0
		v.reset(OpCtz32)
		v0 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v1.AuxInt = int32ToAuxInt(1 << 16)
		v0.AddArg2(x, v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpCtz32(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Ctz32 x)
	// result: (CTZ x)
	for {
		x := v_0
		v.reset(OpLOONG32RCTZ)
		v.AddArg(x)
		return true
	}
}
func rewriteValueLOONG32R_OpCtz8(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Ctz8 x)
	// result: (Ctz32 (Or32 <typ.UInt32> x (MOVWconst [1<<8])))
	for {
		x := v_0
		v.reset(OpCtz32)
		v0 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v1.AuxInt = int32ToAuxInt(1 << 8)
		v0.AddArg2(x, v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpDiv16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16 x y)
	// result: (Select1 (DIV (SignExt16to32 x) (SignExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpLOONG32RDIV, types.NewTuple(typ.Int32, typ.Int32))
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpDiv16u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16u x y)
	// result: (Select1 (DIVU (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpLOONG32RDIVU, types.NewTuple(typ.UInt32, typ.UInt32))
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpDiv32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div32 x y)
	// result: (Select1 (DIV x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpLOONG32RDIV, types.NewTuple(typ.Int32, typ.Int32))
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpDiv32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div32u x y)
	// result: (Select1 (DIVU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpLOONG32RDIVU, types.NewTuple(typ.UInt32, typ.UInt32))
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpDiv8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8 x y)
	// result: (Select1 (DIV (SignExt8to32 x) (SignExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpLOONG32RDIV, types.NewTuple(typ.Int32, typ.Int32))
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpDiv8u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8u x y)
	// result: (Select1 (DIVU (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpLOONG32RDIVU, types.NewTuple(typ.UInt32, typ.UInt32))
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpEq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq16 x y)
	// result: (SGTUconst [1] (XOR (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RSGTUconst)
		v.AuxInt = int32ToAuxInt(1)
		v0 := b.NewValue0(v.Pos, OpLOONG32RXOR, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpEq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq32 x y)
	// result: (SGTUconst [1] (XOR x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RSGTUconst)
		v.AuxInt = int32ToAuxInt(1)
		v0 := b.NewValue0(v.Pos, OpLOONG32RXOR, typ.UInt32)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpEq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Eq32F x y)
	// result: (FPFlagTrue (CMPEQF x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpLOONG32RCMPEQF, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpEq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Eq64F x y)
	// result: (FPFlagTrue (CMPEQD x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpLOONG32RCMPEQD, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpEq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq8 x y)
	// result: (SGTUconst [1] (XOR (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RSGTUconst)
		v.AuxInt = int32ToAuxInt(1)
		v0 := b.NewValue0(v.Pos, OpLOONG32RXOR, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpEqB(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (EqB x y)
	// result: (XORconst [1] (XOR <typ.Bool> x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RXORconst)
		v.AuxInt = int32ToAuxInt(1)
		v0 := b.NewValue0(v.Pos, OpLOONG32RXOR, typ.Bool)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpEqPtr(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (EqPtr x y)
	// result: (SGTUconst [1] (XOR x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RSGTUconst)
		v.AuxInt = int32ToAuxInt(1)
		v0 := b.NewValue0(v.Pos, OpLOONG32RXOR, typ.UInt32)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpHmul32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Hmul32 x y)
	// result: (Select0 (MULT x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpLOONG32RMULT, types.NewTuple(typ.Int32, typ.Int32))
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpHmul32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Hmul32u x y)
	// result: (Select0 (MULTU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpLOONG32RMULTU, types.NewTuple(typ.UInt32, typ.UInt32))
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpIsInBounds(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (IsInBounds idx len)
	// result: (SGTU len idx)
	for {
		idx := v_0
		len := v_1
		v.reset(OpLOONG32RSGTU)
		v.AddArg2(len, idx)
		return true
	}
}
func rewriteValueLOONG32R_OpIsNonNil(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (IsNonNil ptr)
	// result: (SGTU ptr (MOVWconst [0]))
	for {
		ptr := v_0
		v.reset(OpLOONG32RSGTU)
		v0 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(0)
		v.AddArg2(ptr, v0)
		return true
	}
}
func rewriteValueLOONG32R_OpIsSliceInBounds(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (IsSliceInBounds idx len)
	// result: (XORconst [1] (SGTU idx len))
	for {
		idx := v_0
		len := v_1
		v.reset(OpLOONG32RXORconst)
		v.AuxInt = int32ToAuxInt(1)
		v0 := b.NewValue0(v.Pos, OpLOONG32RSGTU, typ.Bool)
		v0.AddArg2(idx, len)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpLOONG32RADD(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ADD x (MOVWconst <t> [c]))
	// cond: !t.IsPtr()
	// result: (ADDconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpLOONG32RMOVWconst {
				continue
			}
			t := v_1.Type
			c := auxIntToInt32(v_1.AuxInt)
			if !(!t.IsPtr()) {
				continue
			}
			v.reset(OpLOONG32RADDconst)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ADD x (NEG y))
	// result: (SUB x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpLOONG32RNEG {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpLOONG32RSUB)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RADDconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ADDconst [off1] (MOVWaddr [off2] {sym} ptr))
	// result: (MOVWaddr [off1+off2] {sym} ptr)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RMOVWaddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		v.reset(OpLOONG32RMOVWaddr)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg(ptr)
		return true
	}
	// match: (ADDconst [0] x)
	// result: x
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		x := v_0
		v.copyOf(x)
		return true
	}
	// match: (ADDconst [c] (MOVWconst [d]))
	// result: (MOVWconst [int32(c+d)])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RMOVWconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(int32(c + d))
		return true
	}
	// match: (ADDconst [c] (ADDconst [d] x))
	// result: (ADDconst [c+d] x)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RADDconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		v.reset(OpLOONG32RADDconst)
		v.AuxInt = int32ToAuxInt(c + d)
		v.AddArg(x)
		return true
	}
	// match: (ADDconst [c] (SUBconst [d] x))
	// result: (ADDconst [c-d] x)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RSUBconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		v.reset(OpLOONG32RADDconst)
		v.AuxInt = int32ToAuxInt(c - d)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RAND(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (AND x (MOVWconst [c]))
	// result: (ANDconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpLOONG32RMOVWconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			v.reset(OpLOONG32RANDconst)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (AND x x)
	// result: x
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (AND (SGTUconst [1] x) (SGTUconst [1] y))
	// result: (SGTUconst [1] (OR <x.Type> x y))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpLOONG32RSGTUconst || auxIntToInt32(v_0.AuxInt) != 1 {
				continue
			}
			x := v_0.Args[0]
			if v_1.Op != OpLOONG32RSGTUconst || auxIntToInt32(v_1.AuxInt) != 1 {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpLOONG32RSGTUconst)
			v.AuxInt = int32ToAuxInt(1)
			v0 := b.NewValue0(v.Pos, OpLOONG32ROR, x.Type)
			v0.AddArg2(x, y)
			v.AddArg(v0)
			return true
		}
		break
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RANDconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ANDconst [0] _)
	// result: (MOVWconst [0])
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	// match: (ANDconst [-1] x)
	// result: x
	for {
		if auxIntToInt32(v.AuxInt) != -1 {
			break
		}
		x := v_0
		v.copyOf(x)
		return true
	}
	// match: (ANDconst [c] (MOVWconst [d]))
	// result: (MOVWconst [c&d])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RMOVWconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(c & d)
		return true
	}
	// match: (ANDconst [c] (ANDconst [d] x))
	// result: (ANDconst [c&d] x)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RANDconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		v.reset(OpLOONG32RANDconst)
		v.AuxInt = int32ToAuxInt(c & d)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RCMOVZ(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (CMOVZ _ f (MOVWconst [0]))
	// result: f
	for {
		f := v_1
		if v_2.Op != OpLOONG32RMOVWconst || auxIntToInt32(v_2.AuxInt) != 0 {
			break
		}
		v.copyOf(f)
		return true
	}
	// match: (CMOVZ a _ (MOVWconst [c]))
	// cond: c!=0
	// result: a
	for {
		a := v_0
		if v_2.Op != OpLOONG32RMOVWconst {
			break
		}
		c := auxIntToInt32(v_2.AuxInt)
		if !(c != 0) {
			break
		}
		v.copyOf(a)
		return true
	}
	// match: (CMOVZ a (MOVWconst [0]) c)
	// result: (CMOVZzero a c)
	for {
		a := v_0
		if v_1.Op != OpLOONG32RMOVWconst || auxIntToInt32(v_1.AuxInt) != 0 {
			break
		}
		c := v_2
		v.reset(OpLOONG32RCMOVZzero)
		v.AddArg2(a, c)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RCMOVZzero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (CMOVZzero _ (MOVWconst [0]))
	// result: (MOVWconst [0])
	for {
		if v_1.Op != OpLOONG32RMOVWconst || auxIntToInt32(v_1.AuxInt) != 0 {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	// match: (CMOVZzero a (MOVWconst [c]))
	// cond: c!=0
	// result: a
	for {
		a := v_0
		if v_1.Op != OpLOONG32RMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		if !(c != 0) {
			break
		}
		v.copyOf(a)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RLoweredAtomicAdd(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (LoweredAtomicAdd ptr (MOVWconst [c]) mem)
	// cond: is16Bit(int64(c))
	// result: (LoweredAtomicAddconst [c] ptr mem)
	for {
		ptr := v_0
		if v_1.Op != OpLOONG32RMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		mem := v_2
		if !(is16Bit(int64(c))) {
			break
		}
		v.reset(OpLOONG32RLoweredAtomicAddconst)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RLoweredAtomicStore32(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (LoweredAtomicStore32 ptr (MOVWconst [0]) mem)
	// result: (LoweredAtomicStorezero ptr mem)
	for {
		ptr := v_0
		if v_1.Op != OpLOONG32RMOVWconst || auxIntToInt32(v_1.AuxInt) != 0 {
			break
		}
		mem := v_2
		v.reset(OpLOONG32RLoweredAtomicStorezero)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RLoweredPanicBoundsRC(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (LoweredPanicBoundsRC [kind] {p} (MOVWconst [c]) mem)
	// result: (LoweredPanicBoundsCC [kind] {PanicBoundsCC{Cx:int64(c), Cy:p.C}} mem)
	for {
		kind := auxIntToInt64(v.AuxInt)
		p := auxToPanicBoundsC(v.Aux)
		if v_0.Op != OpLOONG32RMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		mem := v_1
		v.reset(OpLOONG32RLoweredPanicBoundsCC)
		v.AuxInt = int64ToAuxInt(kind)
		v.Aux = panicBoundsCCToAux(PanicBoundsCC{Cx: int64(c), Cy: p.C})
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RLoweredPanicBoundsRR(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (LoweredPanicBoundsRR [kind] x (MOVWconst [c]) mem)
	// result: (LoweredPanicBoundsRC [kind] x {PanicBoundsC{C:int64(c)}} mem)
	for {
		kind := auxIntToInt64(v.AuxInt)
		x := v_0
		if v_1.Op != OpLOONG32RMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		mem := v_2
		v.reset(OpLOONG32RLoweredPanicBoundsRC)
		v.AuxInt = int64ToAuxInt(kind)
		v.Aux = panicBoundsCToAux(PanicBoundsC{C: int64(c)})
		v.AddArg2(x, mem)
		return true
	}
	// match: (LoweredPanicBoundsRR [kind] (MOVWconst [c]) y mem)
	// result: (LoweredPanicBoundsCR [kind] {PanicBoundsC{C:int64(c)}} y mem)
	for {
		kind := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpLOONG32RMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		y := v_1
		mem := v_2
		v.reset(OpLOONG32RLoweredPanicBoundsCR)
		v.AuxInt = int64ToAuxInt(kind)
		v.Aux = panicBoundsCToAux(PanicBoundsC{C: int64(c)})
		v.AddArg2(y, mem)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RLoweredPanicExtendRC(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (LoweredPanicExtendRC [kind] {p} (MOVWconst [hi]) (MOVWconst [lo]) mem)
	// result: (LoweredPanicBoundsCC [kind] {PanicBoundsCC{Cx:int64(hi)<<32+int64(uint32(lo)), Cy:p.C}} mem)
	for {
		kind := auxIntToInt64(v.AuxInt)
		p := auxToPanicBoundsC(v.Aux)
		if v_0.Op != OpLOONG32RMOVWconst {
			break
		}
		hi := auxIntToInt32(v_0.AuxInt)
		if v_1.Op != OpLOONG32RMOVWconst {
			break
		}
		lo := auxIntToInt32(v_1.AuxInt)
		mem := v_2
		v.reset(OpLOONG32RLoweredPanicBoundsCC)
		v.AuxInt = int64ToAuxInt(kind)
		v.Aux = panicBoundsCCToAux(PanicBoundsCC{Cx: int64(hi)<<32 + int64(uint32(lo)), Cy: p.C})
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RLoweredPanicExtendRR(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (LoweredPanicExtendRR [kind] hi lo (MOVWconst [c]) mem)
	// result: (LoweredPanicExtendRC [kind] hi lo {PanicBoundsC{C:int64(c)}} mem)
	for {
		kind := auxIntToInt64(v.AuxInt)
		hi := v_0
		lo := v_1
		if v_2.Op != OpLOONG32RMOVWconst {
			break
		}
		c := auxIntToInt32(v_2.AuxInt)
		mem := v_3
		v.reset(OpLOONG32RLoweredPanicExtendRC)
		v.AuxInt = int64ToAuxInt(kind)
		v.Aux = panicBoundsCToAux(PanicBoundsC{C: int64(c)})
		v.AddArg3(hi, lo, mem)
		return true
	}
	// match: (LoweredPanicExtendRR [kind] (MOVWconst [hi]) (MOVWconst [lo]) y mem)
	// result: (LoweredPanicBoundsCR [kind] {PanicBoundsC{C:int64(hi)<<32 + int64(uint32(lo))}} y mem)
	for {
		kind := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpLOONG32RMOVWconst {
			break
		}
		hi := auxIntToInt32(v_0.AuxInt)
		if v_1.Op != OpLOONG32RMOVWconst {
			break
		}
		lo := auxIntToInt32(v_1.AuxInt)
		y := v_2
		mem := v_3
		v.reset(OpLOONG32RLoweredPanicBoundsCR)
		v.AuxInt = int64ToAuxInt(kind)
		v.Aux = panicBoundsCToAux(PanicBoundsC{C: int64(hi)<<32 + int64(uint32(lo))})
		v.AddArg2(y, mem)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RMOVBUload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBUload [off1] {sym} x:(ADDconst [off2] ptr) mem)
	// cond: (is16Bit(int64(off1+off2)) || x.Uses == 1)
	// result: (MOVBUload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		x := v_0
		if x.Op != OpLOONG32RADDconst {
			break
		}
		off2 := auxIntToInt32(x.AuxInt)
		ptr := x.Args[0]
		mem := v_1
		if !(is16Bit(int64(off1+off2)) || x.Uses == 1) {
			break
		}
		v.reset(OpLOONG32RMOVBUload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBUload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVBUload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpLOONG32RMOVWaddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpLOONG32RMOVBUload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBUload [off] {sym} ptr (MOVBstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: (MOVBUreg x)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpLOONG32RMOVBstore {
			break
		}
		off2 := auxIntToInt32(v_1.AuxInt)
		sym2 := auxToSym(v_1.Aux)
		x := v_1.Args[1]
		ptr2 := v_1.Args[0]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpLOONG32RMOVBUreg)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RMOVBUreg(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (MOVBUreg x:(MOVBUload _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpLOONG32RMOVBUload {
			break
		}
		v.reset(OpLOONG32RMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBUreg x:(MOVBUreg _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpLOONG32RMOVBUreg {
			break
		}
		v.reset(OpLOONG32RMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBUreg <t> x:(MOVBload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVBUload <t> [off] {sym} ptr mem)
	for {
		t := v.Type
		x := v_0
		if x.Op != OpLOONG32RMOVBload {
			break
		}
		off := auxIntToInt32(x.AuxInt)
		sym := auxToSym(x.Aux)
		mem := x.Args[1]
		ptr := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpLOONG32RMOVBUload, t)
		v.copyOf(v0)
		v0.AuxInt = int32ToAuxInt(off)
		v0.Aux = symToAux(sym)
		v0.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBUreg (ANDconst [c] x))
	// result: (ANDconst [c&0xff] x)
	for {
		if v_0.Op != OpLOONG32RANDconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		v.reset(OpLOONG32RANDconst)
		v.AuxInt = int32ToAuxInt(c & 0xff)
		v.AddArg(x)
		return true
	}
	// match: (MOVBUreg (MOVWconst [c]))
	// result: (MOVWconst [int32(uint8(c))])
	for {
		if v_0.Op != OpLOONG32RMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(int32(uint8(c)))
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RMOVBload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBload [off1] {sym} x:(ADDconst [off2] ptr) mem)
	// cond: (is16Bit(int64(off1+off2)) || x.Uses == 1)
	// result: (MOVBload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		x := v_0
		if x.Op != OpLOONG32RADDconst {
			break
		}
		off2 := auxIntToInt32(x.AuxInt)
		ptr := x.Args[0]
		mem := v_1
		if !(is16Bit(int64(off1+off2)) || x.Uses == 1) {
			break
		}
		v.reset(OpLOONG32RMOVBload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVBload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpLOONG32RMOVWaddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpLOONG32RMOVBload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBload [off] {sym} ptr (MOVBstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: (MOVBreg x)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpLOONG32RMOVBstore {
			break
		}
		off2 := auxIntToInt32(v_1.AuxInt)
		sym2 := auxToSym(v_1.Aux)
		x := v_1.Args[1]
		ptr2 := v_1.Args[0]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpLOONG32RMOVBreg)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RMOVBreg(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (MOVBreg x:(MOVBload _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpLOONG32RMOVBload {
			break
		}
		v.reset(OpLOONG32RMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg x:(MOVBreg _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpLOONG32RMOVBreg {
			break
		}
		v.reset(OpLOONG32RMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg <t> x:(MOVBUload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVBload <t> [off] {sym} ptr mem)
	for {
		t := v.Type
		x := v_0
		if x.Op != OpLOONG32RMOVBUload {
			break
		}
		off := auxIntToInt32(x.AuxInt)
		sym := auxToSym(x.Aux)
		mem := x.Args[1]
		ptr := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpLOONG32RMOVBload, t)
		v.copyOf(v0)
		v0.AuxInt = int32ToAuxInt(off)
		v0.Aux = symToAux(sym)
		v0.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBreg (ANDconst [c] x))
	// cond: c & 0x80 == 0
	// result: (ANDconst [c&0x7f] x)
	for {
		if v_0.Op != OpLOONG32RANDconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		if !(c&0x80 == 0) {
			break
		}
		v.reset(OpLOONG32RANDconst)
		v.AuxInt = int32ToAuxInt(c & 0x7f)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg (MOVWconst [c]))
	// result: (MOVWconst [int32(int8(c))])
	for {
		if v_0.Op != OpLOONG32RMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(int32(int8(c)))
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RMOVBstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBstore [off1] {sym} x:(ADDconst [off2] ptr) val mem)
	// cond: (is16Bit(int64(off1+off2)) || x.Uses == 1)
	// result: (MOVBstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		x := v_0
		if x.Op != OpLOONG32RADDconst {
			break
		}
		off2 := auxIntToInt32(x.AuxInt)
		ptr := x.Args[0]
		val := v_1
		mem := v_2
		if !(is16Bit(int64(off1+off2)) || x.Uses == 1) {
			break
		}
		v.reset(OpLOONG32RMOVBstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVBstore [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVBstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpLOONG32RMOVWaddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpLOONG32RMOVBstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVWconst [0]) mem)
	// result: (MOVBstorezero [off] {sym} ptr mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpLOONG32RMOVWconst || auxIntToInt32(v_1.AuxInt) != 0 {
			break
		}
		mem := v_2
		v.reset(OpLOONG32RMOVBstorezero)
		v.AuxInt = int32ToAuxInt(off)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVBreg x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpLOONG32RMOVBreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpLOONG32RMOVBstore)
		v.AuxInt = int32ToAuxInt(off)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, x, mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVBUreg x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpLOONG32RMOVBUreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpLOONG32RMOVBstore)
		v.AuxInt = int32ToAuxInt(off)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, x, mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVHreg x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpLOONG32RMOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpLOONG32RMOVBstore)
		v.AuxInt = int32ToAuxInt(off)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, x, mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVHUreg x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpLOONG32RMOVHUreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpLOONG32RMOVBstore)
		v.AuxInt = int32ToAuxInt(off)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, x, mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVWreg x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpLOONG32RMOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpLOONG32RMOVBstore)
		v.AuxInt = int32ToAuxInt(off)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, x, mem)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RMOVBstorezero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBstorezero [off1] {sym} x:(ADDconst [off2] ptr) mem)
	// cond: (is16Bit(int64(off1+off2)) || x.Uses == 1)
	// result: (MOVBstorezero [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		x := v_0
		if x.Op != OpLOONG32RADDconst {
			break
		}
		off2 := auxIntToInt32(x.AuxInt)
		ptr := x.Args[0]
		mem := v_1
		if !(is16Bit(int64(off1+off2)) || x.Uses == 1) {
			break
		}
		v.reset(OpLOONG32RMOVBstorezero)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBstorezero [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVBstorezero [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpLOONG32RMOVWaddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpLOONG32RMOVBstorezero)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RMOVDload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVDload [off1] {sym} x:(ADDconst [off2] ptr) mem)
	// cond: (is16Bit(int64(off1+off2)) || x.Uses == 1)
	// result: (MOVDload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		x := v_0
		if x.Op != OpLOONG32RADDconst {
			break
		}
		off2 := auxIntToInt32(x.AuxInt)
		ptr := x.Args[0]
		mem := v_1
		if !(is16Bit(int64(off1+off2)) || x.Uses == 1) {
			break
		}
		v.reset(OpLOONG32RMOVDload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVDload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVDload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpLOONG32RMOVWaddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpLOONG32RMOVDload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVDload [off] {sym} ptr (MOVDstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: x
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpLOONG32RMOVDstore {
			break
		}
		off2 := auxIntToInt32(v_1.AuxInt)
		sym2 := auxToSym(v_1.Aux)
		x := v_1.Args[1]
		ptr2 := v_1.Args[0]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RMOVDstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVDstore [off1] {sym} x:(ADDconst [off2] ptr) val mem)
	// cond: (is16Bit(int64(off1+off2)) || x.Uses == 1)
	// result: (MOVDstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		x := v_0
		if x.Op != OpLOONG32RADDconst {
			break
		}
		off2 := auxIntToInt32(x.AuxInt)
		ptr := x.Args[0]
		val := v_1
		mem := v_2
		if !(is16Bit(int64(off1+off2)) || x.Uses == 1) {
			break
		}
		v.reset(OpLOONG32RMOVDstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVDstore [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVDstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpLOONG32RMOVWaddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpLOONG32RMOVDstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RMOVFload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVFload [off] {sym} ptr (MOVWstore [off] {sym} ptr val _))
	// result: (MOVWgpfp val)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpLOONG32RMOVWstore || auxIntToInt32(v_1.AuxInt) != off || auxToSym(v_1.Aux) != sym {
			break
		}
		val := v_1.Args[1]
		if ptr != v_1.Args[0] {
			break
		}
		v.reset(OpLOONG32RMOVWgpfp)
		v.AddArg(val)
		return true
	}
	// match: (MOVFload [off1] {sym} x:(ADDconst [off2] ptr) mem)
	// cond: (is16Bit(int64(off1+off2)) || x.Uses == 1)
	// result: (MOVFload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		x := v_0
		if x.Op != OpLOONG32RADDconst {
			break
		}
		off2 := auxIntToInt32(x.AuxInt)
		ptr := x.Args[0]
		mem := v_1
		if !(is16Bit(int64(off1+off2)) || x.Uses == 1) {
			break
		}
		v.reset(OpLOONG32RMOVFload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVFload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVFload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpLOONG32RMOVWaddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpLOONG32RMOVFload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVFload [off] {sym} ptr (MOVFstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: x
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpLOONG32RMOVFstore {
			break
		}
		off2 := auxIntToInt32(v_1.AuxInt)
		sym2 := auxToSym(v_1.Aux)
		x := v_1.Args[1]
		ptr2 := v_1.Args[0]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RMOVFstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVFstore [off] {sym} ptr (MOVWgpfp val) mem)
	// result: (MOVWstore [off] {sym} ptr val mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpLOONG32RMOVWgpfp {
			break
		}
		val := v_1.Args[0]
		mem := v_2
		v.reset(OpLOONG32RMOVWstore)
		v.AuxInt = int32ToAuxInt(off)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVFstore [off1] {sym} x:(ADDconst [off2] ptr) val mem)
	// cond: (is16Bit(int64(off1+off2)) || x.Uses == 1)
	// result: (MOVFstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		x := v_0
		if x.Op != OpLOONG32RADDconst {
			break
		}
		off2 := auxIntToInt32(x.AuxInt)
		ptr := x.Args[0]
		val := v_1
		mem := v_2
		if !(is16Bit(int64(off1+off2)) || x.Uses == 1) {
			break
		}
		v.reset(OpLOONG32RMOVFstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVFstore [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVFstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpLOONG32RMOVWaddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpLOONG32RMOVFstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RMOVHUload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHUload [off1] {sym} x:(ADDconst [off2] ptr) mem)
	// cond: (is16Bit(int64(off1+off2)) || x.Uses == 1)
	// result: (MOVHUload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		x := v_0
		if x.Op != OpLOONG32RADDconst {
			break
		}
		off2 := auxIntToInt32(x.AuxInt)
		ptr := x.Args[0]
		mem := v_1
		if !(is16Bit(int64(off1+off2)) || x.Uses == 1) {
			break
		}
		v.reset(OpLOONG32RMOVHUload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHUload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVHUload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpLOONG32RMOVWaddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpLOONG32RMOVHUload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHUload [off] {sym} ptr (MOVHstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: (MOVHUreg x)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpLOONG32RMOVHstore {
			break
		}
		off2 := auxIntToInt32(v_1.AuxInt)
		sym2 := auxToSym(v_1.Aux)
		x := v_1.Args[1]
		ptr2 := v_1.Args[0]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpLOONG32RMOVHUreg)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RMOVHUreg(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (MOVHUreg x:(MOVBUload _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpLOONG32RMOVBUload {
			break
		}
		v.reset(OpLOONG32RMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg x:(MOVHUload _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpLOONG32RMOVHUload {
			break
		}
		v.reset(OpLOONG32RMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg x:(MOVBUreg _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpLOONG32RMOVBUreg {
			break
		}
		v.reset(OpLOONG32RMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg x:(MOVHUreg _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpLOONG32RMOVHUreg {
			break
		}
		v.reset(OpLOONG32RMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg <t> x:(MOVHload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVHUload <t> [off] {sym} ptr mem)
	for {
		t := v.Type
		x := v_0
		if x.Op != OpLOONG32RMOVHload {
			break
		}
		off := auxIntToInt32(x.AuxInt)
		sym := auxToSym(x.Aux)
		mem := x.Args[1]
		ptr := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpLOONG32RMOVHUload, t)
		v.copyOf(v0)
		v0.AuxInt = int32ToAuxInt(off)
		v0.Aux = symToAux(sym)
		v0.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHUreg (ANDconst [c] x))
	// result: (ANDconst [c&0xffff] x)
	for {
		if v_0.Op != OpLOONG32RANDconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		v.reset(OpLOONG32RANDconst)
		v.AuxInt = int32ToAuxInt(c & 0xffff)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg (MOVWconst [c]))
	// result: (MOVWconst [int32(uint16(c))])
	for {
		if v_0.Op != OpLOONG32RMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(int32(uint16(c)))
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RMOVHload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHload [off1] {sym} x:(ADDconst [off2] ptr) mem)
	// cond: (is16Bit(int64(off1+off2)) || x.Uses == 1)
	// result: (MOVHload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		x := v_0
		if x.Op != OpLOONG32RADDconst {
			break
		}
		off2 := auxIntToInt32(x.AuxInt)
		ptr := x.Args[0]
		mem := v_1
		if !(is16Bit(int64(off1+off2)) || x.Uses == 1) {
			break
		}
		v.reset(OpLOONG32RMOVHload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVHload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpLOONG32RMOVWaddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpLOONG32RMOVHload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHload [off] {sym} ptr (MOVHstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: (MOVHreg x)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpLOONG32RMOVHstore {
			break
		}
		off2 := auxIntToInt32(v_1.AuxInt)
		sym2 := auxToSym(v_1.Aux)
		x := v_1.Args[1]
		ptr2 := v_1.Args[0]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpLOONG32RMOVHreg)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RMOVHreg(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (MOVHreg x:(MOVBload _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpLOONG32RMOVBload {
			break
		}
		v.reset(OpLOONG32RMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBUload _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpLOONG32RMOVBUload {
			break
		}
		v.reset(OpLOONG32RMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVHload _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpLOONG32RMOVHload {
			break
		}
		v.reset(OpLOONG32RMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBreg _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpLOONG32RMOVBreg {
			break
		}
		v.reset(OpLOONG32RMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBUreg _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpLOONG32RMOVBUreg {
			break
		}
		v.reset(OpLOONG32RMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVHreg _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpLOONG32RMOVHreg {
			break
		}
		v.reset(OpLOONG32RMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg <t> x:(MOVHUload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVHload <t> [off] {sym} ptr mem)
	for {
		t := v.Type
		x := v_0
		if x.Op != OpLOONG32RMOVHUload {
			break
		}
		off := auxIntToInt32(x.AuxInt)
		sym := auxToSym(x.Aux)
		mem := x.Args[1]
		ptr := x.Args[0]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(x.Pos, OpLOONG32RMOVHload, t)
		v.copyOf(v0)
		v0.AuxInt = int32ToAuxInt(off)
		v0.Aux = symToAux(sym)
		v0.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHreg (ANDconst [c] x))
	// cond: c & 0x8000 == 0
	// result: (ANDconst [c&0x7fff] x)
	for {
		if v_0.Op != OpLOONG32RANDconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		if !(c&0x8000 == 0) {
			break
		}
		v.reset(OpLOONG32RANDconst)
		v.AuxInt = int32ToAuxInt(c & 0x7fff)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg (MOVWconst [c]))
	// result: (MOVWconst [int32(int16(c))])
	for {
		if v_0.Op != OpLOONG32RMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(int32(int16(c)))
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RMOVHstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHstore [off1] {sym} x:(ADDconst [off2] ptr) val mem)
	// cond: (is16Bit(int64(off1+off2)) || x.Uses == 1)
	// result: (MOVHstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		x := v_0
		if x.Op != OpLOONG32RADDconst {
			break
		}
		off2 := auxIntToInt32(x.AuxInt)
		ptr := x.Args[0]
		val := v_1
		mem := v_2
		if !(is16Bit(int64(off1+off2)) || x.Uses == 1) {
			break
		}
		v.reset(OpLOONG32RMOVHstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVHstore [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVHstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpLOONG32RMOVWaddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpLOONG32RMOVHstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVWconst [0]) mem)
	// result: (MOVHstorezero [off] {sym} ptr mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpLOONG32RMOVWconst || auxIntToInt32(v_1.AuxInt) != 0 {
			break
		}
		mem := v_2
		v.reset(OpLOONG32RMOVHstorezero)
		v.AuxInt = int32ToAuxInt(off)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVHreg x) mem)
	// result: (MOVHstore [off] {sym} ptr x mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpLOONG32RMOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpLOONG32RMOVHstore)
		v.AuxInt = int32ToAuxInt(off)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, x, mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVHUreg x) mem)
	// result: (MOVHstore [off] {sym} ptr x mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpLOONG32RMOVHUreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpLOONG32RMOVHstore)
		v.AuxInt = int32ToAuxInt(off)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, x, mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVWreg x) mem)
	// result: (MOVHstore [off] {sym} ptr x mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpLOONG32RMOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpLOONG32RMOVHstore)
		v.AuxInt = int32ToAuxInt(off)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, x, mem)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RMOVHstorezero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHstorezero [off1] {sym} x:(ADDconst [off2] ptr) mem)
	// cond: (is16Bit(int64(off1+off2)) || x.Uses == 1)
	// result: (MOVHstorezero [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		x := v_0
		if x.Op != OpLOONG32RADDconst {
			break
		}
		off2 := auxIntToInt32(x.AuxInt)
		ptr := x.Args[0]
		mem := v_1
		if !(is16Bit(int64(off1+off2)) || x.Uses == 1) {
			break
		}
		v.reset(OpLOONG32RMOVHstorezero)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHstorezero [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVHstorezero [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpLOONG32RMOVWaddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpLOONG32RMOVHstorezero)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RMOVWload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWload [off] {sym} ptr (MOVFstore [off] {sym} ptr val _))
	// result: (MOVWfpgp val)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpLOONG32RMOVFstore || auxIntToInt32(v_1.AuxInt) != off || auxToSym(v_1.Aux) != sym {
			break
		}
		val := v_1.Args[1]
		if ptr != v_1.Args[0] {
			break
		}
		v.reset(OpLOONG32RMOVWfpgp)
		v.AddArg(val)
		return true
	}
	// match: (MOVWload [off1] {sym} x:(ADDconst [off2] ptr) mem)
	// cond: (is16Bit(int64(off1+off2)) || x.Uses == 1)
	// result: (MOVWload [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		x := v_0
		if x.Op != OpLOONG32RADDconst {
			break
		}
		off2 := auxIntToInt32(x.AuxInt)
		ptr := x.Args[0]
		mem := v_1
		if !(is16Bit(int64(off1+off2)) || x.Uses == 1) {
			break
		}
		v.reset(OpLOONG32RMOVWload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVWload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVWload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpLOONG32RMOVWaddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpLOONG32RMOVWload)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVWload [off] {sym} ptr (MOVWstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: x
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpLOONG32RMOVWstore {
			break
		}
		off2 := auxIntToInt32(v_1.AuxInt)
		sym2 := auxToSym(v_1.Aux)
		x := v_1.Args[1]
		ptr2 := v_1.Args[0]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RMOVWnop(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MOVWnop (MOVWconst [c]))
	// result: (MOVWconst [c])
	for {
		if v_0.Op != OpLOONG32RMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(c)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RMOVWreg(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MOVWreg x)
	// cond: x.Uses == 1
	// result: (MOVWnop x)
	for {
		x := v_0
		if !(x.Uses == 1) {
			break
		}
		v.reset(OpLOONG32RMOVWnop)
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg (MOVWconst [c]))
	// result: (MOVWconst [c])
	for {
		if v_0.Op != OpLOONG32RMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(c)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RMOVWstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWstore [off] {sym} ptr (MOVWfpgp val) mem)
	// result: (MOVFstore [off] {sym} ptr val mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpLOONG32RMOVWfpgp {
			break
		}
		val := v_1.Args[0]
		mem := v_2
		v.reset(OpLOONG32RMOVFstore)
		v.AuxInt = int32ToAuxInt(off)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVWstore [off1] {sym} x:(ADDconst [off2] ptr) val mem)
	// cond: (is16Bit(int64(off1+off2)) || x.Uses == 1)
	// result: (MOVWstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		x := v_0
		if x.Op != OpLOONG32RADDconst {
			break
		}
		off2 := auxIntToInt32(x.AuxInt)
		ptr := x.Args[0]
		val := v_1
		mem := v_2
		if !(is16Bit(int64(off1+off2)) || x.Uses == 1) {
			break
		}
		v.reset(OpLOONG32RMOVWstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVWstore [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVWstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpLOONG32RMOVWaddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpLOONG32RMOVWstore)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVWstore [off] {sym} ptr (MOVWconst [0]) mem)
	// result: (MOVWstorezero [off] {sym} ptr mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpLOONG32RMOVWconst || auxIntToInt32(v_1.AuxInt) != 0 {
			break
		}
		mem := v_2
		v.reset(OpLOONG32RMOVWstorezero)
		v.AuxInt = int32ToAuxInt(off)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVWstore [off] {sym} ptr (MOVWreg x) mem)
	// result: (MOVWstore [off] {sym} ptr x mem)
	for {
		off := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		ptr := v_0
		if v_1.Op != OpLOONG32RMOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpLOONG32RMOVWstore)
		v.AuxInt = int32ToAuxInt(off)
		v.Aux = symToAux(sym)
		v.AddArg3(ptr, x, mem)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RMOVWstorezero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWstorezero [off1] {sym} x:(ADDconst [off2] ptr) mem)
	// cond: (is16Bit(int64(off1+off2)) || x.Uses == 1)
	// result: (MOVWstorezero [off1+off2] {sym} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym := auxToSym(v.Aux)
		x := v_0
		if x.Op != OpLOONG32RADDconst {
			break
		}
		off2 := auxIntToInt32(x.AuxInt)
		ptr := x.Args[0]
		mem := v_1
		if !(is16Bit(int64(off1+off2)) || x.Uses == 1) {
			break
		}
		v.reset(OpLOONG32RMOVWstorezero)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVWstorezero [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVWstorezero [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := auxIntToInt32(v.AuxInt)
		sym1 := auxToSym(v.Aux)
		if v_0.Op != OpLOONG32RMOVWaddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym2 := auxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpLOONG32RMOVWstorezero)
		v.AuxInt = int32ToAuxInt(off1 + off2)
		v.Aux = symToAux(mergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RMUL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MUL (MOVWconst [0]) _ )
	// result: (MOVWconst [0])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpLOONG32RMOVWconst || auxIntToInt32(v_0.AuxInt) != 0 {
				continue
			}
			v.reset(OpLOONG32RMOVWconst)
			v.AuxInt = int32ToAuxInt(0)
			return true
		}
		break
	}
	// match: (MUL (MOVWconst [1]) x )
	// result: x
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpLOONG32RMOVWconst || auxIntToInt32(v_0.AuxInt) != 1 {
				continue
			}
			x := v_1
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (MUL (MOVWconst [-1]) x )
	// result: (NEG x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpLOONG32RMOVWconst || auxIntToInt32(v_0.AuxInt) != -1 {
				continue
			}
			x := v_1
			v.reset(OpLOONG32RNEG)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (MUL (MOVWconst [c]) x )
	// cond: isUnsignedPowerOfTwo(uint32(c))
	// result: (SLLconst [int32(log32u(uint32(c)))] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpLOONG32RMOVWconst {
				continue
			}
			c := auxIntToInt32(v_0.AuxInt)
			x := v_1
			if !(isUnsignedPowerOfTwo(uint32(c))) {
				continue
			}
			v.reset(OpLOONG32RSLLconst)
			v.AuxInt = int32ToAuxInt(int32(log32u(uint32(c))))
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (MUL (MOVWconst [c]) (MOVWconst [d]))
	// result: (MOVWconst [c*d])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpLOONG32RMOVWconst {
				continue
			}
			c := auxIntToInt32(v_0.AuxInt)
			if v_1.Op != OpLOONG32RMOVWconst {
				continue
			}
			d := auxIntToInt32(v_1.AuxInt)
			v.reset(OpLOONG32RMOVWconst)
			v.AuxInt = int32ToAuxInt(c * d)
			return true
		}
		break
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RNEG(v *Value) bool {
	v_0 := v.Args[0]
	// match: (NEG (SUB x y))
	// result: (SUB y x)
	for {
		if v_0.Op != OpLOONG32RSUB {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpLOONG32RSUB)
		v.AddArg2(y, x)
		return true
	}
	// match: (NEG (NEG x))
	// result: x
	for {
		if v_0.Op != OpLOONG32RNEG {
			break
		}
		x := v_0.Args[0]
		v.copyOf(x)
		return true
	}
	// match: (NEG (MOVWconst [c]))
	// result: (MOVWconst [-c])
	for {
		if v_0.Op != OpLOONG32RMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(-c)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RNOR(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (NOR x (MOVWconst [c]))
	// result: (NORconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpLOONG32RMOVWconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			v.reset(OpLOONG32RNORconst)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg(x)
			return true
		}
		break
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RNORconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (NORconst [c] (MOVWconst [d]))
	// result: (MOVWconst [^(c|d)])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RMOVWconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(^(c | d))
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32ROR(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (OR x (MOVWconst [c]))
	// result: (ORconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpLOONG32RMOVWconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			v.reset(OpLOONG32RORconst)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (OR x x)
	// result: x
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (OR (SGTUzero x) (SGTUzero y))
	// result: (SGTUzero (OR <x.Type> x y))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpLOONG32RSGTUzero {
				continue
			}
			x := v_0.Args[0]
			if v_1.Op != OpLOONG32RSGTUzero {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpLOONG32RSGTUzero)
			v0 := b.NewValue0(v.Pos, OpLOONG32ROR, x.Type)
			v0.AddArg2(x, y)
			v.AddArg(v0)
			return true
		}
		break
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RORconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ORconst [0] x)
	// result: x
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		x := v_0
		v.copyOf(x)
		return true
	}
	// match: (ORconst [-1] _)
	// result: (MOVWconst [-1])
	for {
		if auxIntToInt32(v.AuxInt) != -1 {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(-1)
		return true
	}
	// match: (ORconst [c] (MOVWconst [d]))
	// result: (MOVWconst [c|d])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RMOVWconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(c | d)
		return true
	}
	// match: (ORconst [c] (ORconst [d] x))
	// result: (ORconst [c|d] x)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RORconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		v.reset(OpLOONG32RORconst)
		v.AuxInt = int32ToAuxInt(c | d)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RSGT(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SGT (MOVWconst [c]) x)
	// result: (SGTconst [c] x)
	for {
		if v_0.Op != OpLOONG32RMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpLOONG32RSGTconst)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (SGT x (MOVWconst [0]))
	// result: (SGTzero x)
	for {
		x := v_0
		if v_1.Op != OpLOONG32RMOVWconst || auxIntToInt32(v_1.AuxInt) != 0 {
			break
		}
		v.reset(OpLOONG32RSGTzero)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RSGTU(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SGTU (MOVWconst [c]) x)
	// result: (SGTUconst [c] x)
	for {
		if v_0.Op != OpLOONG32RMOVWconst {
			break
		}
		c := auxIntToInt32(v_0.AuxInt)
		x := v_1
		v.reset(OpLOONG32RSGTUconst)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (SGTU x (MOVWconst [0]))
	// result: (SGTUzero x)
	for {
		x := v_0
		if v_1.Op != OpLOONG32RMOVWconst || auxIntToInt32(v_1.AuxInt) != 0 {
			break
		}
		v.reset(OpLOONG32RSGTUzero)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RSGTUconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SGTUconst [c] (MOVWconst [d]))
	// cond: uint32(c) > uint32(d)
	// result: (MOVWconst [1])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RMOVWconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		if !(uint32(c) > uint32(d)) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(1)
		return true
	}
	// match: (SGTUconst [c] (MOVWconst [d]))
	// cond: uint32(c) <= uint32(d)
	// result: (MOVWconst [0])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RMOVWconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		if !(uint32(c) <= uint32(d)) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	// match: (SGTUconst [c] (MOVBUreg _))
	// cond: 0xff < uint32(c)
	// result: (MOVWconst [1])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RMOVBUreg || !(0xff < uint32(c)) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(1)
		return true
	}
	// match: (SGTUconst [c] (MOVHUreg _))
	// cond: 0xffff < uint32(c)
	// result: (MOVWconst [1])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RMOVHUreg || !(0xffff < uint32(c)) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(1)
		return true
	}
	// match: (SGTUconst [c] (ANDconst [m] _))
	// cond: uint32(m) < uint32(c)
	// result: (MOVWconst [1])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RANDconst {
			break
		}
		m := auxIntToInt32(v_0.AuxInt)
		if !(uint32(m) < uint32(c)) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(1)
		return true
	}
	// match: (SGTUconst [c] (SRLconst _ [d]))
	// cond: uint32(d) <= 31 && 0xffffffff>>uint32(d) < uint32(c)
	// result: (MOVWconst [1])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RSRLconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		if !(uint32(d) <= 31 && 0xffffffff>>uint32(d) < uint32(c)) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(1)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RSGTUzero(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SGTUzero (MOVWconst [d]))
	// cond: d != 0
	// result: (MOVWconst [1])
	for {
		if v_0.Op != OpLOONG32RMOVWconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		if !(d != 0) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(1)
		return true
	}
	// match: (SGTUzero (MOVWconst [d]))
	// cond: d == 0
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpLOONG32RMOVWconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		if !(d == 0) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RSGTconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SGTconst [c] (MOVWconst [d]))
	// cond: c > d
	// result: (MOVWconst [1])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RMOVWconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		if !(c > d) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(1)
		return true
	}
	// match: (SGTconst [c] (MOVWconst [d]))
	// cond: c <= d
	// result: (MOVWconst [0])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RMOVWconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		if !(c <= d) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	// match: (SGTconst [c] (MOVBreg _))
	// cond: 0x7f < c
	// result: (MOVWconst [1])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RMOVBreg || !(0x7f < c) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(1)
		return true
	}
	// match: (SGTconst [c] (MOVBreg _))
	// cond: c <= -0x80
	// result: (MOVWconst [0])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RMOVBreg || !(c <= -0x80) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	// match: (SGTconst [c] (MOVBUreg _))
	// cond: 0xff < c
	// result: (MOVWconst [1])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RMOVBUreg || !(0xff < c) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(1)
		return true
	}
	// match: (SGTconst [c] (MOVBUreg _))
	// cond: c < 0
	// result: (MOVWconst [0])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RMOVBUreg || !(c < 0) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	// match: (SGTconst [c] (MOVHreg _))
	// cond: 0x7fff < c
	// result: (MOVWconst [1])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RMOVHreg || !(0x7fff < c) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(1)
		return true
	}
	// match: (SGTconst [c] (MOVHreg _))
	// cond: c <= -0x8000
	// result: (MOVWconst [0])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RMOVHreg || !(c <= -0x8000) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	// match: (SGTconst [c] (MOVHUreg _))
	// cond: 0xffff < c
	// result: (MOVWconst [1])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RMOVHUreg || !(0xffff < c) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(1)
		return true
	}
	// match: (SGTconst [c] (MOVHUreg _))
	// cond: c < 0
	// result: (MOVWconst [0])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RMOVHUreg || !(c < 0) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	// match: (SGTconst [c] (ANDconst [m] _))
	// cond: 0 <= m && m < c
	// result: (MOVWconst [1])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RANDconst {
			break
		}
		m := auxIntToInt32(v_0.AuxInt)
		if !(0 <= m && m < c) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(1)
		return true
	}
	// match: (SGTconst [c] (SRLconst _ [d]))
	// cond: 0 <= c && uint32(d) <= 31 && 0xffffffff>>uint32(d) < uint32(c)
	// result: (MOVWconst [1])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RSRLconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		if !(0 <= c && uint32(d) <= 31 && 0xffffffff>>uint32(d) < uint32(c)) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(1)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RSGTzero(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SGTzero (MOVWconst [d]))
	// cond: d > 0
	// result: (MOVWconst [1])
	for {
		if v_0.Op != OpLOONG32RMOVWconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		if !(d > 0) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(1)
		return true
	}
	// match: (SGTzero (MOVWconst [d]))
	// cond: d <= 0
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpLOONG32RMOVWconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		if !(d <= 0) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RSLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SLL x (MOVWconst [c]))
	// result: (SLLconst x [c&31])
	for {
		x := v_0
		if v_1.Op != OpLOONG32RMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpLOONG32RSLLconst)
		v.AuxInt = int32ToAuxInt(c & 31)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RSLLconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SLLconst [c] (MOVWconst [d]))
	// result: (MOVWconst [d<<uint32(c)])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RMOVWconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(d << uint32(c))
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RSRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SRA x (MOVWconst [c]))
	// result: (SRAconst x [c&31])
	for {
		x := v_0
		if v_1.Op != OpLOONG32RMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpLOONG32RSRAconst)
		v.AuxInt = int32ToAuxInt(c & 31)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RSRAconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SRAconst [c] (MOVWconst [d]))
	// result: (MOVWconst [d>>uint32(c)])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RMOVWconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(d >> uint32(c))
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RSRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SRL x (MOVWconst [c]))
	// result: (SRLconst x [c&31])
	for {
		x := v_0
		if v_1.Op != OpLOONG32RMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpLOONG32RSRLconst)
		v.AuxInt = int32ToAuxInt(c & 31)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RSRLconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SRLconst [c] (MOVWconst [d]))
	// result: (MOVWconst [int32(uint32(d)>>uint32(c))])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RMOVWconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(int32(uint32(d) >> uint32(c)))
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RSUB(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SUB x (MOVWconst [c]))
	// result: (SUBconst [c] x)
	for {
		x := v_0
		if v_1.Op != OpLOONG32RMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpLOONG32RSUBconst)
		v.AuxInt = int32ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (SUB x (NEG y))
	// result: (ADD x y)
	for {
		x := v_0
		if v_1.Op != OpLOONG32RNEG {
			break
		}
		y := v_1.Args[0]
		v.reset(OpLOONG32RADD)
		v.AddArg2(x, y)
		return true
	}
	// match: (SUB x x)
	// result: (MOVWconst [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	// match: (SUB (MOVWconst [0]) x)
	// result: (NEG x)
	for {
		if v_0.Op != OpLOONG32RMOVWconst || auxIntToInt32(v_0.AuxInt) != 0 {
			break
		}
		x := v_1
		v.reset(OpLOONG32RNEG)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RSUBconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SUBconst [0] x)
	// result: x
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		x := v_0
		v.copyOf(x)
		return true
	}
	// match: (SUBconst [c] (MOVWconst [d]))
	// result: (MOVWconst [d-c])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RMOVWconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(d - c)
		return true
	}
	// match: (SUBconst [c] (SUBconst [d] x))
	// result: (ADDconst [-c-d] x)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RSUBconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		v.reset(OpLOONG32RADDconst)
		v.AuxInt = int32ToAuxInt(-c - d)
		v.AddArg(x)
		return true
	}
	// match: (SUBconst [c] (ADDconst [d] x))
	// result: (ADDconst [-c+d] x)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RADDconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		v.reset(OpLOONG32RADDconst)
		v.AuxInt = int32ToAuxInt(-c + d)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RXOR(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (XOR x (MOVWconst [c]))
	// result: (XORconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpLOONG32RMOVWconst {
				continue
			}
			c := auxIntToInt32(v_1.AuxInt)
			v.reset(OpLOONG32RXORconst)
			v.AuxInt = int32ToAuxInt(c)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (XOR x x)
	// result: (MOVWconst [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLOONG32RXORconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (XORconst [0] x)
	// result: x
	for {
		if auxIntToInt32(v.AuxInt) != 0 {
			break
		}
		x := v_0
		v.copyOf(x)
		return true
	}
	// match: (XORconst [-1] x)
	// result: (NORconst [0] x)
	for {
		if auxIntToInt32(v.AuxInt) != -1 {
			break
		}
		x := v_0
		v.reset(OpLOONG32RNORconst)
		v.AuxInt = int32ToAuxInt(0)
		v.AddArg(x)
		return true
	}
	// match: (XORconst [c] (MOVWconst [d]))
	// result: (MOVWconst [c^d])
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RMOVWconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(c ^ d)
		return true
	}
	// match: (XORconst [c] (XORconst [d] x))
	// result: (XORconst [c^d] x)
	for {
		c := auxIntToInt32(v.AuxInt)
		if v_0.Op != OpLOONG32RXORconst {
			break
		}
		d := auxIntToInt32(v_0.AuxInt)
		x := v_0.Args[0]
		v.reset(OpLOONG32RXORconst)
		v.AuxInt = int32ToAuxInt(c ^ d)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLeq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq16 x y)
	// result: (XORconst [1] (SGT (SignExt16to32 x) (SignExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RXORconst)
		v.AuxInt = int32ToAuxInt(1)
		v0 := b.NewValue0(v.Pos, OpLOONG32RSGT, typ.Bool)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpLeq16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq16U x y)
	// result: (XORconst [1] (SGTU (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RXORconst)
		v.AuxInt = int32ToAuxInt(1)
		v0 := b.NewValue0(v.Pos, OpLOONG32RSGTU, typ.Bool)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpLeq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq32 x y)
	// result: (XORconst [1] (SGT x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RXORconst)
		v.AuxInt = int32ToAuxInt(1)
		v0 := b.NewValue0(v.Pos, OpLOONG32RSGT, typ.Bool)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpLeq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Leq32F x y)
	// result: (FPFlagTrue (CMPGEF y x))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpLOONG32RCMPGEF, types.TypeFlags)
		v0.AddArg2(y, x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpLeq32U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq32U x y)
	// result: (XORconst [1] (SGTU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RXORconst)
		v.AuxInt = int32ToAuxInt(1)
		v0 := b.NewValue0(v.Pos, OpLOONG32RSGTU, typ.Bool)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpLeq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Leq64F x y)
	// result: (FPFlagTrue (CMPGED y x))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpLOONG32RCMPGED, types.TypeFlags)
		v0.AddArg2(y, x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpLeq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq8 x y)
	// result: (XORconst [1] (SGT (SignExt8to32 x) (SignExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RXORconst)
		v.AuxInt = int32ToAuxInt(1)
		v0 := b.NewValue0(v.Pos, OpLOONG32RSGT, typ.Bool)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpLeq8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq8U x y)
	// result: (XORconst [1] (SGTU (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RXORconst)
		v.AuxInt = int32ToAuxInt(1)
		v0 := b.NewValue0(v.Pos, OpLOONG32RSGTU, typ.Bool)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpLess16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less16 x y)
	// result: (SGT (SignExt16to32 y) (SignExt16to32 x))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RSGT)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(x)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueLOONG32R_OpLess16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less16U x y)
	// result: (SGTU (ZeroExt16to32 y) (ZeroExt16to32 x))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RSGTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueLOONG32R_OpLess32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Less32 x y)
	// result: (SGT y x)
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RSGT)
		v.AddArg2(y, x)
		return true
	}
}
func rewriteValueLOONG32R_OpLess32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Less32F x y)
	// result: (FPFlagTrue (CMPGTF y x))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpLOONG32RCMPGTF, types.TypeFlags)
		v0.AddArg2(y, x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpLess32U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Less32U x y)
	// result: (SGTU y x)
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RSGTU)
		v.AddArg2(y, x)
		return true
	}
}
func rewriteValueLOONG32R_OpLess64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Less64F x y)
	// result: (FPFlagTrue (CMPGTD y x))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpLOONG32RCMPGTD, types.TypeFlags)
		v0.AddArg2(y, x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpLess8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less8 x y)
	// result: (SGT (SignExt8to32 y) (SignExt8to32 x))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RSGT)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(x)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueLOONG32R_OpLess8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less8U x y)
	// result: (SGTU (ZeroExt8to32 y) (ZeroExt8to32 x))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RSGTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueLOONG32R_OpLoad(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Load <t> ptr mem)
	// cond: t.IsBoolean()
	// result: (MOVBUload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(t.IsBoolean()) {
			break
		}
		v.reset(OpLOONG32RMOVBUload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is8BitInt(t) && t.IsSigned())
	// result: (MOVBload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is8BitInt(t) && t.IsSigned()) {
			break
		}
		v.reset(OpLOONG32RMOVBload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is8BitInt(t) && !t.IsSigned())
	// result: (MOVBUload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is8BitInt(t) && !t.IsSigned()) {
			break
		}
		v.reset(OpLOONG32RMOVBUload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is16BitInt(t) && t.IsSigned())
	// result: (MOVHload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is16BitInt(t) && t.IsSigned()) {
			break
		}
		v.reset(OpLOONG32RMOVHload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is16BitInt(t) && !t.IsSigned())
	// result: (MOVHUload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is16BitInt(t) && !t.IsSigned()) {
			break
		}
		v.reset(OpLOONG32RMOVHUload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is32BitInt(t) || isPtr(t))
	// result: (MOVWload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is32BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpLOONG32RMOVWload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is32BitFloat(t)
	// result: (MOVFload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is32BitFloat(t)) {
			break
		}
		v.reset(OpLOONG32RMOVFload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is64BitFloat(t)
	// result: (MOVDload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is64BitFloat(t)) {
			break
		}
		v.reset(OpLOONG32RMOVDload)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLocalAddr(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (LocalAddr <t> {sym} base mem)
	// cond: t.Elem().HasPointers()
	// result: (MOVWaddr {sym} (SPanchored base mem))
	for {
		t := v.Type
		sym := auxToSym(v.Aux)
		base := v_0
		mem := v_1
		if !(t.Elem().HasPointers()) {
			break
		}
		v.reset(OpLOONG32RMOVWaddr)
		v.Aux = symToAux(sym)
		v0 := b.NewValue0(v.Pos, OpSPanchored, typ.Uintptr)
		v0.AddArg2(base, mem)
		v.AddArg(v0)
		return true
	}
	// match: (LocalAddr <t> {sym} base _)
	// cond: !t.Elem().HasPointers()
	// result: (MOVWaddr {sym} base)
	for {
		t := v.Type
		sym := auxToSym(v.Aux)
		base := v_0
		if !(!t.Elem().HasPointers()) {
			break
		}
		v.reset(OpLOONG32RMOVWaddr)
		v.Aux = symToAux(sym)
		v.AddArg(base)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLsh16x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x16 <t> x y)
	// result: (CMOVZ (SLL <t> x (ZeroExt16to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt16to32 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpLOONG32RCMOVZ)
		v0 := b.NewValue0(v.Pos, OpLOONG32RSLL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v2.AuxInt = int32ToAuxInt(0)
		v3 := b.NewValue0(v.Pos, OpLOONG32RSGTUconst, typ.Bool)
		v3.AuxInt = int32ToAuxInt(32)
		v3.AddArg(v1)
		v.AddArg3(v0, v2, v3)
		return true
	}
}
func rewriteValueLOONG32R_OpLsh16x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x32 <t> x y)
	// result: (CMOVZ (SLL <t> x y) (MOVWconst [0]) (SGTUconst [32] y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpLOONG32RCMOVZ)
		v0 := b.NewValue0(v.Pos, OpLOONG32RSLL, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v1.AuxInt = int32ToAuxInt(0)
		v2 := b.NewValue0(v.Pos, OpLOONG32RSGTUconst, typ.Bool)
		v2.AuxInt = int32ToAuxInt(32)
		v2.AddArg(y)
		v.AddArg3(v0, v1, v2)
		return true
	}
}
func rewriteValueLOONG32R_OpLsh16x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Lsh16x64 x (Const64 [c]))
	// cond: uint32(c) < 16
	// result: (SLLconst x [int32(c)])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint32(c) < 16) {
			break
		}
		v.reset(OpLOONG32RSLLconst)
		v.AuxInt = int32ToAuxInt(int32(c))
		v.AddArg(x)
		return true
	}
	// match: (Lsh16x64 _ (Const64 [c]))
	// cond: uint32(c) >= 16
	// result: (MOVWconst [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint32(c) >= 16) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLsh16x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x8 <t> x y)
	// result: (CMOVZ (SLL <t> x (ZeroExt8to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt8to32 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpLOONG32RCMOVZ)
		v0 := b.NewValue0(v.Pos, OpLOONG32RSLL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v2.AuxInt = int32ToAuxInt(0)
		v3 := b.NewValue0(v.Pos, OpLOONG32RSGTUconst, typ.Bool)
		v3.AuxInt = int32ToAuxInt(32)
		v3.AddArg(v1)
		v.AddArg3(v0, v2, v3)
		return true
	}
}
func rewriteValueLOONG32R_OpLsh32x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x16 <t> x y)
	// result: (CMOVZ (SLL <t> x (ZeroExt16to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt16to32 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpLOONG32RCMOVZ)
		v0 := b.NewValue0(v.Pos, OpLOONG32RSLL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v2.AuxInt = int32ToAuxInt(0)
		v3 := b.NewValue0(v.Pos, OpLOONG32RSGTUconst, typ.Bool)
		v3.AuxInt = int32ToAuxInt(32)
		v3.AddArg(v1)
		v.AddArg3(v0, v2, v3)
		return true
	}
}
func rewriteValueLOONG32R_OpLsh32x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x32 <t> x y)
	// result: (CMOVZ (SLL <t> x y) (MOVWconst [0]) (SGTUconst [32] y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpLOONG32RCMOVZ)
		v0 := b.NewValue0(v.Pos, OpLOONG32RSLL, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v1.AuxInt = int32ToAuxInt(0)
		v2 := b.NewValue0(v.Pos, OpLOONG32RSGTUconst, typ.Bool)
		v2.AuxInt = int32ToAuxInt(32)
		v2.AddArg(y)
		v.AddArg3(v0, v1, v2)
		return true
	}
}
func rewriteValueLOONG32R_OpLsh32x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Lsh32x64 x (Const64 [c]))
	// cond: uint32(c) < 32
	// result: (SLLconst x [int32(c)])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint32(c) < 32) {
			break
		}
		v.reset(OpLOONG32RSLLconst)
		v.AuxInt = int32ToAuxInt(int32(c))
		v.AddArg(x)
		return true
	}
	// match: (Lsh32x64 _ (Const64 [c]))
	// cond: uint32(c) >= 32
	// result: (MOVWconst [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint32(c) >= 32) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLsh32x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x8 <t> x y)
	// result: (CMOVZ (SLL <t> x (ZeroExt8to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt8to32 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpLOONG32RCMOVZ)
		v0 := b.NewValue0(v.Pos, OpLOONG32RSLL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v2.AuxInt = int32ToAuxInt(0)
		v3 := b.NewValue0(v.Pos, OpLOONG32RSGTUconst, typ.Bool)
		v3.AuxInt = int32ToAuxInt(32)
		v3.AddArg(v1)
		v.AddArg3(v0, v2, v3)
		return true
	}
}
func rewriteValueLOONG32R_OpLsh8x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x16 <t> x y)
	// result: (CMOVZ (SLL <t> x (ZeroExt16to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt16to32 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpLOONG32RCMOVZ)
		v0 := b.NewValue0(v.Pos, OpLOONG32RSLL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v2.AuxInt = int32ToAuxInt(0)
		v3 := b.NewValue0(v.Pos, OpLOONG32RSGTUconst, typ.Bool)
		v3.AuxInt = int32ToAuxInt(32)
		v3.AddArg(v1)
		v.AddArg3(v0, v2, v3)
		return true
	}
}
func rewriteValueLOONG32R_OpLsh8x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x32 <t> x y)
	// result: (CMOVZ (SLL <t> x y) (MOVWconst [0]) (SGTUconst [32] y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpLOONG32RCMOVZ)
		v0 := b.NewValue0(v.Pos, OpLOONG32RSLL, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v1.AuxInt = int32ToAuxInt(0)
		v2 := b.NewValue0(v.Pos, OpLOONG32RSGTUconst, typ.Bool)
		v2.AuxInt = int32ToAuxInt(32)
		v2.AddArg(y)
		v.AddArg3(v0, v1, v2)
		return true
	}
}
func rewriteValueLOONG32R_OpLsh8x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Lsh8x64 x (Const64 [c]))
	// cond: uint32(c) < 8
	// result: (SLLconst x [int32(c)])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint32(c) < 8) {
			break
		}
		v.reset(OpLOONG32RSLLconst)
		v.AuxInt = int32ToAuxInt(int32(c))
		v.AddArg(x)
		return true
	}
	// match: (Lsh8x64 _ (Const64 [c]))
	// cond: uint32(c) >= 8
	// result: (MOVWconst [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint32(c) >= 8) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpLsh8x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x8 <t> x y)
	// result: (CMOVZ (SLL <t> x (ZeroExt8to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt8to32 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpLOONG32RCMOVZ)
		v0 := b.NewValue0(v.Pos, OpLOONG32RSLL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v2.AuxInt = int32ToAuxInt(0)
		v3 := b.NewValue0(v.Pos, OpLOONG32RSGTUconst, typ.Bool)
		v3.AuxInt = int32ToAuxInt(32)
		v3.AddArg(v1)
		v.AddArg3(v0, v2, v3)
		return true
	}
}
func rewriteValueLOONG32R_OpMod16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod16 x y)
	// result: (Select0 (DIV (SignExt16to32 x) (SignExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpLOONG32RDIV, types.NewTuple(typ.Int32, typ.Int32))
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpMod16u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod16u x y)
	// result: (Select0 (DIVU (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpLOONG32RDIVU, types.NewTuple(typ.UInt32, typ.UInt32))
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpMod32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod32 x y)
	// result: (Select0 (DIV x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpLOONG32RDIV, types.NewTuple(typ.Int32, typ.Int32))
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpMod32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod32u x y)
	// result: (Select0 (DIVU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpLOONG32RDIVU, types.NewTuple(typ.UInt32, typ.UInt32))
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpMod8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod8 x y)
	// result: (Select0 (DIV (SignExt8to32 x) (SignExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpLOONG32RDIV, types.NewTuple(typ.Int32, typ.Int32))
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpMod8u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod8u x y)
	// result: (Select0 (DIVU (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpLOONG32RDIVU, types.NewTuple(typ.UInt32, typ.UInt32))
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpMove(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (Move [0] _ _ mem)
	// result: mem
	for {
		if auxIntToInt64(v.AuxInt) != 0 {
			break
		}
		mem := v_2
		v.copyOf(mem)
		return true
	}
	// match: (Move [1] dst src mem)
	// result: (MOVBstore dst (MOVBUload src mem) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 1 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpLOONG32RMOVBstore)
		v0 := b.NewValue0(v.Pos, OpLOONG32RMOVBUload, typ.UInt8)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [2] {t} dst src mem)
	// cond: t.Alignment()%2 == 0
	// result: (MOVHstore dst (MOVHUload src mem) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 2 {
			break
		}
		t := auxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.Alignment()%2 == 0) {
			break
		}
		v.reset(OpLOONG32RMOVHstore)
		v0 := b.NewValue0(v.Pos, OpLOONG32RMOVHUload, typ.UInt16)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [2] dst src mem)
	// result: (MOVBstore [1] dst (MOVBUload [1] src mem) (MOVBstore dst (MOVBUload src mem) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 2 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpLOONG32RMOVBstore)
		v.AuxInt = int32ToAuxInt(1)
		v0 := b.NewValue0(v.Pos, OpLOONG32RMOVBUload, typ.UInt8)
		v0.AuxInt = int32ToAuxInt(1)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVBstore, types.TypeMem)
		v2 := b.NewValue0(v.Pos, OpLOONG32RMOVBUload, typ.UInt8)
		v2.AddArg2(src, mem)
		v1.AddArg3(dst, v2, mem)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [4] {t} dst src mem)
	// cond: t.Alignment()%4 == 0
	// result: (MOVWstore dst (MOVWload src mem) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 4 {
			break
		}
		t := auxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.Alignment()%4 == 0) {
			break
		}
		v.reset(OpLOONG32RMOVWstore)
		v0 := b.NewValue0(v.Pos, OpLOONG32RMOVWload, typ.UInt32)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [4] {t} dst src mem)
	// cond: t.Alignment()%2 == 0
	// result: (MOVHstore [2] dst (MOVHUload [2] src mem) (MOVHstore dst (MOVHUload src mem) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 4 {
			break
		}
		t := auxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.Alignment()%2 == 0) {
			break
		}
		v.reset(OpLOONG32RMOVHstore)
		v.AuxInt = int32ToAuxInt(2)
		v0 := b.NewValue0(v.Pos, OpLOONG32RMOVHUload, typ.UInt16)
		v0.AuxInt = int32ToAuxInt(2)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVHstore, types.TypeMem)
		v2 := b.NewValue0(v.Pos, OpLOONG32RMOVHUload, typ.UInt16)
		v2.AddArg2(src, mem)
		v1.AddArg3(dst, v2, mem)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [4] dst src mem)
	// result: (MOVBstore [3] dst (MOVBUload [3] src mem) (MOVBstore [2] dst (MOVBUload [2] src mem) (MOVBstore [1] dst (MOVBUload [1] src mem) (MOVBstore dst (MOVBUload src mem) mem))))
	for {
		if auxIntToInt64(v.AuxInt) != 4 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpLOONG32RMOVBstore)
		v.AuxInt = int32ToAuxInt(3)
		v0 := b.NewValue0(v.Pos, OpLOONG32RMOVBUload, typ.UInt8)
		v0.AuxInt = int32ToAuxInt(3)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVBstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(2)
		v2 := b.NewValue0(v.Pos, OpLOONG32RMOVBUload, typ.UInt8)
		v2.AuxInt = int32ToAuxInt(2)
		v2.AddArg2(src, mem)
		v3 := b.NewValue0(v.Pos, OpLOONG32RMOVBstore, types.TypeMem)
		v3.AuxInt = int32ToAuxInt(1)
		v4 := b.NewValue0(v.Pos, OpLOONG32RMOVBUload, typ.UInt8)
		v4.AuxInt = int32ToAuxInt(1)
		v4.AddArg2(src, mem)
		v5 := b.NewValue0(v.Pos, OpLOONG32RMOVBstore, types.TypeMem)
		v6 := b.NewValue0(v.Pos, OpLOONG32RMOVBUload, typ.UInt8)
		v6.AddArg2(src, mem)
		v5.AddArg3(dst, v6, mem)
		v3.AddArg3(dst, v4, v5)
		v1.AddArg3(dst, v2, v3)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [3] dst src mem)
	// result: (MOVBstore [2] dst (MOVBUload [2] src mem) (MOVBstore [1] dst (MOVBUload [1] src mem) (MOVBstore dst (MOVBUload src mem) mem)))
	for {
		if auxIntToInt64(v.AuxInt) != 3 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpLOONG32RMOVBstore)
		v.AuxInt = int32ToAuxInt(2)
		v0 := b.NewValue0(v.Pos, OpLOONG32RMOVBUload, typ.UInt8)
		v0.AuxInt = int32ToAuxInt(2)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVBstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(1)
		v2 := b.NewValue0(v.Pos, OpLOONG32RMOVBUload, typ.UInt8)
		v2.AuxInt = int32ToAuxInt(1)
		v2.AddArg2(src, mem)
		v3 := b.NewValue0(v.Pos, OpLOONG32RMOVBstore, types.TypeMem)
		v4 := b.NewValue0(v.Pos, OpLOONG32RMOVBUload, typ.UInt8)
		v4.AddArg2(src, mem)
		v3.AddArg3(dst, v4, mem)
		v1.AddArg3(dst, v2, v3)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [8] {t} dst src mem)
	// cond: t.Alignment()%4 == 0
	// result: (MOVWstore [4] dst (MOVWload [4] src mem) (MOVWstore dst (MOVWload src mem) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 8 {
			break
		}
		t := auxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.Alignment()%4 == 0) {
			break
		}
		v.reset(OpLOONG32RMOVWstore)
		v.AuxInt = int32ToAuxInt(4)
		v0 := b.NewValue0(v.Pos, OpLOONG32RMOVWload, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(4)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVWstore, types.TypeMem)
		v2 := b.NewValue0(v.Pos, OpLOONG32RMOVWload, typ.UInt32)
		v2.AddArg2(src, mem)
		v1.AddArg3(dst, v2, mem)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [8] {t} dst src mem)
	// cond: t.Alignment()%2 == 0
	// result: (MOVHstore [6] dst (MOVHload [6] src mem) (MOVHstore [4] dst (MOVHload [4] src mem) (MOVHstore [2] dst (MOVHload [2] src mem) (MOVHstore dst (MOVHload src mem) mem))))
	for {
		if auxIntToInt64(v.AuxInt) != 8 {
			break
		}
		t := auxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.Alignment()%2 == 0) {
			break
		}
		v.reset(OpLOONG32RMOVHstore)
		v.AuxInt = int32ToAuxInt(6)
		v0 := b.NewValue0(v.Pos, OpLOONG32RMOVHload, typ.Int16)
		v0.AuxInt = int32ToAuxInt(6)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVHstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(4)
		v2 := b.NewValue0(v.Pos, OpLOONG32RMOVHload, typ.Int16)
		v2.AuxInt = int32ToAuxInt(4)
		v2.AddArg2(src, mem)
		v3 := b.NewValue0(v.Pos, OpLOONG32RMOVHstore, types.TypeMem)
		v3.AuxInt = int32ToAuxInt(2)
		v4 := b.NewValue0(v.Pos, OpLOONG32RMOVHload, typ.Int16)
		v4.AuxInt = int32ToAuxInt(2)
		v4.AddArg2(src, mem)
		v5 := b.NewValue0(v.Pos, OpLOONG32RMOVHstore, types.TypeMem)
		v6 := b.NewValue0(v.Pos, OpLOONG32RMOVHload, typ.Int16)
		v6.AddArg2(src, mem)
		v5.AddArg3(dst, v6, mem)
		v3.AddArg3(dst, v4, v5)
		v1.AddArg3(dst, v2, v3)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [6] {t} dst src mem)
	// cond: t.Alignment()%2 == 0
	// result: (MOVHstore [4] dst (MOVHload [4] src mem) (MOVHstore [2] dst (MOVHload [2] src mem) (MOVHstore dst (MOVHload src mem) mem)))
	for {
		if auxIntToInt64(v.AuxInt) != 6 {
			break
		}
		t := auxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.Alignment()%2 == 0) {
			break
		}
		v.reset(OpLOONG32RMOVHstore)
		v.AuxInt = int32ToAuxInt(4)
		v0 := b.NewValue0(v.Pos, OpLOONG32RMOVHload, typ.Int16)
		v0.AuxInt = int32ToAuxInt(4)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVHstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(2)
		v2 := b.NewValue0(v.Pos, OpLOONG32RMOVHload, typ.Int16)
		v2.AuxInt = int32ToAuxInt(2)
		v2.AddArg2(src, mem)
		v3 := b.NewValue0(v.Pos, OpLOONG32RMOVHstore, types.TypeMem)
		v4 := b.NewValue0(v.Pos, OpLOONG32RMOVHload, typ.Int16)
		v4.AddArg2(src, mem)
		v3.AddArg3(dst, v4, mem)
		v1.AddArg3(dst, v2, v3)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [12] {t} dst src mem)
	// cond: t.Alignment()%4 == 0
	// result: (MOVWstore [8] dst (MOVWload [8] src mem) (MOVWstore [4] dst (MOVWload [4] src mem) (MOVWstore dst (MOVWload src mem) mem)))
	for {
		if auxIntToInt64(v.AuxInt) != 12 {
			break
		}
		t := auxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.Alignment()%4 == 0) {
			break
		}
		v.reset(OpLOONG32RMOVWstore)
		v.AuxInt = int32ToAuxInt(8)
		v0 := b.NewValue0(v.Pos, OpLOONG32RMOVWload, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(8)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVWstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(4)
		v2 := b.NewValue0(v.Pos, OpLOONG32RMOVWload, typ.UInt32)
		v2.AuxInt = int32ToAuxInt(4)
		v2.AddArg2(src, mem)
		v3 := b.NewValue0(v.Pos, OpLOONG32RMOVWstore, types.TypeMem)
		v4 := b.NewValue0(v.Pos, OpLOONG32RMOVWload, typ.UInt32)
		v4.AddArg2(src, mem)
		v3.AddArg3(dst, v4, mem)
		v1.AddArg3(dst, v2, v3)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [16] {t} dst src mem)
	// cond: t.Alignment()%4 == 0
	// result: (MOVWstore [12] dst (MOVWload [12] src mem) (MOVWstore [8] dst (MOVWload [8] src mem) (MOVWstore [4] dst (MOVWload [4] src mem) (MOVWstore dst (MOVWload src mem) mem))))
	for {
		if auxIntToInt64(v.AuxInt) != 16 {
			break
		}
		t := auxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.Alignment()%4 == 0) {
			break
		}
		v.reset(OpLOONG32RMOVWstore)
		v.AuxInt = int32ToAuxInt(12)
		v0 := b.NewValue0(v.Pos, OpLOONG32RMOVWload, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(12)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVWstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(8)
		v2 := b.NewValue0(v.Pos, OpLOONG32RMOVWload, typ.UInt32)
		v2.AuxInt = int32ToAuxInt(8)
		v2.AddArg2(src, mem)
		v3 := b.NewValue0(v.Pos, OpLOONG32RMOVWstore, types.TypeMem)
		v3.AuxInt = int32ToAuxInt(4)
		v4 := b.NewValue0(v.Pos, OpLOONG32RMOVWload, typ.UInt32)
		v4.AuxInt = int32ToAuxInt(4)
		v4.AddArg2(src, mem)
		v5 := b.NewValue0(v.Pos, OpLOONG32RMOVWstore, types.TypeMem)
		v6 := b.NewValue0(v.Pos, OpLOONG32RMOVWload, typ.UInt32)
		v6.AddArg2(src, mem)
		v5.AddArg3(dst, v6, mem)
		v3.AddArg3(dst, v4, v5)
		v1.AddArg3(dst, v2, v3)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [s] {t} dst src mem)
	// cond: (s > 16 && logLargeCopy(v, s) || t.Alignment()%4 != 0)
	// result: (LoweredMove [int32(t.Alignment())] dst src (ADDconst <src.Type> src [int32(s-moveSize(t.Alignment(), config))]) mem)
	for {
		s := auxIntToInt64(v.AuxInt)
		t := auxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(s > 16 && logLargeCopy(v, s) || t.Alignment()%4 != 0) {
			break
		}
		v.reset(OpLOONG32RLoweredMove)
		v.AuxInt = int32ToAuxInt(int32(t.Alignment()))
		v0 := b.NewValue0(v.Pos, OpLOONG32RADDconst, src.Type)
		v0.AuxInt = int32ToAuxInt(int32(s - moveSize(t.Alignment(), config)))
		v0.AddArg(src)
		v.AddArg4(dst, src, v0, mem)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpNeq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq16 x y)
	// result: (SGTU (XOR (ZeroExt16to32 x) (ZeroExt16to32 y)) (MOVWconst [0]))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RSGTU)
		v0 := b.NewValue0(v.Pos, OpLOONG32RXOR, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v3 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v3.AuxInt = int32ToAuxInt(0)
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValueLOONG32R_OpNeq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq32 x y)
	// result: (SGTU (XOR x y) (MOVWconst [0]))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RSGTU)
		v0 := b.NewValue0(v.Pos, OpLOONG32RXOR, typ.UInt32)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v1.AuxInt = int32ToAuxInt(0)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueLOONG32R_OpNeq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Neq32F x y)
	// result: (FPFlagFalse (CMPEQF x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RFPFlagFalse)
		v0 := b.NewValue0(v.Pos, OpLOONG32RCMPEQF, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpNeq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Neq64F x y)
	// result: (FPFlagFalse (CMPEQD x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RFPFlagFalse)
		v0 := b.NewValue0(v.Pos, OpLOONG32RCMPEQD, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpNeq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq8 x y)
	// result: (SGTU (XOR (ZeroExt8to32 x) (ZeroExt8to32 y)) (MOVWconst [0]))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RSGTU)
		v0 := b.NewValue0(v.Pos, OpLOONG32RXOR, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v3 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v3.AuxInt = int32ToAuxInt(0)
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValueLOONG32R_OpNeqPtr(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (NeqPtr x y)
	// result: (SGTU (XOR x y) (MOVWconst [0]))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RSGTU)
		v0 := b.NewValue0(v.Pos, OpLOONG32RXOR, typ.UInt32)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v1.AuxInt = int32ToAuxInt(0)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueLOONG32R_OpNot(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Not x)
	// result: (XORconst [1] x)
	for {
		x := v_0
		v.reset(OpLOONG32RXORconst)
		v.AuxInt = int32ToAuxInt(1)
		v.AddArg(x)
		return true
	}
}
func rewriteValueLOONG32R_OpOffPtr(v *Value) bool {
	v_0 := v.Args[0]
	// match: (OffPtr [off] ptr:(SP))
	// result: (MOVWaddr [int32(off)] ptr)
	for {
		off := auxIntToInt64(v.AuxInt)
		ptr := v_0
		if ptr.Op != OpSP {
			break
		}
		v.reset(OpLOONG32RMOVWaddr)
		v.AuxInt = int32ToAuxInt(int32(off))
		v.AddArg(ptr)
		return true
	}
	// match: (OffPtr [off] ptr)
	// result: (ADDconst [int32(off)] ptr)
	for {
		off := auxIntToInt64(v.AuxInt)
		ptr := v_0
		v.reset(OpLOONG32RADDconst)
		v.AuxInt = int32ToAuxInt(int32(off))
		v.AddArg(ptr)
		return true
	}
}
func rewriteValueLOONG32R_OpRotateLeft16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft16 <t> x (MOVWconst [c]))
	// result: (Or16 (Lsh16x32 <t> x (MOVWconst [c&15])) (Rsh16Ux32 <t> x (MOVWconst [-c&15])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpLOONG32RMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpOr16)
		v0 := b.NewValue0(v.Pos, OpLsh16x32, t)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v1.AuxInt = int32ToAuxInt(c & 15)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpRsh16Ux32, t)
		v3 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v3.AuxInt = int32ToAuxInt(-c & 15)
		v2.AddArg2(x, v3)
		v.AddArg2(v0, v2)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpRotateLeft32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft32 <t> x (MOVWconst [c]))
	// result: (Or32 (Lsh32x32 <t> x (MOVWconst [c&31])) (Rsh32Ux32 <t> x (MOVWconst [-c&31])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpLOONG32RMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpOr32)
		v0 := b.NewValue0(v.Pos, OpLsh32x32, t)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v1.AuxInt = int32ToAuxInt(c & 31)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpRsh32Ux32, t)
		v3 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v3.AuxInt = int32ToAuxInt(-c & 31)
		v2.AddArg2(x, v3)
		v.AddArg2(v0, v2)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpRotateLeft64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft64 <t> x (MOVWconst [c]))
	// result: (Or64 (Lsh64x32 <t> x (MOVWconst [c&63])) (Rsh64Ux32 <t> x (MOVWconst [-c&63])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpLOONG32RMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpOr64)
		v0 := b.NewValue0(v.Pos, OpLsh64x32, t)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v1.AuxInt = int32ToAuxInt(c & 63)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpRsh64Ux32, t)
		v3 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v3.AuxInt = int32ToAuxInt(-c & 63)
		v2.AddArg2(x, v3)
		v.AddArg2(v0, v2)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpRotateLeft8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft8 <t> x (MOVWconst [c]))
	// result: (Or8 (Lsh8x32 <t> x (MOVWconst [c&7])) (Rsh8Ux32 <t> x (MOVWconst [-c&7])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpLOONG32RMOVWconst {
			break
		}
		c := auxIntToInt32(v_1.AuxInt)
		v.reset(OpOr8)
		v0 := b.NewValue0(v.Pos, OpLsh8x32, t)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v1.AuxInt = int32ToAuxInt(c & 7)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpRsh8Ux32, t)
		v3 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v3.AuxInt = int32ToAuxInt(-c & 7)
		v2.AddArg2(x, v3)
		v.AddArg2(v0, v2)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpRsh16Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux16 <t> x y)
	// result: (CMOVZ (SRL <t> (ZeroExt16to32 x) (ZeroExt16to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt16to32 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpLOONG32RCMOVZ)
		v0 := b.NewValue0(v.Pos, OpLOONG32RSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v3 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v3.AuxInt = int32ToAuxInt(0)
		v4 := b.NewValue0(v.Pos, OpLOONG32RSGTUconst, typ.Bool)
		v4.AuxInt = int32ToAuxInt(32)
		v4.AddArg(v2)
		v.AddArg3(v0, v3, v4)
		return true
	}
}
func rewriteValueLOONG32R_OpRsh16Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux32 <t> x y)
	// result: (CMOVZ (SRL <t> (ZeroExt16to32 x) y) (MOVWconst [0]) (SGTUconst [32] y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpLOONG32RCMOVZ)
		v0 := b.NewValue0(v.Pos, OpLOONG32RSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg2(v1, y)
		v2 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v2.AuxInt = int32ToAuxInt(0)
		v3 := b.NewValue0(v.Pos, OpLOONG32RSGTUconst, typ.Bool)
		v3.AuxInt = int32ToAuxInt(32)
		v3.AddArg(y)
		v.AddArg3(v0, v2, v3)
		return true
	}
}
func rewriteValueLOONG32R_OpRsh16Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux64 x (Const64 [c]))
	// cond: uint32(c) < 16
	// result: (SRLconst (SLLconst <typ.UInt32> x [16]) [int32(c+16)])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint32(c) < 16) {
			break
		}
		v.reset(OpLOONG32RSRLconst)
		v.AuxInt = int32ToAuxInt(int32(c + 16))
		v0 := b.NewValue0(v.Pos, OpLOONG32RSLLconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(16)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16Ux64 _ (Const64 [c]))
	// cond: uint32(c) >= 16
	// result: (MOVWconst [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint32(c) >= 16) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpRsh16Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux8 <t> x y)
	// result: (CMOVZ (SRL <t> (ZeroExt16to32 x) (ZeroExt8to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt8to32 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpLOONG32RCMOVZ)
		v0 := b.NewValue0(v.Pos, OpLOONG32RSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v3 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v3.AuxInt = int32ToAuxInt(0)
		v4 := b.NewValue0(v.Pos, OpLOONG32RSGTUconst, typ.Bool)
		v4.AuxInt = int32ToAuxInt(32)
		v4.AddArg(v2)
		v.AddArg3(v0, v3, v4)
		return true
	}
}
func rewriteValueLOONG32R_OpRsh16x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x16 x y)
	// result: (SRA (SignExt16to32 x) ( CMOVZ <typ.UInt32> (ZeroExt16to32 y) (MOVWconst [31]) (SGTUconst [32] (ZeroExt16to32 y))))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpLOONG32RCMOVZ, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v3.AuxInt = int32ToAuxInt(31)
		v4 := b.NewValue0(v.Pos, OpLOONG32RSGTUconst, typ.Bool)
		v4.AuxInt = int32ToAuxInt(32)
		v4.AddArg(v2)
		v1.AddArg3(v2, v3, v4)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueLOONG32R_OpRsh16x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x32 x y)
	// result: (SRA (SignExt16to32 x) ( CMOVZ <typ.UInt32> y (MOVWconst [31]) (SGTUconst [32] y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpLOONG32RCMOVZ, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v2.AuxInt = int32ToAuxInt(31)
		v3 := b.NewValue0(v.Pos, OpLOONG32RSGTUconst, typ.Bool)
		v3.AuxInt = int32ToAuxInt(32)
		v3.AddArg(y)
		v1.AddArg3(y, v2, v3)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueLOONG32R_OpRsh16x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x64 x (Const64 [c]))
	// cond: uint32(c) < 16
	// result: (SRAconst (SLLconst <typ.UInt32> x [16]) [int32(c+16)])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint32(c) < 16) {
			break
		}
		v.reset(OpLOONG32RSRAconst)
		v.AuxInt = int32ToAuxInt(int32(c + 16))
		v0 := b.NewValue0(v.Pos, OpLOONG32RSLLconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(16)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16x64 x (Const64 [c]))
	// cond: uint32(c) >= 16
	// result: (SRAconst (SLLconst <typ.UInt32> x [16]) [31])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint32(c) >= 16) {
			break
		}
		v.reset(OpLOONG32RSRAconst)
		v.AuxInt = int32ToAuxInt(31)
		v0 := b.NewValue0(v.Pos, OpLOONG32RSLLconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(16)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpRsh16x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x8 x y)
	// result: (SRA (SignExt16to32 x) ( CMOVZ <typ.UInt32> (ZeroExt8to32 y) (MOVWconst [31]) (SGTUconst [32] (ZeroExt8to32 y))))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpLOONG32RCMOVZ, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v3.AuxInt = int32ToAuxInt(31)
		v4 := b.NewValue0(v.Pos, OpLOONG32RSGTUconst, typ.Bool)
		v4.AuxInt = int32ToAuxInt(32)
		v4.AddArg(v2)
		v1.AddArg3(v2, v3, v4)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueLOONG32R_OpRsh32Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux16 <t> x y)
	// result: (CMOVZ (SRL <t> x (ZeroExt16to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt16to32 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpLOONG32RCMOVZ)
		v0 := b.NewValue0(v.Pos, OpLOONG32RSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v2.AuxInt = int32ToAuxInt(0)
		v3 := b.NewValue0(v.Pos, OpLOONG32RSGTUconst, typ.Bool)
		v3.AuxInt = int32ToAuxInt(32)
		v3.AddArg(v1)
		v.AddArg3(v0, v2, v3)
		return true
	}
}
func rewriteValueLOONG32R_OpRsh32Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux32 <t> x y)
	// result: (CMOVZ (SRL <t> x y) (MOVWconst [0]) (SGTUconst [32] y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpLOONG32RCMOVZ)
		v0 := b.NewValue0(v.Pos, OpLOONG32RSRL, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v1.AuxInt = int32ToAuxInt(0)
		v2 := b.NewValue0(v.Pos, OpLOONG32RSGTUconst, typ.Bool)
		v2.AuxInt = int32ToAuxInt(32)
		v2.AddArg(y)
		v.AddArg3(v0, v1, v2)
		return true
	}
}
func rewriteValueLOONG32R_OpRsh32Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Rsh32Ux64 x (Const64 [c]))
	// cond: uint32(c) < 32
	// result: (SRLconst x [int32(c)])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint32(c) < 32) {
			break
		}
		v.reset(OpLOONG32RSRLconst)
		v.AuxInt = int32ToAuxInt(int32(c))
		v.AddArg(x)
		return true
	}
	// match: (Rsh32Ux64 _ (Const64 [c]))
	// cond: uint32(c) >= 32
	// result: (MOVWconst [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint32(c) >= 32) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpRsh32Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux8 <t> x y)
	// result: (CMOVZ (SRL <t> x (ZeroExt8to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt8to32 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpLOONG32RCMOVZ)
		v0 := b.NewValue0(v.Pos, OpLOONG32RSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v2.AuxInt = int32ToAuxInt(0)
		v3 := b.NewValue0(v.Pos, OpLOONG32RSGTUconst, typ.Bool)
		v3.AuxInt = int32ToAuxInt(32)
		v3.AddArg(v1)
		v.AddArg3(v0, v2, v3)
		return true
	}
}
func rewriteValueLOONG32R_OpRsh32x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x16 x y)
	// result: (SRA x ( CMOVZ <typ.UInt32> (ZeroExt16to32 y) (MOVWconst [31]) (SGTUconst [32] (ZeroExt16to32 y))))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RSRA)
		v0 := b.NewValue0(v.Pos, OpLOONG32RCMOVZ, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v2.AuxInt = int32ToAuxInt(31)
		v3 := b.NewValue0(v.Pos, OpLOONG32RSGTUconst, typ.Bool)
		v3.AuxInt = int32ToAuxInt(32)
		v3.AddArg(v1)
		v0.AddArg3(v1, v2, v3)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueLOONG32R_OpRsh32x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x32 x y)
	// result: (SRA x ( CMOVZ <typ.UInt32> y (MOVWconst [31]) (SGTUconst [32] y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RSRA)
		v0 := b.NewValue0(v.Pos, OpLOONG32RCMOVZ, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v1.AuxInt = int32ToAuxInt(31)
		v2 := b.NewValue0(v.Pos, OpLOONG32RSGTUconst, typ.Bool)
		v2.AuxInt = int32ToAuxInt(32)
		v2.AddArg(y)
		v0.AddArg3(y, v1, v2)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueLOONG32R_OpRsh32x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Rsh32x64 x (Const64 [c]))
	// cond: uint32(c) < 32
	// result: (SRAconst x [int32(c)])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint32(c) < 32) {
			break
		}
		v.reset(OpLOONG32RSRAconst)
		v.AuxInt = int32ToAuxInt(int32(c))
		v.AddArg(x)
		return true
	}
	// match: (Rsh32x64 x (Const64 [c]))
	// cond: uint32(c) >= 32
	// result: (SRAconst x [31])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint32(c) >= 32) {
			break
		}
		v.reset(OpLOONG32RSRAconst)
		v.AuxInt = int32ToAuxInt(31)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpRsh32x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x8 x y)
	// result: (SRA x ( CMOVZ <typ.UInt32> (ZeroExt8to32 y) (MOVWconst [31]) (SGTUconst [32] (ZeroExt8to32 y))))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RSRA)
		v0 := b.NewValue0(v.Pos, OpLOONG32RCMOVZ, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v2.AuxInt = int32ToAuxInt(31)
		v3 := b.NewValue0(v.Pos, OpLOONG32RSGTUconst, typ.Bool)
		v3.AuxInt = int32ToAuxInt(32)
		v3.AddArg(v1)
		v0.AddArg3(v1, v2, v3)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueLOONG32R_OpRsh8Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux16 <t> x y)
	// result: (CMOVZ (SRL <t> (ZeroExt8to32 x) (ZeroExt16to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt16to32 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpLOONG32RCMOVZ)
		v0 := b.NewValue0(v.Pos, OpLOONG32RSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v3 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v3.AuxInt = int32ToAuxInt(0)
		v4 := b.NewValue0(v.Pos, OpLOONG32RSGTUconst, typ.Bool)
		v4.AuxInt = int32ToAuxInt(32)
		v4.AddArg(v2)
		v.AddArg3(v0, v3, v4)
		return true
	}
}
func rewriteValueLOONG32R_OpRsh8Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux32 <t> x y)
	// result: (CMOVZ (SRL <t> (ZeroExt8to32 x) y) (MOVWconst [0]) (SGTUconst [32] y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpLOONG32RCMOVZ)
		v0 := b.NewValue0(v.Pos, OpLOONG32RSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg2(v1, y)
		v2 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v2.AuxInt = int32ToAuxInt(0)
		v3 := b.NewValue0(v.Pos, OpLOONG32RSGTUconst, typ.Bool)
		v3.AuxInt = int32ToAuxInt(32)
		v3.AddArg(y)
		v.AddArg3(v0, v2, v3)
		return true
	}
}
func rewriteValueLOONG32R_OpRsh8Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux64 x (Const64 [c]))
	// cond: uint32(c) < 8
	// result: (SRLconst (SLLconst <typ.UInt32> x [24]) [int32(c+24)])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint32(c) < 8) {
			break
		}
		v.reset(OpLOONG32RSRLconst)
		v.AuxInt = int32ToAuxInt(int32(c + 24))
		v0 := b.NewValue0(v.Pos, OpLOONG32RSLLconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(24)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8Ux64 _ (Const64 [c]))
	// cond: uint32(c) >= 8
	// result: (MOVWconst [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint32(c) >= 8) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(0)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpRsh8Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux8 <t> x y)
	// result: (CMOVZ (SRL <t> (ZeroExt8to32 x) (ZeroExt8to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt8to32 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpLOONG32RCMOVZ)
		v0 := b.NewValue0(v.Pos, OpLOONG32RSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v3 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v3.AuxInt = int32ToAuxInt(0)
		v4 := b.NewValue0(v.Pos, OpLOONG32RSGTUconst, typ.Bool)
		v4.AuxInt = int32ToAuxInt(32)
		v4.AddArg(v2)
		v.AddArg3(v0, v3, v4)
		return true
	}
}
func rewriteValueLOONG32R_OpRsh8x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x16 x y)
	// result: (SRA (SignExt8to32 x) ( CMOVZ <typ.UInt32> (ZeroExt16to32 y) (MOVWconst [31]) (SGTUconst [32] (ZeroExt16to32 y))))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpLOONG32RCMOVZ, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v3.AuxInt = int32ToAuxInt(31)
		v4 := b.NewValue0(v.Pos, OpLOONG32RSGTUconst, typ.Bool)
		v4.AuxInt = int32ToAuxInt(32)
		v4.AddArg(v2)
		v1.AddArg3(v2, v3, v4)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueLOONG32R_OpRsh8x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x32 x y)
	// result: (SRA (SignExt8to32 x) ( CMOVZ <typ.UInt32> y (MOVWconst [31]) (SGTUconst [32] y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpLOONG32RCMOVZ, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v2.AuxInt = int32ToAuxInt(31)
		v3 := b.NewValue0(v.Pos, OpLOONG32RSGTUconst, typ.Bool)
		v3.AuxInt = int32ToAuxInt(32)
		v3.AddArg(y)
		v1.AddArg3(y, v2, v3)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueLOONG32R_OpRsh8x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x64 x (Const64 [c]))
	// cond: uint32(c) < 8
	// result: (SRAconst (SLLconst <typ.UInt32> x [24]) [int32(c+24)])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint32(c) < 8) {
			break
		}
		v.reset(OpLOONG32RSRAconst)
		v.AuxInt = int32ToAuxInt(int32(c + 24))
		v0 := b.NewValue0(v.Pos, OpLOONG32RSLLconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(24)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8x64 x (Const64 [c]))
	// cond: uint32(c) >= 8
	// result: (SRAconst (SLLconst <typ.UInt32> x [24]) [31])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint32(c) >= 8) {
			break
		}
		v.reset(OpLOONG32RSRAconst)
		v.AuxInt = int32ToAuxInt(31)
		v0 := b.NewValue0(v.Pos, OpLOONG32RSLLconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(24)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpRsh8x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x8 x y)
	// result: (SRA (SignExt8to32 x) ( CMOVZ <typ.UInt32> (ZeroExt8to32 y) (MOVWconst [31]) (SGTUconst [32] (ZeroExt8to32 y))))
	for {
		x := v_0
		y := v_1
		v.reset(OpLOONG32RSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpLOONG32RCMOVZ, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v3.AuxInt = int32ToAuxInt(31)
		v4 := b.NewValue0(v.Pos, OpLOONG32RSGTUconst, typ.Bool)
		v4.AuxInt = int32ToAuxInt(32)
		v4.AddArg(v2)
		v1.AddArg3(v2, v3, v4)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueLOONG32R_OpSelect0(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Select0 (Add32carry <t> x y))
	// result: (ADD <t.FieldType(0)> x y)
	for {
		if v_0.Op != OpAdd32carry {
			break
		}
		t := v_0.Type
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpLOONG32RADD)
		v.Type = t.FieldType(0)
		v.AddArg2(x, y)
		return true
	}
	// match: (Select0 (Add32carrywithcarry <t> x y c))
	// result: (ADD <t.FieldType(0)> c (ADD <t.FieldType(0)> x y))
	for {
		if v_0.Op != OpAdd32carrywithcarry {
			break
		}
		t := v_0.Type
		c := v_0.Args[2]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpLOONG32RADD)
		v.Type = t.FieldType(0)
		v0 := b.NewValue0(v.Pos, OpLOONG32RADD, t.FieldType(0))
		v0.AddArg2(x, y)
		v.AddArg2(c, v0)
		return true
	}
	// match: (Select0 (Sub32carry <t> x y))
	// result: (SUB <t.FieldType(0)> x y)
	for {
		if v_0.Op != OpSub32carry {
			break
		}
		t := v_0.Type
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpLOONG32RSUB)
		v.Type = t.FieldType(0)
		v.AddArg2(x, y)
		return true
	}
	// match: (Select0 (MULTU (MOVWconst [0]) _ ))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpLOONG32RMULTU {
			break
		}
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpLOONG32RMOVWconst || auxIntToInt32(v_0_0.AuxInt) != 0 {
				continue
			}
			v.reset(OpLOONG32RMOVWconst)
			v.AuxInt = int32ToAuxInt(0)
			return true
		}
		break
	}
	// match: (Select0 (MULTU (MOVWconst [1]) _ ))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpLOONG32RMULTU {
			break
		}
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpLOONG32RMOVWconst || auxIntToInt32(v_0_0.AuxInt) != 1 {
				continue
			}
			v.reset(OpLOONG32RMOVWconst)
			v.AuxInt = int32ToAuxInt(0)
			return true
		}
		break
	}
	// match: (Select0 (MULTU (MOVWconst [-1]) x ))
	// result: (CMOVZ (ADDconst <x.Type> [-1] x) (MOVWconst [0]) x)
	for {
		if v_0.Op != OpLOONG32RMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpLOONG32RMOVWconst || auxIntToInt32(v_0_0.AuxInt) != -1 {
				continue
			}
			x := v_0_1
			v.reset(OpLOONG32RCMOVZ)
			v0 := b.NewValue0(v.Pos, OpLOONG32RADDconst, x.Type)
			v0.AuxInt = int32ToAuxInt(-1)
			v0.AddArg(x)
			v1 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
			v1.AuxInt = int32ToAuxInt(0)
			v.AddArg3(v0, v1, x)
			return true
		}
		break
	}
	// match: (Select0 (MULTU (MOVWconst [c]) x ))
	// cond: isUnsignedPowerOfTwo(uint32(c))
	// result: (SRLconst [int32(32-log32u(uint32(c)))] x)
	for {
		if v_0.Op != OpLOONG32RMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpLOONG32RMOVWconst {
				continue
			}
			c := auxIntToInt32(v_0_0.AuxInt)
			x := v_0_1
			if !(isUnsignedPowerOfTwo(uint32(c))) {
				continue
			}
			v.reset(OpLOONG32RSRLconst)
			v.AuxInt = int32ToAuxInt(int32(32 - log32u(uint32(c))))
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (Select0 (MULTU (MOVWconst [c]) (MOVWconst [d])))
	// result: (MOVWconst [int32((int64(uint32(c))*int64(uint32(d)))>>32)])
	for {
		if v_0.Op != OpLOONG32RMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpLOONG32RMOVWconst {
				continue
			}
			c := auxIntToInt32(v_0_0.AuxInt)
			if v_0_1.Op != OpLOONG32RMOVWconst {
				continue
			}
			d := auxIntToInt32(v_0_1.AuxInt)
			v.reset(OpLOONG32RMOVWconst)
			v.AuxInt = int32ToAuxInt(int32((int64(uint32(c)) * int64(uint32(d))) >> 32))
			return true
		}
		break
	}
	// match: (Select0 (DIV (MOVWconst [c]) (MOVWconst [d])))
	// cond: d != 0
	// result: (MOVWconst [c%d])
	for {
		if v_0.Op != OpLOONG32RDIV {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpLOONG32RMOVWconst {
			break
		}
		c := auxIntToInt32(v_0_0.AuxInt)
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpLOONG32RMOVWconst {
			break
		}
		d := auxIntToInt32(v_0_1.AuxInt)
		if !(d != 0) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(c % d)
		return true
	}
	// match: (Select0 (DIVU (MOVWconst [c]) (MOVWconst [d])))
	// cond: d != 0
	// result: (MOVWconst [int32(uint32(c)%uint32(d))])
	for {
		if v_0.Op != OpLOONG32RDIVU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpLOONG32RMOVWconst {
			break
		}
		c := auxIntToInt32(v_0_0.AuxInt)
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpLOONG32RMOVWconst {
			break
		}
		d := auxIntToInt32(v_0_1.AuxInt)
		if !(d != 0) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(int32(uint32(c) % uint32(d)))
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpSelect1(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Select1 (Add32carry <t> x y))
	// result: (SGTU <typ.Bool> x (ADD <t.FieldType(0)> x y))
	for {
		if v_0.Op != OpAdd32carry {
			break
		}
		t := v_0.Type
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpLOONG32RSGTU)
		v.Type = typ.Bool
		v0 := b.NewValue0(v.Pos, OpLOONG32RADD, t.FieldType(0))
		v0.AddArg2(x, y)
		v.AddArg2(x, v0)
		return true
	}
	// match: (Select1 (Add32carrywithcarry <t> x y c))
	// result: (OR <typ.Bool> (SGTU <typ.Bool> x xy:(ADD <t.FieldType(0)> x y)) (SGTU <typ.Bool> xy (ADD <t.FieldType(0)> c xy)))
	for {
		if v_0.Op != OpAdd32carrywithcarry {
			break
		}
		t := v_0.Type
		c := v_0.Args[2]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpLOONG32ROR)
		v.Type = typ.Bool
		v0 := b.NewValue0(v.Pos, OpLOONG32RSGTU, typ.Bool)
		xy := b.NewValue0(v.Pos, OpLOONG32RADD, t.FieldType(0))
		xy.AddArg2(x, y)
		v0.AddArg2(x, xy)
		v2 := b.NewValue0(v.Pos, OpLOONG32RSGTU, typ.Bool)
		v3 := b.NewValue0(v.Pos, OpLOONG32RADD, t.FieldType(0))
		v3.AddArg2(c, xy)
		v2.AddArg2(xy, v3)
		v.AddArg2(v0, v2)
		return true
	}
	// match: (Select1 (Sub32carry <t> x y))
	// result: (SGTU <typ.Bool> (SUB <t.FieldType(0)> x y) x)
	for {
		if v_0.Op != OpSub32carry {
			break
		}
		t := v_0.Type
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpLOONG32RSGTU)
		v.Type = typ.Bool
		v0 := b.NewValue0(v.Pos, OpLOONG32RSUB, t.FieldType(0))
		v0.AddArg2(x, y)
		v.AddArg2(v0, x)
		return true
	}
	// match: (Select1 (MULTU (MOVWconst [0]) _ ))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpLOONG32RMULTU {
			break
		}
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpLOONG32RMOVWconst || auxIntToInt32(v_0_0.AuxInt) != 0 {
				continue
			}
			v.reset(OpLOONG32RMOVWconst)
			v.AuxInt = int32ToAuxInt(0)
			return true
		}
		break
	}
	// match: (Select1 (MULTU (MOVWconst [1]) x ))
	// result: x
	for {
		if v_0.Op != OpLOONG32RMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpLOONG32RMOVWconst || auxIntToInt32(v_0_0.AuxInt) != 1 {
				continue
			}
			x := v_0_1
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (Select1 (MULTU (MOVWconst [-1]) x ))
	// result: (NEG <x.Type> x)
	for {
		if v_0.Op != OpLOONG32RMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpLOONG32RMOVWconst || auxIntToInt32(v_0_0.AuxInt) != -1 {
				continue
			}
			x := v_0_1
			v.reset(OpLOONG32RNEG)
			v.Type = x.Type
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (Select1 (MULTU (MOVWconst [c]) x ))
	// cond: isUnsignedPowerOfTwo(uint32(c))
	// result: (SLLconst [int32(log32u(uint32(c)))] x)
	for {
		if v_0.Op != OpLOONG32RMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpLOONG32RMOVWconst {
				continue
			}
			c := auxIntToInt32(v_0_0.AuxInt)
			x := v_0_1
			if !(isUnsignedPowerOfTwo(uint32(c))) {
				continue
			}
			v.reset(OpLOONG32RSLLconst)
			v.AuxInt = int32ToAuxInt(int32(log32u(uint32(c))))
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (Select1 (MULTU (MOVWconst [c]) (MOVWconst [d])))
	// result: (MOVWconst [int32(uint32(c)*uint32(d))])
	for {
		if v_0.Op != OpLOONG32RMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpLOONG32RMOVWconst {
				continue
			}
			c := auxIntToInt32(v_0_0.AuxInt)
			if v_0_1.Op != OpLOONG32RMOVWconst {
				continue
			}
			d := auxIntToInt32(v_0_1.AuxInt)
			v.reset(OpLOONG32RMOVWconst)
			v.AuxInt = int32ToAuxInt(int32(uint32(c) * uint32(d)))
			return true
		}
		break
	}
	// match: (Select1 (DIV (MOVWconst [c]) (MOVWconst [d])))
	// cond: d != 0
	// result: (MOVWconst [c/d])
	for {
		if v_0.Op != OpLOONG32RDIV {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpLOONG32RMOVWconst {
			break
		}
		c := auxIntToInt32(v_0_0.AuxInt)
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpLOONG32RMOVWconst {
			break
		}
		d := auxIntToInt32(v_0_1.AuxInt)
		if !(d != 0) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(c / d)
		return true
	}
	// match: (Select1 (DIVU (MOVWconst [c]) (MOVWconst [d])))
	// cond: d != 0
	// result: (MOVWconst [int32(uint32(c)/uint32(d))])
	for {
		if v_0.Op != OpLOONG32RDIVU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpLOONG32RMOVWconst {
			break
		}
		c := auxIntToInt32(v_0_0.AuxInt)
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpLOONG32RMOVWconst {
			break
		}
		d := auxIntToInt32(v_0_1.AuxInt)
		if !(d != 0) {
			break
		}
		v.reset(OpLOONG32RMOVWconst)
		v.AuxInt = int32ToAuxInt(int32(uint32(c) / uint32(d)))
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpSignmask(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Signmask x)
	// result: (SRAconst x [31])
	for {
		x := v_0
		v.reset(OpLOONG32RSRAconst)
		v.AuxInt = int32ToAuxInt(31)
		v.AddArg(x)
		return true
	}
}
func rewriteValueLOONG32R_OpSlicemask(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (Slicemask <t> x)
	// result: (SRAconst (NEG <t> x) [31])
	for {
		t := v.Type
		x := v_0
		v.reset(OpLOONG32RSRAconst)
		v.AuxInt = int32ToAuxInt(31)
		v0 := b.NewValue0(v.Pos, OpLOONG32RNEG, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueLOONG32R_OpStore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Store {t} ptr val mem)
	// cond: t.Size() == 1
	// result: (MOVBstore ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.Size() == 1) {
			break
		}
		v.reset(OpLOONG32RMOVBstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.Size() == 2
	// result: (MOVHstore ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.Size() == 2) {
			break
		}
		v.reset(OpLOONG32RMOVHstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.Size() == 4 && !t.IsFloat()
	// result: (MOVWstore ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.Size() == 4 && !t.IsFloat()) {
			break
		}
		v.reset(OpLOONG32RMOVWstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.Size() == 4 && t.IsFloat()
	// result: (MOVFstore ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.Size() == 4 && t.IsFloat()) {
			break
		}
		v.reset(OpLOONG32RMOVFstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.Size() == 8 && t.IsFloat()
	// result: (MOVDstore ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.Size() == 8 && t.IsFloat()) {
			break
		}
		v.reset(OpLOONG32RMOVDstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpSub32withcarry(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Sub32withcarry <t> x y c)
	// result: (SUB (SUB <t> x y) c)
	for {
		t := v.Type
		x := v_0
		y := v_1
		c := v_2
		v.reset(OpLOONG32RSUB)
		v0 := b.NewValue0(v.Pos, OpLOONG32RSUB, t)
		v0.AddArg2(x, y)
		v.AddArg2(v0, c)
		return true
	}
}
func rewriteValueLOONG32R_OpZero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (Zero [0] _ mem)
	// result: mem
	for {
		if auxIntToInt64(v.AuxInt) != 0 {
			break
		}
		mem := v_1
		v.copyOf(mem)
		return true
	}
	// match: (Zero [1] ptr mem)
	// result: (MOVBstore ptr (MOVWconst [0]) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 1 {
			break
		}
		ptr := v_0
		mem := v_1
		v.reset(OpLOONG32RMOVBstore)
		v0 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(0)
		v.AddArg3(ptr, v0, mem)
		return true
	}
	// match: (Zero [2] {t} ptr mem)
	// cond: t.Alignment()%2 == 0
	// result: (MOVHstore ptr (MOVWconst [0]) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 2 {
			break
		}
		t := auxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(t.Alignment()%2 == 0) {
			break
		}
		v.reset(OpLOONG32RMOVHstore)
		v0 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(0)
		v.AddArg3(ptr, v0, mem)
		return true
	}
	// match: (Zero [2] ptr mem)
	// result: (MOVBstore [1] ptr (MOVWconst [0]) (MOVBstore [0] ptr (MOVWconst [0]) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 2 {
			break
		}
		ptr := v_0
		mem := v_1
		v.reset(OpLOONG32RMOVBstore)
		v.AuxInt = int32ToAuxInt(1)
		v0 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVBstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(0)
		v1.AddArg3(ptr, v0, mem)
		v.AddArg3(ptr, v0, v1)
		return true
	}
	// match: (Zero [4] {t} ptr mem)
	// cond: t.Alignment()%4 == 0
	// result: (MOVWstore ptr (MOVWconst [0]) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 4 {
			break
		}
		t := auxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(t.Alignment()%4 == 0) {
			break
		}
		v.reset(OpLOONG32RMOVWstore)
		v0 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(0)
		v.AddArg3(ptr, v0, mem)
		return true
	}
	// match: (Zero [4] {t} ptr mem)
	// cond: t.Alignment()%2 == 0
	// result: (MOVHstore [2] ptr (MOVWconst [0]) (MOVHstore [0] ptr (MOVWconst [0]) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 4 {
			break
		}
		t := auxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(t.Alignment()%2 == 0) {
			break
		}
		v.reset(OpLOONG32RMOVHstore)
		v.AuxInt = int32ToAuxInt(2)
		v0 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVHstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(0)
		v1.AddArg3(ptr, v0, mem)
		v.AddArg3(ptr, v0, v1)
		return true
	}
	// match: (Zero [4] ptr mem)
	// result: (MOVBstore [3] ptr (MOVWconst [0]) (MOVBstore [2] ptr (MOVWconst [0]) (MOVBstore [1] ptr (MOVWconst [0]) (MOVBstore [0] ptr (MOVWconst [0]) mem))))
	for {
		if auxIntToInt64(v.AuxInt) != 4 {
			break
		}
		ptr := v_0
		mem := v_1
		v.reset(OpLOONG32RMOVBstore)
		v.AuxInt = int32ToAuxInt(3)
		v0 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVBstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(2)
		v2 := b.NewValue0(v.Pos, OpLOONG32RMOVBstore, types.TypeMem)
		v2.AuxInt = int32ToAuxInt(1)
		v3 := b.NewValue0(v.Pos, OpLOONG32RMOVBstore, types.TypeMem)
		v3.AuxInt = int32ToAuxInt(0)
		v3.AddArg3(ptr, v0, mem)
		v2.AddArg3(ptr, v0, v3)
		v1.AddArg3(ptr, v0, v2)
		v.AddArg3(ptr, v0, v1)
		return true
	}
	// match: (Zero [3] ptr mem)
	// result: (MOVBstore [2] ptr (MOVWconst [0]) (MOVBstore [1] ptr (MOVWconst [0]) (MOVBstore [0] ptr (MOVWconst [0]) mem)))
	for {
		if auxIntToInt64(v.AuxInt) != 3 {
			break
		}
		ptr := v_0
		mem := v_1
		v.reset(OpLOONG32RMOVBstore)
		v.AuxInt = int32ToAuxInt(2)
		v0 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVBstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(1)
		v2 := b.NewValue0(v.Pos, OpLOONG32RMOVBstore, types.TypeMem)
		v2.AuxInt = int32ToAuxInt(0)
		v2.AddArg3(ptr, v0, mem)
		v1.AddArg3(ptr, v0, v2)
		v.AddArg3(ptr, v0, v1)
		return true
	}
	// match: (Zero [6] {t} ptr mem)
	// cond: t.Alignment()%2 == 0
	// result: (MOVHstore [4] ptr (MOVWconst [0]) (MOVHstore [2] ptr (MOVWconst [0]) (MOVHstore [0] ptr (MOVWconst [0]) mem)))
	for {
		if auxIntToInt64(v.AuxInt) != 6 {
			break
		}
		t := auxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(t.Alignment()%2 == 0) {
			break
		}
		v.reset(OpLOONG32RMOVHstore)
		v.AuxInt = int32ToAuxInt(4)
		v0 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVHstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(2)
		v2 := b.NewValue0(v.Pos, OpLOONG32RMOVHstore, types.TypeMem)
		v2.AuxInt = int32ToAuxInt(0)
		v2.AddArg3(ptr, v0, mem)
		v1.AddArg3(ptr, v0, v2)
		v.AddArg3(ptr, v0, v1)
		return true
	}
	// match: (Zero [8] {t} ptr mem)
	// cond: t.Alignment()%4 == 0
	// result: (MOVWstore [4] ptr (MOVWconst [0]) (MOVWstore [0] ptr (MOVWconst [0]) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 8 {
			break
		}
		t := auxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(t.Alignment()%4 == 0) {
			break
		}
		v.reset(OpLOONG32RMOVWstore)
		v.AuxInt = int32ToAuxInt(4)
		v0 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVWstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(0)
		v1.AddArg3(ptr, v0, mem)
		v.AddArg3(ptr, v0, v1)
		return true
	}
	// match: (Zero [12] {t} ptr mem)
	// cond: t.Alignment()%4 == 0
	// result: (MOVWstore [8] ptr (MOVWconst [0]) (MOVWstore [4] ptr (MOVWconst [0]) (MOVWstore [0] ptr (MOVWconst [0]) mem)))
	for {
		if auxIntToInt64(v.AuxInt) != 12 {
			break
		}
		t := auxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(t.Alignment()%4 == 0) {
			break
		}
		v.reset(OpLOONG32RMOVWstore)
		v.AuxInt = int32ToAuxInt(8)
		v0 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVWstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(4)
		v2 := b.NewValue0(v.Pos, OpLOONG32RMOVWstore, types.TypeMem)
		v2.AuxInt = int32ToAuxInt(0)
		v2.AddArg3(ptr, v0, mem)
		v1.AddArg3(ptr, v0, v2)
		v.AddArg3(ptr, v0, v1)
		return true
	}
	// match: (Zero [16] {t} ptr mem)
	// cond: t.Alignment()%4 == 0
	// result: (MOVWstore [12] ptr (MOVWconst [0]) (MOVWstore [8] ptr (MOVWconst [0]) (MOVWstore [4] ptr (MOVWconst [0]) (MOVWstore [0] ptr (MOVWconst [0]) mem))))
	for {
		if auxIntToInt64(v.AuxInt) != 16 {
			break
		}
		t := auxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(t.Alignment()%4 == 0) {
			break
		}
		v.reset(OpLOONG32RMOVWstore)
		v.AuxInt = int32ToAuxInt(12)
		v0 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v0.AuxInt = int32ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVWstore, types.TypeMem)
		v1.AuxInt = int32ToAuxInt(8)
		v2 := b.NewValue0(v.Pos, OpLOONG32RMOVWstore, types.TypeMem)
		v2.AuxInt = int32ToAuxInt(4)
		v3 := b.NewValue0(v.Pos, OpLOONG32RMOVWstore, types.TypeMem)
		v3.AuxInt = int32ToAuxInt(0)
		v3.AddArg3(ptr, v0, mem)
		v2.AddArg3(ptr, v0, v3)
		v1.AddArg3(ptr, v0, v2)
		v.AddArg3(ptr, v0, v1)
		return true
	}
	// match: (Zero [s] {t} ptr mem)
	// cond: (s > 16 || t.Alignment()%4 != 0)
	// result: (LoweredZero [int32(t.Alignment())] ptr (ADDconst <ptr.Type> ptr [int32(s-moveSize(t.Alignment(), config))]) mem)
	for {
		s := auxIntToInt64(v.AuxInt)
		t := auxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(s > 16 || t.Alignment()%4 != 0) {
			break
		}
		v.reset(OpLOONG32RLoweredZero)
		v.AuxInt = int32ToAuxInt(int32(t.Alignment()))
		v0 := b.NewValue0(v.Pos, OpLOONG32RADDconst, ptr.Type)
		v0.AuxInt = int32ToAuxInt(int32(s - moveSize(t.Alignment(), config)))
		v0.AddArg(ptr)
		v.AddArg3(ptr, v0, mem)
		return true
	}
	return false
}
func rewriteValueLOONG32R_OpZeromask(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Zeromask x)
	// result: (NEG (SGTU x (MOVWconst [0])))
	for {
		x := v_0
		v.reset(OpLOONG32RNEG)
		v0 := b.NewValue0(v.Pos, OpLOONG32RSGTU, typ.Bool)
		v1 := b.NewValue0(v.Pos, OpLOONG32RMOVWconst, typ.UInt32)
		v1.AuxInt = int32ToAuxInt(0)
		v0.AddArg2(x, v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteBlockLOONG32R(b *Block) bool {
	switch b.Kind {
	case BlockLOONG32REQ:
		// match: (EQ (FPFlagTrue cmp) yes no)
		// result: (FPF cmp yes no)
		for b.Controls[0].Op == OpLOONG32RFPFlagTrue {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.resetWithControl(BlockLOONG32RFPF, cmp)
			return true
		}
		// match: (EQ (FPFlagFalse cmp) yes no)
		// result: (FPT cmp yes no)
		for b.Controls[0].Op == OpLOONG32RFPFlagFalse {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.resetWithControl(BlockLOONG32RFPT, cmp)
			return true
		}
		// match: (EQ (XORconst [1] cmp:(SGT _ _)) yes no)
		// result: (NE cmp yes no)
		for b.Controls[0].Op == OpLOONG32RXORconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpLOONG32RSGT {
				break
			}
			b.resetWithControl(BlockLOONG32RNE, cmp)
			return true
		}
		// match: (EQ (XORconst [1] cmp:(SGTU _ _)) yes no)
		// result: (NE cmp yes no)
		for b.Controls[0].Op == OpLOONG32RXORconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpLOONG32RSGTU {
				break
			}
			b.resetWithControl(BlockLOONG32RNE, cmp)
			return true
		}
		// match: (EQ (XORconst [1] cmp:(SGTconst _)) yes no)
		// result: (NE cmp yes no)
		for b.Controls[0].Op == OpLOONG32RXORconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpLOONG32RSGTconst {
				break
			}
			b.resetWithControl(BlockLOONG32RNE, cmp)
			return true
		}
		// match: (EQ (XORconst [1] cmp:(SGTUconst _)) yes no)
		// result: (NE cmp yes no)
		for b.Controls[0].Op == OpLOONG32RXORconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpLOONG32RSGTUconst {
				break
			}
			b.resetWithControl(BlockLOONG32RNE, cmp)
			return true
		}
		// match: (EQ (XORconst [1] cmp:(SGTzero _)) yes no)
		// result: (NE cmp yes no)
		for b.Controls[0].Op == OpLOONG32RXORconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpLOONG32RSGTzero {
				break
			}
			b.resetWithControl(BlockLOONG32RNE, cmp)
			return true
		}
		// match: (EQ (XORconst [1] cmp:(SGTUzero _)) yes no)
		// result: (NE cmp yes no)
		for b.Controls[0].Op == OpLOONG32RXORconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpLOONG32RSGTUzero {
				break
			}
			b.resetWithControl(BlockLOONG32RNE, cmp)
			return true
		}
		// match: (EQ (SGTUconst [1] x) yes no)
		// result: (NE x yes no)
		for b.Controls[0].Op == OpLOONG32RSGTUconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 1 {
				break
			}
			x := v_0.Args[0]
			b.resetWithControl(BlockLOONG32RNE, x)
			return true
		}
		// match: (EQ (SGTUzero x) yes no)
		// result: (EQ x yes no)
		for b.Controls[0].Op == OpLOONG32RSGTUzero {
			v_0 := b.Controls[0]
			x := v_0.Args[0]
			b.resetWithControl(BlockLOONG32REQ, x)
			return true
		}
		// match: (EQ (SGTconst [0] x) yes no)
		// result: (GEZ x yes no)
		for b.Controls[0].Op == OpLOONG32RSGTconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			x := v_0.Args[0]
			b.resetWithControl(BlockLOONG32RGEZ, x)
			return true
		}
		// match: (EQ (SGTzero x) yes no)
		// result: (LEZ x yes no)
		for b.Controls[0].Op == OpLOONG32RSGTzero {
			v_0 := b.Controls[0]
			x := v_0.Args[0]
			b.resetWithControl(BlockLOONG32RLEZ, x)
			return true
		}
		// match: (EQ (MOVWconst [0]) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpLOONG32RMOVWconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (EQ (MOVWconst [c]) yes no)
		// cond: c != 0
		// result: (First no yes)
		for b.Controls[0].Op == OpLOONG32RMOVWconst {
			v_0 := b.Controls[0]
			c := auxIntToInt32(v_0.AuxInt)
			if !(c != 0) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
	case BlockLOONG32RGEZ:
		// match: (GEZ (MOVWconst [c]) yes no)
		// cond: c >= 0
		// result: (First yes no)
		for b.Controls[0].Op == OpLOONG32RMOVWconst {
			v_0 := b.Controls[0]
			c := auxIntToInt32(v_0.AuxInt)
			if !(c >= 0) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (GEZ (MOVWconst [c]) yes no)
		// cond: c < 0
		// result: (First no yes)
		for b.Controls[0].Op == OpLOONG32RMOVWconst {
			v_0 := b.Controls[0]
			c := auxIntToInt32(v_0.AuxInt)
			if !(c < 0) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
	case BlockLOONG32RGTZ:
		// match: (GTZ (MOVWconst [c]) yes no)
		// cond: c > 0
		// result: (First yes no)
		for b.Controls[0].Op == OpLOONG32RMOVWconst {
			v_0 := b.Controls[0]
			c := auxIntToInt32(v_0.AuxInt)
			if !(c > 0) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (GTZ (MOVWconst [c]) yes no)
		// cond: c <= 0
		// result: (First no yes)
		for b.Controls[0].Op == OpLOONG32RMOVWconst {
			v_0 := b.Controls[0]
			c := auxIntToInt32(v_0.AuxInt)
			if !(c <= 0) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
	case BlockIf:
		// match: (If cond yes no)
		// result: (NE cond yes no)
		for {
			cond := b.Controls[0]
			b.resetWithControl(BlockLOONG32RNE, cond)
			return true
		}
	case BlockLOONG32RLEZ:
		// match: (LEZ (MOVWconst [c]) yes no)
		// cond: c <= 0
		// result: (First yes no)
		for b.Controls[0].Op == OpLOONG32RMOVWconst {
			v_0 := b.Controls[0]
			c := auxIntToInt32(v_0.AuxInt)
			if !(c <= 0) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (LEZ (MOVWconst [c]) yes no)
		// cond: c > 0
		// result: (First no yes)
		for b.Controls[0].Op == OpLOONG32RMOVWconst {
			v_0 := b.Controls[0]
			c := auxIntToInt32(v_0.AuxInt)
			if !(c > 0) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
	case BlockLOONG32RLTZ:
		// match: (LTZ (MOVWconst [c]) yes no)
		// cond: c < 0
		// result: (First yes no)
		for b.Controls[0].Op == OpLOONG32RMOVWconst {
			v_0 := b.Controls[0]
			c := auxIntToInt32(v_0.AuxInt)
			if !(c < 0) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (LTZ (MOVWconst [c]) yes no)
		// cond: c >= 0
		// result: (First no yes)
		for b.Controls[0].Op == OpLOONG32RMOVWconst {
			v_0 := b.Controls[0]
			c := auxIntToInt32(v_0.AuxInt)
			if !(c >= 0) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
	case BlockLOONG32RNE:
		// match: (NE (FPFlagTrue cmp) yes no)
		// result: (FPT cmp yes no)
		for b.Controls[0].Op == OpLOONG32RFPFlagTrue {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.resetWithControl(BlockLOONG32RFPT, cmp)
			return true
		}
		// match: (NE (FPFlagFalse cmp) yes no)
		// result: (FPF cmp yes no)
		for b.Controls[0].Op == OpLOONG32RFPFlagFalse {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.resetWithControl(BlockLOONG32RFPF, cmp)
			return true
		}
		// match: (NE (XORconst [1] cmp:(SGT _ _)) yes no)
		// result: (EQ cmp yes no)
		for b.Controls[0].Op == OpLOONG32RXORconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpLOONG32RSGT {
				break
			}
			b.resetWithControl(BlockLOONG32REQ, cmp)
			return true
		}
		// match: (NE (XORconst [1] cmp:(SGTU _ _)) yes no)
		// result: (EQ cmp yes no)
		for b.Controls[0].Op == OpLOONG32RXORconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpLOONG32RSGTU {
				break
			}
			b.resetWithControl(BlockLOONG32REQ, cmp)
			return true
		}
		// match: (NE (XORconst [1] cmp:(SGTconst _)) yes no)
		// result: (EQ cmp yes no)
		for b.Controls[0].Op == OpLOONG32RXORconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpLOONG32RSGTconst {
				break
			}
			b.resetWithControl(BlockLOONG32REQ, cmp)
			return true
		}
		// match: (NE (XORconst [1] cmp:(SGTUconst _)) yes no)
		// result: (EQ cmp yes no)
		for b.Controls[0].Op == OpLOONG32RXORconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpLOONG32RSGTUconst {
				break
			}
			b.resetWithControl(BlockLOONG32REQ, cmp)
			return true
		}
		// match: (NE (XORconst [1] cmp:(SGTzero _)) yes no)
		// result: (EQ cmp yes no)
		for b.Controls[0].Op == OpLOONG32RXORconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpLOONG32RSGTzero {
				break
			}
			b.resetWithControl(BlockLOONG32REQ, cmp)
			return true
		}
		// match: (NE (XORconst [1] cmp:(SGTUzero _)) yes no)
		// result: (EQ cmp yes no)
		for b.Controls[0].Op == OpLOONG32RXORconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpLOONG32RSGTUzero {
				break
			}
			b.resetWithControl(BlockLOONG32REQ, cmp)
			return true
		}
		// match: (NE (SGTUconst [1] x) yes no)
		// result: (EQ x yes no)
		for b.Controls[0].Op == OpLOONG32RSGTUconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 1 {
				break
			}
			x := v_0.Args[0]
			b.resetWithControl(BlockLOONG32REQ, x)
			return true
		}
		// match: (NE (SGTUzero x) yes no)
		// result: (NE x yes no)
		for b.Controls[0].Op == OpLOONG32RSGTUzero {
			v_0 := b.Controls[0]
			x := v_0.Args[0]
			b.resetWithControl(BlockLOONG32RNE, x)
			return true
		}
		// match: (NE (SGTconst [0] x) yes no)
		// result: (LTZ x yes no)
		for b.Controls[0].Op == OpLOONG32RSGTconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			x := v_0.Args[0]
			b.resetWithControl(BlockLOONG32RLTZ, x)
			return true
		}
		// match: (NE (SGTzero x) yes no)
		// result: (GTZ x yes no)
		for b.Controls[0].Op == OpLOONG32RSGTzero {
			v_0 := b.Controls[0]
			x := v_0.Args[0]
			b.resetWithControl(BlockLOONG32RGTZ, x)
			return true
		}
		// match: (NE (MOVWconst [0]) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpLOONG32RMOVWconst {
			v_0 := b.Controls[0]
			if auxIntToInt32(v_0.AuxInt) != 0 {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (NE (MOVWconst [c]) yes no)
		// cond: c != 0
		// result: (First yes no)
		for b.Controls[0].Op == OpLOONG32RMOVWconst {
			v_0 := b.Controls[0]
			c := auxIntToInt32(v_0.AuxInt)
			if !(c != 0) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
	}
	return false
}
