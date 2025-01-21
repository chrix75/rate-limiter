package limiter

import "time"

// Timer defines an interface for obtaining the current time.
type Timer interface {
	Now() time.Time
}

type ConstantTimer struct {
	t time.Time
}

func (t ConstantTimer) Now() time.Time {
	return t.t
}

func NewConstantTimer(t time.Time) ConstantTimer {
	return ConstantTimer{t: t}
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
