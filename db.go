package main

import (
	"gopkg.in/mgo.v2"
)

const (
	mongoURI = "localhost:1987"
	dbname   = "medlemmar"
)

var db *mgo.Database

func init() {
	ds, err := mgo.Dial(mongoURI)

	if err != nil {
		panic(err)
	}

	db = ds.DB(dbname)
}
