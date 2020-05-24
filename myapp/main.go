package main

import(
	"fmt"
	"github.com/2cloudlab/demo_for_ci/mylib"
)

func main(){
	Aggregate(2, 3)
}

func Aggregate(a int, b int){
	Multiply(mylib.Add(a, b))
}

func Multiply(a int) {
	fmt.Println(a)
}