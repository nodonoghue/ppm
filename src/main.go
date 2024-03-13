package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/nodonoghue/ppm/models"
)

func main() {

	fmt.Println("Generating 10 AllChars Password Examples:")
	fmt.Println("-----------------------------------------")

	ch := make(chan string, 10)

	//generate 10 options
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go createPassword(ch, &wg)
	}
	wg.Wait()
	close(ch)

	optionNum := 1
	for password := range ch {
		fmt.Printf("Option %d: %s\n", optionNum, password)
		optionNum++
	}
}

func createPassword(ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	rand.New(rand.NewSource(time.Now().UnixNano()))

	//use shuffle to create a len20 password and print to terminal:
	buf := make([]byte, models.Length)
	for i := 0; i < models.Length; i++ {
		buf[i] = models.AllChars[rand.Intn(len(models.AllChars))]
	}
	ch <- string(buf)
}
