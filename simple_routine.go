package main

import (
"fmt"
"time")

func count() {
	for i:=0; i<5;i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond*1)
	}
}

func main() {
	for _,val:=range make([]int, 1000000000) {
		fmt.Println(val)
		go count()
    }
	time.Sleep(time.Millisecond*2)
	fmt.Println("Hello World")
	time.Sleep(time.Millisecond*5)
}