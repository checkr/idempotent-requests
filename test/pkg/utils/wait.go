package utils

import (
	"errors"
	"fmt"
	"github.com/avast/retry-go"
	"net/http"
	"time"
)

const Ready = "/-/ready"

func WaitForServerToBeReady(host string) {
	ReadinessURL := host + Ready

	err := retry.Do(
		func() error {
			resp, err := http.Get(ReadinessURL)

			if err != nil {
				return err
			}

			if resp.StatusCode == http.StatusOK {
				return nil
			}

			return errors.New("server is not ready yet")
		},
		retry.Attempts(20),
		retry.MaxDelay(5*time.Second),
		retry.OnRetry(func(n uint, err error) {
			fmt.Printf("server is not ready yet. re-try %d\n", n+1)
		}),
	)

	if err != nil {
		panic(err.Error())
	}
}
