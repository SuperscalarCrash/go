// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build loong32r

// Export the 64-bit operations used by sync/atomic.
//
//go:linkname Xadd64
//go:linkname Xchg64
//go:linkname Cas64
//go:linkname Load64
//go:linkname Store64
//go:linkname Or64
//go:linkname And64

package atomic

import (
	"internal/cpu"
	"unsafe"
)

// LA32R has only 32-bit LL/SC. A cache-line-isolated lock serializes 64-bit
// operations so that they remain atomic on a 32-bit machine.
var lock64 struct {
	state uint32
	pad   [cpu.CacheLinePadSize - 4]byte
}

//go:noescape
func spinLock(state *uint32)

//go:noescape
func spinUnlock(state *uint32)

//go:nosplit
func lock64Addr(addr *uint64) {
	if uintptr(unsafe.Pointer(addr))&7 != 0 {
		panicUnaligned()
	}
	_ = *addr
	spinLock(&lock64.state)
}

//go:nosplit
func unlock64() { spinUnlock(&lock64.state) }

//go:nosplit
func Xadd64(addr *uint64, delta int64) uint64 {
	lock64Addr(addr)
	value := *addr + uint64(delta)
	*addr = value
	unlock64()
	return value
}

//go:nosplit
func Xchg64(addr *uint64, value uint64) uint64 {
	lock64Addr(addr)
	old := *addr
	*addr = value
	unlock64()
	return old
}

//go:nosplit
func Cas64(addr *uint64, old, value uint64) bool {
	lock64Addr(addr)
	if *addr == old {
		*addr = value
		unlock64()
		return true
	}
	unlock64()
	return false
}

//go:nosplit
func Load64(addr *uint64) uint64 {
	lock64Addr(addr)
	value := *addr
	unlock64()
	return value
}

//go:nosplit
func Store64(addr *uint64, value uint64) {
	lock64Addr(addr)
	*addr = value
	unlock64()
}

//go:nosplit
func Or64(addr *uint64, value uint64) uint64 {
	lock64Addr(addr)
	old := *addr
	*addr = old | value
	unlock64()
	return old
}

//go:nosplit
func And64(addr *uint64, value uint64) uint64 {
	lock64Addr(addr)
	old := *addr
	*addr = old & value
	unlock64()
	return old
}

//go:noescape
func Xadd(ptr *uint32, delta int32) uint32

//go:noescape
func Xadduintptr(ptr *uintptr, delta uintptr) uintptr

//go:noescape
func Xchg(ptr *uint32, value uint32) uint32

func Xchg8(ptr *uint8, value uint8) uint8 { return goXchg8(ptr, value) }

//go:noescape
func Xchguintptr(ptr *uintptr, value uintptr) uintptr

//go:noescape
func Load(ptr *uint32) uint32

//go:noescape
func Load8(ptr *uint8) uint8

// The pointed-to value escapes when the result escapes.
func Loadp(ptr unsafe.Pointer) unsafe.Pointer

//go:noescape
func LoadAcq(ptr *uint32) uint32

//go:noescape
func LoadAcquintptr(ptr *uintptr) uintptr

//go:nosplit
func Or8(ptr *uint8, value uint8) {
	word := (*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) &^ 3))
	shift := (uintptr(unsafe.Pointer(ptr)) & 3) * 8
	bits := uint32(value) << shift
	for {
		old := *word
		if Cas(word, old, old|bits) {
			return
		}
	}
}

//go:nosplit
func And8(ptr *uint8, value uint8) {
	word := (*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) &^ 3))
	shift := (uintptr(unsafe.Pointer(ptr)) & 3) * 8
	mask := uint32(0xff) << shift
	bits := uint32(value)<<shift | ^mask
	for {
		old := *word
		if Cas(word, old, old&bits) {
			return
		}
	}
}

//go:noescape
func And(ptr *uint32, value uint32)

//go:noescape
func Or(ptr *uint32, value uint32)

//go:noescape
func And32(ptr *uint32, value uint32) uint32

//go:noescape
func Or32(ptr *uint32, value uint32) uint32

//go:noescape
func Anduintptr(ptr *uintptr, value uintptr) uintptr

//go:noescape
func Oruintptr(ptr *uintptr, value uintptr) uintptr

//go:noescape
func Store(ptr *uint32, value uint32)

//go:noescape
func Store8(ptr *uint8, value uint8)

// The stored pointer escapes.
func StorepNoWB(ptr unsafe.Pointer, value unsafe.Pointer)

//go:noescape
func StoreRel(ptr *uint32, value uint32)

//go:noescape
func StoreReluintptr(ptr *uintptr, value uintptr)

//go:noescape
func CasRel(ptr *uint32, old, value uint32) bool
