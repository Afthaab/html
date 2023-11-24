package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/afthaab/job-portal/config"
	"github.com/afthaab/job-portal/internal/auth"
	"github.com/afthaab/job-portal/internal/cache"
	database "github.com/afthaab/job-portal/internal/database/postgres"
	redisDB "github.com/afthaab/job-portal/internal/database/redis"
	"github.com/afthaab/job-portal/internal/handler"
	"github.com/afthaab/job-portal/internal/repository"
	"github.com/afthaab/job-portal/internal/service"
	"github.com/golang-jwt/jwt/v5"

	"github.com/rs/zerolog/log"
)

func main() {
	err := StartApp()
	if err != nil {
		log.Panic().Err(err).Send()
	}
	log.Info().Msg("Hello this is our app")

}

func StartApp() error {
	cfg := config.GetConfig()
	// =========================================================================
	// initializing the authentication support
	log.Info().Msg("main started : initializing the authentication support")

	//reading the private key file

	privatePem := []byte(cfg.AuthConfig.PrivateKey)

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePem)
	if err != nil {
		return fmt.Errorf("error in parsing auth private key : %w", err) // %w is used for error wraping
	}

	publicPem := []byte(cfg.AuthConfig.PublicKey)

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicPem)
	if err != nil {
		return fmt.Errorf("error in parsing auth public key : %w", err) // %w is used for error wraping
	}
	a, err := auth.NewAuth(privateKey, publicKey)
	if err != nil {
		return fmt.Errorf("error in constructing auth %w", err)
	}
	// =========================================================================
	// start the database

	log.Info().Msg("main started : initializing the data")

	db, err := database.ConnectToDatabase(cfg.DBConfig)
	if err != nil {
		return fmt.Errorf("error in opening the database connection : %w", err)
	}

	pg, err := db.DB()
	if err != nil {
		return fmt.Errorf("error in getting the database instance")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = pg.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("database is not connected: %w", err)
	}

	// redis database connection
	rdb := redisDB.ConnectToRedis(cfg.RedisConfig)

	_, err = rdb.Ping(ctx).Result()

	if err != nil {
		return fmt.Errorf("redis  is not connected: %w", err)
	}

	redisLayer := cache.NewRDBLayer(rdb)

	// =========================================================================
	// initialize the repository layer
	repo, err := repository.NewRepository(db)
	if err != nil {
		return err
	}

	svc, err := service.NewService(repo, a, redisLayer)
	if err != nil {
		return err
	}

	// initializing the http server
	api := http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.APPConfig.Host, cfg.APPConfig.Port),
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.APPConfig.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.IdleTimeout) * time.Second,
		Handler:      handler.SetupApi(a, svc),
	}

	// channel to store any errors while setting up the service
	serverErrors := make(chan error, 1)

	go func() {
		log.Info().Str("Port", api.Addr).Msg("main started : api is listening")
		serverErrors <- api.ListenAndServe()
	}()

	//shutdown channel intercepts ctrl+c signals
	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, os.Interrupt)

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error : %w", err)

	case sig := <-shutdown:
		log.Info().Msgf("main: Start shutdown %s", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := api.Shutdown(ctx)
		if err != nil {
			err := api.Close()
			return fmt.Errorf("could not stop server gracefully : %w", err)
		}
	}
	return nil
}
