/*
Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a
function called BubbleSort() which
takes a slice of integers as an argument and returns nothing.
The BubbleSort() function should modify the slice so that the elements are in sorted
order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.

Submit your Go program source code.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("BubbleSort! Enter up to 10 numbers")
	numbers := GetInputs()
	BubbleSort(numbers)
	Output(numbers)
}

func GetInputs() []int {
	var result []int

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	nums := strings.Split(text, " ")

	for i, num := range nums {
		if i >= 10 {
			break
		}
		n, _ := strconv.Atoi(num)
		result = append(result, n)
	}

	return result
}

func BubbleSort(numbers []int) {
	swapped := false

	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			Swap(numbers, i)
			swapped = true
		}
	}
	if swapped {
		BubbleSort(numbers)
	}
}

func Swap(numbers []int, index int) {
	temp := numbers[index]
	numbers[index] = numbers[index+1]
	numbers[index+1] = temp
}

func Output(numbers []int) {
	builder := strings.Builder{}
	for _, num := range numbers {
		builder.WriteString(strconv.Itoa(num) + " ")
	}
	fmt.Println(builder.String())
}
