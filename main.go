package main

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2"
)

func main() {
	http.HandleFunc("/", catchAll)
	http.HandleFunc("/medlemmar/", CORS(medlemmar))
	http.ListenAndServe(":1987", nil)
}

func catchAll(w http.ResponseWriter, r *http.Request) {
	handleError(w, 404)
}

func medlemmar(w http.ResponseWriter, r *http.Request) {
	segs := GetSegs(r)
	size := len(segs)

	switch {
	case size == 2:
		medlemWithID(w, r)
		return
	case size > 2:
		handleError(w, 404)
		return
	}

	switch r.Method {
	case "GET":
		medlemGet(w, r)
		return
	case "POST":
		medlemPost(w, r)
		return
	}

	handleError(w, 404)
}

func medlemWithID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		medlemShow(w, r)
		return
	case "DELETE":
		medlemDelete(w, r)
		return
	case "PUT":
		medlemUpdate(w, r)
		return
	}

	handleError(w, 404)
}

func medlemGet(w http.ResponseWriter, r *http.Request) {
	medlemmar := []*Medlem{}

	All(&Medlem{}).All(&medlemmar)

	JSON(w, medlemmar, 200)
}

func medlemPost(w http.ResponseWriter, r *http.Request) {
	medlem := &Medlem{}

	json.NewDecoder(r.Body).Decode(&medlem)

	medlem.BeforeCreate()

	if err := Insert(medlem); err != nil {
		handleError(w, 400)
		return
	}

	JSON(w, medlem, 201)
}

func medlemUpdate(w http.ResponseWriter, r *http.Request) {
	medlem := &Medlem{Slug: GetSeg(r, 2)}

	json.NewDecoder(r.Body).Decode(&medlem)

	medlem.BeforeUpdate()

	_, err := Update(medlem)

	switch {
	case err == mgo.ErrNotFound:
		handleError(w, 404)
		return
	case err != nil:
		handleError(w, 400)
		return
	}

	JSON(w, medlem, 200)
}

func medlemShow(w http.ResponseWriter, r *http.Request) {
	medlem := &Medlem{Slug: GetSeg(r, 2)}

	err := Find(medlem).One(&medlem)

	if err == mgo.ErrNotFound {
		handleError(w, 404)
		return
	}

	JSON(w, medlem, 200)
}

func medlemDelete(w http.ResponseWriter, r *http.Request) {
	err := Delete(&Medlem{
		Slug: GetSeg(r, 2),
	})

	if err == mgo.ErrNotFound {
		handleError(w, 404)
	}
}

func handleError(w http.ResponseWriter, code int) {
	JSON(w, Error{codes[code], code}, code)
}
