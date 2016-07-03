package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// Author: HackerZ
// Time  : 2016/7/03 11:45

/*
 * Testing Redigo.
 */

func main() {
	fmt.Println("=== Redis && Go Start ===")

	// Connect Redis on Redigo
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println("Redis Conn Error --> ", err.Error())
	}
	defer c.Close()
	fmt.Println()

	// Get String Value
	GetStringValue(c, "hello")
	fmt.Println()

	// Set String Value
	SetStringValue(c, "name", "HackerZ")
	fmt.Println()

	fmt.Println("=== Redis && Go End ===")
}

// GetStringValue
func GetStringValue(c redis.Conn, key string) {
	// GET --> Get the Value of the Key.
	value, err := redis.String(c.Do("GET", key))
	if err != nil {
		fmt.Println("GET KEY Error --> ", err.Error())
	}

	fmt.Println("GET \"", key, "\" value -->", value)
}

// SetStringValue
func SetStringValue(c redis.Conn, key, value string) {
	// SETNX --> Set the Value of a Key, only if the Key does not exist.
	n, err := c.Do("SETNX", key, value)
	if err != nil {
		fmt.Println("SETNX Error --> ", err.Error())
	}

	// EXPIRE --> Set a Key's time to live in second
	if n == int64(1) {
		n, _ = c.Do("EXPIRE", key, 10)
		// if EXPIRE SET Success
		if n == int64(1) {
			fmt.Println("EXPIRE Set Success!")
		} else {
			fmt.Println("EXPIRE Set False!")
		}
	} else {
		fmt.Println("The Key", key, "has been existed!")
	}
}
