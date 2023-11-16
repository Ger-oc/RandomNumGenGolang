package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func generateRandomNumbers(n, min, max int) []int {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = rand.Intn(max-min+1) + min
	}
	return numbers
}

func writeNumbersToCSV(filename string, numbers []int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, num := range numbers {
		err := writer.Write([]string{strconv.Itoa(num)})
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	// Change the filename, number of random numbers, and range as needed
	filename := "bonus100k.csv"
	numNumbers := 100000
	minRange := -99999
	maxRange := 99999

	// Generate random numbers in the specified range
	randomNumbers := generateRandomNumbers(numNumbers, minRange, maxRange)

	// Write numbers to CSV file
	err := writeNumbersToCSV(filename, randomNumbers)
	if err != nil {
		log.Fatalf("Error writing numbers to CSV: %v", err)
	}

	fmt.Printf("Random numbers written to %s\n", filename)
}
