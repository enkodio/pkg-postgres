package logic

import "time"

const (
	defaultMaxOpenConns = 4
	defaultMaxDelay     = 5
	defaultMaxAttempts  = 5

	maxConnIdleTime = time.Second * 5
)

func doWithAttempts(fn func() error, maxAttempts int, delay int) error {
	var err error
	if maxAttempts == 0 {
		maxAttempts = defaultMaxAttempts
	}
	if delay == 0 {
		delay = defaultMaxDelay
	}
	for maxAttempts > 0 {
		if err = fn(); err != nil {
			time.Sleep(time.Second * time.Duration(delay))
			maxAttempts--
			continue
		}
		return nil
	}
	return err
}
