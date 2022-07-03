package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, 世界")
  
	timerChan := make(chan time.Time)
	
  go func() {
		time.Sleep(5 * time.Second)
		timerChan <- time.Now()  // send value over the channel
	}()
	
  timer := <-timerChan  // recieve value to channel
	fmt.Println("heloo  ", timer)
}
