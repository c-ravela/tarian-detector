// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64

package network_accept

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type acceptEventData struct {
	EventContext struct {
		Ts        uint64
		StartTime uint64
		Pid       uint32
		Tgid      uint32
		Ppid      uint32
		Glpid     uint32
		Uid       uint32
		Gid       uint32
		MountId   int32
		MountNsId uint32
		CgroupId  uint64
		NodeInfo  struct {
			Sysname    [65]uint8
			Nodename   [65]uint8
			Release    [65]uint8
			Version    [65]uint8
			Machine    [65]uint8
			Domainname [65]uint8
		}
		Comm         [16]uint8
		Cwd          [32]uint8
		MountDevname [256]uint8
	}
	_        [2]byte
	Id       int32
	Fd       int32
	Addrlen  int32
	SaFamily uint16
	Port     uint16
	V4Addr   struct{ S_addr uint32 }
	V6Addr   struct{ S6Addr [16]uint8 }
	UnixAddr struct{ Path [108]uint8 }
	Padding  uint32
	Ret      int32
}

// loadAccept returns the embedded CollectionSpec for accept.
func loadAccept() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_AcceptBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load accept: %w", err)
	}

	return spec, err
}

// loadAcceptObjects loads accept and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*acceptObjects
//	*acceptPrograms
//	*acceptMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadAcceptObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadAccept()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// acceptSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type acceptSpecs struct {
	acceptProgramSpecs
	acceptMapSpecs
}

// acceptSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type acceptProgramSpecs struct {
	KprobeAcceptEntry   *ebpf.ProgramSpec `ebpf:"kprobe_accept_entry"`
	KretprobeAcceptExit *ebpf.ProgramSpec `ebpf:"kretprobe_accept_exit"`
}

// acceptMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type acceptMapSpecs struct {
	AcceptEventMap *ebpf.MapSpec `ebpf:"accept_event_map"`
}

// acceptObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadAcceptObjects or ebpf.CollectionSpec.LoadAndAssign.
type acceptObjects struct {
	acceptPrograms
	acceptMaps
}

func (o *acceptObjects) Close() error {
	return _AcceptClose(
		&o.acceptPrograms,
		&o.acceptMaps,
	)
}

// acceptMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadAcceptObjects or ebpf.CollectionSpec.LoadAndAssign.
type acceptMaps struct {
	AcceptEventMap *ebpf.Map `ebpf:"accept_event_map"`
}

func (m *acceptMaps) Close() error {
	return _AcceptClose(
		m.AcceptEventMap,
	)
}

// acceptPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadAcceptObjects or ebpf.CollectionSpec.LoadAndAssign.
type acceptPrograms struct {
	KprobeAcceptEntry   *ebpf.Program `ebpf:"kprobe_accept_entry"`
	KretprobeAcceptExit *ebpf.Program `ebpf:"kretprobe_accept_exit"`
}

func (p *acceptPrograms) Close() error {
	return _AcceptClose(
		p.KprobeAcceptEntry,
		p.KretprobeAcceptExit,
	)
}

func _AcceptClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed accept_bpfel_x86.o
var _AcceptBytes []byte
