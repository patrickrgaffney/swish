package builtins

import "fmt"
import "os"

// TODO: Look into returning the cached value in $PWD if it exists,
// otherwise make the os.Getwd() call --- which may take a long time.
func Pwd() {
	wd, err := os.Getwd()
	if err != nil {
		// err could be *PathError, *SyscallError
		if perr, ok := err.(*os.PathError); ok {
			fmt.Fprintf(os.Stderr, "pwd: %s: %s\n", perr.Err.Error(), perr.Path)
		} else if serr, ok := err.(*os.SyscallError); ok {
			fmt.Fprintf(os.Stderr, "pwd: %s: %s\n", serr.Err.Error(), serr.Syscall)
		} else {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		}
	}

	fmt.Println(wd)
}
