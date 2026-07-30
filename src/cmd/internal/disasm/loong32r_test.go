// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package disasm

import (
	"encoding/binary"
	"strings"
	"testing"
)

func TestDisasmLOONG32R(t *testing.T) {
	tests := []struct {
		word uint32
		want string
	}{
		{0x001010a6, "ADD R4, R5, R6"},
		{0x001290a6, "SGTU R4, R5, R6"},
		{0x02bfe4a6, "ADD $-7, R5, R6"},
		{0x207ffca6, "LL 32764(R5), R6"},
		{0x218000a6, "SC R6, -32768(R5)"},
		{0x288004a6, "MOVW 1(R5), R6"},
		{0x298004a6, "MOVW R6, 1(R5)"},
		{0x2a0004a6, "MOVBU 1(R5), R6"},
		{0x2b0004a6, "MOVF 1(R5), F6"},
		{0x2b4004a6, "MOVF F6, 1(R5)"},
		{0x2b8004a6, "MOVD 1(R5), F6"},
		{0x2bc004a6, "MOVD F6, 1(R5)"},
		{0x4c001081, "JIRL $16, R4, R1"},
		{0x50000400, "JMP 1(PC)"},
		{0x0d0190a6, "FSEL FCC3, F5, F4, F6"},
		{0x0015003f, "MOVW R1, R31"},
	}
	for _, test := range tests {
		var code [4]byte
		binary.LittleEndian.PutUint32(code[:], test.word)
		got, size := disasm_loong32r(code[:], 0x1000, func(uint64) (string, uint64) { return "", 0 }, binary.LittleEndian, false)
		if size != 4 || got != test.want {
			t.Errorf("%08x: got (%q, %d), want (%q, 4)", test.word, got, size, test.want)
		}
		if strings.Contains(got, "MOVV") {
			t.Errorf("%08x: LA32R disassembly used an LA64 move: %q", test.word, got)
		}
	}
}
