package configuration

import "github.com/hashicorp/hcl/v2/hclsimple"

type Config struct {
	Providers []DNSProvider `hcl:"dns,block"`
}

type DNSProvider struct {
	Name   string `hcl:"name"`
	Server Server `hcl:"server,block"`
}

type Server struct {
	Protocol  string `hcl:"protocol"`
	Primary   string `hcl:"primary"`
	Secondary string `hcl:"secondary"`
}

func Load(file string) (*Config, error) {
	var config Config
	err := hclsimple.DecodeFile(file, nil, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
