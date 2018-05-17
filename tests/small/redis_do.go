package main

import (
	_ "fmt"
	"question-bank/utils"
)
import "github.com/garyburd/redigo/redis"
import "time"

func main() {
	redis_conn, err := redis.Dial("tcp", "192.168.100.14:6379")
	if err != nil {
		utils.Logger.Error("dail redis failed")
		return
	} else {
		utils.Logger.Debug("redis connect success")
	}

	lock, ok, err := utils.TryLockWithTimeout(redis_conn, "VoxLock:online", time.Second*3600*24*7)
	if err != nil {
		utils.Logger.Error("get online lock failed, key: %s, redis address: %s, %v", "VoxLock:online", "192.168.100.14:6379", err)

		return
	}

	if !ok {
		utils.Logger.Warn("online lock is held by some one else")
		return
	}

	utils.Logger.Info("get lock")

	defer lock.Unlock()

}
