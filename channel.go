package main 

import ("fmt"
"time")

func printCount(c chan int) {
	num:=0
	for num>=0 {
		num := <-c
		fmt.Println(num, " ")
	}
	
}

func main() {
	c:=make(chan int)
	a:=[]int{8,6,7,5,3,0,9,-1}

	go printCount(c)
	go printCount(c)
	go printCount(c)
	go printCount(c)
	go printCount(c)
	go printCount(c)
	go printCount(c)
	go printCount(c)
	for _,val := range a {
		c <- val
    }
	time.Sleep(time.Millisecond*1)
	fmt.Println("end of main")
}