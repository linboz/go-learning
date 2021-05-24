package main

import (
   "fmt"
   "strings"
)
//实现 StrCount。它应当返回一个映射（图），其中包含字符串 s 中每个“字符”的个数，忽略空格。
func main(){
  str := "benson is a good man for a man"
  fmt.Println(StrCount(str))

}
func StrCount(s string) map[string]int {
   var arr []string
   s = strings.Replace(s," ","",-1)
   ss := []rune(s) //字符串转换成int32整数数组
   for i:=0;i<len(ss);i++{
      arr = append(arr,string(ss[i]))
   }
   //fmt.Println(arr)
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
