package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Model interface {
	Unique() bson.M
	Collection() string
}

type Indexed interface {
	Indexes() []mgo.Index
	Model
}

func All(m Model) *mgo.Query {
	return Where(m, nil)
}

func Find(m Model) *mgo.Query {
	return Where(m, m.Unique()).Limit(1)
}

func Where(m Model, q interface{}) *mgo.Query {
	return db.C(m.Collection()).Find(q)
}

func Update(m Model) (*mgo.ChangeInfo, error) {
	return Find(m).Apply(mgo.Change{
		ReturnNew: true,
		Update: bson.M{
			"$set": m,
		},
	}, m)
}

func Insert(m Model) error {
	return db.C(m.Collection()).Insert(m)
}

func Delete(m Model) error {
	return db.C(m.Collection()).Remove(m.Unique())
}

func ensureIndexes(m Indexed) {
	coll := db.C(m.Collection())
	for _, i := range m.Indexes() {
		coll.EnsureIndex(i)
	}
}
