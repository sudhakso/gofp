package builder

import "strings"

type VirtualMachine struct {
	name     string
	networks []VirtualNetwork
	disks    []Disk
}

func (vm *VirtualMachine) String() string {
	var b strings.Builder

	b.WriteString("VM Name - ")
	b.WriteString(vm.name)

	// Show networks
	// ...

	// Show disks
	// ...
	return b.String()
}

type VirtualNetwork struct {
	name   string //VM Network
	cidr   string //192.168.72.0/24
	ipaddr string //192.168.72.10
}

type Disk struct {
	name      string
	size      int
	mountPath string
}

type ConfigData struct {
	NumCPU       int
	Memory       int
	RootDiskSize int
}

type VmProvider interface {
	CreateVirtualInstance(name string, config *ConfigData) *VirtualMachine
}

type VmNetworkProvider interface {
	//CreateVmNetwork(name string, config *NetworkConfigData) *VirtualElement
	AttachVmToNetwork(vm *VirtualMachine, vmnet string) *VirtualMachine
}

type VmStorageProvider interface {
	//CreateVmStorage(name string, config *StorageConfigData) *VirtualElement
	AllocateVmStorage(vm *VirtualMachine, size int, mount string) *VirtualMachine
}

type VirtualMachineBuilder interface {
	VmProvider
	VmNetworkProvider
	VmStorageProvider

	Reset()
	Create() *VirtualMachine
}

type vSphereVMBuilder struct {
	vsphereEndpoint string
	vm              *VirtualMachine
}

func (v *vSphereVMBuilder) CreateVirtualInstance(name string, config *ConfigData) *VirtualMachine {
	v.vm.name = name
	v.vm.networks = make([]VirtualNetwork, 0)
	v.vm.disks = make([]Disk, 0)

	return v.vm
}

func (v *vSphereVMBuilder) AttachVmToNetwork(vmnet string) *VirtualMachine {
	// Lookup the network by name, and allocate a port for the VM
	net := VirtualNetwork{"VM Network", "192.168.72.0/24", "192.168.72.10"}
	// Do other VMWare type things, to bind the VM to a port on vswitch
	// ...
	v.vm.networks = append(v.vm.networks, net)
	return v.vm
}

func (v *vSphereVMBuilder) AllocateVmStorage(size int, mount string) *VirtualMachine {
	//Create disk element for the VM using VMware apis
	disk := Disk{v.vm.name + "-disk1", size, mount}
	// Do other VMware type of things, to bind the disk to the VM as a PCE device
	// ...
	v.vm.disks = append(v.vm.disks, disk)
	return v.vm
}

func (v *vSphereVMBuilder) Reset() {

	// detach network
	// deallocate disks
	// powerOFF instance
	// delete virtualinstance
}

func (v *vSphereVMBuilder) Create() *VirtualMachine {
	return v.vm
}

// func (v *vSphereVMBuilder) CreateVM(name string, cfg ConfigData) *VirtualMachine {
// 	// multi-step process

// 	//step1 - create virtualinstance
// 	//step2 - attach to vm network
// 	//step3 - attach storage to vm
// 	//step4 - powerON vm
// }

func NewvSphereBuilder(endpoint string) *vSphereVMBuilder {
	return &vSphereVMBuilder{endpoint, &VirtualMachine{}}
}
