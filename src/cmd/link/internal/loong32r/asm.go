// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package loong32r

import (
	"cmd/internal/objabi"
	"cmd/internal/sys"
	"cmd/link/internal/ld"
	"cmd/link/internal/loader"
	"cmd/link/internal/sym"
	"debug/elf"
	"fmt"
	"log"
	"sort"
)

// pcaddLabelName is used for linker-generated local symbols at PCADDU12I
// instructions. Binutils requires each PCADD_LO12 relocation to name the
// exact address of its matching PCADD_HI20 instruction.
const pcaddLabelName = ".Lpcadd_hi "

func gentext(ctxt *ld.Link, ldr *loader.Loader) {
	initfunc, addmoduledata := ld.PrepareAddmoduledata(ctxt)
	if initfunc == nil {
		return
	}

	o := func(op uint32) {
		initfunc.AddUint32(ctxt.Arch, op)
	}

	// Emit the following function:
	//
	//	local.dso_init:
	//		la.pcrel $a0, local.moduledata
	//		b runtime.addmoduledata

	//	00000000 <local.dso_init>:
	//	0:	1c000004	pcaddu12i	$a0, 0
	//				0: R_LARCH_PCADD_HI20	local.moduledata
	o(0x1c000004) // pcaddu12i $a0, 0
	rel, _ := initfunc.AddRel(objabi.R_LOONG32R_ADDR_HI)
	rel.SetOff(0)
	rel.SetSiz(4)
	rel.SetSym(ctxt.Moduledata)

	//	4:	02800084	addi.w	$a0, $a0, 0
	//				4: R_LARCH_PCADD_LO12	local.moduledata
	o(0x02800084) // addi.w $a0, $a0, 0
	rel2, _ := initfunc.AddRel(objabi.R_LOONG32R_ADDR_LO)
	rel2.SetOff(4)
	rel2.SetSiz(4)
	rel2.SetSym(ctxt.Moduledata)

	//	8:	50000000	b	0
	//				8: R_LARCH_B26	runtime.addmoduledata
	o(0x50000000)
	rel3, _ := initfunc.AddRel(objabi.R_CALLLOONG32R)
	rel3.SetOff(8)
	rel3.SetSiz(4)
	rel3.SetSym(addmoduledata)
}

func adddynrel(target *ld.Target, ldr *loader.Loader, syms *ld.ArchSyms, s loader.Sym, r loader.Reloc, rIdx int) bool {
	targ := r.Sym()
	var targType sym.SymKind
	if targ != 0 {
		targType = ldr.SymType(targ)
	}

	switch r.Type() {
	default:
		if r.Type() >= objabi.ElfRelocOffset {
			ldr.Errorf(s, "adddynrel: unexpected reloction type %d (%s)", r.Type(), sym.RelocName(target.Arch, r.Type()))
			return false
		}

	case objabi.ElfRelocOffset + objabi.RelocType(elf.R_LARCH_32):
		if targType == sym.SDYNIMPORT {
			ldr.Errorf(s, "unexpected R_LARCH_32 relocation for dynamic symbol %s", ldr.SymName(targ))
		}
		su := ldr.MakeSymbolUpdater(s)
		su.SetRelocType(rIdx, objabi.R_ADDR)
		if target.IsPIE() && target.IsInternal() {
			// For internal linking PIE, this R_ADDR relocation cannot
			// be resolved statically. We need to generate a dynamic
			// relocation. Let the code below handle it.
			break
		}
		return true

	case objabi.ElfRelocOffset + objabi.RelocType(elf.R_LARCH_B26),
		objabi.ElfRelocOffset + objabi.RelocType(elf.R_LARCH_CALL30):
		if targType == sym.SDYNIMPORT {
			addpltsym(target, ldr, syms, targ)
			su := ldr.MakeSymbolUpdater(s)
			su.SetRelocSym(rIdx, syms.PLT)
			su.SetRelocAdd(rIdx, r.Add()+int64(ldr.SymPlt(targ)))
		}
		if targType == 0 || targType == sym.SXREF {
			ldr.Errorf(s, "unknown symbol %s in callloong32r", ldr.SymName(targ))
		}
		relocType := objabi.R_CALLLOONG32R
		if r.Type() == objabi.ElfRelocOffset+objabi.RelocType(elf.R_LARCH_CALL30) {
			relocType = objabi.R_LOONG32R_CALL30
		}
		su := ldr.MakeSymbolUpdater(s)
		su.SetRelocType(rIdx, relocType)
		return true

	case objabi.ElfRelocOffset + objabi.RelocType(elf.R_LARCH_GOT_PCADD_HI20),
		objabi.ElfRelocOffset + objabi.RelocType(elf.R_LARCH_GOT_PCADD_LO12):
		if targType != sym.SDYNIMPORT {
			// TODO: turn LDR of GOT entry into ADR of symbol itself
		}

		ld.AddGotSym(target, ldr, syms, targ, uint32(elf.R_LARCH_32))
		su := ldr.MakeSymbolUpdater(s)
		if r.Type() == objabi.ElfRelocOffset+objabi.RelocType(elf.R_LARCH_GOT_PCADD_HI20) {
			su.SetRelocType(rIdx, objabi.R_LOONG32R_ADDR_HI)
		} else {
			su.SetRelocType(rIdx, objabi.R_LOONG32R_ADDR_LO)
		}
		su.SetRelocSym(rIdx, syms.GOT)
		su.SetRelocAdd(rIdx, r.Add()+int64(ldr.SymGot(targ)))
		return true

	case objabi.ElfRelocOffset + objabi.RelocType(elf.R_LARCH_PCADD_HI20),
		objabi.ElfRelocOffset + objabi.RelocType(elf.R_LARCH_PCADD_LO12),
		objabi.ElfRelocOffset + objabi.RelocType(elf.R_LARCH_PCREL20_S2):
		if targType == sym.SDYNIMPORT {
			ldr.Errorf(s, "unexpected relocation for dynamic symbol %s", ldr.SymName(targ))
		}
		if targType == 0 || targType == sym.SXREF {
			ldr.Errorf(s, "unknown symbol %s", ldr.SymName(targ))
		}

		var relocType objabi.RelocType
		switch r.Type() - objabi.ElfRelocOffset {
		case objabi.RelocType(elf.R_LARCH_PCADD_HI20):
			relocType = objabi.R_LOONG32R_ADDR_HI
		case objabi.RelocType(elf.R_LARCH_PCADD_LO12):
			relocType = objabi.R_LOONG32R_ADDR_LO
		case objabi.RelocType(elf.R_LARCH_PCREL20_S2):
			relocType = objabi.R_LOONG32R_ADDR_PCREL20_S2
		}
		su := ldr.MakeSymbolUpdater(s)
		su.SetRelocType(rIdx, relocType)
		return true

	case objabi.ElfRelocOffset + objabi.RelocType(elf.R_LARCH_ADD32),
		objabi.ElfRelocOffset + objabi.RelocType(elf.R_LARCH_SUB32):
		su := ldr.MakeSymbolUpdater(s)
		if r.Type() == objabi.ElfRelocOffset+objabi.RelocType(elf.R_LARCH_ADD32) {
			su.SetRelocType(rIdx, objabi.R_LOONG32R_ADD32)
		} else {
			su.SetRelocType(rIdx, objabi.R_LOONG32R_SUB32)
		}
		return true

	case objabi.ElfRelocOffset + objabi.RelocType(elf.R_LARCH_B16),
		objabi.ElfRelocOffset + objabi.RelocType(elf.R_LARCH_B21):
		if targType == sym.SDYNIMPORT {
			addpltsym(target, ldr, syms, targ)
			su := ldr.MakeSymbolUpdater(s)
			su.SetRelocSym(rIdx, syms.PLT)
			su.SetRelocAdd(rIdx, r.Add()+int64(ldr.SymPlt(targ)))
		}
		if targType == 0 || targType == sym.SXREF {
			ldr.Errorf(s, "unknown symbol %s in R_JMPxxLOONG32R", ldr.SymName(targ))
		}
		su := ldr.MakeSymbolUpdater(s)
		if r.Type() == objabi.ElfRelocOffset+objabi.RelocType(elf.R_LARCH_B16) {
			su.SetRelocType(rIdx, objabi.R_JMP16LOONG32R)
		} else {
			su.SetRelocType(rIdx, objabi.R_JMP21LOONG32R)
		}
		return true
	}

	relocs := ldr.Relocs(s)
	r = relocs.At(rIdx)

	switch r.Type() {
	case objabi.R_CALLLOONG32R:
		if targType != sym.SDYNIMPORT {
			return true
		}
		if target.IsExternal() {
			return true
		}

		// Internal linking.
		if r.Add() != 0 {
			ldr.Errorf(s, "PLT call with no-zero addend (%v)", r.Add())
		}

		// Build a PLT entry and change the relocation target to that entry.
		addpltsym(target, ldr, syms, targ)
		su := ldr.MakeSymbolUpdater(s)
		su.SetRelocSym(rIdx, syms.PLT)
		su.SetRelocAdd(rIdx, int64(ldr.SymPlt(targ)))
		return true

	case objabi.R_ADDR:
		if ldr.SymType(s) == sym.STEXT && target.IsElf() {
			// The code is asking for the address of an external
			// function. We provide it with the address of the
			// correspondent GOT symbol.
			ld.AddGotSym(target, ldr, syms, targ, uint32(elf.R_LARCH_32))
			su := ldr.MakeSymbolUpdater(s)
			su.SetRelocSym(rIdx, syms.GOT)
			su.SetRelocAdd(rIdx, r.Add()+int64(ldr.SymGot(targ)))
			return true
		}

		// Process dynamic relocations for the data sections.
		if target.IsPIE() && target.IsInternal() {
			// When internally linking, generate dynamic relocations
			// for all typical R_ADDR relocations. The exception
			// are those R_ADDR that are created as part of generating
			// the dynamic relocations and must be resolved statically.
			//
			// There are three phases relevant to understanding this:
			//
			//	dodata()  // we are here
			//	address() // symbol address assignment
			//	reloc()   // resolution of static R_ADDR relocs
			//
			// At this point symbol addresses have not been
			// assigned yet (as the final size of the .rela section
			// will affect the addresses), and so we cannot write
			// the Elf32_Rela.r_offset now. Instead we delay it
			// until after the 'address' phase of the linker is
			// complete. We do this via Addaddrplus, which creates
			// a new R_ADDR relocation which will be resolved in
			// the 'reloc' phase.
			//
			// These synthetic static R_ADDR relocs must be skipped
			// now, or else we will be caught in an infinite loop
			// of generating synthetic relocs for our synthetic
			// relocs.
			//
			// Furthermore, the rela sections contain dynamic
			// relocations with R_ADDR relocations on
			// Elf32_Rela.r_offset. This field should contain the
			// symbol offset as determined by reloc(), not the
			// final dynamically linked address as a dynamic
			// relocation would provide.
			switch ldr.SymName(s) {
			case ".dynsym", ".rela", ".rela.plt", ".got.plt", ".dynamic":
				return false
			}
		} else {
			// Either internally linking a static executable,
			// in which case we can resolve these relocations
			// statically in the 'reloc' phase, or externally
			// linking, in which case the relocation will be
			// prepared in the 'reloc' phase and passed to the
			// external linker in the 'asmb' phase.
			if ldr.SymType(s) != sym.SDATA && ldr.SymType(s) != sym.SRODATA {
				break
			}
		}

		if target.IsElf() {
			// Generate R_LARCH_RELATIVE relocations for best
			// efficiency in the dynamic linker.
			//
			// As noted above, symbol addresses have not been
			// assigned yet, so we can't generate the final reloc
			// entry yet. We ultimately want:
			//
			// r_offset = s + r.Off
			// r_info = R_LARCH_RELATIVE
			// r_addend = targ + r.Add
			//
			// The dynamic linker will set *offset = base address +
			// addend.
			//
			// AddAddrPlus is used for r_offset and r_addend to
			// generate new R_ADDR relocations that will update
			// these fields in the 'reloc' phase.
			rela := ldr.MakeSymbolUpdater(syms.Rela)
			rela.AddAddrPlus(target.Arch, s, int64(r.Off()))
			if r.Siz() == 4 {
				rela.AddUint32(target.Arch, elf.R_INFO32(0, uint32(elf.R_LARCH_RELATIVE)))
			} else {
				ldr.Errorf(s, "unexpected relocation for dynamic symbol %s", ldr.SymName(targ))
			}
			rela.AddAddrPlus(target.Arch, targ, r.Add())
			return true
		}

	case objabi.R_LOONG32R_GOT_HI,
		objabi.R_LOONG32R_GOT_LO:
		ld.AddGotSym(target, ldr, syms, targ, uint32(elf.R_LARCH_32))
		su := ldr.MakeSymbolUpdater(s)
		if r.Type() == objabi.R_LOONG32R_GOT_HI {
			su.SetRelocType(rIdx, objabi.R_LOONG32R_ADDR_HI)
		} else {
			su.SetRelocType(rIdx, objabi.R_LOONG32R_ADDR_LO)
		}
		su.SetRelocSym(rIdx, syms.GOT)
		su.SetRelocAdd(rIdx, r.Add()+int64(ldr.SymGot(targ)))
		return true
	}
	return false
}

func elfsetupplt(ctxt *ld.Link, ldr *loader.Loader, plt, gotplt *loader.SymbolBuilder, dynamic loader.Sym) {
	if plt.Size() == 0 {
		// pcaddu12i $r14, %pcadd_hi20(.got.plt)
		plt.AddSymRef(ctxt.Arch, gotplt.Sym(), 0, objabi.R_LOONG32R_ADDR_HI, 4)
		plt.SetUint32(ctxt.Arch, plt.Size()-4, 0x1c00000e)

		// sub.w $r13, $r13, $r15
		plt.AddUint32(ctxt.Arch, 0x00113dad)

		// ld.w $r15, $r14, %pcadd_lo12(.got.plt)
		// This low relocation is paired with the PCADDU12I 8 bytes
		// earlier. Its addend compensates for the linker's adjacent-pair
		// convention used by ordinary address sequences.
		plt.AddSymRef(ctxt.Arch, gotplt.Sym(), 4, objabi.R_LOONG32R_ADDR_LO, 4)
		plt.SetUint32(ctxt.Arch, plt.Size()-4, 0x288001cf)

		// addi.w $r13, $r13, -(PLT_HEADER_SIZE+12)
		plt.AddUint32(ctxt.Arch, 0x02bf51ad)

		// addi.w $r12, $r14, %pcadd_lo12(.got.plt)
		plt.AddSymRef(ctxt.Arch, gotplt.Sym(), 12, objabi.R_LOONG32R_ADDR_LO, 4)
		plt.SetUint32(ctxt.Arch, plt.Size()-4, 0x028001cc)

		// srli.w $r13, $r13, 2
		plt.AddUint32(ctxt.Arch, 0x004489ad)

		// ld.w $r12, $r12, 4
		plt.AddUint32(ctxt.Arch, 0x2880118c)

		// jirl $r0, $r15, 0
		plt.AddUint32(ctxt.Arch, 0x4c0001e0)

		// check gotplt.size == 0
		if gotplt.Size() != 0 {
			ctxt.Errorf(gotplt.Sym(), "got.plt is not empty at the very beginning")
		}

		gotplt.AddUint32(ctxt.Arch, 0)
		gotplt.AddUint32(ctxt.Arch, 0)
	}
}

func addpltsym(target *ld.Target, ldr *loader.Loader, syms *ld.ArchSyms, s loader.Sym) {
	if ldr.SymPlt(s) >= 0 {
		return
	}

	ld.Adddynsym(ldr, target, syms, s)

	if target.IsElf() {
		plt := ldr.MakeSymbolUpdater(syms.PLT)
		gotplt := ldr.MakeSymbolUpdater(syms.GOTPLT)
		rela := ldr.MakeSymbolUpdater(syms.RelaPLT)
		if plt.Size() == 0 {
			panic("plt is not set up")
		}

		// pcaddu12i $r15, %pcadd_hi20(.got.plt entry)
		plt.AddAddrPlus4(target.Arch, gotplt.Sym(), gotplt.Size())
		plt.SetUint32(target.Arch, plt.Size()-4, 0x1c00000f)
		relocs := plt.Relocs()
		plt.SetRelocType(relocs.Count()-1, objabi.R_LOONG32R_ADDR_HI)

		// ld.w $r15, $r15, %pcadd_lo12(.got.plt entry)
		plt.AddAddrPlus4(target.Arch, gotplt.Sym(), gotplt.Size())
		plt.SetUint32(target.Arch, plt.Size()-4, 0x288001ef)
		relocs = plt.Relocs()
		plt.SetRelocType(relocs.Count()-1, objabi.R_LOONG32R_ADDR_LO)

		// jirl $r13, $r15, 0; nop
		plt.AddUint32(target.Arch, 0x4c0001ed)
		plt.AddUint32(target.Arch, 0x03400000)

		// add to got.plt: pointer to plt[0]
		gotplt.AddAddrPlus(target.Arch, plt.Sym(), 0)

		// rela
		rela.AddAddrPlus(target.Arch, gotplt.Sym(), gotplt.Size()-4)
		sDynid := ldr.SymDynid(s)
		rela.AddUint32(target.Arch, elf.R_INFO32(uint32(sDynid), uint32(elf.R_LARCH_JUMP_SLOT)))
		rela.AddUint32(target.Arch, 0)

		ldr.SetPlt(s, int32(plt.Size()-16))
	} else {
		ldr.Errorf(s, "addpltsym: unsupport binary format")
	}
}

func genSymsLate(ctxt *ld.Link, ldr *loader.Loader) {
	if ctxt.LinkMode != ld.LinkExternal {
		return
	}
	if ctxt.Textp == nil {
		log.Fatal("loong32r genSymsLate called before Textp has been assigned")
	}

	var pcaddSyms []loader.Sym
	for _, s := range ctxt.Textp {
		relocs := ldr.Relocs(s)
		for ri := 0; ri < relocs.Count(); ri++ {
			r := relocs.At(ri)
			switch r.Type() {
			case objabi.R_LOONG32R_ADDR_LO,
				objabi.R_LOONG32R_GOT_LO,
				objabi.R_LOONG32R_TLS_IE_LO:
			default:
				continue
			}

			hiOff := int64(r.Off()) - 4
			if hiOff < 0 {
				ldr.Errorf(s, "PCADD_LO12 relocation has no preceding PCADD_HI20")
				continue
			}
			if hiOff == 0 && ldr.SymType(s).IsText() {
				continue
			}

			sb := ldr.MakeSymbolBuilder(pcaddLabelName)
			sb.SetType(sym.STEXT)
			sb.SetValue(ldr.SymValue(s) + hiOff)
			sb.SetLocal(true)
			sb.SetReachable(true)
			sb.SetVisibilityHidden(true)
			sb.SetSect(ldr.SymSect(s))
			if outer := ldr.OuterSym(s); outer != 0 {
				ldr.AddInteriorSym(outer, sb.Sym())
			}
			pcaddSyms = append(pcaddSyms, sb.Sym())
		}
	}
	ctxt.Textp = append(ctxt.Textp, pcaddSyms...)
	ldr.SortSyms(ctxt.Textp)
}

func findPCAddHI20Symbol(ctxt *ld.Link, ldr *loader.Loader, value int64) loader.Sym {
	i := sort.Search(len(ctxt.Textp), func(i int) bool {
		return ldr.SymValue(ctxt.Textp[i]) >= value
	})
	if i < len(ctxt.Textp) {
		s := ctxt.Textp[i]
		if ldr.SymValue(s) == value && ldr.SymType(s).IsText() {
			return s
		}
	}
	return 0
}

func elfreloc1(ctxt *ld.Link, out *ld.OutBuf, ldr *loader.Loader, s loader.Sym, r loader.ExtReloc, ri int, sectoff int64) bool {
	elfsym := ld.ElfSymForReloc(ctxt, r.Xsym)
	writeRelaSym := func(typ elf.R_LARCH, symIndex int32, addend int64) {
		out.Write32(uint32(sectoff))
		out.Write32(elf.R_INFO32(uint32(symIndex), uint32(typ)))
		out.Write32(uint32(addend))
	}
	writeRela := func(typ elf.R_LARCH, addend int64) {
		writeRelaSym(typ, elfsym, addend)
	}
	writePCAddLo := func(typ elf.R_LARCH) {
		// PCADD_LO12 does not name the final target. It names the matching
		// PCADD_HI20 instruction, whose relocation carries the target and
		// addend. Every dedicated LA32R address sequence places that high
		// instruction immediately before the low instruction.
		relocs := ldr.Relocs(s)
		hiValue := ldr.SymValue(s) + int64(relocs.At(ri).Off()) - 4
		hiSym := findPCAddHI20Symbol(ctxt, ldr, hiValue)
		if hiSym == 0 {
			ld.Errorf("failed to find LA32R PCADD_HI20 symbol at %d (%x)", sectoff, hiValue)
			writeRelaSym(typ, 0, 0)
			return
		}
		writeRelaSym(typ, ld.ElfSymForReloc(ctxt, hiSym), 0)
	}
	switch r.Type {
	default:
		return false
	case objabi.R_ADDR, objabi.R_DWARFSECREF:
		if r.Size != 4 {
			return false
		}
		writeRela(elf.R_LARCH_32, r.Xadd)
	case objabi.R_LOONG32R_TLS_LE_LO:
		writeRela(elf.R_LARCH_TLS_LE_LO12, r.Xadd)

	case objabi.R_LOONG32R_TLS_LE_HI:
		writeRela(elf.R_LARCH_TLS_LE_HI20, r.Xadd)

	case objabi.R_CALLLOONG32R:
		writeRela(elf.R_LARCH_B26, r.Xadd)

	case objabi.R_LOONG32R_CALL30:
		writeRela(elf.R_LARCH_CALL30, r.Xadd)

	case objabi.R_LOONG32R_TLS_IE_HI:
		writeRela(elf.R_LARCH_TLS_IE_PCADD_HI20, 0)

	case objabi.R_LOONG32R_TLS_IE_LO:
		writePCAddLo(elf.R_LARCH_TLS_IE_PCADD_LO12)

	case objabi.R_LOONG32R_ADDR_LO:
		writePCAddLo(elf.R_LARCH_PCADD_LO12)

	case objabi.R_LOONG32R_ADDR_HI:
		writeRela(elf.R_LARCH_PCADD_HI20, r.Xadd)

	case objabi.R_LOONG32R_ADDR_PCREL20_S2:
		writeRela(elf.R_LARCH_PCREL20_S2, r.Xadd)

	case objabi.R_LOONG32R_GOT_HI:
		writeRela(elf.R_LARCH_GOT_PCADD_HI20, 0)

	case objabi.R_LOONG32R_GOT_LO:
		writePCAddLo(elf.R_LARCH_GOT_PCADD_LO12)

	case objabi.R_JMP16LOONG32R:
		writeRela(elf.R_LARCH_B16, r.Xadd)
	case objabi.R_JMP21LOONG32R:
		writeRela(elf.R_LARCH_B21, r.Xadd)
	case objabi.R_JMPLOONG32R:
		writeRela(elf.R_LARCH_B26, r.Xadd)
	case objabi.R_LOONG32R_ADD32:
		writeRela(elf.R_LARCH_ADD32, r.Xadd)
	case objabi.R_LOONG32R_SUB32:
		writeRela(elf.R_LARCH_SUB32, r.Xadd)
	}

	return true
}

func machoreloc1(*sys.Arch, *ld.OutBuf, *loader.Loader, loader.Sym, loader.ExtReloc, int64) bool {
	return false
}

func archreloc(target *ld.Target, ldr *loader.Loader, syms *ld.ArchSyms, r loader.Reloc, s loader.Sym, val int64) (o int64, nExtReloc int, ok bool) {
	rs := r.Sym()
	if target.IsExternal() {
		switch r.Type() {
		default:
			return val, 0, false
		case objabi.R_LOONG32R_ADDR_HI,
			objabi.R_LOONG32R_ADDR_LO,
			objabi.R_LOONG32R_ADDR_PCREL20_S2:
			// set up addend for eventual relocation via outer symbol.
			rs, _ := ld.FoldSubSymbolOffset(ldr, rs)
			rst := ldr.SymType(rs)
			if rst != sym.SHOSTOBJ && rst != sym.SDYNIMPORT && ldr.SymSect(rs) == nil {
				ldr.Errorf(s, "missing section for %s", ldr.SymName(rs))
			}
			return val, 1, true
		case objabi.R_LOONG32R_TLS_LE_HI,
			objabi.R_LOONG32R_TLS_LE_LO,
			objabi.R_CALLLOONG32R,
			objabi.R_LOONG32R_CALL30,
			objabi.R_JMPLOONG32R,
			objabi.R_LOONG32R_TLS_IE_HI,
			objabi.R_LOONG32R_TLS_IE_LO,
			objabi.R_LOONG32R_GOT_HI,
			objabi.R_LOONG32R_GOT_LO,
			objabi.R_JMP16LOONG32R,
			objabi.R_JMP21LOONG32R,
			objabi.R_LOONG32R_ADD32,
			objabi.R_LOONG32R_SUB32:
			return val, 1, true
		}
	}

	const isOk = true
	const noExtReloc = 0

	switch r.Type() {
	case objabi.R_CONST:
		return r.Add(), noExtReloc, isOk
	case objabi.R_GOTOFF:
		return ldr.SymValue(r.Sym()) + r.Add() - ldr.SymValue(syms.GOT), noExtReloc, isOk
	case objabi.R_LOONG32R_ADDR_HI,
		objabi.R_LOONG32R_ADDR_LO:
		pc := ldr.SymValue(s) + int64(r.Off())
		if r.Type() == objabi.R_LOONG32R_ADDR_LO {
			pc -= 4 // address of the paired PCADDU12I
		}
		t := ldr.SymAddr(rs) + r.Add() - pc
		if r.Type() == objabi.R_LOONG32R_ADDR_LO {
			return val&0xffc003ff | ((t & 0xfff) << 10), noExtReloc, isOk
		}
		return val&0xfe00001f | ((((t + 0x800) >> 12) & 0xfffff) << 5), noExtReloc, isOk
	case objabi.R_LOONG32R_ADDR_PCREL20_S2:
		pc := ldr.SymValue(s) + int64(r.Off())
		t := (ldr.SymAddr(rs) + r.Add() - pc) >> 2
		return val&0xfe00001f | ((t & 0xfffff) << 5), noExtReloc, isOk
	case objabi.R_LOONG32R_TLS_LE_HI,
		objabi.R_LOONG32R_TLS_LE_LO:
		t := ldr.SymAddr(rs) + r.Add()
		if r.Type() == objabi.R_LOONG32R_TLS_LE_LO {
			return val&0xffc003ff | ((t & 0xfff) << 10), noExtReloc, isOk
		}
		return val&0xfe00001f | (((t) >> 12 << 5) & 0x1ffffe0), noExtReloc, isOk
	case objabi.R_CALLLOONG32R,
		objabi.R_JMPLOONG32R:
		pc := ldr.SymValue(s) + int64(r.Off())
		t := ldr.SymAddr(rs) + r.Add() - pc
		return val&0xfc000000 | (((t >> 2) & 0xffff) << 10) | (((t >> 2) & 0x3ff0000) >> 16), noExtReloc, isOk

	case objabi.R_LOONG32R_CALL30:
		pc := ldr.SymValue(s) + int64(r.Off())
		t := ldr.SymAddr(rs) + r.Add() - pc
		// val is PCADDU12I (low word) followed by JIRL (high word).
		pcaddu12i := (val & 0xfe00001f) | (((t >> 12) & 0xfffff) << 5)
		jirl := ((val >> 32) & 0xfc0003ff) | (((t >> 2) & 0x3ff) << 10)
		return pcaddu12i | (jirl << 32), noExtReloc, isOk

	case objabi.R_JMP16LOONG32R,
		objabi.R_JMP21LOONG32R:
		pc := ldr.SymValue(s) + int64(r.Off())
		t := ldr.SymAddr(rs) + r.Add() - pc
		if r.Type() == objabi.R_JMP16LOONG32R {
			return val&0xfc0003ff | (((t >> 2) & 0xffff) << 10), noExtReloc, isOk
		}
		return val&0xfc0003e0 | (((t >> 2) & 0xffff) << 10) | (((t >> 2) & 0x1f0000) >> 16), noExtReloc, isOk

	case objabi.R_LOONG32R_TLS_IE_HI,
		objabi.R_LOONG32R_TLS_IE_LO:
		if target.IsPIE() && target.IsElf() {
			if !target.IsLinux() {
				ldr.Errorf(s, "TLS reloc on unsupported OS %v", target.HeadType)
			}
			t := ldr.SymAddr(rs) + r.Add()
			if r.Type() == objabi.R_LOONG32R_TLS_IE_HI {
				// pcaddu12i -> lu12i.w
				return (0x14000000 | (val & 0x1f) | ((t >> 12) << 5)), noExtReloc, isOk
			}
			// ld.w -> ori
			return (0x03800000 | (val & 0x3ff) | ((t & 0xfff) << 10)), noExtReloc, isOk
		} else {
			log.Fatalf("cannot handle R_LOONG32R_TLS_IE_x (sym %s) when linking internally", ldr.SymName(rs))
		}

	case objabi.R_LOONG32R_ADD32, objabi.R_LOONG32R_SUB32:
		if r.Type() == objabi.R_LOONG32R_ADD32 {
			return val + ldr.SymAddr(rs) + r.Add(), noExtReloc, isOk
		}
		return val - (ldr.SymAddr(rs) + r.Add()), noExtReloc, isOk
	}

	return val, 0, false
}

func archrelocvariant(*ld.Target, *loader.Loader, loader.Reloc, sym.RelocVariant, loader.Sym, int64, []byte) int64 {
	return -1
}

func extreloc(target *ld.Target, ldr *loader.Loader, r loader.Reloc, s loader.Sym) (loader.ExtReloc, bool) {
	switch r.Type() {
	case objabi.R_LOONG32R_ADDR_HI,
		objabi.R_LOONG32R_ADDR_LO,
		objabi.R_LOONG32R_ADDR_PCREL20_S2,
		objabi.R_LOONG32R_GOT_HI,
		objabi.R_LOONG32R_GOT_LO:
		return ld.ExtrelocViaOuterSym(ldr, r, s), true

	case objabi.R_LOONG32R_TLS_LE_HI,
		objabi.R_LOONG32R_TLS_LE_LO,
		objabi.R_CONST,
		objabi.R_GOTOFF,
		objabi.R_CALLLOONG32R,
		objabi.R_LOONG32R_CALL30,
		objabi.R_JMP16LOONG32R,
		objabi.R_JMP21LOONG32R,
		objabi.R_JMPLOONG32R,
		objabi.R_LOONG32R_TLS_IE_HI,
		objabi.R_LOONG32R_TLS_IE_LO,
		objabi.R_LOONG32R_ADD32,
		objabi.R_LOONG32R_SUB32:
		return ld.ExtrelocSimple(ldr, r), true
	}
	return loader.ExtReloc{}, false
}

// Convert the direct jump relocation r to refer to a trampoline if the target is too far.
func trampoline(ctxt *ld.Link, ldr *loader.Loader, ri int, rs, s loader.Sym) {
	relocs := ldr.Relocs(s)
	r := relocs.At(ri)
	switch r.Type() {
	case objabi.ElfRelocOffset + objabi.RelocType(elf.R_LARCH_B26), objabi.R_CALLLOONG32R:
		if ldr.SymType(rs) == sym.SDYNIMPORT {
			// Nothing to do.
			// The plt symbol has not been added. If we add tramp
			// here, plt will not work.
			return
		}

		var t int64
		// ldr.SymValue(rs) == 0 indicates a cross-package jump to a function that is not yet
		// laid out. Conservatively use a trampoline. This should be rare, as we lay out packages
		// in dependency order.
		if ldr.SymValue(rs) != 0 {
			t = ldr.SymValue(rs) + r.Add() - (ldr.SymValue(s) + int64(r.Off()))
		}
		if t >= 1<<27 || t < -1<<27 || ldr.SymValue(rs) == 0 || (*ld.FlagDebugTramp > 1 && (ldr.SymPkg(s) == "" || ldr.SymPkg(s) != ldr.SymPkg(rs))) {
			// direct call too far need to insert trampoline.
			// look up existing trampolines first. if we found one within the range
			// of direct call, we can reuse it. otherwise create a new one.
			var tramp loader.Sym
			for i := 0; ; i++ {
				oName := ldr.SymName(rs)
				name := oName + fmt.Sprintf("%+x-tramp%d", r.Add(), i)
				tramp = ldr.LookupOrCreateSym(name, ldr.SymVersion(rs))
				ldr.SetAttrReachable(tramp, true)
				if ldr.SymType(tramp) == sym.SDYNIMPORT {
					// don't reuse trampoline defined in other module
					continue
				}
				if oName == "runtime.deferreturn" {
					ldr.SetIsDeferReturnTramp(tramp, true)
				}
				if ldr.SymValue(tramp) == 0 {
					// either the trampoline does not exist -- we need to create one,
					// or found one the address which is not assigned -- this will be
					// laid down immediately after the current function. use this one.
					break
				}

				t = ldr.SymValue(tramp) - (ldr.SymValue(s) + int64(r.Off()))
				if t >= -1<<27 && t < 1<<27 {
					// found an existing trampoline that is not too far
					// we can just use it.
					break
				}
			}
			if ldr.SymType(tramp) == 0 {
				// trampoline does not exist, create one
				trampb := ldr.MakeSymbolUpdater(tramp)
				ctxt.AddTramp(trampb, ldr.SymType(s))
				if ldr.SymType(rs) == sym.SDYNIMPORT {
					if r.Add() != 0 {
						ctxt.Errorf(s, "nonzero addend for DYNIMPORT call: %v+%d", ldr.SymName(rs), r.Add())
					}
					gentrampgot(ctxt, ldr, trampb, rs)
				} else {
					gentramp(ctxt, ldr, trampb, rs, r.Add())
				}
			}
			// modify reloc to point to tramp, which will be resolved later
			sb := ldr.MakeSymbolUpdater(s)
			relocs := sb.Relocs()
			r := relocs.At(ri)
			r.SetSym(tramp)
			r.SetAdd(0) // clear the offset embedded in the instruction
		}
	default:
		ctxt.Errorf(s, "trampoline called with non-jump reloc: %d (%s)", r.Type(), sym.RelocName(ctxt.Arch, r.Type()))
	}
}

// generate a trampoline to target+offset.
func gentramp(ctxt *ld.Link, ldr *loader.Loader, tramp *loader.SymbolBuilder, target loader.Sym, offset int64) {
	tramp.SetSize(12) // 3 instructions
	P := make([]byte, tramp.Size())

	o1 := uint32(0x1c00001e) // pcaddu12i $r30, 0
	ctxt.Arch.ByteOrder.PutUint32(P, o1)
	r1, _ := tramp.AddRel(objabi.R_LOONG32R_ADDR_HI)
	r1.SetOff(0)
	r1.SetSiz(4)
	r1.SetSym(target)
	r1.SetAdd(offset)

	o2 := uint32(0x028003de) // addi.w $r30, $r30, 0
	ctxt.Arch.ByteOrder.PutUint32(P[4:], o2)
	r2, _ := tramp.AddRel(objabi.R_LOONG32R_ADDR_LO)
	r2.SetOff(4)
	r2.SetSiz(4)
	r2.SetSym(target)
	r2.SetAdd(offset)

	o3 := uint32(0x4c0003c0) // jirl $r0, $r30, 0
	ctxt.Arch.ByteOrder.PutUint32(P[8:], o3)

	tramp.SetData(P)
}

func gentrampgot(ctxt *ld.Link, ldr *loader.Loader, tramp *loader.SymbolBuilder, target loader.Sym) {
	tramp.SetSize(12) // 3 instructions
	P := make([]byte, tramp.Size())

	o1 := uint32(0x1c00001e) // pcaddu12i $r30, 0
	ctxt.Arch.ByteOrder.PutUint32(P, o1)
	r1, _ := tramp.AddRel(objabi.R_LOONG32R_GOT_HI)
	r1.SetOff(0)
	r1.SetSiz(4)
	r1.SetSym(target)

	o2 := uint32(0x288003de) // ld.w $r30, $r30, 0
	ctxt.Arch.ByteOrder.PutUint32(P[4:], o2)
	r2, _ := tramp.AddRel(objabi.R_LOONG32R_GOT_LO)
	r2.SetOff(4)
	r2.SetSiz(4)
	r2.SetSym(target)

	o3 := uint32(0x4c0003c0) // jirl $r0, $r30, 0
	ctxt.Arch.ByteOrder.PutUint32(P[8:], o3)

	tramp.SetData(P)
}
