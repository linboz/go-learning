/**
输入

输入两个数，a,b

输出

输出a+b的值

样例输入

2 3

样例输出

5
 */
package main

import "fmt"

func main(){
	var a, b int
	fmt.Scanln(&a,&b)
	fmt.Println(a+b)
}