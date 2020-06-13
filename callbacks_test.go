package callbacks

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCallbacks(t *testing.T) {
	type EventOne struct{}
	type EventTwo struct{}

	t.Run("callbacks which are registered are called", func(t *testing.T) {
		called := false

		funcOne := func(ctx context.Context, event EventOne) error {
			called = true
			return nil
		}

		cb := Create()
		cb.Add(funcOne)

		err := cb.Call(context.Background(), EventOne{})

		assert.NoError(t, err)
		assert.True(t, called)
	})

	t.Run("callbacks which are called that error are returned up", func(t *testing.T) {
		expectedError := errors.New("callback error")

		funcOne := func(ctx context.Context, event EventOne) error {
			return expectedError
		}

		cb := Create()
		cb.Add(funcOne)

		err := cb.Call(context.Background(), EventOne{})

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})

	t.Run("callbacks which are not registered do not call registered callbacks", func(t *testing.T) {
		called := false

		funcOne := func(ctx context.Context, event EventOne) error {
			called = true
			return nil
		}

		cb := Create()
		cb.Add(funcOne)

		err := cb.Call(context.Background(), EventTwo{})

		assert.NoError(t, err)
		assert.False(t, called)
	})
}
