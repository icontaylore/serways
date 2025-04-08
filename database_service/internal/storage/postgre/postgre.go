package postgre

import (
	dbConf "db_service/internal/configs/config_db"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	DB *gorm.DB
}

func NewStorage(cfg *dbConf.ConfigDB) (*Storage, error) {
	connStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.DatabaseHost, cfg.DatabaseUser, cfg.DatabaseName, cfg.DatabasePas,
	)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Storage{DB: db}, nil
}
