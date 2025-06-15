package basics

import "fmt"

func main() {
	// Simple iteration over a range
	// for i:=1; i <= 5; i++ {
	//fmt.Println(i)
	//}

	// iterate over collection
	numbers := []int{1, 2, 3, 4, 5, 6}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	//for i:=1; 1<=10; i ++ {
	//	if i%2 == 0 (
	//		continue
	//	)
	//	fmt.Println("Odd number: ", i)
	//	if i == 5 {
	//		fmt.Println("")
	//	}
	//}

	//rows := 5
	//
	//// Outer loop
	//for i := 1; i <= rows; i++ {
	//	// inner loop for spaces before stars
	//	for j := 1; j <= rows-1; j++ {
	//		fmt.Println(" ")
	//	}
	//	// inner loop for stars
	//	for k := 1; k <= 2*i-1; k++ {
	//		fmt.Println("*")
	//	}
	//	fmt.Println()
	//}

	for i := range 10 {
		i++
		fmt.Println(i)
	}
	fmt.Println("We have lift off!")
}
