package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	n := -1
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	studentName, _ := reader.ReadString('\n')

	fmt.Printf("Hello %s Enter Number of Subjects: ", studentName)
	for {
		_, err := fmt.Scanln(&n)
		if err != nil {
			fmt.Println("Invalid input, please enter a number.")
		} else if n <= 0 {
			fmt.Println("Number must be a positive number.")
		} else {
			break
		}
	}
	m := make(map[string]int)
	var grade int
	var name string
	sums := 0
	for i := 0; i < n; i++ {
		fmt.Printf("Enter your %dth Subject name and grade: \n", i+1)
		for {
			_, err := fmt.Scan(&name, &grade)

			if err != nil {
				fmt.Println("Invalid input, please enter a number.")
			} else {
				break
			}
		}
		m[name] = grade
		sums += grade
	}
	fmt.Printf("Your Name is: %s\n", studentName)
	fmt.Print("subjects are: \n")

	for i, val := range m {
		fmt.Printf("%s: %d\n", i, val)
	}
	fmt.Printf("\nAverage: %d\n", sums/n)

}
