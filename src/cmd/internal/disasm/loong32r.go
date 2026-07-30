// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package disasm

import (
	"encoding/binary"
	"fmt"
)

// This decoder is intentionally independent of the LA64 decoder. The opcode
// masks and operand layouts below come from the LA32R reference manual.

type loong32rOpcode struct {
	plan9 string
	gnu   string
}

var loong32rRRR = map[uint32]loong32rOpcode{
	0x020: {"ADD", "add.w"},
	0x022: {"SUB", "sub.w"},
	0x024: {"SGT", "slt"},
	0x025: {"SGTU", "sltu"},
	0x028: {"NOR", "nor"},
	0x029: {"AND", "and"},
	0x02a: {"OR", "or"},
	0x02b: {"XOR", "xor"},
	0x02e: {"SLL", "sll.w"},
	0x02f: {"SRL", "srl.w"},
	0x030: {"SRA", "sra.w"},
	0x038: {"MUL", "mul.w"},
	0x039: {"MULH", "mulh.w"},
	0x03a: {"MULHU", "mulh.wu"},
	0x040: {"DIV", "div.w"},
	0x041: {"REM", "mod.w"},
	0x042: {"DIVU", "div.wu"},
	0x043: {"REMU", "mod.wu"},

	0x201: {"ADDF", "fadd.s"},
	0x202: {"ADDD", "fadd.d"},
	0x205: {"SUBF", "fsub.s"},
	0x206: {"SUBD", "fsub.d"},
	0x209: {"MULF", "fmul.s"},
	0x20a: {"MULD", "fmul.d"},
	0x20d: {"DIVF", "fdiv.s"},
	0x20e: {"DIVD", "fdiv.d"},
	0x211: {"FMAXF", "fmax.s"},
	0x212: {"FMAXD", "fmax.d"},
	0x215: {"FMINF", "fmin.s"},
	0x216: {"FMIND", "fmin.d"},
	0x219: {"FMAXAF", "fmaxa.s"},
	0x21a: {"FMAXAD", "fmaxa.d"},
	0x21d: {"FMINAF", "fmina.s"},
	0x21e: {"FMINAD", "fmina.d"},
	0x225: {"FCOPYSGF", "fcopysign.s"},
	0x226: {"FCOPYSGD", "fcopysign.d"},

	0x0c1<<5 | 0x04: {"CMPEQF", "fcmp.ceq.s"},
	0x0c2<<5 | 0x04: {"CMPEQD", "fcmp.ceq.d"},
	0x0c1<<5 | 0x07: {"CMPGEF", "fcmp.sle.s"},
	0x0c2<<5 | 0x07: {"CMPGED", "fcmp.sle.d"},
	0x0c1<<5 | 0x03: {"CMPGTF", "fcmp.slt.s"},
	0x0c2<<5 | 0x03: {"CMPGTD", "fcmp.slt.d"},
}

var loong32rRR = map[uint32]loong32rOpcode{
	0x0018: {"RDTIMELW", "rdtimel.w"},
	0x0019: {"RDTIMEHW", "rdtimeh.w"},
	0x4501: {"ABSF", "fabs.s"},
	0x4502: {"ABSD", "fabs.d"},
	0x4505: {"NEGF", "fneg.s"},
	0x4506: {"NEGD", "fneg.d"},
	0x450d: {"FCLASSF", "fclass.s"},
	0x450e: {"FCLASSD", "fclass.d"},
	0x4511: {"SQRTF", "fsqrt.s"},
	0x4512: {"SQRTD", "fsqrt.d"},
	0x4515: {"FRECIPF", "frecip.s"},
	0x4516: {"FRECIPD", "frecip.d"},
	0x4519: {"FRSQRTF", "frsqrt.s"},
	0x451a: {"FRSQRTD", "frsqrt.d"},
	0x4525: {"MOVF", "fmov.s"},
	0x4526: {"MOVD", "fmov.d"},
	0x4529: {"MOVW", "movgr2fr.w"},
	0x452d: {"MOVW", "movfr2gr.s"},
	0x4646: {"MOVDF", "fcvt.s.d"},
	0x4649: {"MOVFD", "fcvt.d.s"},
	0x46a1: {"TRUNCFW", "ftintrz.w.s"},
	0x46a2: {"TRUNCDW", "ftintrz.w.d"},
	0x4744: {"MOVWF", "ffint.s.w"},
	0x4748: {"MOVWD", "ffint.d.w"},
}

var loong32rFMA = map[uint32]loong32rOpcode{
	0x081: {"FMADDF", "fmadd.s"},
	0x082: {"FMADDD", "fmadd.d"},
	0x085: {"FMSUBF", "fmsub.s"},
	0x086: {"FMSUBD", "fmsub.d"},
	0x089: {"FNMADDF", "fnmadd.s"},
	0x08a: {"FNMADDD", "fnmadd.d"},
	0x08d: {"FNMSUBF", "fnmsub.s"},
	0x08e: {"FNMSUBD", "fnmsub.d"},
}

func loong32rSignExtend(v uint32, bits uint) int64 {
	return int64(int32(v<<(32-bits)) >> (32 - bits))
}

func loong32rR(r uint32) string   { return fmt.Sprintf("R%d", r) }
func loong32rF(r uint32) string   { return fmt.Sprintf("F%d", r) }
func loong32rGR(r uint32) string  { return fmt.Sprintf("$r%d", r) }
func loong32rGF(r uint32) string  { return fmt.Sprintf("$f%d", r) }
func loong32rFCC(r uint32) string { return fmt.Sprintf("FCC%d", r) }

func loong32rTarget(pc uint64, delta int64, lookup lookupFunc) (string, string) {
	target := uint64(int64(pc) + delta)
	if name, base := lookup(target); name != "" && base == target {
		return name + "(SB)", name
	}
	return fmt.Sprintf("%d(PC)", delta/4), fmt.Sprintf("%#x", target)
}

func loong32rText(plan9, gnu string, gnuAsm bool) string {
	if gnuAsm {
		return fmt.Sprintf("%-36s // %s", plan9, gnu)
	}
	return plan9
}

func disasm_loong32r(code []byte, pc uint64, lookup lookupFunc, ord binary.ByteOrder, gnuAsm bool) (string, int) {
	if len(code) < 4 {
		return "?", 4
	}
	w := ord.Uint32(code)
	rd, rj, rk := w&31, w>>5&31, w>>10&31

	if w == 0x03400000 {
		return loong32rText("NOOP", "nop", gnuAsm), 4
	}

	// Direct, indirect, conditional, and floating-condition branches.
	switch w >> 26 {
	case 0x12:
		imm := (w>>10)&0xffff | (w&31)<<16
		delta := loong32rSignExtend(imm, 21) << 2
		p9target, gtarget := loong32rTarget(pc, delta, lookup)
		fcc := w >> 5 & 7
		name, gname := "BFPF", "bceqz"
		if w>>8&1 != 0 {
			name, gname = "BFPT", "bcnez"
		}
		return loong32rText(fmt.Sprintf("%s %s, %s", name, loong32rFCC(fcc), p9target), fmt.Sprintf("%s $fcc%d, %s", gname, fcc, gtarget), gnuAsm), 4
	case 0x13:
		delta := loong32rSignExtend(w>>10&0xffff, 16) << 2
		if delta == 0 && rd == 0 {
			return loong32rText(fmt.Sprintf("JMP (%s)", loong32rR(rj)), fmt.Sprintf("jr %s", loong32rGR(rj)), gnuAsm), 4
		}
		if delta == 0 && rd == 1 {
			return loong32rText(fmt.Sprintf("CALL (%s)", loong32rR(rj)), fmt.Sprintf("jirl $r1, %s, 0", loong32rGR(rj)), gnuAsm), 4
		}
		return loong32rText(fmt.Sprintf("JIRL $%d, %s, %s", delta, loong32rR(rj), loong32rR(rd)), fmt.Sprintf("jirl %s, %s, %d", loong32rGR(rd), loong32rGR(rj), delta), gnuAsm), 4
	case 0x14, 0x15:
		imm := (w>>10)&0xffff | (w&0x3ff)<<16
		delta := loong32rSignExtend(imm, 26) << 2
		p9target, gtarget := loong32rTarget(pc, delta, lookup)
		name, gname := "JMP", "b"
		if w>>26 == 0x15 {
			name, gname = "CALL", "bl"
		}
		return loong32rText(name+" "+p9target, gname+" "+gtarget, gnuAsm), 4
	case 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b:
		delta := loong32rSignExtend(w>>10&0xffff, 16) << 2
		p9target, gtarget := loong32rTarget(pc, delta, lookup)
		names := [...]loong32rOpcode{{"BEQ", "beq"}, {"BNE", "bne"}, {"BLT", "blt"}, {"BGE", "bge"}, {"BLTU", "bltu"}, {"BGEU", "bgeu"}}
		op := names[w>>26-0x16]
		return loong32rText(fmt.Sprintf("%s %s, %s, %s", op.plan9, loong32rR(rj), loong32rR(rd), p9target), fmt.Sprintf("%s %s, %s, %s", op.gnu, loong32rGR(rj), loong32rGR(rd), gtarget), gnuAsm), 4
	}

	// Upper-immediate instructions have a signed 20-bit immediate.
	switch w >> 25 {
	case 0x0a, 0x0e:
		imm := loong32rSignExtend(w>>5&0xfffff, 20)
		op := loong32rOpcode{"LU12IW", "lu12i.w"}
		if w>>25 == 0x0e {
			op = loong32rOpcode{"PCADDU12I", "pcaddu12i"}
		}
		return loong32rText(fmt.Sprintf("%s $%d, %s", op.plan9, imm, loong32rR(rd)), fmt.Sprintf("%s %s, %d", op.gnu, loong32rGR(rd), imm), gnuAsm), 4
	}

	// LL.W and SC.W use a signed si14 field scaled by four bytes.
	switch w >> 24 {
	case 0x20, 0x21:
		off := loong32rSignExtend(w>>10&0x3fff, 14) << 2
		if w>>24 == 0x20 {
			return loong32rText(fmt.Sprintf("LL %d(%s), %s", off, loong32rR(rj), loong32rR(rd)), fmt.Sprintf("ll.w %s, %s, %d", loong32rGR(rd), loong32rGR(rj), off), gnuAsm), 4
		}
		return loong32rText(fmt.Sprintf("SC %s, %d(%s)", loong32rR(rd), off, loong32rR(rj)), fmt.Sprintf("sc.w %s, %s, %d", loong32rGR(rd), loong32rGR(rj), off), gnuAsm), 4
	}

	// Loads, stores, prefetch, and the 12-bit ALU immediate group.
	op10 := w >> 22
	imm12 := loong32rSignExtend(w>>10&0xfff, 12)
	switch op10 {
	case 0x008, 0x009, 0x00a, 0x00d, 0x00e, 0x00f:
		ops := map[uint32]loong32rOpcode{
			0x008: {"SGT", "slti"}, 0x009: {"SGTU", "sltui"},
			0x00a: {"ADD", "addi.w"}, 0x00d: {"AND", "andi"},
			0x00e: {"OR", "ori"}, 0x00f: {"XOR", "xori"},
		}
		op := ops[op10]
		imm := imm12
		if op10 == 0x00d || op10 == 0x00e || op10 == 0x00f {
			imm = int64(w >> 10 & 0xfff)
		}
		return loong32rText(fmt.Sprintf("%s $%d, %s, %s", op.plan9, imm, loong32rR(rj), loong32rR(rd)), fmt.Sprintf("%s %s, %s, %d", op.gnu, loong32rGR(rd), loong32rGR(rj), imm), gnuAsm), 4
	case 0x0a0, 0x0a1, 0x0a2, 0x0a8, 0x0a9, 0x0ac, 0x0ae:
		ops := map[uint32]loong32rOpcode{0x0a0: {"MOVB", "ld.b"}, 0x0a1: {"MOVH", "ld.h"}, 0x0a2: {"MOVW", "ld.w"}, 0x0a8: {"MOVBU", "ld.bu"}, 0x0a9: {"MOVHU", "ld.hu"}, 0x0ac: {"MOVF", "fld.s"}, 0x0ae: {"MOVD", "fld.d"}}
		op := ops[op10]
		to := loong32rR(rd)
		if op10 == 0x0ac || op10 == 0x0ae {
			to = loong32rF(rd)
		}
		gnuTo := loong32rGR(rd)
		if op10 == 0x0ac || op10 == 0x0ae {
			gnuTo = loong32rGF(rd)
		}
		return loong32rText(fmt.Sprintf("%s %d(%s), %s", op.plan9, imm12, loong32rR(rj), to), fmt.Sprintf("%s %s, %s, %d", op.gnu, gnuTo, loong32rGR(rj), imm12), gnuAsm), 4
	case 0x0a4, 0x0a5, 0x0a6, 0x0ad, 0x0af:
		ops := map[uint32]loong32rOpcode{0x0a4: {"MOVB", "st.b"}, 0x0a5: {"MOVH", "st.h"}, 0x0a6: {"MOVW", "st.w"}, 0x0ad: {"MOVF", "fst.s"}, 0x0af: {"MOVD", "fst.d"}}
		op := ops[op10]
		from := loong32rR(rd)
		if op10 == 0x0ad || op10 == 0x0af {
			from = loong32rF(rd)
		}
		gnuFrom := loong32rGR(rd)
		if op10 == 0x0ad || op10 == 0x0af {
			gnuFrom = loong32rGF(rd)
		}
		return loong32rText(fmt.Sprintf("%s %s, %d(%s)", op.plan9, from, imm12, loong32rR(rj)), fmt.Sprintf("%s %s, %s, %d", op.gnu, gnuFrom, loong32rGR(rj), imm12), gnuAsm), 4
	case 0x0ab:
		hint := rd
		return loong32rText(fmt.Sprintf("PRELD %d(%s), $%d", imm12, loong32rR(rj), hint), fmt.Sprintf("preld %d, %s, %d", hint, loong32rGR(rj), imm12), gnuAsm), 4
	}

	// Four-operand floating fused operations.
	if op, ok := loong32rFMA[w>>20]; ok {
		ra := w >> 15 & 31
		return loong32rText(fmt.Sprintf("%s %s, %s, %s, %s", op.plan9, loong32rF(rk), loong32rF(rj), loong32rF(ra), loong32rF(rd)), fmt.Sprintf("%s %s, %s, %s, %s", op.gnu, loong32rGF(rd), loong32rGF(rj), loong32rGF(rk), loong32rGF(ra)), gnuAsm), 4
	}

	if w>>18 == 0x340 {
		fcc, fk := w>>15&7, w>>10&31
		return loong32rText(fmt.Sprintf("FSEL %s, %s, %s, %s", loong32rFCC(fcc), loong32rF(rj), loong32rF(fk), loong32rF(rd)), fmt.Sprintf("fsel %s, %s, %s, $fcc%d", loong32rGF(rd), loong32rGF(rj), loong32rGF(fk), fcc), gnuAsm), 4
	}

	// Three-register integer and floating-point operations.
	if op, ok := loong32rRRR[w>>15]; ok {
		if w>>15 >= 0x200 {
			if op.plan9[:3] == "CMP" {
				fcc := rd & 7
				return loong32rText(fmt.Sprintf("%s %s, %s, %s", op.plan9, loong32rF(rk), loong32rF(rj), loong32rFCC(fcc)), fmt.Sprintf("%s $fcc%d, %s, %s", op.gnu, fcc, loong32rGF(rj), loong32rGF(rk)), gnuAsm), 4
			}
			return loong32rText(fmt.Sprintf("%s %s, %s, %s", op.plan9, loong32rF(rk), loong32rF(rj), loong32rF(rd)), fmt.Sprintf("%s %s, %s, %s", op.gnu, loong32rGF(rd), loong32rGF(rj), loong32rGF(rk)), gnuAsm), 4
		}
		if op.plan9 == "OR" && rk == 0 {
			return loong32rText(fmt.Sprintf("MOVW %s, %s", loong32rR(rj), loong32rR(rd)), fmt.Sprintf("move %s, %s", loong32rGR(rd), loong32rGR(rj)), gnuAsm), 4
		}
		return loong32rText(fmt.Sprintf("%s %s, %s, %s", op.plan9, loong32rR(rk), loong32rR(rj), loong32rR(rd)), fmt.Sprintf("%s %s, %s, %s", op.gnu, loong32rGR(rd), loong32rGR(rj), loong32rGR(rk)), gnuAsm), 4
	}

	// Shift-immediate, barriers, traps, and system calls all have a fixed
	// 17-bit opcode followed by an immediate field.
	switch w >> 15 {
	case 0x081, 0x089, 0x091:
		ops := map[uint32]loong32rOpcode{0x081: {"SLL", "slli.w"}, 0x089: {"SRL", "srli.w"}, 0x091: {"SRA", "srai.w"}}
		op := ops[w>>15]
		return loong32rText(fmt.Sprintf("%s $%d, %s, %s", op.plan9, rk, loong32rR(rj), loong32rR(rd)), fmt.Sprintf("%s %s, %s, %d", op.gnu, loong32rGR(rd), loong32rGR(rj), rk), gnuAsm), 4
	case 0x054:
		imm := w & 0x7fff
		return loong32rText(fmt.Sprintf("BREAK $%d", imm), fmt.Sprintf("break %d", imm), gnuAsm), 4
	case 0x056:
		imm := w & 0x7fff
		return loong32rText(fmt.Sprintf("SYSCALL $%d", imm), fmt.Sprintf("syscall %d", imm), gnuAsm), 4
	case 0x70e4:
		imm := w & 0x7fff
		return loong32rText(fmt.Sprintf("DBAR $%d", imm), fmt.Sprintf("dbar %d", imm), gnuAsm), 4
	case 0x70e5:
		imm := w & 0x7fff
		return loong32rText(fmt.Sprintf("IBAR $%d", imm), fmt.Sprintf("ibar %d", imm), gnuAsm), 4
	}

	// Timer and two-register floating-point instructions.
	if op, ok := loong32rRR[w>>10]; ok {
		switch w >> 10 {
		case 0x0018:
			if rd == 0 {
				return loong32rText(fmt.Sprintf("RDTIMEID %s", loong32rR(rj)), fmt.Sprintf("rdtime.d %s, $r0", loong32rGR(rj)), gnuAsm), 4
			}
			return loong32rText(fmt.Sprintf("RDTIMELW %s, %s", loong32rR(rd), loong32rR(rj)), fmt.Sprintf("rdtimel.w %s, %s", loong32rGR(rd), loong32rGR(rj)), gnuAsm), 4
		case 0x0019:
			return loong32rText(fmt.Sprintf("RDTIMEHW %s, %s", loong32rR(rd), loong32rR(rj)), fmt.Sprintf("rdtimeh.w %s, %s", loong32rGR(rd), loong32rGR(rj)), gnuAsm), 4
		case 0x4529:
			return loong32rText(fmt.Sprintf("MOVW %s, %s", loong32rR(rj), loong32rF(rd)), fmt.Sprintf("%s %s, %s", op.gnu, loong32rGF(rd), loong32rGR(rj)), gnuAsm), 4
		case 0x452d:
			return loong32rText(fmt.Sprintf("MOVW %s, %s", loong32rF(rj), loong32rR(rd)), fmt.Sprintf("%s %s, %s", op.gnu, loong32rGR(rd), loong32rGF(rj)), gnuAsm), 4
		default:
			return loong32rText(fmt.Sprintf("%s %s, %s", op.plan9, loong32rF(rj), loong32rF(rd)), fmt.Sprintf("%s %s, %s", op.gnu, loong32rGF(rd), loong32rGF(rj)), gnuAsm), 4
		}
	}

	return loong32rText(fmt.Sprintf("WORD $0x%08x", w), fmt.Sprintf(".word 0x%08x", w), gnuAsm), 4
}
