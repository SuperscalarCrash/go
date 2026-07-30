// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build loong32r

#include "textflag.h"
#include "funcdata.h"

#define REGCTXT R29

// makeFuncStub is the code half of the function returned by MakeFunc.
// There is no argument size here; the runtime obtains the argument map from
// the function value installed by MakeFunc.
TEXT ·makeFuncStub(SB),(NOSPLIT|WRAPPER),$20
	NO_LOCAL_POINTERS
	MOVW	REGCTXT, 4(R3)
	MOVW	$argframe+0(FP), R4
	MOVW	R4, 8(R3)
	MOVB	R0, 20(R3)
	ADD	$20, R3, R4
	MOVW	R4, 12(R3)
	MOVW	R0, 16(R3) // no register arguments in the LA32R Go ABI
	CALL	·callReflect(SB)
	RET

// methodValueCall is the code half of the function returned by
// makeMethodValue. Its stack contract is the same as makeFuncStub's.
TEXT ·methodValueCall(SB),(NOSPLIT|WRAPPER),$20
	NO_LOCAL_POINTERS
	MOVW	REGCTXT, 4(R3)
	MOVW	$argframe+0(FP), R4
	MOVW	R4, 8(R3)
	MOVB	R0, 20(R3)
	ADD	$20, R3, R4
	MOVW	R4, 12(R3)
	MOVW	R0, 16(R3) // no register arguments in the LA32R Go ABI
	CALL	·callMethod(SB)
	RET
