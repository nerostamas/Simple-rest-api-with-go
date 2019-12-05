package app

import (
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Repository struct {

}

const SERVER = "localhost:27017"

const DBNAME = "musicstore"

const DOCNAME = "albums"

func (r Repository) getAlbums() Albums  {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server: ", err)
	}

	defer session.Clone()
	c := session.DB(DBNAME).C(DOCNAME)
	results := Albums{}
	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results: ", err)
	}

	return results
}

func (r Repository) AddAlbum(album Album) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Clone()

	album.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(album)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func (r Repository) UpdateAlbum(album Album) bool  {
	session, err := mgo.Dial(SERVER)
	defer session.Clone()
	session.DB(DBNAME).C(DOCNAME).UpdateId(album.ID, album)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func (r Repository) DeleteAlbum(id string) string  {
	session, err := mgo.Dial(SERVER)
	defer session.Clone()

	if !bson.IsObjectIdHex(id) {
		return "NOT FOUND"
	}

	oid := bson.ObjectIdHex(id)

	if err = session.DB(DBNAME).C(DOCNAME).RemoveId(oid); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}

	return "OK"
}