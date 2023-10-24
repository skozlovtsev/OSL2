package main

import (
	"fmt"
	"math"
	"time"

	"github.com/skozlovtsev/OSL2/pkg/config"
	"github.com/skozlovtsev/OSL2/pkg/mthasher"
)

func main() {
	mthasher.Start = config.Config.Start

	MTHasher := mthasher.NewMultithreadHasher(mthasher.SHA256, config.Config.Base, config.Config.Len)

	allVariances := int(math.Pow(float64(config.Config.Base), float64(config.Config.Len)))

	spanSize := allVariances / config.Config.Threads

	lastSpanSize := spanSize + allVariances%config.Config.Threads

	start := 0

	for i := 0; i < config.Config.Threads-1; i++ {
		end := start + spanSize

		MTHasher.Add([2]int{start, end - 1})

		start = end
	}

	end := start + lastSpanSize

	MTHasher.Add([2]int{start, end})

	// Start main loop
	stime := time.Now().UnixMilli()

	MTHasher.Run(&config.Config.Cases)

	timeUnixMilli := time.Now().UnixMilli() - stime

	fmt.Printf("time: %d.%d\n", timeUnixMilli/1000, timeUnixMilli%1000)
}
