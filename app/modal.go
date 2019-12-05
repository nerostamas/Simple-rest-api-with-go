package app

import "gopkg.in/mgo.v2/bson"

type Album struct {
	ID	bson.ObjectId	`bson:"_id"`
	Title	string 		`json:"title"`
	Artist	string		`json:"artist"`
	Year	string		`json:"year"`
}

type Albums []Album