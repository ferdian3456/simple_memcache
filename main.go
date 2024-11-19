package main

import (
	"github.com/bradfitz/gomemcache/memcache"
	"log"
)

func main() {
	client := memcache.New(":11211")

	_, err := client.Get("name")
	if err != nil && err != memcache.ErrCacheMiss {
		log.Println(err)
	}

	key := "name"
	value := "andhika"

	err = setKey(client, key, value)
	if err != nil {
		log.Println(err)
	}

	var item memcache.Item

	item, err = getKey(client, key)
	if err != nil {
		log.Println(err)
	}

	log.Println("Key:", key)
	log.Printf("Value: %s", item.Value)

	err = deleteKey(client, key)
	if err != nil {
		log.Println(err)
	}

	log.Printf("Success to delete key: %s from memcache", key)
}

func setKey(client *memcache.Client, key string, value string) error {
	err := client.Set(&memcache.Item{
		Key:   key,
		Value: []byte(value),
	})

	return err
}

func getKey(client *memcache.Client, key string) (memcache.Item, error) {
	item, err := client.Get(key)
	if err != nil {
		return memcache.Item{}, err
	}
	return *item, nil
}

func deleteKey(client *memcache.Client, key string) error {
	err := client.Delete(key)
	if err != nil {
		return err
	}

	return nil
}
