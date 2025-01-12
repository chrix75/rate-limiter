package limiter

type InMemoryLimitRepo struct {
	limits map[string]int
}

func (r *InMemoryLimitRepo) AddClient(clientName string, maxCalls int) {
	if r.limits == nil {
		r.limits = map[string]int{}
	}
	r.limits[clientName] = maxCalls
}

func (r *InMemoryLimitRepo) Get(clientName string) int {
	return r.limits[clientName]
}
