package main

import (
	"fmt"
	"time"
	"sync"
//	"runtime"
)

var wg sync.WaitGroup

func worker (in,out chan int) {
	defer wg.Done()
	for x:= range in {
		//fmt.Println("=",x)
		out<-x
	}
	//close(out)
}

func reader (in chan int) {
	defer wg.Done()
	for x:=range in {
		fmt.Println("reader>", x)
	}
	// time.Sleep(1000 * time.Millisecond)
	// time.Sleep(100 * time.Millisecond)
	//  for {
	//  	select {
	//  		case x := <-in:
	// 			fmt.Println("reader>", x)
	// 			time.Sleep(100 * time.Millisecond)
	// 		default:
	//  			fmt.Println("*********")
	//  			return
	//  }
//}
}

func main() {
	ci := make(chan int)
	co := make(chan int, 100)
	wg.Add(3)
	go worker(ci,co)
	go worker(ci,co)
	go worker(ci,co)
	
	for i:=1; i<20; i++ {
		ci<-i
		//fmt.Println(">>", i)
	}
	//fmt.Println("=====")
	close(ci)

	wg.Wait()
	go reader(co)
	//fmt.Println(<-co)
	//fmt.Println(<-co)
	
//	time.Sleep(3 * time.Second)
	// for {
	// 	if runtime.NumGoroutine()==1 {
	// 		return
	// 	}
	// //fmt.Println("---",runtime.NumGoroutine())
	//fmt.Println("ppppppp")
	time.Sleep(1000000000)
	//fmt.Println("oooooo")
	//time.After(100 * time.Second)
	
	//}
	
}

