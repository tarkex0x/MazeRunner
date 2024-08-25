package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func PrintMessage(message string) {
	fmt.Println(message)
}

func PrintGameState(playerX, playerY, goalX, goalY int, hasWon bool) {
	fmt.Printf("Player Position: X: %d, Y: %d\n", playerX, playerY)
	fmt.Printf("Goal Position: X: %d, Y: %d\n", goalX, goalY)
	if hasWon {
		fmt.Println("Congratulations! You've reached the goal!")
	} else {
		fmt.Println("Keep going! You haven't reached the goal yet.")
	}
}

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}