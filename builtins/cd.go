package builtins

import "fmt"
import "os"

func Cd(dir string) {
    if e := os.Chdir(dir); e != nil {
        fmt.Println(e.Error())
    }
}
