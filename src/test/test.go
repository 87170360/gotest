package main

import (
	"fmt"
	"time"
)

type Semaphore chan bool

func NewSemaphore(num int) *Semaphore {
	s := make(Semaphore, num)
	return &s
}

func (s Semaphore) P() {
	s <- true
}

func (s Semaphore) V() {
	<-s
}

func main() {
	s := NewSemaphore(0)
	for i := 0; i < 1; i++ {
		go func() {
			fmt.Println("hello")  //1
			s.P()                 //2
			fmt.Println("hello2") //3
		}()
	}

	for i := 0; i < 1; i++ {
		go func() {
			fmt.Println("hlx")  //4
			s.V()               //5
			fmt.Println("hlx2") //6
		}()
	}

	time.Sleep(time.Millisecond * 100000)
}
