package practice

import (
	"fmt"
	"sync"
	"time"
)

// sync.Once confirm that command will run only once time whole the life cycle

type Manager struct {
	Name string
}

var managerInstance *Manager
var managerOnce sync.Once

func GetManager() *Manager {
	managerOnce.Do(func() {
		managerInstance = &Manager{
			Name: "Manager Once",
		}
	})
	return managerInstance
}

// practice

type Database struct {
	Config *DBConfig
}

var dbInstance *Database
var dbOnce sync.Once

func GetDatabase() *Database {
	dbOnce.Do(func() {
		fmt.Println("Database is config... pls wait")
		dbInstance = &Database{
			Config: NewDBConfig("database config", 5, WithMaxIdle(5), WithTimeOut(5*time.Second)),
		}
	})
	return dbInstance
}
