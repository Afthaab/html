package handler

import (
	"log"

	"github.com/afthaab/job-portal/internal/auth"
	"github.com/afthaab/job-portal/internal/middleware"
	"github.com/afthaab/job-portal/internal/service"
	"github.com/gin-gonic/gin"
)

func SetupApi(a auth.Authentication, svc service.UserService) *gin.Engine {
	r := gin.New()

	m, err := middleware.NewMiddleware(a)
	if err != nil {
		log.Panic("middlewares not setup")
	}

	h, err := NewHandler(svc)
	if err != nil {
		log.Panic("handlers not setup")
	}

	r.Use(m.Log(), gin.Recovery())

	r.POST("/api/register", h.SignUp)
	r.POST("/api/login", h.Signin)

	r.POST("/api/forget/password/getmail", h.SendOtp)
	r.POST("/api/forget/password/verify/otp", h.VerifyOtp)

	r.POST("/api/companies", m.Authenticate(h.AddCompany))
	r.GET("/api/companies", m.Authenticate(h.ViewAllCompanies))
	r.GET("/api/companies/:id", m.Authenticate(h.ViewCompany))

	r.POST("/api/companies/:id/jobs", m.Authenticate(h.AddJobs))
	r.GET("/api/companies/:id/jobs", m.Authenticate(h.ViewJob))
	r.GET("/api/jobs", m.Authenticate(h.ViewAllJobs))
	r.GET("/api/jobs/:id", m.Authenticate(h.ViewJobByID))

	r.POST("/api/process/application", m.Authenticate(h.ProcessApplication))

	return r

}
