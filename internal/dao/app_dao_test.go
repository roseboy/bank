package dao

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestCreateApp(t *testing.T) {
	err := InitMySQL()
	if err != nil {
		panic(err)
	}
	var maxAppId uint64
	err = MySQL.Model(&App{}).Select("ifnull(max(app_id),125000000)").Scan(&maxAppId).Error
	if err != nil {
		panic(err)
	}

	secretId := fmt.Sprintf("AKID%s", randChars())
	secretKey := randChars()
	app := App{AppId: maxAppId + 1, SecretId: secretId, SecretKey: secretKey}

	MySQL.Create(&app)

	data, _ := json.Marshal(app)
	fmt.Println(string(data))

}

func randChars() string {
	rand.Seed(time.Now().UnixNano())
	var chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	str := ""
	for i := 0; i < 32; i++ {
		n := rand.Intn(len(chars))
		str = fmt.Sprintf("%s%s", str, chars[n:n+1])
	}
	return str
}
