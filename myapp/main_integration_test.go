// +build integration

package main_test

import(
	"testing"
	"github.com/2cloudlab/demo_for_ci/myapp"
)

func TestAggregate(t *testing.T) {
	main.Aggregate(2, 3)
}