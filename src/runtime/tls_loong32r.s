// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build loong32r

#include "go_asm.h"
#include "go_tls.h"
#include "funcdata.h"
#include "textflag.h"

// In cgo binaries, R22 is mirrored in ELF TLS so callbacks arriving on an
// arbitrary C thread can recover g. LA32R TLS address materialization uses
// the architecture's own TLS_LE/TLS_IE relocations and clobbers only R30.
TEXT runtime·save_g(SB),NOSPLIT|NOFRAME,$0-0
	MOVB	runtime·iscgo(SB), R30
	BEQ	R30, tls_nocgo
	MOVW	g, runtime·tls_g(SB)
tls_nocgo:
	RET

TEXT runtime·load_g(SB),NOSPLIT|NOFRAME,$0-0
	MOVW	runtime·tls_g(SB), g
	RET

GLOBL runtime·tls_g(SB), TLSBSS, $4
