package mthasher

import "context"

type MultithreadHasher struct {
	Close      func()
	AnswerChan chan [32]byte
	threads    []*Thread
	hashFunc   iHashFunc
	ctx        context.Context
	base       int
	len        int
	cid        uint8
}

func NewMultithreadHasher(close func(), hashFunc iHashFunc, answerChan chan [32]byte, ctx context.Context, base int, len int) *MultithreadHasher {
	return &MultithreadHasher{
		Close:      close,
		AnswerChan: answerChan,
		hashFunc:   hashFunc,
		ctx:        ctx,
		base:       base,
		len:        len,
	}
}

func (h *MultithreadHasher) Add(span [2]int) {
	h.threads = append(h.threads, NewThread(h.cid, h.AnswerChan, h.hashFunc, h.ctx, h.base, h.len, span))
	h.cid++
}

func (h *MultithreadHasher) Run() {
	for _, thread := range h.threads {
		go thread.Start()
	}
}
