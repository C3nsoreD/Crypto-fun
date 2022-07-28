package basics

import (
	"context"
	"errors"
	"sync"
	"time"
)

type Circuit func(context.Context) (string, error)

func Breaker(circuit Circuit, thershold int) Circuit {
	var consecutiveFailures int = 0
	var lastAttempt = time.Now()
	var m sync.RWMutex

	return func(ctx context.Context) (string, error) {
		m.RLock() // lock reads

		d := consecutiveFailures - int(thershold)

		if d >= 0 {
			shouldRetryAt := lastAttempt.Add(time.Second * 2 << d) // use expoonential backoff [simple not ideal]
			if !time.Now().After(shouldRetryAt) {
				m.RUnlock()
				return "", errors.New("service unreachable")
			}
		}
		m.RUnlock()

		response, err := circuit(ctx)
		m.Lock()

		defer m.Unlock()

		lastAttempt = time.Now()

		if err != nil {
			consecutiveFailures++
			return response, err
		}
		consecutiveFailures = 0
		return response, nil
	}
}
