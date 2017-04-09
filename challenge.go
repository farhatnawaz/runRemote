package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

func runCmdRemotely(user string, address string, password string, cmd string) (string, error) {

	// Authentication
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
	}
	// Connect
	client, err := ssh.Dial("tcp", address+":22", config)
	if err != nil {
		return "", err
	}
	// Create a session.
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	var a bytes.Buffer
	session.Stdout = &a // to get output

	// Run the command
	err = session.Run(cmd)
	return a.String(), err
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	// I know I had the choice to hard code the values instead of taking input from the user. I just thought
	// it would seem more convenient to do it this way.

	fmt.Println("Please enter the username: ")
	user, _ := reader.ReadString('\n')

	fmt.Println("Please enter the address: ")
	address, _ := reader.ReadString('\n')

	fmt.Println("Please enter the password: ")
	password, _ := reader.ReadString('\n')

	fmt.Println("Please enter the command: ")
	cmd, _ := reader.ReadString('\n')

	fmt.Println(runCmdRemotely(user, address, password, cmd))

}
