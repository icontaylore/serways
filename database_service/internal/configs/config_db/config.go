package config_db

import (
	"github.com/joho/godotenv"
	"os"
)

type ConfigDB struct {
	DatabaseName string
	DatabaseHost string
	DatabasePas  string
	DatabasePort string
	DatabaseUser string
}

func LoadConfigDB() (*ConfigDB, error) {
	err := godotenv.Load("./configs/.env")
	if err != nil {
		return nil, err
	}

	return &ConfigDB{
		DatabaseName: os.Getenv("DB_NAME"),
		DatabaseHost: os.Getenv("DB_HOST"),
		DatabasePas:  os.Getenv("DB_PAS"),
		DatabasePort: os.Getenv("PORT_ONE"),
		DatabaseUser: os.Getenv("DB_USER"),
	}, nil
}

//DB_USER=admin
//DB_NAME=db_service
//DB_PAS=1
//DB_HOST=localhost
//PORT_ONE=5430
//PORT_TWO=5432
