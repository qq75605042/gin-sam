package datasource

import (
	"fmt"
	"gin-sam/conf"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var masterInstance *gorm.DB

var dbLock sync.Mutex

func InstanceDb() *gorm.DB {
	if masterInstance != nil {
		return masterInstance
	}
	dbLock.Lock()
	defer dbLock.Unlock()
	if masterInstance != nil {
		return masterInstance
	}
	return NewDaMaster()
}

func NewDaMaster() *gorm.DB {
	sourceName := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", conf.DbMaster.User, conf.DbMaster.Pwd, conf.DbMaster.Host, conf.DbMaster.Port, conf.DbMaster.Database)

	db, err := gorm.Open(conf.DRIVER, sourceName)
	if err != nil {
		log.Fatal("db init failed,err:", err)
		return nil
	}
	db.LogMode(true)
	masterInstance = db
	return db
}
