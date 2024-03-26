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
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/ahmetb/vaultctx/internal/testutil"
)

func Testvaultconfig_DeleteContextEntry_errors(t *testing.T) {
	kc := new(vaultconfig).WithLoader(WithMockvaultconfigLoader(`[1, 2, 3]`))
	_ = kc.Parse()
	err := kc.DeleteContextEntry("foo")
	if err == nil {
		t.Fatal("supposed to fail on non-mapping nodes")
	}

	kc = new(vaultconfig).WithLoader(WithMockvaultconfigLoader(`a: b`))
	_ = kc.Parse()
	err = kc.DeleteContextEntry("foo")
	if err == nil {
		t.Fatal("supposed to fail if contexts key does not exist")
	}

	kc = new(vaultconfig).WithLoader(WithMockvaultconfigLoader(`contexts: "some string"`))
	_ = kc.Parse()
	err = kc.DeleteContextEntry("foo")
	if err == nil {
		t.Fatal("supposed to fail if contexts key is not an array")
	}
}

func Testvaultconfig_DeleteContextEntry(t *testing.T) {
	test := WithMockvaultconfigLoader(
		testutil.KC().WithCtxs(
			testutil.Ctx("c1"),
			testutil.Ctx("c2"),
			testutil.Ctx("c3")).ToYAML(t))
	kc := new(vaultconfig).WithLoader(test)
	if err := kc.Parse(); err != nil {
		t.Fatal(err)
	}
	if err := kc.DeleteContextEntry("c1"); err != nil {
		t.Fatal(err)
	}
	if err := kc.Save(); err != nil {
		t.Fatal(err)
	}

	expected := testutil.KC().WithCtxs(
		testutil.Ctx("c2"),
		testutil.Ctx("c3")).ToYAML(t)
	out := test.Output()
	if diff := cmp.Diff(expected, out); diff != "" {
		t.Fatalf("diff: %s", diff)
	}
}

func Testvaultconfig_ModifyCurrentContext_fieldExists(t *testing.T) {
	test := WithMockvaultconfigLoader(
		testutil.KC().WithCurrentCtx("abc").Set("field1", "value1").ToYAML(t))
	kc := new(vaultconfig).WithLoader(test)
	if err := kc.Parse(); err != nil {
		t.Fatal(err)
	}
	if err := kc.ModifyCurrentContext("foo"); err != nil {
		t.Fatal(err)
	}
	if err := kc.Save(); err != nil {
		t.Fatal(err)
	}

	expected := testutil.KC().WithCurrentCtx("foo").Set("field1", "value1").ToYAML(t)
	out := test.Output()
	if diff := cmp.Diff(expected, out); diff != "" {
		t.Fatalf("diff: %s", diff)
	}
}

func Testvaultconfig_ModifyCurrentContext_fieldMissing(t *testing.T) {
	test := WithMockvaultconfigLoader(`f1: v1`)
	kc := new(vaultconfig).WithLoader(test)
	if err := kc.Parse(); err != nil {
		t.Fatal(err)
	}
	if err := kc.ModifyCurrentContext("foo"); err != nil {
		t.Fatal(err)
	}
	if err := kc.Save(); err != nil {
		t.Fatal(err)
	}

	expected := `f1: v1
current-context: foo
`
	out := test.Output()
	if diff := cmp.Diff(expected, out); diff != "" {
		t.Fatalf("diff: %s", diff)
	}
}

func Testvaultconfig_ModifyContextName_noContextsEntryError(t *testing.T) {
	// no context entries
	test := WithMockvaultconfigLoader(`a: b`)
	kc := new(vaultconfig).WithLoader(test)
	if err := kc.Parse(); err != nil {
		t.Fatal(err)
	}
	if err := kc.ModifyContextName("c1", "c2"); err == nil {
		t.Fatal("was expecting error for no 'contexts' entry; got nil")
	}
}

func Testvaultconfig_ModifyContextName_contextsEntryNotSequenceError(t *testing.T) {
	// no context entries
	test := WithMockvaultconfigLoader(
		`contexts: "hello"`)
	kc := new(vaultconfig).WithLoader(test)
	if err := kc.Parse(); err != nil {
		t.Fatal(err)
	}
	if err := kc.ModifyContextName("c1", "c2"); err == nil {
		t.Fatal("was expecting error for 'context entry not a sequence'; got nil")
	}
}

func Testvaultconfig_ModifyContextName_noChange(t *testing.T) {
	test := WithMockvaultconfigLoader(testutil.KC().WithCtxs(
		testutil.Ctx("c1"),
		testutil.Ctx("c2"),
		testutil.Ctx("c3")).ToYAML(t))
	kc := new(vaultconfig).WithLoader(test)
	if err := kc.Parse(); err != nil {
		t.Fatal(err)
	}
	if err := kc.ModifyContextName("c5", "c6"); err == nil {
		t.Fatal("was expecting error for 'no changes made'")
	}
}

func Testvaultconfig_ModifyContextName(t *testing.T) {
	test := WithMockvaultconfigLoader(testutil.KC().WithCtxs(
		testutil.Ctx("c1"),
		testutil.Ctx("c2"),
		testutil.Ctx("c3")).ToYAML(t))
	kc := new(vaultconfig).WithLoader(test)
	if err := kc.Parse(); err != nil {
		t.Fatal(err)
	}
	if err := kc.ModifyContextName("c1", "ccc"); err != nil {
		t.Fatal(err)
	}
	if err := kc.Save(); err != nil {
		t.Fatal(err)
	}

	expected := testutil.KC().WithCtxs(
		testutil.Ctx("ccc"),
		testutil.Ctx("c2"),
		testutil.Ctx("c3")).ToYAML(t)
	out := test.Output()
	if diff := cmp.Diff(expected, out); diff != "" {
		t.Fatalf("diff: %s", diff)
	}
}
