package dao

import (
	"github.com/roseboy/bank-server/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	MySQL *gorm.DB
)

// InitMySQL init mysql db
func InitMySQL() (err error) {
	dsn := config.Cfg.MySQL.DSN
	MySQL, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err == nil {
		pingDB(MySQL)
	}
	return err
}

func pingDB(db *gorm.DB) {
	ticker := time.NewTicker(time.Minute)
	go func(t *time.Ticker) {
		for {
			<-t.C
			d, _ := db.DB()
			d.Ping()
		}
	}(ticker)
}
