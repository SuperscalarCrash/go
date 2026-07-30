// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objfile

import (
	"debug/elf"
	"testing"
)

func TestELFLoongArchClass(t *testing.T) {
	tests := []struct {
		class elf.Class
		want  string
	}{
		{elf.ELFCLASS32, "loong32r"},
		{elf.ELFCLASS64, "loong64"},
	}
	for _, test := range tests {
		f := &elfFile{elf: &elf.File{FileHeader: elf.FileHeader{Class: test.class, Machine: elf.EM_LOONGARCH}}}
		if got := f.goarch(); got != test.want {
			t.Errorf("ELF class %v: got GOARCH %q, want %q", test.class, got, test.want)
		}
	}
}
