package main

import (
	"encoding/json"
	"github.com/liveplant/liveplant-server/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	"github.com/liveplant/liveplant-server/Godeps/_workspace/src/github.com/carbocation/interpose"
	gorilla_mux "github.com/liveplant/liveplant-server/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/liveplant/liveplant-server/Godeps/_workspace/src/github.com/tylerb/graceful"
	"net/http"
	"os"
	"time"
)

type Application struct {
}

type CurrentAction struct {
	Action        string `json:"action"`
	UnixTimestamp int64  `json:"unixTimestamp"`
}

func GetCurrentAction(w http.ResponseWriter, r *http.Request) {
	action := &CurrentAction{
		Action:        "water",
		UnixTimestamp: int64(time.Now().Unix()),
	}
	json.NewEncoder(w).Encode(action)
}

func NewApplication() (*Application, error) {
	app := &Application{}
	return app, nil
}

func (app *Application) mux() *gorilla_mux.Router {
	router := gorilla_mux.NewRouter()

	router.HandleFunc("/current_action", GetCurrentAction).Methods("GET")

	return router
}

func main() {
	app, _ := NewApplication()
	middle := interpose.New()

	middle.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			next.ServeHTTP(w, req)
		})
	})

	middle.UseHandler(app.mux())

	ServerPort := "5000" // default port
	if len(os.Getenv("PORT")) > 0 {
		ServerPort = os.Getenv("PORT")
	}
	drainInterval, _ := time.ParseDuration("1s")
	srv := &graceful.Server{
		Timeout: drainInterval,
		Server: &http.Server{
			Addr:    ":" + ServerPort,
			Handler: middle,
		},
	}
	err := srv.ListenAndServe()
	if err != nil {
		logrus.Fatal(err.Error())
	}
}
