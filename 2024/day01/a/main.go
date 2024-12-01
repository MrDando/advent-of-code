package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	arr1, arr2 := readData()
	resultArr := make([]int, 0, 1000)

	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]
	})

	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})

	for i := 0; i < 1000; i++ {
		dif := Abs(arr1[i] - arr2[i])
		resultArr = append(resultArr, dif)
	}

	result := 0
	for i := 0; i < len(resultArr); i++ {
		result += resultArr[i]
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
		// Split on whitespace (handles both spaces and tabs)
		values := strings.Fields(line)

		// Convert strings to integers
		val1, _ := strconv.Atoi(values[0])
		val2, _ := strconv.Atoi(values[1])

		array1 = append(array1, val1)
		array2 = append(array2, val2)
	}

	return array1, array2
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
