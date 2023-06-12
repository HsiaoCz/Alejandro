package conf

var Conf = new(WebApp)

type WebApp struct {
	AC AppConf
	MC MysqlConf
	RC RedisConf
}

type AppConf struct{}

type MysqlConf struct{}

type RedisConf struct{}

func InitConf() error {
	return nil
}
