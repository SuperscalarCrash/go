// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build loong32r && linux

package unix

import "unsafe"

// LA32R uses the 32-bit asm-generic Linux syscall ABI. The asm-generic 3264
// syscall slots select the large-file entry points, and 64-bit scalar
// arguments occupy consecutive little-endian register pairs.

// The LA32R kernel exposes only the large-file fcntl entry at asm-generic slot
// 25. Keep the ordinary name as an API alias; fcntl_linux_32bit.go selects the
// matching Flock_t entry explicitly.
const SYS_FCNTL = SYS_FCNTL64

//sys	EpollWait(epfd int, events []EpollEvent, msec int) (n int, err error) = SYS_EPOLL_PWAIT
//sys	Fadvise(fd int, offset int64, length int64, advice int) (err error) = SYS_FADVISE64_64
//sys	Fchown(fd int, uid int, gid int) (err error)
//sys	Fstat(fd int, stat *Stat_t) (err error) = SYS_FSTAT64
//sys	Fstatat(dirfd int, path string, stat *Stat_t, flags int) (err error) = SYS_FSTATAT64
//sys	Ftruncate(fd int, length int64) (err error) = SYS_FTRUNCATE64
//sysnb	Getegid() (egid int)
//sysnb	Geteuid() (euid int)
//sysnb	Getgid() (gid int)
//sysnb	Getuid() (uid int)
//sys	Listen(s int, n int) (err error)
//sys	MemfdSecret(flags int) (fd int, err error)
//sys	pread(fd int, p []byte, offset int64) (n int, err error) = SYS_PREAD64
//sys	pwrite(fd int, p []byte, offset int64) (n int, err error) = SYS_PWRITE64
//sys	sendfile(outfd int, infd int, offset *int64, count int) (written int, err error) = SYS_SENDFILE64
//sys	setfsgid(gid int) (prev int, err error)
//sys	setfsuid(uid int) (prev int, err error)
//sys	Shutdown(fd int, how int) (err error)
//sys	Splice(rfd int, roff *int64, wfd int, woff *int64, len int, flags int) (n int, err error)
//sys	SyncFileRange(fd int, off int64, n int64, flags int) (err error)
//sys	Truncate(path string, length int64) (err error) = SYS_TRUNCATE64

func Select(nfd int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (n int, err error) {
	var ts *Timespec
	if timeout != nil {
		ts = &Timespec{Sec: timeout.Sec, Nsec: timeout.Usec * 1000}
	}
	return pselect6(nfd, r, w, e, ts, nil)
}

func Stat(path string, stat *Stat_t) error {
	return Fstatat(AT_FDCWD, path, stat, 0)
}

func Lstat(path string, stat *Stat_t) error {
	return Fstatat(AT_FDCWD, path, stat, AT_SYMLINK_NOFOLLOW)
}

func Lchown(path string, uid int, gid int) error {
	return Fchownat(AT_FDCWD, path, uid, gid, AT_SYMLINK_NOFOLLOW)
}

func Fstatfs(fd int, buf *Statfs_t) (err error) {
	_, _, errno := Syscall(SYS_FSTATFS64, uintptr(fd), unsafe.Sizeof(*buf), uintptr(unsafe.Pointer(buf)))
	if errno != 0 {
		err = errnoErr(errno)
	}
	return
}

func Statfs(path string, buf *Statfs_t) (err error) {
	p, err := BytePtrFromString(path)
	if err != nil {
		return err
	}
	_, _, errno := Syscall(SYS_STATFS64, uintptr(unsafe.Pointer(p)), unsafe.Sizeof(*buf), uintptr(unsafe.Pointer(buf)))
	if errno != 0 {
		err = errnoErr(errno)
	}
	return
}

func Seek(fd int, offset int64, whence int) (off int64, err error) {
	_, _, errno := Syscall6(SYS_LLSEEK, uintptr(fd), uintptr(offset>>32), uintptr(offset), uintptr(unsafe.Pointer(&off)), uintptr(whence), 0)
	if errno != 0 {
		err = errnoErr(errno)
	}
	return
}

//sys	accept4(s int, rsa *RawSockaddrAny, addrlen *_Socklen, flags int) (fd int, err error)
//sys	bind(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
//sys	connect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
//sysnb	getgroups(n int, list *_Gid_t) (nn int, err error)
//sysnb	setgroups(n int, list *_Gid_t) (err error)
//sys	getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error)
//sys	setsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error)
//sysnb	socket(domain int, typ int, proto int) (fd int, err error)
//sysnb	socketpair(domain int, typ int, proto int, fd *[2]int32) (err error)
//sysnb	getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sysnb	getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sys	recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error)
//sys	sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error)
//sys	recvmsg(s int, msg *Msghdr, flags int) (n int, err error)
//sys	sendmsg(s int, msg *Msghdr, flags int) (n int, err error)

//sysnb	Gettimeofday(tv *Timeval) (err error)
//sys	mmap2(addr uintptr, length uintptr, prot int, flags int, fd int, pageOffset uintptr) (xaddr uintptr, err error) = SYS_MMAP2

func mmap(addr, length uintptr, prot, flags, fd int, offset int64) (uintptr, error) {
	pageOffset := uintptr(offset / 4096)
	if offset != int64(pageOffset)*4096 {
		return 0, EINVAL
	}
	return mmap2(addr, length, prot, flags, fd, pageOffset)
}

type rlimit32 struct {
	Cur uint32
	Max uint32
}

//sysnb	getrlimit(resource int, rlim *rlimit32) (err error) = SYS_GETRLIMIT

const (
	rlimInf32 = ^uint32(0)
	rlimInf64 = ^uint64(0)
)

func Getrlimit(resource int, rlim *Rlimit) error {
	err := Prlimit(0, resource, nil, rlim)
	if err != ENOSYS {
		return err
	}

	var old rlimit32
	if err := getrlimit(resource, &old); err != nil {
		return err
	}
	if old.Cur == rlimInf32 {
		rlim.Cur = rlimInf64
	} else {
		rlim.Cur = uint64(old.Cur)
	}
	if old.Max == rlimInf32 {
		rlim.Max = rlimInf64
	} else {
		rlim.Max = uint64(old.Max)
	}
	return nil
}

func setTimespec(sec, nsec int64) Timespec {
	return Timespec{Sec: int32(sec), Nsec: int32(nsec)}
}

func setTimeval(sec, usec int64) Timeval {
	return Timeval{Sec: int32(sec), Usec: int32(usec)}
}

func futimesat(dirfd int, path string, tv *[2]Timeval) error {
	if tv == nil {
		return utimensat(dirfd, path, nil, 0)
	}
	ts := [2]Timespec{
		NsecToTimespec(TimevalToNsec(tv[0])),
		NsecToTimespec(TimevalToNsec(tv[1])),
	}
	return utimensat(dirfd, path, &ts, 0)
}

func Time(t *Time_t) (Time_t, error) {
	var tv Timeval
	if err := Gettimeofday(&tv); err != nil {
		return 0, err
	}
	result := Time_t(tv.Sec)
	if t != nil {
		*t = result
	}
	return result, nil
}

func Utime(path string, buf *Utimbuf) error {
	if buf == nil {
		return Utimes(path, nil)
	}
	return Utimes(path, []Timeval{{Sec: buf.Actime}, {Sec: buf.Modtime}})
}

func utimes(path string, tv *[2]Timeval) error {
	if tv == nil {
		return utimensat(AT_FDCWD, path, nil, 0)
	}
	ts := [2]Timespec{
		NsecToTimespec(TimevalToNsec(tv[0])),
		NsecToTimespec(TimevalToNsec(tv[1])),
	}
	return utimensat(AT_FDCWD, path, &ts, 0)
}

func Ustat(dev int, ubuf *Ustat_t) error { return ENOSYS }

func (r *PtraceRegs) PC() uint64      { return uint64(r.Era) }
func (r *PtraceRegs) SetPC(pc uint64) { r.Era = uint32(pc) }

func (iov *Iovec) SetLen(length int) { iov.Len = uint32(length) }

func (msg *Msghdr) SetControllen(length int) { msg.Controllen = uint32(length) }

func (msg *Msghdr) SetIovlen(length int) { msg.Iovlen = uint32(length) }

func (cmsg *Cmsghdr) SetLen(length int) { cmsg.Len = uint32(length) }

func (rsa *RawSockaddrNFCLLCP) SetServiceNameLen(length int) {
	rsa.Service_name_len = uint32(length)
}

func Pause() error {
	_, err := ppoll(nil, 0, nil, nil)
	return err
}

func Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) error {
	return Renameat2(olddirfd, oldpath, newdirfd, newpath, 0)
}

//sys	kexecFileLoad(kernelFd int, initrdFd int, cmdlineLen int, cmdline string, flags int) (err error)

func KexecFileLoad(kernelFd int, initrdFd int, cmdline string, flags int) error {
	cmdlineLen := len(cmdline)
	if cmdlineLen > 0 {
		cmdlineLen++
	}
	return kexecFileLoad(kernelFd, initrdFd, cmdlineLen, cmdline, flags)
}
