package systemparameter

import (
	"encoding/json"
	"errors"

	grm "github.com/gat/necessities/database/gorm"
	"github.com/gat/necessities/logger"
	rds "github.com/gat/necessities/redis"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type SystemParameter struct {
	ID           string `json:"id" gorm:"varchar(36)"`
	Group        string `json:"group" gorm:"varchar(36)"`
	Value        string `json:"value" gorm:"varchar(200)"`
	Descriptions string `json:"descriptions" gorm:"varchar(200)"`
}

var (
	module = "system_parameter"
)

// GetSystemParameter gets system parameter data from Redis.
// If data not found in Redis, will fetch from database and set to Redis.
func GetSystemParameter(dg *grm.DatabaseGorm, cr *rds.ClientRedis, id string) (*SystemParameter, error) {
	logger := logger.NewLogger("")

	parameterData := new(SystemParameter)
	getCache, err := cr.Read(id, module)
	if err == redis.Nil {
		logger.LogError("cache not found", err)
		parameterData.ID = id
		queryResult := dg.DB.First(parameterData)
		if errors.Is(queryResult.Error, gorm.ErrRecordNotFound) {
			logger.LogError("get error from DB", queryResult.Error)
			return nil, queryResult.Error
		}

		// write/cache to Redis
		if err := cr.Write(id, module, parameterData); err != nil {
			logger.LogError("write error to redis", err)
			return nil, err
		}

		return parameterData, err
	} else if err != nil {
		logger.LogError("read cache", err)
		return nil, err
	} else {
		err = json.Unmarshal(getCache, parameterData)
		return parameterData, err
	}
}

// GetSystemParameters gets system parameters data from Redis.
// If data not found in Redis, will fetch from database and set to Redis.
func GetSystemParameters(dg *grm.DatabaseGorm, cr *rds.ClientRedis, filter SystemParameter) (map[string]string, error) {
	logger := logger.NewLogger("")

	id := filter.ID
	if len(filter.Group) > 0 {
		id = "filter/" + filter.Group
	}

	systemParameters := &[]SystemParameter{}
	getCache, err := cr.Read(id, module)

	if err == redis.Nil {
		logger.LogError("cache not found", err)
		if err := dg.DB.Where(&filter).Find(systemParameters).Error; err != nil {
			logger.LogError("get system parameters", err)
			return nil, err
		}

		// write/cache to Redis
		if err := cr.Write(id, module, systemParameters); err != nil {
			logger.LogError("write error to redis", err)
			return nil, err
		}

		return parseSystemParameter(systemParameters), nil
	} else if err != nil {
		logger.LogError("read cache", err)
		return nil, err
	} else {
		err = json.Unmarshal(getCache, systemParameters)
		return parseSystemParameter(systemParameters), err
	}
}

// parseSystemParameter parses slice (array) of system parameter to map[string]string
func parseSystemParameter(sp *[]SystemParameter) map[string]string {
	result := make(map[string]string)
	for _, v := range *sp {
		result[v.ID] = v.Value
	}
	return result
}
