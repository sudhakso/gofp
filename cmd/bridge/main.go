package main

import (
	"fmt"

	"github.com/gofp/pkg/bridge"
)

func main() {

	// Construct the bridge
	b := bridge.BackUpTool{}
	s := bridge.SecurityTool{}

	// Construct the types
	vm := bridge.NewCloudServer("myVM", &b, &s)
	fmt.Println("New cloud virtualmachine created with name - myVM")

	vm.Backup()  // configure backup
	vm.Encrypt() // encrypt vm

	bm := bridge.NewPhysicalServer("myBM", &b, &s)
	fmt.Println("New physical machine created with name - myBM")

	bm.Backup()  // configure backup
	bm.Encrypt() // encrypt bm

	kc := bridge.NewPhysicalServer("myK8s", &b, &s)
	fmt.Println("New K8s cluster created with name - myK8s")

	kc.Backup()  // configure backup
	kc.Encrypt() // encrypt cluster

}
