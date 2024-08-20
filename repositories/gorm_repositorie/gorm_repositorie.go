package gormrepositorie

import (
	playerdomain "ctcwebservice/core/player/domain"
	playersportdomain "ctcwebservice/core/player_sport/domain"
	sportdomain "ctcwebservice/core/sport/domain"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
	}

	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	dbname := os.Getenv("DBNAME")
	password := os.Getenv("PASSWORD")
	sslmode := os.Getenv("SSLMODE")

	dns := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=%s", host, user, dbname, password, sslmode)

	db, _ := gorm.Open(postgres.Open(dns), &gorm.Config{TranslateError: true})

	db.SetupJoinTable(&playerdomain.Player{}, "PlayerSports", &playersportdomain.PlayerSport{})
	db.AutoMigrate(&playerdomain.Player{}, &sportdomain.Sport{}, &playersportdomain.PlayerSport{})

	Db = db
}
