package main

import "github.com/garyburd/redigo/redis"
import "fmt"
import "time"
import "flag"

var redis_ip = flag.String("redis_ip", "000.000.000.000", "IP Address for redis")
var redis_port = flag.String("redis_port", "6379", "Port for redis")
var open_connections = flag.Int("connections", 10, "Open connections at a time")

var server = ""

func main() {
	flag.Parse()
	server = fmt.Sprint(*redis_ip) + ":" + fmt.Sprint(*redis_port)

	for i := 0; i < *open_connections; i++ {
		go dowork(i)
	}
	for {
		fmt.Println("sleeping")
		time.Sleep(time.Second)
	}

}

func dowork(i int) {
	c, err := redis.Dial("tcp", server)
	if err != nil {
		fmt.Println("Error:", err)
	}
	c.Do("del", "poop:"+fmt.Sprint(i))
	for {
		_, err := c.Do("rpush", "poop:"+fmt.Sprint(i), 1)
		if err != nil {
			fmt.Println("Rpush Error:", err)
			c, err = redis.Dial("tcp", server)
			if err != nil {
				fmt.Println("Error:", err)
			}
		}
		time.Sleep(time.Millisecond * 200)
	}
}
