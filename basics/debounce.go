// Debounce stability pattern
package basics

import (
	"context"
	"sync"
	"time"
)

type Circuit func(context.Context) (string, error)

func DebounceFirst(ciruit Circuit, d time.Duration) Circuit {
	var thershold time.Time
	var result string
	var err error
	var m sync.Mutex

	return func(ctx context.Context) (string, error) {
		m.Lock()

		defer func() {
			thershold = time.Now().Add(d)
			m.Unlock()
		}()
		if time.Now().Before(thershold) { // ignore calls after that aren't before threshold
			return result, err
		}

		result, err = ciruit(ctx)

		return result, err
	}
}
