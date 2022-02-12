package unittest

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

// CallMustReturnWithinTimeout is a test helper that invokes the given function and fails the test if the invocation
// does not return prior to the given timeout.
func CallMustReturnWithinTimeout(t *testing.T, f func(), timeout time.Duration, failureMsg string) {
	done := make(chan interface{})

	go func() {
		f()

		close(done)
	}()

	select {
	case <-done:
		return
	case <-time.After(timeout):
		require.Failf(t, "function did not return on time: %s", failureMsg)
	}
}
