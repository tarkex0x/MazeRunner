package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Command int

const (
	MoveUp Command = iota
	MoveDown
	MoveLeft
	MoveRight
	Invalid
)

func convertInputToCommand(input string) Command {
	switch strings.ToLower(input) {
	case "w":
		return MoveUp
	case "s":
		return MoveDown
	case "a":
		return MoveLeft
	case "d":
		return MoveRight
	default:
		return Invalid
	}
}

func readUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err // Return the error to be handled by the caller
	}
	return strings.TrimSpace(input), nil
}

func main() {
	for {
		userInput, err := readUserInput("Enter move (W/A/S/D): ")
		if err != nil {
			fmt.Println("Error reading input:", err) //Handling error here
			continue
		}

		command := convertInputToCommand(userInput)
		switch command {
		case MoveUp:
			fmt.Println("Moving up.")
		case MoveDown:
			fmt.Println("Moving down.")
		case MoveLeft:
			fmt.Println("Moving left.")
		case MoveRight:
			fmt.Println("Moving right.")
		case Invalid:
			fmt.Println("Invalid command, please enter W, A, S, or D.")
		}
	}
}