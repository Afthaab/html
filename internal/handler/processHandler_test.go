package handler

// import (
// 	"context"
// 	"errors"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/afthaab/job-portal/internal/auth"
// 	"github.com/afthaab/job-portal/internal/middleware"
// 	newModels "github.com/afthaab/job-portal/internal/models/requestModels"
// 	"github.com/afthaab/job-portal/internal/service"
// 	"github.com/afthaab/job-portal/internal/service/mockmodels"
// 	"github.com/gin-gonic/gin"
// 	"github.com/go-playground/assert/v2"
// 	"github.com/golang-jwt/jwt/v5"
// 	"go.uber.org/mock/gomock"
// )

// func Test_handler_ProcessApplication(t *testing.T) {
// 	tests := []struct {
// 		name               string
// 		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.UserService)
// 		expectedStatusCode int
// 		expectedResponse   string
// 	}{
// 		{
// 			name: "missing trace id",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
// 				c.Request = httpRequest

// 				return c, rr, nil
// 			},
// 			expectedStatusCode: http.StatusInternalServerError,
// 			expectedResponse:   `{"error":"Internal Server Error"}`,
// 		},
// 		{
// 			name: "missing jwt claims",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest

// 				return c, rr, nil
// 			},
// 			expectedStatusCode: http.StatusUnauthorized,
// 			expectedResponse:   `{"error":"Unauthorized"}`,
// 		},
// 		{
// 			name: "error in validating the json",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`[
// 					{
// 						"name": "afthab",
// 						"age": "22",
// 						"jid": 1,
// 						"job_application": {
// 							"jobName": "senior web developer",
// 							"noticePeriod": 10,
// 							"location": [
// 								1

// 							],
// 							"technologyStack": [
// 								1
// 							],
// 							"experience": 5,
// 							"qualifications": [
// 								1
// 							],
// 							"shifts": [
// 								1
// 							],
// 							"jobtype": "permanent"
// 						}
// 					},
// 					{
// 						"name": "jeevan",
// 						"age": "22",
// 						"jid": 1,
// 						"job_application": {
// 							"jobName": "senior web developer",
// 							"noticePeriod": 10,
// 							"location": [
// 								1,
// 								2

// 							],
// 							"technologyStack": [
// 								1,
// 								2
// 							],
// 							"experience": 5,
// 							"qualifications": [
// 								1,
// 								2
// 							],
// 							"shifts": [
// 								1,
// 								2
// 							],
// 							"jobtype": "permanent
// 						}
// 					}
// 				]`))
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest

// 				mc := gomock.NewController(t)
// 				ms := mockmodels.NewMockUserService(mc)

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusBadRequest,
// 			expectedResponse:   `{"error":"please provide valid details"}`,
// 		},
// 		{
// 			name: "error from service layer",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`[
// 					{
// 						"name": "afthab",
// 						"age": "22",
// 						"jid": 1,
// 						"job_application": {
// 							"jobName": "senior web developer",
// 							"noticePeriod": 10,
// 							"location": [
// 								1
// 							],
// 							"technologyStack": [
// 								1
// 							],
// 							"experience": 5,
// 							"qualifications": [
// 								1
// 							],
// 							"shifts": [
// 								1
// 							],
// 							"jobtype": "permanent"
// 						}
// 					},
// 					{
// 						"name": "jeevan",
// 						"age": "22",
// 						"jid": 1,
// 						"job_application": {
// 							"jobName": "senior web developer",
// 							"noticePeriod": 10,
// 							"location": [
// 								1,
// 								2

// 							],
// 							"technologyStack": [
// 								1,
// 								2
// 							],
// 							"experience": 5,
// 							"qualifications": [
// 								1,
// 								2
// 							],
// 							"shifts": [
// 								1,
// 								2
// 							],
// 							"jobtype": "permanent"
// 						}
// 					}
// 				]`))
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest

// 				mc := gomock.NewController(t)
// 				ms := mockmodels.NewMockUserService(mc)

// 				ms.EXPECT().ProccessApplication(c.Request.Context(), gomock.Any()).Return([]newModels.NewUserApplication{}, errors.New("test service error")).AnyTimes()

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusBadRequest,
// 			expectedResponse:   `{"error":"test service error"}`,
// 		},
// 		{
// 			name: "success case",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`[
// 					{
// 						"name": "afthab",
// 						"age": "22",
// 						"jid": 1,
// 						"job_application": {
// 							"jobName": "senior web developer",
// 							"noticePeriod": 10,
// 							"location": [
// 								1

// 							],
// 							"technologyStack": [
// 								1
// 							],
// 							"experience": 5,
// 							"qualifications": [
// 								1
// 							],
// 							"shifts": [
// 								1
// 							],
// 							"jobtype": "permanent"
// 						}
// 					},
// 					{
// 						"name": "jeevan",
// 						"age": "22",
// 						"jid": 1,
// 						"job_application": {
// 							"jobName": "senior web developer",
// 							"noticePeriod": 10,
// 							"location": [
// 								1,
// 								2

// 							],
// 							"technologyStack": [
// 								1,
// 								2
// 							],
// 							"experience": 5,
// 							"qualifications": [
// 								1,
// 								2
// 							],
// 							"shifts": [
// 								1,
// 								2
// 							],
// 							"jobtype": "permanent"
// 						}
// 					}
// 				]`))
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest

// 				mc := gomock.NewController(t)
// 				ms := mockmodels.NewMockUserService(mc)

// 				ms.EXPECT().ProccessApplication(c.Request.Context(), gomock.Any()).Return([]newModels.NewUserApplication{
// 					{
// 						Name: "Afthab",
// 						Age:  "22",
// 					},
// 					{
// 						Name: "purvi",
// 						Age:  "23",
// 					},
// 					{
// 						Name: "jeevan",
// 						Age:  "22",
// 					},
// 				}, nil).AnyTimes()

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusOK,
// 			expectedResponse:   `[{"name":"Afthab","age":"22","jid":0,"job_application":{"jobName":"","noticePeriod":0,"location":null,"technologyStack":null,"experience":0,"qualifications":null,"shifts":null,"jobtype":""}},{"name":"purvi","age":"23","jid":0,"job_application":{"jobName":"","noticePeriod":0,"location":null,"technologyStack":null,"experience":0,"qualifications":null,"shifts":null,"jobtype":""}},{"name":"jeevan","age":"22","jid":0,"job_application":{"jobName":"","noticePeriod":0,"location":null,"technologyStack":null,"experience":0,"qualifications":null,"shifts":null,"jobtype":""}}]`,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gin.SetMode(gin.TestMode)
// 			c, rr, ms := tt.setup()
// 			h := handler{
// 				service: ms,
// 			}
// 			h.ProcessApplication(c)
// 			assert.Equal(t, tt.expectedStatusCode, rr.Code)
// 			assert.Equal(t, tt.expectedResponse, rr.Body.String())
// 		})
// 	}
// }
