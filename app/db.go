package app

import "gopkg.in/mgo.v2"

var DbName = "mobile"

func InitDB() {
	var err error
	mgoSession, err = mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}

	//TODO: ensure index here
}
