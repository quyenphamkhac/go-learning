package memo

import "sync"

// Memo caches result of calling a Func
type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]result
}

// Func is type of the function helps memorize
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// New memo
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// NOTE: not concurrency-safe

// Get a result with a key
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	defer memo.mu.Unlock()
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}
