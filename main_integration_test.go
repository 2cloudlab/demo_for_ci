// +build integration

package main_test

import(
	"testing"
	"demo_for_ci/myapp"
)

func TestAggregate(t *testing.T) {
	main.Aggregate(2, 3)
}