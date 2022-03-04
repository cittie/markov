package markov

import (
	"math/rand"
	"sync"
	"time"
)

var rPool = sync.Pool{New: func() interface{} {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}}

func Get() *rand.Rand {
	return rPool.Get().(*rand.Rand)
}

func Put(r *rand.Rand) {
	rPool.Put(r)
}

func Int63n(n int64) int64 {
	r := Get()
	defer Put(r)
	return r.Int63n(n)
}
