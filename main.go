// Swish: the james-harden-from-the-free-throw-line shell
package main

import "bufio"
import "fmt"
import "os"
import "strings"

// Write the prompt to the screen.
func prompt() {
	prompt := "pat@swish$ "

	fmt.Print(prompt)
}

// Read user input until newline is entered.
func readInput(stdin *bufio.Reader) []string {
	input, _ := stdin.ReadString('\n')

	// Remove the trailing newline.
	return strings.Split(input[:len(input) - 1], " ")
}

// The main loop of the swish shell.
func swishLoop() {
	stdin := bufio.NewReader(os.Stdin)

	for {
		prompt()
		command := readInput(stdin)

		fmt.Println("--user input:", command)
	}
}

func main() {

	swishLoop()
}
