package main

import (
	"encoding/json"
	"flag"
	log "github.com/liveplant/liveplant-server/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	"github.com/liveplant/liveplant-server/Godeps/_workspace/src/github.com/carbocation/interpose"
	gorilla_mux "github.com/liveplant/liveplant-server/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/liveplant/liveplant-server/Godeps/_workspace/src/github.com/tylerb/graceful"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
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
var VoteCountWater int = 0
var VoteCountNothing int = 0

type Application struct {
}

type CurrentAction struct {
	Action        string `json:"action"`
	UnixTimestamp int64  `json:"unixTimestamp"`
}

func GetCurrentAction(w http.ResponseWriter, r *http.Request) {
	action := &CurrentAction{
		Action:        GetWinningAction(),
		UnixTimestamp: int64(time.Now().Unix()),
	}
	json.NewEncoder(w).Encode(action)
}

func GetWinningAction() string {
	// Return the action that has the greatest number of votes

	var winningAction string

	if VoteCountWater > VoteCountNothing {
		winningAction = ActionWater
	} else {
		winningAction = ActionNothing
	}

	return winningAction
}

func DebugPrintHttpRequest(r *http.Request) {
	// If debug logger is enabled,
	// print out all the details of the supplied HTTP request.
	if log.GetLevel() == log.DebugLevel {
		dump, err := httputil.DumpRequest(r, true)
		if err == nil {
			log.Debug("Request received: \n" + string(dump))
		} else {
			log.Debug("Error reading request: " + err.Error())
		}
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

	log.Debug("PostVotes called")

	DebugPrintHttpRequest(r)

	decoder := json.NewDecoder(r.Body)
	var vote Vote
	err := decoder.Decode(&vote)

	if err == nil {

		if vote.Action == ActionWater {
			VoteCountWater++
			log.Debug("Voted for action \""+ActionWater+"\" ", VoteCountWater)
		} else if vote.Action == ActionNothing {
			VoteCountNothing++
			log.Debug("Voted for action \""+ActionNothing+"\" ", VoteCountNothing)
		} else {
			log.Debug("Encountered unhandled action \"" + vote.Action + "\"")
		}

	} else {
		log.Debug("Error parsing vote body: " + err.Error())
	}

	// TODO - output a json response
	// { "message":string }
}

type ActionInfo struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	VoteCount   int    `json:"voteCount"`
}

type CurrentVoteInfo struct {
	Actions []ActionInfo `json:"actions"`
}

/*
Example response from GET /votes
{
  "actions": [
    {
      "name": "nothing",
      "displayName": "Nothing",
      "voteCount": 123
    },
    {
      "name": "water",
      "displayName": "Water",
      "voteCount": 47
    }
  ]
}
*/
func GetVotes(w http.ResponseWriter, r *http.Request) {

	currentVoteInfo := &CurrentVoteInfo{
		Actions: []ActionInfo{
			ActionInfo{Name: ActionNothing, DisplayName: "Nothing", VoteCount: VoteCountNothing},
			ActionInfo{Name: ActionWater, DisplayName: "Water", VoteCount: VoteCountWater},
		},
	}

	json.NewEncoder(w).Encode(currentVoteInfo)
}
func NewApplication() (*Application, error) {
	app := &Application{}
	return app, nil
}

type preFlightHandler func(http.ResponseWriter, *http.Request)

func NewPreFlightHandler(methods ...string) preFlightHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ", "))
			w.Header().Set("Access-Control-Allow-Headers", "content-type")
		}
	}
}

func (app *Application) mux() *gorilla_mux.Router {
	router := gorilla_mux.NewRouter()

	router.HandleFunc("/current_action", GetCurrentAction).Methods("GET")
	router.HandleFunc("/votes", PostVotes).Methods("POST")
	router.HandleFunc("/votes", GetVotes).Methods("GET")
	router.HandleFunc("/votes", NewPreFlightHandler("GET", "POST")).Methods("OPTIONS")

	return router
}

func InitLogLevel() {
	// Check if --debug argument was supplied on the command line.
	// Check if LIVEPLANTDEBUG environment variable is present.
	// (Environment variable takes precedence over command line flag)
	// Enable or disable debug logger accordingly.

	// Declare and parse command line flag
	boolPtr := flag.Bool("debug", false, "Whether or not to enable debug logger.")
	flag.Parse()

	var debugLoggerEnabled bool = *boolPtr

	if len(os.Getenv("LIVEPLANTDEBUG")) > 0 {
		// Environment variable is present, so
		// debug logger should be enabled.
		// (overrides command line flag)
		debugLoggerEnabled = true
	}

	if debugLoggerEnabled {
		// Log everything
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug logging enabled")
	} else {
		// Only log fatal or panic events
		// (events where the application is terminated)
		log.SetLevel(log.FatalLevel)
	}
}

func main() {

	InitLogLevel()

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
