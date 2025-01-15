package limiter

import "time"

type DayLimiter struct {
	refTime        time.Time
	limiter        *RateLimiter
	timeBoxLimiter *TimeBoxedLimiter
	limits         map[string]int
}

func (l *DayLimiter) SetMaxCallsForClient(clientName string, max int) {
	if l.limiter != nil {
		l.limiter.SetMaxCallsForClient(clientName, max)
	}

	if l.timeBoxLimiter != nil {
		l.timeBoxLimiter.SetMaxCallsForClient(clientName, max)
	}

	l.limits[clientName] = max
}

func (l *DayLimiter) Allow(t time.Time, clientName string) bool {
	refDay := time.Date(l.refTime.Year(), l.refTime.Month(), l.refTime.Day(), 0, 0, 0, 0, l.refTime.Location())
	curDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	if curDay.After(refDay) {
		l.SetMaxCallsForClient(clientName, l.limits[clientName])
	}

	if l.limiter != nil {
		return l.limiter.Allow(clientName)
	}

	return l.timeBoxLimiter.Allow(t, clientName)
}

func NewDayLimiter(refTime time.Time, limiter *RateLimiter) *DayLimiter {
	return &DayLimiter{
		limiter: limiter,
		limits:  make(map[string]int),
		refTime: refTime,
	}
}

func NewDayLimiterWithTimeBox(limiter *TimeBoxedLimiter) *DayLimiter {
	return &DayLimiter{
		timeBoxLimiter: limiter,
		limits:         make(map[string]int),
		refTime:        limiter.refTime,
	}
}
