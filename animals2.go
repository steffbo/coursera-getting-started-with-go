/*
Write a program which allows the user to create a set of animals and to get information about those animals.
Each animal has a name and can be either a cow, bird, or snake.
With each command, the user can either create a new animal of one of the three types,
or the user can request information about an animal that he/she has already created.
Each animal has a unique name, defined by the user. Note that the user can define animals of a chosen type,
but the types of animals are restricted to either cow, bird, or snake.
The following table contains the three types of animals and their associated data.

Animal | Food eaten | Locomotion method | Spoken sound
cow    | grass      | walk				| moo
bird   | worms		| fly			    | peep
snake  | mice       | slither           | hsss

Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line.
Your program should continue in this loop forever. Every command from the user must be either a “newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings. The first string is “newanimal”.
The second string is an arbitrary string which will be the name of the new animal.
The third string is the type of the new animal, either “cow”, “bird”, or “snake”.
Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings. The first string is “query”.
The second string is the name of the animal. The third string is the name of the information
requested about the animal, either “eat”, “move”, or “speak”. Your program should process each
query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal. Specifically,
the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values.
The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak()
method should print the animal’s spoken sound. Define three types Cow, Bird, and Snake. For each of these three types,
define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface.
When the user creates an animal, create an object of the appropriate type. Your program should call the appropriate
method when the user issues a query command.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	animals := map[string]Animal{}

	for {
		fmt.Println("Type command (\"newanimal\" or \"query\")")
		fmt.Print("> ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		split := strings.Split(text, " ")
		command := split[0]

		if command == "newanimal" {
			name := split[1]
			kind := split[2]
			switch strings.ToLower(kind) {
			case "cow":
				animals[name] = Cow{}
				fmt.Println("Created it!")
			case "bird":
				animals[name] = Bird{}
				fmt.Println("Created it!")
			case "snake":
				animals[name] = Snake{}
				fmt.Println("Created it!")
			default:
				fmt.Println("Invalid kind")
			}
			continue
		}

		if command == "query" {
			name := split[1]
			if animals[name] == nil {
				fmt.Println("Name unknown")
				continue
			}
			info := split[2]
			animal := animals[name]
			switch strings.ToLower(info) {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Invalid info")
			}
		}
	}
}

type Animal interface {
	Eat()
	Move()
	Speak()
}
type Cow struct{}

func (c Cow) Eat() {
	fmt.Println("grass")
}
func (c Cow) Move() {
	fmt.Println("walk")
}
func (c Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (b Bird) Eat() {
	fmt.Println("worms")
}
func (b Bird) Move() {
	fmt.Println("fly")
}
func (b Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

func (s Snake) Eat() {
	fmt.Println("mice")
}
func (s Snake) Move() {
	fmt.Println("slither")
}
func (s Snake) Speak() {
	fmt.Println("hsss")
}
