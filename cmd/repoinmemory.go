package cmd

import (
	"github.com/gopheramit/pomo/pomodoro"
	"github.com/gopheramit/pomo/repository"
)

func getRepo() (pomodoro.Repository, error) {
	return repository.NewInMemoryRepo(), nil
}
