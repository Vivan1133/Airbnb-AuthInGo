package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading env vars")
	}

}

func GetString(key string, fallback string) string {

	val, found := os.LookupEnv(key)

	if !found {
		return fallback
	}

	return val

}

func GetInt(key string, fallback int) int {

	val, found := os.LookupEnv(key)

	if !found {
		return fallback
	}

	intVal, err := strconv.Atoi(val)

	if err != nil {
		fmt.Println("error in fetching env vars")
		return fallback
	}

	return intVal

}

func GetBool(key string, fallback bool) bool {

	val, found := os.LookupEnv(key)

	if !found {
		return fallback
	}

	boolVal, err := strconv.ParseBool(val)

	if err != nil {
		return fallback
	}

	return boolVal

}