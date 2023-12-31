package main

import (
	"fmt"

	"github.com/gofp/pkg/flywheel"
)

// Flyweight
// Goal: Optimize on usage of space by sharing constructs amongst many concrete objects.

func main() {

	// VMs
	dummy := &flywheel.VirtualMachine{Name: "dummy"}
	vm1 := &flywheel.VirtualMachine{Name: "vm1"}
	vm2 := &flywheel.VirtualMachine{Name: "vm2"}
	vm3 := &flywheel.VirtualMachine{Name: "vm3"}
	vm4 := &flywheel.VirtualMachine{Name: "vm4"}
	vm5 := &flywheel.VirtualMachine{Name: "vm5"}
	vm6 := &flywheel.VirtualMachine{Name: "vm6"}

	// manager
	manager := flywheel.VirtManager{VM: []*flywheel.VirtualMachine{dummy, vm1, vm2, vm3, vm4, vm5, vm6}}

	// Policy to Encrypt & BackUp
	p_1 := manager.Apply(1, 3)
	p_1.Encrypt = true
	p_1.BackUp = true

	// Policy to Snapshot
	p_2 := manager.Apply(5, 6)
	p_2.MinutelySnapshot = true

	fmt.Printf(manager.Display())
}
