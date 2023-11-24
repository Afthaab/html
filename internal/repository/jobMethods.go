package repository

import (
	"context"
	"errors"

	"github.com/afthaab/job-portal/internal/models"
	newModels "github.com/afthaab/job-portal/internal/models/requestModels"
	"github.com/rs/zerolog/log"
)

func (r *Repo) GetTheJobData(jobid uint) (models.Jobs, error) {
	var jobData models.Jobs

	// Preload related data using GORM's Preload method
	result := r.db.Preload("Company").
		Preload("Location").
		Preload("TechnologyStack").
		Preload("Qualifications").
		Preload("Shift").
		Where("id = ?", jobid).
		Find(&jobData)

	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Jobs{}, result.Error
	}

	return jobData, nil
}

func (r *Repo) CreateJob(ctx context.Context, jobData models.Jobs) (newModels.ResponseNewJobs, error) {
	result := r.db.Create(&jobData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return newModels.ResponseNewJobs{}, errors.New("could not create the jobs")
	}
	return newModels.ResponseNewJobs{
		Jobid: jobData.ID,
	}, nil
}

func (r *Repo) ViewJobDetailsBy(ctx context.Context, jid uint64) (models.Jobs, error) {
	var jobData models.Jobs
	result := r.db.Where("id = ?", jid).Find(&jobData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Jobs{}, errors.New("could not create the jobs")
	}
	return jobData, nil
}

func (r *Repo) FindAllJobs(ctx context.Context) ([]models.Jobs, error) {
	var jobDatas []models.Jobs
	result := r.db.Find(&jobDatas)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("could not find the jobs")
	}
	return jobDatas, nil

}

func (r *Repo) FindJob(ctx context.Context, cid uint64) ([]models.Jobs, error) {
	var jobData []models.Jobs
	result := r.db.Where("cid = ?", cid).Find(&jobData)
	if result.Error != nil || result.RowsAffected == 0 {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("could not find the company")
	}
	return jobData, nil
}
