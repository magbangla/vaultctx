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
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"

	"github.com/ahmetb/vaultctx/internal/cmdutil"
	"github.com/ahmetb/vaultctx/internal/env"
	"github.com/ahmetb/vaultctx/internal/printer"
	"github.com/ahmetb/vaultctx/internal/vaultconfig"
)

type InteractiveSwitchOp struct {
	SelfCmd string
}

type InteractiveDeleteOp struct {
	SelfCmd string
}

func (op InteractiveSwitchOp) Run(_, stderr io.Writer) error {
	// parse vaultconfig just to see if it can be loaded
	kc := new(vaultconfig.vaultconfig).WithLoader(vaultconfig.DefaultLoader)
	if err := kc.Parse(); err != nil {
		if cmdutil.IsNotFoundErr(err) {
			printer.Warning(stderr, "vaultconfig file not found")
			return nil
		}
		return errors.Wrap(err, "vaultconfig error")
	}
	kc.Close()

	cmd := exec.Command("fzf", "--ansi", "--no-preview")
	var out bytes.Buffer
	cmd.Stdin = os.Stdin
	cmd.Stderr = stderr
	cmd.Stdout = &out

	cmd.Env = append(os.Environ(),
		fmt.Sprintf("FZF_DEFAULT_COMMAND=%s", op.SelfCmd),
		fmt.Sprintf("%s=1", env.EnvForceColor))
	if err := cmd.Run(); err != nil {
		if _, ok := err.(*exec.ExitError); !ok {
			return err
		}
	}
	choice := strings.TrimSpace(out.String())
	if choice == "" {
		return errors.New("you did not choose any of the options")
	}
	name, err := switchContext(choice)
	if err != nil {
		return errors.Wrap(err, "failed to switch context")
	}
	printer.Success(stderr, "Switched to context \"%s\".", printer.SuccessColor.Sprint(name))
	return nil
}

func (op InteractiveDeleteOp) Run(_, stderr io.Writer) error {
	// parse vaultconfig just to see if it can be loaded
	kc := new(vaultconfig.vaultconfig).WithLoader(vaultconfig.DefaultLoader)
	if err := kc.Parse(); err != nil {
		if cmdutil.IsNotFoundErr(err) {
			printer.Warning(stderr, "vaultconfig file not found")
			return nil
		}
		return errors.Wrap(err, "vaultconfig error")
	}
	kc.Close()

	if len(kc.ContextNames()) == 0 {
		return errors.New("no contexts found in config")
	}

	cmd := exec.Command("fzf", "--ansi", "--no-preview")
	var out bytes.Buffer
	cmd.Stdin = os.Stdin
	cmd.Stderr = stderr
	cmd.Stdout = &out

	cmd.Env = append(os.Environ(),
		fmt.Sprintf("FZF_DEFAULT_COMMAND=%s", op.SelfCmd),
		fmt.Sprintf("%s=1", env.EnvForceColor))
	if err := cmd.Run(); err != nil {
		if _, ok := err.(*exec.ExitError); !ok {
			return err
		}
	}

	choice := strings.TrimSpace(out.String())
	if choice == "" {
		return errors.New("you did not choose any of the options")
	}

	name, wasActiveContext, err := deleteContext(choice)
	if err != nil {
		return errors.Wrap(err, "failed to delete context")
	}

	if wasActiveContext {
		printer.Warning(stderr, "You deleted the current context. Use \"%s\" to select a new context.",
			selfName())
	}

	printer.Success(stderr, `Deleted context %s.`, printer.SuccessColor.Sprint(name))

	return nil
}