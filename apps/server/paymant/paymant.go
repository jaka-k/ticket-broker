package paymant

import (
	"fmt"
	"math/rand"
	"time"
)

func ProcessOrder(id string) {
	random := rand.Intn(90) + 10
	duration := time.Duration(random) * time.Millisecond

	ticker := time.NewTicker(duration)

	fmt.Printf("The paymant [%s] was successfully process and confirmed. \n __rand__num__is=%v", id, duration)
	<-ticker.C
}
