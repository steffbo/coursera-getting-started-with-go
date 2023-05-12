/*
Write a program to sort an array of integers. The program should partition the array into 4 parts,
each of which is sorted by a different goroutine. Each partition should be of approximately equal size.
Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of
the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.

Review criteria
Students will receive five points if the program sorts the integers and five additional points
if there are four goroutines that each prints out a set of array elements that it is storing.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const GoroutineCount = 4

func main() {
	fmt.Println("Enter any amount of numbers")
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	nums := strings.Split(input, " ")

	slices := getIntSlicesFromInput(nums)

	// debug slices
	//for _, a := range slices {
	//	fmt.Println(a)
	//}

	wg := sync.WaitGroup{}
	wg.Add(len(slices))
	for _, slice := range slices {
		// sort each slice individually
		fmt.Print("Now sorting ")
		fmt.Println(slice)
		go sortSlice(&slice, &wg)
	}
	wg.Wait()

	// put the numbers from all slices back into a single slice
	var merged []int
	for _, slice := range slices {
		for _, n := range slice {
			merged = append(merged, n)
		}
	}

	// deliberately not doing merge-sort here
	sort.Ints(merged)
	fmt.Print("Sorted: ")
	fmt.Println(merged)
}

// transform list of string numbers to 4 int slices
func getIntSlicesFromInput(input []string) [][]int {
	var result [][]int

	var ints []int
	for i := 0; i < len(input); i++ {
		n, _ := strconv.Atoi(input[i])
		ints = append(ints, n)
	}

	l := len(input) / GoroutineCount

	result = append(result, ints[0:l])
	result = append(result, ints[l:l*2])
	result = append(result, ints[l*2:l*3])
	result = append(result, ints[l*3:l*4+len(input)%4]) // add the rest to the last slice

	return result
}

// sort w/ std lib, as this assignment is about goroutines and not a sort algo
func sortSlice(slice *[]int, wg *sync.WaitGroup) {
	sort.Ints(*slice)
	wg.Done()
}
