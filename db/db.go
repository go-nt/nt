package db

import (
	"errors"
	"github.com/go-nt/nt/db/driver"
)

var configs map[string]map[string]any
var dbs map[string]Driver

func SetConfig(name string, config map[string]any) {
	if configs == nil {
		configs = make(map[string]map[string]any)
	}
	configs[name] = config
}

func GetConfig(name string) (map[string]any, error) {
	config, ok := configs[name]
	if ok {
		return config, nil
	}

	return nil, errors.New("config (" + name + ") not found")
}

func GetDb(name string) (Driver, error) {
	db, ok := dbs[name]
	if ok {
		return db, nil
	}

	config, ok := configs[name]
	if ok {
		db, err := GetDbByConfig(config)
		if err != nil {
			return nil, err
		}

		if dbs == nil {
			dbs = make(map[string]Driver)
		}

		dbs[name] = db
		return db, nil
	}

	return nil, errors.New("no available db")
}

func GetDbByConfig(config map[string]any) (Driver, error) {
	driverName, ok := config["driver"]
	if !ok {
		return nil, errors.New("db driver type is missing")
	}

	var db Driver
	switch t := driverName.(type) {
	case string:
		switch t {
		case "mysql":
			db = new(driver.Mysql)
		default:
			return nil, errors.New("unsupported driver type: " + t)
		}
	}

	db.Config(config)
	err := db.Init()
	if err != nil {
		return nil, err
	}

	return db, nil
}
