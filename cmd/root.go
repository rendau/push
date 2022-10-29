package cmd

import (
	"os"
	"time"

	dopDbPg "github.com/rendau/dop/adapters/db/pg"
	dopLoggerZap "github.com/rendau/dop/adapters/logger/zap"
	dopServerHttps "github.com/rendau/dop/adapters/server/https"
	"github.com/rendau/dop/dopTools"
	"github.com/rendau/push/docs"
	"github.com/rendau/push/internal/adapters/prv"
	"github.com/rendau/push/internal/adapters/prv/fcm"
	"github.com/rendau/push/internal/adapters/repo/pg"
	"github.com/rendau/push/internal/adapters/server/rest"
	"github.com/rendau/push/internal/domain/core"
)

func Execute() {
	var err error

	app := struct {
		lg         *dopLoggerZap.St
		db         *dopDbPg.St
		repo       *pg.St
		prv        prv.Prv
		core       *core.St
		restApi    *rest.St
		restApiSrv *dopServerHttps.St
	}{}

	confLoad()

	app.lg = dopLoggerZap.New(conf.LogLevel, conf.Debug)

	app.db, err = dopDbPg.New(conf.Debug, app.lg, dopDbPg.OptionsSt{
		Dsn: conf.PgDsn,
	})
	if err != nil {
		app.lg.Fatal(err)
	}

	app.repo = pg.New(app.db, app.lg)

	app.prv, err = fcm.New(app.lg, conf.FcmCredsPath)

	err = app.prv.Send(&prv.SendReqSt{
		Tokens: []string{"asdfasdfads"},
		Title:  "Hello",
		Body:   "world!",
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
		},
		Badge:      2,
		AndroidTag: "",
	})
	if err != nil {
		app.lg.Fatal(err)
	}

	app.core = core.New(app.lg, app.repo, app.prv)

	docs.SwaggerInfo.Host = conf.SwagHost
	docs.SwaggerInfo.BasePath = conf.SwagBasePath
	docs.SwaggerInfo.Schemes = []string{conf.SwagSchema}
	docs.SwaggerInfo.Title = "Push service"

	// START

	app.lg.Infow("Starting")

	app.restApiSrv = dopServerHttps.Start(
		conf.HttpListen,
		rest.GetHandler(
			app.lg,
			app.core,
			conf.HttpCors,
		),
		app.lg,
	)

	var exitCode int

	select {
	case <-dopTools.StopSignal():
	case <-app.restApiSrv.Wait():
		exitCode = 1
	}

	// STOP

	app.lg.Infow("Shutting down...")

	if !app.restApiSrv.Shutdown(20 * time.Second) {
		exitCode = 1
	}

	app.lg.Infow("Exit")

	os.Exit(exitCode)
}
