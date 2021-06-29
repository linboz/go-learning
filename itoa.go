package main

import "fmt"

func main(){
	const (
		login = iota
		logout
		user = iota + 1
		account = iota +3
	)
	fmt.Println(login,logout,user,account)
}
