package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"wildberries-l2/os/commands"
)

func RunShell() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)

		// Проверка на команду выхода
		if input == "\\quit" {
			break
		}

		CheckCommand(input)

	}
}

func CheckCommand(input string) {
	if strings.HasPrefix(input, "cd") {
		commands.Cd(input)
	} else if input == "pwd" {
		commands.Pwd()
	} else if strings.HasPrefix(input, "echo") {
		commands.Echo(input)
	} else if strings.HasPrefix(input, "kill") {
		commands.Kill(input)
	} else if input == "ps" {
		commands.Ps()
	} else {
		commands.External(input)
	}
}
