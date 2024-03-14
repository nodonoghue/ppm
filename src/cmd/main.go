package main

import (
	"fmt"
	"sync"

	"github.com/nodonoghue/ppm/internal/utils"
)

func main() {

	fmt.Println("Generating 10 AllChars Password Examples:")
	fmt.Println("-----------------------------------------")

	ch := make(chan string, 10)

	//generate 10 options
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go utils.CreatePassword(ch, &wg)
	}
	wg.Wait()
	close(ch)

	optionNum := 1
	for password := range ch {
		fmt.Printf("Option %d: %s\n", optionNum, password)
		optionNum++
	}
}
