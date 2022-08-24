package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// 11
var wg sync.WaitGroup

// test

const kmaxPrimeChanNum = 1000000
const kNum = 1000000

func IsPrimeSqrt(value int) (int, bool) {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return 0, false
		}
	}
	return value, value > 1
}

// write data to inchan
func putNum(intChan chan<- int) {

	for i := 2; i < kNum; i++ {
		intChan <- i
	}
	close(intChan)
	wg.Done()
}

func primeNum(intChan <-chan int, primeChan chan<- int, exitChan chan<- bool) {
	for i := range intChan {
		if v, ok := IsPrimeSqrt(i); ok {
			primeChan <- v
		} else {
			break
		}
	}

	wg.Done()

	exitChan <- true
}

func printPrime(primeChan <-chan int) {
	for v := range primeChan {
		fmt.Printf("Prime number: %v\n", v)
	}
	wg.Done()
}

func main() {
	mySelect()
}

func myPrime() {
	intChan := make(chan int, 10000)
	primeChan := make(chan int, 50000)
	exitChan := make(chan bool, kmaxPrimeChanNum) // exit symbol for primeChan

	start := time.Now().Unix()
	wg.Add(1)
	go putNum(intChan) // store the numbers
	for i := 0; i < kmaxPrimeChanNum; i++ {
		wg.Add(1)
		go primeNum(intChan, primeChan, exitChan) // get the prime numbers
	}
	wg.Add(1)
	go printPrime(primeChan) // print the numbers

	// to determine if chan is full
	wg.Add(1)
	go func() {
		for i := 0; i < kmaxPrimeChanNum; i++ {
			<-exitChan
		}
		close(primeChan)
		wg.Done()
	}()
	wg.Wait()
	end := time.Now().Unix()
	fmt.Printf("Finished with %v seconds\n", (end - start))
}

func mySelect() {
	// to recover
	defer func() { //  to catch panic
		if err := recover(); err != nil {
			fmt.Println("mySelect panic ", err)
		}
	}()

	// define a chan with 10 int
	iChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		iChan <- i
	}

	// define a chan with 5 strings
	sChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		sChan <- "hello" + fmt.Sprintf("%d", i)
	}

	// select does not require chan to be closed
	for {
		select {
		case v := <-iChan:
			fmt.Printf("get value from iChan: %v\n", v)
		case v := <-sChan:
			fmt.Printf("get value from sChan %v\n", v)
		default:
			fmt.Printf("Finished")
			return
		}

	}

}
