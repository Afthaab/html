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

// func Test_handler_AddJobs(t *testing.T) {
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
// 			name: "invalid job id",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Params = append(c.Params, gin.Param{Key: "id", Value: "abc"})

// 				c.Request = httpRequest

// 				mc := gomock.NewController(t)
// 				ms := mockmodels.NewMockUserService(mc)

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusBadRequest,
// 			expectedResponse:   `{"error":"Bad Request"}`,
// 		},
// 		{
// 			name: "error in validating the json",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{
// 					"jobName" : "senior web developer",
// 					"minNoticePeriod" : 7,
// 					"maxNoticePeriod" : 30,
// 					"location" : [
// 						1,
// 						2
// 						],
// 					"technologyStack":[
// 						1,
// 						2
// 					],
// 					"description":" Web designers primarily focus on the visual and user experience aspects of web development. They create mockups, wireframes, and prototypes to communicate design concepts, working closely with web developers to implement designs and maintain a consistent user interface.",
// 					"minExperience":1,
// 					"maxExperience":5,
// 					"qualifications":[
// 						1,
// 						2
// 					],
// 					"shifts":[
// 						1,
// 						2
// 					],
// 					"jobtype":"permanent
// 				}`))
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Params = append(c.Params, gin.Param{Key: "id", Value: "123"})
// 				c.Request = httpRequest

// 				mc := gomock.NewController(t)
// 				ms := mockmodels.NewMockUserService(mc)

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusBadRequest,
// 			expectedResponse:   `{"error":"please provide details"}`,
// 		},
// 		{
// 			name: "error in service layer",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{
// 					"jobName" : "senior web developer",
// 					"minNoticePeriod" : 7,
// 					"maxNoticePeriod" : 30,
// 					"location" : [
// 						1,
// 						2
// 						],
// 					"technologyStack":[
// 						1,
// 						2
// 					],
// 					"description":" Web designers primarily focus on the visual and user experience aspects of web development. They create mockups, wireframes, and prototypes to communicate design concepts, working closely with web developers to implement designs and maintain a consistent user interface.",
// 					"minExperience":1,
// 					"maxExperience":5,
// 					"qualifications":[
// 						1,
// 						2
// 					],
// 					"shifts":[
// 						1,
// 						2
// 					],
// 					"jobtype":"permanent"
// 				}`))
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Params = append(c.Params, gin.Param{Key: "id", Value: "123"})

// 				c.Request = httpRequest

// 				mc := gomock.NewController(t)
// 				ms := mockmodels.NewMockUserService(mc)

// 				ms.EXPECT().AddJobDetails(c.Request.Context(), gomock.Any(), gomock.Any()).Return(newModels.ResponseNewJobs{}, errors.New("test error from mock function"))

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusInternalServerError,
// 			expectedResponse:   `{"error":"test error from mock function"}`,
// 		},
// 		{
// 			name: "success from service layer",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{
// 					"jobName" : "senior web developer",
// 					"minNoticePeriod" : 7,
// 					"maxNoticePeriod" : 30,
// 					"location" : [
// 						1,
// 						2
// 						],
// 					"technologyStack":[
// 						1,
// 						2
// 					],
// 					"description":" Web designers primarily focus on the visual and user experience aspects of web development. They create mockups, wireframes, and prototypes to communicate design concepts, working closely with web developers to implement designs and maintain a consistent user interface.",
// 					"minExperience":1,
// 					"maxExperience":5,
// 					"qualifications":[
// 						1,
// 						2
// 					],
// 					"shifts":[
// 						1,
// 						2
// 					],
// 					"jobtype":"permanent"
// 				}`))
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Params = append(c.Params, gin.Param{Key: "id", Value: "123"})

// 				c.Request = httpRequest

// 				mc := gomock.NewController(t)
// 				ms := mockmodels.NewMockUserService(mc)

// 				ms.EXPECT().AddJobDetails(c.Request.Context(), gomock.Any(), gomock.Any()).Return(newModels.ResponseNewJobs{
// 					Jobid: 1,
// 				}, nil)

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusOK,
// 			expectedResponse:   `{"jobID":1}`,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gin.SetMode(gin.TestMode)
// 			c, rr, ms := tt.setup()
// 			h := handler{
// 				service: ms,
// 			}
// 			h.AddJobs(c)
// 			assert.Equal(t, tt.expectedStatusCode, rr.Code)
// 			assert.Equal(t, tt.expectedResponse, rr.Body.String())
// 		})
// 	}
// }
