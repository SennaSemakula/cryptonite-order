package main

import (
	"context"
	"log"
	"time"
)

var log log.Logger

func main() {
	log.SetPrefix("crypto-order")
	ctx := context.Background()
	context.WithTimeout(ctx, time.Duration*2)

}
