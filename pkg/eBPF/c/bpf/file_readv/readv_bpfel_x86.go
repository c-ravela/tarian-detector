// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64

package file_readv

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type readvEventData struct {
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
		CgroupId  uint64
		NodeInfo  struct {
			Sysname    [65]uint8
			Nodename   [65]uint8
			Release    [65]uint8
			Version    [65]uint8
			Machine    [65]uint8
			Domainname [65]uint8
		}
		MountInfo struct {
			MountId      int32
			MountNsId    uint32
			MountDevname [256]uint8
		}
	}
	_    [2]byte
	Id   int32
	_    [4]byte
	Fd   uint64
	Vlen uint64
	Ret  int64
}

// loadReadv returns the embedded CollectionSpec for readv.
func loadReadv() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_ReadvBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load readv: %w", err)
	}

	return spec, err
}

// loadReadvObjects loads readv and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*readvObjects
//	*readvPrograms
//	*readvMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadReadvObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadReadv()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// readvSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type readvSpecs struct {
	readvProgramSpecs
	readvMapSpecs
}

// readvSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type readvProgramSpecs struct {
	KprobeReadvEntry   *ebpf.ProgramSpec `ebpf:"kprobe_readv_entry"`
	KretprobeReadvExit *ebpf.ProgramSpec `ebpf:"kretprobe_readv_exit"`
}

// readvMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type readvMapSpecs struct {
	Event *ebpf.MapSpec `ebpf:"event"`
}

// readvObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadReadvObjects or ebpf.CollectionSpec.LoadAndAssign.
type readvObjects struct {
	readvPrograms
	readvMaps
}

func (o *readvObjects) Close() error {
	return _ReadvClose(
		&o.readvPrograms,
		&o.readvMaps,
	)
}

// readvMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadReadvObjects or ebpf.CollectionSpec.LoadAndAssign.
type readvMaps struct {
	Event *ebpf.Map `ebpf:"event"`
}

func (m *readvMaps) Close() error {
	return _ReadvClose(
		m.Event,
	)
}

// readvPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadReadvObjects or ebpf.CollectionSpec.LoadAndAssign.
type readvPrograms struct {
	KprobeReadvEntry   *ebpf.Program `ebpf:"kprobe_readv_entry"`
	KretprobeReadvExit *ebpf.Program `ebpf:"kretprobe_readv_exit"`
}

func (p *readvPrograms) Close() error {
	return _ReadvClose(
		p.KprobeReadvEntry,
		p.KretprobeReadvExit,
	)
}

func _ReadvClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed readv_bpfel_x86.o
var _ReadvBytes []byte