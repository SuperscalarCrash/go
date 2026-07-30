// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && loong32r

package linux_test

import (
	"internal/runtime/syscall/linux"
	"os"
	"testing"
)

func TestPreadLargeOffset(t *testing.T) {
	f, err := os.CreateTemp(t.TempDir(), "pread-large-offset")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	const offset = int64(5)<<30 + 123
	want := []byte("loong32r-pread64")
	if _, err := f.WriteAt(want, offset); err != nil {
		t.Fatalf("WriteAt offset %#x: %v", offset, err)
	}

	got := make([]byte, len(want))
	n, errno := linux.Pread(int(f.Fd()), got, offset)
	if errno != 0 {
		t.Fatalf("Pread offset %#x: errno %d", offset, errno)
	}
	if n != len(want) || string(got) != string(want) {
		t.Fatalf("Pread offset %#x = %d, %q; want %d, %q", offset, n, got, len(want), want)
	}
}
