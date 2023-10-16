package main

import (
	"context"
	"fmt"
	"time"

	"github.com/skozlovtsev/OSL2/pkg/mthasher"
	"github.com/skozlovtsev/OSL2/pkg/validator"
)

var (
	base    = 26
	len     = 5
	threads = 1
)

func main() {
	fmt.Scanf("%i", &threads)

	answerChan := make(chan [32]byte, threads*2)

	ctx, close := context.WithCancel(context.Background())

	MTHasher := mthasher.NewMultithreadHasher(close, mthasher.SHA256, answerChan, ctx, base, len)

	Validator := validator.NewValidator(answerChan)

	// Start main loop
	stime := time.Now().Unix()

	MTHasher.Run()

	Validator.Start()

	fmt.Println(time.Now().Unix() - stime)
}
