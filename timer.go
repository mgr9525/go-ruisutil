package ruisUtil

import "time"

type Timer struct {
	dur time.Duration
	tms time.Time
}

func NewTimer(dur time.Duration) *Timer {
	return &Timer{
		dur: dur,
	}
}

func (c *Timer) Reinit() {
	c.tms = time.Time{}
}
func (c *Timer) Reset() {
	c.tms = time.Now()
}
func (c *Timer) Tick() bool {
	if time.Since(c.tms) > c.dur {
		c.Reset()
		return true
	}
	return false
}
func (c *Timer) Tmout() bool {
	return time.Since(c.tms) > c.dur
}
func (c *Timer) SetDur(dur time.Duration) {
	c.dur = dur
}
func (c *Timer) GetDur() time.Duration {
	if time.Since(c.tms) > c.dur {
		return c.dur
	}
	return 0
}
