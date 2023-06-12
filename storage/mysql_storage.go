package storage

type MysqlStorage struct {
	Um userMysqlI
}

func NewMysqlStorage() *MysqlStorage {
	return &MysqlStorage{
		Um: newUserMysql(),
	}
}
