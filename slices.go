package main

import "fmt"

func main(){
	var slice []int
	slice = append(slice, 1)
	slice = append(slice,2)
	slice = append(slice,3,4,5)
	printSlice(slice)
	slice = remove(slice,2)
	printSlice(slice)
	s2 := make([]int,3)
	printSlice(s2)
	s2 = append(s2,4)
	printSlice(s2)
}

func printSlice(s []int){
	fmt.Printf("len = %d, cap = %d, s = %v\n",len(s),cap(s),s)
}

func remove(s []int,pos int) []int{
	return append(s[:pos],s[pos+1:]...)
}
