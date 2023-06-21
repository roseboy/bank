package dao

import (
	"gorm.io/gorm"
	"time"
)

// App app
type App struct {
	AppId     uint64         `gorm:"primaryKey,column:app_id"`
	SecretId  string         `gorm:"column:secret_id"`
	SecretKey string         `gorm:"column:secret_key"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index,column:deleted_at"`
}

// TableName app
func (a *App) TableName() string {
	return "app"
}

// GetAuthInfo get
func GetAuthInfo(ak string) (uint64, string) {
	app := App{}
	err := MySQL.Find(&app, "secret_id = ?", ak).Error
	if err != nil {
		return 0, ""
	}
	return app.AppId, app.SecretKey
}
