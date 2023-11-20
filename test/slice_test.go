package test

import (
	"fmt"
	"studygo/generics"
	"testing"
)

func TestSliceDelete(t *testing.T) {
	var slice = []int{1, 2, 3, 4, 5}
	slice = generics.Delete(1, slice)
	fmt.Printf("val:%v,len:%v,cap:%v\n", slice, len(slice), cap(slice))
	slice = generics.Delete(1, slice)
	fmt.Printf("val:%v,len:%v,cap:%v\n", slice, len(slice), cap(slice))
	slice = generics.Delete(1, slice)
	fmt.Printf("val:%v,len:%v,cap:%v\n", slice, len(slice), cap(slice))
}
