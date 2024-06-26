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
	"fmt"
	"io"

	"github.com/magbangla/vaultctx/internal/vaultconfig"
	"github.com/pkg/errors"
)

// CurrentOp prints the current context
type CurrentOp struct{}

func (_op CurrentOp) Run(stdout, _ io.Writer) error {
	kc := new(vaultconfig.Vaultconfig).WithLoader(vaultconfig.DefaultLoader)
	defer kc.Close()
	if err := kc.Parse(); err != nil {
		return errors.Wrap(err, "vaultconfig error")
	}

	v := kc.GetCurrentContext()
	if v == "" {
		return errors.New("current-context is not set")
	}
	_, err := fmt.Fprintln(stdout, v)
	return errors.Wrap(err, "write error")
}
