package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
	"strconv"
)

var muxRouter = mux.NewRouter()

type Bicycle struct {
	Brand  string
	Color  string
	Serial string
}

func createBike(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	decoder := json.NewDecoder(r.Body)
	var bike Bicycle
	err := decoder.Decode(&bike)
	key := datastore.NewIncompleteKey(ctx, "bicycle", nil)
	k, err := datastore.Put(ctx, key, &bike)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%d", k.IntID())

}

func updateBike(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	vars := mux.Vars(r)
	key_str := vars["key"]
	intId, _ := strconv.ParseInt(key_str, 10, 64)
	key := datastore.NewKey(ctx, "bicycle", "", intId, nil)
	decoder := json.NewDecoder(r.Body)
	var bike Bicycle
	err := decoder.Decode(&bike)
	k, err := datastore.Put(ctx, key, &bike)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%d", k.IntID())

}

func deleteBike(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	vars := mux.Vars(r)
	key_str := vars["key"]
	intId, _ := strconv.ParseInt(key_str, 10, 64)
	key := datastore.NewKey(ctx, "bicycle", "", intId, nil)
	err := datastore.Delete(ctx, key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func retrieveBike(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	vars := mux.Vars(r)
	key_str := vars["key"]
	intId, _ := strconv.ParseInt(key_str, 10, 64)
	key := datastore.NewKey(ctx, "bicycle", "", intId, nil)
	var bike Bicycle
	err := datastore.Get(ctx, key, &bike)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bike)
}

func main() {
	http.Handle("/", muxRouter)
	muxRouter.HandleFunc("/merlin/{key}", retrieveBike).Methods("GET")
	muxRouter.HandleFunc("/merlin/new", createBike).Methods("POST")
	muxRouter.HandleFunc("/merlin/update/{key}", updateBike).Methods("POST")
	muxRouter.HandleFunc("/merlin/delete/{key}", deleteBike).Methods("GET")
	appengine.Main()
}
