package limiter

import "sync"

type InMemoryLimitRepo struct {
	limits map[string]int
}

func (r *InMemoryLimitRepo) DecAndGet(clientName string) int {
	v := r.limits[clientName]
	if v < 1 {
		return 0
	}
	nu := v - 1
	r.limits[clientName] = nu
	return v
}

func (r *InMemoryLimitRepo) AddClient(clientName string, maxCalls int) {
	if r.limits == nil {
		r.limits = map[string]int{}
	}
	r.limits[clientName] = maxCalls
}

type LimitRepository interface {
	AddClient(clientName string, maxCalls int)
	DecAndGet(clientName string) int
}

type ConcurrentLimitRepository struct {
	mu   sync.Mutex
	repo LimitRepository
}

func (r *ConcurrentLimitRepository) AddClient(clientName string, maxCalls int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.repo.AddClient(clientName, maxCalls)
}

func (r *ConcurrentLimitRepository) DecAndGet(clientName string) int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.repo.DecAndGet(clientName)
}
