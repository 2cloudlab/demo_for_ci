package main

import(
	"fmt"
	"demo_for_ci/mylib"
)

// main
func main(){
	Aggregate(2, 3)
}

// Aggreage a + b
func Aggregate(a int, b int){
	Multiply(mylib.Add(a, b))
}

// Multiply multiply by 3
func Multiply(a int) {
	fmt.Println(3 * a)
}