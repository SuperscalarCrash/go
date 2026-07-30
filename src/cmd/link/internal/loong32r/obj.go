// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package loong32r

import (
	"cmd/internal/objabi"
	"cmd/internal/sys"
	"cmd/link/internal/ld"
)

func Init() (*sys.Arch, ld.Arch) {
	arch := sys.ArchLoong32r

	theArch := ld.Arch{
		Funcalign:        funcAlign,
		Maxalign:         maxAlign,
		Minalign:         minAlign,
		Dwarfregsp:       dwarfRegSP,
		Dwarfreglr:       dwarfRegLR,
		TrampLimit:       0x7c00000,                      // 26-bit signed offset * 4, leave room for PLT etc.
		CodePad:          []byte{0x00, 0x00, 0x2a, 0x00}, // BREAK 0
		Adddynrel:        adddynrel,
		Archinit:         archinit,
		Archreloc:        archreloc,
		Archrelocvariant: archrelocvariant,
		Extreloc:         extreloc,
		Machoreloc1:      machoreloc1,
		Gentext:          gentext,
		GenSymsLate:      genSymsLate,
		Trampoline:       trampoline,

		ELF: ld.ELFArch{
			Linuxdynld:     "/lib32/ld-linux-loongarch-ilp32s.so.1",
			LinuxdynldMusl: "/lib/ld-musl-loongarch32-sf.so.1",
			Freebsddynld:   "XXX",
			Openbsddynld:   "XXX",
			Netbsddynld:    "XXX",
			Dragonflydynld: "XXX",
			Solarisdynld:   "XXX",

			Reloc1:    elfreloc1,
			RelocSize: 12,
			SetupPLT:  elfsetupplt,
		},
	}

	return arch, theArch
}

func archinit(ctxt *ld.Link) {
	switch ctxt.HeadType {
	default:
		ld.Exitf("unknown -H option: %v", ctxt.HeadType)
	case objabi.Hlinux: /* loong32r ELF32 */
		ld.Elfinit(ctxt)
		ld.HEADR = ld.ELFRESERVE
		if *ld.FlagRound == -1 {
			*ld.FlagRound = 0x10000
		}
		if *ld.FlagTextAddr == -1 {
			*ld.FlagTextAddr = ld.Rnd(0x10000, *ld.FlagRound) + int64(ld.HEADR)
		}
	}
}
