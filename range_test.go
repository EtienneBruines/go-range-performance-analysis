package main

import (
	"testing"
	"fmt"
	"math/rand"
	"github.com/dustin/randbo"
	"io"
	"bytes"
)

type large_object struct {
	Binary []byte
	Pointer *large_object
	Stringy string
}

var list []float64
var list_large []large_object
var list_large_pointer []*large_object
// random_reader is a source where we get random binary data
var random_reader = randbo.New()
const SIZE = 1024

// dummy_object creates a large_object with different values each time. 
func dummy_object(i int) large_object {
	obj := large_object{}
	obj.Stringy = "abc" + fmt.Sprint(i)
	if i > 0 {
		obj.Pointer = &list_large[0]
	}
	
	var buf bytes.Buffer
	io.CopyN(&buf, random_reader, SIZE)
	copy(obj.Binary, buf.Bytes())

	return obj
}

// fill ensures all slices/arrays are filled
func fill() {
	if len(list) == 0 {
		for i := 0; i < 100000; i++ {
			list = append(list, rand.Float64())
		}
	}
	if len(list_large) == 0 {
		for i := 0; i < 100000; i++ {
			obj := dummy_object(i)
				
			list_large = append(list_large, obj)
			list_large_pointer = append(list_large_pointer, &obj)
		}
	}
}


func BenchmarkLargeValue(b *testing.B) {
	fill()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for index := range list_large {
			fmt.Sprint(list_large[index])
		}
	}
}

func BenchmarkLargeIndex(b *testing.B) {
	fill()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, value := range list_large {
			fmt.Sprint(value)
		}
	}
}

func BenchmarkLargeValueMultiple(b *testing.B) {
	fill()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for index := range list_large {
			fmt.Sprint(list_large[index])
			fmt.Sprint(list_large[index])
			fmt.Sprint(list_large[index])
			fmt.Sprint(list_large[index])
			fmt.Sprint(list_large[index])
		}
	}
}

func BenchmarkLargeIndexMultiple(b *testing.B) {
	fill()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, value := range list_large {
			fmt.Sprint(value)
			fmt.Sprint(value)
			fmt.Sprint(value)
			fmt.Sprint(value)
			fmt.Sprint(value)
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

func BenchmarkValueMultiple(b *testing.B) {
	fill()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for index := range list {
			fmt.Sprint(list[index])
			fmt.Sprint(list[index])
			fmt.Sprint(list[index])
			fmt.Sprint(list[index])
			fmt.Sprint(list[index])
		}
	}
}

func BenchmarkIndexMultiple(b *testing.B) {
	fill()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, value := range list {
			fmt.Sprint(value)
			fmt.Sprint(value)
			fmt.Sprint(value)
			fmt.Sprint(value)
			fmt.Sprint(value)
		}
	}
}
