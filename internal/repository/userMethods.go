package repository

import (
	"context"
	"errors"

	"github.com/afthaab/job-portal/internal/models"
	"github.com/rs/zerolog/log"
)

func (r *Repo) CreateUser(ctx context.Context, UserDetails models.User) (models.User, error) {
	result := r.db.Create(&UserDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.User{}, errors.New("could not create the user")
	}
	return UserDetails, nil
}

func (r *Repo) UpdatePassword(ctx context.Context, email string, psswd string) error {
	var userDetails models.User
	result := r.db.Model(&userDetails).Where("email = ?", email).Update("password_hash", psswd)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return errors.New("could not upadate the password")
	}
	return nil
}

func (r *Repo) CheckEmail(ctx context.Context, email string) (models.User, error) {
	var userDetails models.User
	result := r.db.Where("email = ?", email).First(&userDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.User{}, errors.New("email not found")
	}
	return userDetails, nil

}
