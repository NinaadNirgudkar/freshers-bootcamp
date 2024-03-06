package daytwo

import (
	"fmt"
	"math/rand"
	"sync"
)

var wa = sync.WaitGroup{}

func Ratings() {
	students := make([]int, 200)

	for i := 0; i < 200; i++ {
		wa.Add(1)
		go submitRating(students, i)
	}

	result := []int{0, 0}

	go func(result []int) {
		for i := range students {
			result[0] += students[i]
			result[1]++
		}
		fmt.Println("lol", result)
	}(result)

	// Calculate and print the average rating
	wa.Add(1)
	go func(result []int) {
		fmt.Println(result)
		if result[1] > 0 {
			averageRating := float64(result[1]) / float64(result[1])
			fmt.Printf("Average Rating: %.2f\n", averageRating)
		} else {
			fmt.Println("No ratings received.")
		}
		wa.Done()
	}(result)

	wa.Wait()
}

func submitRating(students []int, i int) {
	defer wa.Done()
	students[i] = rand.Intn(10)
}
