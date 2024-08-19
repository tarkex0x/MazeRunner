package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

func getRandomNumberInRange(min, max int) int {
	if min >= max {
		return min
	}
	return rand.Intn(max-min+1) + min
}

func getEnvironmentVariableAsInteger(variableName string, defaultValue int) int {
	envValue := os.Getenv(variableName)
	if intValue, err := strconv.Atoi(envValue); err == nil {
		return intValue
	}
	return defaultValue
}

func initializeRandomSeed() {
	seed := getEnvironmentVariableAsInteger("MAZE_RUNNER_SEED", int(time.Now().Unix()))
	rand.Seed(int64(seed))
}

func main() {
	initializeRandomSeed()

	minValue := 1
	maxValue := 10
	randomValue := getRandomNumberInRange(minValue, maxValue)
	_ = randomValue
}