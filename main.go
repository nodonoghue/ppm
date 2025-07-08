package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"syscall"
	"time"

	"github.com/nodonoghue/ppm/internal/cli"
	"github.com/nodonoghue/ppm/internal/encryption"
	"github.com/nodonoghue/ppm/internal/generate"
	"github.com/nodonoghue/ppm/internal/models"
	"github.com/nodonoghue/ppm/internal/save"
	"golang.org/x/term"
)

func main() {
	commandFlags := cli.GetFlags()

	var drops []models.BucketDrop

	data, err := save.ReadFile()
	if err != nil {
		log.Fatal("Unable to read vault: ", err.Error())
	}

	password := getPassword(data != nil)

	if data != nil {
		decryptedData, err := encryption.Decrypt(data, password)
		if err != nil {
			log.Fatal("Unable to decrypt vault: ", err.Error())
		}
		if err := json.Unmarshal(decryptedData, &drops); err != nil {
			log.Fatal("Unable to unmarshal vault: ", err.Error())
		}
	}

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
	selectedPassword, err := cli.GetVariant(pwIndex, variants)
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

	drop.Password = selectedPassword
	drop.Name = name
	drop.URI = uri
	drop.Username = username

	drops = append(drops, drop)

	marshalledDrops, err := json.Marshal(drops)
	if err != nil {
		log.Fatal("Unable to marshal struct to json: ", err.Error())
	}

	encryptedDrops, err := encryption.Encrypt(marshalledDrops, password)
	if err != nil {
		log.Fatal("Unable to encrypt vault: ", err.Error())
	}

	if err := save.OverwriteFile(encryptedDrops); err != nil {
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

func getPassword(existingVault bool) string {
	if existingVault {
		fmt.Println("Enter password to decrypt vault:")
	} else {
		fmt.Println("Enter password to encrypt vault:")
	}

	if term.IsTerminal(int(syscall.Stdin)) {
		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatal("Unable to read password: ", err.Error())
		}
		return string(bytePassword)
	} else {
		fmt.Println("Warning: running in a non-interactive terminal, password will be echoed.")
		reader := bufio.NewReader(os.Stdin)
		password, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Unable to read password: ", err.Error())
		}
		return password
	}
}