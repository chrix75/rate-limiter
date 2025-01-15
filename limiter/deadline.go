package limiter

import "time"

type DayLimiter struct {
	refTime time.Time
	limiter *RateLimiter
	limits  map[string]int
}

func (l *DayLimiter) SetMaxCallsForClient(clientName string, max int) {
	l.limiter.SetMaxCallsForClient(clientName, max)
	l.limits[clientName] = max
}

func (l *DayLimiter) Allow(t time.Time, clientName string) bool {
	refDay := time.Date(l.refTime.Year(), l.refTime.Month(), l.refTime.Day(), 0, 0, 0, 0, l.refTime.Location())
	curDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	if curDay.After(refDay) {
		l.SetMaxCallsForClient(clientName, l.limits[clientName])
	}
	return l.limiter.Allow(clientName)
}

func NewDayLimiter(refTime time.Time, limiter *RateLimiter) *DayLimiter {
	return &DayLimiter{
		limiter: limiter,
		limits:  make(map[string]int),
		refTime: refTime,
	}
}
