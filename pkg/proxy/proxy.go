package proxy

import "fmt"

// A type that functions as an interface to a particular resource.
// That resource may be Remote, Expensive to construct, or might need additional
// capabilities like decorator.

// Proxy pattern is useful to contruct the underlying object in a lazy fashion, while the
// clients using do not really care or know.

// Proxy pattern can be used to share the object if required.
// clients won't even know.

// Categories of proxies - All proxies implement the same interface of the type that is being proxoed.
// .. Protection proxy (like AuthZ, AuthN...)
// .. Virtual proxy (like comms proxy...)

// Benefits of Proxy
// 1. Encapsulation / Hide Complexity
// 2. Lazy initializiation
// 3. Like Decorator, enhance the behavior transparently to the clients.

type Direction uint8

const (
	Egress = iota
	Ingress
)

var AllRules []*Rule

type Inspector interface {
	Allow(target string, data string) *Rule
	Deny(target string) *Rule

	IsAllowed(p string) bool

	Send(string, string)
}

type Rule struct {
	Destination         string
	AllowHttp, AllowSsh bool
}

type Firewall struct {
	AllRules []*Rule
}

func (f *Firewall) Allow(h string, d string) *Rule {
	r := &Rule{}
	r.AllowHttp, r.AllowSsh = true, true

	f.AllRules = append(f.AllRules, r)
	return r
}

func (f *Firewall) Deny(h string) *Rule {
	r := &Rule{}
	r.AllowHttp, r.AllowSsh = false, false

	f.AllRules = append(f.AllRules, r)
	return r
}

func (f *Firewall) IsAllowed(h string) bool {
	for _, r := range f.AllRules {
		if r.Destination == h && r.AllowHttp {
			return true
		}
	}
	return false
}

func (f *Firewall) Send(d string, p string) {
	if !f.IsAllowed(d) {
		fmt.Printf("Check the firewall rules. Request to %s not allowed\n", d)
		return
	}
	fmt.Printf("Sending data (%s) to the target (%s)\n", p, d)
}

type HttpClient struct {
	firewallproxy Inspector
}

func NewHttpClientWithFirewall(f Inspector) *HttpClient {
	return &HttpClient{f}
}

// send data
func (h *HttpClient) Send(dest string, payload string) {
	h.firewallproxy.Send(dest, payload)
}
