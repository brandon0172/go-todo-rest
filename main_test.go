package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Welcome %d times\n", i)
	}
}
