// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package loong32r

import (
	"cmd/internal/objabi"
	"cmd/link/internal/loader"
	"debug/elf"
	"testing"
)

func TestPairedPCAddReloc(t *testing.T) {
	elfHi := objabi.ElfRelocOffset + objabi.RelocType(elf.R_LARCH_PCADD_HI20)
	goHi := objabi.R_LOONG32R_ADDR_HI

	tests := []struct {
		name       string
		highType   objabi.RelocType
		interfere  bool
		highOffset int32
		want       bool
	}{
		{name: "ELF high", highType: elfHi, highOffset: 12, want: true},
		{name: "rewritten Go high", highType: goHi, highOffset: 12, want: true},
		{name: "same-offset unrelated relocation", highType: elfHi, interfere: true, highOffset: 12, want: true},
		{name: "high is not immediately before low", highType: elfHi, highOffset: 8, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reporter := &loader.ErrorReporter{}
			ldr := loader.NewLoader(0, reporter)
			text := ldr.LookupOrCreateSym("text", 0)
			target := ldr.LookupOrCreateSym("target", 0)
			localLabel := ldr.CreateStaticSym(".Lpcadd_hi")
			sb := ldr.MakeSymbolUpdater(text)

			hi, _ := sb.AddRel(test.highType)
			hi.SetOff(test.highOffset)
			hi.SetSiz(4)
			hi.SetSym(target)
			hi.SetAdd(37)

			if test.interfere {
				r, _ := sb.AddRel(objabi.R_ADDR)
				r.SetOff(12)
				r.SetSiz(4)
				r.SetSym(localLabel)
			}

			low, lowIdx := sb.AddRel(objabi.ElfRelocOffset + objabi.RelocType(elf.R_LARCH_PCADD_LO12))
			low.SetOff(16)
			low.SetSiz(4)
			low.SetSym(localLabel)

			got, ok := pairedPCAddReloc(ldr, text, lowIdx, elfHi, goHi)
			if ok != test.want {
				t.Fatalf("pairedPCAddReloc ok = %v, want %v", ok, test.want)
			}
			if !ok {
				return
			}
			if got.Type() != test.highType || got.Off() != test.highOffset || got.Sym() != target || got.Add() != 37 {
				t.Fatalf("paired high = {type:%v off:%d sym:%d add:%d}, want {type:%v off:%d sym:%d add:37}",
					got.Type(), got.Off(), got.Sym(), got.Add(), test.highType, test.highOffset, target)
			}
		})
	}
}

func TestPCAddHighReloc(t *testing.T) {
	tests := []struct {
		name       string
		lowType    objabi.RelocType
		highType   objabi.RelocType
		highOffset int32
		highAdd    int64
		lowOffset  int32
		lowAdd     int64
		want       bool
	}{
		{
			name:       "adjacent address pair",
			lowType:    objabi.R_LOONG32R_ADDR_LO,
			highType:   objabi.R_LOONG32R_ADDR_HI,
			highOffset: 12,
			highAdd:    37,
			lowOffset:  16,
			lowAdd:     37,
			want:       true,
		},
		{
			name:       "nonadjacent PLT pair",
			lowType:    objabi.R_LOONG32R_ADDR_LO,
			highType:   objabi.R_LOONG32R_ADDR_HI,
			highOffset: 0,
			lowOffset:  16,
			lowAdd:     12,
			want:       true,
		},
		{
			name:       "adjacent GOT pair",
			lowType:    objabi.R_LOONG32R_GOT_LO,
			highType:   objabi.R_LOONG32R_GOT_HI,
			highOffset: 20,
			lowOffset:  24,
			want:       true,
		},
		{
			name:       "adjacent TLS IE pair",
			lowType:    objabi.R_LOONG32R_TLS_IE_LO,
			highType:   objabi.R_LOONG32R_TLS_IE_HI,
			highOffset: 28,
			lowOffset:  32,
			want:       true,
		},
		{
			name:       "wrong compensated offset",
			lowType:    objabi.R_LOONG32R_ADDR_LO,
			highType:   objabi.R_LOONG32R_ADDR_HI,
			highOffset: 4,
			lowOffset:  16,
			lowAdd:     12,
			want:       false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reporter := &loader.ErrorReporter{}
			ldr := loader.NewLoader(0, reporter)
			text := ldr.LookupOrCreateSym("text", 0)
			target := ldr.LookupOrCreateSym("target", 0)
			sb := ldr.MakeSymbolUpdater(text)

			high, _ := sb.AddRel(test.highType)
			high.SetOff(test.highOffset)
			high.SetSiz(4)
			high.SetSym(target)
			high.SetAdd(test.highAdd)

			low, lowIdx := sb.AddRel(test.lowType)
			low.SetOff(test.lowOffset)
			low.SetSiz(4)
			low.SetSym(target)
			low.SetAdd(test.lowAdd)

			got, ok := pcAddHighReloc(ldr, text, lowIdx)
			if ok != test.want {
				t.Fatalf("pcAddHighReloc ok = %v, want %v", ok, test.want)
			}
			if ok && got.Off() != test.highOffset {
				t.Fatalf("paired high offset = %d, want %d", got.Off(), test.highOffset)
			}
		})
	}
}

func TestValidPCRelativeOffset(t *testing.T) {
	tests := []struct {
		name   string
		offset int64
		bits   uint
		want   bool
	}{
		{name: "B16 minimum", offset: -1 << 17, bits: 18, want: true},
		{name: "B16 maximum", offset: 1<<17 - 4, bits: 18, want: true},
		{name: "B16 below minimum", offset: -1<<17 - 4, bits: 18, want: false},
		{name: "B16 above maximum", offset: 1 << 17, bits: 18, want: false},
		{name: "B21 minimum", offset: -1 << 22, bits: 23, want: true},
		{name: "B21 maximum", offset: 1<<22 - 4, bits: 23, want: true},
		{name: "B26 minimum", offset: -1 << 27, bits: 28, want: true},
		{name: "B26 maximum", offset: 1<<27 - 4, bits: 28, want: true},
		{name: "PCREL20 minimum", offset: -1 << 21, bits: 22, want: true},
		{name: "PCREL20 maximum", offset: 1<<21 - 4, bits: 22, want: true},
		{name: "misaligned", offset: 2, bits: 28, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := validPCRelativeOffset(test.offset, test.bits); got != test.want {
				t.Fatalf("validPCRelativeOffset(%d, %d) = %v, want %v", test.offset, test.bits, got, test.want)
			}
		})
	}
}
