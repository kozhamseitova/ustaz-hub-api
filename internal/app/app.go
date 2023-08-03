package app

import (
	"github.com/kozhamseitova/api-blog/pkg/client/postgres"
	"github.com/kozhamseitova/api-blog/pkg/httpserver"
	"github.com/kozhamseitova/api-blog/pkg/jwttoken"
	"github.com/kozhamseitova/ustaz-hub-api/internal/config"
	"github.com/kozhamseitova/ustaz-hub-api/internal/handler"
	"github.com/kozhamseitova/ustaz-hub-api/internal/repository"
	"github.com/kozhamseitova/ustaz-hub-api/internal/service"
	"log"
	"os"
	"os/signal"
)

func Run(cfg *config.Config) error {
	conn, err := postgres.New(
		postgres.WithHost(cfg.DB.Host),
		postgres.WithPort(cfg.DB.Port),
		postgres.WithDBName(cfg.DB.DBName),
		postgres.WithUsername(cfg.DB.Username),
		postgres.WithPassword(cfg.DB.Password),
	)

	if err != nil {
		//ggygvy
	}

	log.Println("db connection success")

	token := jwttoken.New(cfg.Token.SecretKey)
	repos := repository.NewPostgresRepository(conn.Pool)
	srvcs := service.NewService(repos, cfg, token)
	hndl := handler.NewHandler(srvcs)

	server := httpserver.New(
		hndl.InitRouter(),
		httpserver.WithPort(cfg.HTTP.Port),
		httpserver.WithReadTimeout(cfg.HTTP.ReadTimeout),
		httpserver.WithWriteTimeout(cfg.HTTP.WriteTimeout),
		httpserver.WithShutdownTimeout(cfg.HTTP.ShutdownTimeout),
	)

	server.Start()
	log.Println("server started")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case s := <-interrupt:
		log.Printf("signal received: %s", s.String())
	case err = <-server.Notify():
		log.Printf("server notify: %s", err.Error())
	}

	err = server.Shutdown()
	if err != nil {
		log.Printf("server shutdown err: %s", err)
	}

	return nil

}
