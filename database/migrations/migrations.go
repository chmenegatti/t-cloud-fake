package migrations

import (
	"github.com/chmenegatti/t-cloud-fake/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Alert{})
}
