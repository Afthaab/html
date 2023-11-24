package service

import (
	"context"

	"github.com/afthaab/job-portal/internal/models"
	newModels "github.com/afthaab/job-portal/internal/models/requestModels"
	"gorm.io/gorm"
)

func (s *Service) AddJobDetails(ctx context.Context, bodyjobData newModels.NewJobs, cid uint64) (newModels.ResponseNewJobs, error) {
	jobData := models.Jobs{
		Cid:             uint(cid),
		Jobname:         bodyjobData.Jobname,
		MinNoticePeriod: bodyjobData.MinNoticePeriod,
		MaxNoticePeriod: bodyjobData.MaxNoticePeriod,
		Description:     bodyjobData.Description,
		Jobtype:         bodyjobData.Jobtype,
		MinExperience:   bodyjobData.MinExperience,
		MaxExperience:   bodyjobData.MaxExperience,
	}
	jobData.Location = getLocations(bodyjobData.Location)

	for _, v := range bodyjobData.Qualifications {
		tempQualifications := models.Qualification{
			Model: gorm.Model{
				ID: v,
			},
		}
		jobData.Qualifications = append(jobData.Qualifications, tempQualifications)
	}
	for _, v := range bodyjobData.TechnologyStack {
		tempTechStack := models.TechnologyStack{
			Model: gorm.Model{
				ID: v,
			},
		}
		jobData.TechnologyStack = append(jobData.TechnologyStack, tempTechStack)
	}
	for _, v := range bodyjobData.Shift {
		tempShift := models.Shift{
			Model: gorm.Model{
				ID: v,
			},
		}
		jobData.Shift = append(jobData.Shift, tempShift)
	}

	responseData, err := s.UserRepo.CreateJob(ctx, jobData)
	if err != nil {
		return responseData, err
	}
	return responseData, nil
}

func getLocations(locationIds []uint) (locationData []models.Location) {
	for _, v := range locationIds {
		tempLocation := models.Location{
			Model: gorm.Model{
				ID: v,
			},
		}
		locationData = append(locationData, tempLocation)
	}
	return locationData
}

func (s *Service) ViewJobById(ctx context.Context, jid uint64) (models.Jobs, error) {
	jobData, err := s.UserRepo.ViewJobDetailsBy(ctx, jid)
	if err != nil {
		return models.Jobs{}, err
	}
	return jobData, nil
}

func (s *Service) ViewAllJobs(ctx context.Context) ([]models.Jobs, error) {
	jobDatas, err := s.UserRepo.FindAllJobs(ctx)
	if err != nil {
		return nil, err
	}
	return jobDatas, nil

}

func (s *Service) ViewJob(ctx context.Context, cid uint64) ([]models.Jobs, error) {
	jobData, err := s.UserRepo.FindJob(ctx, cid)
	if err != nil {
		return nil, err
	}
	return jobData, nil
}
