package time

import "time"

type Stub struct {
	T time.Time
}

func (s Stub) Now() time.Time {
	return s.T
}
