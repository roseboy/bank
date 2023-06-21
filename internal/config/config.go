package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Cfg config value
var Cfg config

// config cfg
type config struct {
	MySQL struct {
		DSN string `json:"dsn"`
	} `json:"mysql"`
	Test struct {
		SecretId  string `json:"secretId"`
		SecretKey string `json:"secretKey"`
	} `json:"test"`
}

// InitConfig init
func InitConfig(file string) error {
	jsonStr, err := readText(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(jsonStr), &Cfg)
	if err == nil {
		loadConfig(file)
	}
	return err
}

func loadConfig(file string) {
	ticker := time.NewTicker(10 * time.Second)
	go func(t *time.Ticker) {
		for {
			<-t.C
			jsonStr, _ := readText(file)
			_ = json.Unmarshal([]byte(jsonStr), &Cfg)
		}
	}(ticker)
}

func readText(path string) (string, error) {
	var text = ""

	_, err := os.Stat(path)
	if err != nil {
		return "", err
	}

	f, err := os.Open(path)
	if err != nil {
		return "", nil
	}
	defer func() { _ = f.Close() }()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		text = fmt.Sprintf("%s%s\n", text, line)
	}
	return text, nil
}
