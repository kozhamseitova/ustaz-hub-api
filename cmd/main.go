package cmd

import (
	"fmt"
	"github.com/kozhamseitova/ustaz-hub-api/internal/app"
	"github.com/kozhamseitova/ustaz-hub-api/internal/config"
)

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
