package main

import "lru/cache"

func main() {
	c := cache.New(3)

	c.Set("1", "asd")
	c.Set("2", "asd")
	c.Set("1", "asd")
	c.Set("2", "asd")
	c.Set("3", "dsa")
	c.Set("1", "asd")
	c.Set("5", "sd1")
}
