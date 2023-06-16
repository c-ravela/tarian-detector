// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || amd64p32 || arm || arm64 || mips64le || mips64p32le || mipsle || ppc64le || riscv64
// +build 386 amd64 amd64p32 arm arm64 mips64le mips64p32le mipsle ppc64le riscv64

package process_exit

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type exitEventData struct {
	Pid       uint32
	Tgid      uint32
	Uid       uint32
	Gid       uint32
	SyscallNr int32
	_         [4]byte
	Ret       int64
	Comm      [16]uint8
	Cwd       [32]uint8
}

// loadExit returns the embedded CollectionSpec for exit.
func loadExit() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_ExitBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load exit: %w", err)
	}

	return spec, err
}

// loadExitObjects loads exit and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*exitObjects
//	*exitPrograms
//	*exitMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadExitObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadExit()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// exitSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type exitSpecs struct {
	exitProgramSpecs
	exitMapSpecs
}

// exitSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type exitProgramSpecs struct {
	ExecveExit *ebpf.ProgramSpec `ebpf:"execve_exit"`
}

// exitMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type exitMapSpecs struct {
	Event *ebpf.MapSpec `ebpf:"event"`
}

// exitObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadExitObjects or ebpf.CollectionSpec.LoadAndAssign.
type exitObjects struct {
	exitPrograms
	exitMaps
}

func (o *exitObjects) Close() error {
	return _ExitClose(
		&o.exitPrograms,
		&o.exitMaps,
	)
}

// exitMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadExitObjects or ebpf.CollectionSpec.LoadAndAssign.
type exitMaps struct {
	Event *ebpf.Map `ebpf:"event"`
}

func (m *exitMaps) Close() error {
	return _ExitClose(
		m.Event,
	)
}

// exitPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadExitObjects or ebpf.CollectionSpec.LoadAndAssign.
type exitPrograms struct {
	ExecveExit *ebpf.Program `ebpf:"execve_exit"`
}

func (p *exitPrograms) Close() error {
	return _ExitClose(
		p.ExecveExit,
	)
}

func _ExitClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed exit_bpfel.o
var _ExitBytes []byte
