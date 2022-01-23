package dns

import (
	"context"
	"errors"
	"net"
	"strings"
	"time"
)

type Resolver struct {
	ctx context.Context
	r   *net.Resolver
}

type ResolverOptions struct {
	Server   string
	Protocol string
}

func NewResolver(ctx context.Context, opts ResolverOptions) (*Resolver, error) {
	dnsResolverAddress := opts.Server
	if isEmptyString(dnsResolverAddress) {
		return nil, errors.New("DNS server required.")
	}

	dnsResolverProtocol := opts.Protocol
	if isEmptyString(dnsResolverProtocol) {
		dnsResolverProtocol = "udp"
	}

	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network string, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Second * time.Duration(1),
			}
			return d.DialContext(ctx, dnsResolverProtocol, dnsResolverAddress)
		},
	}

	return &Resolver{ctx, r}, nil
}

func (r Resolver) Resolve(domain string) (string, error) {
	ips, err := r.r.LookupHost(r.ctx, domain)
	if err != nil {
		return "", err
	}

	return ips[0], nil
}

func isEmptyString(s string) bool {
	return strings.TrimSpace(s) == ""
}
