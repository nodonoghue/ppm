package main

import (
	"flag"
	"fmt"
	"sync"

	"github.com/nodonoghue/ppm/internal/cli"
	"github.com/nodonoghue/ppm/internal/generate"
)

func main() {

	commandFlags := cli.GetFlags()
	flag.Parse()

	fmt.Println("Generating 10 AllChars Password Examples:")
	fmt.Println("-----------------------------------------")

	ch := make(chan string, *commandFlags.Options)

	//generate 10 options
	var wg sync.WaitGroup
	for i := 0; i < *commandFlags.Options; i++ {
		wg.Add(1)
		go generate.Password(ch, &wg)
	}
	wg.Wait()
	close(ch)

	optionNum := 1
	for password := range ch {
		fmt.Printf("Option %d: %s\n", optionNum, password)
		optionNum++
	}
}
