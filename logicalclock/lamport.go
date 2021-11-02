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
	mutex     sync.Mutex
}

func (this *LamportClock) Increment() {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.increment()
}

func (this *LamportClock) increment() {
	this.timestamp++
	log.Printf("Incrementing Lamport Timestamp to %d", this.Read())
}

func (this *LamportClock) Read() int64 {
	return this.timestamp
}

func (this *LamportClock) synchronize(other LamportTimer) {
	if other.Read() > this.Read() {
		log.Printf("Received timestamp %d which is larger than current %d. Updating internal timestamp to %d", other.Read(), this.Read(), other.Read())
		this.timestamp = other.Read()
	} else {
		log.Printf("Received timestamp %d which is not larger than current %d. Keeping internal timestamp at %d", other.Read(), this.Read(), this.Read())
	}
}

func (this *LamportClock) Update(other LamportTimer) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	this.synchronize(other)
	this.increment()
}
