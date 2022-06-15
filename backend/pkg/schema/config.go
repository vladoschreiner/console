// Copyright 2022 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file https://github.com/redpanda-data/redpanda/blob/dev/licenses/bsl.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package schema

import (
	"flag"
	"fmt"
)

// Config for using a (Confluent) Schema Registry
type Config struct {
	Enabled bool     `yaml:"enabled" json:"enabled"`
	URLs    []string `yaml:"urls" json:"urls"`

	// Credentials
	Username    string `yaml:"username" json:"username"`
	Password    string `yaml:"password" json:"password"`
	BearerToken string `yaml:"bearerToken" json:"bearerToken"`

	// TLS / Custom CA
	TLS TLSConfig `yaml:"tls" json:"tls"`
}

// RegisterFlags registers all nested config flags.
func (c *Config) RegisterFlags(f *flag.FlagSet) {
	f.StringVar(&c.Password, "schema.registry.password", "", "Password for authenticating against the schema registry (optional)")
	f.StringVar(&c.BearerToken, "schema.registry.token", "", "Bearer token for authenticating against the schema registry (optional)")
}

func (c *Config) Validate() error {
	if c.Enabled == false {
		return nil
	}

	if len(c.URLs) == 0 {
		return fmt.Errorf("schema registry is enabled but no URL is configured")
	}

	return nil
}
