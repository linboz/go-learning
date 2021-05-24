package main

import (
   "fmt"
)
//实现一个 fibonacci 函数，它返回一个函数（闭包），该闭包返回一个斐波纳契数列 `(0, 1, 1, 2, 3, 5, ...)`
// 返回一个“返回int的函数”
func fibonacci() func() int {
   var first,next = 0, 1
   return func() int{
      tmp := first
      first,next = next,(first + next)
      return tmp
   }
}

func main() {
   f := fibonacci()
   for i := 0; i < 10; i++ {
      fmt.Println(f())
   }
}
