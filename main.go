package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/nodonoghue/ppm/internal/cli"
	"github.com/nodonoghue/ppm/internal/generate"
	"github.com/nodonoghue/ppm/internal/models"
	"github.com/nodonoghue/ppm/internal/save"
)

func main() {
	commandFlags := cli.GetFlags()

	fmt.Printf("Generating %d AllChars Password Examples:\n", *commandFlags.NumVariants)
	fmt.Println("-----------------------------------------")

	passwords := generatePasswords(commandFlags)

	variants := make(map[int]string)
	for i, password := range passwords {
		fmt.Printf("Option %d: %s\n", i+1, password)
		variants[i] = password
	}

	fmt.Println("Select an Option to copy to save to your bucket")
	pwIndex := cli.ReadInput()
	password, err := cli.GetVariant(pwIndex, variants)
	if err != nil {
		log.Fatal("Entry must be numeric")
	}

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

func generatePasswords(commandFlags models.CommandFlags) []string {
	ch := make(chan string, *commandFlags.NumVariants)
	var wg sync.WaitGroup
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for range *commandFlags.NumVariants {
		wg.Add(1)
		go generate.Password(ch, &wg, commandFlags, r)
	}

	wg.Wait()
	close(ch)

	passwords := make([]string, 0, *commandFlags.NumVariants)
	for password := range ch {
		passwords = append(passwords, password)
	}

	return passwords
}
