package adapter

import "labix.org/v2/mgo/bson"

// DB Driver interface
type Driver interface {
	Conn() error
	Find(query bson.M, collection string, params ...int) ([]map[string]interface{}, error)
	Insert(collection string, model interface{}) error
	Update(collection string, where bson.M, model interface{}) error
	Copy() Driver
}
