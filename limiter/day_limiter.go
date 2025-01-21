package limiter

import "time"

// DayLimiter provides day-based rate limiting functionality for clients by resetting limits based on the current day.
// It leverages a Timer and a RateLimiter to manage time tracking and rate control, respectively.
// It supports setting max daily calls for each client and checks if a client is allowed based on current limits.
type DayLimiter struct {
	refTime time.Time
	timer   Timer
	limiter RateLimiter
	limits  map[string]int
}

// SetMaxCallsForClient sets the maximum number of daily calls allowed for a specific client in the DayLimiter instance.
func (l *DayLimiter) SetMaxCallsForClient(clientName string, max int) {
	l.limiter.SetMaxCallsForClient(clientName, max)
	l.limits[clientName] = max
}

// Allow determines if a client is permitted to proceed, resetting daily limits if the current day has changed.
func (l *DayLimiter) Allow(clientName string) bool {
	t := l.timer.Now()
	refDay := time.Date(l.refTime.Year(), l.refTime.Month(), l.refTime.Day(), 0, 0, 0, 0, l.refTime.Location())
	curDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	if curDay.After(refDay) {
		l.SetMaxCallsForClient(clientName, l.limits[clientName])
	}

	return l.limiter.Allow(clientName)
}

func NewDayLimiter(timer Timer, limiter RateLimiter) *DayLimiter {
	return &DayLimiter{
		limiter: limiter,
		limits:  make(map[string]int),
		refTime: timer.Now(),
		timer:   timer,
	}
}
