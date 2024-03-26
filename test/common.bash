#!/usr/bin/env bats

# bats setup function
setup() {
  TEMP_HOME="$(mktemp -d)"
  export TEMP_HOME
  export HOME=$TEMP_HOME
  export vaultCONFIG="${TEMP_HOME}/config"
}

# bats teardown function
teardown() {
  rm -rf "$TEMP_HOME"
}

use_config() {
  cp "$BATS_TEST_DIRNAME/testdata/$1" $vaultCONFIG
}

# wrappers around "vaultctl config" command

get_namespace() {
  vaultctl config view -o=jsonpath="{.contexts[?(@.name==\"$(get_context)\")].context.namespace}"
}

get_context() {
  vaultctl config current-context
}

switch_context() {
  vaultctl config use-context "${1}"
}
