package main

import (
	"fmt"

	"github.com/gofp/pkg/command"
)

func main() {

	//create a VM
	vm := &command.VM{Name: "newVM"}

	powerOnCmd := command.NewVMCommand(vm, command.CmdPowerON, 10, true)
	powerOnCmd.Call()
	fmt.Println("VM - ", *vm)

	powerOffCmd := command.NewVMCommand(vm, command.CmdPowerOff, 10, true)
	powerOffCmd.Call()
	fmt.Println("VM - ", *vm)
	// Lets undo the operation
	powerOffCmd.Undo()
	fmt.Println("VM - ", *vm)

	// create 2 VMs
	//create a VM
	vm1 := &command.VM{Name: "newVM1"}
	vm2 := &command.VM{Name: "newVM2"}

	vAppOps := command.NewVAppCommand(vm1, vm2, "192.168.172.65")
	vAppOps.Call()
	fmt.Println(vAppOps)

	//undo
	vAppOps.Undo()
	fmt.Println(vAppOps)
}
