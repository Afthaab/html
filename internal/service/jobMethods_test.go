package service

// import (
// 	"context"
// 	"errors"
// 	"reflect"
// 	"testing"

// 	"github.com/afthaab/job-portal/internal/auth"
// 	"github.com/afthaab/job-portal/internal/cache"
// 	newModels "github.com/afthaab/job-portal/internal/models/requestModels"
// 	mockrepository "github.com/afthaab/job-portal/internal/repository/mockModels"
// 	"go.uber.org/mock/gomock"
// )

// func TestService_AddJobDetails(t *testing.T) {
// 	type args struct {
// 		ctx         context.Context
// 		bodyjobData newModels.NewJobs
// 		cid         uint64
// 	}
// 	tests := []struct {
// 		name             string
// 		args             args
// 		want             newModels.ResponseNewJobs
// 		wantErr          bool
// 		mockRepoResponse func() (newModels.ResponseNewJobs, error)
// 	}{
// 		{
// 			name: "error from repository mock function",
// 			want: newModels.ResponseNewJobs{
// 				Jobid: 1,
// 			},
// 			args: args{
// 				ctx: context.Background(),
// 				bodyjobData: newModels.NewJobs{
// 					Jobname:         "new",
// 					MinNoticePeriod: uint(10),
// 					MaxNoticePeriod: uint(20),
// 					Location: []uint{
// 						uint(1), uint(2),
// 					},
// 					TechnologyStack: []uint{
// 						uint(1), uint(2),
// 					},
// 					Description:   "description",
// 					MinExperience: uint(5),
// 					MaxExperience: uint(10),
// 					Qualifications: []uint{
// 						uint(1), uint(2),
// 					},
// 					Shift: []uint{
// 						uint(1), uint(2),
// 					},
// 					Jobtype: "no type",
// 				},
// 				cid: 1,
// 			},
// 			wantErr: true,
// 			mockRepoResponse: func() (newModels.ResponseNewJobs, error) {
// 				return newModels.ResponseNewJobs{
// 					Jobid: 1,
// 				}, errors.New("test case error from mock function")
// 			},
// 		},
// 		{
// 			name: "success case",
// 			want: newModels.ResponseNewJobs{},
// 			args: args{
// 				ctx: context.Background(),
// 				bodyjobData: newModels.NewJobs{
// 					Jobname:         "new",
// 					MinNoticePeriod: uint(10),
// 					MaxNoticePeriod: uint(20),
// 					Location: []uint{
// 						uint(1), uint(2),
// 					},
// 					TechnologyStack: []uint{
// 						uint(1), uint(2),
// 					},
// 					Description:   "description",
// 					MinExperience: uint(5),
// 					MaxExperience: uint(10),
// 					Qualifications: []uint{
// 						uint(1), uint(2),
// 					},
// 					Shift: []uint{
// 						uint(1), uint(2),
// 					},
// 					Jobtype: "no type",
// 				},
// 				cid: 1,
// 			},
// 			wantErr: false,
// 			mockRepoResponse: func() (newModels.ResponseNewJobs, error) {
// 				return newModels.ResponseNewJobs{}, nil
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mc := gomock.NewController(t)
// 			mockRepo := mockrepository.NewMockUserRepo(mc)
// 			mockRepo.EXPECT().CreateJob(gomock.Any(), gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()

// 			svc, err := NewService(mockRepo, &auth.Auth{}, &cache.RDBLayer{})
// 			if err != nil {
// 				t.Errorf("error is initializing the repo layer")
// 				return
// 			}
// 			got, err := svc.AddJobDetails(tt.args.ctx, tt.args.bodyjobData, tt.args.cid)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Service.AddJobDetails() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Service.AddJobDetails() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
