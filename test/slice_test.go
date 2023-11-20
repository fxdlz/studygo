package test

import (
	"fmt"
	"studygo/generics"
	"testing"
)

func TestSliceDelete(t *testing.T) {
	var sliceInt = []int{1, 2, 3, 4, 5}
	sliceInt = generics.Delete(1, sliceInt)
	fmt.Printf("val:%v,len:%v,cap:%v\n", sliceInt, len(sliceInt), cap(sliceInt))
	sliceInt = generics.Delete(1, sliceInt)
	fmt.Printf("val:%v,len:%v,cap:%v\n", sliceInt, len(sliceInt), cap(sliceInt))
	sliceInt = generics.Delete(1, sliceInt)
	fmt.Printf("val:%v,len:%v,cap:%v\n", sliceInt, len(sliceInt), cap(sliceInt))

	var sliceStr = []string{"1", "2", "3", "4", "5"}
	sliceStr = generics.Delete(1, sliceStr)
	fmt.Printf("val:%v,len:%v,cap:%v\n", sliceStr, len(sliceStr), cap(sliceStr))
	sliceStr = generics.Delete(1, sliceStr)
	fmt.Printf("val:%v,len:%v,cap:%v\n", sliceStr, len(sliceStr), cap(sliceStr))
	sliceStr = generics.Delete(1, sliceStr)
	fmt.Printf("val:%v,len:%v,cap:%v\n", sliceStr, len(sliceStr), cap(sliceStr))

}
