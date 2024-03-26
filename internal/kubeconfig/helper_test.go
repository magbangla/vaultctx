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

package vaultconfig

import (
	"bytes"
	"io"
	"strings"
)

type MockvaultconfigLoader struct {
	in  io.Reader
	out bytes.Buffer
}

func (t *MockvaultconfigLoader) Read(p []byte) (n int, err error)  { return t.in.Read(p) }
func (t *MockvaultconfigLoader) Write(p []byte) (n int, err error) { return t.out.Write(p) }
func (t *MockvaultconfigLoader) Close() error                      { return nil }
func (t *MockvaultconfigLoader) Reset() error                      { return nil }
func (t *MockvaultconfigLoader) Load() ([]ReadWriteResetCloser, error) {
	return []ReadWriteResetCloser{ReadWriteResetCloser(t)}, nil
}
func (t *MockvaultconfigLoader) Output() string { return t.out.String() }

func WithMockvaultconfigLoader(vaultcfg string) *MockvaultconfigLoader {
	return &MockvaultconfigLoader{in: strings.NewReader(vaultcfg)}
}
