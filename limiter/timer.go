package limiter

import "time"

// Timer defines an interface for obtaining the current time.
type Timer interface {
	Now() time.Time
}

type FixedValueTimer struct {
	t time.Time
}

func (t FixedValueTimer) Now() time.Time {
	return t.t
}

func NewFixedValueTimer(t time.Time) FixedValueTimer {
	return FixedValueTimer{t: t}
}

type DynamicTimer struct {
	T time.Time
}

func (t *DynamicTimer) Now() time.Time {
	return t.T
}

func NewDynamicTimer(t time.Time) *DynamicTimer {
	return &DynamicTimer{T: t}
}
