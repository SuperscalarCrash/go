// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package loong32r

import (
	"cmd/internal/obj"
	"fmt"
)

func init() {
	obj.RegisterRegister(obj.RBaseLOONG32R, REG_LAST+1, rconv)
}

func rconv(r int) string {
	switch {
	case r == 0:
		return "NONE"
	case r == REGG:
		return "g"
	case REG_R0 <= r && r <= REG_R31:
		return fmt.Sprintf("R%d", r-REG_R0)
	case REG_F0 <= r && r <= REG_F31:
		return fmt.Sprintf("F%d", r-REG_F0)
	case REG_FCSR0 <= r && r <= REG_FCSR3:
		return fmt.Sprintf("FCSR%d", r-REG_FCSR0)
	case REG_FCC0 <= r && r <= REG_FCC7:
		return fmt.Sprintf("FCC%d", r-REG_FCC0)
	default:
		return fmt.Sprintf("badreg(%d)", r-obj.RBaseLOONG32R)
	}
}

func DRconv(a int) string {
	if a >= C_NONE && a < C_NCLASS {
		return cnames0[a]
	}
	return "C_??"
}
