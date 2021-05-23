/*
描述

无穷数列1，1，2，3，5，8，13，21，34，55...称为Fibonacci数列，它可以递归地定义为 F(n)=1 ...........(n=1或n=2) F(n)=F(n-1)+F(n-2).....(n>2)
现要你来求第n个斐波纳奇数。（第1个、第二个都为1)

输入

第一行是一个整数m(m<5)表示共有m组测试数据 每次测试数据只有一行，且只有一个整形数n(n<20)

输出

对每组输入n，输出第n个Fibonacci数

样例输入

3 1 3 5

样例输出

1 2 5

 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Fibonacci(n int) int{
	var result int
	if(n > 2){
		result = Fibonacci(n-1) + Fibonacci(n-2)
		return result
	}else{
		return 1
	}
}
func main(){
	var str string
	var arr []string
	var res int
	inputReader := bufio.NewReader(os.Stdin)
	str,_ = inputReader.ReadString('\n')
	arr = strings.Fields(str)
	fmt.Println(arr)
	for i,_ := range arr {
		if(i > 0){
			res,_ = strconv.Atoi(arr[i])
			fmt.Printf("%d ",Fibonacci(res))
		}
	}

}

