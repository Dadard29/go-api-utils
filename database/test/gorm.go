package test

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Profile struct {
	//gorm.Model
	ProfileKey string `gorm:"type:varchar(70);index:profile_key;primary_key"`
	Username string `gorm:"type:varchar(70);index:username"`
	HashedPassword string `gorm:"type:varchar(70);index:hashed_password"`
}

func (Profile) TableName() string {
	return "profile"
}

func TestGorm() {
	username := "root"
	password := "admin"
	dbName := "sample_api"

	dsn := fmt.Sprintf("%s:%s@/%s", username, password, dbName)

	db, _ := gorm.Open("mysql", dsn)

	model := Profile{
		ProfileKey:     "pk3",
		Username:       "username1",
		HashedPassword: "hash1",
	}

	// db.NewRecord(model)
	// db.Create(&model)

	var modelQuery = Profile{
		ProfileKey: "pk",
	}
	db.Find(&modelQuery)
	fmt.Printf("%v\n", model)
	db.Close()
}
