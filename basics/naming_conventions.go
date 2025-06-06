package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

type EmployeApple struct {
	FirstName string
	LastName string
	Age int
}

func main() {
	// PascalCase
	// e.g. CalculateArea
	// Structs, interfaces, enum

	// snake_case
	// e.g. user_id
	// variables and filenames

	// UPPERCASE
	// e.g. ALPHABET
	// constants

	// camelCase
	// e.g. javaScript
	// variables and filenames

	const MAXRETRIES = 5

	var employeeID = 1001
	fmt.Println("EmployeeID: ", employeeID)

}