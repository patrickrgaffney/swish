// Swish: the james-harden-from-the-free-throw-line shell
package main

import "bufio"
import "fmt"
import "os"
import "os/exec"
import "strings"

import "github.com/patrickrgaffney/swish/builtins"

// Write the (hardcoded) prompt to the screen.
func prompt() {
	prompt := "pat@swish$ "

	fmt.Print(prompt)
}

// Read user input until newline is entered.
func readInput(stdin *bufio.Reader) []string {
	input, _ := stdin.ReadString('\n')

	// Remove the trailing newline.
	return strings.Split(input[:len(input)-1], " ")
}

func parseInput(cmd []string) {
	switch cmd[0] {
	case "cd":
		builtins.Cd(cmd[1])
	case "pwd":
		builtins.Pwd()
	case "exit":
		builtins.Exit(0)
	case "help":
		builtins.Help()
	case "alias":
		if len(cmd) > 1 {
			builtins.Alias(cmd[1:])
		} else {
			builtins.Alias([]string{})
		}
	default:
		execute(cmd)
	}
}

func execute(program []string) {
	// Make sure the cmd exists in the $PATH.
	_, pathErr := exec.LookPath(program[0])
	if pathErr != nil {
		fmt.Fprintf(os.Stderr, "swish: '%s': command not in $PATH\n", program[0])
		return
	}

	// Initialize and execute the command.
	cmd := exec.Command(program[0], program[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		// Only log errors that cannot be promoted to *ExitError.
		if _, ok := err.(*exec.ExitError); !ok {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		}
	}
}

// The main loop of the swish shell.
func swishLoop() {
	stdin := bufio.NewReader(os.Stdin)

	for {
		prompt()
		cmd := readInput(stdin)
		parseInput(cmd)
	}
}

func main() {
	swishLoop()
}
