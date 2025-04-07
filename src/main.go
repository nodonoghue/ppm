package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/nodonoghue/ppm/internal/cli"
	"github.com/nodonoghue/ppm/internal/generate"
	"github.com/nodonoghue/ppm/internal/models"
	"github.com/nodonoghue/ppm/internal/save"
)

func main() {

	//TODO: Add a tui to the generation, selection, saving, and lookup processes.
	
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
	models.Variants = make(map[int]string)
	for password := range ch {
		fmt.Printf("Option %d: %s\n", optionNum, password)
		models.Variants[optionNum-1] = password
		optionNum++
	}

	fmt.Println("Select an Option to copy to save to your bucket")
	pwIndex := cli.ReadInput()
	password := cli.GetVariant(pwIndex)

	fmt.Println("What is the username?")
	username := cli.ReadInput()

	fmt.Println("What is the URI/URL associated with this login?")
	uri := cli.ReadInput()

	fmt.Println("Give this login a name:")
	name := cli.ReadInput()

	var drop models.BucketDrop

	drop.Password = password
	drop.Name = name
	drop.URI = uri
	drop.Username = username

	fmt.Printf("You have selected: %s\n", password)

	u, err := json.Marshal(drop)
	if err != nil {
		log.Fatal("Unable to marshal struct to json: ", err.Error())
	}

	if err := save.Value(string(u)); err != nil {
		log.Fatal("Unable to save struct to file: ", err.Error())
	}
	fmt.Println("Saved to your bucket")
}
