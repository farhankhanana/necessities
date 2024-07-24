package databases

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gat/necessities/logger"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	DatabaseHost               string
	DatabasePort               string
	DatabaseUsername           string
	DatabasePassword           string
	DatabaseName               string
	DatabaseSystem             string
	DatabaseMaxIdleConnections int
	DatabaseMaxOpenConnections int
}

type DatabaseGorm struct {
	DB    *gorm.DB
	SqlDB *sql.DB
}

// BaseModel defines base model of gorm.
type BaseModel struct {
	CreatedAt time.Time      `json:"created_at,omitempty" csv:"created_at"`
	UpdatedAt time.Time      `json:"updated_at,omitempty" csv:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index" csv:"-"`
}

// AuditModel defines audit model for CES.
type AuditModel struct {
	BaseModel
	CreatedByID    string `json:"created_by_id,omitempty" gorm:"type:varchar(36)" csv:"created_by_id"`
	CreatedByName  string `json:"created_by_name,omitempty" gorm:"type:varchar(50)" csv:"created_by_name"`
	UpdatedByID    string `json:"updated_by_id,omitempty" gorm:"type:varchar(36)" csv:"updated_by_id"`
	UpdatedByName  string `json:"updated_by_name,omitempty" gorm:"type:varchar(50)" csv:"updated_by_name"`
	RevisionNumber int    `json:"revision_number,omitempty" gorm:"type:smallint" csv:"revision_number"`
}

func NewDatabaseGorm(db *gorm.DB, sqlDB *sql.DB) *DatabaseGorm {
	return &DatabaseGorm{
		DB:    db,
		SqlDB: sqlDB,
	}
}

type Option func(s *DatabaseConfig)

func SetMaxIdleConnections(maxIdleConn int) Option {
	return func(dbConfig *DatabaseConfig) {
		dbConfig.DatabaseMaxIdleConnections = maxIdleConn
	}
}

func SetMaxOpenConnections(maxOpenConn int) Option {
	return func(dbConfig *DatabaseConfig) {
		dbConfig.DatabaseMaxOpenConnections = maxOpenConn
	}
}

// NewDatabaseConfig returns object of database configuration.
// `dbSystem` is current database supported by this library. Choose one of available enums in each supported database.
//
//   - PostgreSQL: enum[`postgre`, `postgres`, `postgresql`, `postgre_sql`, `pg`]
//   - MySQL: enum[`mysql`, `my_sql`]
//
// `opts` is optional parameter. You can set max idle connection and max open connection using
// `SetMaxIdleConnections` and `SetMaxOpenConnections` function.
func NewDatabaseConfig(host, port, username, password, dbName, dbSystem string, opts ...Option) *DatabaseConfig {
	dbConfig := &DatabaseConfig{
		DatabaseHost:     host,
		DatabasePort:     port,
		DatabaseUsername: username,
		DatabasePassword: password,
		DatabaseName:     dbName,
		DatabaseSystem:   dbSystem,
	}

	for _, opt := range opts {
		opt(dbConfig)
	}

	return dbConfig
}

// AuthDatabase authenticates and connects to database.
// This function doesn't return error, because when error occurred this function immediately
// call fatal (panic).
func (dbc *DatabaseConfig) AuthDatabase(config *gorm.Config) *DatabaseGorm {
	logger := logger.NewLogger("")

	var err error
	var db *gorm.DB
	switch dbc.DatabaseSystem {
	case "postgre", "postgres", "postgresql", "postgre_sql", "pg":
		logger.LogInfo("init postgre sql")
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
			dbc.DatabaseHost,
			dbc.DatabaseUsername,
			dbc.DatabasePassword,
			dbc.DatabaseName,
			dbc.DatabasePort,
		)
		db, err = gorm.Open(postgres.Open(dsn), config)
	case "mysql", "my_sql":
		logger.LogInfo("init mysql sql")
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=Local",
			dbc.DatabaseUsername,
			dbc.DatabasePassword,
			dbc.DatabaseHost,
			dbc.DatabasePort,
			dbc.DatabaseName,
		)
		db, err = gorm.Open(mysql.Open(dsn), config)
	}
	if err != nil {
		logger.LogPanic("gorm connect to database", err)
	}

	logger.LogInfo(fmt.Sprintf("connected to %s", dbc.DatabaseSystem))

	var sqlDB *sql.DB
	sqlDB, _ = db.DB()
	if dbc.DatabaseMaxIdleConnections > 0 {
		sqlDB.SetMaxIdleConns(dbc.DatabaseMaxIdleConnections)
	}
	if dbc.DatabaseMaxOpenConnections > 0 {
		sqlDB.SetMaxOpenConns(dbc.DatabaseMaxOpenConnections)
	}

	return NewDatabaseGorm(db, sqlDB)
}

// AutoMigrate migrates, creates, and updates (un)existing tables.
// This function doesn't return error, because when error occurred this function immediately
// call fatal (panic).
func (dg *DatabaseGorm) AutoMigrate(models map[string]interface{}) {
	logger := logger.NewLogger("")

	for k := range models {
		err := dg.DB.AutoMigrate(models[k])
		if err != nil {
			logger.LogPanic(fmt.Sprintf("auto migrate fail for model %s", k), err)
		}
	}
	logger.LogInfo("auto migrate finish")
}
