package builder

import (
	"fmt"
	"strings"
)

// Builder facet implementation

// We want to keep the Machine struct hidden from usage by client.
// All use must be able to use the hidden object via the parameterized builder.
type machine struct {
	name   string
	cpu    int
	memory int
}

func (m *machine) string() string {
	var b strings.Builder

	b.WriteString("\n	VM Name - ")
	b.WriteString(m.name)

	b.WriteString("\n	vcpus - ")
	b.WriteString(fmt.Sprint(m.cpu))

	b.WriteString("\n	memory - ")
	b.WriteString(fmt.Sprint(m.memory))

	return b.String()
}

func BuildMachine(buildaction func(*ParameterizedMachineBuilder)) {
	b := ParameterizedMachineBuilder{}
	buildaction(&b)
	buildMachineImpl(&b.machine)
	// print the machine created
	fmt.Printf("	%s", b.machine.string())
}

func buildMachineImpl(machine *machine) {
	//
	fmt.Printf("\nCreating machine\n")
}

type ParameterizedMachineBuilder struct {
	machine machine
}

func (b *ParameterizedMachineBuilder) Name(name string) *ParameterizedMachineBuilder {
	b.machine.name = name
	return b
}

func (b *ParameterizedMachineBuilder) WithCPU(numCPU int) *ParameterizedMachineBuilder {
	if numCPU <= 0 {
		panic("Invalid CPU configuration specified")
	}
	b.machine.cpu = numCPU
	return b
}

func (b *ParameterizedMachineBuilder) WithMemory(memory int) *ParameterizedMachineBuilder {
	if memory <= 0 {
		panic("Invalid Memory configuration specified")
	}
	b.machine.memory = memory
	return b
}
