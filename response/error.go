package response

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	grm "github.com/gat/necessities/database/gorm"
	rds "github.com/gat/necessities/redis"

	"github.com/gat/necessities/logger"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/gat/necessities/utils"
)

var (
	defaultGetSequence = []string{"redis", "gorm"}
)

type ErrorMessageSource struct {
	redisClient  *rds.ClientRedis
	databaseGorm *gorm.DB
	getSequence  []string
}

type Setup func(s *ErrorMessageSource)

var defaultErrorMessageSource *ErrorMessageSource

func NewErrorMessageSource(setup ...Setup) *ErrorMessageSource {
	errorMessageSource := new(ErrorMessageSource)
	for _, s := range setup {
		s(errorMessageSource)
	}

	if len(errorMessageSource.getSequence) <= 0 {
		errorMessageSource.getSequence = defaultGetSequence
	}

	return errorMessageSource
}

func OverwriteDefaultSource() Setup {
	return func(s *ErrorMessageSource) {
		defaultErrorMessageSource = s
	}
}

func SetGetSequence(sequence ...string) Setup {
	return func(source *ErrorMessageSource) {
		source.getSequence = sequence
	}
}

func AddSourceRedis(redisClient *rds.ClientRedis) Setup {
	return func(source *ErrorMessageSource) {
		source.redisClient = redisClient
	}
}

func AddSourceDatabaseGorm(dbGorm *gorm.DB) Setup {
	return func(source *ErrorMessageSource) {
		source.databaseGorm = dbGorm
	}
}

type Error struct {
	Descriptions  map[string]interface{} `json:"descriptions" gorm:"serializer:json"`
	WhatToDo      map[string]interface{} `json:"what_to_do" gorm:"serializer:json"`
	ID            string                 `json:"id" gorm:"primaryKey;type:varchar(36)"`
	ProblemOwner  string                 `json:"problem_owner" gorm:"type:varchar(100)"`
	SeverityLevel int                    `json:"severity_level"`
}

const (
	module = "error"
)

// GetError gets error data from all sequence that already defined in `ErrorMessageSource` object.
// If data not found in first sequence, it will search in next sequence until it found or last sequence.
//
// err object will be returned if the last sequence error occurs. Otherwise, errorData == nil and err == nil if no error occurs and data is
// not found in sequences.
//
// If `redis` is in sequence and data not found in redis, it will write the data in redis.
func (s *ErrorMessageSource) GetError(id string, locale string, parseTemplateData interface{}) (errorData *Error, err error) {
	logger := logger.NewLogger("")

	if len(locale) <= 0 {
		locale = strings.ToUpper("id")
	} else {
		locale = strings.ToUpper(locale)
	}

	writeRedis := false
	errorData = new(Error)
	found := false
	for i, source := range s.getSequence {
		if found {
			break
		}
		switch source {
		case "redis":
			if s.redisClient == nil {
				break
			}
			getCache, err := s.redisClient.Read(id, module)
			if err == redis.Nil {
				logger.LogWarn("cache not found", err)
				writeRedis = true
			} else if err != nil {
				logger.LogError("read cache", err)
			} else {
				err = json.Unmarshal(getCache, errorData)
				found = true
			}
		case "gorm", "postgre", "psql", "postresql", "mysql":
			if s.databaseGorm == nil {
				break
			}
			errorData.ID = id
			queryResult := s.databaseGorm.First(errorData)
			if queryResult.Error != nil {
				logger.LogError("get error from DB", queryResult.Error)
				// check if this is the last sequence or not
				// if this the last sequence and error data not already found, return error
				if i == len(s.getSequence)-1 && !found {
					return errorData, queryResult.Error
				}
			} else {
				found = true
			}
		}
	}

	if !found {
		errorMessage := fmt.Sprintf("error data with id %s not found in any specified sequence", id)
		return nil, errors.New(errorMessage)
	}

	// write/cache to Redis
	if found && s.redisClient != nil && writeRedis {
		if err := s.redisClient.Write(id, module, errorData); err != nil {
			logger.LogWarn("write error to redis", err)
		}
	}

	// if parseTemplateData not nil, then parse template format
	if parseTemplateData != nil {
		resMsg, err := utils.ParseTemplateToString(errorData.Descriptions[locale].(string), parseTemplateData)
		if err != nil {
			logger.LogError("parse template message", err)
		}
		errorData.Descriptions[locale] = resMsg
		return errorData, err
	} else {
		return errorData, nil
	}
}

// InitErrorTable initiates error message data in database.
// This function doesn't return error, because when error occurred this function immediately
// call fatal (panic).
func InitErrorTable(db *grm.DatabaseGorm) {
	logger := logger.NewLogger("")
	db.AutoMigrate(map[string]interface{}{
		"errors": Error{},
	})

	err := db.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&ErrorMessages).Error

	if err != nil {
		logger.LogPanic("upsert error message data to errors table")
	}
	logger.LogInfo("migrate table and initialize data to errors table")
}
