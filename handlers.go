package main

import (
	//"encoding/json"
	//"github.com/gorilla/mux"
	//"github.com/mholt/binding"
	"github.com/unrolled/render"
	//"log"
	"net/http"
)

func PlantIndex(w http.ResponseWriter, r *http.Request) {
	resp := render.New()
	resp.JSON(w, http.StatusOK, &Plant{
    Name: "big-john",
    DisplayName: "Big John",
  })
}
