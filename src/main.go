package main

import (
	"fmt"
	"reflect"
)

func main() {
	list1 := []int{1, 2, 3, 4}
	list2 := []int{1, 2, 3, 4}
	list3 := []int{1, 2, 3, 5}

	fmt.Println("list1 equals list2:", reflect.DeepEqual(list1, list2)) // true
	fmt.Println("list1 equals list3:", reflect.DeepEqual(list1, list3)) // false
}
