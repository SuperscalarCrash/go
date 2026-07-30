// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && loong32r

package syscall_test

import (
	"os"
	"path/filepath"
	"syscall"
	"testing"
	"unsafe"
)

func TestLoong32rEpollEvent(t *testing.T) {
	var event syscall.EpollEvent
	if got, want := unsafe.Sizeof(event), uintptr(16); got != want {
		t.Fatalf("sizeof(EpollEvent) = %d, want %d", got, want)
	}
	if got, want := unsafe.Offsetof(event.Fd), uintptr(8); got != want {
		t.Fatalf("offsetof(EpollEvent.Fd) = %d, want %d", got, want)
	}

	epfd, err := syscall.EpollCreate1(syscall.EPOLL_CLOEXEC)
	if err != nil {
		t.Fatal(err)
	}
	defer syscall.Close(epfd)

	var pipefd [2]int
	if err := syscall.Pipe2(pipefd[:], syscall.O_CLOEXEC); err != nil {
		t.Fatal(err)
	}
	defer syscall.Close(pipefd[0])
	defer syscall.Close(pipefd[1])

	const tag = int32(0x12345678)
	event = syscall.EpollEvent{Events: syscall.EPOLLIN, Fd: tag}
	if err := syscall.EpollCtl(epfd, syscall.EPOLL_CTL_ADD, pipefd[0], &event); err != nil {
		t.Fatal(err)
	}
	if _, err := syscall.Write(pipefd[1], []byte{1}); err != nil {
		t.Fatal(err)
	}

	events := make([]syscall.EpollEvent, 1)
	n, err := syscall.EpollWait(epfd, events, 1000)
	if err != nil {
		t.Fatal(err)
	}
	if n != 1 || events[0].Events&syscall.EPOLLIN == 0 || events[0].Fd != tag {
		t.Fatalf("EpollWait = %d, %#v; want one EPOLLIN event tagged %#x", n, events[0], tag)
	}
}

func TestLoong32rPtraceRegsLayout(t *testing.T) {
	var regs syscall.PtraceRegs
	if got, want := unsafe.Sizeof(regs), uintptr(184); got != want {
		t.Fatalf("sizeof(PtraceRegs) = %d, want %d", got, want)
	}
	if got, want := unsafe.Offsetof(regs.Era), uintptr(132); got != want {
		t.Fatalf("offsetof(PtraceRegs.Era) = %d, want %d", got, want)
	}
	if got, want := unsafe.Offsetof(regs.Badv), uintptr(136); got != want {
		t.Fatalf("offsetof(PtraceRegs.Badv) = %d, want %d", got, want)
	}
}

func TestLoong32rLargeFileSyscalls(t *testing.T) {
	path := filepath.Join(t.TempDir(), "large-file")
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	const offset = int64(5)<<30 + 321
	want := []byte("loong32r-large-offset")
	if n, err := syscall.Pwrite(int(f.Fd()), want, offset); err != nil || n != len(want) {
		t.Fatalf("Pwrite offset %#x = %d, %v", offset, n, err)
	}
	got := make([]byte, len(want))
	if n, err := syscall.Pread(int(f.Fd()), got, offset); err != nil || n != len(want) {
		t.Fatalf("Pread offset %#x = %d, %v", offset, n, err)
	}
	if string(got) != string(want) {
		t.Fatalf("Pread offset %#x = %q, want %q", offset, got, want)
	}

	const ftruncateSize = int64(6)<<30 + 17
	if err := syscall.Ftruncate(int(f.Fd()), ftruncateSize); err != nil {
		t.Fatalf("Ftruncate(%#x): %v", ftruncateSize, err)
	}
	if info, err := f.Stat(); err != nil || info.Size() != ftruncateSize {
		t.Fatalf("size after Ftruncate = %v, %v; want %d", func() int64 {
			if info != nil {
				return info.Size()
			}
			return -1
		}(), err, ftruncateSize)
	}

	const truncateSize = int64(7)<<30 + 19
	if err := syscall.Truncate(path, truncateSize); err != nil {
		t.Fatalf("Truncate(%#x): %v", truncateSize, err)
	}
	if info, err := os.Stat(path); err != nil || info.Size() != truncateSize {
		t.Fatalf("size after Truncate = %v, %v; want %d", func() int64 {
			if info != nil {
				return info.Size()
			}
			return -1
		}(), err, truncateSize)
	}

	const seekOffset = int64(5)<<30 + 0x10203
	if got, err := syscall.Seek(int(f.Fd()), seekOffset, 0); err != nil || got != seekOffset {
		t.Fatalf("Seek(%#x) = %#x, %v", seekOffset, got, err)
	}
	seekData := []byte("loong32r-llseek")
	if n, err := syscall.Write(int(f.Fd()), seekData); err != nil || n != len(seekData) {
		t.Fatalf("Write after Seek(%#x) = %d, %v", seekOffset, n, err)
	}
	seekGot := make([]byte, len(seekData))
	if n, err := syscall.Pread(int(f.Fd()), seekGot, seekOffset); err != nil || n != len(seekGot) {
		t.Fatalf("Pread after Seek(%#x) = %d, %v", seekOffset, n, err)
	}
	if string(seekGot) != string(seekData) {
		t.Fatalf("data after Seek(%#x) = %q, want %q", seekOffset, seekGot, seekData)
	}

	// The LA32R sync_file_range syscall packs both 64-bit arguments into
	// pairs of 32-bit syscall registers. Use an offset whose high half is
	// non-zero so a register-order mistake is visible to the kernel.
	const syncFileRangeWrite = 0x2
	const syncOffset = int64(5)<<30 + 0x12345000
	if err := syscall.SyncFileRange(int(f.Fd()), syncOffset, 4096, syncFileRangeWrite); err != nil {
		t.Fatalf("SyncFileRange(%#x): %v", syncOffset, err)
	}

	// Extending an empty file with a 4 KiB allocation at a large offset
	// validates fallocate's split offset while consuming only one block.
	t.Run("Fallocate", func(t *testing.T) {
		fallocatePath := filepath.Join(t.TempDir(), "fallocate-large-offset")
		fallocateFile, err := os.OpenFile(fallocatePath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
		if err != nil {
			t.Fatal(err)
		}
		defer fallocateFile.Close()

		const (
			fallocateOffset = int64(5)<<30 + 0x20000
			fallocateLength = int64(4096)
		)
		if err := syscall.Fallocate(int(fallocateFile.Fd()), 0, fallocateOffset, fallocateLength); err != nil {
			if err == syscall.ENOSYS || err == syscall.EOPNOTSUPP {
				t.Skipf("filesystem does not support fallocate: %v", err)
			}
			t.Fatalf("Fallocate(%#x, %#x): %v", fallocateOffset, fallocateLength, err)
		}
		info, err := fallocateFile.Stat()
		if err != nil {
			t.Fatal(err)
		}
		if want := fallocateOffset + fallocateLength; info.Size() != want {
			t.Fatalf("size after Fallocate = %#x, want %#x", info.Size(), want)
		}
	})
}

func TestLoong32rSpliceResultABI(t *testing.T) {
	var pipefd [2]int
	if err := syscall.Pipe2(pipefd[:], syscall.O_CLOEXEC); err != nil {
		t.Fatal(err)
	}
	defer syscall.Close(pipefd[0])
	defer syscall.Close(pipefd[1])

	want := []byte("loong32r-splice")
	if n, err := syscall.Write(pipefd[1], want); err != nil || n != len(want) {
		t.Fatalf("Write = %d, %v; want %d", n, err, len(want))
	}

	f, err := os.Create(filepath.Join(t.TempDir(), "splice-output"))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	// Linux splice returns ssize_t, which is int-sized in the ILP32 ABI.
	// Keep an explicit int assignment here so this test also checks the Go
	// wrapper signature rather than only its runtime value.
	var n int
	n, err = syscall.Splice(pipefd[0], nil, int(f.Fd()), nil, len(want), 0)
	if err != nil || n != len(want) {
		t.Fatalf("Splice = %d, %v; want %d", n, err, len(want))
	}
	if _, err := f.Seek(0, 0); err != nil {
		t.Fatal(err)
	}
	got := make([]byte, len(want))
	if _, err := f.Read(got); err != nil {
		t.Fatal(err)
	}
	if string(got) != string(want) {
		t.Fatalf("spliced data = %q, want %q", got, want)
	}
}
