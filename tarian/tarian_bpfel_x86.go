// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64

package tarian

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type tarianEventDataT struct {
	Context struct {
		Ts   uint64
		Task struct {
			StartTime     uint64
			HostPid       uint32
			HostTgid      uint32
			HostPpid      uint32
			Pid           uint32
			Tgid          uint32
			Ppid          uint32
			Uid           uint32
			Gid           uint32
			CgroupId      uint64
			MountNsId     uint64
			PidNsId       uint64
			ExecId        uint64
			ParentExecId  uint64
			EexecId       uint64
			EparentExecId uint64
			Comm          [16]uint8
			Cwd           [4096]uint8
		}
		EventId     uint32
		Syscall     int32
		ProcessorId uint16
	}
	Buf struct {
		NumFields  uint8
		FieldTypes uint64
		Data       [10240]uint8
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

// loadTarian returns the embedded CollectionSpec for tarian.
func loadTarian() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_TarianBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load tarian: %w", err)
	}

	return spec, err
}

// loadTarianObjects loads tarian and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*tarianObjects
//	*tarianPrograms
//	*tarianMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadTarianObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadTarian()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// tarianSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type tarianSpecs struct {
	tarianProgramSpecs
	tarianMapSpecs
}

// tarianSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type tarianProgramSpecs struct {
	KprobeAccept      *ebpf.ProgramSpec `ebpf:"kprobe_accept"`
	KprobeBind        *ebpf.ProgramSpec `ebpf:"kprobe_bind"`
	KprobeClone       *ebpf.ProgramSpec `ebpf:"kprobe_clone"`
	KprobeClose       *ebpf.ProgramSpec `ebpf:"kprobe_close"`
	KprobeConnect     *ebpf.ProgramSpec `ebpf:"kprobe_connect"`
	KprobeExecve      *ebpf.ProgramSpec `ebpf:"kprobe_execve"`
	KprobeExecveat    *ebpf.ProgramSpec `ebpf:"kprobe_execveat"`
	KprobeListen      *ebpf.ProgramSpec `ebpf:"kprobe_listen"`
	KprobeOpen        *ebpf.ProgramSpec `ebpf:"kprobe_open"`
	KprobeOpenat      *ebpf.ProgramSpec `ebpf:"kprobe_openat"`
	KprobeOpenat2     *ebpf.ProgramSpec `ebpf:"kprobe_openat2"`
	KprobeRead        *ebpf.ProgramSpec `ebpf:"kprobe_read"`
	KprobeReadv       *ebpf.ProgramSpec `ebpf:"kprobe_readv"`
	KprobeSocket      *ebpf.ProgramSpec `ebpf:"kprobe_socket"`
	KprobeWrite       *ebpf.ProgramSpec `ebpf:"kprobe_write"`
	KprobeWritev      *ebpf.ProgramSpec `ebpf:"kprobe_writev"`
	KretprobeAccept   *ebpf.ProgramSpec `ebpf:"kretprobe_accept"`
	KretprobeBind     *ebpf.ProgramSpec `ebpf:"kretprobe_bind"`
	KretprobeClone    *ebpf.ProgramSpec `ebpf:"kretprobe_clone"`
	KretprobeClose    *ebpf.ProgramSpec `ebpf:"kretprobe_close"`
	KretprobeConnect  *ebpf.ProgramSpec `ebpf:"kretprobe_connect"`
	KretprobeExecve   *ebpf.ProgramSpec `ebpf:"kretprobe_execve"`
	KretprobeExecveat *ebpf.ProgramSpec `ebpf:"kretprobe_execveat"`
	KretprobeListen   *ebpf.ProgramSpec `ebpf:"kretprobe_listen"`
	KretprobeOpen     *ebpf.ProgramSpec `ebpf:"kretprobe_open"`
	KretprobeOpenat   *ebpf.ProgramSpec `ebpf:"kretprobe_openat"`
	KretprobeOpenat2  *ebpf.ProgramSpec `ebpf:"kretprobe_openat2"`
	KretprobeRead     *ebpf.ProgramSpec `ebpf:"kretprobe_read"`
	KretprobeReadv    *ebpf.ProgramSpec `ebpf:"kretprobe_readv"`
	KretprobeSocket   *ebpf.ProgramSpec `ebpf:"kretprobe_socket"`
	KretprobeWrite    *ebpf.ProgramSpec `ebpf:"kretprobe_write"`
	KretprobeWritev   *ebpf.ProgramSpec `ebpf:"kretprobe_writev"`
}

// tarianMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type tarianMapSpecs struct {
	Events   *ebpf.MapSpec `ebpf:"events"`
	PercpuRb *ebpf.MapSpec `ebpf:"percpu_rb"`
	RbCpu0   *ebpf.MapSpec `ebpf:"rb_cpu0"`
	RbCpu1   *ebpf.MapSpec `ebpf:"rb_cpu1"`
	RbCpu10  *ebpf.MapSpec `ebpf:"rb_cpu10"`
	RbCpu11  *ebpf.MapSpec `ebpf:"rb_cpu11"`
	RbCpu12  *ebpf.MapSpec `ebpf:"rb_cpu12"`
	RbCpu13  *ebpf.MapSpec `ebpf:"rb_cpu13"`
	RbCpu14  *ebpf.MapSpec `ebpf:"rb_cpu14"`
	RbCpu15  *ebpf.MapSpec `ebpf:"rb_cpu15"`
	RbCpu2   *ebpf.MapSpec `ebpf:"rb_cpu2"`
	RbCpu3   *ebpf.MapSpec `ebpf:"rb_cpu3"`
	RbCpu4   *ebpf.MapSpec `ebpf:"rb_cpu4"`
	RbCpu5   *ebpf.MapSpec `ebpf:"rb_cpu5"`
	RbCpu6   *ebpf.MapSpec `ebpf:"rb_cpu6"`
	RbCpu7   *ebpf.MapSpec `ebpf:"rb_cpu7"`
	RbCpu8   *ebpf.MapSpec `ebpf:"rb_cpu8"`
	RbCpu9   *ebpf.MapSpec `ebpf:"rb_cpu9"`
}

// tarianObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadTarianObjects or ebpf.CollectionSpec.LoadAndAssign.
type tarianObjects struct {
	tarianPrograms
	tarianMaps
}

func (o *tarianObjects) Close() error {
	return _TarianClose(
		&o.tarianPrograms,
		&o.tarianMaps,
	)
}

// tarianMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadTarianObjects or ebpf.CollectionSpec.LoadAndAssign.
type tarianMaps struct {
	Events   *ebpf.Map `ebpf:"events"`
	PercpuRb *ebpf.Map `ebpf:"percpu_rb"`
	RbCpu0   *ebpf.Map `ebpf:"rb_cpu0"`
	RbCpu1   *ebpf.Map `ebpf:"rb_cpu1"`
	RbCpu10  *ebpf.Map `ebpf:"rb_cpu10"`
	RbCpu11  *ebpf.Map `ebpf:"rb_cpu11"`
	RbCpu12  *ebpf.Map `ebpf:"rb_cpu12"`
	RbCpu13  *ebpf.Map `ebpf:"rb_cpu13"`
	RbCpu14  *ebpf.Map `ebpf:"rb_cpu14"`
	RbCpu15  *ebpf.Map `ebpf:"rb_cpu15"`
	RbCpu2   *ebpf.Map `ebpf:"rb_cpu2"`
	RbCpu3   *ebpf.Map `ebpf:"rb_cpu3"`
	RbCpu4   *ebpf.Map `ebpf:"rb_cpu4"`
	RbCpu5   *ebpf.Map `ebpf:"rb_cpu5"`
	RbCpu6   *ebpf.Map `ebpf:"rb_cpu6"`
	RbCpu7   *ebpf.Map `ebpf:"rb_cpu7"`
	RbCpu8   *ebpf.Map `ebpf:"rb_cpu8"`
	RbCpu9   *ebpf.Map `ebpf:"rb_cpu9"`
}

func (m *tarianMaps) Close() error {
	return _TarianClose(
		m.Events,
		m.PercpuRb,
		m.RbCpu0,
		m.RbCpu1,
		m.RbCpu10,
		m.RbCpu11,
		m.RbCpu12,
		m.RbCpu13,
		m.RbCpu14,
		m.RbCpu15,
		m.RbCpu2,
		m.RbCpu3,
		m.RbCpu4,
		m.RbCpu5,
		m.RbCpu6,
		m.RbCpu7,
		m.RbCpu8,
		m.RbCpu9,
	)
}

// tarianPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadTarianObjects or ebpf.CollectionSpec.LoadAndAssign.
type tarianPrograms struct {
	KprobeAccept      *ebpf.Program `ebpf:"kprobe_accept"`
	KprobeBind        *ebpf.Program `ebpf:"kprobe_bind"`
	KprobeClone       *ebpf.Program `ebpf:"kprobe_clone"`
	KprobeClose       *ebpf.Program `ebpf:"kprobe_close"`
	KprobeConnect     *ebpf.Program `ebpf:"kprobe_connect"`
	KprobeExecve      *ebpf.Program `ebpf:"kprobe_execve"`
	KprobeExecveat    *ebpf.Program `ebpf:"kprobe_execveat"`
	KprobeListen      *ebpf.Program `ebpf:"kprobe_listen"`
	KprobeOpen        *ebpf.Program `ebpf:"kprobe_open"`
	KprobeOpenat      *ebpf.Program `ebpf:"kprobe_openat"`
	KprobeOpenat2     *ebpf.Program `ebpf:"kprobe_openat2"`
	KprobeRead        *ebpf.Program `ebpf:"kprobe_read"`
	KprobeReadv       *ebpf.Program `ebpf:"kprobe_readv"`
	KprobeSocket      *ebpf.Program `ebpf:"kprobe_socket"`
	KprobeWrite       *ebpf.Program `ebpf:"kprobe_write"`
	KprobeWritev      *ebpf.Program `ebpf:"kprobe_writev"`
	KretprobeAccept   *ebpf.Program `ebpf:"kretprobe_accept"`
	KretprobeBind     *ebpf.Program `ebpf:"kretprobe_bind"`
	KretprobeClone    *ebpf.Program `ebpf:"kretprobe_clone"`
	KretprobeClose    *ebpf.Program `ebpf:"kretprobe_close"`
	KretprobeConnect  *ebpf.Program `ebpf:"kretprobe_connect"`
	KretprobeExecve   *ebpf.Program `ebpf:"kretprobe_execve"`
	KretprobeExecveat *ebpf.Program `ebpf:"kretprobe_execveat"`
	KretprobeListen   *ebpf.Program `ebpf:"kretprobe_listen"`
	KretprobeOpen     *ebpf.Program `ebpf:"kretprobe_open"`
	KretprobeOpenat   *ebpf.Program `ebpf:"kretprobe_openat"`
	KretprobeOpenat2  *ebpf.Program `ebpf:"kretprobe_openat2"`
	KretprobeRead     *ebpf.Program `ebpf:"kretprobe_read"`
	KretprobeReadv    *ebpf.Program `ebpf:"kretprobe_readv"`
	KretprobeSocket   *ebpf.Program `ebpf:"kretprobe_socket"`
	KretprobeWrite    *ebpf.Program `ebpf:"kretprobe_write"`
	KretprobeWritev   *ebpf.Program `ebpf:"kretprobe_writev"`
}

func (p *tarianPrograms) Close() error {
	return _TarianClose(
		p.KprobeAccept,
		p.KprobeBind,
		p.KprobeClone,
		p.KprobeClose,
		p.KprobeConnect,
		p.KprobeExecve,
		p.KprobeExecveat,
		p.KprobeListen,
		p.KprobeOpen,
		p.KprobeOpenat,
		p.KprobeOpenat2,
		p.KprobeRead,
		p.KprobeReadv,
		p.KprobeSocket,
		p.KprobeWrite,
		p.KprobeWritev,
		p.KretprobeAccept,
		p.KretprobeBind,
		p.KretprobeClone,
		p.KretprobeClose,
		p.KretprobeConnect,
		p.KretprobeExecve,
		p.KretprobeExecveat,
		p.KretprobeListen,
		p.KretprobeOpen,
		p.KretprobeOpenat,
		p.KretprobeOpenat2,
		p.KretprobeRead,
		p.KretprobeReadv,
		p.KretprobeSocket,
		p.KretprobeWrite,
		p.KretprobeWritev,
	)
}

func _TarianClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed tarian_bpfel_x86.o
var _TarianBytes []byte