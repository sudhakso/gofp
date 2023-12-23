package main

import (
	"fmt"

	"github.com/gofp/pkg/builder"
)

func main() {

	b := builder.NewvSphereBuilder("http://dummy-vsphere-endpoint:6443")

	b.CreateVirtualInstance("myVM", &builder.ConfigData{NumCPU: 2, Memory: 4, RootDiskSize: 50})
	b.AttachVmToNetwork("VM Network")
	b.AllocateVmStorage(50, "/")

	vm := b.Create()
	// show VM properties
	fmt.Printf("Created VM - %s\n", vm.String())

	// Better builder via aggregation
	bb := builder.NewMachineBuilder()
	bb.Name("myBetterVM").
		Compute().
			WithCPU(4).
			Memory(16).
		Network().
			AttachNetwork("VM Network").
		Storage().
			AllocateDisk(120)
	
	mac :=	bb.Build()
	fmt.Printf("Create better VM - %s", mac.String())

	// builder parameterization to hide inner type
	// Here machine object is not accessible directly to the clients
	builder.BuildMachine(func(pmb *builder.ParameterizedMachineBuilder) {
		pmb.Name("myProtectedVM").
			WithCPU(8).
			WithMemory(64)
	})

	//Lazy initialization implementation using Builder
	lb := builder.LazyMachineBuilder{}
	lb.Name("LazyVM").
		WithCPU(16)
	
	//...After a while
	//Build
	lm := lb.Build()
	fmt.Println("\nLazy machine created - ", *lm)
}
