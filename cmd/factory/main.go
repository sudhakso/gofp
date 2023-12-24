package main

import (
	"fmt"

	"github.com/gofp/pkg/builder/factory"
)

func main() {
	m := factory.NewMachine("myVM", factory.On)

	//Operate on the machine
	m.PowerOn()
	// after a while...
	m.Suspend()

	// functional factories
	perf := factory.NewSpotInstanceFactory(factory.Performance)
	cost := factory.NewSpotInstanceFactory(factory.Cost)

	// Actual construction
	pm := perf("spotPerfVM", factory.On)
	pc := cost("sportCostVM", factory.Off)

	fmt.Println(pm)
	fmt.Println(pc)
}
