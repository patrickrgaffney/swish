// Swish: the james-harden-from-the-free-throw-line shell
package main

import "bufio"
import "fmt"
import "os"
import "os/exec"
import "strings"

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

func execute(program []string) {
	// Make sure the cmd exists in the $PATH.
	_, pathErr := exec.LookPath(program[0])
	if pathErr != nil {
		fmt.Println(pathErr.Error())
		return
	}

	// Initialize and execute the command.
	cmd := exec.Command(program[0], program[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}

	// Fork and execute our program.
	// execErr := syscall.Exec(binary, cmd, os.Environ())
	// if execErr != nil {
	// 	panic(execErr)
	// }
}

// The main loop of the swish shell.
func swishLoop() {
	stdin := bufio.NewReader(os.Stdin)

	for {
		prompt()
		cmd := readInput(stdin)

		fmt.Println("--swish:", cmd)

		if cmd[0] == "cd" {
			change_dir(cmd[1])
		} else {
			execute(cmd)
		}
	}
}

func change_dir(dir string) {
	if e := os.Chdir(dir); e != nil {
		fmt.Println(e.Error())
	}
}

func main() {

	swishLoop()
}
