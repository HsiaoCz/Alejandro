package storage

type Storage struct {
	Ms *MysqlStorage
	Rs *RedisStorage
}

func NewStorage() *Storage {
	return &Storage{
		Ms: NewMysqlStorage(),
		Rs: NewRedisStorage(),
	}
}
