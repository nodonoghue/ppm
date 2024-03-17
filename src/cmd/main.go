package main

import (
	"flag"
	"fmt"
	"sync"

	"github.com/nodonoghue/ppm/internal/generate"
)

func main() {

	opts := commandFlags()

	fmt.Println("Generating 10 AllChars Password Examples:")
	fmt.Println("-----------------------------------------")

	ch := make(chan string, *opts)

	//generate 10 options
	var wg sync.WaitGroup
	for i := 0; i < *opts; i++ {
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

func commandFlags() *int {
	//Note to future self, maybe...
	//each new flag should be registered in it's own function
	//desire to have short and long for the same
	//The trick will be to pass back the handful of CLI options back to the callers
	//or simply execute the action, will depend on use-case and context.
	opts := flag.Int("options", 1, "Sets the number of password options to generate")
	flag.Parse()
	return opts
}
