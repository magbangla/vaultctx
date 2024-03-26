_vault_contexts()
{
  local curr_arg;
  curr_arg=${COMP_WORDS[COMP_CWORD]}
  COMPREPLY=( $(compgen -W "- $(vaultctl config get-contexts --output='name')" -- $curr_arg ) );
}

complete -F _vault_contexts vaultctx kctx
