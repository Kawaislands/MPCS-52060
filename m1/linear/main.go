package main

import (
	"fmt"
	"m1/vector"
)

func main() {

	// v1 := vector.Vec2{1, 3}
	v1:= vector.MakeVec2(1, 3)
	v2 := vector.MakeVec2(2, 4)
	v3 := vector.AddVec2(v1, v2)

	v3.Negate()
	
	// fmt.Printf("v1 = %v, %v", v1.X, v1.Y)
	// fmt.Printf("v2 = %v, %v", v2.X, v2.Y)
	fmt.Printf("v3 = %v, %v", v3.X, v3.Y)
}
