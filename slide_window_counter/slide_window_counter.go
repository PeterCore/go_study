package slide_window_counter

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

var (
	once sync.Once
)

type SlidingWindowCounter struct {
	incurRequests    int32
	durationRequests chan int32
	accuracy         time.Duration
	snippet          time.Duration
	currentRequests  int32
	allowRequests    int32
}

func New(accuracy time.Duration, snippet time.Duration, allowRequests int32) *SlidingWindowCounter {
	return &SlidingWindowCounter{durationRequests: make(chan int32, snippet/accuracy/1000), accuracy: accuracy, snippet: snippet, allowRequests: allowRequests}
}

func (l *SlidingWindowCounter) Take() error {
	once.Do(func() {
		go sliding(l)
		go calculate(l)
	})
	curRequest := atomic.LoadInt32(&l.currentRequests)
	if curRequest >= l.allowRequests {
		return errors.New("exceed limit")
	}
	if !atomic.CompareAndSwapInt32(&l.currentRequests, curRequest, curRequest+1) {
		return errors.New("exceed limit")
	}
	atomic.AddInt32(&l.incurRequests, 1)
	return nil

}

func sliding(l *SlidingWindowCounter) {
	for {
		select {
		case <-time.After(l.accuracy):
			t := atomic.SwapInt32(&l.incurRequests, 0)
			l.durationRequests <- t
		}
	}
}

func calculate(l *SlidingWindowCounter) {
	for {
		<-time.After(l.accuracy)
		if len(l.durationRequests) == cap(l.durationRequests) {
			break
		}
	}
	for {
		<-time.After(l.accuracy)
		t := <-l.durationRequests
		if t != 0 {
			atomic.AddInt32(&l.currentRequests, -t)
		}
	}
}
