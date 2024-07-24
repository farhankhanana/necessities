package redis

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gat/necessities/logger"
	"github.com/go-redis/redis"
)

var (
	defaultExpiration = 24 * time.Hour
)

type ClientRedis struct {
	Client      *redis.Client
	ProjectName string
	Expiration  time.Duration
}

// NewRedisClient creates new client object for Redis.
// Need host, port, password and project name as parameters.
// Parameter `expiration` will be set to default one day if `expiration` value is zero
func NewRedisClient(host, port, password, projectName string, expiration time.Duration) (*ClientRedis, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})
	c := &ClientRedis{
		Client:      redisClient,
		ProjectName: projectName,
		Expiration:  expiration,
	}
	err := redisClient.Ping().Err()
	if err != nil {
		return nil, err
	}

	if expiration == 0 {
		c.Expiration = defaultExpiration
	}

	return c, err
}

// Write writes to Redis
func (c *ClientRedis) Write(id, module string, data interface{}) error {
	logger := logger.NewLogger("")

	key := fmt.Sprintf("%s-%s/%s", c.ProjectName, module, id)

	dataByte, _ := json.Marshal(data)
	err := c.Client.Set(key, string(dataByte), time.Duration(c.Expiration)*time.Second).Err()
	if err != nil {
		logger.LogError("redis set value", err)
		return err
	}

	return err
}

// Read reads cached data from Redis
func (c *ClientRedis) Read(id, module string) ([]byte, error) {
	logger := logger.NewLogger("")

	key := fmt.Sprintf("%s-%s/%s", c.ProjectName, module, id)

	val, err := c.Client.Get(key).Bytes()
	if err != nil {
		logger.LogError("redis get value", err)
		return nil, err
	}
	return val, err
}
