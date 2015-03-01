package main

import (
	"net/http"
)

func middlewareJSON(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// do some stuff before
	rw.Header().Set("Content-Type", "application/json")
	next(rw, r)
	// do some stuff after
}
