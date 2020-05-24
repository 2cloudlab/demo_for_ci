package main

import(
	"fmt"
	"demo_for_ci/mylib"
)

// main
func main(){
	Aggregate(2, 3)
}

// Aggreage
func Aggregate(a int, b int){
	Multiply(mylib.Add(a, b))
}

// Multiply
func Multiply(a int) {
	fmt.Println(a)
}