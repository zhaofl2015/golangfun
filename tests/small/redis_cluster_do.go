package main

import (
	"fmt"

	"time"

	_ "encoding/json"

	"github.com/chasex/redis-go-cluster"
)

var (
	cluster, _ = GetRedisClusterConnection()
)

func GetRedisClusterConnection() (*redis.Cluster, error) {
	cluster, err := redis.NewCluster(
		&redis.Options{
			StartNodes:   []string{"192.168.100.52:7000", "192.168.100.52:7001"},
			ConnTimeout:  50 * time.Millisecond,
			ReadTimeout:  50 * time.Millisecond,
			WriteTimeout: 50 * time.Millisecond,
			KeepAlive:    16,
			AliveTime:    60 * time.Second,
		})
	return cluster, err
}

func main() {
	//	data, err := cluster.Do("set", "newha", "wol")
	//	fmt.Println(data, err)

	data, err := cluster.Do("sadd", "smset", "a")
	fmt.Println(data, err)

	data, err = cluster.Do("sadd", "smset", "b")
	fmt.Println(data, err)

	data, err = cluster.Do("sadd", "smset", "aaaa")
	fmt.Println(data, err)

	data, err = redis.Values(cluster.Do("smembers", "smset2"))
	//	for _, item := range data {
	//		fmt.Println(item)
	//	}
	fmt.Println(data)
	for _, item := range data.([]interface{}) {
		fmt.Println(string(item.([]byte)))
	}
	//	fmt.Println(data.([]interface{}))
	//	var ids []string
	//	fmt.Println(json.Unmarshal(data, ids), err)
}
