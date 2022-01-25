package main

import (
	"context"
	"flag"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/cheynewallace/tabby"
	"github.com/jdecool/dnspropagation/internal/configuration"
	"github.com/jdecool/dnspropagation/internal/dns"
)

var (
	configFile string
)

type DNSResult struct {
	Provider  string
	Primary   string
	Secondary string
}

func main() {
	flag.StringVar(&configFile, "config", "", "Configuration file to use")
	flag.Parse()

	if len(flag.Args()) != 1 {
		panic("Missing domain argument.")
	}

	domain := flag.Arg(0)
	if strings.TrimSpace(domain) == "" {
		panic("Missing domain argument.")
	}

	config, err := configuration.Load(configFile)
	if err != nil {
		panic(err)
	}

	start := time.Now()
	r := handleDNSResolution(*config, domain)
	t := time.Since(start)

	displayOutput(r, t)
}

func handleDNSResolution(config configuration.Config, domain string) map[string]DNSResult {
	var wg sync.WaitGroup

	ch := make(chan DNSResult)
	for _, p := range config.Providers {
		wg.Add(1)
		go resolveProvider(ch, p, domain)
	}

	rm := make(map[string]DNSResult)
	go func() {
		for r := range ch {
			rm[r.Provider] = r
			wg.Done()
		}
	}()

	wg.Wait()

	return rm
}

func resolveProvider(ch chan DNSResult, p configuration.DNSProvider, d string) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	var wg sync.WaitGroup
	f := func(ch chan string, name string, addr string, protocol string) {
		wg.Add(1)
		defer wg.Done()

		r, err := dns.NewResolver(context.Background(), dns.ResolverOptions{
			Server:   addr,
			Protocol: protocol,
		})
		if err != nil {
			ch <- err.Error()
			return
		}

		ip, err := r.Resolve(d)
		if err != nil {
			ch <- err.Error()
			return
		}

		ch <- ip
	}

	go f(ch1, p.Name, p.Server.Primary, p.Server.Protocol)
	go f(ch2, p.Name, p.Server.Secondary, p.Server.Protocol)
	wg.Wait()

	ch <- DNSResult{
		Provider:  p.Name,
		Primary:   <-ch1,
		Secondary: <-ch2,
	}
}

func displayOutput(r map[string]DNSResult, d time.Duration) {
	keys := make([]string, 0, len(r))
	for k := range r {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	t := tabby.New()
	t.AddHeader("Provider", "Primary", "Secondary")
	for _, k := range keys {
		p := r[k]

		t.AddLine(p.Provider, p.Primary, p.Secondary)
	}
	t.Print()

	fmt.Printf("\nTime: %.2f seconds\n", d.Seconds())
}
