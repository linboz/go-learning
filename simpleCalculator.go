package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	var a,b,result float64
	var operator,str string
	var arr []string
	fmt.Println("Please input: ")
	inputReader := bufio.NewReader(os.Stdin)
	str,_ = inputReader.ReadString('\n')
	arr = strings.Fields(str)
	a,_ = strconv.ParseFloat(arr[0],32)
	operator = arr[1]
	b,_ = strconv.ParseFloat(arr[2],32)
	switch operator {
		case "+" :
			result = a + b
		case "-" :
			result = a - b
		case "*" :
			result = a * b
		case "/" :
			result = a / b
		default:
			result = 0
	}
	fmt.Printf("The reuslt is: %.2f",result)
}
