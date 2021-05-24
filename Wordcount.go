package main

import (
   "fmt"
   "strings"
)
//实现 WordCount。它应当返回一个映射（图），其中包含字符串 s 中每个“单词”的个数。
func main(){
  str := "benson is a good man for a man"
  fmt.Println(WordCount(str))

}
func WordCount(s string) map[string]int {
   arr := strings.Fields(s)
   var m map[string]int
   m = make(map[string]int)
   for _,v := range arr{
      _,ok := m[v]
      if ok {
         m[v] = m[v] + 1
      }else{
         m[v] = 1
      }
   }
   return m
}
