package basics

import "fmt"

func main() {

	/// for as while with break
	// sum := 0
	//for {
	//	sum += 10
	//	fmt.Printf("sum: %d\n", sum)
	//	if sum >= 50 {
	//		break
	//	}
	//}

	num := 1
	for num <= 10 {
		if num%2 == 0 {
			num++
			continue
		}
		fmt.Println("Odd number: ", num)
		num++ //++ increment operator increases value by 1 -- decrement operator, decreases value by 1
	}
}
