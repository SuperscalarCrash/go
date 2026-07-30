// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package work

import (
	"cmd/go/internal/cfg"
	"slices"
	"testing"
)

func TestGCCArchArgsLoong32r(t *testing.T) {
	oldGoarch := cfg.Goarch
	defer func() { cfg.Goarch = oldGoarch }()

	cfg.Goarch = "loong32r"
	got := new(Builder).gccArchArgs()
	want := []string{"-mabi=ilp32s"}
	if !slices.Equal(got, want) {
		t.Fatalf("gccArchArgs() = %q, want %q", got, want)
	}
}
