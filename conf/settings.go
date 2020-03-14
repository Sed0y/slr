package conf

import (
	"database/sql"
)

var DB_postgres *sql.DB

const (

	// параметры подключения к БД PostgreSQL

	PostgresUser     = "postgres"
	PostgresPassword = "123456"
	PostgresDB       = "journal"
	PostgresHost     = "127.0.0.1"
	//PostgresHost = "192.168.65.184"
	PostgresPort = 5432

	// *************  НАСТРОЙКИ КЛИЕНТА   ********************************
	// хранение куки
	CookieExpiration = 180 // дней

	// *************  НАСТРОЙКИ ДОСТУПОВ  ********************************
	// логин суперпользователя
	Root = "БочаговКА"

	SolarisService = "https://solaris-inform.com"
	ProxyVSK       = "http://mwgtest.vsk.ru:9090"
	ProxyAuth      = "bochagov:vmmjn28z?W11"
	SolarisToken   = "mJQYQXnIolSwciOV3Km6613iblosnqYCS7t3LvvXYZGUhULVGDBRNxm6yv3jqxhU"
)
