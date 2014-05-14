package bytepool

import (
	"sync/atomic"
)

type ArbitraryPool struct {
	misses    int32
	list      chan *ArbitraryItem
	createNew func() interface{}
	resetItem func(item interface{}) error
}

func NewArbitrary(count int, createNew func() interface{}, resetItem func(item interface{}) error) *ArbitraryPool {
	p := &ArbitraryPool{
		list:      make(chan *ArbitraryItem, count),
		createNew: createNew,
		resetItem: resetItem,
	}
	for i := 0; i < count; i++ {
		p.list <- newArbitraryItem(p, p.createNew())
	}
	return p
}

func (pool *ArbitraryPool) Checkout() *ArbitraryItem {
	var item *ArbitraryItem
	select {
	case item = <-pool.list:
	default:
		atomic.AddInt32(&pool.misses, 1)
		item = newArbitraryItem(nil, pool.createNew())
	}
	return item
}

func (pool *ArbitraryPool) Len() int {
	return len(pool.list)
}

func (pool *ArbitraryPool) Misses() int32 {
	return atomic.LoadInt32(&pool.misses)
}
