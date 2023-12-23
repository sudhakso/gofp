package builder

// Lazy builder implementation

type LazyMachine struct {
	name string
	cpu  int
}

type action func(machine *LazyMachine)

// Builder contains a series of actions to apply to the object
// Builder no longer aggregates the object itself
type LazyMachineBuilder struct {
	actions []action
}

func (b *LazyMachineBuilder) Name(name string) *LazyMachineBuilder {
	b.actions = append(b.actions, func(m *LazyMachine) {
		m.name = name
	})
	return b
}

func (b *LazyMachineBuilder) WithCPU(cpu int) *LazyMachineBuilder {
	b.actions = append(b.actions, func(m *LazyMachine) {
		m.cpu = cpu
	})
	return b
}

func (b *LazyMachineBuilder) Build() *LazyMachine {
	m := LazyMachine{}
	for _, v := range b.actions {
		v(&m)
	}
	return &m
}
