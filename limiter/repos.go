package limiter

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
