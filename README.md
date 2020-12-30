# redisgoExample
基于redisgo的redis客户端简单封装和使用示例

## redisgo包
[redisgo](https://github.com/gomodule/redigo)是一款go语言的redis客户端库。
为了简化对redis的操作，可以使用redisgo对[redis常用命令](http://redis.io/commands)进行封装。
首先在Github上面创建一个仓库[redisgoExample](https://github.com/ccf19881030/redisgoExample)
然后git clone将项目克隆到本地，比如说我的阿里云CentOS8服务器下，
```shell
git clone https://github.com/ccf19881030/redisgoExample.git
```
如下图所示：
![克隆redisgoExample](https://img-blog.csdnimg.cn/20201230141858180.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2NjZjE5ODgxMDMw,size_16,color_FFFFFF,t_70)
当然运行go项目的前提是需要安装golang开发环境
进入到`redisgoExample`目录，执行如下命令：
```shell
go mod init ybu.cn/iot
```
使用`go mod init`命令初始化一个`ybu.cn/iot`的自定义包
然后同样是在`redisgoExample`目录下运行`go get`命令安装`redisgo`客户端：
```shell
go get github.com/gomodule/redigo/redis
```
此时目录下会多出`go.mod`和`go.sum`文件，里面包含了`redisgo`包的引入。
`go.mod`文件内容如下所示：
```mod
module ybu.cn/iot

go 1.14

require github.com/gomodule/redigo v1.8.3

```
`go.sum`文件内容如下所示：
```sum
github.com/davecgh/go-spew v1.1.0 h1:ZDRjVQ15GmhC3fiQ8ni8+OwkZQO4DARzQgrnXU1Liz8=
github.com/davecgh/go-spew v1.1.0/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/gomodule/redigo v1.8.3 h1:HR0kYDX2RJZvAup8CsiJwxB4dTCSC0AaUq6S4SiLwUc=
github.com/gomodule/redigo v1.8.3/go.mod h1:P9dn9mFrCBvWhGE1wpxx6fgq7BAeLBk+UUUzlpkBYO0=
github.com/gomodule/redigo/redis v0.0.0-do-not-use h1:J7XIp6Kau0WoyT4JtXHT3Ei0gA1KkSc6bc87j9v9WIo=
github.com/pmezard/go-difflib v1.0.0 h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=
github.com/pmezard/go-difflib v1.0.0/go.mod h1:iKH77koFhYxTK1pcRnkKkqfTogsbg7gZNVY4sRDYZ/4=
github.com/stretchr/objx v0.1.0/go.mod h1:HFkY916IF+rwdDfMAkV7OtwuqBVzrE8GR6GFx+wExME=
github.com/stretchr/testify v1.5.1 h1:nOGnQDM7FYENwehXlg/kFVnos3rEvtKTjRvOWSzb6H4=
github.com/stretchr/testify v1.5.1/go.mod h1:5W2xD1RspED5o8YsWQXVCued0rvSQ+mT+I5cxcmMvtA=
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/yaml.v2 v2.2.2 h1:ZCJp+EgiOT7lHqUV2J862kp8Qj64Jo6az82+3Td9dZw=
gopkg.in/yaml.v2 v2.2.2/go.mod h1:hI93XBmqTisBFMUTm0b8Fm+jr3Dg1NNxqwp+5A1VGuI=
```

## 自定义的common包
在redisgoExample目录下新建一个`common`目录，再创建`array.go`、`define.go`、`interface.go`这三个go文件，用于一些数组、redis配置、redis数据结构的基本操作，
其内容分别如下：
### 1.array.go
```go
package common

// ArrayOf does the array contain specified item
func ArrayOf(arr []string, dest string) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == dest {
			return true
		}
	}
	return false
}

// ArrayDuplice 数组去重
func ArrayDuplice(arr []string) []string {
	var out []string
	tmp := make(map[string]byte)
	for _, v := range arr {
		tmplen := len(tmp)
		tmp[v] = 0
		if len(tmp) != tmplen {
			out = append(out, v)
		}
	}
	return out
}
```

### 2.define.go
```go
package common

// RedisConnOpt connect redis options
type RedisConnOpt struct {
	Enable   bool
	Host     string
	Port     int32
	Password string
	Index    int32
	TTL      int32
}
```

### 3.interface.go
```go
package common

// RedisData 存储数据结构
type RedisData struct {
	Key		string
	Field 	string
	Value 	string
	Expire 	int64
}

// RedisDataArray RedisData of array
type RedisDataArray []*RedisData


// IRedis redis client interface
type IRedis interface {
	// KEYS get patten key array
	KEYS(patten string) ([]string, error)

	// SCAN get patten key array
	SCAN(patten string) ([]string, error)

	// DEL delete k-v
	DEL(key string) (int, error)

	// DELALL delete key array
	DELALL(key []string) (int, error)

	// GET get k-v
	GET(key string) (string, error)

	// SET set k-v
	//SET(key string, value string) (int64, error)

	// SETEX set k-v expire seconds
	SETEX(key string, sec int, value string) (int64, error)

	// EXPIRE set key expire seconds
	EXPIRE(key string, sec int64) (int64, error)

	// HGETALL get map of key
	HGETALL(key string) (map[string]string, error)

	// HGET get value of key-field
	HGET(key string, field string) (string, error)

	// HSET set value of key-field
	//HSET(key string, field string, value string) (int64, error)

	// Write 向redis中写入多组数据
	Write(data RedisDataArray)
}
```

## redisgo的封装
在redisgoExample目录下新建一个cache目录，在此目录下创建一个`redis.go`的文件，主要用于封装常见的redis命令，其内容如下：
```go
package cache

import (
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
	"ybu.cn/iot/common"
)

// https://godoc.org/github.com/gomodule/redigo/redis#pkg-examples
// https://github.com/gomodule/redigo

// RedisClient redis client instance
type RedisClient struct {
	pool    *redis.Pool
	connOpt common.RedisConnOpt
	// 数据接收
	chanRx chan common.RedisDataArray
	// 是否退出
	isExit bool
}

// NewRedis new redis client
func NewRedis(opt common.RedisConnOpt) *RedisClient {
	return &RedisClient{
		connOpt: opt,
		pool:    newPool(opt),
		chanRx:  make(chan common.RedisDataArray, 100),
	}
}

// newPool 线程池
func newPool(opt common.RedisConnOpt) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		// MaxActive:   10,
		// Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", opt.Host, opt.Port))
			if err != nil {
				log.Fatalf("Redis.Dial: %v", err)
				return nil, err
			}
			if _, err := c.Do("AUTH", opt.Password); err != nil {
				c.Close()
				log.Fatalf("Redis.AUTH: %v", err)
				return nil, err
			}
			if _, err := c.Do("SELECT", opt.Index); err != nil {
				c.Close()
				log.Fatalf("Redis.SELECT: %v", err)
				return nil, err
			}
			return c, nil
		},
	}
}

// Start 启动接收任务协程
func (r *RedisClient) Start() {
	r.isExit = false
	// 开启协程用于循环接收数据
	go r.loopRead()
}

// Stop 停止接收任务
func (r *RedisClient) Stop() {
	r.isExit = true
	// 关闭数据接收通道
	close(r.chanRx)
	// 关闭redis线程池
	r.pool.Close()
}

// Write 向redis中写入多组数据
func (r *RedisClient) Write(data common.RedisDataArray) {
	r.chanRx <- data
}

// loopRead 循环接收数据
func (r *RedisClient) loopRead() {
	for !r.isExit {
		select {
		case rx := <-r.chanRx:
			for _, it := range rx {
				if len(it.Key) > 0 {
					if len(it.Field) > 0 {
						if _, err := r.HSET(it.Key, it.Field, it.Value); err != nil {
							log.Printf("[%s, %s, %s]: %s\n", it.Key, it.Field, it.Value, err.Error())
						}
					} else {
						if _, err := r.SET(it.Key, it.Value); err != nil {
							log.Printf("[%s, %s, %s]: %s\n", it.Key, it.Field, it.Value, err.Error())
						}
					}
					if it.Expire > 0 {
						r.EXPIRE(it.Key, it.Expire)
					}
				}
			}
		}
	}

}

// Error get redis connect error
func (r *RedisClient) Error() error {
	conn := r.pool.Get()
	defer conn.Close()
	return conn.Err()
}

// 常用Redis操作命令的封装
// http://redis.io/commands

// KEYS get patten key array
func (r *RedisClient) KEYS(patten string) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.Strings(conn.Do("KEYS", patten))
}

// SCAN 获取大量key
func (r *RedisClient) SCAN(patten string) ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	var out []string
	var cursor uint64 = 0xffffff
	isfirst := true
	for cursor != 0 {
		if isfirst {
			cursor = 0
			isfirst = false
		}
		arr, err := conn.Do("SCAN", cursor, "MATCH", patten, "COUNT", 100)
		if err != nil {
			return out, err
		}
		switch arr := arr.(type) {
		case []interface{}:
			cursor, _ = redis.Uint64(arr[0], nil)
			it, _ := redis.Strings(arr[1], nil)
			out = append(out, it...)
		}
	}
	out = common.ArrayDuplice(out)
	return out, nil
}

// DEL delete k-v
func (r *RedisClient) DEL(key string) (int, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("DEL", key))
}

// DELALL delete key array
func (r *RedisClient) DELALL(key []string) (int, error) {
	conn := r.pool.Get()
	defer conn.Close()
	arr := make([]interface{}, len(key))
	for i, v := range key {
		arr[i] = v
	}
	return redis.Int(conn.Do("DEL", arr...))
}

// GET get k-v
func (r *RedisClient) GET(key string) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.String(conn.Do("GET", key))
}

// SET set k-v
func (r *RedisClient) SET(key string, value string) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("SET", key, value))
}

// SETEX set k-v expire seconds
func (r *RedisClient) SETEX(key string, sec int, value string) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("SETEX", key, sec, value))
}

// EXPIRE set key expire seconds
func (r *RedisClient) EXPIRE(key string, sec int64) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("EXPIRE", key, sec))
}

// HGETALL get map of key
func (r *RedisClient) HGETALL(key string) (map[string]string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.StringMap(conn.Do("HGETALL", key))
}

// HGET get value of key-field
func (r *RedisClient) HGET(key string, field string) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.String(conn.Do("HGET", key, field))
}

// HSET set value of key-field
func (r *RedisClient) HSET(key string, field string, value string) (int64, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("HSET", key, field, value))
}

```

## 测试redis客户端
在redisgoExample目录下新建一个`redisgoExample.go`文件用于测试，
其内容如下：
```go
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
	_redisCli.Start()
	defer _redisCli.Stop()
	t1 := time.Now().UnixNano() / 1e6
	a1, _ := _redisCli.SCAN("GB212_20*")
	t2 := time.Now().UnixNano() / 1e6
	a2, _ := _redisCli.KEYS("GB212_20*")
	t3 := time.Now().UnixNano() / 1e6
	fmt.Printf("SCAN time: %d\tlen: %d\nKEYS time: %d\tlen: %d\n", t2-t1, len(a1), t3-t2, len(a2))
}
```

我是在自己的阿里云服务器上运行的，并且redis密码改成自己阿里云服务器上的redis配置，运行结果如下图所示：
![测试结果](https://img-blog.csdnimg.cn/20201230144712105.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2NjZjE5ODgxMDMw,size_16,color_FFFFFF,t_70)
最终的代码我已经上传至我的github仓库：[https://github.com/ccf19881030/redisgoExample](https://github.com/ccf19881030/redisgoExample)，需要的话可以自取：
```
git clone https://github.com/ccf19881030/redisgoExample.git
```

## 参考资料
- [https://github.com/gomodule/redigo](https://github.com/gomodule/redigo)
- [](https://godoc.org/github.com/gomodule/redigo/redis#pkg-examples)
- [API Reference](http://godoc.org/github.com/gomodule/redigo/redis)
- [FAQ](https://github.com/gomodule/redigo/wiki/FAQ)
- [Examples](https://godoc.org/github.com/gomodule/redigo/redis#pkg-examples)
