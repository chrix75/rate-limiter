package limiter

import "time"

type TimeBoxedLimiter struct {
	refTime time.Time
	timer   Timer
	limiter RateLimiter
	period  time.Duration
	limits  map[string]int
}

func (l *TimeBoxedLimiter) SetMaxCallsForClient(clientName string, max int) {
	l.limiter.SetMaxCallsForClient(clientName, max)
	l.limits[clientName] = max
}

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
