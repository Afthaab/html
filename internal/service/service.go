package service

import (
	"context"
	"errors"

	"github.com/afthaab/job-portal/internal/auth"
	"github.com/afthaab/job-portal/internal/cache"
	"github.com/afthaab/job-portal/internal/models"
	newModels "github.com/afthaab/job-portal/internal/models/requestModels"
	"github.com/afthaab/job-portal/internal/repository"
)

type Service struct {
	UserRepo repository.UserRepo
	auth     auth.Authentication
	rdb      cache.Caching
}

//go:generate mockgen -source=service.go -destination=mockmodels/service_mock.go -package=mockmodels

type UserService interface {
	UserSignup(ctx context.Context, userData models.NewUser) (models.User, error)
	UserSignIn(ctx context.Context, userData models.NewUser) (string, error)

	AddCompanyDetails(ctx context.Context, companyData models.Company) (models.Company, error)
	ViewAllCompanies(ctx context.Context) ([]models.Company, error)
	ViewCompanyDetails(ctx context.Context, cid uint64) (models.Company, error)
	ViewJob(ctx context.Context, cid uint64) ([]models.Jobs, error)

	AddJobDetails(ctx context.Context, jobData newModels.NewJobs, cid uint64) (newModels.ResponseNewJobs, error)
	ViewAllJobs(ctx context.Context) ([]models.Jobs, error)
	ViewJobById(ctx context.Context, jid uint64) (models.Jobs, error)

	ProccessApplication(ctx context.Context, applicationData []newModels.NewUserApplication) ([]newModels.NewUserApplication, error)
	SendOtp(ctx context.Context, userData newModels.GetEmail) error
	VerifyOtp(ctx context.Context, userData newModels.GetVerifyOtp) error
}

func NewService(userRepo repository.UserRepo, a auth.Authentication, rdb cache.Caching) (UserService, error) {
	if userRepo == nil {
		return nil, errors.New("interface cannot be null")
	}
	return &Service{
		UserRepo: userRepo,
		auth:     a,
		rdb:      rdb,
	}, nil
}
