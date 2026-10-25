package service

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type stubRepo struct {
	users     map[string]User
	err       error
	calls     int
	lastEmail string
}

func (r *stubRepo) ByEmail(email string) (User, error) {
	r.calls++
	r.lastEmail = email
	if r.err != nil {
		return User{}, r.err
	}
	u, ok := r.users[email]
	if !ok {
		return User{}, ErrNotFound
	}
	return u, nil
}
func TestService_FindIDByEmail(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		r := &stubRepo{users: map[string]User{"a@example.com": {ID: 42, Email: "a@example.com"}}}
		got, e := New(r).FindIDByEmail("a@example.com")
		require.NoError(t, e)
		assert.Equal(t, int64(42), got)
		assert.Equal(t, 1, r.calls)
		assert.Equal(t, "a@example.com", r.lastEmail)
	})
	t.Run("missing", func(t *testing.T) {
		r := &stubRepo{}
		got, e := New(r).FindIDByEmail("missing")
		assert.Zero(t, got)
		assert.ErrorIs(t, e, ErrNotFound)
	})
	t.Run("wrapped", func(t *testing.T) {
		r := &stubRepo{err: fmt.Errorf("repo: %w", ErrNotFound)}
		_, e := New(r).FindIDByEmail("a")
		assert.ErrorIs(t, e, ErrNotFound)
	})
	t.Run("repository error", func(t *testing.T) {
		want := errors.New("unavailable")
		r := &stubRepo{err: want}
		got, e := New(r).FindIDByEmail("a")
		assert.Zero(t, got)
		assert.ErrorIs(t, e, want)
	})
}
