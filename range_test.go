package main

import (
	"testing"
	"fmt"
	"math/rand"
)

var list []float64

func fill() {
	if len(list) == 0 {
		for i := 0; i < 100000; i++ {
			list = append(list, rand.Float64())
		}
	}
}

func BenchmarkValue(b *testing.B) {
	fill()
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		for index := range list {
			fmt.Sprint(list[index])
		}
	}
}

func BenchmarkIndex(b *testing.B) {
	fill()
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		for _, value := range list {
			fmt.Sprint(value)
		}
	}
}
