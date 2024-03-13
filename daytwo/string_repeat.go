package daytwo

import (
	"fmt"
	"sync"
)

func StringRepeatCounter() {
	inputStrings := []string{"quickk", "kbrown", "fox", "lazy", "dog"}

	// Create a channel to receive results
	resultChannel := make(chan map[string]int, len(inputStrings))

	var wg sync.WaitGroup
	defer wg.Wait()

	// Launch a goroutine for each string in the input
	for _, str := range inputStrings {
		wg.Add(1)
		go func(str string) {
			defer wg.Done()
			processString(str, resultChannel)
		}(str)
	}

	//club all the individual string calculations
	finalOccurrences := make(map[string]int)
	go func() {
		for result := range resultChannel {
			for key, value := range result {
				finalOccurrences[string(key)] += value
			}
		}
		// Print the final frequencies
		fmt.Println(finalOccurrences)
	}()
	wg.Wait()
	close(resultChannel)
}

func processString(str string, resultChannel chan<- map[string]int) {

	occurrences := make(map[string]int)
	for _, c := range str {
		occurrences[string(c)]++
	}
	resultChannel <- occurrences
}
