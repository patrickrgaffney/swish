// Swish: the james-harden-from-the-free-throw-line shell
package main

import "bufio"
import "fmt"
import "os"

func prompt() {
	prompt := "pat@swish$ "

	fmt.Print(prompt)
}

func readInput(stdin *bufio.Reader) {
	input, _ := stdin.ReadString('\n')

	fmt.Printf("user input: %s", input)
}

func swishLoop() {
	stdin := bufio.NewReader(os.Stdin)

	for {
		prompt()
		readInput(stdin)
	}
}

func main() {

	swishLoop()
}
