package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rendau/push/internal/adapters/db/pg"
	"github.com/rendau/push/internal/adapters/httpapi"
	"github.com/rendau/push/internal/adapters/logger/zap"
	"github.com/rendau/push/internal/domain/core"
	"github.com/rendau/push/internal/interfaces"
	"github.com/spf13/viper"
)

func Execute() {
	var err error

	app := struct {
		lg      interfaces.Logger
		db      interfaces.Db
		core    *core.St
		httpApi *httpapi.St
	}{}

	loadConf()

	app.lg, err = zap.NewLogger(viper.GetString("log_level"), viper.GetBool("debug"), false)
	if err != nil {
		log.Fatal(err)
	}

	app.db, err = pg.New(app.lg, viper.GetString("pg_dsn"))
	if err != nil {
		app.lg.Fatal(err)
	}

	app.core = core.New(app.lg, app.db, viper.GetString("fcm_server_key"))

	app.httpApi = httpapi.New(app.lg, app.core, viper.GetString("usr_auth_url"), viper.GetString("http_listen"))

	app.lg.Infow("Starting", "http_listen", viper.GetString("http_listen"))

	app.httpApi.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	var exitCode int

	select {
	case <-stop:
	case <-app.httpApi.Wait():
		exitCode = 1
	}

	app.lg.Infow("Shutting down...")

	ctx, ctxCancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer ctxCancel()

	err = app.httpApi.Shutdown(ctx)
	if err != nil {
		app.lg.Errorw("Fail to shutdown httpapi-api", err)
		exitCode = 1
	}

	os.Exit(exitCode)
}

func loadConf() {
	viper.SetDefault("debug", "false")
	viper.SetDefault("http_listen", ":80")
	viper.SetDefault("log_level", "debug")

	confFilePath := os.Getenv("CONF_PATH")
	if confFilePath == "" {
		confFilePath = "conf.yml"
	}
	viper.SetConfigFile(confFilePath)
	_ = viper.ReadInConfig()

	viper.AutomaticEnv()
}
