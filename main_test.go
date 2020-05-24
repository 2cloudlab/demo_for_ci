package main_test

import(
	"testing"
	"demo_for_ci/myapp"
)

func TestMultiply(t *testing.T) {
	main.Multiply(2 + 3)
}