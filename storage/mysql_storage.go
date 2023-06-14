package storage

import (
	"alejandro/conf"
	"alejandro/data"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlStorage struct {
	Um userMysqlI
	mc *mysqlConf
}

type mysqlConf struct {
	mysql_user     string
	mysql_password string
	mysql_host     string
	mysql_port     string
	db_name        string
}

func NewMysqlStorage() *MysqlStorage {
	mcy := conf.Conf.MC
	return &MysqlStorage{
		Um: newUserMysql(),
		mc: &mysqlConf{
			mysql_user:     mcy.Mysql_User,
			mysql_password: mcy.Password,
			mysql_host:     mcy.Mysql_Host,
			mysql_port:     mcy.Mysql_Port,
			db_name:        mcy.DB_Name,
		},
	}
}

func (m *MysqlStorage) initStore() error {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.mc.mysql_user, m.mc.mysql_password, m.mc.mysql_host, m.mc.mysql_port, m.mc.db_name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db.AutoMigrate(&data.User{})
	return nil
}
