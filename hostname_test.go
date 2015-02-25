package hostess_test

import (
	"github.com/cbednarski/hostess"
	"net"
	"testing"
)

func TestHostname(t *testing.T) {
	h := hostess.Hostname{}
	h.Domain = domain
	h.Ip = ip
	h.Enabled = enabled

	if h.Domain != domain {
		t.Errorf("Domain should be %s", domain)
	}
	if !h.Ip.Equal(ip) {
		t.Errorf("Ip should be %s", ip)
	}
	if h.Enabled != enabled {
		t.Errorf("Enabled should be %s", enabled)
	}
}

func TestEqual(t *testing.T) {
	a := hostess.NewHostname("localhost", "127.0.0.1", true)
	b := hostess.NewHostname("localhost", "127.0.0.1", false)
	c := hostess.NewHostname("localhost", "127.0.1.1", false)

	if !a.Equal(b) {
		t.Errorf("%s and %s should be equal", a, b)
	}
	if a.Equal(c) {
		t.Errorf("%s and %s should not be equal", a, c)
	}
}

func TestEqualIp(t *testing.T) {
	a := hostess.NewHostname("localhost", "127.0.0.1", true)
	c := hostess.NewHostname("localhost", "127.0.1.1", false)
	ip := net.ParseIP("127.0.0.1")

	if !a.EqualIp(ip) {
		t.Errorf("%s and %s should be equal", a.Ip, ip)
	}
	if a.EqualIp(c.Ip) {
		t.Errorf("%s and %s should not be equal", a.Ip, c.Ip)
	}
}

func TestIsValid(t *testing.T) {
	a := hostess.NewHostname("localhost", "127.0.0.1", true)
	d := hostess.NewHostname("", "127.0.0.1", true)
	e := hostess.NewHostname("localhost", "localhost", true)

	if !a.IsValid() {
		t.Errorf("%s should be a valid hostname", a)
	}
	if d.IsValid() {
		t.Errorf("%s should be invalid because the name is blank", d)
	}
	if e.IsValid() {
		t.Errorf("%s should be invalid because the ip is malformed", e)
	}
}

func TestFormatHostname(t *testing.T) {
	hostname := &hostess.Hostname{domain, ip, enabled, false}

	const exp_enabled = "127.0.0.1 localhost"
	if hostname.Format() != exp_enabled {
		t.Errorf(asserts, hostname.Format(), exp_enabled)
	}

	hostname.Enabled = false
	const exp_disabled = "# 127.0.0.1 localhost"
	if hostname.Format() != exp_disabled {
		t.Errorf(asserts, hostname.Format(), exp_disabled)
	}
}
