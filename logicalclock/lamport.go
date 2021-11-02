package logicalclock

import (
	"log"
	"sync"
)

type LamportTimer interface {
	Increment()
	Read() int64
	Update(other LamportTimer)
}

func NewLamportClock(ts int64) *LamportClock {
	return &LamportClock{timestamp: ts}
}

type LamportClock struct {
	timestamp int64
	mutex      sync.Mutex
}

func (this *LamportClock) Increment() {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	
	log.Println("Time is passing...")
	this.timestamp++
}

func (this *LamportClock) Read() int64 {
	return this.timestamp
}

func (this *LamportClock) synchronize(other LamportTimer) {
	if other.Read() > this.Read() {
		this.timestamp = other.Read()
	}
}

func (this *LamportClock) Update(other LamportTimer) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	this.synchronize(other)
	this.timestamp++
}
