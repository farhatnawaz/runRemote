package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

//This is the actual function that helps access the remote machine and then execute the given command
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

// This function validates the IP address entered by the user
func isValidAddress(address string) bool {

	var regex = "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])\\.(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])\\.(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])\\.(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"

	s := strings.Split(address, "@")
	match, _ := regexp.MatchString(regex, s[1])
	return match
}

//This function validates the format of the address entered by the user
func isValidFormat(address string) bool {
	return strings.Contains(address, "@")
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	var address string

	// I know I had the choice to hard code the values instead of taking input from the user. I just thought
	// it would seem more convenient to do it this way.

	for {
		fmt.Println("Please enter the address (format user@address) : ")
		address, _ := reader.ReadString('\n')

		if isValidFormat(address) && isValidAddress(address) {
			break
		}

	}

	fmt.Println("Please enter the command: ")
	cmd, _ := reader.ReadString('\n')

	runCmdRemotely(address, cmd)

}
