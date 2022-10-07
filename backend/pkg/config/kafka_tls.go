// Copyright 2022 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file https://github.com/redpanda-data/redpanda/blob/dev/licenses/bsl.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package config

import "flag"

// KafkaTLS to connect to Kafka via TLS
type KafkaTLS struct {
	Enabled               bool   `yaml:"enabled"`
	CaFilepath            string `yaml:"caFilepath"`
	CertFilepath          string `yaml:"certFilepath"`
	KeyFilepath           string `yaml:"keyFilepath"`
	Passphrase            string `yaml:"passphrase"`
	InsecureSkipTLSVerify bool   `yaml:"insecureSkipTlsVerify"`
}

// RegisterFlags for all sensitive Kafka TLS configs
func (c *KafkaTLS) RegisterFlags(f *flag.FlagSet) {
	f.StringVar(&c.Passphrase, "kafka.tls.passphrase", "", "Passphrase to optionally decrypt the private key")
}
