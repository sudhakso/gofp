package main

import (
	"fmt"

	"github.com/gofp/pkg/composite"
)

func main() {
	n1 := composite.NewNeuron("n1", "*")
	n2 := composite.NewNeuron("n2", "*")
	n3 := composite.NewNeuron("n3", "*")

	nl1 := composite.NewNeutronLayer(2)
	nl1.Neurons = append(nl1.Neurons, *composite.NewNeuron("nl1.1", "**"))
	nl1.Neurons = append(nl1.Neurons, *composite.NewNeuron("nl1.2", "**"))

	// composite.Connect(n1, nl1)
	composite.Connect(n1, n2)
	composite.Connect(n1, n3)

	composite.Connect(n1, nl1)

	fmt.Println(n1)
}
