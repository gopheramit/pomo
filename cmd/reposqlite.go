//go:build !inmemory
// +build !inmemory

package cmd

import (
	"github.com/gopheramit/pomo/pomodoro"
	"github.com/gopheramit/pomo/repository"
	"github.com/spf13/viper"
)

func getRepo() (pomodoro.Repository, error) {
	repo, err := repository.NewSQLite3Repo(viper.GetString("db"))
	if err != nil {
		return nil, err
	}
	return repo, nil
}
