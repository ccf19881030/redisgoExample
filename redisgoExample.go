package main

import (
	"fmt"
	"time"

	"ybu.cn/iot/cache"
	"ybu.cn/iot/common"
)

func main() {
	fmt.Println("redisgo client demo")
	// redis的配置
	redisOpt := common.RedisConnOpt{
		true,
		"127.0.0.1",
		6379,
		"123456",
		3,
		240,
	}

	_redisCli := cache.NewRedis(redisOpt)
	// KEYS 示例
	keys, err := _redisCli.KEYS("0_last_gb212_2011:*")
	if err != nil {
		fmt.Println("KEYS failed, err: %v", err)
	}
	for index, val := range keys {
		fmt.Printf("第%d个值为：%s\n", index + 1, val)
	}
	// GET 示例
	key, err := _redisCli.GET("username")
	if err != nil {
		fmt.Println("GET failed, err: %v", err)
	}
	fmt.Println("key: ", key)
	// SET 示例
	//i1, err := _redisCli.SET("month", "12")
	//if err != nil {
	//	fmt.Println("SET failed, err: %v, %v", err, i1)
	//}
	// HGET 示例
	name, err := _redisCli.HGET("animals", "name")
	age, err := _redisCli.HGET("animals", "age")
	sex, err := _redisCli.HGET("animals", "sex")
	color, err := _redisCli.HGET("animals", "color")
	fmt.Printf("animals: [name:%v], [age: %v], [sex: %v], [color: %v]\n", name, age, sex, color)
	// HGETALL 示例
	animalsMap, err := _redisCli.HGETALL("animals")
	for k, v :=  range animalsMap {
		fmt.Printf("k : %v, v: %v\t", k, v)
	}

	// redis client
	//_redisCli.Start()
	//defer _redisCli.Stop()
	t1 := time.Now().UnixNano() / 1e6
	a1, _ := _redisCli.SCAN("GB212_20*")
	t2 := time.Now().UnixNano() / 1e6
	a2, _ := _redisCli.KEYS("GB212_20*")
	t3 := time.Now().UnixNano() / 1e6
	fmt.Printf("SCAN time: %d\tlen: %d\nKEYS time: %d\tlen: %d\n", t2-t1, len(a1), t3-t2, len(a2))
}
