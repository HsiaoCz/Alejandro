package storage

type Redis_Storage interface{}

type RedisStorage struct{}

func NewRedisStorage() *RedisStorage {
	return &RedisStorage{}
}
