package metrics

import "sync"

type Stats struct {
	Processed int `json:"processed"`
	Failed    int `json:"failed"`
	Queued    int `json:"queued"`
}

var mu sync.Mutex
var M = &Stats{}

func IncProcessed() {
	mu.Lock()
	defer mu.Unlock()
	M.Processed++
}

func IncFailed() {
	mu.Lock()
	defer mu.Unlock()
	M.Failed++
}

func IncQueued() {
	mu.Lock()
	defer mu.Unlock()
	M.Queued++
}
