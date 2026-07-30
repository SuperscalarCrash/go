// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build loong32r

package cpu

import "testing"

func TestHostByteOrderLoong32r(t *testing.T) {
	order := hostByteOrder()
	if got, want := order.Uint32([]byte{1, 2, 3, 4}), uint32(0x04030201); got != want {
		t.Fatalf("host byte order Uint32 = %#x, want %#x", got, want)
	}
	if got, want := order.Uint64([]byte{1, 2, 3, 4, 5, 6, 7, 8}), uint64(0x0807060504030201); got != want {
		t.Fatalf("host byte order Uint64 = %#x, want %#x", got, want)
	}
}
