package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	studentName, _ := reader.ReadString('\n')
	studentName = strings.TrimSpace(studentName)
	for {
		fmt.Print("Enter the number of subjects: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var err error
		n, err = strconv.Atoi(input)
		if err != nil || n <= 0 {
			fmt.Println("Invalid input, please enter a positive number.")
		} else {
			break
		}
	}
	m := make(map[string]float64)
	sums := 0.0
	for i := 0; i < n; i++ {
		fmt.Printf("Enter the name of subject %d: ", i+1)
		subject, _ := reader.ReadString('\n')
		subject = strings.TrimSpace(subject)
		var grade float64
		for {
			fmt.Printf("Enter the grade for %s: ", subject)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			var err error
			grade, err = strconv.ParseFloat(input, 64)
			if err != nil || grade < 0 || grade > 100 {
				fmt.Println("Invalid grade, please enter a number between 0 and 100.")
			} else {
				break
			}
		}
		m[subject] = grade
		sums = sums + grade
	}
	fmt.Printf("Your Name is: %s\n", studentName)
	fmt.Print("subjects are: \n")

	for i, val := range m {
		fmt.Printf("%s: %.3f\n", i, val)
	}
	av := sums / float64(n)
	fmt.Print(av, n)
	fmt.Printf("\nAverage: %.3f\n", av)

}
