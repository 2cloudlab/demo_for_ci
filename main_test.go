package main_test

import(
	"testing"
	"demo_for_ci"
)

func TestMultiply(t *testing.T) {
	main.Multiply(2 + 3)
}