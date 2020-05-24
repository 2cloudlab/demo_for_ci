package main_test

import(
	"testing"
	"2cloudlab/demo_for_ci/myapp"
)

func TestMultiply(t *testing.T) {
	main.Multiply(2 + 3)
}