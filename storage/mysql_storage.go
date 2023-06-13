package storage

type MysqlStorage struct {
	Um userMysqlI
}

func NewMysqlStorage() *MysqlStorage {
	return &MysqlStorage{
		Um: newUserMysql(),
	}
}

func (m *MysqlStorage) initStore() error {
	return nil
}
