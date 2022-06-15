// Copyright 2022 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file https://github.com/redpanda-data/redpanda/blob/dev/licenses/bsl.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package git

import "flag"

type BasicAuthConfig struct {
	Enabled  bool   `yaml:"enabled" json:"enabled"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
}

// RegisterFlagsWithPrefix for sensitive Basic Auth configs
func (c *BasicAuthConfig) RegisterFlagsWithPrefix(f *flag.FlagSet, prefix string) {
	f.StringVar(&c.Password, prefix+"git.basic-auth.password", "", "Basic Auth password")
}
