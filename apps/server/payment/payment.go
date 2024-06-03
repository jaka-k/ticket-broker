package payment

import (
	"log"
	"math/rand"
	"time"
)

func ProcessOrder(id string) error {
	random := rand.Intn(90) + 10
	duration := time.Duration(random) * time.Millisecond

	log.Printf("Processing payment [%s] with delay: %v", id, duration)
	time.Sleep(duration)

	log.Printf("The payment [%s] was successfully processed and confirmed.", id)
	return nil
}
