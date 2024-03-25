package main

import "fmt"

func updateSlice(mySlice []int) {
	mySlice[0] = 11
}

func addOne(num1 *int, num2 *int, num3 *int) {
	*num1 = *num1 + 1
	*num2 = *num2 + 1
	*num3 = *num3 + 1
}
func main() {

	var number, number2, number3 int
	fmt.Printf("Before (number):%v\n", number)
	fmt.Printf("Before (number2):%v\n", number2)
	fmt.Printf("Before (number3):%v\n", number3)
	fmt.Println("----------")
	addOne(&number, &number2, &number3)
	fmt.Printf("After (number):%v\n", number)
	fmt.Printf("After (number2):%v\n", number2)
	fmt.Printf("After (number3):%v\n", number3)

	/*
		fmt.Println("----------")
		fmt.Printf("Before:%v\n",intSlice)
		updateSlice(intSlice)
		fmt.Printf("After:%v\n",intSlice)
	*/

}
