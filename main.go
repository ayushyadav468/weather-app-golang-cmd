package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"cli-weather-app/apiCall"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	weather_api, exists := os.LookupEnv("WEATHER_API")
	if !exists {
		log.Fatal("No weather api key in env")
	}
	location, err := GetUserInput()
	if err != nil {
		log.Fatal(err)
	}
	apiCall.ApiCalltoOpenWeather(location, weather_api)
}

func GetUserInput() (string, error) {
	fmt.Print("Enter a location: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		return "", err
	}
	return scanner.Text(), nil
}
