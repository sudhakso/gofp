package flywheel

import (
	"fmt"
	"strings"
)

type VirtualMachine struct {
	Name string
}

type Policy struct {
	Start, End                        int
	Encrypt, BackUp, MinutelySnapshot bool
}

func (p *Policy) Covers(index int) bool {
	return index >= p.Start && index <= p.End
}

// Holds representation of VMs with associated policies
type VirtManager struct {
	VM       []*VirtualMachine
	policies []*Policy
}

func (v *VirtManager) Apply(start, end int) *Policy {
	p := &Policy{start, end, false, false, false}
	v.policies = append(v.policies, p)
	return p
}

func (v *VirtManager) Display() string {
	s := strings.Builder{}
	for i := 0; i < len(v.VM); i++ {
		s.WriteString("\n\nVM =>")
		s.WriteString(v.VM[i].Name)
		s.WriteString("  \nApplied policies: \n")
		for j, p := range v.policies {

			s.WriteString(fmt.Sprintf("\nPolicy %d", j))
			//Encrypt
			s.WriteString("   \nIs Encrypted: ")
			if p.Covers(i) && p.Encrypt {
				s.WriteString("    true")
			} else {
				s.WriteString("    false")
			}

			// Backup
			s.WriteString("   \nIs Backed-up: ")
			if p.Covers(i) && p.BackUp {
				s.WriteString("    true")
			} else {
				s.WriteString("    false")
			}

			// Snapshotted
			s.WriteString("   \nIs Snapshotted: ")
			if p.Covers(i) && p.MinutelySnapshot {
				s.WriteString("    true")
			} else {
				s.WriteString("    false")
			}
		}
	}
	return s.String()
}
