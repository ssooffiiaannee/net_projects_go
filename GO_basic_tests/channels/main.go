package main

import (
	"fmt"
	// "time"
)

func channel_testing(c chan<- int){
	c <- 33

	close(c)
}

func main(){
	c := make(chan int, 33) // 33 capacity
	// _, ok := <-c
	go channel_testing(c)
	// time.Sleep(time.Second)
	_, ok := <-c
	fmt.Println(ok) // returns true because there is a remaining unread value
	_, ok = <-c
	fmt.Println(ok) // closed : no more unread data
	// for val := range c{
	// 	fmt.Println(val)
	// }
	// close(c)
}