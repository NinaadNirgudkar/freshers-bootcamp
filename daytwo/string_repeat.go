package daytwo

import (
	"fmt"
	"sync"
)

var pushCounter = 0
var pullCounter = 0

func StringRepeatCounter() {
	inputStrings := []string{"quickk", "kbrown", "fox", "lazy", "dog"}

	// Create a channel to receive results
	resultChannel := make(chan map[string]int)

	var wg sync.WaitGroup

	// Launch a goroutine for each string in the input
	for _, str := range inputStrings {
		wg.Add(1)
		go func(str string) {
			defer wg.Done()
			processString(str, resultChannel)
		}(str)
	}

	finalOccurrences := make(map[string]int)
	for result := range resultChannel {
		pullCounter++
		fmt.Println("pullCounter", pullCounter)
		for key, value := range result {

			fmt.Println(key, value, finalOccurrences)
			finalOccurrences[string(key)] += value
		}
	}

	close(resultChannel)

	// Print the final frequencies
	fmt.Println(finalOccurrences)
	defer wg.Wait()
}

func processString(str string, resultChannel chan<- map[string]int) {

	occurrences := make(map[string]int)
	for _, c := range str {
		occurrences[string(c)]++ //= (occurrences[string(c)] | 0) + 1
		fmt.Println("pushing", string(c))
	}
	pushCounter++
	fmt.Println(pushCounter)
	resultChannel <- occurrences
}
