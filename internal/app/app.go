package app

import (
	"context"
	"errors"
	"github.com/stickpro/go-shop/internal/config"
	"github.com/stickpro/go-shop/internal/router"
	"github.com/stickpro/go-shop/internal/server"
	"github.com/stickpro/go-shop/pkg/database/pgsql"
	"github.com/stickpro/go-shop/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)

	if err != nil {
		logger.Error(err)

		return
	}

	db, _ := pgsql.ConnectionDataBase(cfg.DB.Host, cfg.DB.Username, cfg.DB.Password, cfg.DB.DBName, cfg.DB.Port)

	handler := router.NewRouter()

	srv := server.NewServer(cfg, handler.Init(db))

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logger.Info("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}

}
