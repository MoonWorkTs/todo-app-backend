package main

import (
	"os"
	todoapp "todo-app"
	"todo-app/pkg/handler"
	"todo-app/pkg/repository"
	"todo-app/pkg/service"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error initializing env variables: %s", err.Error())
	}

	db, err := repository.NewPostresDB(repository.Config{
		Host:       viper.GetString("db.host"),
		Port:       viper.GetString("db.port"),
		DBName:     viper.GetString("db.dbname"),
		SSLMode:    viper.GetString("db.sslmode"),
		SearchPath: viper.GetString("db.search_path"),
		Username:   os.Getenv("DB_USERNAME"),
		Password:   os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("error initializing db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todoapp.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error occuerd while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
