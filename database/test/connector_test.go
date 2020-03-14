package test

import (
	"fmt"
	"github.com/Dadard29/go-api-utils/database"
	"testing"
)

type ProfileInfos struct {
	//gorm.Model
	ProfileKey     string `gorm:"type:varchar(70);index:profile_key;primary_key"`
	Username       string `gorm:"type:varchar(70);index:username"`
	HashedPassword string `gorm:"type:varchar(70);index:hashed_password"`
}

func (ProfileInfos) TableName() string {
	return "profile"
}

func TestNewConnector(t *testing.T) {
	config := map[string]string{
		"usernameKey": "username",
		"passwordKey": "password",
		"database":    "sample_api",
		"host":        "127.0.0.1",
		"port":        "3306",
	}
	var profile ProfileInfos

	connector := database.NewConnector(config, true, []interface{}{&profile})

	var p []ProfileInfos

	var f ProfileInfos
	connector.Orm.Find(&p).First(&f)
	fmt.Println(p)
	fmt.Println(f)
}
