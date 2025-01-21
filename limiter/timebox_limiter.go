package limiter

import "time"

// TimeBoxedLimiter implements rate limiting with a specified time window per client.
// It uses a Timer to reference the current time and resets limits after the set period.
type TimeBoxedLimiter struct {
	refTime time.Time
	timer   Timer
	limiter RateLimiter
	period  time.Duration
	limits  map[string]int
}

// SetMaxCallsForClient sets the maximum number of allowed calls for a specific client within the time box period.
func (l *TimeBoxedLimiter) SetMaxCallsForClient(clientName string, max int) {
	l.limiter.SetMaxCallsForClient(clientName, max)
	l.limits[clientName] = max
}

// Allow determines if the specified client is allowed to proceed based on time-boxed rate limiting rules.
// Resets the client's limit if the set time box period has elapsed.
func (l *TimeBoxedLimiter) Allow(clientName string) bool {
	t := l.timer.Now()
	elapseTime := t.Sub(l.refTime)
	if elapseTime >= l.period {
		limit := l.limits[clientName]
		l.limiter.SetMaxCallsForClient(clientName, limit)
	}
	l.refTime = t
	return l.limiter.Allow(clientName)
}

func NewTimeBoxedLimiter(timer Timer, limiter *CounterLimiter, period time.Duration) *TimeBoxedLimiter {
	now := timer.Now()
	return &TimeBoxedLimiter{
		refTime: now,
		timer:   timer,
		limiter: limiter,
		period:  period,
		limits:  make(map[string]int),
	}
}
