package main

import (
	"fmt"
	"m1/vector"
)

func main() {
	v1 := vector.Vec2{X: 1, Y: 2}
	v2 := vector.Vec2{X: 2, Y: 4}
	var v3 vector.Vec2

	fmt.Printf("v1=%v\n", v1)
	fmt.Printf("v2=%v\n", v2)
	fmt.Printf("v3=%v\n", v3)

	fmt.Printf("%v\n", vector.AddVec2(v1, v2))

	fmt.Printf("Before calling Negate v1=%v\n", v1)
	v1.Negate()
	fmt.Printf("After calling Negate v1=%v\n", v1)

}
