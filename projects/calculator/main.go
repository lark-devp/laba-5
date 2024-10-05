package main

import "fmt"

func calculator(firstChan, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	output := make(chan int)
	go func(ch chan int) {
		defer close(ch)
		select {
		case x := <-firstChan:
			ch <- x * x
			fmt.Println("Первый канал")
		case x := <-secondChan:
			ch <- x * 3
			fmt.Println("Второй канал")
		case <-stopChan:
			fmt.Println("Стоп")

		}

	}(output)
	return output
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	stop := make(chan struct{})
	result := calculator(ch1, ch2, stop)
	ch1 <- 5

	fmt.Println(<-result)
}
