package mthasher

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sync"
)

var Start byte = 97 // utf8 encoded of "a"

type iHashFunc func([]byte) [32]byte

type MultithreadHasher struct {
	spans    [][2]int
	hashFunc iHashFunc
	base     int
	len      int
}

func NewMultithreadHasher(hashFunc iHashFunc, base int, len int) *MultithreadHasher {
	return &MultithreadHasher{
		hashFunc: hashFunc,
		base:     base,
		len:      len,
	}
}

func (h *MultithreadHasher) Add(span [2]int) {
	h.spans = append(h.spans, span)
}

func (h *MultithreadHasher) Run(cases *[]string) {
	var wg sync.WaitGroup
	for n, span := range h.spans {
		wg.Add(1)
		fmt.Println(n, " starting with span", span[0], span[1])
		// Starting goroutine
		go func(span [2]int, cases *[]string) {

			defer wg.Done()

			for w := span[0]; w <= span[1]; w++ {
				word := Word(w, h.base, h.len)

				//fmt.Println(string(word))

				a := h.hashFunc(word)

				// Validation
				if len(*cases) == 0 {
					return
				}

				for i, c := range *cases {
					if c == hex.EncodeToString(a[:]) {
						fmt.Println(c, "answer: ", string(word))
						*cases = append((*cases)[:i], (*cases)[i+1:]...)
					}
				}
			}
		}(span, cases)
	}

	wg.Wait()
}

func SHA256(data []byte) [32]byte {
	return sha256.Sum256(data)
}

func Word(data int, base int, len int) []byte {
	a := make([]byte, 0)
	for i := 0; i < len; i++ {
		a = append([]byte{Start + byte(data%base)}, a...)
		data = data / base
	}
	return a
}
