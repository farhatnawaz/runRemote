package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func runCmdRemotely(address string, command string) {

	cmd := exec.Command("ssh", address, command) // exec.command would ask for the password anyway
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(&out)
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	// I know I had the choice to hard code the values instead of taking input from the user. I just thought
	// it would seem more convenient to do it this way.

	fmt.Println("Please enter the address (format user@address) : ")
	address, _ := reader.ReadString('\n')

	fmt.Println("Please enter the command: ")
	cmd, _ := reader.ReadString('\n')

	runCmdRemotely(address, cmd)

}
