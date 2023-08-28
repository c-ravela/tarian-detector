// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64

package network_connect

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type connectEventData struct {
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
	_        [2]byte
	Id       int32
	Fd       int32
	Addrlen  int32
	Ret      int32
	SaFamily uint16
	Port     uint16
	V4Addr   struct{ S_addr uint32 }
	V6Addr   struct{ S6Addr [16]uint8 }
	UnixAddr struct{ Path [108]int8 }
	Padding  uint32
}

// loadConnect returns the embedded CollectionSpec for connect.
func loadConnect() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_ConnectBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load connect: %w", err)
	}

	return spec, err
}

// loadConnectObjects loads connect and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*connectObjects
//	*connectPrograms
//	*connectMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadConnectObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadConnect()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// connectSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type connectSpecs struct {
	connectProgramSpecs
	connectMapSpecs
}

// connectSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type connectProgramSpecs struct {
	KprobeConnectEntry   *ebpf.ProgramSpec `ebpf:"kprobe_connect_entry"`
	KretprobeConnectExit *ebpf.ProgramSpec `ebpf:"kretprobe_connect_exit"`
}

// connectMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type connectMapSpecs struct {
	ConnectEventMap *ebpf.MapSpec `ebpf:"connect_event_map"`
}

// connectObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadConnectObjects or ebpf.CollectionSpec.LoadAndAssign.
type connectObjects struct {
	connectPrograms
	connectMaps
}

func (o *connectObjects) Close() error {
	return _ConnectClose(
		&o.connectPrograms,
		&o.connectMaps,
	)
}

// connectMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadConnectObjects or ebpf.CollectionSpec.LoadAndAssign.
type connectMaps struct {
	ConnectEventMap *ebpf.Map `ebpf:"connect_event_map"`
}

func (m *connectMaps) Close() error {
	return _ConnectClose(
		m.ConnectEventMap,
	)
}

// connectPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadConnectObjects or ebpf.CollectionSpec.LoadAndAssign.
type connectPrograms struct {
	KprobeConnectEntry   *ebpf.Program `ebpf:"kprobe_connect_entry"`
	KretprobeConnectExit *ebpf.Program `ebpf:"kretprobe_connect_exit"`
}

func (p *connectPrograms) Close() error {
	return _ConnectClose(
		p.KprobeConnectEntry,
		p.KretprobeConnectExit,
	)
}

func _ConnectClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed connect_bpfel_x86.o
var _ConnectBytes []byte
