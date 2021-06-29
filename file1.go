package main

import (
	"fmt"
	"os"
)

func main(){
	fd, err := os.OpenFile("a.txt",os.O_RDWR|os.O_CREATE|os.O_TRUNC,0666)
	defer fd.Close()//defer will close the file after main function done
	if err != nil {
		fmt.Println("Failed to Open file!",err)
		return
	}
	//write string
	n, err := fd.WriteString("hello world\n")
	if err != nil {
		fmt.Println("Failed to WriteString",err)
		return
	}
	fmt.Printf("write success %d bytes\n",n)
	//change file read position
	fd.Seek(0,os.SEEK_SET)
    buf := make([]byte,20)
    //read file content
    n, err = fd.Read(buf)
    os.Stdout.Write(buf)
    //fd.Close()
}
