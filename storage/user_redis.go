package storage

type userRedisI interface {
	SubscribeUser(int, int) error
}

type userRedis struct{}

func newUserRedis() *userRedis {
	return &userRedis{}
}

func (u *userRedis) SubscribeUser(userID int, subUserID int) error {
	return nil
}
