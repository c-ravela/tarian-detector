// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64
// +build 386 amd64

package file_openat2

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type openat2EventData struct {
	Pid      uint32
	Tgid     uint32
	Uid      uint32
	Gid      uint32
	Filename [256]uint8
	Fd       int32
}

// loadOpenat2 returns the embedded CollectionSpec for openat2.
func loadOpenat2() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_Openat2Bytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load openat2: %w", err)
	}

	return spec, err
}

// loadOpenat2Objects loads openat2 and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*openat2Objects
//	*openat2Programs
//	*openat2Maps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadOpenat2Objects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadOpenat2()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// openat2Specs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type openat2Specs struct {
	openat2ProgramSpecs
	openat2MapSpecs
}

// openat2Specs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type openat2ProgramSpecs struct {
	KprobeOpenat2 *ebpf.ProgramSpec `ebpf:"kprobe_openat2"`
}

// openat2MapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type openat2MapSpecs struct {
	Event *ebpf.MapSpec `ebpf:"event"`
}

// openat2Objects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadOpenat2Objects or ebpf.CollectionSpec.LoadAndAssign.
type openat2Objects struct {
	openat2Programs
	openat2Maps
}

func (o *openat2Objects) Close() error {
	return _Openat2Close(
		&o.openat2Programs,
		&o.openat2Maps,
	)
}

// openat2Maps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadOpenat2Objects or ebpf.CollectionSpec.LoadAndAssign.
type openat2Maps struct {
	Event *ebpf.Map `ebpf:"event"`
}

func (m *openat2Maps) Close() error {
	return _Openat2Close(
		m.Event,
	)
}

// openat2Programs contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadOpenat2Objects or ebpf.CollectionSpec.LoadAndAssign.
type openat2Programs struct {
	KprobeOpenat2 *ebpf.Program `ebpf:"kprobe_openat2"`
}

func (p *openat2Programs) Close() error {
	return _Openat2Close(
		p.KprobeOpenat2,
	)
}

func _Openat2Close(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed openat2_bpfel_x86.o
var _Openat2Bytes []byte
