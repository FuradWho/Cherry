package mysql

import (
	"fmt"
	"github.com/FuradWho/Cherry/internal/apiserver/store"
	"gorm.io/gorm"
	"sync"
)

type datastore struct {
	db *gorm.DB

	// can include two database instance if needed
	// docker *grom.DB
	// db *gorm.DB
}

func (ds *datastore) Employees() store.EmployeeStore {
	return newUsers(ds)
}

func (ds *datastore) Close() error {
	if ds.db == nil {
		return nil
	}

	db, err := ds.db.DB()
	if err != nil {
		return err
	}

	return db.Close()
}

var (
	mysqlFactory store.Factory
	once         sync.Once
)

// GetMySQLFactoryOr create mysql factory with the given config.
func GetMySQLFactoryOr() (store.Factory, error) {
	var err error
	//var dbIns *gorm.DB
	//once.Do(func() {
	//	options := &db.Options{
	//		Host:                  viper.GetString("db.host"),
	//		Username:              viper.GetString("db.username"),
	//		Password:              viper.GetString("db.password"),
	//		Database:              viper.GetString("db.database"),
	//		MaxIdleConnections:    viper.GetInt("db.max-idle-connections"),
	//		MaxOpenConnections:    viper.GetInt("db.max-open-connections"),
	//		MaxConnectionLifeTime: viper.GetDuration("db.max-connection-life-time"),
	//		LogLevel:              viper.GetInt("db.log-level"),
	//		Logger:                logger.New(viper.GetInt("db.log-level")),
	//	}
	//	dbIns, err = db.New(options)
	//
	//	// uncomment the following line if you need auto migration the given models
	//	// not suggested in production environment.
	//	// migrateDatabase(dbIns)
	//
	//	mysqlFactory = &datastore{dbIns}
	//})

	if mysqlFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql store fatory, mysqlFactory: %+v, error: %w", mysqlFactory, err)
	}

	return mysqlFactory, nil
}
