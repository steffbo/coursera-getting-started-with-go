/*
Write a program which allows the user to get information about a predefined set of animals.
Three animals are predefined, cow, bird, and snake. Each animal can eat, move, and speak.
The user can issue a request to find out one of three things about an animal:
1) the food that it eats, 2) its method of locomotion, and 3) the sound it makes when it speaks.
The following table contains the three animals and their associated data which should be hard-coded into your program.

Animal | Food eaten | Locomotion method | Spoken sound
cow    | grass      | walk				| moo
bird   | worms		| fly			    | peep
snake  | mice       | slither           | hsss

Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
Your program accepts one request at a time from the user, prints out the answer to the request,
and prints out a new prompt. Your program should continue in this loop forever.
Every request from the user must be a single line containing 2 strings.
The first string is the name of an animal, either “cow”, “bird”, or “snake”.
The second string is the name of the information requested about the animal,
either “eat”, “move”, or “speak”. Your program should process each request by printing out the requested data.

You will need a data structure to hold the information about each animal.
Make a type called Animal which is a struct containing three fields:food, locomotion, and noise,
all of which are strings. Make three methods called Eat(), Move(), and Speak().
The receiver type of all of your methods should be your Animal type.
The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion,
and the Speak() method should print the animal’s spoken sound.
Your program should call the appropriate method when the user makes a request.

*/

package main

//import (
//	"bufio"
//	"fmt"
//	"math/rand"
//	"os"
//	"strings"
//	"time"
//)
//
//type Animal struct {
//	food       string
//	locomotion string
//	noise      string
//}
//
//func (a *Animal) Eat() string {
//	return a.food
//}
//
//func (a *Animal) Move() string {
//	return a.locomotion
//}
//
//func (a *Animal) Speak() string {
//	return a.noise
//}
//
//func main() {
//
//	animals := map[string]*Animal{
//		"cow":   {food: "grass", locomotion: "walk", noise: "moo"},
//		"bird":  {food: "worms", locomotion: "fly", noise: "peep"},
//		"snake": {food: "mice", locomotion: "slither", noise: "hsss"},
//	}
//
//	for {
//		randomRequest := GetRandomRequest(animals)
//		fmt.Println("Make a request like (" + randomRequest + ")")
//
//		scanner := bufio.NewScanner(os.Stdin)
//		scanner.Scan()
//		text := scanner.Text()
//
//		// use default (random suggestion)
//		if len(text) == 0 {
//			split := strings.Split(randomRequest, " ")
//			animal := split[0]
//			action := split[1]
//			AnimalAction(animals, animal, action)
//			continue
//		}
//
//		split := strings.Split(text, " ")
//		if len(split) != 2 {
//			fmt.Println("Invalid input, try again")
//			continue
//		}
//
//		animal := strings.ToLower(split[0])
//		validAnimal := animal == "cow" || animal == "bird" || animal == "snake"
//		if !validAnimal {
//			fmt.Println("Invalid animal, try again")
//			continue
//		}
//		action := strings.ToLower(split[1])
//
//		AnimalAction(animals, animal, action)
//	}
//}
//
//func GetRandomRequest(animals map[string]*Animal) string {
//	// Generate a random index within the range of the map's length
//	rand.Seed(time.Now().UnixNano())
//	randomIndex := rand.Intn(len(animals))
//
//	// Iterate over the map to return the key at the random index
//	i := 0
//	var randomAnimal string
//	for k := range animals {
//		if i == randomIndex {
//			randomAnimal = k
//			break
//		}
//		i++
//	}
//
//	// Select a random field based on the random key
//	switch rand.Intn(3) {
//	case 0:
//		return randomAnimal + " eat"
//	case 1:
//		return randomAnimal + " move"
//	case 2:
//		return randomAnimal + " speak"
//	default:
//		return ""
//	}
//}
//
//func AnimalAction(animals map[string]*Animal, animal string, action string) {
//	requestedAnimal := animals[animal]
//	switch action {
//	case "eat":
//		fmt.Println(animal + " eats " + requestedAnimal.Eat())
//	case "move":
//		fmt.Println(animal + "'s move is to " + requestedAnimal.Move())
//	case "speak":
//		fmt.Println(animal + " says \"" + requestedAnimal.Speak() + "\"!")
//	default:
//		fmt.Println("invalid request")
//	}
//}
