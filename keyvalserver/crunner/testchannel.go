package main

import (
	"fmt"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}
func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}
func printer(c chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
		//time.Sleep(time.Millisecond * 5)
	}
}
func main() {
	aMap := make(map[string][]byte)
	aMap["imre"] = []byte("testing")
	fmt.Println(aMap["imre"])
}