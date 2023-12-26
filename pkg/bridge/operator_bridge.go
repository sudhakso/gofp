package bridge

import "fmt"

type Flavor int

const (
	Small Flavor = iota
	Medium
	Large
	Xlarge
)

// Backup Operations
type BackupOperator interface {
	BackUpVirtualMachine() // VM
	BackUpCloudMachine()   // Cloud VM
	BackupBaremetal()      // Baremetal Server
	BackupKubernetes()     // Kubernetes cluster
}

// Bridge
// Utility for backing up different instance types
type BackUpTool struct {
}

func (b *BackUpTool) BackUpVirtualMachine() {
	// panic("not implemented") // TODO: Implement
	fmt.Printf("Configuring backup of vm\n")
}

func (b *BackUpTool) BackUpCloudMachine() {
	// panic("not implemented") // TODO: Implement
	fmt.Printf("Configuring backup of vm\n")
}

func (b *BackUpTool) BackupBaremetal() {
	// panic("not implemented") // TODO: Implement
	fmt.Printf("Configuring backup of baremetal\n")
}

func (b *BackUpTool) BackupKubernetes() {
	// panic("not implemented") // TODO: Implement
	fmt.Printf("Configuring backup of k8s\n")
}

// Security Operations
type SecurityOperator interface {
	SecureVirtualServer()     // VM
	SecureCloudServer()       // Cloud
	SecureBaremetalServer()   // Baremetal
	SecureKubernetesCluster() // Kubernetes
}

// Bridge
// Utility for securing instances
type SecurityTool struct {
}

func (s *SecurityTool) SecureVirtualServer() {
	// panic("not implemented") // TODO: Implement
	fmt.Printf("encrypting vm\n")
}

func (s *SecurityTool) SecureCloudServer() {
	// panic("not implemented") // TODO: Implement
	fmt.Printf("encrypting cloudvm\n")
}

func (s *SecurityTool) SecureBaremetalServer() {
	// panic("not implemented") // TODO: Implement
	fmt.Printf("encrypting baremetal\n")
}

func (s *SecurityTool) SecureKubernetesCluster() {
	// panic("not implemented") // TODO: Implement
	fmt.Printf("encrypting k8s\n")
}

type CloudServer struct {
	backupOps BackupOperator   // BackUpTool is the bridge b/w CloudServer and Backup service
	secOps    SecurityOperator // SecurityTool is the bridge b/w CloudServer and Security service

	name string
}

func NewCloudServer(name string, b BackupOperator, s SecurityOperator) *CloudServer {
	return &CloudServer{b, s, name}
}

func (c *CloudServer) Backup() {
	c.backupOps.BackUpCloudMachine()
}

func (c *CloudServer) Encrypt() {
	c.secOps.SecureVirtualServer()
}

type PhysicalServer struct {
	backupOps BackupOperator
	secOps    SecurityOperator

	// ...Other attributes
	name string
}

func NewPhysicalServer(name string, b BackupOperator, s SecurityOperator) *PhysicalServer {
	return &PhysicalServer{b, s, name}
}

func (c *PhysicalServer) Backup() {
	c.backupOps.BackupBaremetal()
}

func (c *PhysicalServer) Encrypt() {
	c.secOps.SecureBaremetalServer()
}

type KubernetesCluster struct {
	backupOps BackupOperator
	secOps    SecurityOperator

	// ...Other attributes
	name string
}

func NewKubernetesCluster(name string, b BackupOperator, s SecurityOperator) *KubernetesCluster {
	return &KubernetesCluster{b, s, name}
}

func (c *KubernetesCluster) Backup() {
	c.backupOps.BackupKubernetes()
}

func (c *KubernetesCluster) Encrypt() {
	c.secOps.SecureKubernetesCluster()
}
