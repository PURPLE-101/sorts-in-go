package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(numbers[]int) []int{

	arrayLength := len(numbers)

	for i := 0; i < arrayLength; i++ {
		for j := 1; j < arrayLength-1; j++ {
			if(numbers[j-1] > numbers[j]) {
				numbers[j], numbers[j-1] = numbers[j-1], numbers[j]
			}
		}
	}

	return numbers
}

func secondBubbleSort(numbers[]int) []int{

	arrayLength := len(numbers)
	swapped := true

	for swapped {
		swapped = false
		for i := 1; i < arrayLength-1; i++ {
			if(numbers[i-1] > numbers[i]) {
				numbers[i], numbers[i-1] = numbers[i-1], numbers[i]
				swapped = true
			}
		}
	}

	return numbers
}

func generateRandomTableOf(size int) []int{
	
	var numbers = make([]int,size)
	var max = 100
	var min = 1

	for i := 0; i < size; i++ {
		numbers[i] = rand.Intn(max-min) + min
	}

	return numbers
}

func startAndTimeBubbleSort(optionNumber int)  {
	var numbers = generateRandomTableOf(10000)

	start := time.Now()
	
	validateAndExecute(optionNumber, numbers)

	elapsed := time.Since(start)

	fmt.Println("Algorithm ", optionNumber, " took " , elapsed , " to complete.")
}

func validateAndExecute(optionNumber int, numbers[] int) {
	if(optionNumber == 1){
		bubbleSort(numbers)
	}else if(optionNumber == 2){
		secondBubbleSort(numbers)
	}else {
		fmt.Println("Invalid choice")
	}

	return
}

func main(){
	startAndTimeBubbleSort(1)
	
	fmt.Println("----------------------")

	startAndTimeBubbleSort(2)
}
