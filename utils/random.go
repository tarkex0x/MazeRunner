package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

func getRandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func getEnvAsInt(name string, defaultValue int) int {
	valueStr := os.Getenv(name)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func initRandSeed() {
	seedValue := getEnvAsInt("MAZE_RUNNER_SEED", int(time.Now().Unix()))
	rand.Seed(int64(seedValue))
}

func main() {
	initRandSeed()

	min := 1
	max := 10
	randomNumber := getRandomInt(min, max)
}