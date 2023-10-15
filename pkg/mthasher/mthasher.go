package mthasher

import "context"

type MultithreadHasher struct {
	threads  []*Thread
	ctx      context.Context
	close    func()
	hashFunc iHashFunc
	answerChan chan 
}

func (h *MultithreadHasher) Add(thread *Thread) {
	h.threads = append(h.threads, thread)
}

func (h *MultithreadHasher) Run() {
	for _, thread := range h.threads {
		go thread.Start()
	}
}
