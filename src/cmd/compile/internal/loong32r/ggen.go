// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package loong32r

import (
	"cmd/compile/internal/base"
	"cmd/compile/internal/objw"
	"cmd/compile/internal/types"
	"cmd/internal/obj"
	"cmd/internal/obj/loong32r"
)

func padframe(frame int64) int64 {
	// The LA32R ABI requires SP to be 16-byte aligned. The assembler adds
	// FixedFrameSize bytes for the saved link register, so pad the compiler's
	// local frame such that the complete frame is aligned. Doing this here is
	// important: stack maps and zeroing code are generated from this size and
	// must agree with the frame installed by the assembler.
	const stackAlign = int64(16)
	frame += base.Ctxt.Arch.FixedFrameSize
	frame = (frame + stackAlign - 1) &^ (stackAlign - 1)
	return frame - base.Ctxt.Arch.FixedFrameSize
}

func zerorange(pp *objw.Progs, p *obj.Prog, off, cnt int64, _ *uint32) *obj.Prog {
	if cnt%int64(types.PtrSize) != 0 {
		panic("zeroed region not aligned")
	}

	for cnt != 0 {
		p = pp.Append(p, loong32r.AMOVW, obj.TYPE_REG, loong32r.REGZERO, 0, obj.TYPE_MEM, loong32r.REGSP, base.Ctxt.Arch.FixedFrameSize+off)
		cnt -= int64(types.PtrSize)
		off += int64(types.PtrSize)
	}

	return p
}

func ginsnop(pp *objw.Progs) *obj.Prog {
	p := pp.Prog(loong32r.ANOOP)
	return p
}
