package test

import (
	"MVC_DI/util/stream"
	"fmt"
)

func TestListStream() {
	source := []int{1, 2, 3}
	fmt.Printf("source: %v\n", source)
	doubled := stream.NewListStream(source).Map(func(item int) any { return item * 2 }).ToList()
	fmt.Printf("double: %v\n", doubled)
	even := stream.NewListStream(source).Filter(func(item int) bool { return item%2 == 0 }).ToList()
	fmt.Printf("even: %v\n", even)
}

func TestMapStream() {
	source := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Printf("source: %v\n", source)
	doubled := stream.NewMapStream(source).Map(func(key string, val int) (string, any) { return key + "_double", val * 2 }).ToMap()
	fmt.Printf("double: %v\n", doubled)
	even := stream.NewMapStream(source).Filter(func(key string, val int) bool { return val%2 == 0 }).ToMap()
	fmt.Printf("even: %v\n", even)
}
