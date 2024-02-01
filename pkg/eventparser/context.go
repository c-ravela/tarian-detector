// SPDX-License-Identifier: Apache-2.0
// Copyright 2024 Authors of Tarian & the Organization created Tarian

package eventparser

type TarianMetaData struct {
	MetaData struct {
		Event     int32
		Nparams   uint8
		Syscall   int32
		Ts        uint64
		Processor uint16
		Task      struct {
			StartTime    uint64
			HostPid      uint32
			HostTgid     uint32
			HostPpid     uint32
			Pid          uint32
			Tgid         uint32
			Ppid         uint32
			Uid          uint32
			Gid          uint32
			CgroupId     uint64
			MountNsId    uint64
			PidNsId      uint64
			ExecId       uint64
			ParentExecId uint64
			Comm         [16]uint8
			Cwd          [256]uint8
		}
	}
	SystemInfo struct {
		Sysname    [65]uint8
		Nodename   [65]uint8
		Release    [65]uint8
		Version    [65]uint8
		Machine    [65]uint8
		Domainname [65]uint8
	}
}

type TarianParamType uint32

const (
	TDT_NONE      TarianParamType = 0
	TDT_U8        TarianParamType = 1
	TDT_U16       TarianParamType = 2
	TDT_U32       TarianParamType = 3
	TDT_U64       TarianParamType = 4
	TDT_S8        TarianParamType = 5
	TDT_S16       TarianParamType = 6
	TDT_S32       TarianParamType = 7
	TDT_S64       TarianParamType = 8
	TDT_IPV6      TarianParamType = 9
	TDT_STR       TarianParamType = 10
	TDT_STR_ARR   TarianParamType = 11
	TDT_BYTE_ARR  TarianParamType = 12
	TDT_IOVEC_ARR TarianParamType = 15
	TDT_SOCKADDR  TarianParamType = 14
)

type TarianEventsE int

const (
	TDE_SYSCALL_EXECVE_E = 2
	TDE_SYSCALL_EXECVE_R = 3

	TDE_SYSCALL_EXECVEAT_E = 4
	TDE_SYSCALL_EXECVEAT_R = 5

	TDE_SYSCALL_CLONE_E = 6
	TDE_SYSCALL_CLONE_R = 7

	TDE_SYSCALL_CLOSE_E = 8
	TDE_SYSCALL_CLOSE_R = 9

	TDE_SYSCALL_READ_E = 10
	TDE_SYSCALL_READ_R = 11

	TDE_SYSCALL_WRITE_E = 12
	TDE_SYSCALL_WRITE_R = 13

	TDE_SYSCALL_OPEN_E = 14
	TDE_SYSCALL_OPEN_R = 15

	TDE_SYSCALL_READV_E = 16
	TDE_SYSCALL_READV_R = 17

	TDE_SYSCALL_WRITEV_E = 18
	TDE_SYSCALL_WRITEV_R = 19

	TDE_SYSCALL_OPENAT_E = 20
	TDE_SYSCALL_OPENAT_R = 21

	TDE_SYSCALL_OPENAT2_E = 22
	TDE_SYSCALL_OPENAT2_R = 23

	TDE_SYSCALL_LISTEN_E = 24
	TDE_SYSCALL_LISTEN_R = 25

	TDE_SYSCALL_SOCKET_E = 26
	TDE_SYSCALL_SOCKET_R = 27

	TDE_SYSCALL_ACCEPT_E = 28
	TDE_SYSCALL_ACCEPT_R = 29

	TDE_SYSCALL_BIND_E = 30
	TDE_SYSCALL_BIND_R = 31

	TDE_SYSCALL_CONNECT_E = 32
	TDE_SYSCALL_CONNECT_R = 33
)
