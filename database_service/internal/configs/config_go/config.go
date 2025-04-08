package config_go

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Configure struct {
	Env         string        `yaml:"env" env-default:"local"`
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	GRPC        GRPCConfig    `yaml:"grpc" env-required:"true"`
	TokenTTL    time.Duration `yaml:"token_ttl"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port" env-required:"50052"`
	Timeout time.Duration `yaml:"timeout" env-required:"5s"`
}

func MustLoad() *Configure {
	path := SearchPath()
	if path == "" {
		panic("нужен путь")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("файл конфига не найден")
	}
	var cfg Configure

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("не удалось считать конфиг")
	}

	return &cfg
}

func SearchPath() string {
	p, _ := os.Getwd()
	fmt.Println(p)
	var res string
	flag.StringVar(&res, "config_go", "", "path to config_go")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}
