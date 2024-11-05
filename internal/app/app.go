package app

import (
	"net/http"
	"net/http/pprof"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/splashk1e/jet/internal"
	"github.com/splashk1e/jet/internal/config"
	"github.com/splashk1e/jet/internal/handlers"
)

type App struct {
	Config  config.Config
	Server  *internal.Server
	Worker  *internal.Worker
	Handler *handlers.Handler
}

func (app *App) Run() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		app.Server.Run(app.Config.Port, app.Handler)
	}()
	logrus.Info("server started")
	ServePProf()
	wg.Add(1)
	go func() {
		defer wg.Done()
		app.Worker.Run()
	}()
	logrus.Info("worker started")
	wg.Wait()
}
func (app *App) Shutdown() {

}
func ServePProf() {
	srv := http.Server{
		Addr:         ":6060",
		Handler:      pprofHandler(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Minute,
		IdleTimeout:  60 * time.Second,
	}
	logrus.Infof("start pprof on port :6060")
	go func() { _ = srv.ListenAndServe() }()
}
func pprofHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.Handle("/debug/pprof/block", pprof.Handler("block"))
	mux.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	mux.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	mux.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	return mux
}
