package mthasher

import (
	"context"
	"crypto/sha256"
)

const start = 97 // utf8 encoded of "a"

type iHashFunc func([]byte) [32]byte

type Thread struct {
	answerChan chan [32]byte   // buffered answer channel
	hashFunc   iHashFunc       // hash function
	ctx        context.Context // goroutine context for goroutine termination
	base       int             // base
	len        int             // length of password
	span       [2]int          // span for thread
	id         uint8           // thread identifier
}

func NewThread(
	id uint8, // thread identifier
	answerChan chan [32]byte, // buffered answer channel
	hashFunc iHashFunc, // hash function
	ctx context.Context, // goroutine context for goroutine termination
	base int, // base
	len int, // length of password
	span [2]int, // span for thread
) *Thread {
	return &Thread{
		id:         id,
		answerChan: answerChan,
		hashFunc:   hashFunc,
		ctx:        ctx,
		base:       base,
		len:        len,
		span:       span,
	}
}

func (t *Thread) Start() {
	for w := t.span[0]; w <= t.span[1]; w++ {
		select {
		case <-t.ctx.Done():
			return
		default:
			t.answerChan <- t.hashFunc(word(w, t.base, t.len))
		}
	}
}

func SHA256(data []byte) [32]byte {
	return sha256.Sum256(data)
}

func word(data int, base int, len int) []byte {
	a := make([]byte, 0)
	for i := 0; i < len; i++ {
		a = append([]byte{start + byte(data%base)}, a...)
		data = data / base
	}
	return a
}
