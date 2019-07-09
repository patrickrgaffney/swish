package builtins

import "fmt"

func Help() {
	const help string = `                         /$$           /$$
                        |__/          | $$
  /$$$$$$$ /$$  /$$  /$$ /$$  /$$$$$$$| $$$$$$$
 /$$_____/| $$ | $$ | $$| $$ /$$_____/| $$__  $$
|  $$$$$$ | $$ | $$ | $$| $$|  $$$$$$ | $$  \ $$
 \____  $$| $$ | $$ | $$| $$ \____  $$| $$  | $$
 /$$$$$$$/|  $$$$$/$$$$/| $$ /$$$$$$$/| $$  | $$
|_______/  \_____/\___/ |__/|_______/ |__/  |__/

swish, version: 0.0.1
the james-harden-from-three shell

usage: swish

builtins:
    cd      change the current working directory
    pwd     print the current working directory
    exit    exit the shell
    help    print this message

Created with no midrange game and golang by 
Pat Gaffney <pat@hypepat.com>
`

	fmt.Print(help)
}
