package builtins

import "fmt"
import "os"
import "strings"

var aliases = make(map[string]string)

func Alias(operands []string) {
    // Either list all aliases, add one, or print a single alias.
    if len(operands) > 0 {
        alias := strings.SplitN(operands[0], "=", 2)

        // Add an alias.
        if len(alias) > 1 {
            aliases[alias[0]] = alias[1]
        } else {
            if value, ok := LookUpAlias(alias[0]); !ok {
                fmt.Fprintf(os.Stderr, "swish: '%s' alias not found\n", alias[0])
            } else {
                fmt.Printf("alias '%s'='%s'\n", alias[0], value)
            }
        }
    } else {
        for alias, value := range aliases {
            fmt.Printf("alias '%s'='%s'\n", alias, value)
        }
    }
}

func LookUpAlias(name string) (string, bool) {
    alias, ok := aliases[name]
    return alias, ok
}
