package pomodoro_test

import (
	"testing"

	"github.com/gopheramit/pomo/pomodoro"
	"github.com/gopheramit/pomo/repository"
)

func getRepo(t *testing.T) (pomodoro.Repository, func()) {
	t.Helper()
	return repository.NewInMemoryRepo(), func() {}

}
