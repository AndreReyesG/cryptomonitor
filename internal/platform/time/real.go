package time

import "time"

type Real struct{}

func (r Real) Now() time.Time {
	return time.Now()
}
