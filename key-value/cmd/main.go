package main

import (
	"fmt"
	"key-value/internals"
	"time"
)

func main() {
	cache := internals.NewCache()

	cache.Set("key1", "PavanIlla", 10)
	cache.Set("key2", "TejaIlla", 0)

	if value, found := cache.Get("key1"); found {
		fmt.Println("Found:", value)
	}

	cache.Delete("item2")

	time.Sleep(11 * time.Second)

	if _, found := cache.Get("key1"); !found {
		fmt.Println("Key1 got expired")
	}
}
