package structpadding

import "sync/atomic"

type NoPad struct {
	a uint64
	b uint64
	c uint64
}

func (myatomic *NoPad) IncreaseAllEles() {
	atomic.AddUint64(&myatomic.a, 1)
	atomic.AddUint64(&myatomic.b, 1)
	atomic.AddUint64(&myatomic.c, 1)
}

// adding padding based on the cache line of the cpu
//makes it faster because it prevents false sharing and cache evicting
// the cache line is 64 bytes
// uint64 is 4 bytes, therefore we will the rest of the cache line with 60 bytes
type Pad struct {
	a   uint64
	_p1 [60]byte
	b   uint64
	_p2 [60]byte
	c   uint64
	_p3 [60]byte
}

func (myatomic *Pad) IncreaseAllEles() {
	atomic.AddUint64(&myatomic.a, 1)
	atomic.AddUint64(&myatomic.b, 1)
	atomic.AddUint64(&myatomic.c, 1)
}