# vaultns

function __fish_vaultns_arg_number -a number
    set -l cmd (commandline -opc)
    test (count $cmd) -eq $number
end

complete -f -c vaultns
complete -f -x -c vaultns -n '__fish_vaultns_arg_number 1' -a "(vaultctl get ns -o=custom-columns=NAME:.metadata.name --no-headers)"
complete -f -x -c vaultns -n '__fish_vaultns_arg_number 1' -a "-" -d "switch to the previous namespace in this context"
complete -f -x -c vaultns -n '__fish_vaultns_arg_number 1' -s c -l current -d "show the current namespace"
complete -f -x -c vaultns -n '__fish_vaultns_arg_number 1' -s h -l help -d "show the help message"
