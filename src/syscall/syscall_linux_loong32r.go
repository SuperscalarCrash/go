// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

import "unsafe"

const (
	_SYS_setgroups  = SYS_SETGROUPS
	_SYS_clone3     = 435
	_SYS_faccessat2 = 439
	_SYS_fchmodat2  = 452

	// Linux maps the asm-generic 3264 syscall slots to these LA32R
	// interfaces. The aliases are part of the public syscall constants.
	SYS_STATFS64    = 43
	SYS_FSTATFS64   = 44
	SYS_TRUNCATE64  = 45
	SYS_FTRUNCATE64 = 46
	SYS__LLSEEK     = 62
	SYS_SENDFILE64  = 71
	SYS_FSTATAT64   = 79
	SYS_FSTAT64     = 80
	SYS_MMAP2       = 222
)

//sys EpollWait(epfd int, events []EpollEvent, msec int) (n int, err error) = SYS_EPOLL_PWAIT
//sys Fchown(fd int, uid int, gid int) (err error)
//sys Fstat(fd int, stat *Stat_t) (err error) = SYS_FSTAT64
//sys fstatat(dirfd int, path string, stat *Stat_t, flags int) (err error) = SYS_FSTATAT64
//sys Ftruncate(fd int, length int64) (err error) = SYS_FTRUNCATE64
//sysnb Getegid() (egid int)
//sysnb Geteuid() (euid int)
//sysnb Getgid() (gid int)
//sysnb Getuid() (uid int)
//sys Listen(s int, n int) (err error)
//sys pread(fd int, p []byte, offset int64) (n int, err error) = SYS_PREAD64
//sys pwrite(fd int, p []byte, offset int64) (n int, err error) = SYS_PWRITE64
//sys renameat2(olddirfd int, oldpath string, newdirfd int, newpath string, flags uint) (err error) = SYS_RENAMEAT2
//sys sendfile(outfd int, infd int, offset *int64, count int) (written int, err error) = SYS_SENDFILE64
//sys Setfsgid(gid int) (err error)
//sys Setfsuid(uid int) (err error)
//sys Shutdown(fd int, how int) (err error)
//sys Splice(rfd int, roff *int64, wfd int, woff *int64, len int, flags int) (n int, err error)
//sys Truncate(path string, length int64) (err error) = SYS_TRUNCATE64
//sys accept4(s int, rsa *RawSockaddrAny, addrlen *_Socklen, flags int) (fd int, err error)
//sys bind(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
//sys connect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
//sysnb getgroups(n int, list *_Gid_t) (nn int, err error)
//sys getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error)
//sys setsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error)
//sysnb socket(domain int, typ int, proto int) (fd int, err error)
//sysnb socketpair(domain int, typ int, proto int, fd *[2]int32) (err error)
//sysnb getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sysnb getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sys recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error)
//sys sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error)
//sys recvmsg(s int, msg *Msghdr, flags int) (n int, err error)
//sys sendmsg(s int, msg *Msghdr, flags int) (n int, err error)

//sysnb Gettimeofday(tv *Timeval) (err error)

//sys mmap2(addr uintptr, length uintptr, prot int, flags int, fd int, pageOffset uintptr) (xaddr uintptr, err error) = SYS_MMAP2

func Fstatat(fd int, path string, stat *Stat_t, flags int) error {
	return fstatat(fd, path, stat, flags)
}

func Stat(path string, stat *Stat_t) error {
	return fstatat(_AT_FDCWD, path, stat, 0)
}

func Lstat(path string, stat *Stat_t) error {
	return fstatat(_AT_FDCWD, path, stat, _AT_SYMLINK_NOFOLLOW)
}

func Lchown(path string, uid, gid int) error {
	return Fchownat(_AT_FDCWD, path, uid, gid, _AT_SYMLINK_NOFOLLOW)
}

func Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) error {
	return renameat2(olddirfd, oldpath, newdirfd, newpath, 0)
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
	_, _, errno := Syscall6(SYS__LLSEEK, uintptr(fd), uintptr(offset>>32), uintptr(offset), uintptr(unsafe.Pointer(&off)), uintptr(whence), 0)
	if errno != 0 {
		err = errnoErr(errno)
	}
	return
}

func SyncFileRange(fd int, off int64, n int64, flags int) error {
	_, _, errno := Syscall6(SYS_SYNC_FILE_RANGE,
		uintptr(fd), uintptr(off), uintptr(off>>32),
		uintptr(n), uintptr(n>>32), uintptr(flags))
	if errno != 0 {
		return errnoErr(errno)
	}
	return nil
}

func mmap(addr, length uintptr, prot, flags, fd int, offset int64) (uintptr, error) {
	pageOffset := uintptr(offset / 4096)
	if offset != int64(pageOffset)*4096 {
		return 0, EINVAL
	}
	return mmap2(addr, length, prot, flags, fd, pageOffset)
}

type sigset_t struct {
	X__val [32]uint32
}

//sys pselect(nfd int, r *FdSet, w *FdSet, e *FdSet, timeout *Timespec, sigmask *sigset_t) (n int, err error) = SYS_PSELECT6

func Select(nfd int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (n int, err error) {
	var ts *Timespec
	if timeout != nil {
		ts = &Timespec{Sec: timeout.Sec, Nsec: timeout.Usec * 1000}
	}
	return pselect(nfd, r, w, e, ts, nil)
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
	tv := []Timeval{{Sec: buf.Actime}, {Sec: buf.Modtime}}
	return Utimes(path, tv)
}

func utimes(path string, tv *[2]Timeval) error {
	if tv == nil {
		return utimensat(_AT_FDCWD, path, nil, 0)
	}
	ts := [2]Timespec{
		NsecToTimespec(TimevalToNsec(tv[0])),
		NsecToTimespec(TimevalToNsec(tv[1])),
	}
	return utimensat(_AT_FDCWD, path, &ts, 0)
}

func (r *PtraceRegs) PC() uint64        { return uint64(r.Era) }
func (r *PtraceRegs) SetPC(pc uint64)   { r.Era = uint32(pc) }
func (r *PtraceRegs) GetEra() uint64    { return uint64(r.Era) }
func (r *PtraceRegs) SetEra(era uint64) { r.Era = uint32(era) }

func (iov *Iovec) SetLen(length int)         { iov.Len = uint32(length) }
func (msg *Msghdr) SetControllen(length int) { msg.Controllen = uint32(length) }
func (msg *Cmsghdr) SetLen(length int)       { msg.Len = uint32(length) }
func InotifyInit() (fd int, err error)       { return InotifyInit1(0) }

//sys ppoll(fds *pollFd, nfds int, timeout *Timespec, sigmask *sigset_t) (n int, err error) = SYS_PPOLL

func Pause() error {
	_, err := ppoll(nil, 0, nil, nil)
	return err
}
