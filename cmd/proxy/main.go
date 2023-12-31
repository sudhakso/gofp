package main

import "github.com/gofp/pkg/proxy"

func main() {

	// Create a firewall instance with a set of rules.
	f := proxy.Firewall{AllRules: []*proxy.Rule{
		{Destination: "http://google.com", AllowHttp: true, AllowSsh: true},
		{Destination: "http://microsoft.com", AllowHttp: true, AllowSsh: true},
		{Destination: "http://hackingsite.com", AllowHttp: false, AllowSsh: false},
		{Destination: "http://insecure.com", AllowHttp: false, AllowSsh: false},
	}}

	// others
	f.Allow("http://myskills.com", "")
	f.Deny("http://hackerrank.com")

	p := proxy.NewHttpClientWithFirewall(&f)
	p.Send("http://google.com", "Feeling lucky!") // Allowed

	p.Send("http://somerandom.com", "Feeling unlucky!") // Not Allowed

}
