package time

import "time"

type Provider interface {
	Now() time.Time
}
