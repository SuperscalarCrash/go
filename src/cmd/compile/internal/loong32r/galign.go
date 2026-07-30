// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package loong32r

import (
	"cmd/compile/internal/ssa"
	"cmd/compile/internal/ssagen"
	"cmd/internal/obj/loong32r"
)

func Init(arch *ssagen.ArchInfo) {
	arch.LinkArch = &loong32r.Linkloong32r
	arch.REGSP = loong32r.REGSP
	arch.MAXWIDTH = (1 << 31) - 1
	// The floating-point extension is optional in LA32R. Keep the baseline
	// GOARCH usable on implementations without an FPU.
	arch.SoftFloat = true
	arch.ZeroRange = zerorange
	arch.Ginsnop = ginsnop
	arch.SSAMarkMoves = func(s *ssagen.State, b *ssa.Block) {}
	arch.SSAGenValue = ssaGenValue
	arch.SSAGenBlock = ssaGenBlock
}
