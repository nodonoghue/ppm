package utils

import (
	"math/rand"
	"sync"
	"time"

	"github.com/nodonoghue/ppm/internal/models"
)

func CreatePassword(ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	rand.New(rand.NewSource(time.Now().UnixNano()))

	//use shuffle to create a len20 password and print to terminal:
	buf := make([]byte, models.Length)
	for i := 0; i < models.Length; i++ {
		buf[i] = models.AllChars[rand.Intn(len(models.AllChars))]
	}
	ch <- string(buf)
}
