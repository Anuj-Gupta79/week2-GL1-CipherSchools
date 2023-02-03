package main

import (
	"fmt"
	"time"
)

func mainJs() {
	channel := make(chan string, 1)
	go func(ch chan<- string) {
		ch <- "2"
		fmt.Println(1)
	}(channel)
	message := <-channel
	time.Sleep(time.Second * 5)
	fmt.Println(message)
}

func mai2n() {
	channel := make(chan string, 1)
	go func(ch chan string) {
		mess := <-ch
		fmt.Println(mess)
		fmt.Println(1)

	}(channel)
	message := "Hello from MaAIN FUNCTION"
	channel <- message
	time.Sleep(time.Second * 5)
	message = <-channel
	fmt.Println(message)
}

// func main() {
// 	channel := make(chan string, 1)
// 	go func(ch chan<- string) {
// 		ch <- "2"
// 		// channel <- "3"
// 		//time.Sleep(time.Second * 5)
// 		fmt.Println(1)

// 	}(channel)
// 	message := <-channel
// 	//mess := <-channel
// 	time.Sleep(time.Second * 5)
// 	fmt.Println(message)
// 	// fmt.Println(mess)
// }

// func main() {
// 	channel := make(chan string, 1)
// 	go func() {
// 		channel <- "2"
// 		// channel <- "3"
// 		//time.Sleep(time.Second * 5)
// 		fmt.Println(1)

// 	}()
// 	message := <-channel
// 	//mess := <-channel
// 	time.Sleep(time.Second * 5)
// 	fmt.Println(message)
// 	// fmt.Println(mess)
// }

// func main() {
// 	channel := make(chan string)
// 	//var wait sync.WaitGroup
// 	//wait.Add(1)
// 	go func() {
// 		channel <- "Hello from Anonymous"
// 		//time.Sleep(time.Second * 5)
// 		fmt.Println(1)
// 		//wait.Done()
// 	}()
// 	go func() {
// 		message := <-channel
// 		// // time.Sleep(time.Second * 5)
// 		fmt.Println(message)
// 		//wait.Done()
// 	}()
// 	time.Sleep(time.Second * 5)
// 	message := <-channel
// 	// // time.Sleep(time.Second * 5)
// 	fmt.Println(message)
// 	//wait.Wait()
// }