package main

import (
	"bufio"
	"encoding/json"
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
	// Open our jsonFile
	jsonFile, err := os.Open("/home/colin/Downloads/planetarySystem.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var system System

	json.Unmarshal(byteValue, &system)

	fmt.Println("Welcome to the " + system.Name)
	fmt.Println("There are " + strconv.Itoa(len(system.Planets)) + " planets")

	name := GetInput("What is your name?")

	fmt.Println("Nice to meet you " + name + "!")
	fmt.Println("Let's go on an adventure!")
	randomChoice := ""

	for {
		randomChoice = strings.ToLower(GetInput("Shall I randomly choose a planet for you? (Y or N)"))
		if randomChoice == "y" || randomChoice == "n" {
			break
		}
		fmt.Println("Invalid input")
	}

	if randomChoice == "y" {
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

func GetInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt + " ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimRight(input, "\n")
	return input
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
