#compdef vaultns kns=vaultns
_arguments "1: :(- $(vaultctl get namespaces -o=jsonpath='{range .items[*].metadata.name}{@}{"\n"}{end}'))"
