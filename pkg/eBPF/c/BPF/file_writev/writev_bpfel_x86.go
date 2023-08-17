// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64

package file_writev

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type writevEventData struct {
	EventContext struct {
		Ts        uint64
		StartTime uint64
		Pid       uint32
		Tgid      uint32
		Ppid      uint32
		Glpid     uint32
		Uid       uint32
		Gid       uint32
		Comm      [16]uint8
		Cwd       [32]uint8
		NodeInfo  struct {
			Sysname    [65]uint8
			Nodename   [65]uint8
			Release    [65]uint8
			Version    [65]uint8
			Machine    [65]uint8
			Domainname [65]uint8
		}
	}
	_    [2]byte
	Id   int32
	_    [4]byte
	Fd   uint64
	Vlen uint64
	Ret  int64
}

// loadWritev returns the embedded CollectionSpec for writev.
func loadWritev() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_WritevBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load writev: %w", err)
	}

	return spec, err
}

// loadWritevObjects loads writev and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*writevObjects
//	*writevPrograms
//	*writevMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadWritevObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadWritev()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// writevSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type writevSpecs struct {
	writevProgramSpecs
	writevMapSpecs
}

// writevSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type writevProgramSpecs struct {
	KprobeWritevEntry   *ebpf.ProgramSpec `ebpf:"kprobe_writev_entry"`
	KretprobeWritevExit *ebpf.ProgramSpec `ebpf:"kretprobe_writev_exit"`
}

// writevMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type writevMapSpecs struct {
	Event *ebpf.MapSpec `ebpf:"event"`
}

// writevObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadWritevObjects or ebpf.CollectionSpec.LoadAndAssign.
type writevObjects struct {
	writevPrograms
	writevMaps
}

func (o *writevObjects) Close() error {
	return _WritevClose(
		&o.writevPrograms,
		&o.writevMaps,
	)
}

// writevMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadWritevObjects or ebpf.CollectionSpec.LoadAndAssign.
type writevMaps struct {
	Event *ebpf.Map `ebpf:"event"`
}

func (m *writevMaps) Close() error {
	return _WritevClose(
		m.Event,
	)
}

// writevPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadWritevObjects or ebpf.CollectionSpec.LoadAndAssign.
type writevPrograms struct {
	KprobeWritevEntry   *ebpf.Program `ebpf:"kprobe_writev_entry"`
	KretprobeWritevExit *ebpf.Program `ebpf:"kretprobe_writev_exit"`
}

func (p *writevPrograms) Close() error {
	return _WritevClose(
		p.KprobeWritevEntry,
		p.KretprobeWritevExit,
	)
}

func _WritevClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed writev_bpfel_x86.o
var _WritevBytes []byte
