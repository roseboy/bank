package main

import (
	"fmt"
	act "github.com/roseboy/bank-server/internal/action"
	"github.com/roseboy/bank-server/internal/config"
	"github.com/roseboy/bank-server/internal/dao"
	"github.com/roseboy/go-ng/ng"
	"github.com/roseboy/go-ng/plugin"
	"github.com/roseboy/go-ng/plugin/action"
	"time"
)

func main() {
	err := config.InitConfig("./bank.json")
	panicErr(err)

	err = dao.InitMySQL()
	panicErr(err)

	actionPlg := plugin.NewActionPlugin("/api", true, dao.GetAuthInfo)
	actionPlg.RegisterAction("OpenAccount", act.OpenAccount,
		&act.OpenAccountRequest{}, &act.OpenAccountResponse{})

	go TestApi()

	err = ng.NewServer(18080).RegisterPlugins(actionPlg).Start()
	panicErr(err)

}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func TestApi() {
	secretId := config.Cfg.Test.SecretId
	secretKey := config.Cfg.Test.SecretKey
	timestamp := time.Now().Unix() + 5
	body := fmt.Sprintf(`{"Action":"OpenAccount","Timestamp":%d,"AppId":111,"BankType":"CCB"}`, timestamp)
	sign := action.CalcSignature(&action.CalcSignatureArgs{
		Service:   "/api",
		Timestamp: timestamp,
		Method:    "POST",
		Host:      "localhost:18080",
		URI:       "/api",
		Payload:   body,
		SecretKey: secretKey,
	})
	authorization := fmt.Sprintf("%s;%s", secretId, sign)

	fmt.Println("execute the following command in 10s:")
	fmt.Printf("curl localhost:18080/api -d'%s' -H'Authorization:%s'\n", body, authorization)
	fmt.Println()
}
