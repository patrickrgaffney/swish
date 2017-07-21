// Swish: the james-harden-from-the-free-throw-line shell
package main

import "bufio"
import "fmt"
import "os"
import "os/exec"
import "strings"
import "syscall"

// Write the (hardcoded) prompt to the screen.
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

func execute(cmd []string) {
	// Make sure the cmd exists in the $PATH.
	binary, pathErr := exec.LookPath(cmd[0])
	if pathErr != nil {
		panic(pathErr)
	}

	// Fork and execute our program.
	execErr := syscall.Exec(binary, cmd, os.Environ())
	if execErr != nil {
		panic(execErr)
	}
}

// The main loop of the swish shell.
func swishLoop() {
	stdin := bufio.NewReader(os.Stdin)

	for {
		prompt()
		cmd := readInput(stdin)

		fmt.Println("--user input:", cmd)
		execute(cmd)
	}
}

func main() {

	swishLoop()
}
