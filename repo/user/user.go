package user

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const collection = "users"

type User struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Username string        `bson:"username" json:"username"`
	Password string        `bson:"password" json:"password"`
}

func NewUser() *User {
	return new(User)
}

func (self *User) Add(s *mgo.Session) (err error) {

	//	self.Username = username
	//	self.Password = password

	c := s.DB("mobile").C(collection)
	index := mgo.Index{
		Key:    []string{"username"},
		Unique: true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		return
	}

	err = c.Insert(self)
	return
}

func (self *User) Find(s *mgo.Session) (user *User, err error) {

	c := s.DB("mobile").C(collection)
	q := c.Find(bson.M{"username": self.Username})
	x := &User{}
	err = q.One(x)
	user = x
	return
}
