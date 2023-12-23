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

}
