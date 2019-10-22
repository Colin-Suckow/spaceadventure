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

	fmt.Print("What is your name? ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimRight(name, "\n")

	fmt.Println("Nice to meet you " + name + "!")
	fmt.Println("Let's go on an adventure!")
	fmt.Println("Shall I randomly choose a planet for you? (Y or N)")

	randomChoice := ""

	for {
		reader := bufio.NewReader(os.Stdin)
		randomChoice, _ = reader.ReadString('\n')
		randomChoice = strings.ToLower(strings.TrimRight(randomChoice, "\n"))
		if randomChoice == "y" || randomChoice == "n" {
			break
		}
		fmt.Println("Invalid input. Try again (Y or N)")

	}

	if randomChoice == "y" {
		rand.Seed(time.Now().UnixNano())
		randomNumber := rand.Intn(len(system.Planets))
		planet := system.Planets[randomNumber]
		fmt.Println("Traveling to " + planet.Name)
		fmt.Println("Arrived at " + planet.Name + ". " + planet.Description)
	} else {
		reader := bufio.NewReader(os.Stdin)
		planetChoice, _ := reader.ReadString('\n')
		planetChoice = strings.TrimRight(planetChoice, "\n")
		planet := FindPlanet(system.Planets, planetChoice)
		fmt.Println("Traveling to " + planet.Name)
		fmt.Println("Arrived at " + planet.Name + ". " + planet.Description)

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
