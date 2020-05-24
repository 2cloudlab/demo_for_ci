package mylib_test

import (
	"fmt"
	"testing"
	"demo_for_ci/mylib"
)

func TestAdd(t *testing.T) {
	result := mylib.Add(1, 1)
	fmt.Println(result)
}