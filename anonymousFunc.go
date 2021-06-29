package main

import "fmt"

func main(){
	nextNum := getSequence()
	fmt.Println(nextNum())
	fmt.Println(nextNum())
	f := getSequence()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(nextNum())
}

func getSequence() func() int{
	i := 0
	return func() int{
		i += 1
		return i
	}
}
