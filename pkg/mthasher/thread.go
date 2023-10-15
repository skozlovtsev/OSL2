package mthasher

import (
	"context"
	"crypto/sha256"
)

const start = 97

type iHashFunc func([]byte) [32]byte

type Thread struct {
	answerChan chan [32]byte   // buffered answer channel
	hashFunc   iHashFunc       // hash function
	ctx        context.Context // goroutine context for goroutine termination
	alphabet   [2]byte         // span of alphabet
	span       [][2]byte       // span for thread
	id         uint8           // thread identifier
}

func NewThread(
	id uint8, // thread identifier
	answerChan chan [32]byte, // buffered answer channel
	hashFunc iHashFunc, // hash function
	ctx context.Context, // goroutine context for goroutine termination
	alphabet [2]byte, // span of alphabet
	span [][2]byte, // span for thread
) *Thread {
	return &Thread{
		id:         id,
		answerChan: answerChan,
		hashFunc:   hashFunc,
		ctx:        ctx,
		alphabet:   alphabet,
		span:       span,
	}
}

func (t *Thread) Start() {

}

func SHA256(data []byte) [32]byte {
	return sha256.Sum256(data)
}

func word(data int, base int, len int) []byte {
	a := make([]byte, 5)
	for i := 0; i < len; i++ {
		a = append([]byte{start + byte(data%base)}, a...)
		data = data / base
	}
	return a
}
