package main

import (
	"github.com/codegangsta/negroni"
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func newPool() *redis.Pool {
	redisURL := ":6379"
	if len(os.Getenv("REDISTOGO_URL")) > 0 {
		redisURL = os.Getenv("REDISTOGO_URL")
	}
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisURL)
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

var pool = newPool()

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/plants", PlantIndex).
		Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
	})

	n := negroni.Classic()
	n.Use(c)
	n.Use(negroni.HandlerFunc(middlewareJSON))
	n.UseHandler(router)

	ServerPort := "9001"
	if len(os.Getenv("PORT")) > 0 {
		ServerPort = os.Getenv("PORT")
	}

	log.Panic(http.ListenAndServe("0.0.0.0:"+ServerPort, n))
}
