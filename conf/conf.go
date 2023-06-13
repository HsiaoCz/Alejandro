package conf

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(WebApp)

type WebApp struct {
	AC AppConf   `mapstructure:"app"`
	MC MysqlConf `mapstructure:"mysql"`
	RC RedisConf `mapstructure:"redis"`
}

type AppConf struct {
	Addr string `mapstructure:"addr"`
	Port string `mapstructure:"port"`
}

type MysqlConf struct {
	Mysql_User string `mapstructure:"mysql_user"`
	Password   string `mapstructure:"password"`
	Mysql_Host string `mapstructure:"mysql_host"`
	Mysql_Port string `mapstructure:"mysql_port"`
	DB_Name    string `mapstructure:"db_name"`
}

type RedisConf struct {
	Redis_Host   string `mapstructure:"redis_host"`
	Redis_Port   string `mapstructure:"redis_port"`
	Redis_DB     string `mapstructure:"redis_db"`
	Redis_Passwd string `mapstructure:"redis_passwd"`
}

func InitConf() error {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(Conf)
	if err != nil {
		log.Fatal(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("the config is changed", e.Name)
		err = viper.Unmarshal(Conf)
		if err != nil {
			log.Fatal(err)
		}
	})
	return err
}
