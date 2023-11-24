package service

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/afthaab/job-portal/internal/models"
	newModels "github.com/afthaab/job-portal/internal/models/requestModels"
	"github.com/afthaab/job-portal/internal/pkg"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func (s *Service) VerifyOtp(ctx context.Context, userData newModels.GetVerifyOtp) error {
	if userData.ConfirmPassword != userData.NewPassword {
		return errors.New("the passwords did not match")
	}

	otp, err := s.rdb.CheckTheOtp(ctx, userData.Email)
	if err != nil {
		return err
	}
	if otp != userData.Otp {
		return errors.New("the otp did not match")
	}

	err = s.rdb.DeleteTheCache(ctx, userData.Email)
	if err != nil {
		return err
	}
	psswd, err := pkg.HashPassword(userData.NewPassword)
	if err != nil {
		return err
	}
	err = s.UserRepo.UpdatePassword(ctx, userData.Email, psswd)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) SendOtp(ctx context.Context, userData newModels.GetEmail) error {
	// check the email and username from the database
	userDetails, err := s.UserRepo.CheckEmail(ctx, userData.Email)
	if err != nil {
		return err
	}

	if userDetails.Dob != userData.DateofBirth {
		return errors.New("the date of birth is invalid")
	}

	otp, err := pkg.GenerateOneTimePassword(userDetails.Email)
	if err != nil {
		log.Error().Err(err).Msg("could not generate the password")
		return err
	}

	err = s.rdb.CacheTheOtp(ctx, userDetails.Email, otp)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UserSignIn(ctx context.Context, userData models.NewUser) (string, error) {
	// checking the email in the db
	userDetails, err := s.UserRepo.CheckEmail(ctx, userData.Email)
	if err != nil {
		return "", err
	}

	// comaparing the password and hashed password
	err = pkg.CheckHashedPassword(userData.Password, userDetails.PasswordHash)
	if err != nil {
		log.Info().Err(err).Send()
		return "", errors.New("entered password is not wrong")
	}

	// setting up the claims
	claims := jwt.RegisteredClaims{
		Issuer:    "job portal project",
		Subject:   strconv.FormatUint(uint64(userDetails.ID), 10),
		Audience:  jwt.ClaimStrings{"users"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token, err := s.auth.GenerateAuthToken(claims)
	if err != nil {
		return "", err
	}

	return token, nil

}

func (s *Service) UserSignup(ctx context.Context, userData models.NewUser) (models.User, error) {
	hashedPass, err := pkg.HashPassword(userData.Password)
	if err != nil {
		return models.User{}, err
	}
	userDetails := models.User{
		Username:     userData.Username,
		Email:        userData.Email,
		PasswordHash: hashedPass,
		Dob:          userData.DateOfBirth,
	}

	userDetails, err = s.UserRepo.CreateUser(ctx, userDetails)
	if err != nil {
		return models.User{}, err
	}
	return userDetails, nil
}
