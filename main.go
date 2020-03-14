package main

import (
	"solaris/controllers"
	_ "solaris/routers"

	"github.com/astaxie/beego"

	"database/sql"
	"solaris/conf"
	"solaris/models"

	"fmt"
	"log"
	"time"

	"math/rand"
)

var App models.Application

func init() {

	var err error

	// Инициализация подключения в БД PostgreSQL
	// host=%s port=%d
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.PostgresHost, conf.PostgresPort,
		conf.PostgresUser, conf.PostgresPassword, conf.PostgresDB)

	conf.DB_postgres, err = sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}

	//App.Filials.Load()

	App.Emloyers.Load()
	fmt.Println(App.Emloyers)
	App.Permissions.Init()

	fmt.Println("ok")
	App.Solaris.Init(conf.SolarisService, conf.ProxyVSK, conf.ProxyAuth, conf.SolarisToken)

	controllers.App = &App

}

var letterRunes = []rune("абвгдеёжзиклмнопрстуфхцчшщъыьэюя")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func randate() time.Time {
	min := time.Date(1950, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(1990, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func main() {

	beego.Run()
}
