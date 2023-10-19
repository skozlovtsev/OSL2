package main

import (
	"fmt"
	"math"
	"time"

	"github.com/skozlovtsev/OSL2/pkg/mthasher"
)

var (
	base    = 26
	len     = 5
	threads = 1

	cases = []string{
		"1115dd800feaacefdf481f1f9070374a2a81e27880f187396db67958b207cbad",
		"3a7bd3e2360a3d29eea436fcfb7e44c735d117c42d1c1835420b6b9942dd4f1b",
		"74e1bb62f8dabb8125a58852b63bdf6eaef667cb56ac7f7cdba6d7305c50a22f",
	}
)

func main() {
	fmt.Scanf("%d", &threads)

	MTHasher := mthasher.NewMultithreadHasher(mthasher.SHA256, base, len)

	allVariances := int(math.Pow(float64(base), float64(len)))

	spanSize := allVariances / threads

	lastSpanSize := spanSize + allVariances%threads

	start := 0

	for i := 0; i < threads-1; i++ {
		end := start + spanSize

		MTHasher.Add([2]int{start, end})

		start = end
	}

	end := start + lastSpanSize

	MTHasher.Add([2]int{start, end})

	// Start main loop
	stime := time.Now().Unix()

	MTHasher.Run(cases)

	fmt.Println("time: ", time.Now().Unix()-stime)
}
