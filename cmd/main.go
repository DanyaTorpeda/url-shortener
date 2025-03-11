package main

import (
	"os"
	shortener "url-shortener"
	"url-shortener/pkg/handler"
	"url-shortener/pkg/repository"
	"url-shortener/pkg/service"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Errorf("error occured loading env variables: %s", err.Error())
	}

	if err := initConfig(); err != nil {
		logrus.Errorf("error occured initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: os.Getenv("PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Errorf("error occured connecting to db: %s", err.Error())
	}
	defer db.Close()

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(shortener.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		logrus.Errorf("error ocurred running server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
