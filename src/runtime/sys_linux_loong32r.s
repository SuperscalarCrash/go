// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && loong32r

#include "go_asm.h"
#include "textflag.h"

#define SYS_openat 56
#define SYS_close 57
#define SYS_pipe2 59
#define SYS_read 63
#define SYS_write 64
#define SYS_exit 93
#define SYS_exit_group 94
#define SYS_futex 98
#define SYS_nanosleep 101
#define SYS_setitimer 103
#define SYS_timer_create 107
#define SYS_timer_settime 110
#define SYS_timer_delete 111
#define SYS_sched_getaffinity 123
#define SYS_sched_yield 124
#define SYS_kill 129
#define SYS_tgkill 131
#define SYS_sigaltstack 132
#define SYS_rt_sigaction 134
#define SYS_rt_sigprocmask 135
#define SYS_getpid 172
#define SYS_gettid 178
#define SYS_brk 214
#define SYS_munmap 215
#define SYS_clone 220
#define SYS_mmap2 222
#define SYS_mincore 232
#define SYS_madvise 233
#define SYS_clock_gettime64 403
#define SYS_timer_settime64 409
#define SYS_futex_time64 422

#define AT_FDCWD -100

TEXT runtime·exit(SB),NOSPLIT|NOFRAME,$0-4
	MOVW	code+0(FP), R4
	MOVW	$SYS_exit_group, R11
	SYSCALL
	UNDEF
	RET

TEXT runtime·exitThread(SB),NOSPLIT|NOFRAME,$0-4
	MOVW	wait+0(FP), R12
	DBAR	$0x12
	MOVW	R0, 0(R12)
	DBAR	$0x12
	MOVW	R0, R4
	MOVW	$SYS_exit, R11
	SYSCALL
	UNDEF
	JMP	0(PC)

TEXT runtime·open(SB),NOSPLIT|NOFRAME,$0-16
	MOVW	name+0(FP), R5
	MOVW	mode+4(FP), R6
	MOVW	perm+8(FP), R7
	MOVW	$AT_FDCWD, R4
	MOVW	$SYS_openat, R11
	SYSCALL
	MOVW	$-4096, R10
	BGEU	R10, R4, open_done
	MOVW	$-1, R4
open_done:
	MOVW	R4, ret+12(FP)
	RET

TEXT runtime·closefd(SB),NOSPLIT|NOFRAME,$0-8
	MOVW	fd+0(FP), R4
	MOVW	$SYS_close, R11
	SYSCALL
	MOVW	$-4096, R10
	BGEU	R10, R4, close_done
	MOVW	$-1, R4
close_done:
	MOVW	R4, ret+4(FP)
	RET

TEXT runtime·write1(SB),NOSPLIT|NOFRAME,$0-16
	MOVW	fd+0(FP), R4
	MOVW	p+4(FP), R5
	MOVW	n+8(FP), R6
	MOVW	$SYS_write, R11
	SYSCALL
	MOVW	R4, ret+12(FP)
	RET

TEXT runtime·read(SB),NOSPLIT|NOFRAME,$0-16
	MOVW	fd+0(FP), R4
	MOVW	p+4(FP), R5
	MOVW	n+8(FP), R6
	MOVW	$SYS_read, R11
	SYSCALL
	MOVW	R4, ret+12(FP)
	RET

// func pipe2(flags int32) (r, w int32, errno int32)
TEXT runtime·pipe2(SB),NOSPLIT|NOFRAME,$0-16
	ADD	$4, R3, R4
	MOVW	flags+0(FP), R5
	MOVW	$SYS_pipe2, R11
	SYSCALL
	MOVW	R4, errno+12(FP)
	RET

TEXT runtime·usleep(SB),NOSPLIT,$16-4
	MOVW	usec+0(FP), R6
	MOVW	$1000000, R7
	DIVU	R7, R6, R8
	REMU	R7, R6, R9
	MOVW	$1000, R10
	MUL	R10, R9
	MOVW	R8, 8(R3)
	MOVW	R9, 12(R3)
	ADD	$8, R3, R4
	MOVW	R0, R5
	MOVW	$SYS_nanosleep, R11
	SYSCALL
	RET

TEXT runtime·gettid(SB),NOSPLIT|NOFRAME,$0-4
	MOVW	$SYS_gettid, R11
	SYSCALL
	MOVW	R4, ret+0(FP)
	RET

TEXT runtime·raise(SB),NOSPLIT|NOFRAME,$0-4
	MOVW	$SYS_getpid, R11
	SYSCALL
	MOVW	R4, R12
	MOVW	$SYS_gettid, R11
	SYSCALL
	MOVW	R4, R5
	MOVW	R12, R4
	MOVW	sig+0(FP), R6
	MOVW	$SYS_tgkill, R11
	SYSCALL
	RET

TEXT runtime·raiseproc(SB),NOSPLIT|NOFRAME,$0-4
	MOVW	$SYS_getpid, R11
	SYSCALL
	MOVW	sig+0(FP), R5
	MOVW	$SYS_kill, R11
	SYSCALL
	RET

TEXT ·getpid(SB),NOSPLIT|NOFRAME,$0-4
	MOVW	$SYS_getpid, R11
	SYSCALL
	MOVW	R4, ret+0(FP)
	RET

TEXT ·tgkill(SB),NOSPLIT|NOFRAME,$0-12
	MOVW	tgid+0(FP), R4
	MOVW	tid+4(FP), R5
	MOVW	sig+8(FP), R6
	MOVW	$SYS_tgkill, R11
	SYSCALL
	RET

TEXT runtime·setitimer(SB),NOSPLIT|NOFRAME,$0-12
	MOVW	mode+0(FP), R4
	MOVW	new+4(FP), R5
	MOVW	old+8(FP), R6
	MOVW	$SYS_setitimer, R11
	SYSCALL
	RET

TEXT runtime·timer_create(SB),NOSPLIT|NOFRAME,$0-16
	MOVW	clockid+0(FP), R4
	MOVW	sevp+4(FP), R5
	MOVW	timerid+8(FP), R6
	MOVW	$SYS_timer_create, R11
	SYSCALL
	MOVW	R4, ret+12(FP)
	RET

TEXT runtime·timer_settime32(SB),NOSPLIT|NOFRAME,$0-20
	MOVW	timerid+0(FP), R4
	MOVW	flags+4(FP), R5
	MOVW	new+8(FP), R6
	MOVW	old+12(FP), R7
	MOVW	$SYS_timer_settime, R11
	SYSCALL
	MOVW	R4, ret+16(FP)
	RET

TEXT runtime·timer_settime64(SB),NOSPLIT|NOFRAME,$0-20
	MOVW	timerid+0(FP), R4
	MOVW	flags+4(FP), R5
	MOVW	new+8(FP), R6
	MOVW	old+12(FP), R7
	MOVW	$SYS_timer_settime64, R11
	SYSCALL
	MOVW	R4, ret+16(FP)
	RET

TEXT runtime·timer_delete(SB),NOSPLIT|NOFRAME,$0-8
	MOVW	timerid+0(FP), R4
	MOVW	$SYS_timer_delete, R11
	SYSCALL
	MOVW	R4, ret+4(FP)
	RET

TEXT runtime·mincore(SB),NOSPLIT|NOFRAME,$0-16
	MOVW	addr+0(FP), R4
	MOVW	n+4(FP), R5
	MOVW	dst+8(FP), R6
	MOVW	$SYS_mincore, R11
	SYSCALL
	MOVW	R4, ret+12(FP)
	RET

TEXT runtime·rtsigprocmask(SB),NOSPLIT|NOFRAME,$0-16
	MOVW	how+0(FP), R4
	MOVW	new+4(FP), R5
	MOVW	old+8(FP), R6
	MOVW	size+12(FP), R7
	MOVW	$SYS_rt_sigprocmask, R11
	SYSCALL
	MOVW	$-4096, R10
	BGEU	R10, R4, 2(PC)
	UNDEF
	RET

TEXT runtime·rt_sigaction(SB),NOSPLIT|NOFRAME,$0-20
	MOVW	sig+0(FP), R4
	MOVW	new+4(FP), R5
	MOVW	old+8(FP), R6
	MOVW	size+12(FP), R7
	MOVW	$SYS_rt_sigaction, R11
	SYSCALL
	MOVW	R4, ret+16(FP)
	RET

TEXT runtime·futex_time32(SB),NOSPLIT|NOFRAME,$0-28
	MOVW	addr+0(FP), R4
	MOVW	op+4(FP), R5
	MOVW	val+8(FP), R6
	MOVW	ts+12(FP), R7
	MOVW	addr2+16(FP), R8
	MOVW	val3+20(FP), R9
	MOVW	$SYS_futex, R11
	SYSCALL
	MOVW	R4, ret+24(FP)
	RET

TEXT runtime·futex_time64(SB),NOSPLIT|NOFRAME,$0-28
	MOVW	addr+0(FP), R4
	MOVW	op+4(FP), R5
	MOVW	val+8(FP), R6
	MOVW	ts+12(FP), R7
	MOVW	addr2+16(FP), R8
	MOVW	val3+20(FP), R9
	MOVW	$SYS_futex_time64, R11
	SYSCALL
	MOVW	R4, ret+24(FP)
	RET

// int32 clone(int32 flags, void *stk, M *mp, G *gp, void (*fn)(void))
TEXT runtime·clone(SB),NOSPLIT|NOFRAME,$0-24
	MOVW	flags+0(FP), R4
	MOVW	stk+4(FP), R5
	MOVW	mp+8(FP), R16
	MOVW	gp+12(FP), R17
	MOVW	fn+16(FP), R18
	ADD	$-16, R5
	MOVW	R16, 0(R5)
	MOVW	R17, 4(R5)
	MOVW	R18, 8(R5)
	MOVW	$1234, R19
	MOVW	R19, 12(R5)
	MOVW	R0, R6
	MOVW	R0, R7
	MOVW	R0, R8
	MOVW	$SYS_clone, R11
	SYSCALL
	BEQ	R4, clone_child
	MOVW	R4, ret+20(FP)
	RET

clone_child:
	NOP	R3
	MOVW	12(R3), R19
	MOVW	$1234, R20
	BEQ	R19, R20, 2(PC)
	UNDEF
	MOVW	$SYS_gettid, R11
	SYSCALL
	MOVW	0(R3), R16
	MOVW	4(R3), R17
	MOVW	8(R3), R18
	BEQ	R16, clone_nog
	BEQ	R17, clone_nog
	MOVW	R4, m_procid(R16)
	MOVW	R16, g_m(R17)
	MOVW	R17, g
clone_nog:
	ADD	$16, R3
	CALL	(R18)
	MOVW	$111, R4
	MOVW	$SYS_exit, R11
	SYSCALL
	UNDEF

TEXT runtime·sigaltstack(SB),NOSPLIT|NOFRAME,$0-8
	MOVW	new+0(FP), R4
	MOVW	old+4(FP), R5
	MOVW	$SYS_sigaltstack, R11
	SYSCALL
	MOVW	$-4096, R10
	BGEU	R10, R4, 2(PC)
	UNDEF
	RET

TEXT runtime·sched_getaffinity(SB),NOSPLIT|NOFRAME,$0-16
	MOVW	pid+0(FP), R4
	MOVW	len+4(FP), R5
	MOVW	buf+8(FP), R6
	MOVW	$SYS_sched_getaffinity, R11
	SYSCALL
	MOVW	R4, ret+12(FP)
	RET

TEXT runtime·sbrk0(SB),NOSPLIT|NOFRAME,$0-4
	MOVW	R0, R4
	MOVW	$SYS_brk, R11
	SYSCALL
	MOVW	R4, ret+0(FP)
	RET

// func sigfwd(fn uintptr, sig uint32, info *siginfo, ctx unsafe.Pointer)
// sigfwd calls a C ABI function, so it needs a real frame to preserve R1 and
// keep the stack 16-byte aligned across the call.
TEXT runtime·sigfwd(SB),NOSPLIT,$0-16
	MOVW	fn+0(FP), R20
	MOVW	sig+4(FP), R4
	MOVW	info+8(FP), R5
	MOVW	ctx+12(FP), R6
	CALL	(R20)
	RET

// func mmap(addr unsafe.Pointer, n uintptr, prot, flags, fd int32, off uint32)
//     (p unsafe.Pointer, err int)
TEXT runtime·mmap(SB),NOSPLIT|NOFRAME,$0-32
	MOVW	addr+0(FP), R4
	MOVW	n+4(FP), R5
	MOVW	prot+8(FP), R6
	MOVW	flags+12(FP), R7
	MOVW	fd+16(FP), R8
	MOVW	off+20(FP), R9
	SRL	$12, R9
	MOVW	$SYS_mmap2, R11
	SYSCALL
	MOVW	$-4096, R10
	BGEU	R10, R4, mmap_ok
	SUB	R4, R0, R10
	MOVW	R0, p+24(FP)
	MOVW	R10, err+28(FP)
	RET
mmap_ok:
	MOVW	R4, p+24(FP)
	MOVW	R0, err+28(FP)
	RET

TEXT runtime·munmap(SB),NOSPLIT|NOFRAME,$0-8
	MOVW	addr+0(FP), R4
	MOVW	n+4(FP), R5
	MOVW	$SYS_munmap, R11
	SYSCALL
	MOVW	$-4096, R10
	BGEU	R10, R4, 2(PC)
	UNDEF
	RET

TEXT runtime·madvise(SB),NOSPLIT|NOFRAME,$0-16
	MOVW	addr+0(FP), R4
	MOVW	n+4(FP), R5
	MOVW	flags+8(FP), R6
	MOVW	$SYS_madvise, R11
	SYSCALL
	MOVW	R4, ret+12(FP)
	RET

TEXT runtime·osyield(SB),NOSPLIT|NOFRAME,$0-0
	MOVW	$SYS_sched_yield, R11
	SYSCALL
	RET

// walltime and nanotime use the time64 ABI unconditionally. LA32R kernels in
// this tree expose the asm-generic clock_gettime64 syscall and a 16-byte
// timespec containing int64 seconds followed by int32 nanoseconds.
TEXT runtime·walltime(SB),NOSPLIT,$24-12
	MOVW	$0, R4
	ADD	$8, R3, R5
	MOVW	$SYS_clock_gettime64, R11
	SYSCALL
	MOVW	8(R3), R6
	MOVW	12(R3), R7
	MOVW	16(R3), R8
	MOVW	R6, sec_lo+0(FP)
	MOVW	R7, sec_hi+4(FP)
	MOVW	R8, nsec+8(FP)
	RET

TEXT runtime·nanotime1(SB),NOSPLIT,$24-8
	MOVW	$1, R4
	ADD	$8, R3, R5
	MOVW	$SYS_clock_gettime64, R11
	SYSCALL
	MOVW	8(R3), R6
	MOVW	12(R3), R7
	MOVW	16(R3), R8
	MOVW	$1000000000, R9
	MUL	R9, R6, R10
	MULHU	R9, R6, R11
	MUL	R9, R7, R12
	ADD	R12, R11
	ADD	R8, R10, R12
	SGTU	R10, R12, R13
	ADD	R13, R11
	MOVW	R12, ret_lo+0(FP)
	MOVW	R11, ret_hi+4(FP)
	RET

// sigtramp is entered using the Linux LA32 C ABI, with the signal number,
// siginfo pointer, and ucontext pointer in R4, R5, and R6 respectively.
TEXT runtime·sigtramp(SB),NOSPLIT|TOPFRAME,$52
	// A signal may interrupt C code. Preserve every ILP32S callee-save
	// integer register before calling into the Go runtime. LA32R is
	// soft-float, so there are no callee-save floating-point registers.
	MOVW	g, 16(R3)
	MOVW	R23, 20(R3)
	MOVW	R24, 24(R3)
	MOVW	R25, 28(R3)
	MOVW	R26, 32(R3)
	MOVW	R27, 36(R3)
	MOVW	R28, 40(R3)
	MOVW	R29, 44(R3)
	MOVW	R30, 48(R3)
	MOVW	R31, 52(R3)

	MOVB	runtime·iscgo(SB), R12
	BEQ	R12, 2(PC)
	CALL	runtime·load_g(SB)

	MOVW	R4, 4(R3)
	MOVW	R5, 8(R3)
	MOVW	R6, 12(R3)
	CALL	runtime·sigtrampgo(SB)

	MOVW	16(R3), g
	MOVW	20(R3), R23
	MOVW	24(R3), R24
	MOVW	28(R3), R25
	MOVW	32(R3), R26
	MOVW	36(R3), R27
	MOVW	40(R3), R28
	MOVW	44(R3), R29
	MOVW	48(R3), R30
	MOVW	52(R3), R31
	RET

TEXT runtime·cgoSigtramp(SB),NOSPLIT|NOFRAME,$0
	JMP	runtime·sigtramp(SB)
