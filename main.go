package main

import (
	"encoding/json"
	log "github.com/liveplant/liveplant-server/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	"github.com/liveplant/liveplant-server/Godeps/_workspace/src/github.com/carbocation/interpose"
	gorilla_mux "github.com/liveplant/liveplant-server/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/liveplant/liveplant-server/Godeps/_workspace/src/github.com/tylerb/graceful"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

// Define string constants that correspond
// to each supported action here
const (
	ActionWater   string = "water"
	ActionNothing string = "nothing"
)

// Variables for keeping track of the current vote count
// for each action.
// These should probably be stored in redis at some point.
var VoteCountWater   int = 0
var VoteCountNothing int = 0

type Application struct {
}

type CurrentAction struct {
	Action        string `json:"action"`
	UnixTimestamp int64  `json:"unixTimestamp"`
}

func GetCurrentAction(w http.ResponseWriter, r *http.Request) {
	action := &CurrentAction{
		Action:        ActionWater,
		UnixTimestamp: int64(time.Now().Unix()),
	}
	json.NewEncoder(w).Encode(action)
}

func PrintHttpRequest(r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err == nil {
		log.Println("Request received: \n" + string(dump))
	} else {
		log.Error("Error reading request: " + err.Error())
	}
}

/*
Example cURL commands for invoking the API on the commandline:

curl -H "Content-Type: application/json" -X POST -d '{"someKey1":"someValue1","someKey2":"someValue2"}' http://127.0.0.1:5000/votes
curl -H "Content-Type: application/json" -X POST -d '{"action":"water"}' http://127.0.0.1:5000/votes
curl -H "Content-Type: application/json" -X POST -d '{"action":"nothing"}' http://127.0.0.1:5000/votes

curl http://127.0.0.1:5000/current_action
*/

type Vote struct {
	Action string `json:"action"`
}

func PostVotes(w http.ResponseWriter, r *http.Request) {

	log.Info("PostVotes called")

	// PrintHttpRequest(r)

	decoder := json.NewDecoder(r.Body)
	var vote Vote
	err := decoder.Decode(&vote)

	if err == nil {

		if vote.Action == ActionWater {
			VoteCountWater++
			log.Info("Voted for action \"" + ActionWater + "\" ", VoteCountWater)
		} else if vote.Action == ActionNothing {
			VoteCountNothing++
			log.Info("Voted for action \"" + ActionNothing + "\" ", VoteCountNothing)
		} else {
			log.Error("Encountered unhandled action \"" + vote.Action + "\"")
		}

	} else {
		log.Error("Error parsing vote body: " + err.Error())
	}

	// TODO - output a json response
	// { "message":string }
}

/*
Example response from GET /votes
A json dictionary that shows the current number
of votes that have been cast for each action.
{
  "actions": {
    "water": 1,
    "nothing": 0
  }
}
*/
type CurrentVoteCount struct {
	Actions map[string]int `json:"actions"`
}

func GetVotes(w http.ResponseWriter, r *http.Request) {

	log.Info("GetVotes called")

	currentVotes := &CurrentVoteCount{
		Actions: make(map[string]int),
	}

	currentVotes.Actions[ActionWater]   = VoteCountWater
	currentVotes.Actions[ActionNothing] = VoteCountNothing

	json.NewEncoder(w).Encode(currentVotes)
}

func NewApplication() (*Application, error) {
	app := &Application{}
	return app, nil
}

func (app *Application) mux() *gorilla_mux.Router {
	router := gorilla_mux.NewRouter()

	router.HandleFunc("/current_action", GetCurrentAction).Methods("GET")
	router.HandleFunc("/votes", PostVotes).Methods("POST")
	router.HandleFunc("/votes", GetVotes).Methods("GET")

	return router
}

func main() {
	log.Println("Launching liveplant server")

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

	log.Println("Running liveplant server on port " + ServerPort)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
