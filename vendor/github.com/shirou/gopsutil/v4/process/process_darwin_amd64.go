// SPDX-License-Identifier: BSD-3-Clause
// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs types_darwin.go

package process

const (
	sizeofPtr      = 0x8
	sizeofShort    = 0x2
	sizeofInt      = 0x4
	sizeofLong     = 0x8
	sizeofLongLong = 0x8
)

type (
	_C_short     int16
	_C_int       int32
	_C_long      int64
	_C_long_long int64
)

type Timespec struct {
	Sec  int64
	Nsec int64
}

type Timeval struct {
	Sec       int64
	Usec      int32
	Pad_cgo_0 [4]byte
}

type Rusage struct {
	Utime    Timeval
	Stime    Timeval
	Maxrss   int64
	Ixrss    int64
	Idrss    int64
	Isrss    int64
	Minflt   int64
	Majflt   int64
	Nswap    int64
	Inblock  int64
	Oublock  int64
	Msgsnd   int64
	Msgrcv   int64
	Nsignals int64
	Nvcsw    int64
	Nivcsw   int64
}

type Rlimit struct {
	Cur uint64
	Max uint64
}

type UGid_t uint32

type KinfoProc struct {
	Proc  ExternProc
	Eproc Eproc
}

type Eproc struct {
	Paddr     *uint64
	Sess      *Session
	Pcred     Upcred
	Ucred     Uucred
	Pad_cgo_0 [4]byte
	Vm        Vmspace
	Ppid      int32
	Pgid      int32
	Jobc      int16
	Pad_cgo_1 [2]byte
	Tdev      int32
	Tpgid     int32
	Pad_cgo_2 [4]byte
	Tsess     *Session
	Wmesg     [8]int8
	Xsize     int32
	Xrssize   int16
	Xccount   int16
	Xswrss    int16
	Pad_cgo_3 [2]byte
	Flag      int32
	Login     [12]int8
	Spare     [4]int32
	Pad_cgo_4 [4]byte
}

type Proc struct{}

type Session struct{}

type ucred struct {
	Link  _Ctype_struct___0
	Ref   uint64
	Posix Posix_cred
	Label *Label
	Audit Au_session
}

type Uucred struct {
	Ref       int32
	UID       uint32
	Ngroups   int16
	Pad_cgo_0 [2]byte
	Groups    [16]uint32
}

type Upcred struct {
	Pc_lock   [72]int8
	Pc_ucred  *ucred
	P_ruid    uint32
	P_svuid   uint32
	P_rgid    uint32
	P_svgid   uint32
	P_refcnt  int32
	Pad_cgo_0 [4]byte
}

type Vmspace struct {
	Dummy     int32
	Pad_cgo_0 [4]byte
	Dummy2    *int8
	Dummy3    [5]int32
	Pad_cgo_1 [4]byte
	Dummy4    [3]*int8
}

type Sigacts struct{}

type ExternProc struct {
	P_un        [16]byte
	P_vmspace   uint64
	P_sigacts   uint64
	Pad_cgo_0   [3]byte
	P_flag      int32
	P_stat      int8
	P_pid       int32
	P_oppid     int32
	P_dupfd     int32
	Pad_cgo_1   [4]byte
	User_stack  uint64
	Exit_thread uint64
	P_debugger  int32
	Sigwait     int32
	P_estcpu    uint32
	P_cpticks   int32
	P_pctcpu    uint32
	Pad_cgo_2   [4]byte
	P_wchan     uint64
	P_wmesg     uint64
	P_swtime    uint32
	P_slptime   uint32
	P_realtimer Itimerval
	P_rtime     Timeval
	P_uticks    uint64
	P_sticks    uint64
	P_iticks    uint64
	P_traceflag int32
	Pad_cgo_3   [4]byte
	P_tracep    uint64
	P_siglist   int32
	Pad_cgo_4   [4]byte
	P_textvp    uint64
	P_holdcnt   int32
	P_sigmask   uint32
	P_sigignore uint32
	P_sigcatch  uint32
	P_priority  uint8
	P_usrpri    uint8
	P_nice      int8
	P_comm      [17]int8
	Pad_cgo_5   [4]byte
	P_pgrp      uint64
	P_addr      uint64
	P_xstat     uint16
	P_acflag    uint16
	Pad_cgo_6   [4]byte
	P_ru        uint64
}

type Itimerval struct {
	Interval Timeval
	Value    Timeval
}

type Vnode struct{}

type Pgrp struct{}

type UserStruct struct{}

type Au_session struct {
	Aia_p *AuditinfoAddr
	Mask  AuMask
}

type Posix_cred struct {
	UID       uint32
	Ruid      uint32
	Svuid     uint32
	Ngroups   int16
	Pad_cgo_0 [2]byte
	Groups    [16]uint32
	Rgid      uint32
	Svgid     uint32
	Gmuid     uint32
	Flags     int32
}

type Label struct{}

type ProcTaskInfo struct {
	Virtual_size      uint64
	Resident_size     uint64
	Total_user        uint64
	Total_system      uint64
	Threads_user      uint64
	Threads_system    uint64
	Policy            int32
	Faults            int32
	Pageins           int32
	Cow_faults        int32
	Messages_sent     int32
	Messages_received int32
	Syscalls_mach     int32
	Syscalls_unix     int32
	Csw               int32
	Threadnum         int32
	Numrunning        int32
	Priority          int32
}

type AuditinfoAddr struct {
	Auid   uint32
	Mask   AuMask
	Termid AuTidAddr
	Asid   int32
	Flags  uint64
}

type AuMask struct {
	Success uint32
	Failure uint32
}

type AuTidAddr struct {
	Port int32
	Type uint32
	Addr [4]uint32
}

type UcredQueue struct {
	Next *ucred
	Prev **ucred
}
