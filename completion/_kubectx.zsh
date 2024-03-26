#compdef vaultctx kctx=vaultctx

local vaultCTX="${HOME}/.vault/vaultctx"
PREV=""

local context_array=("${(@f)$(vaultctl config get-contexts --output='name')}")
local all_contexts=(\'${^context_array}\')

if [ -f "$vaultCTX" ]; then
    # show '-' only if there's a saved previous context
    local PREV=$(cat "${vaultCTX}")

    _arguments \
      "-d:*: :(${all_contexts})" \
      "(- *): :(- ${all_contexts})"
else
    _arguments \
      "-d:*: :(${all_contexts})" \
      "(- *): :(${all_contexts})"
fi
