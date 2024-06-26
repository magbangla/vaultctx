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
	"io"

	"github.com/pkg/errors"

	"github.com/magbangla/vaultctx/internal/printer"
	"github.com/magbangla/vaultctx/internal/vaultconfig"
)

// UnsetOp indicates intention to remove current-context preference.
type UnsetOp struct{}

func (_ UnsetOp) Run(_, stderr io.Writer) error {
	kc := new(vaultconfig.Vaultconfig).WithLoader(vaultconfig.DefaultLoader)
	defer kc.Close()
	if err := kc.Parse(); err != nil {
		return errors.Wrap(err, "vaultconfig error")
	}

	if err := kc.UnsetCurrentContext(); err != nil {
		return errors.Wrap(err, "error while modifying current-context")
	}
	if err := kc.Save(); err != nil {
		return errors.Wrap(err, "failed to save vaultconfig file after modification")
	}

	err := printer.Success(stderr, "Active context unset for vaultctl.")
	return errors.Wrap(err, "write error")
}
