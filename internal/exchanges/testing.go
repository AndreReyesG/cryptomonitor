package exchanges

import "time"

type StubTime struct {
	T time.Time
}

func (s StubTime) Now() time.Time {
	return s.T
}
