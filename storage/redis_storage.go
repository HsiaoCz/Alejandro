package storage

type RedisStorage struct {
	Ur userRedisI
}

func NewRedisStorage() *RedisStorage {
	return &RedisStorage{
		Ur: newUserRedis(),
	}
}

func (r *RedisStorage) initStore() error {
	return nil
}
