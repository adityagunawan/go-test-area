package main

import (
	"fmt"
	"strconv"
)

func main() {
	var j, k int
	// const PIE float32 = 3.14

	j = 8
	k = 3

	var grade string = "B"
	var marks int = 90

	//print string
	// fmt.Println(x + y)

	//get type of variable
	// fmt.Printf("1.) tipe data dari variable x adalah  %T\n", x)

	//convert int to string
	printInteger(j, k)

	//decision
	if j > 5 {
		fmt.Println("variable j lebih besar dari 5")
	}

	fmt.Printf("Your grade is  %s\n", getGrade(marks, grade))
}

func printInteger(j, k int) {
	fmt.Println("hasil j + k = " + strconv.Itoa(j+k))
}

func getGrade(marks int, grade string) string {
	var result string

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	switch {
	case grade == "A":
		result = "Excellent!\n"
	case grade == "B", grade == "C":
		result = "Well done\n"
	case grade == "D":
		result = "You passed\n"
	case grade == "F":
		result = "Better try again\n"
	default:
		result = "Invalid grade\n"
	}

	return result
}
