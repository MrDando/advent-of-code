package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	arr1, arr2 := readData()

	countMap := make(map[int]int)

	for _, num := range arr2 {
		countMap[num]++
	}

	result := 0
	for _, num := range arr1 {
		result += num * countMap[num]
	}

	fmt.Printf("Result: %v\n", result)
}

func readData() ([]int, []int) {
	file, err := os.Open("input.txt")
	defer file.Close()

	array1 := make([]int, 0, 1000)
	array2 := make([]int, 0, 1000)

	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)

		val1, _ := strconv.Atoi(values[0])
		val2, _ := strconv.Atoi(values[1])

		array1 = append(array1, val1)
		array2 = append(array2, val2)
	}

	return array1, array2
}
