// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build loong32r

#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

#define REGCTXT R29

// _rt0_loong32r_lib is called by an ELF .init_array entry in c-archive and
// c-shared binaries. It follows the LA32 ILP32S C ABI. The target is
// soft-float, so only the integer callee-save registers R22-R31 are preserved.
// Slots 4(SP) and 8(SP) are reserved for newosproc0's ABI0 arguments.
TEXT _rt0_loong32r_lib(SB),NOSPLIT,$64-0
	// Address materialization may use R30, so preserve it before touching a
	// global symbol. The fixed 64-byte prologue itself needs no temporary.
	MOVW	R30, 44(R3)
	MOVW	g, 12(R3)
	MOVW	R23, 16(R3)
	MOVW	R24, 20(R3)
	MOVW	R25, 24(R3)
	MOVW	R26, 28(R3)
	MOVW	R27, 32(R3)
	MOVW	R28, 36(R3)
	MOVW	R29, 40(R3)
	MOVW	R31, 48(R3)

	// g must be nil until this thread has entered the Go runtime.
	MOVW	R0, g
	MOVW	R4, _rt0_loong32r_lib_argc<>(SB)
	MOVW	R5, _rt0_loong32r_lib_argv<>(SB)

	CALL	runtime·libpreinit(SB)

	// Start runtime initialization on a new thread, then return to the
	// dynamic loader. cgo supplies the preferred C ABI thread creator.
	MOVW	_cgo_sys_thread_create(SB), R20
	BEQ	R20, lib_nocgo
	MOVW	$_rt0_loong32r_lib_go(SB), R4
	MOVW	R0, R5
	CALL	(R20)
	JMP	lib_restore

lib_nocgo:
	MOVW	$0x800000, R20
	MOVW	$_rt0_loong32r_lib_go(SB), R21
	MOVW	R20, 4(R3)
	MOVW	R21, 8(R3)
	CALL	runtime·newosproc0(SB)

lib_restore:
	MOVW	12(R3), g
	MOVW	16(R3), R23
	MOVW	20(R3), R24
	MOVW	24(R3), R25
	MOVW	28(R3), R26
	MOVW	32(R3), R27
	MOVW	36(R3), R28
	MOVW	40(R3), R29
	MOVW	48(R3), R31
	MOVW	44(R3), R30
	RET

TEXT _rt0_loong32r_lib_go(SB),NOSPLIT|NOFRAME,$0
	MOVW	_rt0_loong32r_lib_argc<>(SB), R4
	MOVW	_rt0_loong32r_lib_argv<>(SB), R5
	JMP	runtime·rt0_go(SB)

DATA _rt0_loong32r_lib_argc<>(SB)/4, $0
GLOBL _rt0_loong32r_lib_argc<>(SB),NOPTR,$4
DATA _rt0_loong32r_lib_argv<>(SB)/4, $0
GLOBL _rt0_loong32r_lib_argv<>(SB),NOPTR,$4

// rt0_go bootstraps the first goroutine from the Linux process stack.
// On entry R3 is SP, R4 is argc, and R5 is argv.
TEXT runtime·rt0_go(SB),NOSPLIT|TOPFRAME,$0
	// Keep the initial stack aligned for both the Go and ILP32S C ABIs.
	ADD	$-16, R3
	MOVW	R4, 4(R3)
	MOVW	R5, 8(R3)

	MOVW	$runtime·g0(SB), g
	MOVW	$(-64*1024), R20
	ADD	R20, R3, R21
	MOVW	R21, g_stackguard0(g)
	MOVW	R21, g_stackguard1(g)
	MOVW	R21, (g_stack+stack_lo)(g)
	MOVW	R3, (g_stack+stack_hi)(g)

	// Let runtime/cgo initialize its thread state when it is linked in.
	// The call follows the LA32 ILP32S ABI and is kept entirely separate
	// from the LA64 runtime startup path.
	MOVW	_cgo_init(SB), R25
	BEQ	R25, nocgo
	MOVW	R0, R7
	MOVW	R0, R6
	MOVW	$setg_gcc<>(SB), R5
	MOVW	g, R4
	CALL	(R25)

nocgo:
	CALL	runtime·save_g(SB)

	MOVW	(g_stack+stack_lo)(g), R20
	ADD	$const_stackGuard, R20
	MOVW	R20, g_stackguard0(g)
	MOVW	R20, g_stackguard1(g)

	MOVW	$runtime·m0(SB), R20
	MOVW	g, m_g0(R20)
	MOVW	R20, g_m(g)

	CALL	runtime·check(SB)
	CALL	runtime·args(SB)
	CALL	runtime·osinit(SB)
	CALL	runtime·schedinit(SB)

	MOVW	$runtime·mainPC(SB), R20
	ADD	$-16, R3
	MOVW	R20, 4(R3)
	MOVW	R0, 0(R3)
	CALL	runtime·newproc(SB)
	ADD	$16, R3

	CALL	runtime·mstart(SB)
	UNDEF
	RET

DATA runtime·mainPC+0(SB)/4,$runtime·main(SB)
GLOBL runtime·mainPC(SB),RODATA,$4

TEXT runtime·breakpoint(SB),NOSPLIT|NOFRAME,$0-0
	UNDEF
	RET

TEXT runtime·asminit(SB),NOSPLIT|NOFRAME,$0-0
	RET

// mstart is the ABI0 entry used when a new thread enters the Go runtime.
TEXT runtime·mstart(SB),NOSPLIT|TOPFRAME,$0
	CALL	runtime·mstart0(SB)
	RET

// func gogo(buf *gobuf)
TEXT runtime·gogo(SB),NOSPLIT|NOFRAME,$0-4
	MOVW	buf+0(FP), R4
	MOVW	gobuf_g(R4), R5
	MOVW	0(R5), R6
	JMP	gogo<>(SB)

TEXT gogo<>(SB),NOSPLIT|NOFRAME,$0
	MOVW	R5, g
	CALL	runtime·save_g(SB)
	MOVW	gobuf_sp(R4), R3
	MOVW	gobuf_lr(R4), R1
	MOVW	gobuf_ctxt(R4), REGCTXT
	MOVW	R0, gobuf_sp(R4)
	MOVW	R0, gobuf_lr(R4)
	MOVW	R0, gobuf_ctxt(R4)
	MOVW	gobuf_pc(R4), R6
	JMP	(R6)

// func mcall(fn func(*g))
TEXT runtime·mcall(SB),NOSPLIT|NOFRAME,$0-4
	MOVW	R3, (g_sched+gobuf_sp)(g)
	MOVW	R1, (g_sched+gobuf_pc)(g)
	MOVW	R0, (g_sched+gobuf_lr)(g)

	MOVW	g, R4
	MOVW	g_m(g), R5
	MOVW	m_g0(R5), g
	CALL	runtime·save_g(SB)
	BNE	g, R4, 2(PC)
	JMP	runtime·badmcall(SB)
	MOVW	fn+0(FP), REGCTXT
	MOVW	0(REGCTXT), R6
	MOVW	(g_sched+gobuf_sp)(g), R3
	ADD	$-16, R3
	MOVW	R4, 4(R3)
	MOVW	R0, 0(R3)
	CALL	(R6)
	JMP	runtime·badmcall2(SB)

TEXT runtime·systemstack_switch(SB),NOSPLIT,$0-0
	UNDEF
	CALL	(R1)
	RET

// func systemstack(fn func())
TEXT runtime·systemstack(SB),NOSPLIT,$0-4
	MOVW	fn+0(FP), REGCTXT
	MOVW	g_m(g), R4
	MOVW	m_gsignal(R4), R5
	BEQ	g, R5, systemstack_noswitch
	MOVW	m_g0(R4), R5
	BEQ	g, R5, systemstack_noswitch
	MOVW	m_curg(R4), R6
	BEQ	g, R6, systemstack_switch
	MOVW	$runtime·badsystemstack(SB), R7
	CALL	(R7)
	CALL	runtime·abort(SB)

systemstack_switch:
	CALL	gosave_systemstack_switch<>(SB)
	MOVW	R5, g
	CALL	runtime·save_g(SB)
	MOVW	(g_sched+gobuf_sp)(g), R3
	MOVW	0(REGCTXT), R6
	CALL	(R6)
	MOVW	g_m(g), R4
	MOVW	m_curg(R4), g
	CALL	runtime·save_g(SB)
	MOVW	(g_sched+gobuf_sp)(g), R3
	MOVW	R0, (g_sched+gobuf_sp)(g)
	RET

systemstack_noswitch:
	MOVW	0(REGCTXT), R4
	MOVW	0(R3), R1
	ADD	$16, R3
	JMP	(R4)

TEXT gosave_systemstack_switch<>(SB),NOSPLIT|NOFRAME,$0
	MOVW	$runtime·systemstack_switch(SB), R20
	ADD	$8, R20 // past the stack-decrement instruction in the prologue
	MOVW	R20, (g_sched+gobuf_pc)(g)
	MOVW	R3, (g_sched+gobuf_sp)(g)
	MOVW	R0, (g_sched+gobuf_lr)(g)
	MOVW	(g_sched+gobuf_ctxt)(g), R20
	BEQ	R20, 2(PC)
	CALL	runtime·abort(SB)
	RET

// func switchToCrashStack0(fn func())
TEXT runtime·switchToCrashStack0(SB),NOSPLIT,$0-4
	MOVW	fn+0(FP), REGCTXT
	MOVW	g_m(g), R4
	MOVW	$runtime·gcrash(SB), g
	CALL	runtime·save_g(SB)
	MOVW	R4, g_m(g)
	MOVW	g, m_g0(R4)
	MOVW	(g_stack+stack_hi)(g), R5
	ADD	$-32, R5, R3
	MOVW	0(REGCTXT), R6
	CALL	(R6)
	CALL	runtime·abort(SB)
	UNDEF

// The stack-growth call sequence passes the interrupted function's caller PC
// in R31 while R1 contains the return PC in the interrupted function.
TEXT runtime·morestack(SB),NOSPLIT|NOFRAME,$0-0
	MOVW	R3, (g_sched+gobuf_sp)(g)
	MOVW	R1, (g_sched+gobuf_pc)(g)
	MOVW	R31, (g_sched+gobuf_lr)(g)
	MOVW	REGCTXT, (g_sched+gobuf_ctxt)(g)

	MOVW	g_m(g), R7
	MOVW	m_g0(R7), R8
	BNE	g, R8, 3(PC)
	CALL	runtime·badmorestackg0(SB)
	CALL	runtime·abort(SB)
	MOVW	m_gsignal(R7), R8
	BNE	g, R8, 3(PC)
	CALL	runtime·badmorestackgsignal(SB)
	CALL	runtime·abort(SB)

	MOVW	R31, (m_morebuf+gobuf_pc)(R7)
	MOVW	R3, (m_morebuf+gobuf_sp)(R7)
	MOVW	g, (m_morebuf+gobuf_g)(R7)

	MOVW	m_g0(R7), g
	CALL	runtime·save_g(SB)
	MOVW	(g_sched+gobuf_sp)(g), R3
	MOVW	R0, -16(R3)
	ADD	$-16, R3
	CALL	runtime·newstack(SB)
	UNDEF

TEXT runtime·morestack_noctxt(SB),NOSPLIT|NOFRAME,$0-0
	MOVW	R3, R3
	MOVW	R0, REGCTXT
	JMP	runtime·morestack(SB)

// reflectcall calls a function using an ABI0 argument frame. LA32R has no
// register arguments in the Go ABI, so all arguments and results are copied
// through the stack frame described by stackArgs.
#define DISPATCH(NAME,MAXSIZE) \
	MOVW	$MAXSIZE, R20; \
	SGTU	R4, R20, R20; \
	BNE	R20, 3(PC); \
	MOVW	$NAME(SB), R21; \
	JMP	(R21)

TEXT ·reflectcall(SB),NOSPLIT|NOFRAME,$0-28
	MOVW	frameSize+20(FP), R4
	DISPATCH(runtime·call16, 16)
	DISPATCH(runtime·call32, 32)
	DISPATCH(runtime·call64, 64)
	DISPATCH(runtime·call128, 128)
	DISPATCH(runtime·call256, 256)
	DISPATCH(runtime·call512, 512)
	DISPATCH(runtime·call1024, 1024)
	DISPATCH(runtime·call2048, 2048)
	DISPATCH(runtime·call4096, 4096)
	DISPATCH(runtime·call8192, 8192)
	DISPATCH(runtime·call16384, 16384)
	DISPATCH(runtime·call32768, 32768)
	DISPATCH(runtime·call65536, 65536)
	DISPATCH(runtime·call131072, 131072)
	DISPATCH(runtime·call262144, 262144)
	DISPATCH(runtime·call524288, 524288)
	DISPATCH(runtime·call1048576, 1048576)
	DISPATCH(runtime·call2097152, 2097152)
	DISPATCH(runtime·call4194304, 4194304)
	DISPATCH(runtime·call8388608, 8388608)
	DISPATCH(runtime·call16777216, 16777216)
	DISPATCH(runtime·call33554432, 33554432)
	DISPATCH(runtime·call67108864, 67108864)
	DISPATCH(runtime·call134217728, 134217728)
	DISPATCH(runtime·call268435456, 268435456)
	DISPATCH(runtime·call536870912, 536870912)
	DISPATCH(runtime·call1073741824, 1073741824)
	MOVW	$runtime·badreflectcall(SB), R20
	JMP	(R20)

#define CALLFN(NAME,MAXSIZE) \
TEXT NAME(SB),WRAPPER,$MAXSIZE-28; \
	NO_LOCAL_POINTERS; \
	/* Copy arguments into the new stack frame. */ \
	MOVW	stackArgs+8(FP), R4; \
	MOVW	stackArgsSize+12(FP), R5; \
	ADD	$4, R3, R6; \
	ADD	R5, R6, R7; \
	BEQ	R6, R7, 6(PC); \
	MOVBU	(R4), R8; \
	ADD	$1, R4; \
	MOVB	R8, (R6); \
	ADD	$1, R6; \
	JMP	-5(PC); \
	/* Invoke the function value with its closure context in R29. */ \
	MOVW	fn+4(FP), REGCTXT; \
	MOVW	(REGCTXT), R20; \
	PCDATA	$PCDATA_StackMapIndex, $0; \
	CALL	(R20); \
	/* Copy stack results back with write barriers where required. */ \
	MOVW	stackArgsType+0(FP), R9; \
	MOVW	stackArgs+8(FP), R4; \
	MOVW	stackArgsSize+12(FP), R5; \
	MOVW	stackRetOffset+16(FP), R8; \
	ADD	$4, R3, R6; \
	ADD	R8, R6; \
	ADD	R8, R4; \
	SUB	R8, R5; \
	CALL	callRet<>(SB); \
	RET

// callRet has a fixed frame for runtime.reflectcallmove's ABI0 arguments.
// On entry R9=type, R4=destination, R6=source, and R5=size.
TEXT callRet<>(SB),NOSPLIT,$20-0
	MOVW	R9, 4(R3)
	MOVW	R4, 8(R3)
	MOVW	R6, 12(R3)
	MOVW	R5, 16(R3)
	MOVW	R0, 20(R3) // no register-result area in the LA32R Go ABI
	CALL	runtime·reflectcallmove(SB)
	RET

CALLFN(·call16, 16)
CALLFN(·call32, 32)
CALLFN(·call64, 64)
CALLFN(·call128, 128)
CALLFN(·call256, 256)
CALLFN(·call512, 512)
CALLFN(·call1024, 1024)
CALLFN(·call2048, 2048)
CALLFN(·call4096, 4096)
CALLFN(·call8192, 8192)
CALLFN(·call16384, 16384)
CALLFN(·call32768, 32768)
CALLFN(·call65536, 65536)
CALLFN(·call131072, 131072)
CALLFN(·call262144, 262144)
CALLFN(·call524288, 524288)
CALLFN(·call1048576, 1048576)
CALLFN(·call2097152, 2097152)
CALLFN(·call4194304, 4194304)
CALLFN(·call8388608, 8388608)
CALLFN(·call16777216, 16777216)
CALLFN(·call33554432, 33554432)
CALLFN(·call67108864, 67108864)
CALLFN(·call134217728, 134217728)
CALLFN(·call268435456, 268435456)
CALLFN(·call536870912, 536870912)
CALLFN(·call1073741824, 1073741824)

TEXT runtime·procyieldAsm(SB),NOSPLIT|NOFRAME,$0-4
	RET

// func cputicks() int64
TEXT runtime·cputicks(SB),NOSPLIT|NOFRAME,$0-8
cputicks_retry:
	RDTIMEHW	R4, R0
	RDTIMELW	R5, R0
	RDTIMEHW	R6, R0
	BNE	R4, R6, cputicks_retry
	MOVW	R5, ret_lo+0(FP)
	MOVW	R4, ret_hi+4(FP)
	RET

TEXT runtime·publicationBarrier(SB),NOSPLIT|NOFRAME,$0-0
	DBAR	$0
	RET

// func asmcgocall(fn, arg unsafe.Pointer) int32
// Call fn(arg) on the scheduler stack, aligned for the LA32 ILP32S C ABI.
TEXT ·asmcgocall(SB),NOSPLIT,$0-12
	MOVW	fn+0(FP), R23
	MOVW	arg+4(FP), R24
	MOVW	R3, R25
	MOVW	g, R26

	// Calls made while creating or tearing down an OS thread can already be
	// on g0 or gsignal. Only switch stacks when called from an ordinary g.
	MOVW	g_m(g), R20
	MOVW	m_gsignal(R20), R21
	BEQ	R21, g, asmcgo_g0
	MOVW	m_g0(R20), R21
	BEQ	R21, g, asmcgo_g0

	CALL	gosave_systemstack_switch<>(SB)
	MOVW	R21, g
	CALL	runtime·save_g(SB)
	MOVW	(g_sched+gobuf_sp)(g), R3

asmcgo_g0:
	// Reserve metadata space and restore the ABI's 16-byte stack alignment.
	ADD	$-16, R3
	AND	$~15, R3
	MOVW	R26, 0(R3)
	MOVW	(g_stack+stack_hi)(R26), R20
	SUB	R25, R20
	MOVW	R20, 4(R3)

	MOVW	R24, R4
	CALL	(R23)
	MOVW	R4, R27

	// Restore by depth rather than by the old SP: a callback may have moved
	// the goroutine stack while C was running.
	MOVW	0(R3), g
	CALL	runtime·save_g(SB)
	MOVW	(g_stack+stack_hi)(g), R20
	MOVW	4(R3), R21
	SUB	R21, R20
	MOVW	R20, R3
	MOVW	R27, ret+8(FP)
	RET

// cgocallback(fn, frame unsafe.Pointer, ctxt uintptr)
// Switch from a C stack to the current goroutine stack and invoke
// runtime.cgocallbackg. The frame layout is the 32-bit Go ABI0 layout.
TEXT ·cgocallback(SB),NOSPLIT,$12-12
	NO_LOCAL_POINTERS

	MOVW	fn+0(FP), R20
	BNE	R20, R0, cgocallback_loadg
	MOVW	frame+4(FP), g
	JMP	cgocallback_dropm

cgocallback_loadg:
	MOVB	runtime·iscgo(SB), R20
	BEQ	R20, R0, cgocallback_nocgo
	CALL	runtime·load_g(SB)

cgocallback_nocgo:
	BEQ	g, R0, cgocallback_needm
	MOVW	g_m(g), R20
	MOVW	R20, savedm-4(SP)
	JMP	cgocallback_havem

cgocallback_needm:
	MOVW	R0, savedm-4(SP)
	MOVW	$runtime·needAndBindM(SB), R20
	CALL	(R20)
	MOVW	g_m(g), R20
	MOVW	m_g0(R20), R21
	MOVW	R3, (g_sched+gobuf_sp)(R21)

cgocallback_havem:
	// Save g0.sched.sp where unwindm expects to find it, then make the
	// current C stack pointer the active g0 scheduling stack pointer.
	MOVW	m_g0(R20), R21
	MOVW	(g_sched+gobuf_sp)(R21), R23
	MOVW	R23, savedsp-12(SP)
	MOVW	R3, (g_sched+gobuf_sp)(R21)

	// Take over m.curg without resuming its saved execution. Place its saved
	// PC below the callback frame so traceback crosses this stack switch.
	MOVW	m_curg(R20), g
	CALL	runtime·save_g(SB)
	MOVW	(g_sched+gobuf_sp)(g), R23
	MOVW	(g_sched+gobuf_pc)(g), R24
	MOVW	R24, -16(R23)

	MOVW	fn+0(FP), R25
	MOVW	frame+4(FP), R26
	MOVW	ctxt+8(FP), R27
	ADD	$-16, R23, R3
	MOVW	R25, 4(R3)
	MOVW	R26, 8(R3)
	MOVW	R27, 12(R3)
	CALL	runtime·cgocallbackg(SB)

	MOVW	0(R3), R24
	MOVW	R24, (g_sched+gobuf_pc)(g)
	ADD	$16, R3, R23
	MOVW	R23, (g_sched+gobuf_sp)(g)

	MOVW	g_m(g), R20
	MOVW	m_g0(R20), g
	CALL	runtime·save_g(SB)
	MOVW	(g_sched+gobuf_sp)(g), R3
	MOVW	savedsp-12(SP), R23
	MOVW	R23, (g_sched+gobuf_sp)(g)

	MOVW	savedm-4(SP), R20
	BNE	R20, R0, cgocallback_done
	MOVW	_cgo_pthread_key_created(SB), R20
	BEQ	R20, R0, cgocallback_dropm
	MOVW	(R20), R20
	BNE	R20, R0, cgocallback_done

cgocallback_dropm:
	MOVW	$runtime·dropm(SB), R20
	CALL	(R20)

cgocallback_done:
	RET

// func setg(gg *g)
TEXT runtime·setg(SB),NOSPLIT,$0-4
	MOVW	gg+0(FP), g
	CALL	runtime·save_g(SB)
	RET

// setg_gcc is called with the LA32 C ABI and records g in both R22 and TLS.
TEXT setg_gcc<>(SB),NOSPLIT,$4-0
	MOVW	R4, g
	MOVW	R30, savedR30-4(SP)
	CALL	runtime·save_g(SB)
	MOVW	savedR30-4(SP), R30
	RET

// _cgo_topofstack is called by cgo-generated C wrappers. It follows the C ABI
// and preserves the callee-save registers touched by load_g.
TEXT _cgo_topofstack(SB),NOSPLIT,$8-0
	MOVW	g, saveG-4(SP)
	MOVW	R30, saveR30-8(SP)
	CALL	runtime·load_g(SB)
	MOVW	g_m(g), R20
	MOVW	m_curg(R20), R20
	MOVW	(g_stack+stack_hi)(R20), R4
	MOVW	saveR30-8(SP), R30
	MOVW	saveG-4(SP), g
	RET

TEXT runtime·abort(SB),NOSPLIT|NOFRAME,$0-0
	UNDEF

TEXT runtime·memhash(SB),NOSPLIT|NOFRAME,$0-16
	JMP	runtime·memhashFallback(SB)
TEXT runtime·strhash(SB),NOSPLIT|NOFRAME,$0-12
	JMP	runtime·strhashFallback(SB)
TEXT runtime·memhash32(SB),NOSPLIT|NOFRAME,$0-12
	JMP	runtime·memhash32Fallback(SB)
TEXT runtime·memhash64(SB),NOSPLIT|NOFRAME,$0-12
	JMP	runtime·memhash64Fallback(SB)

TEXT ·checkASM(SB),NOSPLIT|NOFRAME,$0-1
	MOVW	$1, R4
	MOVB	R4, ret+0(FP)
	RET

// This is called from .init_array and follows the LA32 ILP32S C ABI, not the
// Go ABI. The linker-generated caller passes the moduledata pointer in R4.
TEXT runtime·addmoduledata(SB),NOSPLIT|NOFRAME,$0-0
	ADD	$-16, R3
	// Address materialization for the globals below may clobber R30, which is
	// callee-save in ILP32S even though the Go assembler reserves it as REGTMP.
	MOVW	R30, 12(R3)
	MOVW	runtime·lastmoduledatap(SB), R12
	MOVW	R4, moduledata_next(R12)
	MOVW	R4, runtime·lastmoduledatap(SB)
	MOVW	12(R3), R30
	ADD	$16, R3
	RET

// gcWriteBarrier is a private compiler/runtime convention. R25 carries the
// requested byte count on entry and the write-barrier buffer pointer on exit.
// All ordinary allocatable registers are caller-save for LoweredWB.
TEXT gcWriteBarrier<>(SB),NOSPLIT,$16
	MOVW	R25, 4(R3)
gcwb_retry:
	MOVW	g_m(g), R4
	MOVW	m_p(R4), R4
	MOVW	(p_wbBuf+wbBuf_next)(R4), R5
	MOVW	(p_wbBuf+wbBuf_end)(R4), R6
	MOVW	4(R3), R25
	ADD	R25, R5, R7
	SGTU	R7, R6, R8
	BNE	R8, gcwb_flush
	MOVW	R7, (p_wbBuf+wbBuf_next)(R4)
	MOVW	R5, R25
	RET

gcwb_flush:
	CALL	runtime·wbBufFlush<ABIInternal>(SB)
	JMP	gcwb_retry

TEXT runtime·gcWriteBarrier1<ABIInternal>(SB),NOSPLIT|NOFRAME,$0
	MOVW	$4, R25
	JMP	gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier2<ABIInternal>(SB),NOSPLIT|NOFRAME,$0
	MOVW	$8, R25
	JMP	gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier3<ABIInternal>(SB),NOSPLIT|NOFRAME,$0
	MOVW	$12, R25
	JMP	gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier4<ABIInternal>(SB),NOSPLIT|NOFRAME,$0
	MOVW	$16, R25
	JMP	gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier5<ABIInternal>(SB),NOSPLIT|NOFRAME,$0
	MOVW	$20, R25
	JMP	gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier6<ABIInternal>(SB),NOSPLIT|NOFRAME,$0
	MOVW	$24, R25
	JMP	gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier7<ABIInternal>(SB),NOSPLIT|NOFRAME,$0
	MOVW	$28, R25
	JMP	gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier8<ABIInternal>(SB),NOSPLIT|NOFRAME,$0
	MOVW	$32, R25
	JMP	gcWriteBarrier<>(SB)

// Bounds failures encode their operands as indices into R4..R19. Preserve
// that exact bank and pass it to the architecture-neutral decoder.
TEXT runtime·panicBounds<ABIInternal>(SB),NOSPLIT,$72-0
	NO_LOCAL_POINTERS
	MOVW	R4, 12(R3)
	MOVW	R5, 16(R3)
	MOVW	R6, 20(R3)
	MOVW	R7, 24(R3)
	MOVW	R8, 28(R3)
	MOVW	R9, 32(R3)
	MOVW	R10, 36(R3)
	MOVW	R11, 40(R3)
	MOVW	R12, 44(R3)
	MOVW	R13, 48(R3)
	MOVW	R14, 52(R3)
	MOVW	R15, 56(R3)
	MOVW	R16, 60(R3)
	MOVW	R17, 64(R3)
	MOVW	R18, 68(R3)
	MOVW	R19, 72(R3)
	MOVW	R1, 4(R3)
	ADD	$12, R3, R4
	MOVW	R4, 8(R3)
	CALL	runtime·panicBounds32<ABIInternal>(SB)
	RET

TEXT runtime·panicExtend<ABIInternal>(SB),NOSPLIT,$72-0
	NO_LOCAL_POINTERS
	MOVW	R4, 12(R3)
	MOVW	R5, 16(R3)
	MOVW	R6, 20(R3)
	MOVW	R7, 24(R3)
	MOVW	R8, 28(R3)
	MOVW	R9, 32(R3)
	MOVW	R10, 36(R3)
	MOVW	R11, 40(R3)
	MOVW	R12, 44(R3)
	MOVW	R13, 48(R3)
	MOVW	R14, 52(R3)
	MOVW	R15, 56(R3)
	MOVW	R16, 60(R3)
	MOVW	R17, 64(R3)
	MOVW	R18, 68(R3)
	MOVW	R19, 72(R3)
	MOVW	R1, 4(R3)
	ADD	$12, R3, R4
	MOVW	R4, 8(R3)
	CALL	runtime·panicBounds32X<ABIInternal>(SB)
	RET

// asyncPreempt is entered by rewriting a signal context as if the interrupted
// code had called this function. The signal setup saves the interrupted LR at
// 0(SP), installs the resume PC in R1, and moves SP down by one aligned frame.
//
// Save every compiler-visible general register except g (R22), SP (R3), and
// the architectural zero register. R30 is reserved as the linker temporary,
// so it may be used for the final indirect jump.
TEXT runtime·asyncPreempt(SB),NOSPLIT|NOFRAME,$0-0
	// pushCall has already consumed a 16-byte synthetic frame. Reserve a
	// further 128 bytes so both this entry and the call to asyncPreempt2
	// observe the architecture's StackAlign.
	MOVW	R1, -128(R3)
	SUB	$128, R3
	MOVW	R2, 4(R3)
	MOVW	R4, 8(R3)
	MOVW	R5, 12(R3)
	MOVW	R6, 16(R3)
	MOVW	R7, 20(R3)
	MOVW	R8, 24(R3)
	MOVW	R9, 28(R3)
	MOVW	R10, 32(R3)
	MOVW	R11, 36(R3)
	MOVW	R12, 40(R3)
	MOVW	R13, 44(R3)
	MOVW	R14, 48(R3)
	MOVW	R15, 52(R3)
	MOVW	R16, 56(R3)
	MOVW	R17, 60(R3)
	MOVW	R18, 64(R3)
	MOVW	R19, 68(R3)
	MOVW	R20, 72(R3)
	MOVW	R21, 76(R3)
	MOVW	R23, 80(R3)
	MOVW	R24, 84(R3)
	MOVW	R25, 88(R3)
	MOVW	R26, 92(R3)
	MOVW	R27, 96(R3)
	MOVW	R28, 100(R3)
	MOVW	R29, 104(R3)
	MOVW	R30, 108(R3)
	MOVW	R31, 112(R3)

	CALL	runtime·asyncPreempt2(SB)

	MOVW	112(R3), R31
	MOVW	104(R3), R29
	MOVW	100(R3), R28
	MOVW	96(R3), R27
	MOVW	92(R3), R26
	MOVW	88(R3), R25
	MOVW	84(R3), R24
	MOVW	80(R3), R23
	MOVW	76(R3), R21
	MOVW	72(R3), R20
	MOVW	68(R3), R19
	MOVW	64(R3), R18
	MOVW	60(R3), R17
	MOVW	56(R3), R16
	MOVW	52(R3), R15
	MOVW	48(R3), R14
	MOVW	44(R3), R13
	MOVW	40(R3), R12
	MOVW	36(R3), R11
	MOVW	32(R3), R10
	MOVW	28(R3), R9
	MOVW	24(R3), R8
	MOVW	20(R3), R7
	MOVW	16(R3), R6
	MOVW	12(R3), R5
	MOVW	8(R3), R4
	MOVW	4(R3), R2
	// The first slot holds the resume PC saved from R1. pushCall left the
	// interrupted function's LR at the bottom of its 16-byte synthetic frame.
	// Restore that LR to R1, then use reserved REGTMP to jump back to the
	// interrupted PC.
	MOVW	128(R3), R1
	MOVW	0(R3), R30
	ADD	$144, R3
	JMP	(R30)

// The top-most function running on a goroutine returns to goexit+PCQuantum.
TEXT runtime·goexit(SB),NOSPLIT|NOFRAME|TOPFRAME,$0-0
	OR	R0, R0, R0
	CALL	runtime·goexit1(SB)
	OR	R0, R0, R0
