package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/nodonoghue/ppm/internal/cli"
	"github.com/nodonoghue/ppm/internal/generate"
	"github.com/nodonoghue/ppm/internal/save"
)

func main() {

	commandFlags := cli.GetFlags()
	flag.Parse()

	fmt.Printf("Generating %d AllChars Password Examples:\n", *commandFlags.NumVariants)
	fmt.Println("-----------------------------------------")

	ch := make(chan string, *commandFlags.NumVariants)

	var wg sync.WaitGroup
	for i := 0; i < *commandFlags.NumVariants; i++ {
		wg.Add(1)
		go generate.Password(ch, &wg, commandFlags)
	}
	wg.Wait()
	close(ch)

	optionNum := 1
	variants := make(map[int]string)
	for password := range ch {
		fmt.Printf("Option %d: %s\n", optionNum, password)
		variants[optionNum-1] = password
		optionNum++
	}

	fmt.Println("Select an Option to copy to the clipboard")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Unable to read input")
	}
	//for windows use "/r/n" instead of "/n"  need to add code to check current environment and determine the correct new line char(s)
	num, err := strconv.Atoi(strings.Replace(input, "\n", "", -1))
	if err != nil {
		log.Fatal("Input must be numeric. ", err.Error())
	}

	selected := variants[num-1]

	fmt.Printf("You have selected: %s\n", selected)

	save.SaveValue(selected)
	fmt.Println("Saved to your bucket")
}
