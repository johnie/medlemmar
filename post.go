package medlemmar

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Medlem struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Slug      string        `bson:",omitempty" json:"slug"`
	Firstname string        `bson:",omitempty" json:"firstname"`
	Lastname  string        `bson:",omitempty" json:"lastname"`
	Email     string        `bson:",omitempty" json:"email"`
	Personnr  string        `bson:",omitempty" json:"personnummer"`
	Updated   time.Time     `bson:",omitempty" json:"updated"`
	Created   time.Time     `bson:",omitempty" json:"created"`
}

func init() {
	ensureIndexes(&Medlem{})
}

func (p *Medlem) Collection() string {
	return "medlemmar"
}

func (p *Medlem) Unique() bson.M {
	return bson.M{"slug": p.Slug}
}

func (p *Medlem) Indexes() []mgo.Index {
	slug := mgo.Index{
		Unique:   true,
		DropDups: true,
		Key: []string{
			"slug",
		},
	}
	return []mgo.Index{
		slug,
	}
}

func (p *Medlem) BeforeCreate() {
	p.ID = bson.NewObjectId()
	p.Slug = genUniqSlug(8)
	p.Created = time.Now()
	p.Updated = p.Created
}

func (p *Medlem) BeforeUpdate() {
	p.Updated = time.Now()
}

func genUniqSlug(n int) string {
	var medlem *Medlem

	slug := RandSeq(n)

	Where(medlem, slug).One(&medlem)

	if medlem != nil {
		return genUniqSlug(n)
	}

	return slug
}
