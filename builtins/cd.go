package builtins

import "fmt"
import "os"

func Cd(dir string) {
	if e := os.Chdir(dir); e != nil {
		// Try to cast the error as a *PathError.
		if perr, ok := e.(*os.PathError); ok {
			fmt.Fprintf(os.Stderr, "cd: %s: %s\n", perr.Err.Error(), perr.Path)
		} else {
			fmt.Fprintf(os.Stderr, "%s\n", e.Error())
		}
	}
}
