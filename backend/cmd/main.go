package main

import (
	"encoding/json"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"main/api"
	"main/models"
	"main/models/db"
	"main/providers"
	"main/services"
	"os"
)

var (
	DB   *gorm.DB
	conf map[string]string
)

func init() {
	conf = readConfig()
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	log.Println("Preparing database...")

	DB.AutoMigrate(&db.Label{})
	DB.AutoMigrate(&db.Ticket{})
	DB.AutoMigrate(&db.User{})
	DB.AutoMigrate(&db.ServicePoint{})
	DB.AutoMigrate(&db.Branch{})
}

func main() {
	sess := services.NewSessionBasicService()
	apiConf := &api.API{
		AccountProvider: providers.NewAccountProvider(DB, sess),
		TicketProvider:  providers.NewTicketProvider(DB),
		OrgProvider:     providers.NewOrgProvider(DB),
		Config:          conf,
		Host:            conf["app.addr"],
	}

	// Prepopulate admins
	if DB.Model(&db.User{}).Where(&db.User{Username: conf["adm.name"]}).First(&db.User{}).Error != nil {
		log.Println("Creating admin...")

		acc := apiConf.AccountProvider.New()
		uid, err := acc.Register(conf["adm.name"], conf["adm.pass"])
		if err != nil {
			panic(err)
		}

		role := &db.Role{
			Name:       "Admin",
			Meta:       make(models.JSONMap),
			Privileges: models.JSONMap{"admin": 1},
		}
		if DB.Model(&db.Role{}).Create(&role).Error != nil {
			panic("Failed to create admin role")
		}
		u := db.User{}
		u.ID = uid
		DB.Model(&u).Updates(&db.User{RoleID: role.ID})
	}

	log.Println(api.StartWS(apiConf))
}

func readConfig() map[string]string {
	conf := make(map[string]string)
	d, err := os.ReadFile("config.json")
	if err != nil {
		return defaultConfig()
	}
	err = json.Unmarshal(d, &conf)
	if err != nil {
		return defaultConfig()
	}
	return conf
}

func defaultConfig() map[string]string {
	return map[string]string{
		"org.name": "Билеты в Бобруйск",
		"org.desc": "Это надолго",
		"app.addr": ":5000",
		"adm.name": "admin",
		"adm.pass": "Adminster0",
	}
}
