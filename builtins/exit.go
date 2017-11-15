package builtins

import "os"

func Exit(status int) {
	os.Exit(status)
}
