package main

import "fmt"

func zerovalue(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial", i)

	zerovalue(i)
	fmt.Println("zeroval", i)
	zeroptr(&i)
	fmt.Println("zeroptr", i)
	fmt.Println("pointer", &i)

	var number int = 58
	var ptr *int = &number // Stores the address of 'number' in 'ptr'

	fmt.Println("Address of 'number':", &number)      // Outputs the memory address of 'number'
	fmt.Println("Address held by 'ptr':", ptr)        // Outputs the address stored in 'ptr'
	fmt.Println("Value of 'number' via 'ptr':", *ptr) // Accesses the value of 'number' through 'ptr'

	var a int = 20  // A normal integer variable
	var p *int = &a // 'p' holds the address of 'a'

	fmt.Println("Value of 'a':", a)    // Outputs the value of 'a' (20)
	fmt.Println("Address of 'a':", &a) // Outputs the memory address of 'a'

	*p = 10                             // Assigns 10 to the address held by 'p', thereby changing the value of 'a'
	fmt.Println("New value of 'a':", a) // Outputs the new value of 'a' (10)
}
