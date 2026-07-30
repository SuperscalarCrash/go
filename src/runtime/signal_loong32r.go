// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && loong32r

package runtime

import (
	"internal/abi"
	"internal/runtime/sys"
	"unsafe"
)

func dumpregs(c *sigctxt) {
	regs := c.regs()
	for i := 0; i < len(regs.sc_regs); i++ {
		print("r", i, "   ", hex(regs.sc_regs[i]))
		if i&1 == 0 {
			print("\t")
		} else {
			print("\n")
		}
	}
	print("pc   ", hex(c.pc()), "\tlink ", hex(c.link()), "\n")
}

//go:nosplit
//go:nowritebarrierrec
func (c *sigctxt) sigpc() uintptr { return uintptr(c.pc()) }

func (c *sigctxt) sigsp() uintptr { return uintptr(c.sp()) }
func (c *sigctxt) siglr() uintptr { return uintptr(c.link()) }
func (c *sigctxt) fault() uintptr { return uintptr(c.sigaddr()) }

func (c *sigctxt) preparePanic(sig uint32, gp *g) {
	// Save the interrupted link at the bottom of a complete aligned frame.
	// sigpanic is entered as an ordinary Go function, so its incoming SP must
	// still satisfy the LA32R 16-byte stack alignment.
	sp := c.sp() - sys.StackAlign
	c.set_sp(sp)
	*(*uint32)(unsafe.Pointer(uintptr(sp))) = c.link()

	pc := gp.sigpc
	if shouldPushSigpanic(gp, pc, uintptr(c.link())) {
		c.set_link(uint32(pc))
	}
	c.set_r22(uint32(uintptr(unsafe.Pointer(gp))))
	c.set_pc(uint32(abi.FuncPCABIInternal(sigpanic)))
}

func (c *sigctxt) pushCall(targetPC, resumePC uintptr) {
	// Keep the injected function's entry SP aligned. The unwinder also
	// advances over this synthetic frame by alignUp(MinFrameSize,
	// StackAlign), so reserving only the link word would corrupt the caller's
	// SP during traceback.
	sp := c.sp() - sys.StackAlign
	c.set_sp(sp)
	*(*uint32)(unsafe.Pointer(uintptr(sp))) = c.link()
	c.set_link(uint32(resumePC))
	c.set_pc(uint32(targetPC))
}
