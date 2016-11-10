package mongo

import (
	"testing"

	"gopkg.in/mgo.v2/mgo/bson"
)

func TestConnectionAndInsertCollection(t *testing.T) {
	db := New("127.0.0.1:27017", "test")

	if db == nil {
		t.Fatal("error creating db: ", db)
	}

	// Connect to the DB
	err := db.Conn()

	if err != nil {
		t.Fatal(err)
	}

	type Teste struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	test := &Teste{"Teste", 21}

	err = db.Insert("data", test)
	if err != nil {
		t.Fatal(err)
	}

	objects, err := db.Find(bson.M{"name": "Teste"}, "data")
	if err != nil {
		t.Fatal(err)
	}

	if len(objects) != 1 {
		t.Fatalf("expected 1, got %v", len(objects))
	}

	db.DropCollection("data")
	if err != nil {
		t.Fatal(err)
	}

}
