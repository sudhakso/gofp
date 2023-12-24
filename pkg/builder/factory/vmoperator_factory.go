package factory

import "fmt"

type Operator interface {
	PowerOn()
	PowerOFF()
	Suspend()
}

type PowerState int

const (
	On PowerState = iota
	Off
	Fuzzy
)

type machine struct {
	name  string
	state PowerState
	band  PowerBand
}

func (o *machine) PowerOn() {
	fmt.Printf("Powering ON %s\n", o.name)
	o.state = On
}

func (o *machine) PowerOFF() {
	fmt.Printf("Powering OFF %s\n", o.name)
	o.state = Off
}

func (o *machine) Suspend() {
	fmt.Printf("Suspending %s\n", o.name)
	o.state = Fuzzy
}

type warmedUpmachine struct {
	name  string
	state PowerState
	band  PowerBand
}

func (o *warmedUpmachine) PowerOn() {
	fmt.Printf("Its always ON %s\n", o.name)
}

func (o *warmedUpmachine) PowerOFF() {
	fmt.Printf("Its always ON %s, cannot be turned OFF\n", o.name)
}

func (o *warmedUpmachine) Suspend() {
	fmt.Printf("Its always ON %s, cannot be turned Suspended\n", o.name)
}

// type:machine is hidden, it is exposed via a constructor that returns an interface
func NewMachine(name string, initialState PowerState) Operator {
	if initialState == On {
		//Pick a warmedUp machine
		return &warmedUpmachine{name, On, Cost}
	}
	return &machine{name, initialState, Cost}
}
