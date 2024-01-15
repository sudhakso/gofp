package command

type CompositeCommand struct {
	c []VMCommand
}

func (c *CompositeCommand) Call() {
	for _, cmd := range c.c {
		cmd.Call()
	}
}

func (c *CompositeCommand) Undo() {
	for idx := range c.c {
		c.c[len(c.c)-idx-1].Undo()
	}
}

// vApp command as the composite command

type VAppCommand struct {
	CompositeCommand
	lb_ip string
}

func NewVAppCommand(first *VM, second *VM, lb string) *VAppCommand {
	vmFirst := VMCommand{first, PowerON, 10, true, false}
	vmSecond := VMCommand{second, PowerON, 10, true, false}

	vapp := VAppCommand{lb_ip: lb}
	vapp.c = append(vapp.c, vmFirst)
	vapp.c = append(vapp.c, vmSecond)

	return &vapp
}
