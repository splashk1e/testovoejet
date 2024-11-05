package main

import (
	"flag"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	licensev1 "github.com/splashk1e/jet/gen"
	"github.com/splashk1e/jet/internal"
	"github.com/splashk1e/jet/internal/app"
	"github.com/splashk1e/jet/internal/config"
	"github.com/splashk1e/jet/internal/handlers"
	"github.com/splashk1e/jet/internal/services"
)

func main() {
	testlicense := &licensev1.License{
		Uid:                       "12345",
		CreatedAt:                 time.Now().Unix(),
		UpdatedAt:                 time.Now().Unix(),
		CheckDate:                 time.Now().Unix(),
		RecheckDate:               time.Now().Unix(),
		Worktime:                  3600,
		Modules:                   []licensev1.Module{licensev1.Module_B, licensev1.Module_C},
		Version:                   "1.0.0",
		ReadOnly:                  false,
		RecheckNeeded:             true,
		WarningNotice:             []*licensev1.WarningNotice{{Notice: "Warning 1"}},
		CriticalNotice:            []*licensev1.CriticalNotice{{Notice: "Critical 1"}},
		Problems:                  []*licensev1.Problem{{Error: "Error 1", Date: time.Now().Unix()}},
		MaxBasicConn:              100,
		MaxComplianceConn:         50,
		ConnSoftLimit:             true,
		ConnLimitExcess:           []int64{1, 2, 3},
		ComplianceConnLimitExcess: []int64{4, 5},
		PublicKey:                 "your_public_key",
	}

	logrus.SetFormatter(new(logrus.JSONFormatter))
	mode := flag.String("mode", "dev", "Mode of the application: dev or deploy")
	flag.Parse()

	var cfg config.Config
	if err := cfg.InitConfig(); err != nil {
		logrus.Fatalf("can not init config with error:%s", err.Error())
	}
	switch *mode {
	case "deploy":
		cfg.FilePath = cfg.FilePath
	case "dev":
		cfg.FilePath = cfg.FilePathDev
	default:
		logrus.Fatal("wrong flag type, use 'dev' or 'deploy'")
	}

	workerService := services.NewWorkerService(cfg)
	logrus.Infof("worker service created")

	worker := internal.NewWorker(workerService)
	logrus.Infof("worker created")

	serverService := services.NewServerService(cfg)
	logrus.Infof("server service created")

	handler := handlers.NewHandler(serverService)
	logrus.Infof("handler created")

	server := internal.Server{}
	app := app.App{
		Config:  cfg,
		Server:  &server,
		Worker:  worker,
		Handler: handler,
	}
	logrus.Infof("app created")
	if _, err := os.Stat(cfg.FilePath); os.IsNotExist(err) {
		if *mode == "deploy" {
			logrus.Fatal("license file is not extists")
		}
		if *mode == "dev" {
			if _, err := os.Create("license.txt"); err != nil {
				logrus.Errorf("can not create file with error:%s", err.Error())
			}
			if err := workerService.FileWrite(testlicense); err != nil {
				logrus.Errorf("can not write file with error:%s", err.Error())
			}
		}
	}
	app.Run()
}
