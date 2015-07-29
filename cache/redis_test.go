package redis

import "testing"

func SetupRedis() *Redis {
	conn := Connection{"tcp", ":6379", "6"}
	rdb, _ := conn.Dial()

	return rdb
}

func TearDown(rdb *Redis) {
	defer rdb.Conn.Close()

	rdb.Conn.Do("FLUSHDB")
}

func TestSet(t *testing.T) {
	rds := SetupRedis()
	key := "key"
	val := "val"

	err := rds.Set(key, val)

	if err != nil {
		t.Fail()
	}
	// getting
	resp, err := rds.Exists(key)
	if resp != true {
		t.Fail()
	}

	TearDown(rds)
}

func TestGet(t *testing.T) {
	rds := SetupRedis()
	key := "key"
	val := "val"

	err := rds.Set(key, val)

	if err != nil {
		t.Fail()
	}
	// getting
	resp, err := rds.Get(key)
	if resp != val {
		t.Fail()
	}

	TearDown(rds)
}

func TestHSet(t *testing.T) {
	rds := SetupRedis()
	hash := "hashname"
	key := "key"
	val := "val"

	err := rds.HSet(hash, key, val)

	if err != nil {
		t.Fail()
	}
	// getting
	resp, err := rds.HExists(hash, key)
	if resp != true {
		t.Fail()
	}

	TearDown(rds)
}

func TestHGet(t *testing.T) {
	rds := SetupRedis()
	hash := "hashname"
	key := "key"
	val := "val"

	err := rds.HSet(hash, key, val)

	if err != nil {
		t.Fail()
	}
	// getting
	resp, err := rds.HGet(hash, key)
	if resp != val {
		t.Fail()
	}

	TearDown(rds)
}
