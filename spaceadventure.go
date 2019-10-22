package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type System struct {
	Name    string   `json:"name"`
	Planets []Planet `json:"planets"`
}

type Planet struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func main() {

	if len(os.Args) <= 1 {
		fmt.Println("No file path provided")
		return
	}

	system, err := ReadJson(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Welcome to the " + system.Name)
	fmt.Println("There are " + strconv.Itoa(len(system.Planets)) + " planets")

	name := GetInput("What is your name?")

	fmt.Println("Nice to meet you " + name + "!")
	fmt.Println("Let's go on an adventure!")

	if GetYNResponse("Would you like me to randomly choose a planet for you? (Y/N)", "Sorry, I don't understand.") {
		rand.Seed(time.Now().UnixNano())
		randomNumber := rand.Intn(len(system.Planets))
		planet := system.Planets[randomNumber]
		fmt.Println("Traveling to " + planet.Name)
		fmt.Println("Arrived at " + planet.Name + ". " + planet.Description)
	} else {
		planetChoice := GetInput("What planet would you like to visit?")
		planet := FindPlanet(system.Planets, planetChoice)
		fmt.Println("Traveling to " + planet.Name)
		fmt.Println("Arrived at " + planet.Name + ". " + planet.Description)

	}

}

func ReadJson(path string) (System, error) {
	jsonFile, err := os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {

		return System{}, errors.New("Cannot open file")
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var system System

	json.Unmarshal(byteValue, &system)

	return system, nil
}

func GetInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt + " ")
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	input = strings.TrimRight(input, "\n")
	return input
}

func GetYNResponse(prompt string, errorMessage string) bool {
	input := strings.ToLower(GetInput(prompt))
	if input == "y" {
		return true
	} else if input == "n" {
		return false
	} else {
		fmt.Println(errorMessage)
		return GetYNResponse(prompt, errorMessage)
	}
}

func FindPlanet(planets []Planet, name string) Planet {
	for _, planet := range planets {
		if planet.Name == name {
			return planet
		}
	}
	return Planet{
		Name:        "Null",
		Description: "It is very quiet here...",
	}
}
