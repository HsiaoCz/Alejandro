package storage

type RedisStorage struct{}

func NewRedisStorage() *RedisStorage {
	return &RedisStorage{}
}
