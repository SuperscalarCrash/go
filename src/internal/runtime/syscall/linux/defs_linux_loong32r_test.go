// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build loong32r

package linux_test

import (
	"internal/runtime/syscall/linux"
	"testing"
	"unsafe"
)

func TestEpollEventLayout(t *testing.T) {
	var event linux.EpollEvent
	if got, want := unsafe.Sizeof(event), uintptr(16); got != want {
		t.Errorf("sizeof(EpollEvent) = %d, want %d", got, want)
	}
	if got, want := unsafe.Offsetof(event.Data), uintptr(8); got != want {
		t.Errorf("offsetof(EpollEvent.Data) = %d, want %d", got, want)
	}
}
