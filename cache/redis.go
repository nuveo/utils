package redis

import (
	"fmt"
	"log"

	gredis "github.com/garyburd/redigo/redis"
)

type Redis struct {
	uri  string
	Conn gredis.Conn
}

type Connection struct {
	Network string
	Address string
	Db      string
}

func (c *Connection) Dial() (*Redis, error) {
	conn, err := gredis.Dial(c.Network, c.Address)
	if err != nil {
		return &Redis{}, err
	}

	if c.Db != "" {
		_, err = conn.Do("SELECT", c.Db)
		if err != nil {
			return &Redis{}, err
		}
	}

	uri := fmt.Sprintf("%s-%s", c.Network, c.Address)
	return &Redis{uri: uri, Conn: conn}, err
}

func (r *Redis) Set(key string, value interface{}) error {
	_, err := r.Conn.Do("SET", key, value)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (r *Redis) Get(key string) (string, error) {
	reply, err := r.Conn.Do("GET", key)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	result, err := gredis.String(reply, err)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return result, err
}

func (r *Redis) Exists(key string) (bool, error) {
	reply, err := r.Conn.Do("EXISTS", key)
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	result, err := gredis.Bool(reply, err)
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	return result, err
}

func (r *Redis) HSet(hashname, key, value string) error {
	_, err := r.Conn.Do("HSET", hashname, key, value)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}

func (r *Redis) HGet(hashname, key string) (string, error) {
	reply, err := r.Conn.Do("HGET", hashname, key)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	result, err := gredis.String(reply, err)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return result, err
}

func (r *Redis) HExists(hashname, key string) (bool, error) {
	reply, err := r.Conn.Do("HEXISTS", hashname, key)
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	result, err := gredis.Bool(reply, err)
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	return result, err
}
