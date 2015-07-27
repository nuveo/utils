package elastic

import (
	"reflect"
	"testing"
)

func TestSessionAndConnection(t *testing.T) {
	// Create ElasticSearch driver
	db := New("http://192.168.59.103:9200", "test")

	if db == nil {
		t.Fatalf("error creating db: ", db)
	}

	// Connect to the DB
	err := db.Conn()

	if err != nil {
		t.Fatal(err)
	}

	//Delete Index
	db.DeleteIndex()

}

func TestInsertAndQuery(t *testing.T) {
	// Create ElasticSearch driver
	db := New("http://192.168.59.103:9200", "test")
	db.Conn()

	type Teste struct {
		Name string `json:"name"`
	}

	test := &Teste{"Teste"}

	err := db.Insert("data", test)

	if err != nil {
		t.Fatal(err)
	}

	objects, err := db.Query("Teste", "data", &Teste{})
	if err != nil {
		t.Fatal(err)
	}

	if len(objects) != 1 {
		t.Fatalf("expected 1, got %v", len(objects))
	}

	y := reflect.ValueOf(objects[0]).Elem().FieldByName("Name")
	if y.Interface() != "Teste" {
		t.Fatalf("expected 'Teste', got %v", y.Interface())
	}

	//Delete Index
	db.DeleteIndex()
}
