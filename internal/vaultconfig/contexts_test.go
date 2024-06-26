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

	"github.com/magbangla/vaultctx/internal/testutil"
)

func Testvaultconfig_ContextNames(t *testing.T) {
	tl := WithMockvaultconfigLoader(
		testutil.KC().WithCtxs(
			testutil.Ctx("abc"),
			testutil.Ctx("def"),
			testutil.Ctx("ghi")).Set("field1", map[string]string{"bar": "zoo"}).ToYAML(t))
	kc := new(Vaultconfig).WithLoader(tl)
	if err := kc.Parse(); err != nil {
		t.Fatal(err)
	}

	ctx := kc.ContextNames()
	expected := []string{"abc", "def", "ghi"}
	if diff := cmp.Diff(expected, ctx); diff != "" {
		t.Fatalf("%s", diff)
	}
}

func Testvaultconfig_ContextNames_noContextsEntry(t *testing.T) {
	tl := WithMockvaultconfigLoader(`a: b`)
	kc := new(Vaultconfig).WithLoader(tl)
	if err := kc.Parse(); err != nil {
		t.Fatal(err)
	}
	ctx := kc.ContextNames()
	var expected []string = nil
	if diff := cmp.Diff(expected, ctx); diff != "" {
		t.Fatalf("%s", diff)
	}
}

func Testvaultconfig_ContextNames_nonArrayContextsEntry(t *testing.T) {
	tl := WithMockvaultconfigLoader(`contexts: "hello"`)
	kc := new(Vaultconfig).WithLoader(tl)
	if err := kc.Parse(); err != nil {
		t.Fatal(err)
	}
	ctx := kc.ContextNames()
	var expected []string = nil
	if diff := cmp.Diff(expected, ctx); diff != "" {
		t.Fatalf("%s", diff)
	}
}

func Testvaultconfig_CheckContextExists(t *testing.T) {
	tl := WithMockvaultconfigLoader(
		testutil.KC().WithCtxs(
			testutil.Ctx("c1"),
			testutil.Ctx("c2")).ToYAML(t))

	kc := new(Vaultconfig).WithLoader(tl)
	if err := kc.Parse(); err != nil {
		t.Fatal(err)
	}

	if !kc.ContextExists("c1") {
		t.Fatal("c1 actually exists; reported false")
	}
	if !kc.ContextExists("c2") {
		t.Fatal("c2 actually exists; reported false")
	}
	if kc.ContextExists("c3") {
		t.Fatal("c3 does not exist; but reported true")
	}
}
