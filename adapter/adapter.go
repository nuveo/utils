package adapter

import "gopkg.in/mgo.v2/bson"

// Driver DB interface
type Driver interface {
	Conn() error
	Find(collection string, query bson.M, params ...int) ([]map[string]interface{}, error)
	Insert(collection string, model interface{}) error
	Update(collection string, where bson.M, model interface{}) error
	UpdateAll(collection string, where bson.M, model interface{}) error
	Copy() Driver
	Close()
}
