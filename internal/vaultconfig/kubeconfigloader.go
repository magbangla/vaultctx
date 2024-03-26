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
	"os"
	"path/filepath"

	"github.com/ahmetb/vaultctx/internal/cmdutil"

	"github.com/pkg/errors"
)

var (
	DefaultLoader Loader = new(StandardvaultconfigLoader)
)

type StandardvaultconfigLoader struct{}

type vaultconfigFile struct{ *os.File }

func (*StandardvaultconfigLoader) Load() ([]ReadWriteResetCloser, error) {
	cfgPath, err := vaultconfigPath()
	if err != nil {
		return nil, errors.Wrap(err, "cannot determine vaultconfig path")
	}

	f, err := os.OpenFile(cfgPath, os.O_RDWR, 0)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.Wrap(err, "vaultconfig file not found")
		}
		return nil, errors.Wrap(err, "failed to open file")
	}

	// TODO we'll return all vaultconfig files when we start implementing multiple vaultconfig support
	return []ReadWriteResetCloser{ReadWriteResetCloser(&vaultconfigFile{f})}, nil
}

func (kf *vaultconfigFile) Reset() error {
	if err := kf.Truncate(0); err != nil {
		return errors.Wrap(err, "failed to truncate file")
	}
	_, err := kf.Seek(0, 0)
	return errors.Wrap(err, "failed to seek in file")
}

func vaultconfigPath() (string, error) {
	// vaultconfig env var
	if v := os.Getenv("VAULTCONFIG"); v != "" {
		list := filepath.SplitList(v)
		if len(list) > 1 {
			// TODO vaultconfig=file1:file2 currently not supported
			return "", errors.New("multiple files in vaultconfig are currently not supported")
		}
		return v, nil
	}

	// default path
	home := cmdutil.HomeDir()
	if home == "" {
		return "", errors.New("HOME or USERPROFILE environment variable not set")
	}
	return filepath.Join(home, ".vault", "config"), nil
}
