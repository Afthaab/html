package service

// import (
// 	"context"
// 	"errors"
// 	"reflect"
// 	"testing"

// 	"github.com/afthaab/job-portal/internal/auth"
// 	"github.com/afthaab/job-portal/internal/cache/mockmodels"
// 	"github.com/afthaab/job-portal/internal/models"
// 	newModels "github.com/afthaab/job-portal/internal/models/requestModels"
// 	cachMock "github.com/afthaab/job-portal/internal/repository/mockModels"
// 	"github.com/redis/go-redis/v9"
// 	"go.uber.org/mock/gomock"
// 	"gorm.io/gorm"
// )

// func TestService_ProccessApplication(t *testing.T) {
// 	type args struct {
// 		ctx             context.Context
// 		applicationData []newModels.NewUserApplication
// 	}
// 	tests := []struct {
// 		name                 string
// 		args                 args
// 		want                 []newModels.NewUserApplication
// 		mockRepoResponse     func() (models.Jobs, error)
// 		mockCacheResponse    func() (string, error)
// 		mockAddCacheResponse func() error
// 		wantErr              bool
// 	}{
// 		{
// 			name: "not found in redis and in database",
// 			args: args{
// 				ctx: context.Background(),
// 				applicationData: []newModels.NewUserApplication{
// 					{
// 						Name: "afthab",
// 						Age:  "21",
// 						Jid:  1,
// 						Jobs: newModels.Requestfield{
// 							Jobname:         "senior web developer",
// 							NoticePeriod:    7,
// 							Location:        []uint{1, 2},
// 							TechnologyStack: []uint{1, 2},
// 							Experience:      uint(5),
// 							Qualifications:  []uint{1, 2},
// 							Shift:           []uint{1, 2},
// 							Jobtype:         "permanent",
// 						},
// 					},
// 				},
// 			},
// 			want: nil,
// 			mockCacheResponse: func() (string, error) {
// 				return "", redis.Nil
// 			},
// 			mockRepoResponse: func() (models.Jobs, error) {
// 				return models.Jobs{}, errors.New("test error from cache mock")
// 			},
// 			wantErr: false,
// 			mockAddCacheResponse: func() error {
// 				return errors.New("test")
// 			},
// 		},
// 		{
// 			name: "not found in redis and found in database, and error in adding to the cache",
// 			args: args{
// 				ctx: context.Background(),
// 				applicationData: []newModels.NewUserApplication{
// 					{
// 						Name: "afthab",
// 						Age:  "21",
// 						Jid:  1,
// 						Jobs: newModels.Requestfield{
// 							Jobname:         "senior web developer",
// 							NoticePeriod:    7,
// 							Location:        []uint{1, 2},
// 							TechnologyStack: []uint{1, 2},
// 							Experience:      uint(5),
// 							Qualifications:  []uint{1, 2},
// 							Shift:           []uint{1, 2},
// 							Jobtype:         "permanent",
// 						},
// 					},
// 				},
// 			},
// 			want: nil,
// 			mockCacheResponse: func() (string, error) {
// 				return "", redis.Nil
// 			},
// 			mockAddCacheResponse: func() error {
// 				return errors.New("test error")
// 			},
// 			mockRepoResponse: func() (models.Jobs, error) {
// 				return models.Jobs{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Jobname:         "senior web developer",
// 					MinNoticePeriod: uint(5),
// 					MaxNoticePeriod: uint(10),
// 					Location: []models.Location{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					TechnologyStack: []models.TechnologyStack{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					MinExperience: uint(3),
// 					MaxExperience: uint(10),
// 					Qualifications: []models.Qualification{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					Shift: []models.Shift{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					Jobtype: "permanent",
// 				}, nil
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "no error in adding the cache and error in experience",
// 			args: args{
// 				ctx: context.Background(),
// 				applicationData: []newModels.NewUserApplication{
// 					{
// 						Name: "afthab",
// 						Age:  "21",
// 						Jid:  1,
// 						Jobs: newModels.Requestfield{
// 							Jobname:         "senior web developer",
// 							NoticePeriod:    7,
// 							Location:        []uint{1, 2},
// 							TechnologyStack: []uint{1, 2},
// 							Experience:      uint(2),
// 							Qualifications:  []uint{1, 2},
// 							Shift:           []uint{1, 2},
// 							Jobtype:         "permanent",
// 						},
// 					},
// 				},
// 			},
// 			want: nil,
// 			mockCacheResponse: func() (string, error) {
// 				return "", redis.Nil
// 			},
// 			mockAddCacheResponse: func() error {
// 				return nil
// 			},
// 			mockRepoResponse: func() (models.Jobs, error) {
// 				return models.Jobs{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Jobname:         "senior web developer",
// 					MinNoticePeriod: uint(5),
// 					MaxNoticePeriod: uint(10),
// 					Location: []models.Location{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					TechnologyStack: []models.TechnologyStack{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					MinExperience: uint(3),
// 					MaxExperience: uint(10),
// 					Qualifications: []models.Qualification{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					Shift: []models.Shift{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					Jobtype: "permanent",
// 				}, nil
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "error in notice period",
// 			args: args{
// 				ctx: context.Background(),
// 				applicationData: []newModels.NewUserApplication{
// 					{
// 						Name: "afthab",
// 						Age:  "21",
// 						Jid:  1,
// 						Jobs: newModels.Requestfield{
// 							Jobname:         "senior web developer",
// 							NoticePeriod:    0,
// 							Location:        []uint{1, 2},
// 							TechnologyStack: []uint{1, 2},
// 							Experience:      uint(5),
// 							Qualifications:  []uint{1, 2},
// 							Shift:           []uint{1, 2},
// 							Jobtype:         "permanent",
// 						},
// 					},
// 				},
// 			},
// 			want: nil,
// 			mockCacheResponse: func() (string, error) {
// 				return "", redis.Nil
// 			},
// 			mockAddCacheResponse: func() error {
// 				return nil
// 			},
// 			mockRepoResponse: func() (models.Jobs, error) {
// 				return models.Jobs{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Jobname:         "senior web developer",
// 					MinNoticePeriod: uint(5),
// 					MaxNoticePeriod: uint(10),
// 					Location: []models.Location{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					TechnologyStack: []models.TechnologyStack{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					MinExperience: uint(3),
// 					MaxExperience: uint(10),
// 					Qualifications: []models.Qualification{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					Shift: []models.Shift{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					Jobtype: "permanent",
// 				}, nil
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "error in location id's",
// 			args: args{
// 				ctx: context.Background(),
// 				applicationData: []newModels.NewUserApplication{
// 					{
// 						Name: "afthab",
// 						Age:  "21",
// 						Jid:  1,
// 						Jobs: newModels.Requestfield{
// 							Jobname:         "senior web developer",
// 							NoticePeriod:    7,
// 							Location:        []uint{},
// 							TechnologyStack: []uint{1, 2},
// 							Experience:      uint(5),
// 							Qualifications:  []uint{1, 2},
// 							Shift:           []uint{1, 2},
// 							Jobtype:         "permanent",
// 						},
// 					},
// 				},
// 			},
// 			want: nil,
// 			mockCacheResponse: func() (string, error) {
// 				return "", redis.Nil
// 			},
// 			mockAddCacheResponse: func() error {
// 				return nil
// 			},
// 			mockRepoResponse: func() (models.Jobs, error) {
// 				return models.Jobs{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Jobname:         "senior web developer",
// 					MinNoticePeriod: uint(5),
// 					MaxNoticePeriod: uint(10),
// 					Location: []models.Location{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					TechnologyStack: []models.TechnologyStack{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					MinExperience: uint(3),
// 					MaxExperience: uint(10),
// 					Qualifications: []models.Qualification{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					Shift: []models.Shift{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					Jobtype: "permanent",
// 				}, nil
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "error in qualifications",
// 			args: args{
// 				ctx: context.Background(),
// 				applicationData: []newModels.NewUserApplication{
// 					{
// 						Name: "afthab",
// 						Age:  "21",
// 						Jid:  1,
// 						Jobs: newModels.Requestfield{
// 							Jobname:         "senior web developer",
// 							NoticePeriod:    7,
// 							Location:        []uint{1, 2},
// 							TechnologyStack: []uint{1, 2},
// 							Experience:      uint(5),
// 							Qualifications:  []uint{},
// 							Shift:           []uint{1, 2},
// 							Jobtype:         "permanent",
// 						},
// 					},
// 				},
// 			},
// 			want: nil,
// 			mockCacheResponse: func() (string, error) {
// 				return "", redis.Nil
// 			},
// 			mockAddCacheResponse: func() error {
// 				return nil
// 			},
// 			mockRepoResponse: func() (models.Jobs, error) {
// 				return models.Jobs{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Jobname:         "senior web developer",
// 					MinNoticePeriod: uint(5),
// 					MaxNoticePeriod: uint(10),
// 					Location: []models.Location{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					TechnologyStack: []models.TechnologyStack{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					MinExperience: uint(3),
// 					MaxExperience: uint(10),
// 					Qualifications: []models.Qualification{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					Shift: []models.Shift{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					Jobtype: "permanent",
// 				}, nil
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "error in tech stacks",
// 			args: args{
// 				ctx: context.Background(),
// 				applicationData: []newModels.NewUserApplication{
// 					{
// 						Name: "afthab",
// 						Age:  "21",
// 						Jid:  1,
// 						Jobs: newModels.Requestfield{
// 							Jobname:         "senior web developer",
// 							NoticePeriod:    7,
// 							Location:        []uint{1, 2},
// 							TechnologyStack: []uint{},
// 							Experience:      uint(5),
// 							Qualifications:  []uint{1, 2},
// 							Shift:           []uint{1, 2},
// 							Jobtype:         "permanent",
// 						},
// 					},
// 				},
// 			},
// 			want: nil,
// 			mockCacheResponse: func() (string, error) {
// 				return "", redis.Nil
// 			},
// 			mockAddCacheResponse: func() error {
// 				return nil
// 			},
// 			mockRepoResponse: func() (models.Jobs, error) {
// 				return models.Jobs{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Jobname:         "senior web developer",
// 					MinNoticePeriod: uint(5),
// 					MaxNoticePeriod: uint(10),
// 					Location: []models.Location{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					TechnologyStack: []models.TechnologyStack{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					MinExperience: uint(3),
// 					MaxExperience: uint(10),
// 					Qualifications: []models.Qualification{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					Shift: []models.Shift{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					Jobtype: "permanent",
// 				}, nil
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "error in shifts",
// 			args: args{
// 				ctx: context.Background(),
// 				applicationData: []newModels.NewUserApplication{
// 					{
// 						Name: "afthab",
// 						Age:  "21",
// 						Jid:  1,
// 						Jobs: newModels.Requestfield{
// 							Jobname:         "senior web developer",
// 							NoticePeriod:    7,
// 							Location:        []uint{1, 2},
// 							TechnologyStack: []uint{1, 2},
// 							Experience:      uint(5),
// 							Qualifications:  []uint{1, 2},
// 							Shift:           []uint{},
// 							Jobtype:         "permanent",
// 						},
// 					},
// 				},
// 			},
// 			want: nil,
// 			mockCacheResponse: func() (string, error) {
// 				return "", redis.Nil
// 			},
// 			mockAddCacheResponse: func() error {
// 				return nil
// 			},
// 			mockRepoResponse: func() (models.Jobs, error) {
// 				return models.Jobs{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Jobname:         "senior web developer",
// 					MinNoticePeriod: uint(5),
// 					MaxNoticePeriod: uint(10),
// 					Location: []models.Location{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					TechnologyStack: []models.TechnologyStack{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					MinExperience: uint(3),
// 					MaxExperience: uint(10),
// 					Qualifications: []models.Qualification{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					Shift: []models.Shift{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					Jobtype: "permanent",
// 				}, nil
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "success from database",
// 			args: args{
// 				ctx: context.Background(),
// 				applicationData: []newModels.NewUserApplication{
// 					{
// 						Name: "afthab",
// 						Age:  "21",
// 						Jid:  1,
// 						Jobs: newModels.Requestfield{
// 							Jobname:         "senior web developer",
// 							NoticePeriod:    7,
// 							Location:        []uint{1, 2},
// 							TechnologyStack: []uint{1, 2},
// 							Experience:      uint(5),
// 							Qualifications:  []uint{1, 2},
// 							Shift:           []uint{1, 2},
// 							Jobtype:         "permanent",
// 						},
// 					},
// 				},
// 			},
// 			want: []newModels.NewUserApplication{
// 				{
// 					Name: "afthab",
// 					Age:  "21",
// 					Jid:  1,
// 					Jobs: newModels.Requestfield{
// 						Jobname:         "senior web developer",
// 						NoticePeriod:    7,
// 						Location:        []uint{1, 2},
// 						TechnologyStack: []uint{1, 2},
// 						Experience:      uint(5),
// 						Qualifications:  []uint{1, 2},
// 						Shift:           []uint{1, 2},
// 						Jobtype:         "permanent",
// 					},
// 				},
// 			},
// 			mockCacheResponse: func() (string, error) {
// 				return "", redis.Nil
// 			},
// 			mockAddCacheResponse: func() error {
// 				return nil
// 			},
// 			mockRepoResponse: func() (models.Jobs, error) {
// 				return models.Jobs{
// 					Model: gorm.Model{
// 						ID: 1,
// 					},
// 					Jobname:         "senior web developer",
// 					MinNoticePeriod: uint(5),
// 					MaxNoticePeriod: uint(10),
// 					Location: []models.Location{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					TechnologyStack: []models.TechnologyStack{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					MinExperience: uint(3),
// 					MaxExperience: uint(10),
// 					Qualifications: []models.Qualification{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					Shift: []models.Shift{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 					},
// 					Jobtype: "permanent",
// 				}, nil
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "success from cache",
// 			args: args{
// 				ctx: context.Background(),
// 				applicationData: []newModels.NewUserApplication{
// 					{
// 						Name: "afthab",
// 						Age:  "21",
// 						Jid:  1,
// 						Jobs: newModels.Requestfield{
// 							Jobname:         "senior web developer",
// 							NoticePeriod:    7,
// 							Location:        []uint{1, 2},
// 							TechnologyStack: []uint{1, 2},
// 							Experience:      uint(5),
// 							Qualifications:  []uint{1, 2},
// 							Shift:           []uint{1, 2},
// 							Jobtype:         "permanent",
// 						},
// 					},
// 				},
// 			},
// 			want: nil,
// 			mockCacheResponse: func() (string, error) {
// 				return `some random data`, nil
// 			},
// 			mockAddCacheResponse: func() error {
// 				return nil
// 			},
// 			mockRepoResponse: func() (models.Jobs, error) {
// 				return models.Jobs{}, errors.New("test error")
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "success from cache",
// 			args: args{
// 				ctx: context.Background(),
// 				applicationData: []newModels.NewUserApplication{
// 					{
// 						Name: "afthab",
// 						Age:  "21",
// 						Jid:  1,
// 						Jobs: newModels.Requestfield{
// 							Jobname:         "senior web developer",
// 							NoticePeriod:    7,
// 							Location:        []uint{1, 2},
// 							TechnologyStack: []uint{1, 2},
// 							Experience:      uint(5),
// 							Qualifications:  []uint{1, 2},
// 							Shift:           []uint{1, 2},
// 							Jobtype:         "permanent",
// 						},
// 					},
// 				},
// 			},
// 			want: []newModels.NewUserApplication{
// 				{
// 					Name: "afthab",
// 					Age:  "21",
// 					Jid:  1,
// 					Jobs: newModels.Requestfield{
// 						Jobname:         "senior web developer",
// 						NoticePeriod:    7,
// 						Location:        []uint{1, 2},
// 						TechnologyStack: []uint{1, 2},
// 						Experience:      uint(5),
// 						Qualifications:  []uint{1, 2},
// 						Shift:           []uint{1, 2},
// 						Jobtype:         "permanent",
// 					},
// 				},
// 			},
// 			mockCacheResponse: func() (string, error) {
// 				return `{"ID":1,"CreatedAt":"2023-11-08T17:02:42.174821+05:30","UpdatedAt":"2023-11-08T17:02:42.174821+05:30","DeletedAt":null,"company":{"ID":1,"CreatedAt":"2023-11-08T16:59:43.329653+05:30","UpdatedAt":"2023-11-08T16:59:43.329653+05:30","DeletedAt":null,"name":"Tek Systems","location":"Bellandur, Bengaluru","field":"Information Technology"},"cid":1,"jobname":"senior web developer","min_notice_period":7,"max_notice_period":30,"location":[{"ID":1,"CreatedAt":"2023-11-08T10:00:00+05:30","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"place_name":"bengaluru"},{"ID":2,"CreatedAt":"2023-11-08T11:30:00+05:30","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"place_name":"mumbai"}],"technologyStack":[{"ID":1,"CreatedAt":"2023-11-08T10:00:00+05:30","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"stack_name":"golang"},{"ID":2,"CreatedAt":"2023-11-08T10:00:00+05:30","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"stack_name":"java"}],"description":" Web designers primarily focus on the visual and user experience aspects of web development. They create mockups, wireframes, and prototypes to communicate design concepts, working closely with web developers to implement designs and maintain a consistent user interface.","min_experience":1,"max_experience":5,"qualifications":[{"ID":1,"CreatedAt":"2023-11-08T10:00:00+05:30","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"qualification_required":"bachelors"},{"ID":2,"CreatedAt":"2023-11-08T10:00:00+05:30","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"qualification_required":"post graduate"}],"shifts":[{"ID":1,"CreatedAt":"2023-11-08T08:00:00+05:30","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"shift_type":"Morning Shift"},{"ID":2,"CreatedAt":"2023-11-08T15:00:00+05:30","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"shift_type":"Afternoon Shift"}],"jobtype":"permanent"}`, nil
// 			},
// 			mockAddCacheResponse: func() error {
// 				return nil
// 			},
// 			mockRepoResponse: func() (models.Jobs, error) {
// 				return models.Jobs{}, errors.New("test error")
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mc := gomock.NewController(t)
// 			mockCach := mockmodels.NewMockCaching(mc)
// 			mockCach.EXPECT().GetTheCacheData(gomock.Any(), gomock.Any()).Return(tt.mockCacheResponse()).AnyTimes()
// 			mockCach.EXPECT().AddToTheCache(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.mockAddCacheResponse()).AnyTimes()

// 			mockRepo := cachMock.NewMockUserRepo(mc)
// 			mockRepo.EXPECT().GetTheJobData(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()

// 			svc, err := NewService(mockRepo, &auth.Auth{}, mockCach)
// 			if err != nil {
// 				t.Error("could not initialize the mock repo layer")
// 				return
// 			}

// 			got, err := svc.ProccessApplication(tt.args.ctx, tt.args.applicationData)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Service.ProccessApplication() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Service.ProccessApplication() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
