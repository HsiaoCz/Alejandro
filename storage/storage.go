package storage

type Storage struct {
	*MysqlStorage
	*RedisStorage
}

func NewStorage() *Storage {
	return &Storage{
		MysqlStorage: NewMysqlStorage(),
		RedisStorage: NewRedisStorage(),
	}
}
