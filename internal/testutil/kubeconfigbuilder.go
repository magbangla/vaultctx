// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testutil

import (
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

type Context struct {
	Name    string `yaml:"name,omitempty"`
	Context struct {
		Namespace string `yaml:"namespace,omitempty"`
	} `yaml:"context,omitempty"`
}

func Ctx(name string) *Context           { return &Context{Name: name} }
func (c *Context) Ns(ns string) *Context { c.Context.Namespace = ns; return c }

type vaultconfig map[string]interface{}

func KC() *vaultconfig {
	return &vaultconfig{
		"apiVersion": "v1",
		"kind":       "Config"}
}

func (k *vaultconfig) Set(key string, v interface{}) *vaultconfig { (*k)[key] = v; return k }
func (k *vaultconfig) WithCurrentCtx(s string) *vaultconfig       { (*k)["current-context"] = s; return k }
func (k *vaultconfig) WithCtxs(c ...*Context) *vaultconfig        { (*k)["contexts"] = c; return k }

func (k *vaultconfig) ToYAML(t *testing.T) string {
	t.Helper()
	var v strings.Builder
	if err := yaml.NewEncoder(&v).Encode(*k); err != nil {
		t.Fatalf("failed to encode mock vaultconfig: %v", err)
	}
	return v.String()
}
