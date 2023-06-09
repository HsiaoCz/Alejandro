package storage

type Mysql_Storage interface{}

type MysqlStorage struct{}

func NewMysqlStorage() *MysqlStorage {
	return &MysqlStorage{}
}
