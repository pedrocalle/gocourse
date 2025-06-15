package basics

import "fmt"

func main() {
	// if condition {
	// block of code
	// }

	age := 25

	if age >= 18 {
		fmt.Println("You are and adult")
	}

	// if condition{
	// block of code
	// } else if {
	// block of code
	// } else {
	// block of code
	// }

	temperature := 25
	if temperature >= 30 {
		fmt.Println("It's hot outside")
	} else {
		fmt.Println("It's cold outside")
	}

	score := 85

	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 80 {
		fmt.Println("Grade B")
	} else if score >= 70 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade D")
	}

	// this line will be executed after one of the condition is met

	// if condition1 {
	// Code block
	// if condition2{
	//code block 2
	// }
	// }

	//num := 15
	//if num%2 == 0 {
	//	if num%3 == 0 {
	//		fmt.Println("The number is divisible by 3 and 2")
	//	} else {
	//		fmt.Println("The number is divisible by 2 but not 3")
	//	} else {
	//		fmt.Println("The number is not divisible by 2 nor 3")
	//	}
	//}
}
