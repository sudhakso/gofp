package command

import "fmt"

type PowerState int

const (
	PowerOff = iota
	PowerON
	Initialized
)

type CmdPowerAction int

const (
	CmdPowerOff = iota
	CmdPowerON
)

type VM struct {
	Name  string
	state PowerState
}

func (v *VM) PowerOff(gracefulShutOff bool) {
	fmt.Printf("Powering OFF the VM %s\n", v.Name)
	v.state = PowerOff
}

func (v *VM) PowerON(startupDelay int) {
	fmt.Printf("Powering ON the VM %s\n", v.Name)
	v.state = PowerON
}

type Commander interface {
	Call()
	Undo()
}

type VMCommand struct {
	v                   *VM
	action              CmdPowerAction
	startDelay          int
	gracefulShutOff     bool
	lastCommandSucceded bool
}

func (vc *VMCommand) Call() {
	switch vc.action {
	case PowerON:
		vc.v.PowerON(vc.startDelay)
		vc.lastCommandSucceded = true
	case PowerOff:
		vc.v.PowerOff(vc.gracefulShutOff)
		vc.lastCommandSucceded = true
	}
}

func (vc *VMCommand) Undo() {
	if vc.lastCommandSucceded {
		switch vc.action {
		case PowerOff:
			vc.v.PowerON(vc.startDelay)
			vc.lastCommandSucceded = true
		case PowerON:
			vc.v.PowerOff(vc.gracefulShutOff)
			vc.lastCommandSucceded = true
		}
	}
}

func NewVMCommand(v *VM, a CmdPowerAction, delay int, graceful bool) *VMCommand {
	return &VMCommand{v, a, delay, graceful, false}
}
