package builder

// Builder facet implementation

import (
	"fmt"
	"strings"
)

type Machine struct {
	name         string
	cpu          int
	memory       int
	rootDiskSize int
	network      string
}

func (m *Machine) String() string {
	var b strings.Builder

	b.WriteString("\n	VM Name - ")
	b.WriteString(m.name)

	b.WriteString("\n	vcpus - ")
	b.WriteString(fmt.Sprint(m.cpu))

	b.WriteString("\n	memory - ")
	b.WriteString(fmt.Sprint(m.memory))

	b.WriteString("\n	disksize - ")
	b.WriteString(fmt.Sprint(m.rootDiskSize))

	b.WriteString("\n	network - ")
	b.WriteString(m.network)

	return b.String()
}

type MachineBuilder struct {
	machine *Machine
}

type MachineComputeBuilder struct {
	MachineBuilder
}

func NewMachineBuilder() *MachineBuilder {
	return &MachineBuilder{&Machine{}}
}

func (b *MachineBuilder) Compute() *MachineComputeBuilder {
	return &MachineComputeBuilder{*b}
}

func (b *MachineBuilder) Storage() *MachineStorageBuilder {
	return &MachineStorageBuilder{*b}
}

func (b *MachineBuilder) Network() *MachineNetworkBuilder {
	return &MachineNetworkBuilder{*b}
}

func (b *MachineBuilder) Name(name string) *MachineBuilder {
	b.machine.name = name
	return b
}

func (b *MachineBuilder) Build() *Machine {
	return b.machine
}

func (b *MachineComputeBuilder) WithCPU(numCPU int) *MachineComputeBuilder {
	b.machine.cpu = numCPU
	return b
}

func (b *MachineComputeBuilder) Memory(memory int) *MachineComputeBuilder {
	b.machine.memory = memory
	return b
}

type MachineStorageBuilder struct {
	MachineBuilder
}

func (b *MachineStorageBuilder) AllocateDisk(size int) *MachineStorageBuilder {
	b.machine.rootDiskSize = size
	return b
}

type MachineNetworkBuilder struct {
	MachineBuilder
}

func (b *MachineNetworkBuilder) AttachNetwork(netname string) *MachineNetworkBuilder {
	b.machine.network = netname
	return b
}
