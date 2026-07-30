// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && loong32r

package runtime

import "unsafe"

// Constants and layouts follow the Linux LA32 UAPI and asm-generic ABI.
const (
	_EINTR  = 0x4
	_EAGAIN = 0xb
	_ENOMEM = 0xc
	_ENOSYS = 0x26

	_PROT_NONE  = 0
	_PROT_READ  = 0x1
	_PROT_WRITE = 0x2
	_PROT_EXEC  = 0x4

	_MAP_ANON    = 0x20
	_MAP_PRIVATE = 0x2
	_MAP_FIXED   = 0x10

	_MADV_DONTNEED   = 0x4
	_MADV_FREE       = 0x8
	_MADV_HUGEPAGE   = 0xe
	_MADV_NOHUGEPAGE = 0xf
	_MADV_COLLAPSE   = 0x19

	_SA_RESTART  = 0x10000000
	_SA_ONSTACK  = 0x08000000
	_SA_SIGINFO  = 0x4
	_SA_RESTORER = 0

	_SI_KERNEL = 0x80
	_SI_TIMER  = -2

	_SIGHUP    = 1
	_SIGINT    = 2
	_SIGQUIT   = 3
	_SIGILL    = 4
	_SIGTRAP   = 5
	_SIGABRT   = 6
	_SIGBUS    = 7
	_SIGFPE    = 8
	_SIGKILL   = 9
	_SIGUSR1   = 10
	_SIGSEGV   = 11
	_SIGUSR2   = 12
	_SIGPIPE   = 13
	_SIGALRM   = 14
	_SIGSTKFLT = 16
	_SIGCHLD   = 17
	_SIGCONT   = 18
	_SIGSTOP   = 19
	_SIGTSTP   = 20
	_SIGTTIN   = 21
	_SIGTTOU   = 22
	_SIGURG    = 23
	_SIGXCPU   = 24
	_SIGXFSZ   = 25
	_SIGVTALRM = 26
	_SIGPROF   = 27
	_SIGWINCH  = 28
	_SIGIO     = 29
	_SIGPWR    = 30
	_SIGSYS    = 31
	_SIGRTMIN  = 32

	_FPE_INTDIV = 1
	_FPE_INTOVF = 2
	_FPE_FLTDIV = 3
	_FPE_FLTOVF = 4
	_FPE_FLTUND = 5
	_FPE_FLTRES = 6
	_FPE_FLTINV = 7
	_FPE_FLTSUB = 8

	_BUS_ADRALN = 1
	_BUS_ADRERR = 2
	_BUS_OBJERR = 3

	_SEGV_MAPERR = 1
	_SEGV_ACCERR = 2

	_ITIMER_REAL    = 0
	_ITIMER_VIRTUAL = 1
	_ITIMER_PROF    = 2

	_CLOCK_THREAD_CPUTIME_ID = 3
	_SIGEV_THREAD_ID         = 4

	_O_RDONLY   = 0
	_O_WRONLY   = 1
	_O_CREAT    = 0x40
	_O_TRUNC    = 0x200
	_O_NONBLOCK = 0x800
	_O_CLOEXEC  = 0x80000

	_AF_UNIX    = 1
	_SOCK_DGRAM = 2
)

type timespec32 struct {
	tv_sec  int32
	tv_nsec int32
}

//go:nosplit
func (ts *timespec32) setNsec(ns int64) {
	ts.tv_sec = int32(ns / 1e9)
	ts.tv_nsec = int32(ns % 1e9)
}

// The time64 ABI aligns its signed 64-bit seconds field to eight bytes.
type timespec struct {
	tv_sec  int64
	tv_nsec int32
	_       int32
}

//go:nosplit
func (ts *timespec) setNsec(ns int64) {
	ts.tv_sec = ns / 1e9
	ts.tv_nsec = int32(ns % 1e9)
}

type timeval struct {
	tv_sec  int32
	tv_usec int32
}

//go:nosplit
func (tv *timeval) set_usec(x int32) { tv.tv_usec = x }

type itimerspec32 struct {
	it_interval timespec32
	it_value    timespec32
}

type itimerspec struct {
	it_interval timespec
	it_value    timespec
}

type itimerval struct {
	it_interval timeval
	it_value    timeval
}

type sigeventFields struct {
	value                  uintptr
	signo                  int32
	notify                 int32
	sigev_notify_thread_id int32
}

type sigevent struct {
	sigeventFields
	_ [_sigev_max_size - unsafe.Sizeof(sigeventFields{})]byte
}

type siginfoFields struct {
	si_signo int32
	si_errno int32
	si_code  int32
	si_addr  uint32
}

type siginfo struct {
	siginfoFields
	_ [_si_max_size - unsafe.Sizeof(siginfoFields{})]byte
}

// LA32 has no kernel sa_restorer field. Keep a trailing field because common
// runtime code names it; the kernel consumes only the preceding layout.
type sigactiont struct {
	sa_handler  uintptr
	sa_flags    uint32
	sa_mask     uint64
	sa_restorer uintptr
}

type stackt struct {
	ss_sp    *byte
	ss_flags int32
	ss_size  uintptr
}

type sigcontext struct {
	// The LA32R kernel signal ABI uses 64-bit save slots for the PC and
	// general registers even though user pointers and register values are
	// 32-bit. Keep these slots at their kernel widths so the register indices
	// line up with the signal frame restored by rt_sigreturn.
	sc_pc         uint64
	sc_regs       [32]uint64
	sc_flags      uint32
	_             uint32
	sc_extcontext [0]uint64
}

// The kernel reserves 128 bytes for a future expanded signal mask and aligns
// the following sigcontext to 16 bytes.
type ucontext struct {
	uc_flags    uint32
	uc_link     *ucontext
	uc_stack    stackt
	uc_sigmask  [2]uint32
	uc_unused   [120]byte
	uc_align    [12]byte
	uc_mcontext sigcontext
}

type sockaddr_un struct {
	family uint16
	path   [108]byte
}
