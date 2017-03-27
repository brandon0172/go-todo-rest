package main

import (
	"fmt"
	"testing"
)

func BenchmarkXxx(testing.B)

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}

func Test() {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}
