// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arch

import (
	"cmd/internal/obj"
	"cmd/internal/obj/loong32r"
	"strings"
)

func jumpLoong32r(word string) bool {
	return word == "CALL" || word == "JAL" || word == "JMP" || strings.HasPrefix(word, "B")
}

func loong32rRegisterNumber(name string, n int16) (int16, bool) {
	switch name {
	case "R":
		if 0 <= n && n < loong32r.NREG {
			return loong32r.REG_R0 + n, true
		}
	case "F":
		if 0 <= n && n < loong32r.NFREG {
			return loong32r.REG_F0 + n, true
		}
	case "FCSR":
		if 0 <= n && n < 4 {
			return loong32r.REG_FCSR0 + n, true
		}
	case "FCC":
		if 0 <= n && n < 8 {
			return loong32r.REG_FCC0 + n, true
		}
	}
	return 0, false
}

// IsLoong32rRDTIME reports the two-output timer instructions in LA32R.
func IsLoong32rRDTIME(op obj.As) bool {
	return op == loong32r.ARDTIMELW || op == loong32r.ARDTIMEHW
}

// IsLoong32rPRELD reports the LA32R prefetch instruction. Its hint operand
// is an immediate rather than a register, so the generic two-operand parser
// must preserve it as an additional source operand.
func IsLoong32rPRELD(op obj.As) bool {
	return op == loong32r.APRELD
}
