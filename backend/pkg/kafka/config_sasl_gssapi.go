// Copyright 2022 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file https://github.com/redpanda-data/redpanda/blob/dev/licenses/bsl.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package kafka

import (
	"flag"
	"fmt"
)

// SASLGSSAPIConfig represents the Kafka Kerberos config
type SASLGSSAPIConfig struct {
	AuthType           string `yaml:"authType" json:"authType"`
	KeyTabPath         string `yaml:"keyTabPath" json:"keyTabPath"`
	KerberosConfigPath string `yaml:"kerberosConfigPath" json:"kerberosConfigPath"`
	ServiceName        string `yaml:"serviceName" json:"serviceName"`
	Username           string `yaml:"username" json:"username"`
	Password           string `yaml:"password" json:"password"`
	Realm              string `yaml:"realm" json:"realm"`

	// EnableFAST enables FAST, which is a pre-authentication framework for Kerberos.
	// It includes a mechanism for tunneling pre-authentication exchanges using armoured KDC messages.
	// FAST provides increased resistance to passive password guessing attacks.
	EnableFast bool `yaml:"enableFast" json:"enableFast"`
}

// RegisterFlags registers all sensitive Kerberos settings as flag
func (c *SASLGSSAPIConfig) RegisterFlags(f *flag.FlagSet) {
	f.StringVar(&c.Password, "kafka.sasl.gssapi.password", "", "Kerberos password if auth type user auth is used")
}

func (c *SASLGSSAPIConfig) Validate() error {
	if c.AuthType != "USER_AUTH" && c.AuthType != "KEYTAB_AUTH" {
		return fmt.Errorf("auth type '%v' is invalid", c.AuthType)
	}

	return nil
}

func (s *SASLGSSAPIConfig) SetDefaults() {
	s.EnableFast = true
}
