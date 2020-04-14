package conf

import "github.com/omigo/log"

type conf struct {
	Name     string
	LogLevel string

	DB struct {
		Dialect  string
		Host     string
		Port     int
		Username string
		Password string
		DBName   string
		Params   string
	}
}

var Conf conf

func init() {
	Conf.LogLevel = "debug"

	Conf.DB.Dialect = "mysql"
	Conf.DB.Host = "127.0.0.1"
	Conf.DB.Port = 3306
	Conf.DB.Username = "test"
	Conf.DB.Password = "123456"
	Conf.DB.DBName = "test"
	Conf.DB.Params = "charset=utf8&parseTime=true&loc=Local"

	log.SetLevelString(Conf.LogLevel)
	// log.Json(Conf)
}
