package main

import (
	"fmt"
	"github.com/kozhamseitova/ustaz-hub-api/internal/app"
	"github.com/kozhamseitova/ustaz-hub-api/internal/config"
)

// @title           Ustaz Hub API
// @version         0.0.1
// @description     Api for simple blog for teachers.

// @contact.name   Aisha
// @contact.email  kozhamseitova91@gmail.com

// @host      localhost:8080
// @BasePath  /api/v1

// @securitydefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%#v", cfg))

	err = app.Run(cfg)
	if err != nil {
		panic(err)
	}
}
