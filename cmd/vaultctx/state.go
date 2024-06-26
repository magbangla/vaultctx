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

package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/magbangla/vaultctx/internal/cmdutil"
)

func vaultctxPrevCtxFile() (string, error) {
	home := cmdutil.HomeDir()
	if home == "" {
		return "", errors.New("HOME or USERPROFILE environment variable not set")
	}
	return filepath.Join(home, ".vault", "vaultctx"), nil
}

// readLastContext returns the saved previous context
// if the state file exists, otherwise returns "".
func readLastContext(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if os.IsNotExist(err) {
		return "", nil
	}
	return string(b), err
}

// writeLastContext saves the specified value to the state file.
// It creates missing parent directories.
func writeLastContext(path, value string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return errors.Wrap(err, "failed to create parent directories")
	}
	return ioutil.WriteFile(path, []byte(value), 0644)
}
