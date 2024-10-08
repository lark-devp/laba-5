package main

// реализовать removeDuplicates(in, out chan string)
import "fmt"

func removeDuplicates(in chan string, out chan string) {
	var stp string
	defer close(out)
	for st := range in {
		if st != stp {
			out <- st
			stp = st
		}
	}
}

func main() {
	// здесь должен быть код для проверки правильности работы функции removeDuplicates(in, out chan string)
	inputStream := make(chan string)
	outputStream := make(chan string)
	go removeDuplicates(inputStream, outputStream)
	var str string
	fmt.Print("Text has been captured: ")
	fmt.Scanln(&str)
	go func() {
		defer close(inputStream)
		for _, st := range str {
			inputStream <- string(st)
		}
	}()
	fmt.Print("Result: ")
	for st := range outputStream {
		fmt.Print(st)
	}
	fmt.Println("\n", "End")
}
