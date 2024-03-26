# vaultctx

function __fish_vaultctx_arg_number -a number
    set -l cmd (commandline -opc)
    test (count $cmd) -eq $number
end

complete -f -c vaultctx
complete -f -x -c vaultctx -n '__fish_vaultctx_arg_number 1' -a "(vaultctl config get-contexts --output='name')"
complete -f -x -c vaultctx -n '__fish_vaultctx_arg_number 1' -a "-" -d "switch to the previous namespace in this context"
