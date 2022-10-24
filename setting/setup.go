package setting

import (
	"chain-api-imgo/config"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// 待编译注入
var RunEnv string

func Setup(env string) error {
	if env == "" {
		env = RunEnv
	}
	if env == "" {
		env = "dev"
	}
	if env == "dev" {
		log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	}
	configFile := "asserts/" + env + "/config.yaml"

	if stat, err := os.Stat("config.yml"); err == nil && !stat.IsDir() {
		configFile = "config.yml"
	}

	conf, err := config.LoadAndSet(configFile)
	log.Println("load config file:", configFile)
	config.SetConfig(conf)
	return err
}
