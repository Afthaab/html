package database

import (
	"fmt"

	"github.com/afthaab/job-portal/config"
	"github.com/afthaab/job-portal/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabase(cfg config.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", cfg.Host, cfg.User, cfg.Password, cfg.DbName, cfg.Port, cfg.Sslmode, cfg.Timezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// AutoMigrate function will ONLY create tables, missing columns and missing indexes, and WON'T change existing column's type or delete unused columns
	err = db.Migrator().AutoMigrate(&models.User{}, &models.Company{}, &models.Jobs{}, &models.Location{}, &models.Qualification{}, &models.Shift{}, &models.TechnologyStack{})
	if err != nil {
		// If there is an error while migrating, log the error message and stop the program
		return nil, err
	}

	return db, nil
}
